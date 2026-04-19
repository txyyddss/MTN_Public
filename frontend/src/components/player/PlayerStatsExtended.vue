<script setup lang="ts">
import { computed } from 'vue'

import PlayerCollapsiblePanel from '@/components/player/PlayerCollapsiblePanel.vue'
import StatBox from '@/components/StatBox.vue'
import { siteContent } from '@/content/siteContent'
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
}>()

function rankForStat(name: string): number | undefined {
  return props.ranks[`stat:${props.selectedCategory}:${name}`]
}

const activeCategoryLabel = computed(() => props.selectedCategory.replace('minecraft:', '').replace(/_/g, ' '))
</script>

<template>
  <PlayerCollapsiblePanel v-if="stats" class="panel-card" :title="siteContent.playerDetail.sections.extendedStats">
    <template #summary>
      <span class="meta-chip">{{ activeCategoryLabel }}</span>
    </template>

    <div class="panel-head">
      <input
        :value="statSearch"
        :placeholder="siteContent.playerDetail.sections.searchPlaceholder"
        class="search-input"
        @input="emit('update:statSearch', ($event.target as HTMLInputElement).value)"
      />
    </div>

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

    <div v-if="groupCategories.length > 0" class="category-tabs">
      <button
        v-for="category in groupCategories"
        :key="category"
        :class="['category-tab', { active: selectedCategory === category }]"
        type="button"
        @click="emit('update:selectedCategory', category)"
      >
        {{ category.replace('minecraft:', '').replace(/_/g, ' ') }}
      </button>
    </div>

    <div v-if="Object.keys(filteredStats).length > 0" class="stat-grid">
      <StatBox
        v-for="(value, name) in filteredStats"
        :key="name"
        :name="formatStatName(name)"
        :value="formatStatValue(value, name)"
        :icon="getStatIconPath(selectedCategory, name)"
        :rank="rankForStat(name)"
      />
    </div>
    <p v-else class="empty-copy">{{ siteContent.playerDetail.sections.emptyStats }}</p>
  </PlayerCollapsiblePanel>
</template>

<style scoped>
.panel-card {
  display: grid;
  gap: 1rem;
}

.panel-head {
  display: flex;
  justify-content: flex-end;
}

.search-input {
  width: min(360px, 100%);
  min-height: 3rem;
  padding: 0 1rem;
  border: 1px solid var(--glass-border);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.03);
  color: var(--text-main);
}

.group-tabs,
.category-tabs {
  display: flex;
  flex-wrap: wrap;
  gap: 0.6rem;
}

.group-tab,
.category-tab {
  border: 1px solid var(--glass-border);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.03);
  color: var(--text-muted);
  padding: 0.65rem 0.9rem;
  text-transform: capitalize;
}

.group-tab.active,
.category-tab.active {
  border-color: rgba(91, 113, 246, 0.42);
  color: var(--text-strong);
}

.stat-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 0.75rem;
}

.empty-copy {
  color: var(--text-muted);
}

.meta-chip {
  padding: 0.45rem 0.7rem;
  border-radius: 999px;
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.74rem;
  border: 1px solid rgba(255, 255, 255, 0.08);
  text-transform: capitalize;
}

@media (max-width: 980px) {
  .panel-head,
  .stat-grid {
    grid-template-columns: 1fr;
  }

  .panel-head {
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
