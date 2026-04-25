<script setup lang="ts">
import type { WhitelistEntry, WhitelistMutationInput } from '@/types/api'

defineProps<{
  entries: WhitelistEntry[]
  busy: boolean
}>()

const emit = defineEmits<{
  remove: [input: WhitelistMutationInput]
}>()

function removeEntry(entry: WhitelistEntry): void {
  emit('remove', {
    edition: entry.edition,
    playerName: entry.player_name,
    qqId: entry.qq_id
  })
}

function formatDate(value: string): string {
  if (!value) {
    return 'Unavailable'
  }

  const normalized = value.includes('T') ? value : value.replace(' ', 'T')
  const date = new Date(normalized)
  if (Number.isNaN(date.getTime())) {
    return value
  }

  return new Intl.DateTimeFormat(undefined, {
    month: 'short',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date)
}
</script>

<template>
  <section class="entries-panel hud-panel">
    <div class="table-head">
      <span class="hud-kicker">Mirror</span>
      <h2 class="table-title">Active Entries</h2>
    </div>

    <div v-if="entries.length === 0" class="empty-state">
      <h3>No active entries</h3>
      <p>The MySQL mirror has no active rows for the current filter.</p>
    </div>

    <div v-else class="entries-table" role="table" aria-label="Active whitelist entries">
      <div class="table-row table-row-head" role="row">
        <span role="columnheader">Player</span>
        <span role="columnheader">Edition</span>
        <span role="columnheader">QQ Owner</span>
        <span role="columnheader">Updated</span>
        <span role="columnheader">Action</span>
      </div>

      <div v-for="entry in entries" :key="entry.id" class="table-row" role="row">
        <span class="player-cell minecraft-font" role="cell">{{ entry.player_name }}</span>
        <span role="cell">
          <span :class="['edition-pill', entry.edition]">{{ entry.edition }}</span>
        </span>
        <span class="mono-cell" role="cell">{{ entry.qq_id || 'admin' }}</span>
        <span class="date-cell" role="cell">{{ formatDate(entry.updated_at) }}</span>
        <span role="cell">
          <button class="remove-btn action-inline action-press" type="button" :disabled="busy" @click="removeEntry(entry)">
            Remove
          </button>
        </span>
      </div>
    </div>
  </section>
</template>

<style scoped>
.entries-panel,
.table-head {
  display: grid;
  gap: 1rem;
}

.table-title {
  font-size: 2rem;
  max-width: none;
}

.empty-state {
  display: grid;
  gap: 0.5rem;
  padding: 1rem;
  border: 1px solid var(--glass-border-soft);
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.03);
}

.empty-state h3 {
  font-size: 1.1rem;
  letter-spacing: 0;
}

.empty-state p {
  color: var(--text-muted);
}

.entries-table {
  display: grid;
  gap: 0.45rem;
  overflow-x: auto;
}

.table-row {
  display: grid;
  grid-template-columns: minmax(150px, 1.4fr) minmax(90px, 0.7fr) minmax(130px, 1fr) minmax(130px, 0.9fr) minmax(100px, auto);
  gap: 0.75rem;
  align-items: center;
  min-width: 720px;
  padding: 0.78rem 0.85rem;
  border: 1px solid var(--glass-border-soft);
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.03);
}

.table-row-head {
  min-height: 2.4rem;
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.68rem;
  letter-spacing: 0;
  text-transform: uppercase;
  background: rgba(76, 147, 251, 0.06);
}

.player-cell {
  color: var(--text-strong);
  font-size: 1rem;
}

.mono-cell,
.date-cell {
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.78rem;
}

.edition-pill {
  display: inline-flex;
  min-width: 4.5rem;
  justify-content: center;
  padding: 0.28rem 0.55rem;
  border-radius: 8px;
  border: 1px solid var(--chip-border);
  color: var(--text-main);
  font-family: var(--mono);
  font-size: 0.68rem;
  letter-spacing: 0;
  text-transform: uppercase;
}

.edition-pill.java {
  border-color: rgba(131, 211, 167, 0.26);
  background: rgba(131, 211, 167, 0.08);
  color: var(--success);
}

.edition-pill.bedrock {
  border-color: rgba(242, 200, 110, 0.28);
  background: rgba(242, 200, 110, 0.08);
  color: var(--warning);
}

.remove-btn {
  min-height: 2.25rem;
  padding: 0.45rem 0.75rem;
  border: 1px solid rgba(255, 139, 139, 0.22);
  border-radius: 8px;
  background: rgba(255, 139, 139, 0.08);
  color: var(--danger);
  font-size: 0.82rem;
  font-weight: 600;
}

.remove-btn:disabled {
  opacity: 0.45;
}
</style>
