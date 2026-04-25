import { ref, toValue, watch, type MaybeRefOrGetter } from 'vue'

import { API_BASE_URL } from '@/config'
import { fetchWithCache } from '@/utils/dataCache'
import type {
  LinkedAccount,
  McMMOSkills,
  OreStat,
  PlayerAdvancement,
  PlayerDetailResponse,
  PlayerInfo,
  PlayerOnlineHeatmap,
  PlayerStatBuckets,
  WhitelistAccount
} from '@/types/api'

export function usePlayerDetail(uuid: MaybeRefOrGetter<string>) {
  const loading = ref(true)
  const info = ref<PlayerInfo | null>(null)
  const stats = ref<PlayerStatBuckets | null>(null)
  const advancements = ref<PlayerAdvancement[] | null>(null)
  const mcmmo = ref<McMMOSkills | null>(null)
  const linkedAccount = ref<LinkedAccount | null>(null)
  const whitelistAccount = ref<WhitelistAccount | null>(null)
  const oreStats = ref<OreStat[]>([])
  const ranks = ref<Record<string, number>>({})
  const onlineHeatmap = ref<PlayerOnlineHeatmap | null>(null)

  async function fetchDetail(playerUuid = toValue(uuid)): Promise<void> {
    if (!playerUuid) {
      info.value = null
      stats.value = null
      advancements.value = null
      mcmmo.value = null
      linkedAccount.value = null
      whitelistAccount.value = null
      oreStats.value = []
      ranks.value = {}
      onlineHeatmap.value = null
      loading.value = false
      return
    }

    loading.value = true

    try {
      const json = await fetchWithCache<PlayerDetailResponse>(`${API_BASE_URL}/api/players/${playerUuid}`)
      info.value = json.info
      stats.value = json.stats?.stats ?? null
      advancements.value = json.advancements?.advancements ?? []
      mcmmo.value = json.mcmmo
      linkedAccount.value = json.linked_account
      whitelistAccount.value = json.whitelist_account
      oreStats.value = json.ore_stats ?? []
      ranks.value = json.ranks ?? {}
      onlineHeatmap.value = json.online_heatmap ?? null
    } catch (error) {
      console.error('Failed to load player detail', error)
      info.value = null
      stats.value = null
      advancements.value = null
      mcmmo.value = null
      linkedAccount.value = null
      whitelistAccount.value = null
      oreStats.value = []
      ranks.value = {}
      onlineHeatmap.value = null
    } finally {
      loading.value = false
    }
  }

  watch(
    () => toValue(uuid),
    (currentUuid) => {
      void fetchDetail(currentUuid)
    },
    { immediate: true }
  )

  return {
    loading,
    info,
    stats,
    advancements,
    mcmmo,
    linkedAccount,
    whitelistAccount,
    oreStats,
    ranks,
    onlineHeatmap,
    fetchDetail
  }
}
