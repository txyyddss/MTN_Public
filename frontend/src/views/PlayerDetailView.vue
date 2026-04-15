<script setup lang="ts">
import { ref, onMounted, onUnmounted, nextTick, computed } from 'vue'
import { useRoute } from 'vue-router'
import { SkinViewer, IdleAnimation } from 'skinview3d'
import Chart from 'chart.js/auto'

const route = useRoute()
const uuid = route.params.uuid as string
const loading = ref(true)

const info = ref<any>(null)
const stats = ref<any>(null)
const advancements = ref<any>(null)
const mcmmo = ref<any>(null)
const linkedAccount = ref<any>(null)
const oreStats = ref<any>([])

const showExactValues = ref(false)
const selectedCategory = ref<string>('minecraft:custom')

const filteredMcmmo = computed(() => {
  if (!mcmmo.value) return {}
  const result: Record<string, any> = {}
  for (const [key, val] of Object.entries(mcmmo.value)) {
    if (['user_id', 'user', 'uuid'].includes(key)) continue
    result[key.charAt(0).toUpperCase() + key.slice(1)] = val
  }
  return result
})

const formatNumber = (num: number) => {
  if (showExactValues.value) return num.toLocaleString()
  if (num >= 1000000) return (num / 1000000).toFixed(1) + 'M'
  if (num >= 1000) return (num / 1000).toFixed(1) + 'K'
  return num.toString()
}

const formatStatName = (name: string) => {
  return name.replace('minecraft:', '').replace(/_/g, ' ')
}

const viewerCanvas = ref<HTMLCanvasElement | null>(null)
const pieChartCanvas = ref<HTMLCanvasElement | null>(null)
let viewer: SkinViewer | null = null
let chartInstance: Chart | null = null

const initSkinViewer = () => {
  if (!viewerCanvas.value || !info.value) return
  
  const cleanName = info.value.last_known_name?.startsWith('be_')
    ? info.value.last_known_name.substring(3)
    : (info.value.last_known_name || 'Steve')
    
  viewer = new SkinViewer({
    canvas: viewerCanvas.value,
    width: 260,
    height: 380,
    skin: `https://mineskin.eu/skin/${cleanName}`
  })
  viewer.animation = new IdleAnimation()
  viewer.autoRotate = true
  viewer.autoRotateSpeed = 0.5
}

const initPieChart = () => {
  if (!pieChartCanvas.value || oreStats.value.length === 0) return
  
  if (chartInstance) chartInstance.destroy()

  const labels = oreStats.value.map((o: any) => o.name)
  const minedData = oreStats.value.map((o: any) => o.mined)

  chartInstance = new Chart(pieChartCanvas.value, {
    type: 'pie',
    data: {
      labels,
      datasets: [
        {
          label: 'Mined',
          data: minedData,
          backgroundColor: [
            '#00e676', '#6366f1', '#f43f5e', '#f59e0b', '#3b82f6', 
            '#ec4899', '#8b5cf6', '#14b8a6', '#f97316', '#64748b'
          ],
          borderWidth: 1,
          borderColor: 'rgba(0,0,0,0.5)'
        }
      ]
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: {
          position: 'right',
          labels: { color: '#ccc' }
        }
      }
    }
  })
}

const fetchDetail = async () => {
  try {
    const res = await fetch(`/api/players/${uuid}`)
    const json = await res.json()
    info.value = json.info
    stats.value = json.stats?.stats || null
    advancements.value = json.advancements?.advancements || []
    mcmmo.value = json.mcmmo
    linkedAccount.value = json.linked_account
    oreStats.value = json.ore_stats
    
    if (stats.value && Object.keys(stats.value).length > 0) {
      const keys = Object.keys(stats.value)
      if (keys.includes('minecraft:custom')) selectedCategory.value = 'minecraft:custom'
      else selectedCategory.value = keys[0]
    }
    
    nextTick(() => {
      initSkinViewer()
      initPieChart()
    })
  } catch (e) {
    console.error('Failed to load player detail', e)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchDetail()
})

onUnmounted(() => {
  if (viewer) viewer.dispose()
  if (chartInstance) chartInstance.destroy()
})

const formatDate = (ms: number) => {
  if (!ms) return 'N/A'
  return new Date(ms).toLocaleString()
}

const formatPlaytime = (ticks: number) => {
  if (!ticks) return '0h'
  const hours = (ticks / 20 / 3600).toFixed(1)
  return `${hours}h`
}

// Compute total advancement completion
const totalAdvancements = computed(() => advancements.value?.length || 0)
const completedAdvancements = computed(() => advancements.value?.filter((a: any) => a.done).length || 0)

const showLockedAdvancements = ref(false)

const categorizedAdvancements = computed(() => {
  if (!advancements.value) return {}
  const result: Record<string, any[]> = {}
  for (const adv of advancements.value) {
    if (!showLockedAdvancements.value && !adv.done) continue
    
    let category = 'Others'
    const parts = adv.key.split('/')
    if (parts.length > 1) {
       category = parts[0].replace('minecraft:', '').replace(/[_-]/g, ' ')
    }
    category = category.charAt(0).toUpperCase() + category.slice(1)
    
    if (!result[category]) result[category] = []
    result[category].push(adv)
  }
  return result
})
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
      <!-- Sidebar Profile Card -->
      <aside class="profile-card glass-card">
        <h2 class="profile-name">{{ info.last_known_name }}</h2>
        <span class="uuid">{{ info.uuid }}</span>
        
        <div class="canvas-wrapper">
          <canvas ref="viewerCanvas"></canvas>
        </div>
        
        <div class="basic-info">
          <div class="info-row">
            <span class="label">First Join</span>
            <span class="val">{{ formatDate(info.first_played) }}</span>
          </div>
          <div class="info-row">
            <span class="label">Last Activity</span>
            <span class="val">{{ formatDate(info.last_seen) }}</span>
          </div>
          <div class="info-row">
            <span class="label">Playtime</span>
            <span class="val">{{ formatPlaytime(info.ticks_lived) }}</span>
          </div>
          <div class="info-row">
            <span class="label">Health</span>
            <span class="val">{{ (info.health || 0).toFixed(1) }} / 20</span>
          </div>
          <div class="info-row">
            <span class="label">XP Level</span>
            <span class="val badge lvl-badge">{{ info.xp_level }}</span>
          </div>
          <div class="info-row" v-if="linkedAccount">
            <span class="label">Cross-linked</span>
            <span class="val text-success">✔️ Active</span>
          </div>
        </div>
      </aside>

      <!-- Main Content Details -->
      <div class="details-section">
        
        <!-- Ores Pie Chart -->
        <section class="panel glass-card" v-if="oreStats && oreStats.length > 0">
          <h3>Ore Mined Statistics</h3>
          <div class="chart-container">
            <canvas ref="pieChartCanvas"></canvas>
          </div>
          <div class="ore-list">
            <div class="ore-item" v-for="ore in oreStats" :key="ore.name">
              <span class="ore-name">{{ ore.name }}</span>
              <div class="ore-stats">
                <span class="stat" title="Mined">⛏️ {{ formatNumber(ore.mined) }}</span>
                <span class="stat" title="Used/Crafted">🔨 {{ formatNumber(ore.used) }}</span>
              </div>
            </div>
          </div>
        </section>

        <!-- McMMO Stats -->
        <section class="panel glass-card" v-if="mcmmo">
          <h3>McMMO Progress <span class="total-badge">Lvl {{ mcmmo.total }}</span></h3>
          <div class="skill-grid">
            <div class="skill-item" v-for="(level, skill) in filteredMcmmo" :key="skill">
              <span class="skill-name">{{ skill }}</span>
              <span class="skill-level">{{ level }}</span>
              <div class="progress-bar">
                <div class="progress-fill" :style="{ width: `${Math.min((level as number) / 1000 * 100, 100)}%` }"></div>
              </div>
            </div>
          </div>
        </section>

        <!-- Advancements Detail -->
        <section class="panel glass-card" v-if="advancements && advancements.length > 0">
          <div class="panel-header" style="margin-bottom: 1rem; border-bottom: none;">
            <h3>Advancements <small class="text-muted">({{ completedAdvancements }}/{{ totalAdvancements }})</small></h3>
            <label class="toggle-container">
              <span class="toggle-label text-muted" style="margin-right: 8px; font-size: 0.85rem;">Show Locked</span>
              <input type="checkbox" v-model="showLockedAdvancements">
            </label>
          </div>
          
          <div class="adv-category" v-for="(items, category) in categorizedAdvancements" :key="category">
            <h4 class="category-name">{{ category }}</h4>
            <div class="advancements-grid">
              <div 
                :class="['adv-item', { locked: !adv.done }]" 
                v-for="adv in items" 
                :key="adv.key"
              >
                <div class="adv-icon">{{ adv.done ? '⭐' : '🔒' }}</div>
                <div class="adv-text-group">
                  <span class="adv-name" :title="adv.key">{{ adv.key.split('/').pop()?.replace(/_/g, ' ') }}</span>
                  <span class="adv-progress" v-if="!adv.done && adv.criteria">
                    {{ Object.keys(adv.criteria).length }} requirements
                  </span>
                </div>
              </div>
            </div>
          </div>
        </section>

        <!-- Raw Stats Viewer -->
        <section class="panel glass-card" v-if="stats && Object.keys(stats).length > 0">
          <div class="panel-header">
            <h3>Extended Statistics</h3>
            <button class="btn-secondary toggle-btn" @click="showExactValues = !showExactValues">
              {{ showExactValues ? 'Show K/M' : 'Show Exact' }}
            </button>
          </div>
          
          <div class="tabs">
            <button 
              v-for="(_, category) in stats" 
              :key="category"
              :class="['tab-btn', { active: selectedCategory === category }]"
              @click="selectedCategory = category as string"
            >
              {{ formatStatName(category as string) }}
            </button>
          </div>
          
          <div class="stat-grid" v-if="selectedCategory && stats[selectedCategory]">
            <div class="stat-item" v-for="(value, name) in stats[selectedCategory]" :key="name">
              <span class="stat-name" :title="name as string">{{ formatStatName(name as string) }}</span>
              <span class="stat-value">{{ formatNumber(value as number) }}</span>
            </div>
          </div>
        </section>
        
      </div>
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
    width: 320px;
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

.profile-name {
  color: var(--primary);
  font-size: 1.8rem;
  margin-bottom: 0.25rem;
}

.uuid {
  font-family: var(--mono);
  color: var(--text-muted);
  font-size: 0.85rem;
  margin-bottom: 1.5rem;
  background: rgba(0,0,0,0.2);
  padding: 4px 8px;
  border-radius: 4px;
}

.canvas-wrapper {
  margin-bottom: 2rem;
  background: rgba(0, 0, 0, 0.4);
  border-radius: var(--radius-md);
  padding: 10px;
  box-shadow: inset 0 2px 10px rgba(0,0,0,0.5);
  cursor: grab;
}

.canvas-wrapper:active {
  cursor: grabbing;
}

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
  font-size: 0.95rem;
}

.info-row .val {
  font-weight: 600;
  color: var(--text-main);
  text-align: right;
}

.lvl-badge {
  background: rgba(0, 230, 118, 0.2);
  color: var(--primary);
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 0.85rem;
}

.text-success {
  color: var(--primary);
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
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
  padding-bottom: 0.75rem;
  border-bottom: 1px solid var(--glass-border);
  font-size: 1.4rem;
}

.panel-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
  padding-bottom: 0.75rem;
  border-bottom: 1px solid var(--glass-border);
}

.panel-header h3 {
  margin: 0;
  padding: 0;
  border: none;
}

.chart-container {
  height: 250px;
  position: relative;
  margin-bottom: 1.5rem;
}

.ore-list {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.ore-item {
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid var(--glass-border);
  padding: 8px 12px;
  border-radius: var(--radius-sm);
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.ore-name {
  font-weight: 600;
  color: var(--primary);
  font-size: 0.9rem;
}

.ore-stats {
  display: flex;
  gap: 12px;
  font-size: 0.8rem;
  color: var(--text-muted);
}

.skill-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
  gap: 1rem;
}

.skill-item {
  background: rgba(0,0,0,0.2);
  padding: 12px;
  border-radius: var(--radius-sm);
  border: 1px solid rgba(255,255,255,0.05);
}

.skill-name {
  color: var(--text-muted);
  font-size: 0.9rem;
  display: block;
  margin-bottom: 4px;
}

.skill-level {
  font-weight: 700;
  color: var(--secondary);
  font-size: 1.2rem;
  display: block;
  margin-bottom: 8px;
}

.progress-bar {
  height: 4px;
  background: rgba(255,255,255,0.1);
  border-radius: 2px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: var(--secondary);
}

.total-badge {
  background: rgba(99, 102, 241, 0.2);
  color: var(--secondary);
  padding: 4px 12px;
  border-radius: 16px;
  font-size: 1rem;
}

.advancements-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.adv-item {
  display: flex;
  align-items: center;
  gap: 6px;
  background: rgba(255, 171, 0, 0.1);
  border: 1px solid rgba(255, 171, 0, 0.3);
  padding: 6px 12px;
  border-radius: 20px;
}

.adv-name {
  font-size: 0.85rem;
  color: #ffca28;
  text-transform: capitalize;
}

.adv-more {
  display: flex;
  align-items: center;
  padding: 6px 12px;
  font-size: 0.85rem;
  color: var(--text-muted);
  background: rgba(255,255,255,0.05);
  border-radius: 20px;
}

.adv-category {
  margin-bottom: 1.5rem;
}

.category-name {
  font-size: 1.1rem;
  color: var(--text-main);
  margin-bottom: 0.75rem;
  border-bottom: 1px solid rgba(255,255,255,0.05);
  padding-bottom: 4px;
}

.adv-item.locked {
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.adv-item.locked .adv-name {
  color: var(--text-muted);
}

.adv-text-group {
  display: flex;
  flex-direction: column;
}

.adv-progress {
  font-size: 0.7rem;
  color: var(--text-muted);
}


.toggle-btn {
  padding: 6px 14px;
  font-size: 0.85rem;
}

.tabs {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 1.5rem;
}

.tab-btn {
  background: transparent;
  border: 1px solid var(--glass-border);
  color: var(--text-muted);
  padding: 6px 14px;
  border-radius: 20px;
  cursor: pointer;
  text-transform: capitalize;
  font-size: 0.9rem;
  transition: all 0.2s;
}

.tab-btn:hover {
  background: rgba(255,255,255,0.05);
  color: #fff;
}

.tab-btn.active {
  background: rgba(0, 230, 118, 0.1);
  border-color: var(--primary);
  color: var(--primary);
}

.stat-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
  gap: 12px;
  max-height: 400px;
  overflow-y: auto;
  padding-right: 10px;
}

.stat-item {
  background: rgba(0,0,0,0.15);
  border: 1px solid rgba(255,255,255,0.02);
  border-radius: var(--radius-sm);
  padding: 10px 14px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.stat-name {
  font-size: 0.85rem;
  color: var(--text-muted);
  text-transform: capitalize;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  margin-right: 8px;
}

.stat-value {
  font-weight: 700;
  color: var(--primary);
  font-size: 1rem;
}

.loading-state, .empty-state {
  text-align: center;
  padding: 5rem 2rem;
}

.spinner {
  width: 40px;
  height: 40px;
  border: 4px solid rgba(0, 230, 118, 0.2);
  border-left-color: var(--primary);
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 1.5rem;
}

</style>
