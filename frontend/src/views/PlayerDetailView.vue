<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { storeToRefs } from 'pinia'

import PlayerPieChart from '@/components/PlayerPieChart.vue'
import PlayerAdvancementList from '@/components/player/PlayerAdvancementList.vue'
import PlayerSidebar from '@/components/player/PlayerSidebar.vue'
import PlayerSkills from '@/components/player/PlayerSkills.vue'
import PlayerStatsExtended from '@/components/player/PlayerStatsExtended.vue'
import { preloadImages, PreloadPriority } from '@/utils/preloader'
import { siteContent, playerStatGroups } from '@/content/siteContent'
import { useAdvancements } from '@/composables/useAdvancements'
import { usePlayerDetail } from '@/composables/usePlayerDetail'
import { usePlayerStats } from '@/composables/usePlayerStats'
import { useServerStatusStore } from '@/stores/serverStatus'
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

const onlinePlayers = computed(() => status.value?.online_players ?? [])
const isOnline = computed(() => onlinePlayers.value.includes(uuid.value))

const filteredStats = computed(() => getFilteredStats(selectedCategory.value, statSearch.value))
const filteredMcmmo = computed(() => getFilteredMcmmo(mcmmo.value))

const groupCategories = computed(() => {
  if (!stats.value) {
    return []
  }

  return activeGroup.value.categories.filter((category) => stats.value?.[category] && category !== 'minecraft:custom')
})

watch(activeGroup, (group) => {
  const firstAvailable = group.categories.find((category) => category === 'minecraft:custom' || stats.value?.[category])
  if (firstAvailable) {
    selectedCategory.value = firstAvailable
  }
})

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

    <div v-else class="content-grid">
      <PlayerSidebar
        :info="info"
        :is-online="isOnline"
        :linked-account="linkedAccount"
        :format-date="formatDate"
        :format-playtime="formatPlaytime"
      />

      <main class="details-grid">
        <PlayerPieChart :ore-stats="oreStats" />
        <PlayerSkills :mcmmo="mcmmo" :ranks="ranks" :filtered-mcmmo="filteredMcmmo" />
        <PlayerAdvancementList
          :advancements="advancements"
          :completed-advancements="completedAdvancements"
          :total-advancements="totalAdvancements"
          :categorized-advancements="categorizedAdvancements"
          :get-advancement-metadata="getAdvancementMetadata"
          :get-adv-icon-path="getAdvIconPath"
        />
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
      </main>
    </div>
  </div>
</template>

<style scoped>
.content-grid {
  display: grid;
  grid-template-columns: 320px 1fr;
  gap: 1rem;
}

.details-grid {
  display: grid;
  gap: 1rem;
}

.state-card {
  display: grid;
  gap: 1rem;
}

@media (max-width: 1024px) {
  .content-grid {
    grid-template-columns: 1fr;
  }
}
</style>
