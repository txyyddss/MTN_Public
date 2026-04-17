import { ref, onMounted, onUnmounted } from 'vue'
import { API_BASE_URL } from '@/config'

export function useServerStatus() {
    const status = ref<any>(null)
    const connInfo = ref<any>(null)
    let pollInterval: any = null

    const fetchStatus = async () => {
        try {
            const res = await fetch(`${API_BASE_URL}/api/status`)
            if (res.ok) {
                status.value = await res.json()
            }
        } catch (err) {
            console.error('Failed to fetch status', err)
        }
    }

    const fetchConnection = async () => {
        try {
            const res = await fetch(`${API_BASE_URL}/api/connection`)
            if (res.ok) {
                connInfo.value = await res.json()
            }
        } catch (err) {
            console.error('Failed to fetch connection info', err)
        }
    }

    onMounted(() => {
        fetchStatus()
        fetchConnection()
        pollInterval = setInterval(fetchStatus, 5000)
    })

    onUnmounted(() => {
        if (pollInterval) clearInterval(pollInterval)
    })

    return {
        status,
        connInfo,
        fetchStatus,
        fetchConnection
    }
}
