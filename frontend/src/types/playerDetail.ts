import type { LeaderboardTarget } from '@/types/api'

export type PlayerDetailTab = 'overview' | 'advancements' | 'statistics'

export interface PlayerDetailTabOption {
  value: PlayerDetailTab
  label: string
}

export interface PlayerRankHighlight {
  label: string
  value: string
  target: LeaderboardTarget
}
