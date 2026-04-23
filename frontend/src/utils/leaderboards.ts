import {
  formatDistanceKilometers,
  formatDistanceMeters,
  formatDurationHoursWide,
  formatLeaderboardTitle,
  getLeaderboardLabel,
  getLocaleValue
} from '@/content/siteContent'
import type { FixedLeaderboardType, LeaderboardTarget, McMMOSkillKey } from '@/types/api'

export function formatFixedLeaderboardValue(key: FixedLeaderboardType, value: number): string {
  if (key === 'playtime') {
    return formatDurationHoursWide((value / 20 / 3600).toFixed(1))
  }

  if (key === 'walking') {
    const kilometers = value / 100000
    return kilometers >= 1 ? formatDistanceKilometers(kilometers.toFixed(1)) : formatDistanceMeters((value / 100).toFixed(0))
  }

  return value.toLocaleString(getLocaleValue())
}

export function createFixedLeaderboardTarget(key: FixedLeaderboardType): LeaderboardTarget {
  const label = getLeaderboardLabel(key)

  return {
    key,
    title: formatLeaderboardTitle(label),
    scoreLabel: label,
    formatValue: (value: number) => formatFixedLeaderboardValue(key, value)
  }
}

export function createMcmmoLeaderboardTarget(skill: McMMOSkillKey, label: string): LeaderboardTarget {
  return {
    key: `mcmmo:${skill}`,
    title: formatLeaderboardTitle(label),
    scoreLabel: label,
    formatValue: (value: number) => value.toLocaleString(getLocaleValue())
  }
}

export function createStatLeaderboardTarget(
  category: string,
  stat: string,
  label: string,
  formatValue: (value: number) => string
): LeaderboardTarget {
  return {
    key: `stat:${category}:${stat}`,
    title: formatLeaderboardTitle(label),
    scoreLabel: label,
    formatValue
  }
}
