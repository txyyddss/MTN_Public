<script setup lang="ts">
import type { PlayerDetailTab, PlayerDetailTabOption } from '@/types/playerDetail'

defineProps<{
  tabs: PlayerDetailTabOption[]
}>()

const activeTab = defineModel<PlayerDetailTab>('activeTab', { required: true })
</script>

<template>
  <nav class="detail-tabs glass-card animate-entry-soft" aria-label="Player detail sections">
    <button
      v-for="tab in tabs"
      :key="tab.value"
      :class="['detail-tab', { active: activeTab === tab.value }]"
      type="button"
      @click="activeTab = tab.value"
    >
      <span class="detail-tab-label">{{ tab.label }}</span>
    </button>
  </nav>
</template>

<style scoped>
.detail-tabs {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 0.55rem;
  padding: 0.55rem;
}

.detail-tab {
  position: relative;
  overflow: hidden;
  min-height: 3rem;
  border: 1px solid transparent;
  border-radius: 18px;
  background: rgba(8, 18, 34, 0.72);
  color: var(--text-muted);
  padding: 0.75rem 0.9rem;
  text-align: center;
  transition:
    transform var(--transition-panel),
    color var(--transition-fast),
    border-color var(--transition-fast),
    background var(--transition-fast);
}

.detail-tab::after {
  content: '';
  position: absolute;
  inset: auto 16px 10px;
  height: 2px;
  border-radius: 999px;
  background: var(--primary);
  transform: scaleX(0);
  transform-origin: center;
  transition: transform var(--transition-panel);
  box-shadow: 0 0 16px rgba(61, 120, 255, 0.45);
}

.detail-tab:hover {
  color: var(--text-main);
  transform: translateY(-1px);
}

.detail-tab.active {
  border-color: rgba(61, 120, 255, 0.28);
  background: rgba(9, 21, 38, 0.94);
  color: var(--text-strong);
}

.detail-tab.active::after {
  transform: scaleX(1);
}

.detail-tab-label {
  font-family: var(--mono);
  font-size: 0.76rem;
  letter-spacing: 0.11em;
  text-transform: uppercase;
}

@media (max-width: 720px) {
  .detail-tabs {
    grid-template-columns: repeat(3, minmax(0, 1fr));
    gap: 0.45rem;
    padding: 0.45rem;
  }

  .detail-tab {
    min-height: 2.8rem;
    padding: 0.7rem 0.45rem;
    border-radius: 16px;
  }

  .detail-tab::after {
    inset: auto 12px 9px;
  }

  .detail-tab-label {
    font-size: 0.68rem;
    letter-spacing: 0.09em;
  }
}
</style>
