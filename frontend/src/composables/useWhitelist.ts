import { computed, onMounted, shallowRef, watch } from 'vue'

import { API_BASE_URL } from '@/config'
import type {
  WhitelistEditionFilter,
  WhitelistEntry,
  WhitelistListResponse,
  WhitelistMutationInput,
  WhitelistOperationResponse
} from '@/types/api'

const tokenStorageKey = 'mtn.whitelist.token'

export function useWhitelist() {
  const token = shallowRef(sessionStorage.getItem(tokenStorageKey) ?? '')
  const editionFilter = shallowRef<WhitelistEditionFilter>('all')
  const entries = shallowRef<WhitelistEntry[]>([])
  const loading = shallowRef(false)
  const mutating = shallowRef(false)
  const error = shallowRef('')
  const notice = shallowRef('')
  let loadSequence = 0
  let mutationSequence = 0

  const hasToken = computed(() => token.value.trim().length > 0)
  const activeCount = computed(() => entries.value.length)
  const javaCount = computed(() => entries.value.filter((entry) => entry.edition === 'java').length)
  const bedrockCount = computed(() => entries.value.filter((entry) => entry.edition === 'bedrock').length)

  watch(token, (value) => {
    const trimmed = value.trim()
    if (trimmed) {
      sessionStorage.setItem(tokenStorageKey, trimmed)
      return
    }

    sessionStorage.removeItem(tokenStorageKey)
  })

  async function loadEntries(): Promise<void> {
    const requestId = ++loadSequence

    if (!hasToken.value) {
      if (requestId === loadSequence) {
        entries.value = []
        error.value = 'Token required'
      }
      return
    }

    loading.value = true
    error.value = ''

    try {
      const params = new URLSearchParams({ edition: editionFilter.value })
      const response = await fetch(`${API_BASE_URL}/api/whitelist?${params.toString()}`, {
        headers: authHeaders()
      })

      if (!response.ok) {
        throw await responseError(response)
      }

      const json = (await response.json()) as WhitelistListResponse
      if (requestId === loadSequence) {
        entries.value = json.entries ?? []
      }
    } catch (caught) {
      if (requestId === loadSequence) {
        error.value = caught instanceof Error ? caught.message : 'Whitelist request failed'
        entries.value = []
      }
    } finally {
      if (requestId === loadSequence) {
        loading.value = false
      }
    }
  }

  async function addEntry(input: WhitelistMutationInput): Promise<void> {
    await mutate('/api/whitelist/add', input)
  }

  async function removeEntry(input: WhitelistMutationInput): Promise<void> {
    await mutate('/api/whitelist/remove', input)
  }

  async function mutate(path: string, input: WhitelistMutationInput): Promise<void> {
    const requestId = ++mutationSequence

    if (!hasToken.value) {
      if (requestId === mutationSequence) {
        error.value = 'Token required'
      }
      return
    }

    mutating.value = true
    error.value = ''
    notice.value = ''

    try {
      const response = await fetch(`${API_BASE_URL}${path}`, {
        method: 'POST',
        headers: {
          ...authHeaders(),
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          edition: input.edition,
          player_name: input.playerName,
          qq_id: input.qqId?.trim() || undefined
        })
      })

      if (!response.ok) {
        throw await responseError(response)
      }

      const json = (await response.json()) as WhitelistOperationResponse
      if (requestId === mutationSequence) {
        notice.value = formatOperationNotice(json)
        await loadEntries()
      }
    } catch (caught) {
      if (requestId === mutationSequence) {
        error.value = caught instanceof Error ? caught.message : 'Whitelist request failed'
      }
    } finally {
      if (requestId === mutationSequence) {
        mutating.value = false
      }
    }
  }

  function authHeaders(): HeadersInit {
    return {
      Authorization: `Bearer ${token.value.trim()}`
    }
  }

  onMounted(() => {
    if (hasToken.value) {
      void loadEntries()
    }
  })

  watch(editionFilter, () => {
    if (hasToken.value) {
      void loadEntries()
    }
  })

  return {
    token,
    editionFilter,
    entries,
    loading,
    mutating,
    error,
    notice,
    hasToken,
    activeCount,
    javaCount,
    bedrockCount,
    loadEntries,
    addEntry,
    removeEntry
  }
}

function formatOperationNotice(response: WhitelistOperationResponse): string {
  if (!response.quota) {
    return response.message
  }
  if (response.quota.exempt) {
    return `${response.message}. Quota exempt`
  }
  return `${response.message}. Quota ${response.quota.used}/${response.quota.limit} used, ${response.quota.remaining} remaining`
}

async function responseError(response: Response): Promise<Error> {
  try {
    const payload = (await response.json()) as { error?: string; code?: string }
    return new Error(payload.error || payload.code || `Request failed: ${response.status}`)
  } catch {
    return new Error(`Request failed: ${response.status}`)
  }
}
