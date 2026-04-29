export type PlayerPlatform = 'Java' | 'Bedrock' | string

export interface EditionStatus {
  online: boolean
  players: number
}

export interface ConnectionLatency {
  online: boolean
  latency: string
  latency_ms: number
}

export interface SystemStats {
  platform: string
  cpu_model: string
  mem_used: number
  mem_percent: number
  load_1: number
  load_5: number
  load_15: number
}

export interface StatusResponse {
  java: EditionStatus | null
  bedrock: EditionStatus | null
  system: SystemStats | null
  online_players: string[]
  updated: string
}

export interface ConnectionAddress {
  domain?: string
  ip: string
  port?: string
  type: string
}

export interface ConnectionInfo {
  java_ipv4?: ConnectionAddress
  bedrock_ipv4?: ConnectionAddress
  java_ipv6?: ConnectionAddress
  bedrock_ipv6?: ConnectionAddress
  java_srv?: ConnectionAddress
}

export interface ConnectionAddresses {
  java_ipv6: string
  bedrock_ipv6: string
  java_ipv4_srv: string
  java_ipv6_srv: string
}

export interface ConnectionResponse {
  connection: ConnectionInfo | null
  addresses: ConnectionAddresses
}

export interface HeatmapDay {
  date: string
  weekday: string
}

export interface PlayerOnlineHeatmap {
  timezone: string
  days: HeatmapDay[]
  hours: number[]
  cells: boolean[][]
}

export interface PlayerInfo {
  uuid: string
  xp_level: number
  ticks_lived: number
  last_seen: number
  first_played: number
  last_known_name: string
  type: PlayerPlatform
}

export type PlayerStatBuckets = Record<string, Record<string, number>>

export interface PlayerStatsPayload {
  uuid: string
  stats: PlayerStatBuckets
}

export interface PlayerAdvancement {
  key: string
  done: boolean
}

export interface PlayerAdvancementsPayload {
  advancements: PlayerAdvancement[]
}

export interface LinkedAccount {
  bedrock_uuid: string
  bedrock_username?: string
  java_uuid: string
  java_username: string
}

export interface WhitelistAccount {
  edition: WhitelistEdition
  player_name: string
  qq_id_masked: string
}

export interface OreStat {
  name: string
  mined: number
  used: number
}

export const MCMMO_SKILL_KEYS = [
  'taming',
  'mining',
  'woodcutting',
  'repair',
  'unarmed',
  'herbalism',
  'excavation',
  'archery',
  'swords',
  'axes',
  'acrobatics',
  'fishing',
  'alchemy',
  'crossbows',
  'tridents',
  'maces',
  'spears'
] as const

export type McMMOSkillKey = (typeof MCMMO_SKILL_KEYS)[number]

export interface McMMOSkills {
  user_id: number
  user: string
  uuid: string
  taming: number
  mining: number
  woodcutting: number
  repair: number
  unarmed: number
  herbalism: number
  excavation: number
  archery: number
  swords: number
  axes: number
  acrobatics: number
  fishing: number
  alchemy: number
  crossbows: number
  tridents: number
  maces: number
  spears: number
  total: number
}

export interface PlayerDetailResponse {
  info: PlayerInfo | null
  stats: PlayerStatsPayload | null
  advancements: PlayerAdvancementsPayload | null
  mcmmo: McMMOSkills | null
  linked_account: LinkedAccount | null
  whitelist_account: WhitelistAccount | null
  ore_stats: OreStat[]
  ranks: Record<string, number>
  online_heatmap?: PlayerOnlineHeatmap | null
}

export interface PlayerListResponse {
  players: PlayerInfo[]
  count: number
  active_days?: number
}

export interface RandomPlayerResponse {
  uuid: string
}

export type FixedLeaderboardType =
  | 'skills'
  | 'playtime'
  | 'mining'
  | 'killing'
  | 'deaths'
  | 'walking'
  | 'pvp'

export type LeaderboardType = FixedLeaderboardType

export type DynamicLeaderboardType = `mcmmo:${McMMOSkillKey}` | `stat:${string}:${string}`

export type LeaderboardKey = FixedLeaderboardType | DynamicLeaderboardType

export interface LeaderboardTarget {
  key: LeaderboardKey
  title: string
  scoreLabel: string
  formatValue?: (value: number) => string
}

export interface LeaderboardEntry {
  uuid: string
  name: string
  value: number
  rank: number
}

export interface LeaderboardResponse {
  type: LeaderboardKey
  entries: LeaderboardEntry[]
  count: number
}

export interface StatGroup {
  name: string
  categories: string[]
}

export interface AdvancementMetadata {
  name: string
  description: string
  icon: string
  type: 'task' | 'goal' | 'challenge'
}

export interface FormattedSkillEntry {
  key: McMMOSkillKey
  label: string
  level: number
}

export type WhitelistEdition = 'java' | 'bedrock'
export type WhitelistEditionFilter = WhitelistEdition | 'all'
export type WhitelistAction = 'add' | 'remove'

export interface WhitelistEntry {
  id: number
  qq_id?: string
  edition: WhitelistEdition
  player_name: string
  normalized_player_name: string
  active: boolean
  created_at: string
  updated_at: string
  removed_at?: string
}

export interface WhitelistListResponse {
  entries: WhitelistEntry[]
  count: number
}

export interface WhitelistMutationInput {
  edition: WhitelistEdition
  playerName: string
  qqId?: string
}

export interface WhitelistOperationResponse {
  entry?: WhitelistEntry
  rcon_output?: string
  message: string
  changed: boolean
  quota?: {
    used: number
    limit: number
    remaining: number
    exempt?: boolean
  }
}
