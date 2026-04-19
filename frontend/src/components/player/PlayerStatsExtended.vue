<script setup lang="ts">
import StatBox from '@/components/StatBox.vue'

defineProps<{
  stats: any
  statGroups: any[]
  activeGroup: any
  groupCategories: string[]
  selectedCategory: string
  statSearch: string
  filteredStats: Record<string, number>
  formatStatName: (name: string) => string
  formatStatValue: (val: number, name: string) => string
  getStatIconPath: (category: string, name: string) => string | undefined
}>()

const emit = defineEmits(['update:activeGroup', 'update:selectedCategory', 'update:statSearch'])
</script>

<template>
  <section class="panel glass-card extended-stats" v-if="stats">
    <div class="stats-header">
      <h3><img src="/mc_icons/blocks/special/bookshelf^48.png" class="header-icon" /> Extended Statistics</h3>
      <div class="search-wrap">
        <input 
          type="text" 
          :value="statSearch" 
          @input="emit('update:statSearch', ($event.target as HTMLInputElement).value)"
          placeholder="Search stats... (e.g. 'diamond')" 
          class="stat-search-input"
        />
      </div>
    </div>

    <!-- Category Groups Tabs -->
    <div class="stat-groups">
      <button 
        v-for="group in statGroups" 
        :key="group.name"
        :class="['group-tab', { active: activeGroup.name === group.name }]"
        @click="emit('update:activeGroup', group)"
      >
        {{ group.name }}
      </button>
    </div>

    <!-- Secondary Category Pills -->
    <div class="category-pills" v-if="groupCategories.length > 0">
      <button 
        v-for="cat in groupCategories" 
        :key="cat"
        :class="['pill', { active: selectedCategory === cat }]"
        @click="emit('update:selectedCategory', cat)"
      >
        {{ cat.replace('minecraft:', '').replace(/_/g, ' ') }}
      </button>
    </div>

    <!-- Stats Display Area -->
    <div class="stats-viewport">
        <div v-if="Object.keys(filteredStats).length > 0" class="stat-grid">
            <StatBox 
                v-for="(val, name) in filteredStats" 
                :key="name"
                :name="formatStatName(String(name))"
                :value="formatStatValue(Number(val), String(name))"
                :icon="getStatIconPath(selectedCategory, String(name))"
            />
        </div>
        <div v-else class="no-results">
            <p>No statistics found in this category matching your search.</p>
        </div>
    </div>
  </section>
</template>

<style scoped>
.stats-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 1rem;
  margin-bottom: 2rem;
}

.stats-header h3 {
  display: flex;
  align-items: center;
  gap: 12px;
  margin: 0;
}

.header-icon {
  width: 24px;
  height: 24px;
}

.stat-search-input {
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid var(--glass-border);
  color: #fff;
  padding: 10px 20px;
  border-radius: 100px;
  width: 100%;
  min-width: 280px;
  outline: none;
  transition: all 0.3s;
}

.stat-search-input:focus {
  border-color: var(--primary);
  background: rgba(255, 255, 255, 0.08);
  box-shadow: 0 0 20px rgba(59, 130, 246, 0.15);
}

.stat-groups {
  display: flex;
  gap: 0.5rem;
  margin-bottom: 1.5rem;
  overflow-x: auto;
  padding-bottom: 5px;
}

.group-tab {
  padding: 10px 20px;
  background: transparent;
  border: 1px solid var(--glass-border);
  border-radius: 8px;
  color: var(--text-muted);
  font-family: var(--heading);
  font-weight: 700;
  cursor: pointer;
  white-space: nowrap;
  transition: all 0.3s;
}

.group-tab.active {
  background: var(--primary);
  color: #000;
  border-color: var(--primary);
}

.category-pills {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin-bottom: 2rem;
}

.pill {
  padding: 6px 14px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid var(--glass-border);
  border-radius: 100px;
  color: var(--text-muted);
  font-size: 0.85rem;
  cursor: pointer;
  text-transform: capitalize;
  transition: all 0.3s;
}

.pill.active {
  background: rgba(59, 130, 246, 0.1);
  color: var(--primary);
  border-color: var(--primary);
}

.stat-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 1.5rem;
}

.no-results {
  text-align: center;
  padding: 4rem 2rem;
  color: var(--text-muted);
  background: rgba(0, 0, 0, 0.2);
  border-radius: var(--radius-lg);
}
</style>
