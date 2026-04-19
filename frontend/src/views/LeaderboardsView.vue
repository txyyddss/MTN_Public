<script setup lang="ts">
import { storeToRefs } from 'pinia'

import LeaderboardPagination from '@/components/leaderboards/LeaderboardPagination.vue'
import LeaderboardTable from '@/components/leaderboards/LeaderboardTable.vue'
import { useLeaderboards } from '@/composables/useLeaderboards'
import { leaderboardLabels, siteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'
import type { LeaderboardType } from '@/types/api'

const statusStore = useServerStatusStore()
const { status } = storeToRefs(statusStore)

const leaderboardTypes = Object.keys(leaderboardLabels) as LeaderboardType[]
const { currentPage, currentType, entries, goToNextPage, goToPreviousPage, loading, totalPages, setType } =
  useLeaderboards('mining')

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

    <div class="leaderboard-tabs glass-card animate-entry delay-100">
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

    <div v-if="loading" class="glass-card state-card">{{ siteContent.leaderboards.loading }}</div>

    <div v-else-if="entries.length === 0" class="glass-card state-card">
      <h2>{{ siteContent.leaderboards.emptyTitle }}</h2>
      <p>{{ siteContent.leaderboards.emptyBody }}</p>
    </div>

    <div v-else class="glass-card animate-entry delay-200">
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
.leaderboard-tabs {
  display: flex;
  flex-wrap: wrap;
  gap: 0.6rem;
  margin-bottom: 1rem;
}

.tab-chip {
  border: 1px solid var(--glass-border);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.03);
  color: var(--text-muted);
  padding: 0.62rem 0.92rem;
}

.tab-chip.active {
  border-color: rgba(55, 111, 255, 0.42);
  color: var(--text-strong);
}

.state-card {
  display: grid;
  gap: 1rem;
}
</style>
