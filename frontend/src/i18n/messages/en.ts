import advancementData from '@/assets/advancements.json'

const advancementNames = Object.fromEntries(
  Object.entries(advancementData).map(([key, value]) => [
    key,
    value.name.replace(/\u807d/g, ' ').replace(/\s+/g, ' ').trim()
  ])
)

const en = {
  siteContent: {
    app: {
      brandSubline: 'Potato server',
      menuToggleAria: 'Toggle navigation',
      locale: {
        label: 'Language',
        options: [
          { value: 'zh-CN', label: '中文' },
          { value: 'en', label: 'EN' }
        ]
      },
      nav: [
        { id: 'home', label: 'Home', to: '/' },
        { id: 'players', label: 'Players', to: '/players' },
        { id: 'wiki', label: 'Documentation', to: 'https://docs.mcmtn.net/', external: true, emphasize: true }
      ]
    },
    seo: {
      siteName: 'MTNetwork',
      default: {
        title: 'MTNetwork | MTN Potato Server',
        description: 'MTNetwork Potato Server is a non-profit Minecraft Java and Bedrock community survival server with live status, player records, and a long-running world archive.'
      },
      routes: {
        home: {
          title: 'MTNetwork | MTN Potato Server',
          description: 'Join MTNetwork Potato Server for long-running vanilla-style survival, Java and Bedrock cross-play, live connection status, and shared community memories.'
        },
        players: {
          title: 'Player Directory | MTNetwork',
          description: 'Search MTNetwork public player records, recent activity, and linked archive pages for the MTN Potato Server community.'
        },
        playerDetail: {
          title: 'Player Record | MTNetwork',
          description: 'View a public MTNetwork player record with profile data, progression summaries, advancements, and statistics.'
        },
        whitelist: {
          title: 'Whitelist Console | MTNetwork',
          description: 'Private whitelist tools for MTNetwork server operators.'
        },
        wiki: {
          title: 'MTNetwork Wiki',
          description: 'Open the external MTNetwork documentation and gameplay guide.'
        }
      }
    },
    home: {
      hero: {
        eyebrow: 'Non-profit vanilla-flavored optimized survival',
        title: 'MTNetwork | MTN Potato Server',
        tagline: 'A community potato server for vanilla-style survival.',
        imageAlt: 'MTN Potato Server in-game build at night',
        projectLine: 'Project MTN long-season archive',
        jumpDetailsCta: 'Status panel',
        jumpJoinCta: 'QQ group',
        scrollCta: 'Scroll to introduction',
        accentLine: 'The potato is pushing harder again.',
        body: 'The hardware is modest, but it is enough for friends to play together. The potato has grown stronger recently, and it keeps working for everyone.',
        primaryCta: 'QQ Group 1064494318',
        primaryHref: '#join-mtn',
        secondaryCta: 'Visit Documentation',
        secondaryHref: 'https://docs.mcmtn.net/',
        facts: [
          { label: 'Operations', value: 'Non-profit' },
          { label: 'Access', value: 'Java / Bedrock cross-play' },
          { label: 'World Goal', value: 'Long-running season' }
        ]
      },
      preloader: {
        steps: ['Opening node link...', 'Checking potato voltage...', 'Loading world archive...', 'Ready.']
      },
      follow: {
        label: 'Follow',
        links: [
          { label: 'Bilibili', short: 'B', href: 'https://space.bilibili.com/1066536503' },
          { label: 'QQ Group', short: 'QQ', href: 'https://qm.qq.com/q/7muRxmmQvu' }
        ]
      },
      heroAside: {
        kicker: 'Server Mascot',
        title: 'The big potato is on duty',
        imageAlt: 'MTN Potato Server mascot character',
        statusLabel: 'Current goal',
        statusValue: 'Share the original joy with more players',
        chips: ['Vanilla survival', 'Community run', 'Cross-device access']
      },
      story: {
        kicker: 'About MTN Potato Server',
        title: 'A big potato planted at the owner’s home.',
        imageCaption: 'Server mascot / MTN Potato Server',
        paragraphs: [
          'As the name suggests, the server performance is modest. It is not professional-grade hardware, but it is enough for friends to play together. The potato has grown stronger recently, lives at the owner’s home, and keeps developing through the owner’s effort and community support.',
          'The server began as a place where the owner and a few roommates could relax together at any time. Opening it now is a way to share that joy with more people.',
          'The goal is long-running operation rather than short resets that erase memories. The server is meant to be social, entertaining, and able to leave players with real emotion and fun.',
          'With Java survival as the main baseline, practical and playful plugins add room for building, redstone engineering, adventure, economy, and trading while keeping the experience close to vanilla.'
        ],
        calloutTitle: 'Player-community operations',
        calloutBody: 'The server is non-profit, maintained by the player community, and built together. Voluntary support is appreciated and goes back into helping the potato grow.'
      },
      quickRoutes: {
        kicker: 'Quick routes',
        title: 'Move through the archive.',
        fallback: 'Open this route.',
        descriptions: {
          players: 'Browse the current archive and jump into individual records.',
          wiki: 'Read the server handbook and gameplay details.'
        }
      },
      news: {
        kicker: 'Latest updates',
        title: 'Recent Server Notes',
        items: [
          { date: '2026.04.28', title: 'The public web archive enters a new visual season.', badge: 'NEW', href: '#details' },
          { date: '2026.04.19', title: 'Fresh world screenshots now rotate through the home hero.', href: '#intro-section' },
          { date: '2026.04.15', title: 'Live status, route checks, and player records are now connected.', href: '#details' },
          { date: '2026.04.10', title: 'The potato server mascot joined the site identity.', href: '#intro-section' }
        ]
      },
      quickCards: {
        title: 'Quick navigation',
        items: [
          { title: 'Feature System', href: '#features' },
          { title: 'Player Archive', href: '/players' },
          { title: 'Server Handbook', href: 'https://docs.mcmtn.net/', external: true }
        ]
      },
      serverIntro: {
        kicker: 'Operations surface',
        title: 'Live infrastructure and connection routing.',
        body: 'Check the node state, recent player activity, and the right address for your client before joining.'
      },
      featuresIntro: {
        kicker: 'What makes the server different',
        title: 'Built for a longer-lived world.'
      },
      featureMode: 'System',
      features: [
        {
          title: 'Server philosophy',
          description: 'Started as a place for friends to relax together, then opened to share that joy with more people.',
          icon: '01',
          accent: 'copper'
        },
        {
          title: 'Player-community operations',
          description: 'Non-profit, maintained together by the community, and shaped by players who build the world.',
          icon: '02',
          accent: 'moss'
        },
        {
          title: 'Vanilla-first survival',
          description: 'The experience stays close to vanilla, with light tuning for fairness and carefully added gameplay.',
          icon: '03',
          accent: 'redstone'
        },
        {
          title: 'Cross-version access',
          description: 'Java and Bedrock can both join smoothly, so players can use the edition they prefer.',
          icon: '04',
          accent: 'copper'
        }
      ],
      featureBackdrops: ['HARDWARE', 'COMMUNITY', 'VANILLA', 'CROSSPLAY'],
      featureQuotes: [
        'A small machine can still carry a real world.',
        'The best archive is the one players keep building together.',
        'Stay close to survival, then add only what helps the story.',
        'One world should be easy to reach from the device you have.'
      ],
      details: {
        slogan: 'Ready to settle in?',
        subSlogan: 'Leave a real footprint in the potato world.',
        groups: [
          {
            title: '// Server information',
            rows: [
              { label: 'Name', value: 'MTNetwork Potato Server' },
              { label: 'Mode', value: 'Java / Bedrock cross-play' },
              { label: 'Operation', value: 'Non-profit community server' },
              { label: 'Open time', value: 'Long-running public season' }
            ]
          },
          {
            title: '// Server setting',
            rows: [
              { label: 'Core play', value: 'Survival / building / redstone / adventure' },
              { label: 'Mechanics', value: 'Vanilla baseline with practical plugins' },
          { label: 'Theme color', value: '#246FCF MTN blue' },
              { label: 'Goal', value: 'A stable world worth remembering' }
            ]
          }
        ]
      },
      cta: {
        kicker: 'Join us now',
        title: 'QQ Group: 1064494318',
        body: 'Join MTN Potato Server for long-running non-profit survival, building, redstone engineering, adventure, trading, and shared memories.',
        primaryCta: 'Join QQ group',
        qqGroup: '1064494318',
        qqGroupUrl: 'https://qm.qq.com/q/7muRxmmQvu',
        siteCta: 'Visit Documentation',
        siteUrl: 'https://docs.mcmtn.net/',
        bilibiliCta: 'Bilibili',
        bilibiliUrl: 'https://space.bilibili.com/1066536503',
        note: 'Documentation: https://docs.mcmtn.net/'
      },
      floatingJoin: {
        label: 'Join Server',
        detail: 'QQ 1064494318',
        href: 'https://qm.qq.com/q/7muRxmmQvu'
      },
      footer: {
        tagline: 'A modest home-hosted potato server for vanilla-first survival, shared builds, and long-running memories.',
        navTitle: 'Explore',
        supportTitle: 'Support',
        links: [
          { label: 'Latest updates', href: '#news', external: true },
          { label: 'Features', href: '#features', external: true },
          { label: 'Players', href: '/players' }
        ],
        legal: 'MTNetwork is an independent Minecraft community and is not affiliated with Mojang or Microsoft.'
      }
    },
    serverPanels: {
      liveStatusTitle: 'Live server status',
      liveStatusPanelTitle: 'Node load',
      liveStatusFallback: 'Waiting for live telemetry...',
      liveStatusRefresh: 'Refresh cadence: every 5 seconds',
      operational: 'Operational',
      standby: 'Standby',
      historyTitle: 'Weekly online heatmap',
      historyLoading: 'Loading recent online history...',
      historyEmpty: 'No activity',
      historyPeak: 'Peak {peak}',
      connectionTitle: 'Connection guide',
      connectionPanelTitle: 'Connection routing',
      connectionHint: 'Prefer IPv6 when your device and network support it.',
      javaTitle: 'Java Edition',
      bedrockTitle: 'Bedrock Edition',
      copy: 'Copy',
      copied: 'Copied',
      lastUpdate: 'Last update',
      labels: {
        platform: 'Platform',
        cpu: 'CPU',
        load: 'Load',
        ram: 'RAM',
        offline: 'Offline'
      }
    },
    players: {
      kicker: 'Public player records',
      title: 'Player Directory',
      body: 'Search the roster, browse the archive, or jump to a random record.',
      controlsKicker: 'Roster query',
      controlsNote: 'Search the live archive or move between recent and full-history views.',
      searchLabel: 'Search',
      loading: 'Loading the latest player records...',
      emptyTitle: 'No players matched this search',
      emptyBody: 'Try a different name fragment or return to the recent-player view.',
      reset: 'Reset Filters',
      searchPlaceholder: 'Search by player name or UUID',
      recentLabel: 'Recent',
      recentDetail: 'last {days}d',
      allLabel: 'Archive',
      allDetail: 'full log',
      randomLabel: 'Random',
      randomTitle: 'Open a random player record',
      onlineNow: 'Online now',
      firstSeenFallback: 'Unknown player',
      visibleInLiveList: 'Visible in the live player list.',
      resultMatching: 'matching "{query}"',
      resultActiveWindow: 'active within the last {days} days',
      resultArchive: 'in the full archive',
      lastSeen: 'Last seen {value}',
      lastSeenUnavailable: 'Last seen unavailable'
    },
    playerDetail: {
      loading: 'Loading the player record...',
      emptyTitle: 'Player record unavailable',
      emptyBody: 'The requested UUID does not have a complete record in the public archive.',
      back: 'Back To Players',
      header: {
        kicker: 'Public player dossier',
        liveBody: 'Live presence detected in the server status feed.',
        archiveBody: 'Historical record and progression summary.',
        uuid: 'UUID',
        edition: 'Edition',
        identityAria: 'Player identity details',
        metricsAria: 'Player summary metrics'
      },
      tabsAria: 'Player detail sections',
      tabs: {
        overview: 'Overview',
        advancements: 'Advancements',
        statistics: 'Statistics'
      },
      summary: {
        onlineNow: 'Online now',
        archiveRecord: 'Archive record',
        advancements: 'Advancements',
        bestRank: 'Best Rank',
        ranks: 'Leaderboard ranks',
        noSkillData: 'No public skill data',
        noHistoryData: 'No recent online samples',
        unranked: 'Unranked'
      },
      profile: {
        firstJoin: 'First Join',
        lastSeen: 'Last Seen',
        playtime: 'Playtime',
        xpLevel: 'XP Level',
        linkedTo: 'Linked Account',
        qqOwner: 'QQ Owner'
      },
      sections: {
        ores: 'Ore Mining Record',
        onlineHistory: 'Weekly Online History',
        skills: 'McMMO Skills',
        advancements: 'Completed Advancements',
        extendedStats: 'Extended Statistics',
        total: 'Total',
        top: 'Top {count}',
        searchPlaceholder: 'Search a stat, for example diamond or walk',
        emptyStats: 'No statistics matched the current category and search.'
      },
      timezoneFallback: 'Local'
    },
    heatmap: {
      online: 'Online',
      offline: 'Offline',
      weekdays: {
        mon: 'Mon',
        tue: 'Tue',
        wed: 'Wed',
        thu: 'Thu',
        fri: 'Fri',
        sat: 'Sat',
        sun: 'Sun',
        monday: 'Monday',
        tuesday: 'Tuesday',
        wednesday: 'Wednesday',
        thursday: 'Thursday',
        friday: 'Friday',
        saturday: 'Saturday',
        sunday: 'Sunday'
      }
    },
    leaderboards: {
      kicker: 'Leaderboard',
      loading: 'Refreshing leaderboard data...',
      visibleEntries: '{count} visible entries',
      emptyTitle: 'No leaderboard data yet',
      emptyBody: 'This category does not have enough public data to rank players right now.',
      close: 'Close',
      paginationAria: 'Leaderboard pagination',
      paginationPrev: 'Prev',
      paginationNext: 'Next',
      paginationStatus: 'Page {page} / {total}',
      columns: {
        rank: 'Rank',
        player: 'Player',
        score: 'Score'
      }
    },
    formats: {
      playerCount: '{count} {noun}',
      playerSingular: 'player',
      playerPlural: 'players',
      leaderboardTitle: '{label} Leaderboard',
      durationHoursShort: '{value}h',
      durationHoursWide: '{value} hrs',
      durationMinutesShort: '{value}m',
      distanceKilometers: '{value} km',
      distanceMeters: '{value} m'
    },
    gallery: {
      eyebrow: 'World archive',
      title: 'Captured moments from the server',
      body: 'A timestamped visual log of the world.',
      close: 'Close',
      frameLabel: 'Frame',
      previousImage: 'Previous image',
      nextImage: 'Next image'
    },
    wikiRedirect: {
      kicker: 'Wiki',
      title: 'Redirecting to the MTNetwork wiki.',
      body: 'If the redirect does not start automatically, use the link below.',
      cta: 'Open Wiki'
    },
    teamPlaceholder: {
      eyebrow: 'Placeholder',
      title: 'This page is still being assembled.',
      body: 'The surface is reserved and will be filled with the final content once it is ready.',
      primaryCta: 'Open Wiki',
      secondaryCta: 'Return Home',
      note: 'Navigation and layout stay in place so the route can ship safely.'
    }
  },
  leaderboardLabels: {
    skills: 'McMMO Total',
    playtime: 'Playtime',
    mining: 'Blocks Mined',
    killing: 'Mob Kills',
    deaths: 'Deaths',
    walking: 'Distance Walked',
    pvp: 'Player Kills'
  },
  playerStatGroups: [
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
  ],
  statCategories: {
    custom: 'Global',
    mined: 'Mined',
    broken: 'Broken',
    crafted: 'Crafted',
    used: 'Used',
    picked_up: 'Picked Up',
    dropped: 'Dropped',
    killed: 'Killed',
    killed_by: 'Killed By'
  },
  statTranslations: {
    play_one_minute: 'Time Played',
    play_time: 'Time Played',
    jump: 'Jumps',
    damage_dealt: 'Damage Dealt',
    damage_taken: 'Damage Taken',
    deaths: 'Deaths',
    mob_kills: 'Mob Kills',
    player_kills: 'Player Kills',
    walk_one_cm: 'Distance Walked',
    sprint_one_cm: 'Distance Sprinted',
    fly_one_cm: 'Distance Flown',
    climb_one_cm: 'Distance Climbed',
    fall_one_cm: 'Distance Fallen',
    minecart_one_cm: 'Distance by Minecart',
    boat_one_cm: 'Distance by Boat',
    pig_one_cm: 'Distance by Pig',
    horse_one_cm: 'Distance by Horse',
    strider_one_cm: 'Distance by Strider',
    aviate_one_cm: 'Distance by Elytra',
    swim_one_cm: 'Distance Swum',
    walk_on_water_one_cm: 'Distance on Water',
    walk_under_water_one_cm: 'Distance Under Water',
    time_since_death: 'Time Since Last Death',
    time_since_rest: 'Time Since Last Rest',
    sneak_time: 'Sneak Time',
    total_world_time: 'Total World Time',
    leave_game: 'Games Quit',
    dropped: 'Items Dropped',
    interact_with_beacon: 'Beacon Interactions',
    inspect_hopper: 'Hopper Inspections',
    interact_with_blast_furnace: 'Blast Furnace Interactions',
    interact_with_smoker: 'Smoker Interactions',
    interact_with_camp_fire: 'Campfire Interactions',
    interact_with_campfire: 'Campfire Interactions',
    talked_to_villager: 'Villager Talks',
    traded_with_villager: 'Villager Trades',
    fish_caught: 'Fish Caught',
    sleep_in_bed: 'Times Slept',
    raid_win: 'Raids Won',
    raid_trigger: 'Raids Triggered',
    trigger_trapped_chest: 'Trapped Chests Triggered',
    damage_absorbed: 'Damage Absorbed',
    interact_with_furnace: 'Furnace Interactions',
    crouch_one_cm: 'Distance Crouched',
    interact_with_stonecutter: 'Stonecutter Interactions',
    damage_resisted: 'Damage Resisted',
    damage_blocked_by_shield: 'Damage Blocked by Shield',
    interact_with_crafting_table: 'Crafting Table Interactions',
    inspect_dropper: 'Dropper Inspections',
    target_hit: 'Targets Hit',
    fill_cauldron: 'Cauldrons Filled',
    interact_with_grindstone: 'Grindstone Interactions',
    open_shulker_box: 'Shulker Boxes Opened',
    open_enderchest: 'Ender Chests Opened',
    damage_dealt_absorbed: 'Damage Dealt (Absorbed)',
    interact_with_brewingstand: 'Brewing Stand Interactions',
    inspect_dispenser: 'Dispenser Inspections',
    interact_with_loom: 'Loom Interactions',
    play_noteblock: 'Note Blocks Played',
    interact_with_lectern: 'Lectern Interactions',
    drop: 'Items Dropped',
    use_cauldron: 'Cauldrons Used',
    bell_ring: 'Bells Rung',
    open_barrel: 'Barrels Opened',
    interact_with_cartography_table: 'Cartography Table Interactions',
    open_chest: 'Chests Opened',
    tune_noteblock: 'Note Blocks Tuned',
    interact_with_anvil: 'Anvil Interactions',
    animals_bred: 'Animals Bred',
    play_record: 'Music Discs Played',
    interact_with_smithing_table: 'Smithing Table Interactions',
    pot_flower: 'Flowers Potted',
    clean_banner: 'Banners Cleaned',
    enchant_item: 'Items Enchanted',
    clean_armor: 'Armor Pieces Cleaned',
    clean_shulker_box: 'Shulker Boxes Cleaned',
    eat_cake_slice: 'Cake Slices Eaten',
    interact_with_beehive: 'Beehive Interactions',
    interact_with_lodestone: 'Lodestone Interactions',
    interact_with_respawn_anchor: 'Respawn Anchor Interactions'
  },
  skillLabels: {
    taming: 'Taming',
    mining: 'Mining',
    woodcutting: 'Woodcutting',
    repair: 'Repair',
    unarmed: 'Unarmed',
    herbalism: 'Herbalism',
    excavation: 'Excavation',
    archery: 'Archery',
    swords: 'Swords',
    axes: 'Axes',
    acrobatics: 'Acrobatics',
    fishing: 'Fishing',
    alchemy: 'Alchemy',
    crossbows: 'Crossbows',
    tridents: 'Tridents',
    maces: 'Maces',
    spears: 'Spears'
  },
  platformLabels: {
    Java: 'Java',
    Bedrock: 'Bedrock'
  },
  oreLabels: {
    quartz: 'Quartz',
    coal: 'Coal',
    redstone: 'Redstone',
    iron: 'Iron',
    copper: 'Copper',
    diamond: 'Diamond',
    gold: 'Gold',
    lapis: 'Lapis',
    emerald: 'Emerald',
    nether_gold_ore: 'Nether Gold',
    ancient_debris: 'Ancient Debris'
  },
  advancementTypes: {
    task: 'Task',
    goal: 'Goal',
    challenge: 'Challenge'
  },
  advancementCategories: {
    adventure: 'Adventure',
    husbandry: 'Husbandry',
    minecraft: 'Minecraft',
    nether: 'Nether',
    story: 'Story',
    end: 'The End',
    others: 'Others'
  },
  advancementNames
} as const

export default en
