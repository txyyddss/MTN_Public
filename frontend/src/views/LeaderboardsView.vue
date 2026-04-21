<script setup lang="ts">
import { computed } from 'vue'
import { storeToRefs } from 'pinia'

import LeaderboardPagination from '@/components/leaderboards/LeaderboardPagination.vue'
import LeaderboardTable from '@/components/leaderboards/LeaderboardTable.vue'
import { useRevealOnScroll } from '@/composables/useRevealOnScroll'
import { useLeaderboards } from '@/composables/useLeaderboards'
import { leaderboardLabels, siteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'
import type { LeaderboardType } from '@/types/api'

const statusStore = useServerStatusStore()
const { status } = storeToRefs(statusStore)

const leaderboardTypes = Object.keys(leaderboardLabels) as LeaderboardType[]
const { currentPage, currentType, entries, goToNextPage, goToPreviousPage, loading, totalPages, setType } =
  useLeaderboards('mining')

const leaderboardDescriptor = computed(() => `${entries.value.length} visible entries`)
const { revealed: tabsRevealed } = useRevealOnScroll<HTMLDivElement>('leaderboardTabs')
const { revealed: boardRevealed } = useRevealOnScroll<HTMLDivElement>('leaderboardBoard')

function isOnline(uuid: string): boolean {
  return status.value?.online_players?.includes(uuid) ?? false
}

function formatValue(value: number, type: LeaderboardType): string {
  if (type === 'playtime') {
    return `${(value / 20 / 3600).toFixed(1)} hrs`
  }

  if (type === 'walking') {
    const kilometers = value / 100000
    return kilometers >= 1 ? `${kilometers.toFixed(1)} km` : `${(value / 100).toFixed(0)} m`
  }

  return value.toLocaleString()
}
</script>

<template>
  <div class="leaderboards-view container page-shell">
    <header class="page-header animate-entry">
      <span class="page-kicker">Live rankings</span>
      <h1>{{ siteContent.leaderboards.title }}</h1>
      <p class="page-lede">{{ siteContent.leaderboards.body }}</p>
    </header>

    <div
      ref="leaderboardTabs"
      :class="['leaderboard-tabs', 'glass-card', 'scroll-reveal', { 'is-revealed': tabsRevealed }]"
    >
      <div class="leaderboard-tabs-copy">
        <span class="hud-kicker">Board select</span>
        <p class="tabs-note">{{ leaderboardDescriptor }}</p>
      </div>

      <div class="tab-row">
        <button
          v-for="type in leaderboardTypes"
          :key="type"
          :class="['tab-chip', { active: currentType === type }]"
          type="button"
          @click="setType(type)"
        >
          {{ leaderboardLabels[type] }}
        </button>
      </div>
    </div>

    <div v-if="loading" class="glass-card state-card">{{ siteContent.leaderboards.loading }}</div>

    <div v-else-if="entries.length === 0" class="glass-card state-card">
      <h2>{{ siteContent.leaderboards.emptyTitle }}</h2>
      <p>{{ siteContent.leaderboards.emptyBody }}</p>
    </div>

    <div
      v-else
      ref="leaderboardBoard"
      :class="['glass-card', 'leaderboard-board', 'scroll-reveal', { 'is-revealed': boardRevealed }]"
    >
      <LeaderboardTable
        :current-type="currentType"
        :entries="entries"
        :format-value="formatValue"
        :is-online="isOnline"
        :player-label="siteContent.leaderboards.columns.player"
        :rank-label="siteContent.leaderboards.columns.rank"
        :score-label="siteContent.leaderboards.columns.score"
      />
      <LeaderboardPagination
        :current-page="currentPage"
        :total-pages="totalPages"
        @next="goToNextPage"
        @previous="goToPreviousPage"
      />
    </div>
  </div>
</template>

<style scoped>
.leaderboards-view {
  display: grid;
  gap: 0.9rem;
}

.leaderboard-tabs {
  display: grid;
  gap: 0.8rem;
  margin-bottom: 0;
}

.leaderboard-tabs-copy {
  display: grid;
  gap: 0.45rem;
}

.tabs-note {
  color: var(--text-muted);
}

.tab-row {
  display: flex;
  flex-wrap: wrap;
  gap: 0.45rem;
}

.tab-chip {
  position: relative;
  overflow: hidden;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.03);
  color: var(--text-muted);
  padding: 0.6rem 0.9rem;
  font-family: var(--sans);
  font-size: 0.88rem;
  font-weight: 500;
  letter-spacing: -0.01em;
  transition:
    transform var(--transition-panel),
    border-color var(--transition-fast),
    color var(--transition-fast),
    background var(--transition-fast);
}

.tab-chip::after {
  content: '';
  position: absolute;
  inset: auto 14px 7px;
  height: 1px;
  border-radius: 999px;
  background: var(--primary);
  transform: scaleX(0);
  transform-origin: center;
  transition: transform var(--transition-panel);
}

.tab-chip:hover {
  color: var(--text-main);
  transform: translateY(-1px);
}

.tab-chip.active {
  border-color: rgba(76, 147, 251, 0.3);
  background: rgba(76, 147, 251, 0.08);
  color: var(--text-strong);
}

.tab-chip.active::after {
  transform: scaleX(1);
}

.leaderboard-board {
  display: grid;
  gap: 0.35rem;
}

.state-card {
  display: grid;
  gap: 0.85rem;
}
</style>
