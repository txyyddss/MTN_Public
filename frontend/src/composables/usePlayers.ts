import { computed, shallowRef, toValue, watch, type MaybeRefOrGetter } from 'vue'

import { API_BASE_URL } from '@/config'
import { fetchWithCache } from '@/utils/dataCache'
import type { PlayerInfo, PlayerListResponse } from '@/types/api'

export function usePlayers(onlinePlayerIds: MaybeRefOrGetter<string[]> = []) {
  const players = shallowRef<PlayerInfo[]>([])
  const count = shallowRef(0)
  const activeDays = shallowRef(0)
  const searchQuery = shallowRef('')
  const showAll = shallowRef(false)
  const loading = shallowRef(true)

  const sortedPlayers = computed(() => {
    const online = new Set(toValue(onlinePlayerIds))

    return [...players.value].sort((left, right) => {
      const leftOnline = online.has(left.uuid)
      const rightOnline = online.has(right.uuid)

      if (leftOnline !== rightOnline) {
        return leftOnline ? -1 : 1
      }

      return right.last_seen - left.last_seen
    })
  })

  async function fetchPlayers(): Promise<void> {
    loading.value = true

    try {
      const params = new URLSearchParams()

      if (searchQuery.value) {
        params.append('search', searchQuery.value)
      }

      if (showAll.value) {
        params.append('all', 'true')
      }

      const url = `${API_BASE_URL}/api/players?${params.toString()}`
      const json = await fetchWithCache<PlayerListResponse>(url)

      players.value = json.players ?? []
      count.value = json.count ?? 0
      activeDays.value = json.active_days ?? 0
    } catch (error) {
      console.error('Failed to fetch players', error)
    } finally {
      loading.value = false
    }
  }

  let refreshInterval: ReturnType<typeof setInterval> | null = null

  function startAutoRefresh(ms = 30000): void {
    if (refreshInterval) {
      stopAutoRefresh()
    }

    refreshInterval = setInterval(() => {
      if (!searchQuery.value) {
        void fetchPlayers()
      }
    }, ms)
  }

  function stopAutoRefresh(): void {
    if (!refreshInterval) {
      return
    }

    clearInterval(refreshInterval)
    refreshInterval = null
  }

  let searchTimeout: ReturnType<typeof setTimeout> | null = null

  watch(searchQuery, () => {
    if (searchTimeout) {
      clearTimeout(searchTimeout)
    }

    searchTimeout = setTimeout(() => {
      void fetchPlayers()
    }, 300)
  })

  watch(showAll, () => {
    void fetchPlayers()
  })

  return {
    players,
    count,
    activeDays,
    searchQuery,
    showAll,
    loading,
    sortedPlayers,
    fetchPlayers,
    startAutoRefresh,
    stopAutoRefresh
  }
}
