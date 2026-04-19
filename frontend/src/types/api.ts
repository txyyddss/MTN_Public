export type PlayerPlatform = 'Java' | 'Bedrock' | string

export interface EditionStatus {
  online: boolean
  players: number
  player_list: string[]
  version: string
  motd: string
}

export interface ConnectionLatency {
  online: boolean
  latency: string
  latency_ms: number
}

export interface SystemStats {
  platform: string
  virtualization: string
  cpu_model: string
  cpu_percent: number
  mem_total: number
  mem_used: number
  mem_percent: number
  disk_read: number
  disk_write: number
  disk_read_rate: number
  disk_write_rate: number
  net_sent: number
  net_recv: number
  net_sent_rate: number
  net_recv_rate: number
  load_1: number
  load_5: number
  load_15: number
}

export interface StatusResponse {
  java: EditionStatus | null
  bedrock: EditionStatus | null
  system: SystemStats | null
  connections: Record<string, ConnectionLatency>
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

export interface PlayerInfo {
  uuid: string
  xp_level: number
  ticks_lived: number
  food_level: number
  health: number
  last_seen: number
  first_played: number
  last_known_name: string
  clean_name: string
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
  criteria: Record<string, string>
}

export interface PlayerAdvancementsPayload {
  uuid: string
  advancements: PlayerAdvancement[]
}

export interface LinkedAccount {
  bedrock_uuid: string
  bedrock_username?: string
  java_uuid: string
  java_username: string
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
  ore_stats: OreStat[]
  ranks: Record<string, number>
}

export interface PlayerListResponse {
  players: PlayerInfo[]
  count: number
  active_days?: number
}

export interface RandomPlayerResponse {
  uuid: string
}

export type LeaderboardType =
  | 'skills'
  | 'playtime'
  | 'mining'
  | 'killing'
  | 'deaths'
  | 'walking'
  | 'pvp'

export interface LeaderboardEntry {
  uuid: string
  name: string
  value: number
  rank: number
}

export interface LeaderboardResponse {
  type: LeaderboardType
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
