<script setup lang="ts">
import { computed, nextTick, ref, shallowRef, useTemplateRef, watch } from 'vue'
import { useRoute } from 'vue-router'
import { storeToRefs } from 'pinia'

import PlayerAdvancementList from '@/components/player/PlayerAdvancementList.vue'
import PlayerDetailHeader from '@/components/player/PlayerDetailHeader.vue'
import PlayerInlineLeaderboard from '@/components/player/PlayerInlineLeaderboard.vue'
import PlayerDetailOverview from '@/components/player/PlayerDetailOverview.vue'
import PlayerDetailTabs from '@/components/player/PlayerDetailTabs.vue'
import PlayerStatsExtended from '@/components/player/PlayerStatsExtended.vue'
import { useAdvancements } from '@/composables/useAdvancements'
import { usePlayerDetail } from '@/composables/usePlayerDetail'
import { usePlayerStats } from '@/composables/usePlayerStats'
import { preloadImages, PreloadPriority } from '@/utils/preloader'
import { playerStatGroups, siteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'
import type { LeaderboardTarget } from '@/types/api'
import type { FixedLeaderboardType } from '@/types/api'
import type { PlayerDetailTab, PlayerDetailTabOption, PlayerRankHighlight } from '@/types/playerDetail'
import { createFixedLeaderboardTarget } from '@/utils/leaderboards'
import { getSkinUrl } from '@/utils/minecraft'

const route = useRoute()
const statusStore = useServerStatusStore()
const { status } = storeToRefs(statusStore)

const uuid = computed(() => route.params.uuid as string)
const { loading, info, stats, advancements, mcmmo, linkedAccount, oreStats, ranks, onlineHeatmap } = usePlayerDetail(uuid)
const { formatStatName, formatStatValue, getStatIconPath, getFilteredStats, getFilteredMcmmo } = usePlayerStats(stats)
const { totalAdvancements, completedAdvancements, categorizedAdvancements, getAdvancementMetadata, getAdvIconPath } =
  useAdvancements(advancements)

const selectedCategory = ref('minecraft:custom')
const statSearch = ref('')
const activeGroup = ref(playerStatGroups[0])
const activeTab = ref<PlayerDetailTab>('overview')
const selectedLeaderboard = shallowRef<LeaderboardTarget | null>(null)
const leaderboardPanelAnchor = useTemplateRef<HTMLDivElement>('leaderboardPanelAnchor')

const tabOptions: PlayerDetailTabOption[] = [
  { value: 'overview', label: siteContent.playerDetail.tabs.overview },
  { value: 'advancements', label: siteContent.playerDetail.tabs.advancements },
  { value: 'statistics', label: siteContent.playerDetail.tabs.statistics }
]

const onlinePlayers = computed(() => status.value?.online_players ?? [])
const isOnline = computed(() => onlinePlayers.value.includes(uuid.value))
const isLeaderboardPlayerOnline = (playerUuid: string): boolean => onlinePlayers.value.includes(playerUuid)

const filteredStats = computed(() => getFilteredStats(selectedCategory.value, statSearch.value))
const filteredMcmmo = computed(() => getFilteredMcmmo(mcmmo.value))

const groupCategories = computed(() => {
  if (!stats.value) {
    return []
  }

  return activeGroup.value.categories.filter((category) => stats.value?.[category] && category !== 'minecraft:custom')
})

const rankHighlights = computed<PlayerRankHighlight[]>(() =>
  [
    { label: 'McMMO', rank: ranks.value.skills, key: 'skills' as FixedLeaderboardType },
    { label: 'Mining', rank: ranks.value.mining, key: 'mining' as FixedLeaderboardType },
    { label: 'Playtime', rank: ranks.value.playtime, key: 'playtime' as FixedLeaderboardType },
    { label: 'Kills', rank: ranks.value.killing, key: 'killing' as FixedLeaderboardType },
    { label: 'Walking', rank: ranks.value.walking, key: 'walking' as FixedLeaderboardType },
    { label: 'PvP', rank: ranks.value.pvp, key: 'pvp' as FixedLeaderboardType }
  ]
    .filter((entry): entry is { label: string; rank: number; key: FixedLeaderboardType } => typeof entry.rank === 'number' && entry.rank > 0)
    .sort((left, right) => left.rank - right.rank)
    .slice(0, 4)
    .map((entry) => ({
      label: entry.label,
      value: `#${entry.rank}`,
      target: createFixedLeaderboardTarget(entry.key)
    }))
)

const topRankHighlight = computed(() => rankHighlights.value[0] ?? null)

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
    selectedLeaderboard.value = null
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

async function handleSelectLeaderboard(target: LeaderboardTarget): Promise<void> {
  selectedLeaderboard.value = target
  await nextTick()
  leaderboardPanelAnchor.value?.scrollIntoView({ behavior: 'smooth', block: 'start' })
}

function closeLeaderboard(): void {
  selectedLeaderboard.value = null
}
</script>

<template>
  <div class="player-detail container page-shell">
    <div v-if="loading" class="detail-layout" aria-hidden="true">
      <section class="glass-card detail-header-skeleton">
        <span class="skeleton-line skeleton-title-line"></span>
        <span class="skeleton-line skeleton-name-line"></span>
        <span class="skeleton-line skeleton-copy-line"></span>
        <div class="detail-skeleton-badges">
          <span class="skeleton-chip"></span>
          <span class="skeleton-chip"></span>
          <span class="skeleton-chip"></span>
        </div>
      </section>

      <section class="glass-card detail-tabs-skeleton">
        <span v-for="index in 3" :key="index" class="skeleton-block detail-tab-skeleton"></span>
      </section>

      <section class="glass-card detail-panel-skeleton">
        <div class="detail-skeleton-grid">
          <span v-for="index in 6" :key="index" class="skeleton-block detail-card-skeleton"></span>
        </div>
      </section>
    </div>

    <div v-else-if="!info" class="glass-card state-card">
      <h1>{{ siteContent.playerDetail.emptyTitle }}</h1>
      <p>{{ siteContent.playerDetail.emptyBody }}</p>
      <RouterLink class="btn-primary" to="/players">{{ siteContent.playerDetail.back }}</RouterLink>
    </div>

    <div v-else class="detail-layout">
      <PlayerDetailHeader
        :info="info"
        :is-online="isOnline"
        :linked-account="linkedAccount"
      />

      <PlayerDetailTabs v-model:active-tab="activeTab" :tabs="tabOptions" />

      <div v-if="selectedLeaderboard" ref="leaderboardPanelAnchor">
        <PlayerInlineLeaderboard :target="selectedLeaderboard" :is-online="isLeaderboardPlayerOnline" @close="closeLeaderboard" />
      </div>

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
          :online-heatmap="onlineHeatmap"
          :rank-highlights="rankHighlights"
          :top-rank-highlight="topRankHighlight"
          :format-date="formatDate"
          :format-playtime="formatPlaytime"
          @select-leaderboard="handleSelectLeaderboard"
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
            @select-leaderboard="handleSelectLeaderboard"
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

.detail-header-skeleton,
.detail-panel-skeleton {
  display: grid;
  gap: 0.85rem;
}

.detail-tabs-skeleton {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 0.55rem;
}

.detail-tab-skeleton {
  min-height: 3rem;
  border-radius: 18px;
}

.detail-skeleton-badges {
  display: flex;
  flex-wrap: wrap;
  gap: 0.55rem;
}

.skeleton-title-line {
  width: 10rem;
}

.skeleton-name-line {
  width: min(24rem, 70%);
  height: 2.6rem;
}

.skeleton-copy-line {
  width: min(28rem, 85%);
}

.detail-skeleton-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 0.75rem;
}

.detail-card-skeleton {
  min-height: 8rem;
  border-radius: 18px;
}

@media (max-width: 860px) {
  .detail-skeleton-grid {
    grid-template-columns: 1fr;
  }
}
</style>
