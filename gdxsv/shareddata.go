package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"go.uber.org/zap"
	"sync"
	"time"
)

// sharing temporary data between lbs and mcs

var sharedData struct {
	sync.Mutex
	battleUsers map[string]McsUser // session_id -> user info
	battleGames map[string]McsGame // battle_code -> game info

	lbsStatusCacheTime time.Time
	lbsStatusCache     []byte
}

func init() {
	sharedData.battleUsers = map[string]McsUser{}
	sharedData.battleGames = map[string]McsGame{}
	go func() {
		for {
			removeOldSharedData()
			time.Sleep(time.Minute)
		}
	}()
}

const (
	McsGameStateCreated = 0
	McsGameStateOpened  = 1
	McsGameStateClosed  = 2

	McsUserStateCreated = 0
	McsUserStateJoined  = 1
	McsUserStateLeft    = 2
)

type McsUser struct {
	BattleCode  string `json:"battle_code,omitempty"`
	McsRegion   string `json:"mcs_region,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	Name        string `json:"name,omitempty"`
	PilotName   string `json:"pilot_name,omitempty"`
	GameParam   []byte `json:"game_param,omitempty"`
	BattleCount int    `json:"battle_count,omitempty"`
	WinCount    int    `json:"win_count,omitempty"`
	LoseCount   int    `json:"lose_count,omitempty"`
	Side        uint16 `json:"side,omitempty"`
	SessionID   string `json:"session_id,omitempty"`

	State     int       `json:"state,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type McsGame struct {
	BattleCode string `json:"battle_code,omitempty"`
	McsAddr    string `json:"mcs_addr,omitempty"`
	GameDisk   int    `json:"game_disk"`
	Rule       Rule   `json:"rule,omitempty"`

	State     int       `json:"state,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type McsStatus struct {
	Region     string    `json:"region,omitempty"`
	PublicAddr string    `json:"public_addr,omitempty"`
	Users      []McsUser `json:"users,omitempty"`
	Games      []McsGame `json:"games,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

type LbsStatus struct {
	Users []McsUser `json:"users,omitempty"`
	Games []McsGame `json:"games,omitempty"`
}

func ShareMcsGame(g McsGame) {
	sharedData.Lock()
	defer sharedData.Unlock()
	sharedData.battleGames[g.BattleCode] = g
}

func ShareUserWhoIsGoingToBattle(u McsUser) {
	sharedData.Lock()
	defer sharedData.Unlock()
	sharedData.battleUsers[u.SessionID] = u
}

func SyncSharedDataMcsToLbs(status *McsStatus) {
	sharedData.Lock()
	defer sharedData.Unlock()

	for _, u := range status.Users {
		_, ok := sharedData.battleUsers[u.SessionID]
		if ok {
			u.UpdatedAt = status.UpdatedAt
			sharedData.battleUsers[u.SessionID] = u

			if u.State == McsUserStateLeft {
				delete(sharedData.battleUsers, u.SessionID)
				logger.Info("remove mcs user", zap.String("session_id", u.SessionID))
			}
		}
	}

	closedBattleCodes := map[string]bool{}
	for _, g := range status.Games {
		_, ok := sharedData.battleGames[g.BattleCode]
		if ok {
			g.UpdatedAt = status.UpdatedAt
			sharedData.battleGames[g.BattleCode] = g

			if g.State == McsGameStateClosed {
				closedBattleCodes[g.BattleCode] = true
				delete(sharedData.battleGames, g.BattleCode)
				logger.Info("remove mcs game", zap.String("battle_code", g.BattleCode))
			}
		}
	}

	for k, u := range sharedData.battleUsers {
		if closedBattleCodes[u.BattleCode] {
			delete(sharedData.battleUsers, k)
		}
	}
}

func SyncSharedDataLbsToMcs(status *LbsStatus) {
	sharedData.Lock()
	defer sharedData.Unlock()

	activeBattleCodes := map[string]bool{}
	activeSessionIDs := map[string]bool{}

	for _, g := range status.Games {
		activeBattleCodes[g.BattleCode] = true

		h, ok := sharedData.battleGames[g.BattleCode]
		if ok {
			continue
		}
		if h.McsAddr == conf.BattlePublicAddr {
			sharedData.battleGames[g.BattleCode] = g
		}
	}

	for _, u := range status.Users {
		activeSessionIDs[u.SessionID] = true

		v, ok := sharedData.battleUsers[u.SessionID]
		if ok {
			continue
		}
		_, ok = sharedData.battleGames[v.BattleCode]
		if ok {
			sharedData.battleUsers[u.SessionID] = u
		}
	}

	for k, g := range sharedData.battleGames {
		if !activeBattleCodes[g.BattleCode] {
			delete(sharedData.battleGames, k)
		}
	}

	for k, u := range sharedData.battleUsers {
		if !activeSessionIDs[u.SessionID] {
			delete(sharedData.battleUsers, k)
		}
	}
}

func GetMcsUsers() []McsUser {
	sharedData.Lock()
	defer sharedData.Unlock()

	var ret []McsUser

	for _, u := range sharedData.battleUsers {
		ret = append(ret, u)
	}

	return ret
}

func GetSerializedLbsStatus() []byte {
	sharedData.Lock()
	defer sharedData.Unlock()

	if 1 <= time.Since(sharedData.lbsStatusCacheTime).Seconds() {
		st := new(LbsStatus)
		for _, u := range sharedData.battleUsers {
			st.Users = append(st.Users, u)
		}

		for _, g := range sharedData.battleGames {
			st.Games = append(st.Games, g)
		}

		var buf bytes.Buffer
		zw := gzip.NewWriter(&buf)
		jw := json.NewEncoder(zw)

		err := jw.Encode(st)
		if err != nil {
			logger.Error("jw.Encode", zap.Error(err))
			return nil
		}

		err = zw.Close()
		if err != nil {
			logger.Error("zw.Close", zap.Error(err))
			return nil
		}

		if (1 << 16) <= buf.Len() {
			logger.Error("too large data", zap.Int("size", buf.Len()))
			return nil
		}

		sharedData.lbsStatusCache = buf.Bytes()
		sharedData.lbsStatusCacheTime = time.Now()
	}

	return sharedData.lbsStatusCache
}

func NotifyLatestLbsStatus(mcs *LbsPeer) {
	mcs.SendMessage(NewServerNotice(lbsExtSyncSharedData).Writer().WriteBytes(GetSerializedLbsStatus()).Msg())
}

func getBattleGameInfo(battleCode string) (McsGame, bool) {
	sharedData.Lock()
	defer sharedData.Unlock()
	g, ok := sharedData.battleGames[battleCode]
	return g, ok
}

func getBattleUserInfo(sessionID string) (McsUser, bool) {
	sharedData.Lock()
	defer sharedData.Unlock()
	u, ok := sharedData.battleUsers[sessionID]
	return u, ok
}

func updateMcsGameState(battleCode string, newState int) {
	sharedData.Lock()
	defer sharedData.Unlock()
	g := sharedData.battleGames[battleCode]
	if g.State < newState {
		logger.Info("updateMcsGameState",
			zap.String("battle_code", battleCode),
			zap.Int("from", g.State),
			zap.Int("to", newState))
		g.State = newState
		sharedData.battleGames[battleCode] = g
	}
}

func updateMcsUserState(sessionID string, newState int) {
	sharedData.Lock()
	defer sharedData.Unlock()
	u := sharedData.battleUsers[sessionID]
	if u.State < newState {
		logger.Info("updateMcsUserState",
			zap.String("session_id", sessionID),
			zap.Int("from", u.State),
			zap.Int("to", newState))
		u.State = newState
		sharedData.battleUsers[sessionID] = u
	}
}

func removeBattleGameInfo(battleCode string) {
	sharedData.Lock()
	defer sharedData.Unlock()
	delete(sharedData.battleGames, battleCode)
	logger.Info("remove mcs game", zap.String("battle_code", battleCode))
}

func removeBattleUserInfo(battleCode string) {
	sharedData.Lock()
	defer sharedData.Unlock()
	for key, u := range sharedData.battleUsers {
		if u.BattleCode == battleCode {
			delete(sharedData.battleUsers, key)
			logger.Info("remove mcs user", zap.String("session_id", u.SessionID))
		}
	}
}

func removeOldSharedData() {
	sharedData.Lock()
	defer sharedData.Unlock()

	for key, u := range sharedData.battleUsers {
		if 1.0 <= time.Since(u.UpdatedAt).Hours() {
			delete(sharedData.battleUsers, key)
			logger.Warn("remove old zombie battle user", zap.String("session_id", key))
		}
	}

	for key, g := range sharedData.battleGames {
		if 1.0 <= time.Since(g.UpdatedAt).Hours() {
			delete(sharedData.battleGames, key)
			logger.Warn("remove old zombie game", zap.String("battle_code", key))
		}
	}
}
