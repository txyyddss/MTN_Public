import { defineStore } from 'pinia'
import { ref } from 'vue'

import { API_BASE_URL } from '@/config'
import type { ConnectionResponse, ServerHistoryResponse, StatusResponse } from '@/types/api'
import { fetchWithCache } from '@/utils/dataCache'

export const useServerStatusStore = defineStore('serverStatus', () => {
  const status = ref<StatusResponse | null>(null)
  const connection = ref<ConnectionResponse | null>(null)
  const history = ref<ServerHistoryResponse | null>(null)
  const isLoading = ref(false)

  let pollTimer: ReturnType<typeof setInterval> | null = null
  let historyTimer: ReturnType<typeof setInterval> | null = null
  let refreshInFlight: Promise<void> | null = null

  async function fetchStatus(): Promise<void> {
    const response = await fetch(`${API_BASE_URL}/api/status`)
    if (!response.ok) {
      throw new Error(`Status request failed: ${response.status}`)
    }

    status.value = (await response.json()) as StatusResponse
  }

  async function fetchConnection(): Promise<void> {
    const response = await fetch(`${API_BASE_URL}/api/connection`)
    if (!response.ok) {
      throw new Error(`Connection request failed: ${response.status}`)
    }

    connection.value = (await response.json()) as ConnectionResponse
  }

  async function fetchHistory(): Promise<void> {
    history.value = await fetchWithCache<ServerHistoryResponse>(`${API_BASE_URL}/api/status/history`, 60000)
  }

  async function refresh(): Promise<void> {
    if (refreshInFlight) {
      return refreshInFlight
    }

    isLoading.value = true
    refreshInFlight = Promise.allSettled([fetchStatus(), fetchConnection(), fetchHistory()])
      .then(() => undefined)
      .finally(() => {
        isLoading.value = false
        refreshInFlight = null
      })

    return refreshInFlight
  }

  async function refreshStatusOnly(): Promise<void> {
    try {
      await fetchStatus()
    } catch (error) {
      console.error('Failed to refresh live status', error)
    }
  }

  async function refreshHistoryOnly(): Promise<void> {
    try {
      await fetchHistory()
    } catch (error) {
      console.error('Failed to refresh server history', error)
    }
  }

  function startPolling(intervalMs = 5000, historyIntervalMs = 60000): void {
    if (pollTimer || historyTimer) {
      stopPolling()
    }

    pollTimer = setInterval(() => {
      void refreshStatusOnly()
    }, intervalMs)

    historyTimer = setInterval(() => {
      void refreshHistoryOnly()
    }, historyIntervalMs)
  }

  function stopPolling(): void {
    if (pollTimer) {
      clearInterval(pollTimer)
      pollTimer = null
    }

    if (historyTimer) {
      clearInterval(historyTimer)
      historyTimer = null
    }
  }

  return {
    status,
    connection,
    history,
    isLoading,
    refresh,
    startPolling,
    stopPolling
  }
})
