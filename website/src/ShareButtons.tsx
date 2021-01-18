import React from 'react';
import { useIntl } from 'react-intl';
import { WebsiteUrl } from './Const';

interface ILink {
    href: string;
    ariaLabel: string;
    children: React.ReactNode;
}

const Link = ({ href, ariaLabel, children }: ILink) => (
    <a
        href={href}
        target="_blank"
        aria-label={ariaLabel}
        rel="noopener noreferrer nofollow"
        style={{ margin: '0 0.2rem' }}
    >
        {children}
    </a>
);

const TwitterButton = () => {
    const intl = useIntl();
    const text = intl.formatMessage({ id: "footer.tweet.data-text" });
    const url = escape(WebsiteUrl);
    const hashTags = ([
        'gdxsv',
        '連ジ',
        'ガンダム',
        'Gundam',
    ].join(','));
    const href = `https://twitter.com/share?text=${text}&url=${url}&hashtags=${hashTags}`;

    return (
        <Link href={href} ariaLabel="Share on Twitter">
            <svg style={{ width: '2rem' }} xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M12 0C5.38 0 0 5.38 0 12s5.38 12 12 12 12-5.38 12-12S18.62 0 12 0zm5.26 9.38v.34c0 3.48-2.64 7.5-7.48 7.5-1.48 0-2.87-.44-4.03-1.2 1.37.17 2.77-.2 3.9-1.08-1.16-.02-2.13-.78-2.46-1.83.38.1.8.07 1.17-.03-1.2-.24-2.1-1.3-2.1-2.58v-.05c.35.2.75.32 1.18.33-.7-.47-1.17-1.28-1.17-2.2 0-.47.13-.92.36-1.3C7.94 8.85 9.88 9.9 12.06 10c-.04-.2-.06-.4-.06-.6 0-1.46 1.18-2.63 2.63-2.63.76 0 1.44.3 1.92.82.6-.12 1.95-.27 1.95-.27-.35.53-.72 1.66-1.24 2.04z" /></svg>
        </Link>
    );
};

export { TwitterButton };
