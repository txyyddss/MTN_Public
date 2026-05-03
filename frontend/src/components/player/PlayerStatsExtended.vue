<script setup lang="ts">
import { computed } from 'vue'

import PlayerCollapsiblePanel from '@/components/player/PlayerCollapsiblePanel.vue'
import StatBox from '@/components/StatBox.vue'
import { getStatCategoryLabel, useSiteContent } from '@/content/siteContent'
import { createStatLeaderboardTarget } from '@/utils/leaderboards'
import type { PlayerStatBuckets, StatGroup } from '@/types/api'

const props = defineProps<{
  stats: PlayerStatBuckets | null
  statGroups: StatGroup[]
  activeGroup: StatGroup
  groupCategories: string[]
  selectedCategory: string
  statSearch: string
  filteredStats: Record<string, number>
  ranks: Record<string, number>
  formatStatName: (name: string) => string
  formatStatValue: (value: number, name: string) => string
  getStatIconPath: (category: string, name: string) => string | undefined
}>()

const emit = defineEmits<{
  (event: 'update:activeGroup', value: StatGroup): void
  (event: 'update:selectedCategory', value: string): void
  (event: 'update:statSearch', value: string): void
  (event: 'selectLeaderboard', value: ReturnType<typeof createStatLeaderboardTarget>): void
}>()

interface StatEntry {
  name: string
  label: string
  value: string
  icon: string | undefined
  rank: number | undefined
  target: ReturnType<typeof createStatLeaderboardTarget>
}

const siteContent = useSiteContent()
const activeCategoryLabel = computed(() => {
  const key = props.selectedCategory.replace('minecraft:', '')
  return getStatCategoryLabel(key) ?? key.replace(/_/g, ' ')
})

const statEntries = computed<StatEntry[]>(() =>
  Object.entries(props.filteredStats).map(([name, value]) => {
    const label = props.formatStatName(name)

    return {
      name,
      label,
      value: props.formatStatValue(value, name),
      icon: props.getStatIconPath(props.selectedCategory, name),
      rank: props.ranks[`stat:${props.selectedCategory}:${name}`],
      target: createStatLeaderboardTarget(props.selectedCategory, name, label, (rawValue) => props.formatStatValue(rawValue, name))
    }
  })
)
</script>

<template>
  <PlayerCollapsiblePanel v-if="stats" class="panel-card" :title="siteContent.playerDetail.sections.extendedStats">
    <template #summary>
      <span class="meta-chip">{{ activeCategoryLabel }}</span>
    </template>

    <div class="stats-controls">
      <div class="group-tabs">
        <button
          v-for="group in statGroups"
          :key="group.name"
          :class="['group-tab', { active: activeGroup.name === group.name }]"
          type="button"
          @click="emit('update:activeGroup', group)"
        >
          {{ group.name }}
        </button>
      </div>

      <input
        id="player-stat-search"
        :value="statSearch"
        name="player-stat-search"
        :placeholder="siteContent.playerDetail.sections.searchPlaceholder"
        class="search-input"
        @input="emit('update:statSearch', ($event.target as HTMLInputElement).value)"
      />
    </div>

    <div v-if="groupCategories.length > 0" class="category-tabs">
      <button
        v-for="category in groupCategories"
        :key="category"
        :class="['category-tab', { active: selectedCategory === category }]"
        type="button"
        @click="emit('update:selectedCategory', category)"
      >
        {{ getStatCategoryLabel(category.replace('minecraft:', '')) ?? category.replace('minecraft:', '').replace(/_/g, ' ') }}
      </button>
    </div>

    <div v-if="statEntries.length > 0" class="stat-grid">
      <StatBox
        v-for="entry in statEntries"
        :key="entry.name"
        :name="entry.label"
        :value="entry.value"
        clickable
        :icon="entry.icon"
        :rank="entry.rank"
        @select="emit('selectLeaderboard', entry.target)"
      />
    </div>
    <p v-else class="empty-copy">{{ siteContent.playerDetail.sections.emptyStats }}</p>
  </PlayerCollapsiblePanel>
</template>

<style scoped>
.panel-card {
  display: grid;
  gap: 0.85rem;
}

.stats-controls {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
}

.search-input {
  width: min(300px, 100%);
  min-height: 2.5rem;
  padding: 0 0.9rem;
  border: 1px solid var(--player-glass-border-soft);
  border-radius: 999px;
  background: var(--player-glass-tile-bg);
  color: var(--text-main);
  font-size: 0.9rem;
  box-shadow: var(--glass-inset);
}

.search-input:focus {
  outline: none;
  border-color: var(--player-glass-border-strong);
  background: var(--player-glass-tile-bg-hover);
}

.group-tabs,
.category-tabs {
  display: flex;
  flex-wrap: wrap;
  gap: 0.45rem;
}

.group-tab,
.category-tab {
  position: relative;
  overflow: hidden;
  border: 1px solid var(--player-glass-border-soft);
  border-radius: 999px;
  background: var(--player-glass-tile-bg);
  color: var(--control-text);
  padding: 0.58rem 0.86rem;
  font-family: var(--sans);
  font-size: 0.84rem;
  font-weight: 500;
  letter-spacing: 0;
  text-transform: capitalize;
  transition:
    transform var(--transition-panel),
    border-color var(--transition-fast),
    color var(--transition-fast),
    background var(--transition-fast);
}

.group-tab::after,
.category-tab::after {
  content: '';
  position: absolute;
  inset: auto 14px 8px;
  height: 1px;
  border-radius: 999px;
  background: var(--control-line);
  transform: scaleX(0);
  transform-origin: center;
  transition: transform var(--transition-panel);
}

.group-tab:hover,
.category-tab:hover {
  color: var(--control-text-hover);
  border-color: var(--player-glass-border-strong);
  background: var(--player-glass-tile-bg-hover);
  transform: translateY(-1px);
}

.group-tab.active,
.category-tab.active {
  border-color: var(--control-border-active);
  background: var(--control-bg-active);
  color: var(--control-text-active);
}

.group-tab.active::after,
.category-tab.active::after {
  transform: scaleX(1);
}

.stat-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 0.65rem;
}

.empty-copy {
  color: var(--text-muted);
}

.meta-chip {
  padding: 0.38rem 0.62rem;
  border-radius: 999px;
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.68rem;
  border: 1px solid var(--player-glass-border-soft);
  background: rgba(var(--secondary-rgb), 0.1);
  text-transform: capitalize;
}

@media (max-width: 980px) {
  .stats-controls,
  .stat-grid {
    grid-template-columns: 1fr;
  }

  .stats-controls {
    align-items: stretch;
    flex-direction: column;
  }
}

@media (max-width: 640px) {
  .stat-grid {
    grid-template-columns: 1fr;
  }
}
</style>
