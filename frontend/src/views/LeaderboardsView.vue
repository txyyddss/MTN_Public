<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { storeToRefs } from 'pinia'

import { leaderboardLabels, siteContent } from '@/content/siteContent'
import { API_BASE_URL } from '@/config'
import { fetchWithCache } from '@/utils/dataCache'
import { useServerStatusStore } from '@/stores/serverStatus'
import type { LeaderboardEntry, LeaderboardResponse, LeaderboardType } from '@/types/api'
import { getAvatarUrl } from '@/utils/minecraft'

const statusStore = useServerStatusStore()
const { status } = storeToRefs(statusStore)

const leaderboardTypes = Object.keys(leaderboardLabels) as LeaderboardType[]
const currentType = ref<LeaderboardType>('mining')
const entries = ref<LeaderboardEntry[]>([])
const loading = ref(false)

function isOnline(uuid: string): boolean {
  return status.value?.online_players?.includes(uuid) ?? false
}

async function fetchLeaderboard(type: LeaderboardType): Promise<void> {
  loading.value = true
  currentType.value = type

  try {
    const json = await fetchWithCache<LeaderboardResponse>(`${API_BASE_URL}/api/leaderboards/${type}`)
    entries.value = json.entries ?? []
  } catch (error) {
    console.error('Failed to load leaderboard data', error)
    entries.value = []
  } finally {
    loading.value = false
  }
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

onMounted(() => {
  void fetchLeaderboard(currentType.value)
})
</script>

<template>
  <div class="leaderboards-view container page-shell">
    <header class="page-header animate-entry">
      <span class="page-kicker">Ranked survival data</span>
      <h1>{{ siteContent.leaderboards.title }}</h1>
      <p class="page-lede">{{ siteContent.leaderboards.body }}</p>
    </header>

    <div class="leaderboard-tabs glass-card animate-entry delay-100">
      <button
        v-for="type in leaderboardTypes"
        :key="type"
        :class="['tab-chip', { active: currentType === type }]"
        type="button"
        @click="fetchLeaderboard(type)"
      >
        {{ leaderboardLabels[type] }}
      </button>
    </div>

    <div v-if="loading" class="glass-card state-card">{{ siteContent.leaderboards.loading }}</div>

    <div v-else-if="entries.length === 0" class="glass-card state-card">
      <h2>{{ siteContent.leaderboards.emptyTitle }}</h2>
      <p>{{ siteContent.leaderboards.emptyBody }}</p>
    </div>

    <div v-else class="glass-card table-wrap animate-entry delay-200">
      <table class="leaderboard-table">
        <thead>
          <tr>
            <th>{{ siteContent.leaderboards.columns.rank }}</th>
            <th>{{ siteContent.leaderboards.columns.player }}</th>
            <th>{{ siteContent.leaderboards.columns.score }}</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="entry in entries" :key="entry.uuid">
            <td class="rank-cell">#{{ entry.rank }}</td>
            <td>
              <RouterLink :to="`/player/${entry.uuid}`" class="player-link">
                <img :src="getAvatarUrl(entry.name)" :alt="entry.name" class="avatar" loading="lazy" />
                <span :class="['minecraft-font', { online: isOnline(entry.uuid) }]">{{ entry.name }}</span>
              </RouterLink>
            </td>
            <td class="value-cell">{{ formatValue(entry.value, currentType) }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped>
.leaderboard-tabs {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
  margin-bottom: 1rem;
}

.tab-chip {
  border: 1px solid var(--glass-border);
  border-radius: 999px;
  background: rgba(255, 248, 234, 0.03);
  color: var(--text-muted);
  padding: 0.7rem 1rem;
}

.tab-chip.active {
  border-color: rgba(91, 113, 246, 0.42);
  color: var(--text-strong);
}

.state-card {
  display: grid;
  gap: 1rem;
}

.table-wrap {
  overflow-x: auto;
}

.leaderboard-table {
  width: 100%;
  border-collapse: collapse;
}

.leaderboard-table th,
.leaderboard-table td {
  padding: 1rem 0.75rem;
  border-bottom: 1px solid rgba(255, 248, 234, 0.08);
  text-align: left;
}

.leaderboard-table th {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.78rem;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.rank-cell,
.value-cell {
  font-family: var(--mono);
}

.value-cell {
  text-align: right;
  color: var(--primary);
}

.player-link {
  display: inline-flex;
  align-items: center;
  gap: 0.85rem;
  color: var(--text-main);
}

.avatar {
  width: 36px;
  height: 36px;
  border-radius: 10px;
  image-rendering: pixelated;
}

.online {
  color: #a9d08e;
}
</style>
