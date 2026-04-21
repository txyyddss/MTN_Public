import { computed, ref, shallowRef, watch } from 'vue'

import { API_BASE_URL } from '@/config'
import { fetchWithCache } from '@/utils/dataCache'
import type { LeaderboardEntry, LeaderboardKey, LeaderboardResponse } from '@/types/api'

const PAGE_SIZE = 20

export function useLeaderboards(initialKey: LeaderboardKey) {
  const currentKey = shallowRef<LeaderboardKey>(initialKey)
  const currentPage = shallowRef(1)
  const entries = ref<LeaderboardEntry[]>([])
  const totalCount = shallowRef(0)
  const loading = shallowRef(false)

  const totalPages = computed(() => Math.max(1, Math.ceil(totalCount.value / PAGE_SIZE)))

  async function fetchLeaderboard(): Promise<void> {
    loading.value = true

    try {
      const encodedKey = encodeURIComponent(currentKey.value)
      const url = `${API_BASE_URL}/api/leaderboards/${encodedKey}?page=${currentPage.value}`
      const json = await fetchWithCache<LeaderboardResponse>(url)
      entries.value = json.entries ?? []
      totalCount.value = json.count ?? 0
    } catch (error) {
      console.error('Failed to load leaderboard data', error)
      entries.value = []
      totalCount.value = 0
    } finally {
      loading.value = false
    }
  }

  function setKey(key: LeaderboardKey): void {
    if (currentKey.value === key) {
      return
    }

    currentKey.value = key
    currentPage.value = 1
  }

  function goToPreviousPage(): void {
    if (currentPage.value <= 1) {
      return
    }

    currentPage.value -= 1
  }

  function goToNextPage(): void {
    if (currentPage.value >= totalPages.value) {
      return
    }

    currentPage.value += 1
  }

  watch(
    () => [currentKey.value, currentPage.value],
    () => {
      void fetchLeaderboard()
    },
    { immediate: true }
  )

  return {
    currentKey,
    currentPage,
    entries,
    loading,
    pageSize: PAGE_SIZE,
    totalCount,
    totalPages,
    goToNextPage,
    goToPreviousPage,
    setKey
  }
}
