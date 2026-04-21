import { leaderboardLabels } from '@/content/siteContent'
import type { FixedLeaderboardType, LeaderboardTarget, McMMOSkillKey } from '@/types/api'

export function formatFixedLeaderboardValue(key: FixedLeaderboardType, value: number): string {
  if (key === 'playtime') {
    return `${(value / 20 / 3600).toFixed(1)} hrs`
  }

  if (key === 'walking') {
    const kilometers = value / 100000
    return kilometers >= 1 ? `${kilometers.toFixed(1)} km` : `${(value / 100).toFixed(0)} m`
  }

  return value.toLocaleString()
}

export function createFixedLeaderboardTarget(key: FixedLeaderboardType): LeaderboardTarget {
  return {
    key,
    title: `${leaderboardLabels[key]} Leaderboard`,
    scoreLabel: leaderboardLabels[key],
    formatValue: (value: number) => formatFixedLeaderboardValue(key, value)
  }
}

export function createMcmmoLeaderboardTarget(skill: McMMOSkillKey, label: string): LeaderboardTarget {
  return {
    key: `mcmmo:${skill}`,
    title: `${label} Leaderboard`,
    scoreLabel: label,
    formatValue: (value: number) => value.toLocaleString()
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
    title: `${label} Leaderboard`,
    scoreLabel: label,
    formatValue
  }
}
