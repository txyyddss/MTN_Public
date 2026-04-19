<script setup lang="ts">
defineProps<{
  currentPage: number
  totalPages: number
}>()

defineEmits<{
  (event: 'previous'): void
  (event: 'next'): void
}>()
</script>

<template>
  <nav class="pager animate-entry delay-300" aria-label="Leaderboard pagination">
    <button
      type="button"
      class="pager-btn"
      :disabled="currentPage <= 1"
      @click="$emit('previous')"
    >
      Prev
    </button>

    <span class="pager-status">Page {{ currentPage }} / {{ totalPages }}</span>

    <button
      type="button"
      class="pager-btn"
      :disabled="currentPage >= totalPages"
      @click="$emit('next')"
    >
      Next
    </button>
  </nav>
</template>

<style scoped>
.pager {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 0.55rem;
  margin-top: 0.85rem;
}

.pager-status {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.7rem;
  letter-spacing: 0.1em;
  text-transform: uppercase;
}

.pager-btn {
  min-width: 4.2rem;
  border: 1px solid var(--glass-border);
  border-radius: 999px;
  background: rgba(9, 18, 33, 0.92);
  color: var(--text-main);
  padding: 0.54rem 0.84rem;
  font-family: var(--mono);
  font-size: 0.72rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  transition:
    border-color var(--transition-fast),
    transform var(--transition-fast),
    opacity var(--transition-fast);
}

.pager-btn:hover:not(:disabled) {
  border-color: var(--glass-border-strong);
  transform: translateY(-1px);
}

.pager-btn:disabled {
  opacity: 0.35;
  cursor: not-allowed;
}

@media (max-width: 720px) {
  .pager {
    justify-content: space-between;
  }
}
</style>
