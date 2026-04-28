<script setup lang="ts">
import { computed, nextTick, ref, shallowRef, useTemplateRef, watch } from 'vue'
import { useRoute } from 'vue-router'
import { storeToRefs } from 'pinia'

import DossierLoadingPanel from '@/components/common/DossierLoadingPanel.vue'
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
import {
  formatDurationHoursShort,
  getLeaderboardLabel,
  getLocaleValue,
  useCurrentLocale,
  usePlayerStatGroups,
  useSiteContent
} from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'
import type { LeaderboardTarget } from '@/types/api'
import type { FixedLeaderboardType } from '@/types/api'
import type { PlayerDetailTab, PlayerDetailTabOption, PlayerRankHighlight } from '@/types/playerDetail'
import { createFixedLeaderboardTarget } from '@/utils/leaderboards'
import { getSkinUrl } from '@/utils/minecraft'

const route = useRoute()
const statusStore = useServerStatusStore()
const { status } = storeToRefs(statusStore)
const siteContent = useSiteContent()
const currentLocale = useCurrentLocale()
const playerStatGroups = usePlayerStatGroups()

const uuid = computed(() => route.params.uuid as string)
const { loading, info, stats, advancements, mcmmo, linkedAccount, whitelistAccount, oreStats, ranks, onlineHeatmap } =
  usePlayerDetail(uuid)
const { formatStatName, formatStatValue, getStatIconPath, getFilteredStats, getFilteredMcmmo } = usePlayerStats(stats)
const { totalAdvancements, completedAdvancements, categorizedAdvancements, getAdvancementMetadata, getAdvIconPath } =
  useAdvancements(advancements)

const selectedCategory = ref('minecraft:custom')
const statSearch = ref('')
const activeGroup = ref(playerStatGroups.value[0])
const activeTab = ref<PlayerDetailTab>('overview')
const selectedLeaderboard = shallowRef<LeaderboardTarget | null>(null)
const leaderboardPanelAnchor = useTemplateRef<HTMLDivElement>('leaderboardPanelAnchor')

const tabOptions = computed<PlayerDetailTabOption[]>(() => [
  { value: 'overview', label: siteContent.value.playerDetail.tabs.overview },
  { value: 'advancements', label: siteContent.value.playerDetail.tabs.advancements },
  { value: 'statistics', label: siteContent.value.playerDetail.tabs.statistics }
])

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

const rankHighlights = computed<PlayerRankHighlight[]>(() => {
  void currentLocale.value

  return [
    { rank: ranks.value.skills, key: 'skills' as FixedLeaderboardType },
    { rank: ranks.value.mining, key: 'mining' as FixedLeaderboardType },
    { rank: ranks.value.playtime, key: 'playtime' as FixedLeaderboardType },
    { rank: ranks.value.killing, key: 'killing' as FixedLeaderboardType },
    { rank: ranks.value.walking, key: 'walking' as FixedLeaderboardType },
    { rank: ranks.value.pvp, key: 'pvp' as FixedLeaderboardType }
  ]
    .filter((entry): entry is { rank: number; key: FixedLeaderboardType } => typeof entry.rank === 'number' && entry.rank > 0)
    .sort((left, right) => left.rank - right.rank)
    .slice(0, 4)
    .map((entry) => ({
      label: getLeaderboardLabel(entry.key),
      value: `#${entry.rank}`,
      target: createFixedLeaderboardTarget(entry.key)
    }))
})

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
  playerStatGroups,
  (groups) => {
    const activeSignature = activeGroup.value.categories.join('|')
    activeGroup.value = groups.find((group) => group.categories.join('|') === activeSignature) ?? groups[0]
  },
  { immediate: true }
)

watch(
  uuid,
  () => {
    activeTab.value = 'overview'
    activeGroup.value = playerStatGroups.value[0]
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
  if (!value) {
    return siteContent.value.players.lastSeenUnavailable
  }

  return new Intl.DateTimeFormat(getLocaleValue(), {
    year: 'numeric',
    month: 'numeric',
    day: 'numeric'
  }).format(new Date(value))
}

function formatDateTime(value: number): string {
  if (!value) {
    return siteContent.value.players.lastSeenUnavailable
  }

  return new Intl.DateTimeFormat(getLocaleValue(), {
    year: 'numeric',
    month: 'numeric',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(new Date(value))
}

function formatPlaytime(ticks: number): string {
  const hours = ticks / 20 / 3600
  void currentLocale.value
  return formatDurationHoursShort(hours.toFixed(1))
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
  <div class="player-detail container page-shell route-page-shell">
    <div v-if="loading" class="detail-loading-stack" aria-hidden="true">
      <DossierLoadingPanel class="detail-loading-hero" label="MTN PLAYER DOSSIER">
        <span class="skeleton-line skeleton-title-line"></span>
        <span class="skeleton-line skeleton-name-line"></span>
        <span class="skeleton-line skeleton-copy-line"></span>
        <div class="detail-skeleton-badges">
          <span class="skeleton-chip"></span>
          <span class="skeleton-chip"></span>
          <span class="skeleton-chip"></span>
        </div>
      </DossierLoadingPanel>

      <DossierLoadingPanel class="detail-tabs-skeleton" label="PROFILE TABS" compact>
        <span v-for="index in 3" :key="index" class="skeleton-block detail-tab-skeleton"></span>
      </DossierLoadingPanel>

      <DossierLoadingPanel class="detail-panel-skeleton" label="ARCHIVE STREAM">
        <div class="detail-skeleton-grid">
          <span v-for="index in 6" :key="index" class="skeleton-block detail-card-skeleton"></span>
        </div>
      </DossierLoadingPanel>
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
          :whitelist-account="whitelistAccount"
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
          :format-date-time="formatDateTime"
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
.detail-loading-stack,
.tab-panel {
  display: grid;
  gap: 1rem;
}

.tab-panel {
  align-items: start;
}

.state-card {
  display: grid;
  gap: 0.85rem;
  padding: 1.4rem;
}

.detail-loading-hero,
.detail-panel-skeleton {
  display: grid;
  gap: 0.85rem;
}

.detail-tabs-skeleton {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 0.55rem;
  min-height: auto;
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

.detail-loading-hero {
  min-height: 15rem;
  align-content: center;
}

.detail-loading-stack .skeleton-line,
.detail-loading-stack .skeleton-block,
.detail-loading-stack .skeleton-chip {
  background: rgba(141, 184, 255, 0.12);
}

@media (max-width: 860px) {
  .detail-skeleton-grid {
    grid-template-columns: 1fr;
  }
}
</style>
