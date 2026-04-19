import { type Ref } from 'vue'
import { CUSTOM_STAT_TRANSLATIONS } from '@/utils/statTranslations'
import iconMap from '@/assets/icon_map.json'

export function usePlayerStats(stats: Ref<any>) {
    const fuzzyMatch = (text: string, query: string) => {
        let i = 0, j = 0;
        while (i < text.length && j < query.length) {
            if (text[i].toLowerCase() === query[j].toLowerCase()) j++;
            i++;
        }
        return j === query.length;
    }

    const formatNumber = (num: number) => {
        if (num >= 1000000) return (num / 1000000).toFixed(1) + 'M'
        if (num >= 1000) return (num / 1000).toFixed(1) + 'K'
        return num.toLocaleString()
    }

    const formatStatName = (name: string) => {
        const id = name.replace('minecraft:', '')
        if (id === 'custom') return 'Global'
        if (CUSTOM_STAT_TRANSLATIONS[id]) return CUSTOM_STAT_TRANSLATIONS[id]
        return id.replace(/_/g, ' ')
    }

    const formatStatValue = (val: number, name: string) => {
        const id = name.replace('minecraft:', '')
        if (id.includes('time') || id.includes('one_minute') || id.includes('since')) {
            const hours = val / 20 / 3600
            if (hours > 1) return hours.toFixed(1) + 'h'
            const mins = val / 20 / 60
            return mins.toFixed(0) + 'm'
        }
        if (id.endsWith('_one_cm')) {
            const km = val / 100000;
            if (km > 1) return km.toFixed(2) + 'km'
            const m = val / 100
            return m.toFixed(1) + 'm'
        }
        return formatNumber(val)
    }

    const getStatIconPath = (category: string, name: string) => {
        const id = name.replace('minecraft:', '')
        if (category === 'minecraft:custom') return undefined
        if ((iconMap as any)[id]) return (iconMap as any)[id]
        if (category === 'minecraft:killed' || category === 'minecraft:killed_by') {
            const eggId = `${id}_spawn_egg`
            if ((iconMap as any)[eggId]) return (iconMap as any)[eggId]
        }
        const cleanId = id.replace('_spawn_egg', '')
        if ((iconMap as any)[cleanId]) return (iconMap as any)[cleanId]
        if (category === 'minecraft:mined' || category === 'minecraft:broken') {
            return `/mc_icons/blocks/misc/${id}^32.png`
        }
        if (['minecraft:crafted', 'minecraft:used', 'minecraft:picked_up', 'minecraft:dropped'].includes(category)) {
            return `/mc_icons/items/${id}.png`
        }
        return undefined
    }

    const getFilteredStats = (category: string, query: string = '') => {
        if (!category || !stats.value || !stats.value[category]) return {}
        const pool = stats.value[category]
        let entries = Object.entries(pool)

        if (category === 'minecraft:custom') {
            entries = entries.filter(([name]) => {
                const id = name.replace('minecraft:', '')
                return !!CUSTOM_STAT_TRANSLATIONS[id]
            })
        }

        if (query) {
            const q = query.toLowerCase().trim()
            entries = entries.filter(([name]) => {
                const id = name.replace('minecraft:', '')
                const displayName = (id === 'custom' ? 'Global' : (CUSTOM_STAT_TRANSLATIONS[id] || id.replace(/_/g, ' ')))
                return fuzzyMatch(displayName, q) || fuzzyMatch(name, q)
            })
        }

        entries.sort((a, b) => (b[1] as number) - (a[1] as number))

        const result: Record<string, number> = {}
        for (const [name, val] of entries) {
            result[name] = val as number
        }
        return result
    }

    const getFilteredMcmmo = (mcmmo: any) => {
        if (!mcmmo) return {}
        let entries = Object.entries(mcmmo).filter(([key]) => !['user_id', 'user', 'uuid', 'total'].includes(key))
        entries.sort((a, b) => (b[1] as number) - (a[1] as number))
        const result: Record<string, any> = {}
        for (const [key, val] of entries) {
            result[key.charAt(0).toUpperCase() + key.slice(1)] = val
        }
        return result
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
