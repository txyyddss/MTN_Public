import type { LeaderboardType, StatGroup } from '@/types/api'

interface NavItem {
  label: string
  to: string
  external?: boolean
  emphasize?: boolean
}

interface FeatureCardContent {
  title: string
  description: string
  icon: string
  accent: 'copper' | 'moss' | 'redstone'
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
      { label: 'Core Members', to: '/core-members' },
      { label: 'Gallery', to: '/gallery' },
      { label: 'Wiki', to: 'https://mtn.1919801.xyz/', external: true, emphasize: true }
    ] satisfies NavItem[]
  },
  home: {
    hero: {
      eyebrow: 'Community-run survival, founded in 2024',
      title: 'Long-term survival. No shortcut economy.',
      body:
        'Java and Bedrock share one world, one history, and one public record.',
      primaryCta: 'Explore Players',
      secondaryCta: 'Open Wiki',
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
      primaryCta: 'Open Wiki',
      note: 'Connection details sit below.'
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
      'A dedicated member archive has not been compiled yet. Use the wiki for current server documentation and public references.',
    primaryCta: 'Open Wiki',
    secondaryCta: 'Back To Home',
    note:
      'The wiki is now the primary public reference for MTNetwork.'
  }
} as const
