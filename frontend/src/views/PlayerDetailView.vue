<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue'
import { useRoute } from 'vue-router'
import { preloadImages, PreloadPriority } from '@/utils/preloader'
import { useServerStatus } from '@/composables/useServerStatus'
import { usePlayerDetail } from '@/composables/usePlayerDetail'
import { usePlayerStats } from '@/composables/usePlayerStats'
import { useAdvancements } from '@/composables/useAdvancements'
import { getSkinUrl } from '@/utils/minecraft'

// Import new sub-components
import PlayerSidebar from '@/components/player/PlayerSidebar.vue'
import PlayerSkills from '@/components/player/PlayerSkills.vue'
import PlayerAdvancementList from '@/components/player/PlayerAdvancementList.vue'
import PlayerStatsExtended from '@/components/player/PlayerStatsExtended.vue'
import PlayerPieChart from '@/components/PlayerPieChart.vue'

const route = useRoute()
const { status } = useServerStatus()
const uuid = computed(() => route.params.uuid as string)

const {
  loading, info, stats, advancements, mcmmo, linkedAccount, oreStats, ranks, fetchDetail
} = usePlayerDetail(uuid.value)

const {
  formatStatName, formatStatValue, getStatIconPath, getFilteredStats, getFilteredMcmmo
} = usePlayerStats(stats)

const {
  totalAdvancements, completedAdvancements, categorizedAdvancements, getAdvancementMetadata, getAdvIconPath
} = useAdvancements(advancements)

const onlinePlayers = computed<string[]>(() => status.value?.online_players || [])
const isOnline = computed(() => onlinePlayers.value.includes(uuid.value))

const selectedCategory = ref<string>('minecraft:custom')
const statSearch = ref('')

const filteredStats = computed(() => getFilteredStats(selectedCategory.value, statSearch.value))
const filteredMcmmo = computed(() => getFilteredMcmmo(mcmmo.value))

onMounted(() => {
  fetchDetail()
})

watch(uuid, (newUuid) => {
  if (newUuid) {
    fetchDetail()
  }
})

const formatDate = (ms: number) => ms ? new Date(ms).toLocaleDateString() : 'N/A'
const formatPlaytime = (ticks: number) => ticks ? `${(ticks / 20 / 3600).toFixed(1)}h` : '0h'

const statGroups = [
  { name: 'Global Statistics', categories: ['minecraft:custom'] },
  { name: 'Items & Blocks', categories: ['minecraft:mined', 'minecraft:broken', 'minecraft:crafted', 'minecraft:used', 'minecraft:picked_up', 'minecraft:dropped'] },
  { name: 'Combat & Mobs', categories: ['minecraft:killed', 'minecraft:killed_by'] }
]

const activeGroup = ref(statGroups[0])

const groupCategories = computed(() => {
  if (!stats.value) return []
  return activeGroup.value.categories.filter(cat => stats.value[cat] && cat !== 'minecraft:custom')
})

watch(activeGroup, (newGroup) => {
  if (!stats.value) return
  const cats = newGroup.categories.filter(cat => stats.value[cat])
  if (cats.length > 0 && !cats.includes(selectedCategory.value)) {
    selectedCategory.value = cats[0]
  }
})

// REALTIME PRELOAD LOGIC
watch(info, (newInfo) => {
  if (newInfo) {
    preloadImages([getSkinUrl(newInfo.last_known_name, newInfo.type)], PreloadPriority.HIGH)
  }
})

watch(categorizedAdvancements, (categories) => {
  if (categories) {
    const urls: string[] = []
    Object.values(categories).forEach(items => {
      items.forEach(adv => urls.push(getAdvIconPath(adv.key)))
    })
    preloadImages(urls, PreloadPriority.MEDIUM_HIGH)
  }
}, { immediate: true })
</script>

<template>
  <div class="player-detail container animate-entry">
    <div v-if="loading" class="loading-state">
      <div class="spinner"></div>
      <p>Loading player data...</p>
    </div>
    
    <div v-else-if="!info" class="empty-state glass-card">
      <h3>Player not found</h3>
      <p>Data could not be loaded for the given UUID.</p>
      <router-link to="/players" class="btn-primary mt-4">Back to Players</router-link>
    </div>
    
    <div v-else class="content-grid">
      <!-- Left Column: Profile Sidebar -->
      <PlayerSidebar 
        :info="info" 
        :is-online="isOnline" 
        :linked-account="linkedAccount"
        :format-date="formatDate"
        :format-playtime="formatPlaytime"
      />

      <!-- Right Column: Stats Sections -->
      <main class="details-section">
        <!-- Ores Pie Chart -->
        <PlayerPieChart :oreStats="oreStats" />

        <!-- McMMO Skills -->
        <PlayerSkills 
          :mcmmo="mcmmo"
          :ranks="ranks"
          :filtered-mcmmo="filteredMcmmo"
        />

        <!-- Advancements -->
        <PlayerAdvancementList 
          :advancements="advancements"
          :completed-advancements="completedAdvancements"
          :total-advancements="totalAdvancements"
          :categorized-advancements="categorizedAdvancements"
          :get-advancement-metadata="getAdvancementMetadata"
          :get-adv-icon-path="getAdvIconPath"
        />

        <!-- Extended Stats -->
        <PlayerStatsExtended 
          v-model:activeGroup="activeGroup"
          v-model:selectedCategory="selectedCategory"
          v-model:statSearch="statSearch"
          :stats="stats"
          :statGroups="statGroups"
          :groupCategories="groupCategories"
          :filtered-stats="filteredStats"
          :format-stat-name="formatStatName"
          :format-stat-value="formatStatValue"
          :get-stat-icon-path="getStatIconPath"
        />
      </main>
    </div>
  </div>
</template>

<style scoped>
.player-detail {
  padding-top: 2rem;
  padding-bottom: 4rem;
}

.content-grid {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

@media (min-width: 900px) {
  .content-grid {
    flex-direction: row;
    align-items: flex-start;
  }
  .profile-card {
    width: 260px;
    flex-shrink: 0;
    position: sticky;
    top: 100px;
  }
}

.profile-card {
  padding: 2rem;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.avatar-header {
    text-align: center;
    margin-bottom: 2rem;
}

.skin-wrapper {
    width: 250px;
    height: 300px;
    margin-bottom: 1rem;
    filter: drop-shadow(0 15px 30px rgba(0,0,0,0.5));
}

.profile-name {
  color: #fff;
  font-size: 1.8rem;
  margin-bottom: 0.5rem;
  transition: color 0.3s ease;
  word-wrap: break-word;
  overflow-wrap: break-word;
  max-width: 100%;
}

.profile-name.online {
  color: #10b981;
  text-shadow: 0 0 15px rgba(16, 185, 129, 0.3);
}

.type-tag {
    font-size: 0.75rem;
    padding: 2px 8px;
    border-radius: 4px;
    font-weight: 800;
}
.type-tag.java { background: rgba(59, 130, 246, 0.2); color: #93c5fd; }
.type-tag.bedrock { background: rgba(79, 70, 229, 0.2); color: #c7d2fe; }

.basic-info {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.info-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 8px;
  border-bottom: 1px solid var(--glass-border);
}

.info-row .label {
  color: var(--text-muted);
  font-size: 0.9rem;
}

.info-row .val {
  font-weight: 600;
  color: var(--text-main);
  text-align: right;
  font-size: 0.9rem;
}

.account-link {
    color: var(--primary) !important;
}

.lvl-badge {
  background: rgba(59, 130, 246, 0.2);
  color: var(--primary);
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 0.8rem;
}

.details-section {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 2rem;
  min-width: 0;
}

.panel h3 {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 1.5rem;
  font-size: 1.4rem;
  color: #fff;
  font-family: var(--heading);
  letter-spacing: -0.5px;
}

.header-icon {
    width: 24px;
    height: 24px;
    image-rendering: auto;
}

.panel-header-simple {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}
.panel-header-simple h3 { margin: 0; }

.ore-content {
    display: flex;
    gap: 2rem;
    align-items: center;
}
@media (max-width: 600px) {
    .ore-content { flex-direction: column; }
    .mobile-hide { display: none !important; }
}

.chart-container {
  height: 200px;
  width: 200px;
  flex-shrink: 0;
}

.ore-list {
  flex: 1;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 8px;
}

.ore-item {
  background: rgba(255, 255, 255, 0.03);
  padding: 8px 12px;
  border-radius: 8px;
  display: flex;
  justify-content: space-between;
}

.ore-name { color: var(--text-muted); font-size: 0.85rem; }
.ore-val { font-weight: 700; color: #fff; font-size: 0.85rem; }

.skill-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
  gap: 4px;
}

.total-badge {
    color: var(--primary);
    font-weight: 800;
}

.rank-badge {
    background: rgba(245, 158, 11, 0.1);
    color: #fcd34d;
    padding: 2px 10px;
    border-radius: 20px;
    font-size: 0.9rem;
    font-weight: 700;
}

.rank-group {
    display: flex;
    gap: 12px;
    color: var(--text-muted);
}

.adv-category { margin-bottom: 2rem; }
.category-name {
  font-size: 0.85rem;
  color: var(--text-muted);
  margin-bottom: 1rem;
  text-transform: uppercase;
  letter-spacing: 1.5px;
  font-weight: 700;
  opacity: 0.8;
}

.advancements-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 10px;
}

.adv-card {
  display: flex;
  align-items: center;
  gap: 12px;
  background: rgba(255, 255, 255, 0.03);
  padding: 10px 14px;
  border-radius: var(--radius-md);
  border: 1px solid var(--glass-border);
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  cursor: default;
}

.adv-card:hover { 
  background: rgba(255, 255, 255, 0.07);
  border-color: rgba(255, 255, 255, 0.2);
  transform: translateY(-2px);
  box-shadow: 0 10px 20px rgba(0,0,0,0.2);
}

.adv-icon-wrap {
  width: 40px;
  height: 40px;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(0,0,0,0.4);
  border-radius: 8px;
  padding: 6px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  position: relative;
}

.adv-icon-wrap::after {
  content: '';
  position: absolute;
  inset: 2px;
  border-radius: 6px;
  border: 1px solid rgba(255, 255, 255, 0.05);
  pointer-events: none;
}

.adv-icon-wrap.goal { border-color: rgba(251, 191, 36, 0.4); box-shadow: inset 0 0 10px rgba(251, 191, 36, 0.1); }
.adv-icon-wrap.challenge { border-color: rgba(244, 63, 94, 0.4); box-shadow: inset 0 0 10px rgba(244, 63, 94, 0.1); }

.adv-icon { 
  width: 24px; 
  height: 24px; 
  image-rendering: auto; /* Changed from pixelated to auto for smoother look of some icons, but keeping it optional */
  filter: drop-shadow(0 2px 4px rgba(0,0,0,0.5));
}

.adv-info { 
  display: flex; 
  flex-direction: column; 
  justify-content: center; 
  min-width: 0; 
  flex: 1; 
}

.adv-name { 
  font-weight: 600; 
  color: #fff; 
  font-size: 0.85rem;
  line-height: 1.3;
  word-wrap: break-word;
  overflow-wrap: break-word;
  white-space: normal;
}

.stat-value { font-weight: 800; color: #fff; font-size: 1.1rem; font-family: var(--heading); }

.rank-badge.mini { font-size: 0.7rem; padding: 2px 8px; }


/* Duplicate class removal and consolidation */

.tabs-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 1rem;
    margin-bottom: 1.5rem;
    flex-wrap: wrap;
}

.search-wrapper {
    position: relative;
    display: flex;
    align-items: center;
    flex: 1;
    max-width: 250px;
    min-width: 150px;
}

@media (max-width: 600px) {
    .search-wrapper {
        min-width: 100%;
        max-width: none;
    }
}

.search-icon {
    position: absolute;
    left: 12px;
    width: 16px;
    height: 16px;
    color: var(--text-muted);
    pointer-events: none;
}

.stat-search-box {
    background: rgba(0,0,0,0.2);
    border: 1px solid var(--glass-border);
    padding: 8px 16px 8px 36px;
    border-radius: 20px;
    font-size: 0.9rem;
    color: #fff;
    width: 100%;
    transition: all 0.3s;
}

.stat-search-box:focus {
    outline: none;
    border-color: var(--primary);
    background: rgba(0,0,0,0.4);
    box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.2);
}

.tabs {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.group-tabs {
    display: flex;
    gap: 4px;
    background: rgba(0,0,0,0.2);
    padding: 4px;
    border-radius: 12px;
    margin-bottom: 1.5rem;
    width: fit-content;
}

.group-btn {
    background: transparent;
    border: none;
    color: var(--text-muted);
    padding: 8px 16px;
    border-radius: 8px;
    cursor: pointer;
    font-size: 0.9rem;
    font-weight: 600;
    transition: all 0.2s;
}

.group-btn.active {
    background: var(--primary);
    color: #000;
}

.tab-btn {
  background: transparent;
  border: 1px solid var(--glass-border);
  color: var(--text-muted);
  padding: 6px 14px;
  border-radius: 20px;
  cursor: pointer;
  font-size: 0.85rem;
}

.tab-btn:hover { background: rgba(255,255,255,0.05); }

.tab-btn.active {
  background: rgba(59, 130, 246, 0.1);
  border-color: var(--primary);
  color: var(--primary);
}

.stat-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(130px, 1fr));
  gap: 8px;
  max-height: 400px;
  overflow-y: auto;
  padding-right: 8px;
}

.empty-mini { text-align: center; color: var(--text-muted); padding: 2rem; }

.stat-item {
  background: rgba(255,255,255,0.02);
  border-radius: 6px;
  padding: 8px 12px;
  display: flex;
  justify-content: space-between;
}

.stat-name { font-size: 0.8rem; color: var(--text-muted); text-transform: capitalize; }
.stat-value { font-weight: 700; color: #fff; font-size: 0.9rem; }

.custom-stat-list {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 12px;
    max-height: 500px;
    overflow-y: auto;
    padding-right: 12px;
}

.custom-stat-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background: rgba(255, 255, 255, 0.03);
    padding: 10px 16px;
    border-radius: 8px;
    border: 1px solid var(--glass-border);
    transition: all 0.2s;
}

.custom-stat-item:hover {
    background: rgba(255, 255, 255, 0.05);
    border-color: rgba(255, 255, 255, 0.2);
    transform: translateX(4px);
}

.stat-label {
    font-size: 0.9rem;
    color: var(--text-muted);
    font-weight: 500;
}

.stat-value-wrap {
    display: flex;
    align-items: center;
    gap: 12px;
}

.stat-rank {
    font-size: 0.7rem;
    background: rgba(245, 158, 11, 0.1);
    color: #fcd34d;
    padding: 2px 8px;
    border-radius: 4px;
    font-weight: 700;
}

.stat-val {
    font-weight: 700;
    color: var(--primary);
    font-family: var(--mono);
    font-size: 1rem;
}

.loading-state, .empty-state { text-align: center; padding: 5rem 2rem;}

.spinner {
  width: 40px; height: 40px;
  border: 3px solid rgba(59, 130, 246, 0.1);
  border-left-color: var(--primary);
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 1rem;
}
@keyframes spin { 100% { transform: rotate(360deg); } }
</style>
