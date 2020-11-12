package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net"
	"strings"
	"sync"
	"time"
)

const (
	maxLobbyCount = 22
	maxRoomCount  = 5
)

const (
	GameDiskUnk = 0 // Unknown
	GameDiskDC1 = 1 // Dreamcast
	GameDiskDC2 = 2 // Dreamcast DX
	GameDiskPS2 = 3 // PS2 DX
)

const (
	TeamNone  = 0
	TeamRenpo = 1
	TeamZeon  = 2
)

type Lbs struct {
	handlers  map[CmdID]LbsHandler
	userPeers map[string]*LbsPeer
	mcsPeers  map[string]*LbsPeer
	lobbies   map[byte]map[uint16]*LbsLobby
	chEvent   chan interface{}
	chQuit    chan interface{}
}

func NewLbs() *Lbs {
	app := &Lbs{
		handlers:  defaultLbsHandlers,
		userPeers: make(map[string]*LbsPeer),
		mcsPeers:  make(map[string]*LbsPeer),
		lobbies:   make(map[byte]map[uint16]*LbsLobby),
		chEvent:   make(chan interface{}, 64),
		chQuit:    make(chan interface{}),
	}

	app.lobbies[GameDiskPS2] = make(map[uint16]*LbsLobby)
	app.lobbies[GameDiskDC1] = make(map[uint16]*LbsLobby)
	app.lobbies[GameDiskDC2] = make(map[uint16]*LbsLobby)

	for i := 1; i <= maxLobbyCount; i++ {
		app.lobbies[GameDiskPS2][uint16(i)] = NewLobby(app, GameDiskPS2, uint16(i))
	}
	for i := 1; i <= maxLobbyCount; i++ {
		app.lobbies[GameDiskDC1][uint16(i)] = NewLobby(app, GameDiskDC1, uint16(i))
	}
	for i := 1; i <= maxLobbyCount; i++ {
		app.lobbies[GameDiskDC2][uint16(i)] = NewLobby(app, GameDiskDC2, uint16(i))
	}

	return app
}

func (lbs *Lbs) GetLobby(platform uint8, lobbyID uint16) *LbsLobby {
	lobbies, ok := lbs.lobbies[platform]
	if !ok {
		return nil
	}

	lobby, ok := lobbies[lobbyID]
	if !ok {
		return nil
	}

	return lobby
}

func (lbs *Lbs) ListenAndServe(addr string) error {
	logger.Info("lbs.ListenAndServe", zap.String("addr", addr))

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}
	go lbs.eventLoop()
	if len(conf.DiscordLiveStatusWebhookURL) > 0 {
		go lbs.publishLiveStatusToDiscordLoop()
	}
	for {
		tcpConn, err := listener.AcceptTCP()
		if err != nil {
			logger.Error("failed to accept", zap.Error(err))
			continue
		}
		logger.Info("a new connection open", zap.String("addr", tcpConn.RemoteAddr().String()))
		peer := lbs.NewPeer(tcpConn)
		go peer.serve()
	}
}

func (lbs *Lbs) NewPeer(conn *net.TCPConn) *LbsPeer {
	return &LbsPeer{
		app:          lbs,
		conn:         conn,
		PlatformInfo: map[string]string{},
		chWrite:      make(chan bool, 1),
		chDispatch:   make(chan bool, 1),
		outbuf:       make([]byte, 0, 1024),
		inbuf:        make([]byte, 0, 1024),
		logger:       logger.With(zap.String("addr", conn.RemoteAddr().String())),
	}
}

func (lbs *Lbs) FindMcs(region string) *McsStatus {
	for _, p := range lbs.mcsPeers {
		if p.mcsStatus != nil {
			if strings.HasPrefix(p.mcsStatus.Region, region) &&
				p.mcsStatus.PublicAddr != "" {
				return p.mcsStatus
			}
		}
	}
	return nil
}

func (lbs *Lbs) FindPeer(userID string) *LbsPeer {
	p, ok := lbs.userPeers[userID]
	if !ok {
		return nil
	}
	return p
}

func (lbs *Lbs) FindMcsPeer(mcsAddr string) *LbsPeer {
	p, ok := lbs.mcsPeers[mcsAddr]
	if !ok {
		return nil
	}
	return p
}

func (lbs *Lbs) Locked(f func(*Lbs)) {
	c := make(chan interface{})
	lbs.chEvent <- eventFunc{
		f: f,
		c: c,
	}
	<-c
}

func (lbs *Lbs) Quit() {
	lbs.Locked(func(app *Lbs) {
		for _, p := range app.userPeers {
			SendServerShutDown(p)
		}
	})
	time.Sleep(200 * time.Millisecond)
	close(lbs.chQuit)
}

func stripHost(addr string) string {
	_, port, err := net.SplitHostPort(addr)
	if err != nil {
		logger.DPanic("failed to split host port", zap.Error(err))
	}
	return ":" + fmt.Sprint(port)
}

type eventPeerCome struct {
	peer *LbsPeer
}

type eventPeerLeave struct {
	peer *LbsPeer
}

type eventPeerMessage struct {
	peer *LbsPeer
	msg  *LbsMessage
}

type eventFunc struct {
	f func(*Lbs)
	c chan<- interface{}
}

func (lbs *Lbs) cleanPeer(p *LbsPeer) {
	if p.UserID != "" {
		if p.Room != nil {
			p.Room.Exit(p.UserID)
			lbs.BroadcastRoomState(p.Room)
			p.Room = nil
		}
		if p.Lobby != nil {
			p.Lobby.Exit(p.UserID)
			lbs.BroadcastLobbyUserCount(p.Lobby)
			lbs.BroadcastLobbyMatchEntryUserCount(p.Lobby)
			p.Lobby = nil
		}
		delete(lbs.userPeers, p.UserID)
	}

	if p.mcsStatus != nil {
		if len(p.mcsStatus.Games) != 0 || len(p.mcsStatus.Users) != 0 {
			logger.Warn("mcs closed during game",
				zap.Any("games", p.mcsStatus.Games), zap.Any("users", p.mcsStatus.Users))
			for _, g := range p.mcsStatus.Games {
				sharedData.RemoveBattleGameInfo(g.BattleCode)
				sharedData.RemoveBattleUserInfo(g.BattleCode)
			}
		}
		delete(p.app.mcsPeers, p.mcsStatus.PublicAddr)
		p.mcsStatus = nil
	}

	p.conn.Close()
	p.left = true
	discordLiveStatusUpdateAvailable = true
}

func (lbs *Lbs) eventLoop() {
	tick := time.Tick(1 * time.Second)
	peers := map[string]*LbsPeer{}

	for {
		select {
		case <-lbs.chQuit:
			return
		case e := <-lbs.chEvent:
			switch args := e.(type) {
			case eventPeerCome:
				args.peer.logger.Info("eventPeerCome")
				args.peer.lastRecvTime = time.Now()
				peers[args.peer.Address()] = args.peer
				StartLoginFlow(args.peer)
			case eventPeerMessage:
				args.peer.logger.Info("eventPeerMessage", zap.Any("msg", args.msg))
				if args.peer.left {
					args.peer.logger.Warn("got message after left", zap.Any("msg", args.msg))
					continue
				}

				args.peer.lastRecvTime = time.Now()
				if f, ok := lbs.handlers[args.msg.Command]; ok {
					f(args.peer, args.msg)
				} else {
					logger.Warn("handler not found",
						zap.String("cmd", args.msg.Command.String()),
						zap.String("cmd_id", fmt.Sprintf("0x%04x", uint16(args.msg.Command))),
						zap.String("msg", args.msg.String()),
						zap.Binary("body", args.msg.Body),
					)
					if args.msg.Category == CategoryQuestion {
						args.peer.SendMessage(NewServerAnswer(args.msg))
					}
				}
			case eventPeerLeave:
				args.peer.logger.Info("eventPeerLeave")
				lbs.cleanPeer(args.peer)
				delete(peers, args.peer.Address())
			case eventFunc:
				args.f(lbs)
				args.c <- struct{}{}
			}
		case <-tick:
			for _, p := range peers {
				lastRecvSince := time.Since(p.lastRecvTime)
				if 1 <= lastRecvSince.Minutes() {
					logger.Info("kick peer", zap.String("addr", p.Address()))
					lbs.cleanPeer(p)
					delete(peers, p.Address())
				} else if 10 <= lastRecvSince.Seconds() {
					RequestLineCheck(p)
				}
			}
			for _, pfLobbies := range lbs.lobbies {
				for _, lobby := range pfLobbies {
					lobby.CheckLobbyBattleStart()
					lobby.CheckRoomBattleStart()
				}
			}
		}
	}
}

func (lbs *Lbs) BroadcastLobbyUserCount(lobby *LbsLobby) {
	if lobby == nil {
		return
	}
	// For lobby select scene.
	ps2msg := NewServerNotice(lbsPlazaJoin).Writer().
		Write16(lobby.ID).Write16(uint16(len(lobby.Users))).Msg()

	dclobby1 := lbs.GetLobby(GameDiskDC1, lobby.ID)
	dclobby2 := lbs.GetLobby(GameDiskDC2, lobby.ID)
	dcmsg := NewServerNotice(lbsPlazaJoin).Writer().
		Write16(lobby.ID).
		Write16(uint16(len(dclobby1.Users))).
		Write16(uint16(len(dclobby2.Users))).Msg()

	for _, u := range lbs.userPeers {
		if u.IsPS2() {
			u.SendMessage(ps2msg)
		} else if u.IsDC() {
			u.SendMessage(dcmsg)
		}
	}

	// For lobby chat scene.
	if lobby.GameDisk == GameDiskPS2 {
		renpo, zeon := lobby.GetUserCountBySide()
		msgSum1 := NewServerNotice(lbsLobbyJoin).Writer().Write16(TeamRenpo).Write16(renpo + zeon).Msg()
		msgSum2 := NewServerNotice(lbsLobbyJoin).Writer().Write16(TeamZeon).Write16(renpo + zeon).Msg()
		msgRenpo := NewServerNotice(lbsLobbyJoin).Writer().Write16(TeamRenpo).Write16(renpo).Msg()
		msgZeon := NewServerNotice(lbsLobbyJoin).Writer().Write16(TeamZeon).Write16(zeon).Msg()
		for userID := range lobby.Users {
			if p := lbs.FindPeer(userID); p != nil {
				if p.InLobbyChat() {
					p.SendMessage(msgSum1)
					p.SendMessage(msgSum2)
				} else {
					p.SendMessage(msgRenpo)
					p.SendMessage(msgZeon)
				}
			}
		}
	} else if lobby.GameDisk == GameDiskDC1 || lobby.GameDisk == GameDiskDC2 {
		lobby1 := lbs.GetLobby(GameDiskDC1, lobby.ID)
		lobby2 := lbs.GetLobby(GameDiskDC2, lobby.ID)
		if lobby1 == nil || lobby2 == nil {
			return
		}

		renpo1, zeon1 := lobby1.GetUserCountBySide()
		renpo2, zeon2 := lobby2.GetUserCountBySide()
		msgSum1 := NewServerNotice(lbsLobbyJoin).Writer().
			Write16(TeamRenpo).
			Write16(renpo1 + zeon1).
			Write16(renpo2 + zeon2).Msg()
		msgSum2 := NewServerNotice(lbsLobbyJoin).Writer().
			Write16(TeamZeon).
			Write16(renpo1 + zeon1).
			Write16(renpo2 + zeon2).Msg()
		msgRenpo := NewServerNotice(lbsLobbyJoin).Writer().
			Write16(TeamRenpo).
			Write16(renpo1).
			Write16(renpo2).Msg()
		msgZeon := NewServerNotice(lbsLobbyJoin).Writer().
			Write16(TeamZeon).
			Write16(zeon1).
			Write16(zeon2).Msg()

		for userID := range lobby1.Users {
			if p := lbs.FindPeer(userID); p != nil {
				if p.InLobbyChat() {
					p.SendMessage(msgSum1)
					p.SendMessage(msgSum2)
				} else {
					p.SendMessage(msgRenpo)
					p.SendMessage(msgZeon)
				}
			}
		}

		for userID := range lobby2.Users {
			if p := lbs.FindPeer(userID); p != nil {
				if p.InLobbyChat() {
					p.SendMessage(msgSum1)
					p.SendMessage(msgSum2)
				} else {
					p.SendMessage(msgRenpo)
					p.SendMessage(msgZeon)
				}
			}
		}
	}
	discordLiveStatusUpdateAvailable = true
}

func (lbs *Lbs) BroadcastLobbyMatchEntryUserCount(lobby *LbsLobby) {
	renpo, zeon := lobby.GetLobbyMatchEntryUserCount()
	msg1 := NewServerNotice(lbsLobbyMatchingJoin).Writer().Write16(TeamRenpo).Write16(renpo).Msg()
	msg2 := NewServerNotice(lbsLobbyMatchingJoin).Writer().Write16(TeamZeon).Write16(zeon).Msg()
	for userID := range lobby.Users {
		if p := lbs.FindPeer(userID); p != nil {
			p.SendMessage(msg1)
			p.SendMessage(msg2)
		}
	}
	discordLiveStatusUpdateAvailable = true
}

func (lbs *Lbs) BroadcastRoomState(room *LbsRoom) {
	if room == nil || room.lobby == nil {
		return
	}
	msg1 := NewServerNotice(lbsRoomStatus).Writer().Write16(room.ID).Write8(room.Status).Msg()
	msg2 := NewServerNotice(lbsRoomTitle).Writer().Write16(room.ID).WriteString(room.Name).Msg()
	for userID := range room.lobby.Users {
		if p := lbs.FindPeer(userID); p != nil {
			if p.Team == room.Team {
				p.SendMessage(msg1)
				p.SendMessage(msg2)
			}
		}
	}
	discordLiveStatusUpdateAvailable = true
}

func (lbs *Lbs) RegisterBattleResult(p *LbsPeer, result *BattleResult) {
	js, err := json.Marshal(result)
	if err != nil {
		logger.Error("failed to marshal battle result",
			zap.Error(err),
			zap.String("battle_code", result.BattleCode),
			zap.Any("battle_result", result))
		return
	}

	record, err := getDB().GetBattleRecordUser(result.BattleCode, p.UserID)
	if err != nil {
		logger.Error("failed to load battle record",
			zap.Error(err),
			zap.String("battle_code", result.BattleCode),
			zap.Any("battle_result", result))
		return
	}

	record.Round = int(result.BattleCount)
	record.Win = int(result.WinCount)
	record.Lose = int(result.LoseCount)
	record.Kill = int(result.KillCount)
	record.Death = 0 // missing in gdxsv
	record.Frame = 0 // missing in gdxsv
	record.Result = string(js)

	err = getDB().UpdateBattleRecord(record)
	if err != nil {
		logger.Error("failed to save battle record",
			zap.Error(err),
			zap.String("battle_code", result.BattleCode),
			zap.Any("battle_result", result))
		return
	}

	logger.Info("update battle count",
		zap.String("user_id", p.UserID),
		zap.Any("before", p.DBUser))

	rec, err := getDB().CalculateUserTotalBattleCount(p.UserID, 0)
	if err != nil {
		logger.Error("failed to calculate battle count", zap.Error(err))
		return
	}

	p.DBUser.BattleCount = rec.Battle
	p.DBUser.WinCount = rec.Win
	p.DBUser.LoseCount = rec.Lose
	p.DBUser.KillCount = rec.Kill
	p.DBUser.DeathCount = rec.Death

	rec, err = getDB().CalculateUserTotalBattleCount(p.UserID, 1)
	if err != nil {
		logger.Error("failed to calculate battle count", zap.Error(err))
		return
	}

	p.DBUser.RenpoBattleCount = rec.Battle
	p.DBUser.RenpoWinCount = rec.Win
	p.DBUser.RenpoLoseCount = rec.Lose
	p.DBUser.RenpoKillCount = rec.Kill
	p.DBUser.RenpoDeathCount = rec.Death

	rec, err = getDB().CalculateUserTotalBattleCount(p.UserID, 2)
	if err != nil {
		logger.Error("failed to calculate battle count", zap.Error(err))
		return
	}

	p.DBUser.ZeonBattleCount = rec.Battle
	p.DBUser.ZeonWinCount = rec.Win
	p.DBUser.ZeonLoseCount = rec.Lose
	p.DBUser.ZeonKillCount = rec.Kill
	p.DBUser.ZeonDeathCount = rec.Death

	rec, err = getDB().CalculateUserDailyBattleCount(p.UserID)
	if err != nil {
		logger.Error("failed to calculate battle count", zap.Error(err))
		return
	}

	p.DBUser.DailyBattleCount = rec.Battle
	p.DBUser.DailyWinCount = rec.Win
	p.DBUser.DailyLoseCount = rec.Lose

	err = getDB().UpdateUser(&p.DBUser)
	if err != nil {
		logger.Error("failed to update user", zap.Error(err))
		return
	}

	logger.Info("update battle count",
		zap.String("user_id", p.UserID),
		zap.Any("after", p.DBUser))
}

type LbsPeer struct {
	DBUser
	logger *zap.Logger

	conn   *net.TCPConn
	app    *Lbs
	Room   *LbsRoom
	Lobby  *LbsLobby
	Battle *LbsBattle

	GameDisk     byte
	PlatformInfo map[string]string
	Team         uint16
	GameParam    []byte
	PilotName    string
	Rank         int

	lastSessionID string
	lastRecvTime  time.Time
	left          bool

	chWrite    chan bool
	chDispatch chan bool
	chQuit     chan interface{}

	mOutbuf sync.Mutex
	outbuf  []byte

	mInbuf sync.Mutex
	inbuf  []byte

	// used only mcs peer
	mcsStatus *McsStatus
}

func (p *LbsPeer) InLobbyChat() bool {
	return p.Lobby != nil && p.Room == nil && p.Team != TeamNone
}

func (p *LbsPeer) IsPS2() bool {
	return p.GameDisk == GameDiskPS2
}

func (p *LbsPeer) IsDC() bool {
	return p.GameDisk == GameDiskDC1 || p.GameDisk == GameDiskDC2
}

func (p *LbsPeer) IsDC1() bool {
	return p.GameDisk == GameDiskDC1
}

func (p *LbsPeer) IsDC2() bool {
	return p.GameDisk == GameDiskDC2
}

func (p *LbsPeer) serve() {
	defer p.conn.Close()
	defer func() {
		p.app.chEvent <- eventPeerLeave{p}
	}()

	ctx, cancel := context.WithCancel(context.Background())

	go p.dispatchLoop(ctx, cancel)
	go p.writeLoop(ctx, cancel)
	go p.readLoop(ctx, cancel)

	p.app.chEvent <- eventPeerCome{p}
	<-ctx.Done()
}

func (p *LbsPeer) SendMessage(msg *LbsMessage) {
	logger.Debug("lobby -> client",
		zap.String("addr", p.Address()),
		zap.Any("msg", msg),
	)

	p.mOutbuf.Lock()
	p.outbuf = append(p.outbuf, msg.Serialize()...)
	p.mOutbuf.Unlock()
	select {
	case p.chWrite <- true:
	default:
	}
}

func (p *LbsPeer) Address() string {
	return p.conn.RemoteAddr().String()
}

func (p *LbsPeer) readLoop(ctx context.Context, cancel func()) {
	defer cancel()

	buf := make([]byte, 4096)
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		p.conn.SetReadDeadline(time.Now().Add(time.Second * 30))
		n, err := p.conn.Read(buf)
		if err != nil {
			logger.Info("tcp read error", zap.Error(err))
			return
		}
		if n == 0 {
			logger.Info("tcp read zero")
			return
		}
		p.mInbuf.Lock()
		p.inbuf = append(p.inbuf, buf[:n]...)
		p.mInbuf.Unlock()

		select {
		case p.chDispatch <- true:
		default:
		}
	}
}

func (p *LbsPeer) writeLoop(ctx context.Context, cancel func()) {
	defer cancel()

	buf := make([]byte, 0, 128)
	for {
		select {
		case <-ctx.Done():
			return
		case <-p.chWrite:
			p.mOutbuf.Lock()
			if len(p.outbuf) == 0 {
				p.mOutbuf.Unlock()
				continue
			}
			buf = append(buf, p.outbuf...)
			p.outbuf = p.outbuf[:0]
			p.mOutbuf.Unlock()

			sum := 0
			size := len(buf)
			for sum < size {
				p.conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
				n, err := p.conn.Write(buf[sum:])
				if err != nil {
					p.logger.Info("tcp write error", zap.Error(err))
					break
				}
				sum += n
			}
			buf = buf[:0]
		}
	}
}

func (p *LbsPeer) dispatchLoop(ctx context.Context, cancel func()) {
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		case <-p.chDispatch:
			p.mInbuf.Lock()
			for len(p.inbuf) >= HeaderSize {
				n, msg := Deserialize(p.inbuf)
				if n == 0 {
					// not enough data coming
					break
				}

				p.inbuf = p.inbuf[n:]
				if msg != nil {
					p.app.chEvent <- eventPeerMessage{peer: p, msg: msg}
				}
			}
			p.mInbuf.Unlock()
		}
	}
}
