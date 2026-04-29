import { defineStore } from 'pinia'
import { ref } from 'vue'

import { API_BASE_URL } from '@/config'
import type { ConnectionResponse, StatusResponse } from '@/types/api'

export const useServerStatusStore = defineStore('serverStatus', () => {
  const status = ref<StatusResponse | null>(null)
  const connection = ref<ConnectionResponse | null>(null)
  const isLoading = ref(false)

  let pollTimer: ReturnType<typeof setInterval> | null = null
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

  async function refresh(): Promise<void> {
    if (refreshInFlight) {
      return refreshInFlight
    }

    isLoading.value = true
    refreshInFlight = Promise.allSettled([fetchStatus(), fetchConnection()])
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

  function startPolling(intervalMs = 5000): void {
    if (pollTimer) {
      stopPolling()
    }

    pollTimer = setInterval(() => {
      void refreshStatusOnly()
    }, intervalMs)
  }

  function stopPolling(): void {
    if (pollTimer) {
      clearInterval(pollTimer)
      pollTimer = null
    }

  }

  return {
    status,
    connection,
    isLoading,
    refresh,
    startPolling,
    stopPolling
  }
})
