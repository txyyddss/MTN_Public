import { type Ref } from 'vue'

import iconMap from '@/assets/icon_map.json'
import {
  formatDistanceKilometers,
  formatDistanceMeters,
  formatDurationHoursShort,
  formatDurationMinutesShort,
  getLocaleValue,
  getSkillLabel,
  getStatCategoryLabel,
  getStatTranslation,
  useCurrentLocale
} from '@/content/siteContent'
import type { FormattedSkillEntry, McMMOSkills, PlayerStatBuckets } from '@/types/api'
import { MCMMO_SKILL_KEYS } from '@/types/api'

const typedIconMap = iconMap as Record<string, string>

export function usePlayerStats(stats: Ref<PlayerStatBuckets | null>) {
  const currentLocale = useCurrentLocale()

  const fuzzyMatch = (text: string, query: string): boolean => {
    let queryIndex = 0

    for (let textIndex = 0; textIndex < text.length && queryIndex < query.length; textIndex += 1) {
      if (text[textIndex].toLowerCase() === query[queryIndex].toLowerCase()) {
        queryIndex += 1
      }
    }

    return queryIndex === query.length
  }

  const formatNumber = (value: number): string => {
    return new Intl.NumberFormat(getLocaleValue(), {
      notation: value >= 10_000 ? 'compact' : 'standard',
      maximumFractionDigits: value >= 10_000 ? 1 : 0
    }).format(value)
  }

  const formatStatName = (name: string): string => {
    const id = name.replace('minecraft:', '')

    if (id === 'custom') {
      return getStatCategoryLabel('custom') ?? 'Global'
    }

    return getStatTranslation(id) ?? id.replace(/_/g, ' ')
  }

  const formatStatValue = (value: number, name: string): string => {
    const id = name.replace('minecraft:', '')
    void currentLocale.value

    if (id.includes('time') || id.includes('one_minute') || id.includes('since')) {
      const hours = value / 20 / 3600
      if (hours > 1) {
        return formatDurationHoursShort(hours.toFixed(1))
      }

      const minutes = value / 20 / 60
      return formatDurationMinutesShort(minutes.toFixed(0))
    }

    if (id.endsWith('_one_cm')) {
      const kilometers = value / 100000
      if (kilometers > 1) {
        return formatDistanceKilometers(kilometers.toFixed(2))
      }

      const meters = value / 100
      return formatDistanceMeters(meters.toFixed(1))
    }

    return formatNumber(value)
  }

  const getStatIconPath = (category: string, name: string): string | undefined => {
    const id = name.replace('minecraft:', '')

    if (category === 'minecraft:custom') {
      return undefined
    }

    if (typedIconMap[id]) {
      return typedIconMap[id]
    }

    if (category === 'minecraft:killed' || category === 'minecraft:killed_by') {
      const eggId = `${id}_spawn_egg`
      if (typedIconMap[eggId]) {
        return typedIconMap[eggId]
      }
    }

    const cleanId = id.replace('_spawn_egg', '')
    if (typedIconMap[cleanId]) {
      return typedIconMap[cleanId]
    }

    if (category === 'minecraft:mined' || category === 'minecraft:broken') {
      return `/mc_icons/blocks/misc/${id}^32.png`
    }

    if (['minecraft:crafted', 'minecraft:used', 'minecraft:picked_up', 'minecraft:dropped'].includes(category)) {
      return `/mc_icons/items/${id}.png`
    }

    return undefined
  }

  const getFilteredStats = (category: string, query = ''): Record<string, number> => {
    if (!category || !stats.value || !stats.value[category]) {
      return {}
    }

    let entries = Object.entries(stats.value[category])

    if (category === 'minecraft:custom') {
      entries = entries.filter(([name]) => {
        const id = name.replace('minecraft:', '')
        return Boolean(getStatTranslation(id))
      })
    }

    if (query.trim()) {
      const normalizedQuery = query.toLowerCase().trim()
      entries = entries.filter(([name]) => {
        const displayName = formatStatName(name)
        return fuzzyMatch(displayName, normalizedQuery) || fuzzyMatch(name, normalizedQuery)
      })
    }

    entries.sort((left, right) => right[1] - left[1])

    return Object.fromEntries(entries)
  }

  const getFilteredMcmmo = (mcmmo: McMMOSkills | null): FormattedSkillEntry[] => {
    if (!mcmmo) {
      return []
    }

    return [...MCMMO_SKILL_KEYS]
      .map((key) => ({
        key,
        label: getSkillLabel(key) ?? key.charAt(0).toUpperCase() + key.slice(1),
        level: mcmmo[key]
      }))
      .filter((entry) => entry.level > 0)
      .sort((left, right) => right.level - left.level)
  }

  return {
    formatNumber,
    formatStatName,
    formatStatValue,
    getStatIconPath,
    getFilteredStats,
    getFilteredMcmmo
  }
}
