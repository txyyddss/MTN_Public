import { ref, computed, watch } from 'vue'
import { API_BASE_URL } from '@/config'
import { fetchWithCache } from '@/utils/dataCache'

export interface PlayerInfo {
    uuid: string
    last_known_name: string
    last_seen: number
    type: string
}

export function usePlayers() {
    const players = ref<PlayerInfo[]>([])
    const count = ref(0)
    const activeDays = ref(0)
    const searchQuery = ref('')
    const showAll = ref(false)
    const loading = ref(true)
    const onlinePlayers = ref<string[]>([])

    const sortedPlayers = computed(() => {
        return [...players.value].sort((a, b) => {
            const aOnline = onlinePlayers.value.includes(a.uuid)
            const bOnline = onlinePlayers.value.includes(b.uuid)
            if (aOnline && !bOnline) return -1
            if (!aOnline && bOnline) return 1
            return 0
        })
    })

    const fetchPlayers = async () => {
        loading.value = true
        try {
            const params = new URLSearchParams()
            if (searchQuery.value) params.append('search', searchQuery.value)
            if (showAll.value) params.append('all', 'true')

            const url = `${API_BASE_URL}/api/players?${params.toString()}`
            const json = await fetchWithCache(url)
            players.value = json.players || []
            count.value = json.count || 0
            activeDays.value = json.active_days || 0

            // Preload details for fetched players
            if (players.value.length > 0) {
                const { preloadData } = await import('@/utils/preloader')
                const detailUrls = players.value.map(p => `${API_BASE_URL}/api/players/${p.uuid}`)
                preloadData(detailUrls)
            }
        } catch (e) {
            console.error('Failed to fetch players:', e)
        } finally {
            loading.value = false
        }
    }

    const fetchOnline = async () => {
        try {
            const json = await fetchWithCache(`${API_BASE_URL}/api/status`)
            onlinePlayers.value = json.online_players || []
        } catch (e) { }
    }

    let refreshInterval: ReturnType<typeof setInterval> | null = null

    const startAutoRefresh = (ms = 5000) => {
        if (refreshInterval) stopAutoRefresh()
        refreshInterval = setInterval(() => {
            fetchOnline()
            // Only refresh full player list if not searching to avoid UI jumps while typing
            if (!searchQuery.value) {
                fetchPlayers()
            }
        }, ms)
    }

    const stopAutoRefresh = () => {
        if (refreshInterval) {
            clearInterval(refreshInterval)
            refreshInterval = null
        }
    }

    let searchTimeout: ReturnType<typeof setTimeout> | null = null
    watch(searchQuery, () => {
        if (searchTimeout) clearTimeout(searchTimeout)
        searchTimeout = setTimeout(fetchPlayers, 300)
    })

    watch(showAll, fetchPlayers)

    return {
        players,
        count,
        activeDays,
        searchQuery,
        showAll,
        loading,
        onlinePlayers,
        sortedPlayers,
        fetchPlayers,
        fetchOnline,
        startAutoRefresh,
        stopAutoRefresh
    }
}
