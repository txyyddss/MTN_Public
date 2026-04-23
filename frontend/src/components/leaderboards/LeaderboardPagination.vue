<script setup lang="ts">
import { computed } from 'vue'

import { useSiteContent } from '@/content/siteContent'

const props = defineProps<{
  currentPage: number
  totalPages: number
}>()

defineEmits<{
  (event: 'previous'): void
  (event: 'next'): void
}>()

const siteContent = useSiteContent()
const statusLabel = computed(() =>
  siteContent.value.leaderboards.paginationStatus
    .replace('{page}', String(props.currentPage))
    .replace('{total}', String(props.totalPages))
)
</script>

<template>
  <nav class="pager animate-entry delay-300" :aria-label="siteContent.leaderboards.paginationAria">
    <button
      type="button"
      class="pager-btn"
      :disabled="currentPage <= 1"
      @click="$emit('previous')"
    >
      {{ siteContent.leaderboards.paginationPrev }}
    </button>

    <span class="pager-status">{{ statusLabel }}</span>

    <button
      type="button"
      class="pager-btn"
      :disabled="currentPage >= totalPages"
      @click="$emit('next')"
    >
      {{ siteContent.leaderboards.paginationNext }}
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
  border: 1px solid var(--control-border);
  border-radius: 999px;
  background: var(--control-bg);
  color: var(--control-text);
  padding: 0.54rem 0.84rem;
  font-family: var(--sans);
  font-size: 0.86rem;
  font-weight: 500;
  letter-spacing: -0.01em;
  transition:
    border-color var(--transition-fast),
    transform var(--transition-fast),
    opacity var(--transition-fast);
}

.pager-btn:hover:not(:disabled) {
  border-color: var(--control-border-active);
  background: var(--control-bg-hover);
  color: var(--control-text-hover);
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
