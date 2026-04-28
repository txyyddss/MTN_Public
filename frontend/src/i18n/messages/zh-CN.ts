const zhCN = {
  siteContent: {
    app: {
      brandSubline: '土豆服务器',
      menuToggleAria: '切换导航菜单',
      locale: {
        label: '语言',
        options: [
          { value: 'zh-CN', label: '中文' },
          { value: 'en', label: 'EN' }
        ]
      },
      nav: [
        { id: 'home', label: '首页', to: '/' },
        { id: 'players', label: '玩家', to: '/players' },
        { id: 'gallery', label: '画廊', to: '/gallery' },
        { id: 'wiki', label: '文档', to: 'https://docs.mcmtn.net/', external: true, emphasize: true }
      ]
    },
    home: {
      hero: {
        eyebrow: '公益原版风味生存优化之服务器',
        title: 'MTNetwork | MTN土豆服务器',
        tagline: '「公益神人土豆原版风味生存优化之服务器」',
        imageAlt: 'MTN 土豆服务器夜景建筑截图',
        projectLine: 'PROJECT MTN LONG-SEASON ARCHIVE',
        scrollCta: '滚动到介绍',
        accentLine: '土豆又发力了！！',
        body: '性能一般般，但足够朋友间一同玩耍。最近土豆变强了不少，大土豆会努力服务大家的～～',
        primaryCta: 'QQ群 1064494318',
        primaryHref: '#join-mtn',
        secondaryCta: '访问文档',
        secondaryHref: 'https://docs.mcmtn.net/',
        facts: [
          { label: '运营方式', value: '公益非营利' },
          { label: '接入方式', value: 'Je / Be 跨版本互通' },
          { label: '世界目标', value: '长周目运营' }
        ]
      },
      preloader: {
        steps: ['正在建立节点连接...', '正在检查土豆电压...', '正在加载世界档案...', '就绪。']
      },
      follow: {
        label: 'FOLLOW',
        links: [
          { label: 'Bilibili', short: 'B', href: 'https://space.bilibili.com/1066536503' },
          { label: 'QQ 群', short: 'QQ', href: 'https://qm.qq.com/q/7muRxmmQvu' }
        ]
      },
      heroAside: {
        kicker: '服务器娘',
        title: '大土豆待命中',
        imageAlt: 'MTN 土豆服务器的服务器娘形象',
        statusLabel: '当前目标',
        statusValue: '把欢乐分享给更多人',
        chips: ['原版生存', '社区驱动', '跨设备互通']
      },
      story: {
        kicker: '了解 MTN 土豆服务器',
        title: '一颗种在腐竹家的大土豆。',
        imageCaption: '服务器娘 / MTN 土豆服务器',
        paragraphs: [
          '如同其名字，服务器的性能一般般，虽然并非专业服务器，但是足够朋友间一同玩耍。（最近土豆变强了不少）土豆平时种植在腐竹家，平时靠腐竹攒米发育，以及几个原始咕咚喂养，大土豆会努力服务大家的～～',
          '服务器建立之初的愿景是腐竹与几位舍友为了能在任何时间一起消遣而建立的，如今开放是希望将当初那份欢乐分享与更多人。',
          '服务器期望进行长周目运营，而非伴随短期存档的重置而导致记忆的消失。这里希望搭建一个富有社交性与娱乐性，使人们收获真面情绪与乐趣的地方。',
          '在以 Java 版生存为主基调上，服务器增添了不少更实用、趣味的插件，使其兼具建筑、生电、冒险、经济交易等更多主题。愿你在游玩时，也能收获你的那份欢乐。'
        ],
        calloutTitle: '玩家社区驱动的运营',
        calloutBody: '服务器实行非营利运营、由玩家社区共同维护，共同创造。一定要投喂腐竹的话我们也会很高兴哦：），代表我们的大土豆感谢您～～'
      },
      quickRoutes: {
        kicker: '快捷入口',
        title: '快速穿梭整个档案库。',
        fallback: '打开此页面。',
        descriptions: {
          players: '浏览当前档案，并跳转到任意玩家记录。',
          gallery: '打开世界影像档案。',
          wiki: '阅读服务器手册与玩法说明。'
        }
      },
      news: {
        kicker: '最新动态',
        title: '服务器近况',
        items: [
          { date: '2026.04.28', title: '公开网页档案进入新的视觉阶段。', badge: 'NEW', href: '#details' },
          { date: '2026.04.19', title: '新的世界截图已加入画廊档案。', href: '#gallery' },
          { date: '2026.04.15', title: '实时状态、连接路由与玩家记录已经接入。', href: '#details' },
          { date: '2026.04.10', title: '土豆服务器看板娘加入站点识别。', href: '#intro-section' }
        ]
      },
      quickCards: {
        title: '快速导航',
        items: [
          { title: '特色系统', href: '#features' },
          { title: '世界画廊', href: '#gallery' },
          { title: '服务器手册', href: 'https://docs.mcmtn.net/', external: true }
        ]
      },
      serverIntro: {
        kicker: '运行面板',
        title: '实时基础设施与连接路由。',
        body: '在进入服务器之前，先检查节点状态、最近玩家活动以及适合你客户端的连接地址。'
      },
      featuresIntro: {
        kicker: '服务器与众不同的点在于',
        title: '为长寿存档和真实乐趣服务。',
        body: '公益跨版本互通、原版优先、多样玩法与社区规划，一起让这个世界更适合长期留下记忆。'
      },
      featureMode: '系统',
      features: [
        {
          title: '服务器理念',
          description: '从朋友间随时消遣出发，开放后希望把当初那份欢乐分享给更多人。',
          icon: '01',
          accent: 'copper'
        },
        {
          title: '玩家社区驱动的运营',
          description: '非营利运营，由玩家社区共同维护、共同创造，所有支持都会回到大土豆的成长里。',
          icon: '02',
          accent: 'moss'
        },
        {
          title: '原版优先的生存',
          description: '体验尽量靠近原版，以微调维持公平，并适度替代或增加有趣的新玩法。',
          icon: '03',
          accent: 'redstone'
        },
        {
          title: '跨版本互通',
          description: '双版本无障碍互通【Je/be】，无论你喜欢哪一种版本都可以畅玩本服。',
          icon: '04',
          accent: 'copper'
        },
        {
          title: '公益跨设备覆盖',
          description: '完全公益跨版本互通，覆盖几乎所有设备，让更多玩家能进入同一个世界。',
          icon: '05',
          accent: 'moss'
        },
        {
          title: '长周目与社区规划',
          description: '规则不是为了处罚玩家，而是为了让存档更长寿，并用档案规划减少玩家纠纷。',
          icon: '06',
          accent: 'redstone'
        }
      ],
      featureBackdrops: ['HARDWARE', 'COMMUNITY', 'VANILLA', 'CROSSPLAY'],
      featureQuotes: [
        '小机器也可以承载一个真实的世界。',
        '最好的档案，是玩家一起继续建设的档案。',
        '尽量靠近生存本身，只加入真正有帮助的内容。',
        '同一个世界，应该能被更多设备抵达。'
      ],
      homeGallery: {
        title: '精选截图',
        note: '图片来自 MTN 世界截图，并已为首页展示优化。',
        items: [
          { caption: '出生点周边档案' },
          { caption: '共同路线工程' },
          { caption: '长周目建筑记录' }
        ]
      },
      details: {
        slogan: '准备好定居了吗',
        subSlogan: '在土豆世界里留下真实足迹。',
        groups: [
          {
            title: '// 服务器信息',
            rows: [
              { label: '名称', value: 'MTNetwork 土豆服务器' },
              { label: '模式', value: 'Java / Bedrock 互通' },
              { label: '运营', value: '公益社区服务器' },
              { label: '开放', value: '长期公开周目' }
            ]
          },
          {
            title: '// 服务器设定',
            rows: [
              { label: '核心玩法', value: '生存 / 建筑 / 红石 / 冒险' },
              { label: '核心机制', value: '原版基底 + 实用插件' },
              { label: '主题色', value: '#4C93FB MTN 蓝' },
              { label: '目标', value: '值得被记住的稳定世界' }
            ]
          }
        ]
      },
      cta: {
        kicker: '立即加入我们',
        title: 'QQ群号：1064494318',
        body: '加入 MTN 土豆服务器，在长期运营的公益生存世界里建筑、生电、冒险、交易，并留下属于你的欢乐。',
        primaryCta: '复制QQ群号',
        copiedLabel: '已复制QQ群号',
        qqGroup: '1064494318',
        siteCta: '访问文档',
        siteUrl: 'https://docs.mcmtn.net/',
        note: '文档：https://docs.mcmtn.net/'
      },
      floatingJoin: {
        label: '加入服务器',
        detail: 'QQ群 1064494318',
        href: 'https://qm.qq.com/q/yLvHNwPXZ8'
      },
      footer: {
        tagline: '一个家用土豆机器上的公益服务器，服务于原版优先的生存、共同建筑和长久记忆。',
        navTitle: '探索',
        supportTitle: '支持',
        links: [
          { label: '最新动态', href: '#news', external: true },
          { label: '特色玩法', href: '#features', external: true },
          { label: '画廊', href: '/gallery' },
          { label: '玩家', href: '/players' }
        ],
        legal: 'MTNetwork 是独立 Minecraft 社区，与 Mojang 或 Microsoft 没有关联。'
      }
    },
    serverPanels: {
      liveStatusTitle: '实时服务器状态',
      liveStatusPanelTitle: '节点负载',
      liveStatusFallback: '正在等待实时遥测数据...',
      liveStatusRefresh: '刷新频率：每 5 秒',
      operational: '运行中',
      standby: '待机',
      historyTitle: '每周在线热力图',
      historyHint: '基于实时采样器生成的最近 7 天逐小时在线情况。',
      historyLoading: '正在加载最近在线历史...',
      historyEmpty: '暂无活动',
      historyPeak: '峰值 {peak}',
      connectionTitle: '连接指南',
      connectionPanelTitle: '连接路由',
      connectionLoading: '正在加载连接记录...',
      connectionHint: '如果设备和网络都支持，优先使用 IPv6。',
      javaTitle: 'Java 版',
      bedrockTitle: '基岩版',
      copy: '复制',
      copied: '已复制',
      lastUpdate: '最后更新',
      labels: {
        platform: '平台',
        cpu: 'CPU',
        load: '负载',
        ram: '内存',
        offline: '离线'
      }
    },
    players: {
      kicker: '公开玩家档案',
      title: '玩家目录',
      body: '搜索玩家名册、浏览历史档案，或随机跳转到一份玩家记录。',
      controlsKicker: '名册查询',
      controlsNote: '搜索实时档案，或在近期视图与完整历史之间切换。',
      searchLabel: '搜索',
      loading: '正在加载最新玩家记录...',
      emptyTitle: '没有玩家匹配当前搜索',
      emptyBody: '试试不同的名字片段，或返回近期玩家视图。',
      reset: '重置筛选',
      searchPlaceholder: '按玩家名或 UUID 搜索',
      recentLabel: '近期',
      recentDetail: '近 {days} 天',
      allLabel: '档案',
      allDetail: '完整记录',
      randomLabel: '随机',
      randomTitle: '随机打开一份玩家记录',
      onlineNow: '当前在线',
      firstSeenFallback: '未知玩家',
      visibleInLiveList: '当前出现在实时在线列表中。',
      resultMatching: '匹配 “{query}”',
      resultActiveWindow: '最近 {days} 天内活跃',
      resultArchive: '完整档案',
      lastSeen: '最后在线 {value}',
      lastSeenUnavailable: '暂无最后在线时间'
    },
    playerDetail: {
      loading: '正在加载玩家记录...',
      emptyTitle: '无法获取玩家记录',
      emptyBody: '请求的 UUID 在公开档案中没有完整记录。',
      back: '返回玩家列表',
      header: {
        kicker: '公开玩家档案',
        liveBody: '该玩家当前出现在服务器状态的实时在线信息中。',
        archiveBody: '这里展示这名玩家的历史记录与进度概览。'
      },
      profileCardTitle: '档案',
      tabsAria: '玩家详情分区',
      tabs: {
        overview: '概览',
        advancements: '进度',
        statistics: '统计'
      },
      summary: {
        onlineNow: '当前在线',
        archiveRecord: '档案记录',
        linkedTo: '关联到',
        playtime: '游玩时间',
        advancements: '进度',
        skillTotal: 'McMMO 总等级',
        skillLeaderboard: '技能榜排名',
        xpLevel: '经验等级',
        ranks: '排行榜名次',
        noSkillData: '暂无公开技能数据',
        noHistoryData: '暂无最近在线采样',
        unranked: '未上榜'
      },
      profile: {
        firstJoin: '首次加入',
        lastSeen: '最后在线',
        playtime: '游玩时间',
        xpLevel: '经验等级',
        linkedTo: '关联账号',
        qqOwner: 'QQ'
      },
      sections: {
        ores: '矿物开采记录',
        onlineHistory: '每周在线记录',
        skills: 'McMMO 技能',
        advancements: '已完成进度',
        extendedStats: '扩展统计',
        total: '总计',
        top: '前 {count} 项',
        searchPlaceholder: '搜索统计，例如 diamond 或 walk',
        emptyStats: '当前分类与搜索条件下没有匹配的统计数据。'
      },
      timezoneFallback: '本地'
    },
    heatmap: {
      online: '在线',
      offline: '离线',
      weekdays: {
        mon: '周一',
        tue: '周二',
        wed: '周三',
        thu: '周四',
        fri: '周五',
        sat: '周六',
        sun: '周日',
        monday: '星期一',
        tuesday: '星期二',
        wednesday: '星期三',
        thursday: '星期四',
        friday: '星期五',
        saturday: '星期六',
        sunday: '星期日'
      }
    },
    leaderboards: {
      title: '排行榜',
      body: '基于实时后端数据生成的排名。',
      kicker: '排行榜',
      loading: '正在刷新排行榜数据...',
      visibleEntries: '当前显示 {count} 条记录',
      emptyTitle: '暂无排行榜数据',
      emptyBody: '这个分类目前没有足够的公开数据来生成排名。',
      close: '关闭',
      paginationAria: '排行榜分页',
      paginationPrev: '上一页',
      paginationNext: '下一页',
      paginationStatus: '第 {page} / {total} 页',
      columns: {
        rank: '排名',
        player: '玩家',
        score: '分数'
      }
    },
    formats: {
      playerCount: '{count} 名玩家',
      playerSingular: '玩家',
      playerPlural: '玩家',
      leaderboardTitle: '{label} 排行榜',
      durationHoursShort: '{value} 小时',
      durationHoursWide: '{value} 小时',
      durationMinutesShort: '{value} 分钟',
      distanceKilometers: '{value} 公里',
      distanceMeters: '{value} 米'
    },
    gallery: {
      eyebrow: '世界档案',
      title: '服务器中的定格瞬间',
      body: '以时间顺序记录世界变化的影像日志。',
      close: '关闭',
      frameLabel: '画面',
      previousImage: '上一张图片',
      nextImage: '下一张图片'
    },
    wikiRedirect: {
      kicker: '百科',
      title: '正在跳转到 MTNetwork 百科。',
      body: '如果没有自动跳转，请使用下面的链接。',
      cta: '打开百科'
    },
    teamPlaceholder: {
      eyebrow: '占位页面',
      title: '这个页面还在整理中。',
      body: '页面入口已经保留，等内容准备完成后会直接接入。',
      primaryCta: '打开百科',
      secondaryCta: '返回首页',
      note: '当前先保留路由与布局，避免后续上线时产生额外改动。'
    }
  },
  leaderboardLabels: {
    skills: 'McMMO 总等级',
    playtime: '游玩时间',
    mining: '开采方块数',
    killing: '击杀生物数',
    deaths: '死亡次数',
    walking: '步行距离',
    pvp: '玩家击杀数'
  },
  playerStatGroups: [
    { name: '全局统计', categories: ['minecraft:custom'] },
    {
      name: '物品与方块',
      categories: [
        'minecraft:mined',
        'minecraft:broken',
        'minecraft:crafted',
        'minecraft:used',
        'minecraft:picked_up',
        'minecraft:dropped'
      ]
    },
    { name: '战斗与生物', categories: ['minecraft:killed', 'minecraft:killed_by'] }
  ],
  statCategories: {
    custom: '全局',
    mined: '挖掘',
    broken: '损坏',
    crafted: '合成',
    used: '使用',
    picked_up: '拾取',
    dropped: '丢弃',
    killed: '击杀',
    killed_by: '被击杀'
  },
  statTranslations: {
    play_one_minute: '游玩时间',
    play_time: '游玩时间',
    jump: '跳跃次数',
    damage_dealt: '造成伤害',
    damage_taken: '受到伤害',
    deaths: '死亡次数',
    mob_kills: '生物击杀数',
    player_kills: '玩家击杀数',
    walk_one_cm: '步行距离',
    sprint_one_cm: '疾跑距离',
    fly_one_cm: '飞行距离',
    climb_one_cm: '攀爬距离',
    fall_one_cm: '下落距离',
    minecart_one_cm: '矿车移动距离',
    boat_one_cm: '船只移动距离',
    pig_one_cm: '骑猪移动距离',
    horse_one_cm: '骑马移动距离',
    strider_one_cm: '炽足兽移动距离',
    aviate_one_cm: '鞘翅飞行距离',
    swim_one_cm: '游泳距离',
    walk_on_water_one_cm: '水上行走距离',
    walk_under_water_one_cm: '水下行走距离',
    time_since_death: '距上次死亡时间',
    time_since_rest: '距上次睡眠时间',
    sneak_time: '潜行时间',
    total_world_time: '世界总时间',
    leave_game: '退出游戏次数',
    dropped: '丢弃物品数',
    interact_with_beacon: '信标交互次数',
    inspect_hopper: '漏斗查看次数',
    interact_with_blast_furnace: '高炉交互次数',
    interact_with_smoker: '烟熏炉交互次数',
    interact_with_camp_fire: '营火交互次数',
    interact_with_campfire: '营火交互次数',
    talked_to_villager: '与村民交谈次数',
    traded_with_villager: '与村民交易次数',
    fish_caught: '钓鱼次数',
    sleep_in_bed: '睡觉次数',
    raid_win: '赢得袭击次数',
    raid_trigger: '触发袭击次数',
    trigger_trapped_chest: '触发陷阱箱次数',
    damage_absorbed: '吸收伤害量',
    interact_with_furnace: '熔炉交互次数',
    crouch_one_cm: '潜行距离',
    interact_with_stonecutter: '切石机交互次数',
    damage_resisted: '抵抗伤害量',
    damage_blocked_by_shield: '盾牌格挡伤害量',
    interact_with_crafting_table: '工作台交互次数',
    inspect_dropper: '投掷器查看次数',
    target_hit: '命中靶心次数',
    fill_cauldron: '装满炼药锅次数',
    interact_with_grindstone: '砂轮交互次数',
    open_shulker_box: '潜影盒开启次数',
    open_enderchest: '末影箱开启次数',
    damage_dealt_absorbed: '造成伤害（被吸收）',
    interact_with_brewingstand: '酿造台交互次数',
    inspect_dispenser: '发射器查看次数',
    interact_with_loom: '织布机交互次数',
    play_noteblock: '音符盒播放次数',
    interact_with_lectern: '讲台交互次数',
    drop: '丢出物品数',
    use_cauldron: '使用炼药锅次数',
    bell_ring: '敲钟次数',
    open_barrel: '木桶开启次数',
    interact_with_cartography_table: '制图台交互次数',
    open_chest: '箱子开启次数',
    tune_noteblock: '音符盒调音次数',
    interact_with_anvil: '铁砧交互次数',
    animals_bred: '繁殖动物次数',
    play_record: '唱片播放次数',
    interact_with_smithing_table: '锻造台交互次数',
    pot_flower: '插花次数',
    clean_banner: '清洗旗帜次数',
    enchant_item: '附魔次数',
    clean_armor: '清洗盔甲次数',
    clean_shulker_box: '清洗潜影盒次数',
    eat_cake_slice: '吃蛋糕片次数',
    interact_with_beehive: '蜂箱交互次数',
    interact_with_lodestone: '磁石交互次数',
    interact_with_respawn_anchor: '重生锚交互次数'
  },
  skillLabels: {
    taming: '驯兽',
    mining: '采矿',
    woodcutting: '伐木',
    repair: '修理',
    unarmed: '徒手',
    herbalism: '草药学',
    excavation: '挖掘',
    archery: '箭术',
    swords: '剑术',
    axes: '斧术',
    acrobatics: '杂技',
    fishing: '钓鱼',
    alchemy: '炼药',
    crossbows: '弩术',
    tridents: '三叉戟',
    maces: '狼牙棒',
    spears: '长矛'
  },
  platformLabels: {
    Java: 'Java 版',
    Bedrock: '基岩版'
  },
  oreLabels: {
    quartz: '石英',
    coal: '煤炭',
    redstone: '红石',
    iron: '铁',
    copper: '铜',
    diamond: '钻石',
    gold: '金',
    lapis: '青金石',
    emerald: '绿宝石',
    nether_gold_ore: '下界金',
    ancient_debris: '远古残骸'
  },
  advancementTypes: {
    task: '任务',
    goal: '目标',
    challenge: '挑战'
  },
  advancementCategories: {
    adventure: '冒险',
    husbandry: '农牧',
    minecraft: 'Minecraft',
    nether: '下界',
    story: '主线',
    end: '末地',
    others: '其他'
  },
  advancementNames: {
    'minecraft:adventure/root': '冒险',
    'minecraft:adventure/sleep_in_bed': '甜蜜的梦',
    'minecraft:adventure/kill_a_mob': '怪物猎人',
    'minecraft:adventure/shoot_arrow': '精准打击',
    'minecraft:adventure/sniper_duel': '狙击对决',
    'minecraft:adventure/ol_betsy': '老伙计',
    'minecraft:adventure/whos_the_pillager_now': '现在谁才是掠夺者？',
    'minecraft:adventure/trade': '这买卖不赖！',
    'minecraft:adventure/trade_at_world_height': '星空交易者',
    'minecraft:adventure/summon_iron_golem': '雇佣帮手',
    'minecraft:adventure/bullseye': '正中靶心',
    'minecraft:adventure/read_power_of_chiseled_bookshelf': '书本的力量',
    'minecraft:adventure/two_birds_one_arrow': '一箭双雕',
    'minecraft:adventure/arbalistic': '弩箭齐发',
    'minecraft:adventure/fall_from_world_height': '洞穴与山崖',
    'minecraft:adventure/play_jukebox_in_meadows': '音乐之声',
    'minecraft:adventure/brush_armadillo': '鳞甲真可爱',
    'minecraft:adventure/salvage_sherd': '尊重遗迹',
    'minecraft:adventure/craft_decorated_pot_using_only_sherds': '小心修复',
    'minecraft:adventure/trim_with_any_armor_pattern': '打造新造型',
    'minecraft:adventure/very_very_frightening': '电闪雷鸣',
    'minecraft:adventure/totem_of_undying': '超越生死',
    'minecraft:adventure/voluntary_exile': '自找放逐',
    'minecraft:adventure/hero_of_the_village': '村庄英雄',
    'minecraft:adventure/honey_block_slide': '像滑冰一样',
    'minecraft:adventure/walk_on_powder_snow_with_leather_boots': '轻装踏雪',
    'minecraft:adventure/lightning_rod_with_villager_no_fire': '避雷有术',
    'minecraft:adventure/spyglass_at_parrot': '鹦鹉学舌',
    'minecraft:adventure/spyglass_at_ghast': '这也太远了吧',
    'minecraft:adventure/spyglass_at_dragon': '这就是尽头',
    'minecraft:adventure/fall_from_world_height_and_survive': '高处不胜寒',
    'minecraft:adventure/kill_mob_near_sculk_catalyst': '生日歌',
    'minecraft:adventure/avoid_vibration': '安静潜行者',
    'minecraft:adventure/salvage_sherd_from_pot': '完璧归赵',
    'minecraft:adventure/crafters_crafting_crafters': '套娃工匠',
    'minecraft:end/root': '末地',
    'minecraft:end/kill_dragon': '解放末地',
    'minecraft:end/dragon_egg': '下一世代',
    'minecraft:end/enter_end_gateway': '远程传送门',
    'minecraft:end/respawn_dragon': '再次开始',
    'minecraft:end/find_end_city': '天空之城',
    'minecraft:end/elytra': '自由如风',
    'minecraft:end/levitate': '极限漂浮',
    'minecraft:husbandry/root': '农牧',
    'minecraft:husbandry/safely_harvest_honey': '蜜蜂我们的朋友',
    'minecraft:husbandry/breed_an_animal': '鹦鹉学鸡叫',
    'minecraft:husbandry/balanced_diet': '均衡饮食',
    'minecraft:husbandry/obtain_netherite_hoe': '认真过头的锄头',
    'minecraft:husbandry/tame_an_animal': '最好的朋友永远系着牵引绳',
    'minecraft:husbandry/fishy_business': '渔获',
    'minecraft:husbandry/tactical_fishing': '战术捕鱼',
    'minecraft:husbandry/complete_catalogue': '完美收藏',
    'minecraft:husbandry/plant_seed': '播种之时',
    'minecraft:husbandry/bred_all_animals': '双人成行',
    'minecraft:husbandry/silk_touch_nest': '丝滑采蜜',
    'minecraft:husbandry/ride_a_boat_with_a_goat': '同船山羊',
    'minecraft:husbandry/make_a_sign_glow': '熠熠生辉',
    'minecraft:husbandry/allay_deliver_item_to_player': '有求必应',
    'minecraft:husbandry/leash_all_frog_variants': '青蛙大集合',
    'minecraft:husbandry/froglights': '蛙明灯',
    'minecraft:husbandry/wax_on': '上蜡',
    'minecraft:husbandry/wax_off': '除蜡',
    'minecraft:husbandry/obtain_sniffer_egg': '闻闻新朋友',
    'minecraft:husbandry/feed_snifflet': '花样喂养',
    'minecraft:husbandry/plant_any_sniffer_seed': '古老的种子',
    'minecraft:nether/root': '下界',
    'minecraft:nether/return_to_sender': '奉还于你',
    'minecraft:nether/find_bastion': '并不是“曾经是”',
    'minecraft:nether/obtain_ancient_debris': '深藏不露',
    'minecraft:nether/fast_travel': '引路下界',
    'minecraft:nether/find_fortress': '可怕的堡垒',
    'minecraft:nether/loot_bastion': '战利品收藏家',
    'minecraft:nether/use_lodestone': '指南针重定位',
    'minecraft:nether/ride_strider': '翻滚的炽足兽',
    'minecraft:nether/obtain_crying_obsidian': '谁在切洋葱？',
    'minecraft:nether/distract_piglin': '闪亮亮！',
    'minecraft:nether/explore_nether': '热门景点',
    'minecraft:nether/summon_wither': '凋零之召',
    'minecraft:nether/obtain_blaze_rod': '烈焰棒到手',
    'minecraft:nether/create_beacon': '带回信标',
    'minecraft:nether/brew_potion': '本地酿造厂',
    'minecraft:nether/all_potions': '本地酿造师',
    'minecraft:nether/create_full_beacon': '信标工程师',
    'minecraft:nether/charge_respawn_anchor': '不只是复活点',
    'minecraft:nether/ride_strider_in_overworld_lava': '不适合这里',
    'minecraft:nether/uneasy_alliance': '艰难同盟',
    'minecraft:nether/netherite_armor': '全副下界合金',
    'minecraft:story/root': 'Minecraft',
    'minecraft:story/mine_stone': '石器时代',
    'minecraft:story/upgrade_tools': '得到升级',
    'minecraft:story/smelt_iron': '来硬的',
    'minecraft:story/obtain_armor': '穿戴整齐',
    'minecraft:story/lava_bucket': '火热话题',
    'minecraft:story/iron_tools': '不止是木头',
    'minecraft:story/deflect_arrow': '见招拆招',
    'minecraft:story/form_obsidian': '冰桶挑战',
    'minecraft:story/mine_diamond': '钻石！',
    'minecraft:story/enter_the_nether': '深入火海',
    'minecraft:story/shiny_gear': '闪亮装备',
    'minecraft:story/enchant_item': '附魔时刻',
    'minecraft:story/cure_zombie_villager': '僵尸医生',
    'minecraft:story/follow_ender_eye': '照着眼睛走',
    'minecraft:story/enter_the_end': '终局之地'
  }
} as const

export default zhCN
