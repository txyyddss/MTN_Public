<script setup lang="ts">
import { getAvatarUrl } from '@/utils/minecraft'
import type { LeaderboardEntry, LeaderboardType } from '@/types/api'

defineProps<{
  entries: LeaderboardEntry[]
  currentType: LeaderboardType
  formatValue: (value: number, type: LeaderboardType) => string
  isOnline: (uuid: string) => boolean
  rankLabel: string
  playerLabel: string
  scoreLabel: string
}>()
</script>

<template>
  <div class="table-wrap animate-entry delay-200">
    <table class="leaderboard-table">
      <thead>
        <tr>
          <th>{{ rankLabel }}</th>
          <th>{{ playerLabel }}</th>
          <th>{{ scoreLabel }}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="entry in entries" :key="entry.uuid">
          <td class="rank-cell">#{{ entry.rank }}</td>
          <td class="player-cell">
            <RouterLink :to="`/player/${entry.uuid}`" class="player-link">
              <img :src="getAvatarUrl(entry.name)" :alt="entry.name" class="avatar" loading="lazy" />
              <span class="player-meta">
                <span :class="['player-name', 'minecraft-font', { online: isOnline(entry.uuid) }]">{{ entry.name }}</span>
                <span class="mobile-value">{{ formatValue(entry.value, currentType) }}</span>
              </span>
            </RouterLink>
          </td>
          <td class="value-cell">{{ formatValue(entry.value, currentType) }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.table-wrap {
  overflow-x: auto;
}

.leaderboard-table {
  width: 100%;
  border-collapse: collapse;
}

.leaderboard-table th,
.leaderboard-table td {
  padding: 0.95rem 0.7rem;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  text-align: left;
}

.leaderboard-table th {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.72rem;
  letter-spacing: 0.14em;
  text-transform: uppercase;
}

.rank-cell,
.value-cell {
  font-family: var(--mono);
}

.rank-cell {
  width: 4.5rem;
  color: var(--text-strong);
}

.value-cell {
  text-align: right;
  color: var(--primary);
}

.player-link {
  display: inline-flex;
  align-items: center;
  gap: 0.8rem;
  color: var(--text-main);
}

.avatar {
  width: 34px;
  height: 34px;
  border-radius: 10px;
  image-rendering: pixelated;
}

.player-meta {
  display: grid;
  gap: 0.18rem;
}

.player-name {
  color: var(--text-main);
}

.mobile-value {
  display: none;
  color: var(--primary);
  font-family: var(--mono);
  font-size: 0.78rem;
}

.online {
  color: #8fe3b3;
}

@media (max-width: 720px) {
  .leaderboard-table,
  .leaderboard-table tbody,
  .leaderboard-table tr,
  .leaderboard-table td {
    display: block;
    width: 100%;
  }

  .leaderboard-table thead,
  .value-cell {
    display: none;
  }

  .leaderboard-table tr {
    padding: 0.85rem 0;
    border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  }

  .leaderboard-table td {
    border: 0;
    padding: 0;
  }

  .rank-cell {
    margin-bottom: 0.45rem;
    color: var(--text-dim);
    font-size: 0.78rem;
  }

  .player-link {
    width: 100%;
  }

  .mobile-value {
    display: block;
  }
}
</style>
