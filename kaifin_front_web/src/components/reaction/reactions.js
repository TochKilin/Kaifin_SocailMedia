export const REACTIONS = [
  {
    key: 'private_like',
    label: 'Private Like',
    private: true,
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
      <path d="M7 11v9H4v-9h3Zm3 9h8a2 2 0 0 0 2-2l1.5-5a2 2 0 0 0-2-2.6H15l.7-4A2 2 0 0 0 13.8 4L10 10v10Z" fill="currentColor"/>
    </svg>`,
  },
  {
    key: 'sad',
    label: 'Sade',
    icon: `<svg viewBox="0 0 36 36">
      <circle cx="18" cy="18" r="17" fill="#F2C94C"/>
      <ellipse cx="12.5" cy="15.5" rx="2.2" ry="2.8" fill="#3A2A1A"/>
      <ellipse cx="23.5" cy="15.5" rx="2.2" ry="2.8" fill="#3A2A1A"/>
      <path d="M11 25c1.8-3 5-4.5 7-4.5s5.2 1.5 7 4.5" stroke="#3A2A1A" stroke-width="2" fill="none" stroke-linecap="round"/>
      <path d="M23.8 17c1.2 1.6 1.4 3.6.4 5.4-.6 1-1.8 1.2-2.3.2-.4-.8.1-1.6.6-2.3.6-.9 1-2.1 1.3-3.3Z" fill="#4FA8D8"/>
    </svg>`,
  },
  {
    key: 'wow',
    label: 'Supprise',
    icon: `<svg viewBox="0 0 36 36">
      <circle cx="18" cy="18" r="17" fill="#F2C94C"/>
      <ellipse cx="12.5" cy="14.5" rx="2.4" ry="3" fill="#3A2A1A"/>
      <ellipse cx="23.5" cy="14.5" rx="2.4" ry="3" fill="#3A2A1A"/>
      <ellipse cx="18" cy="24" rx="3.6" ry="4.4" fill="#3A2A1A"/>
    </svg>`,
  },
  {
    key: 'love',
    label: 'Love',
    icon: `<svg viewBox="0 0 36 36">
      <circle cx="18" cy="18" r="17" fill="#F2C94C"/>
      <path d="M12.5 13.2c-1.6 0-2.9 1.2-2.9 2.8 0 2.1 2.9 4 2.9 4s2.9-1.9 2.9-4c0-1.6-1.3-2.8-2.9-2.8Z" fill="#E8543A"/>
      <path d="M23.5 13.2c-1.6 0-2.9 1.2-2.9 2.8 0 2.1 2.9 4 2.9 4s2.9-1.9 2.9-4c0-1.6-1.3-2.8-2.9-2.8Z" fill="#E8543A"/>
      <path d="M11 24c1.8 2.6 4.6 4 7 4s5.2-1.4 7-4" stroke="#3A2A1A" stroke-width="2" fill="none" stroke-linecap="round"/>
    </svg>`,
  },
  {
    key: 'haha',
    label: 'HaHa',
    icon: `<svg viewBox="0 0 36 36">
      <circle cx="18" cy="18" r="17" fill="#F2C94C"/>
      <path d="M9.5 15.5c1-1.6 2.4-2.3 3.5-2.3s2.5.7 3.5 2.3" stroke="#3A2A1A" stroke-width="2" fill="none" stroke-linecap="round"/>
      <path d="M19.5 15.5c1-1.6 2.4-2.3 3.5-2.3s2.5.7 3.5 2.3" stroke="#3A2A1A" stroke-width="2" fill="none" stroke-linecap="round"/>
      <path d="M9.5 20c1.5 4 5 6.5 8.5 6.5s7-2.5 8.5-6.5Z" fill="#3A2A1A"/>
      <path d="M13.5 20h9c-.4 2-2.2 3.5-4.5 3.5s-4.1-1.5-4.5-3.5Z" fill="#fff"/>
    </svg>`,
  },
  {
    key: 'like',
    label: 'Like',
    icon: `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
      <path d="M7 11v9H4v-9h3Zm3 9h8a2 2 0 0 0 2-2l1.5-5a2 2 0 0 0-2-2.6H15l.7-4A2 2 0 0 0 13.8 4L10 10v10Z" fill="currentColor"/>
    </svg>`,
  },
  {
    key: 'heart',
    label: 'Heart',
    icon: `<svg viewBox="0 0 24 24">
      <path d="M12 21s-7.5-4.6-10-9.1C.4 8.6 2 5 5.6 5 8 5 10 6.4 12 9c2-2.6 4-4 6.4-4C22 5 23.6 8.6 22 11.9 19.5 16.4 12 21 12 21Z" fill="#E8543A"/>
    </svg>`,
  },
  {
    key: 'rose',
    label: 'Rose',
    icon: `<svg viewBox="0 0 24 24">
      <path d="M12 3.5c1.5 0 4 1.2 4 3.7 0 1.4-.8 2.4-1.6 3-.2.1-.1.4.1.4 1.2.1 2.6-.5 3.5-1.6.2-.2.5 0 .4.2-.7 2.4-3 4.1-5.4 4.1-1.7 0-3-.9-4-.9s-2.3.9-4 .9c-2.4 0-4.7-1.7-5.4-4.1-.1-.2.2-.4.4-.2.9 1.1 2.3 1.7 3.5 1.6.2 0 .3-.3.1-.4C2.8 9.6 2 8.6 2 7.2c0-2.5 2.5-3.7 4-3.7 2.3 0 3.6 1.6 4 2.1.4-.5 1.7-2.1 2-2.1Z" fill="#C6402E"/>
      <path d="M12 12.5v8" stroke="#2E7D32" stroke-width="1.8" stroke-linecap="round"/>
      <path d="M12 16c-1.5 0-2.6-1-3-2" stroke="#2E7D32" stroke-width="1.6" fill="none" stroke-linecap="round"/>
    </svg>`,
  },
  {
    key: 'clap',
    label: 'Congrate',
    icon: `<svg viewBox="0 0 24 24">
      <path d="M9 14.5l-2.6-4.3a1.5 1.5 0 1 1 2.6-1.5L11 12" fill="#F2C48B"/>
      <path d="M15 14.5l2.6-4.3a1.5 1.5 0 1 0-2.6-1.5L13 12" fill="#F2C48B"/>
      <path d="M8 15c-.5-2 .5-4 2-4.5s3 0 3.5 1.5" stroke="#D9A15C" stroke-width="1" fill="none"/>
      <path d="M4.5 6.5l1.5 1M19.5 6.5l-1.5 1M12 4v1.6" stroke="#4A4A4E" stroke-width="1.6" stroke-linecap="round"/>
      <path d="M8 16c0 3 1.8 5 4 5s4-2 4-5" fill="#F2C48B"/>
    </svg>`,
  },
  {
    key: 'pray',
    label: 'Pray',
    icon: `<svg viewBox="0 0 24 24">
      <path d="M12 3l-1 6-3.5 4v7c0 .8.7 1.5 1.5 1.5s1.5-.7 1.5-1.5v-4h3v4c0 .8.7 1.5 1.5 1.5s1.5-.7 1.5-1.5v-7l-3.5-4-1-6Z" fill="#7C6FE8"/>
      <path d="M12 3v17.5" stroke="#5C4FC7" stroke-width="1" />
    </svg>`,
  },
]