import { storeToRefs } from 'pinia'

import { useServerStatusStore } from '@/stores/serverStatus'

export function useServerStatus() {
  const store = useServerStatusStore()
  const { status, connection, isLoading } = storeToRefs(store)

  return {
    status,
    connection,
    isLoading,
    refresh: store.refresh,
    startPolling: store.startPolling,
    stopPolling: store.stopPolling
  }
}
