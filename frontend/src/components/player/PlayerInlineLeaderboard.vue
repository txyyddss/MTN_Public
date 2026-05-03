<script setup lang="ts">
import { computed, watch } from 'vue'

import LeaderboardPagination from '@/components/leaderboards/LeaderboardPagination.vue'
import LeaderboardTable from '@/components/leaderboards/LeaderboardTable.vue'
import { useLeaderboards } from '@/composables/useLeaderboards'
import { useSiteContent } from '@/content/siteContent'
import type { LeaderboardTarget } from '@/types/api'

const props = defineProps<{
  target: LeaderboardTarget
  isOnline: (uuid: string) => boolean
}>()

const emit = defineEmits<{
  (event: 'close'): void
}>()
const siteContent = useSiteContent()

const { currentPage, entries, goToNextPage, goToPreviousPage, loading, totalPages, setKey } = useLeaderboards(props.target.key)

watch(
  () => props.target.key,
  (key) => {
    setKey(key)
  }
)

const resultLabel = computed(() => {
  if (loading.value) {
    return siteContent.value.leaderboards.loading
  }

  return siteContent.value.leaderboards.visibleEntries.replace('{count}', String(entries.value.length))
})

function formatEntryValue(value: number): string {
  if (props.target.formatValue) {
    return props.target.formatValue(value)
  }

  return value.toLocaleString()
}
</script>

<template>
  <section class="glass-card inline-board player-glass-card player-glass-reveal-soft">
    <div class="inline-board-head">
      <div class="inline-board-copy">
        <span class="hud-kicker">{{ siteContent.leaderboards.kicker }}</span>
        <h3 class="inline-board-title">{{ target.title }}</h3>
        <p class="inline-board-note">{{ resultLabel }}</p>
      </div>

      <button type="button" class="btn-secondary inline-board-close" @click="emit('close')">
        {{ siteContent.leaderboards.close }}
      </button>
    </div>

    <div v-if="loading" class="inline-board-skeleton" aria-hidden="true">
      <div v-for="index in 5" :key="index" class="skeleton-board-row">
        <span class="skeleton-block skeleton-rank"></span>
        <div class="skeleton-player">
          <span class="skeleton-avatar"></span>
          <span class="skeleton-line skeleton-player-line"></span>
        </div>
        <span class="skeleton-line skeleton-score"></span>
      </div>
    </div>

    <div v-else-if="entries.length === 0" class="hud-state-card">
      <h3>{{ siteContent.leaderboards.emptyTitle }}</h3>
      <p>{{ siteContent.leaderboards.emptyBody }}</p>
    </div>

    <div v-else class="inline-board-body">
      <LeaderboardTable
        :entries="entries"
        :format-value="formatEntryValue"
        :is-online="isOnline"
        :player-label="siteContent.leaderboards.columns.player"
        :rank-label="siteContent.leaderboards.columns.rank"
        :score-label="target.scoreLabel"
      />
      <LeaderboardPagination
        :current-page="currentPage"
        :total-pages="totalPages"
        @next="goToNextPage"
        @previous="goToPreviousPage"
      />
    </div>
  </section>
</template>

<style scoped>
.inline-board {
  display: grid;
  gap: 0.9rem;
  border-color: var(--player-glass-border);
  background: var(--player-glass-bg);
  box-shadow: var(--player-glass-shadow), var(--glass-inset);
}

.inline-board-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 1rem;
}

.inline-board-copy {
  display: grid;
  gap: 0.45rem;
}

.inline-board-title {
  font-size: 1.4rem;
}

.inline-board-note {
  color: var(--text-muted);
}

.inline-board-close {
  min-height: 2.55rem;
  padding-inline: 1rem;
}

.inline-board-skeleton {
  display: grid;
  gap: 0.65rem;
}

.skeleton-board-row {
  display: grid;
  grid-template-columns: 4.5rem minmax(0, 1fr) 7rem;
  gap: 0.85rem;
  align-items: center;
  padding: 0.9rem 0.95rem;
  border-radius: 16px;
  background: var(--player-glass-tile-bg);
  border: 1px solid var(--player-glass-border-soft);
  box-shadow: var(--glass-inset);
}

.skeleton-player {
  display: flex;
  align-items: center;
  gap: 0.8rem;
}

.skeleton-rank {
  width: 3.5rem;
  height: 0.85rem;
}

.skeleton-player-line {
  width: min(180px, 100%);
}

.skeleton-score {
  width: 100%;
}

@media (max-width: 720px) {
  .inline-board-head {
    flex-direction: column;
  }

  .skeleton-board-row {
    grid-template-columns: 1fr;
  }
}
</style>
