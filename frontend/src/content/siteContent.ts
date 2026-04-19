import type { LeaderboardType, StatGroup } from '@/types/api'

interface NavItem {
  label: string
  to: string
  external?: boolean
}

interface FeatureCardContent {
  title: string
  description: string
  icon: string
  accent: 'copper' | 'moss' | 'redstone'
}

interface PrincipleContent {
  title: string
  description: string
}

interface IntroStoryBlock {
  heading: string
  paragraphs: string[]
}

export const leaderboardLabels: Record<LeaderboardType, string> = {
  skills: 'McMMO Total',
  playtime: 'Playtime',
  mining: 'Blocks Mined',
  killing: 'Mob Kills',
  deaths: 'Deaths',
  walking: 'Distance Walked',
  pvp: 'Player Kills'
}

export const playerStatGroups: StatGroup[] = [
  { name: 'Global Statistics', categories: ['minecraft:custom'] },
  {
    name: 'Items And Blocks',
    categories: [
      'minecraft:mined',
      'minecraft:broken',
      'minecraft:crafted',
      'minecraft:used',
      'minecraft:picked_up',
      'minecraft:dropped'
    ]
  },
  { name: 'Combat And Mobs', categories: ['minecraft:killed', 'minecraft:killed_by'] }
]

export const siteContent = {
  app: {
    loader: {
      kicker: 'SURVIVAL ATLAS',
      title: 'Mapping the world...'
    },
    nav: [
      { label: 'Home', to: '/' },
      { label: 'Players', to: '/players' },
      { label: 'Leaderboards', to: '/leaderboards' },
      { label: 'Server Intro', to: '/server-intro' },
      { label: 'Core Members', to: '/core-members' },
      { label: 'Gallery', to: '/gallery' },
      { label: 'Wiki', to: '/wiki', external: true }
    ] satisfies NavItem[]
  },
  home: {
    hero: {
      eyebrow: 'Community-run survival, founded in 2024',
      title: 'Long-term survival. No shortcut economy.',
      body:
        'Java and Bedrock share one world, one history, and one public record.',
      primaryCta: 'Explore Players',
      secondaryCta: 'Read The Story',
      facts: [
        { label: 'Operations', value: 'Non-profit community run' },
        { label: 'Access', value: 'Java and Bedrock' },
        { label: 'Node', value: 'Dedicated hardware in Shenzhen' }
      ]
    },
    features: [
      {
        title: 'Community-run operations',
        description:
          'Non-profit, community-run, and built around equal footing.',
        icon: '01',
        accent: 'copper'
      },
      {
        title: 'Vanilla-first survival',
        description:
          'Core survival stays intact. Tweaks support fairness, not replacement gameplay.',
        icon: '02',
        accent: 'moss'
      },
      {
        title: 'Cross-play access',
        description:
          'Java and Bedrock join the same world with shared public records.',
        icon: '03',
        accent: 'redstone'
      },
      {
        title: 'Dedicated infrastructure',
        description:
          'Dedicated hardware in Shenzhen, surfaced live by the site.',
        icon: '04',
        accent: 'copper'
      },
      {
        title: 'Builders with range',
        description:
          'Railways, cities, farms, and large shared builds over resets.',
        icon: '05',
        accent: 'moss'
      },
      {
        title: 'Public records',
        description:
          'Players, rankings, stats, and screenshots stay public.',
        icon: '06',
        accent: 'redstone'
      }
    ] satisfies FeatureCardContent[],
    cta: {
      title: 'Join, build, leave a mark.',
      body:
        'If you want fair survival with visible history, this is the pace.',
      primaryCta: 'Open Intro',
      note: 'Connection details sit below.'
    }
  },
  intro: {
    hero: {
      eyebrow: 'Server introduction',
      title: 'From private homeserver to public world.',
      body:
        'The stack grew, the hardware improved, and the server opened up.'
    },
    story: [
      {
        heading: 'Where it started',
        paragraphs: [
          'MTNetwork began in October 2024 on a homeserver in Shenzhen. The early core was small: mc9957, wor114514, and later txrog.',
          'It started as a private space for friends, then moved toward a stronger Java-oriented stack and a public community.'
        ]
      },
      {
        heading: 'How it changed',
        paragraphs: [
          'Railway builder HEHAWellgood helped push the world toward public systems, routes, and larger projects.',
          'That arc still defines MTNetwork: survival first, large work welcomed, and technical effort in service of stability.'
        ]
      }
    ] satisfies IntroStoryBlock[],
    principles: [
      {
        title: 'Fair by default',
        description:
          'Non-profit and non-pay-to-win. Everyone starts on equal footing.'
      },
      {
        title: 'Vanilla at the center',
        description:
          'The world stays rooted in survival, exploration, and building.'
      },
      {
        title: 'Cross-play where it matters',
        description:
          'Java and Bedrock both remain part of the operating model.'
      },
      {
        title: 'Automation friendly',
        description:
          'Technical play, automation, and infrastructure builds are welcome.'
      }
    ] satisfies PrincipleContent[],
    infrastructure: {
      title: 'Live operating context',
      body:
        'Status, addresses, and node load stay visible.'
    },
    community: {
      title: 'Official community group',
      body:
        'The QR code links to the official QQ group for announcements and questions.',
      caption: 'Scan to join the MTNetwork QQ group.'
    }
  },
  serverPanels: {
    liveStatusTitle: 'Live server status',
    liveStatusFallback: 'Waiting for live telemetry...',
    liveStatusRefresh: 'Refresh cadence: every 5 seconds',
    connectionTitle: 'Connection guide',
    connectionLoading: 'Loading connection records...',
    connectionHint:
      'Prefer IPv6 when your device and network support it.',
    simpleView: 'Compact',
    fullView: 'Detailed',
    javaTitle: 'Java Edition',
    bedrockTitle: 'Bedrock Edition',
    copyPrefix: 'Copied',
    labels: {
      platform: 'Platform',
      cpu: 'CPU',
      load: 'Load',
      ram: 'RAM',
      offline: 'Offline'
    }
  },
  players: {
    title: 'Player Directory',
    body:
      'Search the roster, browse the archive, or jump to a random record.',
    loading: 'Loading the latest player records...',
    emptyTitle: 'No players matched this search',
    emptyBody: 'Try a different name fragment or return to the recent-player view.',
    reset: 'Reset Filters',
    searchPlaceholder: 'Search by player name or UUID',
    recentLabel: 'Recent',
    allLabel: 'Archive',
    randomTitle: 'Open a random player record',
    onlineNow: 'Online now',
    firstSeenFallback: 'Unknown player'
  },
  playerDetail: {
    loading: 'Loading the player record...',
    emptyTitle: 'Player record unavailable',
    emptyBody: 'The requested UUID does not have a complete record in the public archive.',
    back: 'Back To Players',
    tabs: {
      overview: 'Overview',
      advancements: 'Advancements',
      statistics: 'Statistics'
    },
    summary: {
      onlineNow: 'Online now',
      archiveRecord: 'Archive record',
      linkedTo: 'Linked to',
      playtime: 'Playtime',
      advancements: 'Advancements',
      skillTotal: 'McMMO Total',
      xpLevel: 'XP Level',
      ranks: 'Leaderboard ranks',
      noSkillData: 'No public skill data'
    },
    profile: {
      firstJoin: 'First Join',
      lastSeen: 'Last Seen',
      playtime: 'Playtime',
      xpLevel: 'XP Level',
      linkedTo: 'Linked Account'
    },
    sections: {
      ores: 'Ore Mining Record',
      skills: 'McMMO Skills',
      advancements: 'Completed Advancements',
      extendedStats: 'Extended Statistics',
      total: 'Total',
      searchPlaceholder: 'Search a stat, for example diamond or walk',
      emptyStats: 'No statistics matched the current category and search.',
      challenge: 'Challenge',
      goal: 'Goal',
      task: 'Task'
    }
  },
  leaderboards: {
    title: 'Leaderboards',
    body:
      'Rankings generated from live backend data.',
    loading: 'Refreshing leaderboard data...',
    emptyTitle: 'No leaderboard data yet',
    emptyBody: 'This category does not have enough public data to rank players right now.',
    columns: {
      rank: 'Rank',
      player: 'Player',
      score: 'Score'
    }
  },
  gallery: {
    eyebrow: 'World archive',
    title: 'Captured moments from the server',
    body:
      'A timestamped visual log of the world.',
    capturesLabel: 'captures',
    action: 'Open Frame',
    close: 'Close',
    frameLabel: 'Frame'
  },
  teamPlaceholder: {
    eyebrow: 'Core members',
    title: 'The member index is still being prepared.',
    body:
      'The server story already names several of the people who shaped MTNetwork, but a dedicated member archive has not been compiled yet.',
    primaryCta: 'Read The Server Intro',
    secondaryCta: 'Open The Community Panel',
    note:
      'For now, use the server introduction for the early project history and the official community group for current contact.'
  }
} as const
