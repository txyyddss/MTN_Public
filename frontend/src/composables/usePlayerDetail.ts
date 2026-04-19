import { ref } from 'vue'
import { API_BASE_URL } from '@/config'
import { fetchWithCache } from '@/utils/dataCache'

export interface PlayerInfo {
    uuid: string
    last_known_name: string
    first_played: number
    last_seen: number
    ticks_lived: number
    xp_level: number
    type: string
}

export function usePlayerDetail(uuid: string) {
    const loading = ref(true)
    const info = ref<PlayerInfo | null>(null)
    const stats = ref<any>(null)
    const advancements = ref<any[] | null>(null)
    const mcmmo = ref<any>(null)
    const linkedAccount = ref<any>(null)
    const oreStats = ref<any[]>([])
    const ranks = ref<Record<string, number>>({})

    const fetchDetail = async () => {
        if (!uuid) return
        loading.value = true
        try {
            const json = await fetchWithCache(`${API_BASE_URL}/api/players/${uuid}`)
            info.value = json.info
            stats.value = json.stats?.stats || null
            advancements.value = json.advancements?.advancements || []
            mcmmo.value = json.mcmmo
            linkedAccount.value = json.linked_account
            oreStats.value = json.ore_stats
            ranks.value = json.ranks || {}
        } catch (e) {
            console.error('Failed to load player detail', e)
        } finally {
            loading.value = false
        }
    }

    return {
        loading,
        info,
        stats,
        advancements,
        mcmmo,
        linkedAccount,
        oreStats,
        ranks,
        fetchDetail
    }
}
