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
    brand: 'MT Network',
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
    ] satisfies NavItem[],
    footer:
      'MT Network is a community-run Minecraft survival server founded in 2024. It focuses on fair play, cross-play access, and long-term building.'
  },
  home: {
    hero: {
      eyebrow: 'Community-run survival server, founded in 2024',
      title: 'A long-horizon survival world built for fair play.',
      body:
        'MT Network pairs a vanilla-first ruleset with Java and Bedrock access, a dedicated node in Shenzhen, and a public-facing archive of players, rankings, and world moments.',
      primaryCta: 'Explore Players',
      secondaryCta: 'Read The Server Story',
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
          'The project is maintained as a non-profit community server. The ruleset is built around equal footing rather than paid shortcuts.',
        icon: '01',
        accent: 'copper'
      },
      {
        title: 'Vanilla-first survival',
        description:
          'The core Minecraft survival loop stays intact. Tweaks are there to support fairness, clarity, and long-term play rather than replace the game.',
        icon: '02',
        accent: 'moss'
      },
      {
        title: 'Cross-play access',
        description:
          'Java and Bedrock players can enter the same world, with linked identity data supporting shared progress where the backend exposes it.',
        icon: '03',
        accent: 'redstone'
      },
      {
        title: 'Dedicated infrastructure',
        description:
          'The live node information surfaced by the site points to Shenzhen, China Mobile connectivity, and dedicated hardware rather than a disposable host.',
        icon: '04',
        accent: 'copper'
      },
      {
        title: 'Builders with range',
        description:
          'The server story is tied to city building, railway work, survival automation, and large shared projects instead of short-cycle resets.',
        icon: '05',
        accent: 'moss'
      },
      {
        title: 'Public records',
        description:
          'Players, leaderboards, advancements, statistics, and screenshots are surfaced through the site so the world has a visible memory.',
        icon: '06',
        accent: 'redstone'
      }
    ] satisfies FeatureCardContent[],
    cta: {
      title: 'Join the world, then leave a trace worth keeping.',
      body:
        'If you want a survival server with fair rules, visible history, and room for large-scale building, MT Network is built for that pace.',
      primaryCta: 'Open Server Intro',
      note: 'Connection details for both editions are listed below.'
    }
  },
  intro: {
    hero: {
      eyebrow: 'Server introduction',
      title: 'From a private homeserver to a public survival world.',
      body:
        'The project started small and became public through iteration: better software architecture, stronger infrastructure, and a clearer community identity.'
    },
    story: [
      {
        heading: 'Where it started',
        paragraphs: [
          'MT Network began in October 2024 on a homeserver in Shenzhen. The early core was small: hardware specialist mc9957, veteran player wor114514, and later txrog, who helped reshape the software architecture.',
          'At first it was only a private space for a few friends. As the stack matured, the server moved from a closed Bedrock setup toward a more capable Java-oriented architecture and then opened itself to a wider community.'
        ]
      },
      {
        heading: 'How it changed',
        paragraphs: [
          'Railway builder HEHAWellgood became part of that next phase, bringing a strong infrastructure mindset to the world itself. The identity of the server shifted from a private save into a shared place with public systems, routes, and records.',
          'That arc still defines MT Network today: survival first, large projects welcomed, and technical work in service of a stable long-term world.'
        ]
      }
    ] satisfies IntroStoryBlock[],
    principles: [
      {
        title: 'Fair by default',
        description:
          'The server presents itself as non-profit and non-pay-to-win. Players are meant to enter on equal footing, without paid privilege tiers.'
      },
      {
        title: 'Vanilla at the center',
        description:
          'The goal is not to replace Minecraft with a custom minigame stack. The world stays rooted in survival play, exploration, and building.'
      },
      {
        title: 'Cross-play where it matters',
        description:
          'Java and Bedrock access are both part of the operating model, with linked account data available in the backend when present.'
      },
      {
        title: 'Automation friendly',
        description:
          'The project explicitly positions itself as friendly to technical play, including high-frequency contraptions and larger infrastructure builds.'
      }
    ] satisfies PrincipleContent[],
    infrastructure: {
      title: 'Live operating context',
      body:
        'The site exposes live status, connection records, and system telemetry so players can see how the world is running instead of treating the server as a black box.'
    },
    community: {
      title: 'Official community group',
      body:
        'The QR code on this page links to the official QQ group. It is the clearest verified entry point for announcements, coordination, and questions.',
      caption: 'Scan to join the official MT Network community group.'
    }
  },
  serverPanels: {
    liveStatusTitle: 'Live server status',
    liveStatusFallback: 'Waiting for live telemetry...',
    liveStatusRefresh: 'Refresh cadence: every 5 seconds',
    connectionTitle: 'Connection guide',
    connectionLoading: 'Loading connection records...',
    connectionHint:
      'If your device and network support IPv6, use it first. Keep the wiki link in the navigation for external documentation if it exists outside this repo.',
    simpleView: 'Condensed View',
    fullView: 'Detailed View',
    javaTitle: 'Java Edition',
    bedrockTitle: 'Bedrock Edition',
    copyPrefix: 'Copied',
    labels: {
      platform: 'Platform',
      cpu: 'CPU',
      load: 'Load',
      ram: 'RAM',
      network: 'Network',
      offline: 'Offline'
    }
  },
  players: {
    title: 'Player Directory',
    body:
      'Search the active roster, browse the full archive, or jump directly to a random player record.',
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
      'Rankings are generated from live backend data, covering McMMO totals and major survival statistics.',
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
      'The gallery is a visual log of the world as it changes, using timestamped screenshots stored directly in the repo.',
    capturesLabel: 'captures',
    action: 'Open Frame',
    close: 'Close',
    frameLabel: 'Frame'
  },
  teamPlaceholder: {
    eyebrow: 'Core members',
    title: 'The member index is still being prepared.',
    body:
      'The server story already names several of the people who shaped MT Network, but a dedicated member archive has not been compiled yet.',
    primaryCta: 'Read The Server Intro',
    secondaryCta: 'Open The Community Panel',
    note:
      'For now, use the server introduction for the early project history and the official community group for current contact.'
  }
} as const
