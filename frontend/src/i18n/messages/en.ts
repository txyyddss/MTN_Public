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
      brandSubline: 'Survival atlas',
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
        { id: 'gallery', label: 'Gallery', to: '/gallery' },
        { id: 'wiki', label: 'Wiki', to: 'https://mtn.1919801.xyz/', external: true, emphasize: true }
      ]
    },
    home: {
      hero: {
        eyebrow: 'Community-run survival, founded in 2024',
        title: 'Long-term survival. No shortcut economy.',
        body: 'Java and Bedrock share one world, one history, and one public record.',
        primaryCta: 'Explore Players',
        secondaryCta: 'Open Wiki',
        facts: [
          { label: 'Operations', value: 'Non-profit community run' },
          { label: 'Access', value: 'Java and Bedrock' },
          { label: 'Node', value: 'Dedicated hardware in Shenzhen' }
        ]
      },
      quickRoutes: {
        kicker: 'Quick routes',
        title: 'Move through the archive.',
        fallback: 'Open this route.',
        descriptions: {
          players: 'Browse the current archive and jump into individual records.',
          gallery: 'Open the visual world log.',
          wiki: 'Read the server handbook and gameplay details.'
        }
      },
      serverIntro: {
        kicker: 'Operations surface',
        title: 'Live infrastructure and connection routing.',
        body: 'Check the node state, recent player activity, and the right address for your client before joining.'
      },
      featuresIntro: {
        kicker: 'Operating principles',
        title: 'What this world is built for.',
        body: 'Fair rules, stable survival, shared history made visible.'
      },
      featureMode: 'System',
      features: [
        {
          title: 'Community-run operations',
          description: 'Non-profit, community-run, and built around equal footing.',
          icon: '01',
          accent: 'copper'
        },
        {
          title: 'Vanilla-first survival',
          description: 'Core survival stays intact. Tweaks support fairness, not replacement gameplay.',
          icon: '02',
          accent: 'moss'
        },
        {
          title: 'Cross-play access',
          description: 'Java and Bedrock join the same world with shared public records.',
          icon: '03',
          accent: 'redstone'
        },
        {
          title: 'Dedicated infrastructure',
          description: 'Dedicated hardware in Shenzhen, surfaced live by the site.',
          icon: '04',
          accent: 'copper'
        },
        {
          title: 'Builders with range',
          description: 'Railways, cities, farms, and large shared builds over resets.',
          icon: '05',
          accent: 'moss'
        },
        {
          title: 'Public records',
          description: 'Players, rankings, stats, and screenshots stay public.',
          icon: '06',
          accent: 'redstone'
        }
      ],
      cta: {
        kicker: 'Next step',
        title: 'Join, build, leave a mark.',
        body: 'If you want fair survival with visible history, this is the pace.',
        primaryCta: 'Open Wiki',
        note: 'Connection details sit below.'
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
      historyHint: 'Rolling 7-day hourly presence based on the live sampler.',
      historyLoading: 'Loading recent online history...',
      historyEmpty: 'No activity',
      historyPeak: 'Peak {peak}',
      connectionTitle: 'Connection guide',
      connectionPanelTitle: 'Connection routing',
      connectionLoading: 'Loading connection records...',
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
        archiveBody: 'Historical record and progression summary.'
      },
      profileCardTitle: 'Profile',
      tabsAria: 'Player detail sections',
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
        skillLeaderboard: 'Skill leaderboard',
        xpLevel: 'XP Level',
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
        linkedTo: 'Linked Account'
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
      title: 'Leaderboards',
      body: 'Rankings generated from live backend data.',
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
