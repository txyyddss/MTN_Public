<script setup lang="ts">
import { ref, onMounted, onUnmounted, nextTick, computed } from 'vue'
import { useRoute } from 'vue-router'
import Chart from 'chart.js/auto'
import { API_BASE_URL } from '@/config'
import advancementData from '@/assets/advancements.json'

const route = useRoute()
const uuid = route.params.uuid as string
const loading = ref(true)

const info = ref<any>(null)
const stats = ref<any>(null)
const advancements = ref<any>(null)
const mcmmo = ref<any>(null)
const linkedAccount = ref<any>(null)
const oreStats = ref<any>([])
const ranks = ref<Record<string, number>>({})

const selectedCategory = ref<string>('minecraft:custom')
const statSearch = ref('')

const filteredStats = computed(() => {
  if (!selectedCategory.value || !stats.value || !stats.value[selectedCategory.value]) return {}
  const pool = stats.value[selectedCategory.value]
  
  let entries = Object.entries(pool)
  
  if (statSearch.value) {
    const query = statSearch.value.toLowerCase()
    entries = entries.filter(([name]) => name.toLowerCase().includes(query))
  }
  
  // Sort by value descending
  entries.sort((a, b) => (b[1] as number) - (a[1] as number))
  
  const result: Record<string, number> = {}
  for (const [name, val] of entries) {
    result[name] = val as number
  }
  return result
})

const filteredMcmmo = computed(() => {
  if (!mcmmo.value) return {}
  let entries = Object.entries(mcmmo.value).filter(([key]) => !['user_id', 'user', 'uuid', 'total'].includes(key))
  
  // Sort by level descending
  entries.sort((a, b) => (b[1] as number) - (a[1] as number))
  
  const result: Record<string, any> = {}
  for (const [key, val] of entries) {
    result[key.charAt(0).toUpperCase() + key.slice(1)] = val
  }
  return result
})

const formatNumber = (num: number) => {
  if (num >= 1000000) return (num / 1000000).toFixed(1) + 'M'
  if (num >= 1000) return (num / 1000).toFixed(1) + 'K'
  return num.toLocaleString()
}

const formatStatName = (name: string) => {
  return name.replace('minecraft:', '').replace(/_/g, ' ')
}

const pieChartCanvas = ref<HTMLCanvasElement | null>(null)
let chartInstance: Chart | null = null

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
          data: minedData,
          backgroundColor: [
            '#3B82F6', '#6366f1', '#8b5cf6', '#ec4899', '#f43f5e', 
            '#f59e0b', '#10b981', '#14b8a6', '#64748b', '#475569'
          ],
          borderWidth: 0
        }
      ]
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: {
          position: 'right',
          labels: { color: '#94a3b8', font: { family: 'Noto Sans SC' } }
        }
      }
    }
  })
}

const fetchDetail = async () => {
  try {
    const res = await fetch(`${API_BASE_URL}/api/players/${uuid}`)
    const json = await res.json()
    info.value = json.info
    stats.value = json.stats?.stats || null
    advancements.value = json.advancements?.advancements || []
    mcmmo.value = json.mcmmo
    linkedAccount.value = json.linked_account
    oreStats.value = json.ore_stats
    ranks.value = json.ranks || {}
    
    if (stats.value && Object.keys(stats.value).length > 0) {
      const keys = Object.keys(stats.value)
      if (keys.includes('minecraft:custom')) selectedCategory.value = 'minecraft:custom'
      else selectedCategory.value = keys[0]
    }
    
    nextTick(() => {
      initPieChart()
    })
  } catch (e) {
    console.error('Failed to load player detail', e)
  } finally {
    loading.value = false
  }
}

const getAvatarUrl = (p: any) => {
  if (!p?.last_known_name) return 'https://mineskin.eu/helm/Steve'
  let cleanName = p.last_known_name
  if (p.type === 'Bedrock' || cleanName.startsWith('.')) {
      cleanName = cleanName.replace(/^\./, '').replace(/^BE_/, '')
  }
  return `https://mineskin.eu/helm/${cleanName}`
}

onMounted(() => {
  fetchDetail()
})

onUnmounted(() => {
  if (chartInstance) chartInstance.destroy()
})

const formatDate = (ms: number) => {
  if (!ms) return 'N/A'
  return new Date(ms).toLocaleDateString()
}

const formatPlaytime = (ticks: number) => {
  if (!ticks) return '0h'
  const hours = (ticks / 20 / 3600).toFixed(1)
  return `${hours}h`
}

const totalAdvancements = computed(() => advancements.value?.length || 0)
const completedAdvancements = computed(() => advancements.value?.filter((a: any) => a.done).length || 0)

const categorizedAdvancements = computed(() => {
  if (!advancements.value) return {}
  const result: Record<string, any[]> = {}
  for (const adv of advancements.value) {
    if (!adv.done) continue
    
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

const getAdvancementMetadata = (key: string) => {
  return (advancementData as any)[key] || { name: key.split('/').pop(), icon: key.split('/').pop(), description: '', type: 'task' }
}

const getAdvIconPath = (advKey: string) => {
  const meta = getAdvancementMetadata(advKey)
  const category = advKey.split(':')[1]?.split('/')[0] || 'minecraft'
  const iconName = meta.icon
  
  // Try to use a base path that works in both dev and prod
  // If icons are moved to public, use /mc_icons/...
  // For now, keep as src assets but use a relative import-friendly style if possible.
  // Actually, standardizing on /assets/mc_icons/ might be better if they are in public.
  // But they are in src/assets. 
  return `/mc_icons/advancements/${category}/${iconName}.png`
}
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
        <div class="avatar-header">
            <img :src="getAvatarUrl(info)" :alt="info.last_known_name" class="detail-avatar" />
            <h2 class="profile-name minecraft-font">{{ info.last_known_name }}</h2>
            <span :class="['type-tag', info.type?.toLowerCase()]">{{ info.type === 'Bedrock' ? 'Bedrock' : 'Java' }}</span>
        </div>
        
        <div class="basic-info">
          <div class="info-row">
            <span class="label">First Join</span>
            <span class="val">{{ formatDate(info.first_played) }}</span>
          </div>
          <div class="info-row">
            <span class="label">Last Seen</span>
            <span class="val">{{ formatDate(info.last_seen) }}</span>
          </div>
          <div class="info-row">
            <span class="label">Playtime</span>
            <span class="val">{{ formatPlaytime(info.ticks_lived) }}</span>
          </div>
          <div class="info-row">
            <span class="label">XP Level</span>
            <span class="val badge lvl-badge">{{ info.xp_level }}</span>
          </div>
          <div class="info-row linked-row" v-if="linkedAccount">
            <span class="label">Linked to</span>
            <span class="val account-link">
                {{ linkedAccount.bedrock_username || linkedAccount.java_username }}
            </span>
          </div>
        </div>
      </aside>

      <!-- Main Content Details -->
      <div class="details-section">
        
        <!-- Ores Pie Chart -->
        <section class="panel glass-card" v-if="oreStats && oreStats.length > 0">
          <h3><img src="/icons/all_blocks.ico" class="header-icon" /> Ore Mined Statistics</h3>
          <div class="ore-content">
              <div class="chart-container">
                <canvas ref="pieChartCanvas"></canvas>
              </div>
              <div class="ore-list">
                <div class="ore-item" v-for="ore in oreStats.slice(0, 10)" :key="ore.name">
                  <span class="ore-name">{{ ore.name }}</span>
                  <span class="ore-val">{{ formatNumber(ore.mined) }}</span>
                </div>
              </div>
          </div>
        </section>

        <!-- McMMO Skills -->
        <section class="panel glass-card" v-if="mcmmo">
          <div class="panel-header-simple">
            <h3><img src="/icons/monsters_hunted.ico" class="header-icon" /> McMMO Skills</h3>
            <div class="rank-badge" v-if="ranks.skills">Rank #{{ ranks.skills }}</div>
            <div class="total-badge">Total {{ mcmmo.total }}</div>
          </div>
          <div class="skill-grid">
            <div class="skill-item" v-for="(level, skill) in filteredMcmmo" :key="skill">
              <span class="skill-name">{{ skill }}</span>
              <span class="skill-level">{{ level }}</span>
            </div>
          </div>
        </section>

        <section class="panel glass-card" v-if="advancements && advancements.length > 0">
          <h3><img src="/icons/all_advancements.ico" class="header-icon" /> Advancements <small class="text-muted">({{ completedAdvancements }}/{{ totalAdvancements }})</small></h3>
          
          <div class="adv-category" v-for="(items, category) in categorizedAdvancements" :key="category">
            <h4 class="category-name">{{ category }}</h4>
            <div class="advancements-grid">
              <div class="adv-card" v-for="adv in items" :key="adv.key">
                <div class="adv-icon-wrap" :class="getAdvancementMetadata(adv.key).type">
                  <img :src="getAdvIconPath(adv.key)" :alt="getAdvancementMetadata(adv.key).name" class="adv-icon" @error="(e: any) => e.target.style.display='none'" />
                </div>
                <div class="adv-info">
                  <span class="adv-name">{{ getAdvancementMetadata(adv.key).name }}</span>
                  <p class="adv-desc" v-if="getAdvancementMetadata(adv.key).description">{{ getAdvancementMetadata(adv.key).description }}</p>
                </div>
              </div>
            </div>
          </div>
        </section>

        <!-- Extended Statistics -->
        <section class="panel glass-card" v-if="stats && Object.keys(stats).length > 0">
          <div class="panel-header-simple">
            <h3><img src="/src/assets/icons/all_blocks.ico" class="header-icon" /> Extended Statistics</h3>
            <div class="rank-group">
                <div class="rank-badge mini" v-if="ranks.playtime">Playtime #{{ ranks.playtime }}</div>
                <div class="rank-badge mini" v-if="ranks.mining">Mining #{{ ranks.mining }}</div>
            </div>
          </div>
          
          <div class="tabs-header">
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
            <input v-model="statSearch" placeholder="Filter stats..." class="stat-search-box" />
          </div>
          
          <div class="stat-grid" v-if="selectedCategory && Object.keys(filteredStats).length > 0">
            <div class="stat-box" v-for="(value, name) in filteredStats" :key="name">
              <div class="stat-main">
                <span class="stat-name">{{ formatStatName(name as string) }}</span>
                <span class="stat-value">{{ formatNumber(value as number) }}</span>
              </div>
              <div class="stat-rank" v-if="ranks['stat:' + selectedCategory + ':' + name]">
                  #{{ ranks['stat:' + selectedCategory + ':' + name] }}
              </div>
            </div>
          </div>
          <div v-else class="empty-mini">No stats matching search</div>
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
    width: 300px;
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

.detail-avatar {
    width: 96px;
    height: 96px;
    border-radius: 12px;
    background: #111;
    image-rendering: pixelated;
    margin-bottom: 1rem;
    box-shadow: 0 8px 16px rgba(0,0,0,0.5);
}

.profile-name {
  color: #fff;
  font-size: 1.8rem;
  margin-bottom: 0.5rem;
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
  grid-template-columns: repeat(auto-fill, minmax(130px, 1fr));
  gap: 1rem;
}

.skill-item {
  background: rgba(255,255,255,0.03);
  padding: 12px;
  border-radius: 8px;
  border-left: 3px solid var(--primary);
}

.skill-name {
  color: var(--text-muted);
  font-size: 0.8rem;
  display: block;
}

.skill-level {
  font-weight: 700;
  color: #fff;
  font-size: 1.2rem;
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

.advancements-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 1rem;
}

.adv-card {
  display: flex;
  gap: 12px;
  background: rgba(255, 255, 255, 0.03);
  padding: 12px;
  border-radius: 12px;
  border: 1px solid var(--glass-border);
  transition: all 0.3s ease;
}
.adv-card:hover { border-color: var(--primary); background: rgba(255, 255, 255, 0.05); }

.adv-icon-wrap {
  width: 48px;
  height: 48px;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(0,0,0,0.3);
  border-radius: 8px;
  padding: 6px;
  position: relative;
}
.adv-icon-wrap.goal { border: 2px solid #fcd34d; box-shadow: 0 0 10px rgba(252, 211, 77, 0.3); }
.adv-icon-wrap.challenge { border: 2px solid #f43f5e; box-shadow: 0 0 10px rgba(244, 63, 94, 0.3); }

.adv-icon { width: 32px; height: 32px; image-rendering: pixelated; }

.adv-info { display: flex; flex-direction: column; justify-content: center; min-width: 0; }
.adv-name { font-weight: 700; color: #fff; font-size: 0.95rem; }
.adv-desc { font-size: 0.75rem; color: var(--text-muted); white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }

.stat-box {
  background: rgba(255, 255, 255, 0.02);
  border-radius: 8px;
  padding: 10px 14px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-left: 2px solid transparent;
  transition: all 0.2s;
}
.stat-box:hover { border-left-color: var(--primary); background: rgba(255, 255, 255, 0.05); }

.stat-main { display: flex; flex-direction: column; }
.stat-name { font-size: 0.75rem; color: var(--text-muted); text-transform: uppercase; letter-spacing: 0.5px; }
.stat-value { font-weight: 800; color: #fff; font-size: 1.1rem; font-family: var(--heading); }

.stat-rank { font-size: 0.8rem; font-weight: 800; color: #fcd34d; background: rgba(245, 158, 11, 0.1); padding: 2px 8px; border-radius: 4px; }

.rank-badge.mini { font-size: 0.7rem; padding: 2px 8px; }

.adv-category { margin-bottom: 1.5rem; }
.category-name {
  font-size: 1rem;
  color: var(--text-muted);
  margin-bottom: 0.75rem;
  text-transform: uppercase;
  letter-spacing: 1px;
}

.advancements-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.adv-item {
  background: rgba(59, 130, 246, 0.1);
  border: 1px solid rgba(59, 130, 246, 0.2);
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.85rem;
  color: #93c5fd;
}

.tabs-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 1rem;
    margin-bottom: 1.5rem;
    flex-wrap: wrap;
}

.stat-search-box {
    background: rgba(0,0,0,0.2);
    border: 1px solid var(--glass-border);
    padding: 6px 12px;
    border-radius: 20px;
    font-size: 0.85rem;
    color: #fff;
    min-width: 150px;
}

.tabs {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
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
  grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
  gap: 12px;
  max-height: 450px;
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
