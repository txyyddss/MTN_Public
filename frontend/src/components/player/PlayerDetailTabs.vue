<script setup lang="ts">
import { useSiteContent } from '@/content/siteContent'
import type { PlayerDetailTab, PlayerDetailTabOption } from '@/types/playerDetail'

defineProps<{
  tabs: PlayerDetailTabOption[]
}>()

const activeTab = defineModel<PlayerDetailTab>('activeTab', { required: true })
const siteContent = useSiteContent()
</script>

<template>
  <nav class="detail-tabs glass-card player-glass-card player-glass-reveal-soft" :aria-label="siteContent.playerDetail.tabsAria">
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
  border-color: var(--player-glass-border);
  background: var(--player-glass-bg);
  box-shadow: var(--player-glass-shadow), var(--glass-inset);
}

.detail-tab {
  position: relative;
  overflow: hidden;
  min-height: 3rem;
  border: 1px solid var(--player-glass-border-soft);
  border-radius: 18px;
  background: var(--player-glass-tile-bg);
  color: var(--control-text);
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
  height: 1px;
  border-radius: 999px;
  background: var(--control-line);
  transform: scaleX(0);
  transform-origin: center;
  transition: transform var(--transition-panel);
}

.detail-tab:hover {
  color: var(--control-text-hover);
  border-color: var(--player-glass-border-strong);
  background: var(--player-glass-tile-bg-hover);
  transform: translateY(-1px);
}

.detail-tab.active {
  border-color: var(--control-border-active);
  background:
    linear-gradient(135deg, rgba(var(--primary-rgb), 0.24), rgba(var(--secondary-rgb), 0.11)),
    rgba(255, 255, 255, 0.05);
  color: var(--control-text-active);
  box-shadow: var(--glass-inset), 0 10px 24px rgba(0, 0, 0, 0.18);
}

.detail-tab.active::after {
  transform: scaleX(1);
}

.detail-tab-label {
  font-family: var(--sans);
  font-size: 0.88rem;
  font-weight: 500;
  letter-spacing: 0;
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
    font-size: 0.78rem;
    letter-spacing: 0;
  }
}
</style>
