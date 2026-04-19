import { type Ref } from 'vue'

import { CUSTOM_STAT_TRANSLATIONS } from '@/utils/statTranslations'
import iconMap from '@/assets/icon_map.json'
import type { FormattedSkillEntry, McMMOSkills, PlayerStatBuckets } from '@/types/api'
import { MCMMO_SKILL_KEYS } from '@/types/api'

const typedIconMap = iconMap as Record<string, string>

export function usePlayerStats(stats: Ref<PlayerStatBuckets | null>) {
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
    if (value >= 1_000_000) {
      return `${(value / 1_000_000).toFixed(1)}M`
    }

    if (value >= 1_000) {
      return `${(value / 1_000).toFixed(1)}K`
    }

    return value.toLocaleString()
  }

  const formatStatName = (name: string): string => {
    const id = name.replace('minecraft:', '')

    if (id === 'custom') {
      return 'Global'
    }

    return CUSTOM_STAT_TRANSLATIONS[id] ?? id.replace(/_/g, ' ')
  }

  const formatStatValue = (value: number, name: string): string => {
    const id = name.replace('minecraft:', '')

    if (id.includes('time') || id.includes('one_minute') || id.includes('since')) {
      const hours = value / 20 / 3600
      if (hours > 1) {
        return `${hours.toFixed(1)}h`
      }

      return `${(value / 20 / 60).toFixed(0)}m`
    }

    if (id.endsWith('_one_cm')) {
      const kilometers = value / 100000
      if (kilometers > 1) {
        return `${kilometers.toFixed(2)}km`
      }

      return `${(value / 100).toFixed(1)}m`
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
        return Boolean(CUSTOM_STAT_TRANSLATIONS[id])
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
        label: key.charAt(0).toUpperCase() + key.slice(1),
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
