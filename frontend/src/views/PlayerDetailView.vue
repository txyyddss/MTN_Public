<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { storeToRefs } from 'pinia'

import PlayerAdvancementList from '@/components/player/PlayerAdvancementList.vue'
import PlayerDetailHero from '@/components/player/PlayerDetailHero.vue'
import PlayerDetailOverview from '@/components/player/PlayerDetailOverview.vue'
import PlayerDetailTabs from '@/components/player/PlayerDetailTabs.vue'
import PlayerStatsExtended from '@/components/player/PlayerStatsExtended.vue'
import { useAdvancements } from '@/composables/useAdvancements'
import { usePlayerDetail } from '@/composables/usePlayerDetail'
import { usePlayerStats } from '@/composables/usePlayerStats'
import { preloadImages, PreloadPriority } from '@/utils/preloader'
import { playerStatGroups, siteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'
import type { PlayerDetailTab, PlayerDetailTabOption } from '@/types/playerDetail'
import { getSkinUrl } from '@/utils/minecraft'

const route = useRoute()
const statusStore = useServerStatusStore()
const { status } = storeToRefs(statusStore)

const uuid = computed(() => route.params.uuid as string)
const { loading, info, stats, advancements, mcmmo, linkedAccount, oreStats, ranks } = usePlayerDetail(uuid)
const { formatStatName, formatStatValue, getStatIconPath, getFilteredStats, getFilteredMcmmo } = usePlayerStats(stats)
const { totalAdvancements, completedAdvancements, categorizedAdvancements, getAdvancementMetadata, getAdvIconPath } =
  useAdvancements(advancements)

const selectedCategory = ref('minecraft:custom')
const statSearch = ref('')
const activeGroup = ref(playerStatGroups[0])
const activeTab = ref<PlayerDetailTab>('overview')

const tabOptions: PlayerDetailTabOption[] = [
  { value: 'overview', label: siteContent.playerDetail.tabs.overview },
  { value: 'advancements', label: siteContent.playerDetail.tabs.advancements },
  { value: 'statistics', label: siteContent.playerDetail.tabs.statistics }
]

const onlinePlayers = computed(() => status.value?.online_players ?? [])
const isOnline = computed(() => onlinePlayers.value.includes(uuid.value))

const filteredStats = computed(() => getFilteredStats(selectedCategory.value, statSearch.value))
const filteredMcmmo = computed(() => getFilteredMcmmo(mcmmo.value))
const skillTotal = computed(() => mcmmo.value?.total ?? 0)

const groupCategories = computed(() => {
  if (!stats.value) {
    return []
  }

  return activeGroup.value.categories.filter((category) => stats.value?.[category] && category !== 'minecraft:custom')
})

const rankHighlights = computed(() =>
  [
    { label: 'McMMO', rank: ranks.value.skills },
    { label: 'Mining', rank: ranks.value.mining },
    { label: 'Playtime', rank: ranks.value.playtime },
    { label: 'Kills', rank: ranks.value.killing },
    { label: 'Walking', rank: ranks.value.walking },
    { label: 'PvP', rank: ranks.value.pvp }
  ]
    .filter((entry): entry is { label: string; rank: number } => typeof entry.rank === 'number' && entry.rank > 0)
    .sort((left, right) => left.rank - right.rank)
    .slice(0, 4)
    .map((entry) => ({ label: entry.label, value: `#${entry.rank}` }))
)

watch(
  [activeGroup, stats],
  ([group]) => {
    const firstAvailable = group.categories.find((category) => category === 'minecraft:custom' || stats.value?.[category])
    if (firstAvailable) {
      selectedCategory.value = firstAvailable
    }
  },
  { immediate: true }
)

watch(
  uuid,
  () => {
    activeTab.value = 'overview'
    activeGroup.value = playerStatGroups[0]
    selectedCategory.value = 'minecraft:custom'
    statSearch.value = ''
  },
  { immediate: true }
)

watch(info, (playerInfo) => {
  if (playerInfo) {
    preloadImages([getSkinUrl(playerInfo.last_known_name, playerInfo.type)], PreloadPriority.HIGH)
  }
})

watch(
  categorizedAdvancements,
  (categories) => {
    const urls = Object.values(categories).flatMap((items) => items.map((advancement) => getAdvIconPath(advancement.key)))
    if (urls.length > 0) {
      preloadImages(urls, PreloadPriority.MEDIUM_HIGH)
    }
  },
  { immediate: true }
)

function formatDate(value: number): string {
  return value ? new Date(value).toLocaleDateString() : 'N/A'
}

function formatPlaytime(ticks: number): string {
  if (!ticks) {
    return '0h'
  }

  return `${(ticks / 20 / 3600).toFixed(1)}h`
}
</script>

<template>
  <div class="player-detail container page-shell">
    <div v-if="loading" class="glass-card state-card">
      {{ siteContent.playerDetail.loading }}
    </div>

    <div v-else-if="!info" class="glass-card state-card">
      <h1>{{ siteContent.playerDetail.emptyTitle }}</h1>
      <p>{{ siteContent.playerDetail.emptyBody }}</p>
      <RouterLink class="btn-primary" to="/players">{{ siteContent.playerDetail.back }}</RouterLink>
    </div>

    <div v-else class="detail-layout">
      <PlayerDetailHero
        :info="info"
        :is-online="isOnline"
        :linked-account="linkedAccount"
        :completed-advancements="completedAdvancements"
        :total-advancements="totalAdvancements"
        :skill-total="skillTotal"
        :rank-highlights="rankHighlights"
        :format-playtime="formatPlaytime"
      />

      <PlayerDetailTabs v-model:active-tab="activeTab" :tabs="tabOptions" />

      <Transition name="panel-swap" mode="out-in">
        <PlayerDetailOverview
          v-if="activeTab === 'overview'"
          key="overview"
          :info="info"
          :is-online="isOnline"
          :linked-account="linkedAccount"
          :ore-stats="oreStats"
          :mcmmo="mcmmo"
          :ranks="ranks"
          :filtered-mcmmo="filteredMcmmo"
          :completed-advancements="completedAdvancements"
          :total-advancements="totalAdvancements"
          :rank-highlights="rankHighlights"
          :format-date="formatDate"
          :format-playtime="formatPlaytime"
        />

        <section v-else-if="activeTab === 'advancements'" key="advancements" class="tab-panel">
          <PlayerAdvancementList
            :advancements="advancements"
            :completed-advancements="completedAdvancements"
            :total-advancements="totalAdvancements"
            :categorized-advancements="categorizedAdvancements"
            :get-advancement-metadata="getAdvancementMetadata"
            :get-adv-icon-path="getAdvIconPath"
          />
        </section>

        <section v-else key="statistics" class="tab-panel">
          <PlayerStatsExtended
            v-model:active-group="activeGroup"
            v-model:selected-category="selectedCategory"
            v-model:stat-search="statSearch"
            :stats="stats"
            :stat-groups="playerStatGroups"
            :group-categories="groupCategories"
            :filtered-stats="filteredStats"
            :ranks="ranks"
            :format-stat-name="formatStatName"
            :format-stat-value="formatStatValue"
            :get-stat-icon-path="getStatIconPath"
          />
        </section>
      </Transition>
    </div>
  </div>
</template>

<style scoped>
.player-detail,
.detail-layout,
.tab-panel {
  display: grid;
  gap: 0.9rem;
}

.tab-panel {
  align-items: start;
}

.state-card {
  display: grid;
  gap: 0.85rem;
}
</style>
