<script setup lang="ts">
import { useMediaQuery } from '@/composables/useMediaQuery'
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

const { matches: isMobile } = useMediaQuery('(max-width: 720px)')
</script>

<template>
  <div class="table-wrap animate-entry delay-200">
    <TransitionGroup v-if="isMobile" name="list-fade" tag="div" class="mobile-board-list">
      <article v-for="entry in entries" :key="entry.uuid" class="mobile-board-card hover-rise">
        <div class="mobile-board-topline">
          <span class="mobile-rank">#{{ entry.rank }}</span>
          <span v-if="isOnline(entry.uuid)" class="mobile-presence">Online</span>
        </div>

        <RouterLink :to="`/player/${entry.uuid}`" class="mobile-player-link">
          <img :src="getAvatarUrl(entry.name)" :alt="entry.name" class="avatar" loading="lazy" />
          <div class="mobile-player-copy">
            <span :class="['player-name', 'minecraft-font', { online: isOnline(entry.uuid) }]">{{ entry.name }}</span>
            <span class="mobile-score-label">{{ scoreLabel }}</span>
            <strong class="mobile-card-score">{{ formatValue(entry.value, currentType) }}</strong>
          </div>
        </RouterLink>
      </article>
    </TransitionGroup>

    <table v-else class="leaderboard-table">
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
  border-collapse: separate;
  border-spacing: 0 0.45rem;
}

.leaderboard-table th,
.leaderboard-table td {
  padding: 0.82rem 0.72rem;
  text-align: left;
}

.leaderboard-table thead th {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.68rem;
  letter-spacing: 0.14em;
  text-transform: uppercase;
}

.leaderboard-table tbody tr {
  background: rgba(255, 255, 255, 0.03);
  transition:
    transform var(--transition-panel),
    background var(--transition-fast),
    box-shadow var(--transition-panel);
}

.leaderboard-table tbody tr:hover {
  background: rgba(255, 255, 255, 0.05);
  transform: translateY(-1px);
  box-shadow: 0 16px 32px rgba(0, 0, 0, 0.22);
}

.leaderboard-table tbody td:first-child {
  border-radius: 16px 0 0 16px;
}

.leaderboard-table tbody td:last-child {
  border-radius: 0 16px 16px 0;
}

.rank-cell,
.value-cell {
  font-family: var(--mono);
}

.rank-cell {
  width: 5rem;
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
  width: 36px;
  height: 36px;
  border-radius: 12px;
  image-rendering: pixelated;
}

.player-meta {
  display: grid;
  gap: 0.12rem;
}

.player-name {
  color: var(--text-main);
}

.online {
  color: var(--success);
}

@media (max-width: 720px) {
  .mobile-board-list {
    display: grid;
    gap: 0.75rem;
  }

.mobile-board-card {
    display: grid;
    gap: 0.7rem;
    padding: 0.85rem 0.9rem;
    border-radius: 20px;
    border: 1px solid rgba(255, 255, 255, 0.08);
    background: rgba(255, 255, 255, 0.03);
    box-shadow: var(--panel-shadow);
  }

  .mobile-board-topline {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 0.75rem;
  }

  .mobile-rank,
  .mobile-presence,
  .mobile-score-label {
    font-family: var(--mono);
    font-size: 0.68rem;
    letter-spacing: 0.12em;
    text-transform: uppercase;
  }

  .mobile-rank,
  .mobile-score-label {
    color: var(--text-dim);
  }

  .mobile-presence {
    color: var(--success);
  }

  .mobile-player-link {
    display: flex;
    align-items: center;
    gap: 0.8rem;
    width: 100%;
  }

  .mobile-player-copy {
    display: grid;
    gap: 0.18rem;
    min-width: 0;
  }

  .mobile-card-score {
    color: var(--primary);
    font-size: 1rem;
    font-weight: 600;
  }
}
</style>
