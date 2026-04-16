<script setup lang="ts">
import { ref, onMounted, onUnmounted, nextTick, computed } from 'vue'
import { useRoute } from 'vue-router'
import { watch } from 'vue'
import Chart from 'chart.js/auto'
import { API_BASE_URL } from '@/config'
import advancementData from '@/assets/advancements.json'
import StatBox from '@/components/StatBox.vue'
import SkillItem from '@/components/SkillItem.vue'
import SkinViewer from '@/components/SkinViewer.vue'

const route = useRoute()
const uuid = computed(() => route.params.uuid as string)
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
  console.log('Initializing pie chart with oreStats:', oreStats.value)
  if (!pieChartCanvas.value || oreStats.value.length === 0) {
    console.warn('Cannot init pie chart: canvas missing or no oreStats')
    return
  }
  
  if (chartInstance) chartInstance.destroy()

  const labels = oreStats.value.map((o: any) => o.name)
  const minedData = oreStats.value.map((o: any) => o.mined)
  console.log('Chart labels:', labels, 'Chart data:', minedData)

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
  loading.value = true
  try {
    const res = await fetch(`${API_BASE_URL}/api/players/${uuid.value}`)
    const json = await res.json()
    info.value = json.info
    stats.value = json.stats?.stats || null
    advancements.value = json.advancements?.advancements || []
    mcmmo.value = json.mcmmo
    linkedAccount.value = json.linked_account
    oreStats.value = json.ore_stats
    ranks.value = json.ranks || {}
    
    console.log('Player detail fetched:', {
      hasStats: !!stats.value,
      oreStatsCount: oreStats.value?.length,
      oreStats: oreStats.value
    })
    
    if (stats.value && Object.keys(stats.value).length > 0) {
      const keys = Object.keys(stats.value)
      if (keys.includes('minecraft:custom')) selectedCategory.value = 'minecraft:custom'
      else selectedCategory.value = keys[0]
    }
  } catch (e) {
    console.error('Failed to load player detail', e)
  } finally {
    loading.value = false
    nextTick(() => {
      initPieChart()
    })
  }
}

const getSkinUrl = (p: any) => {
  if (!p?.last_known_name) return 'https://mineskin.eu/skin/Steve'
  let cleanName = p.last_known_name
  if (p.type === 'Bedrock' || cleanName.startsWith('.')) {
      cleanName = cleanName.replace(/^\./, '').replace(/^BE_/, '')
  }
  return `https://mineskin.eu/skin/${cleanName}`
}

onMounted(() => {
  fetchDetail()
})

watch(uuid, (newUuid) => {
  if (newUuid) {
    fetchDetail()
  }
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
  let category = advKey.split(':')[1]?.split('/')[0] || 'minecraft'
  
  if (category === 'story') category = 'minecraft'
  
  const iconName = meta.icon
  return `/mc_icons/advancements/${category}/${iconName}.png`
}

const customStatIcons: Record<string, string> = {
  'play_one_minute': '/mc_icons/blocks/special/observer^32.png',
  'jump': '/mc_icons/blocks/misc/scaffolding^32.png',
  'damage_dealt': '/mc_icons/items/iron_sword.png',
  'damage_taken': '/mc_icons/items/iron_sword.png', // Fallback to sword for combat
  'deaths': '/mc_icons/items/skull_and_beacon^32.png',
  'mob_kills': '/mc_icons/items/iron_sword.png',
  'player_kills': '/mc_icons/items/iron_sword.png',
  'walk_one_cm': '/mc_icons/blocks/dirt/dirt_path^32.png',
  'sprint_one_cm': '/mc_icons/blocks/dirt/dirt_path^32.png',
  'fly_one_cm': '/mc_icons/blocks/misc/cobweb^32.png',
  'climb_one_cm': '/mc_icons/blocks/misc/ladder^32.png',
  'fall_one_cm': '/mc_icons/blocks/misc/cobweb^32.png',
  'minecart_one_cm': '/mc_icons/items/conduit^32.png', // Generic tech
  'boat_one_cm': '/mc_icons/items/water_bucket.png',
  'pig_one_cm': '/mc_icons/blocks/misc/armor_stand^32.png',
  'horse_one_cm': '/mc_icons/blocks/misc/armor_stand^32.png',
  'strider_one_cm': '/mc_icons/blocks/misc/armor_stand^32.png',
  'aviate_one_cm': '/mc_icons/blocks/misc/cobweb^32.png',
  'swim_one_cm': '/mc_icons/items/water_bucket.png',
  'walk_on_water_one_cm': '/mc_icons/items/water_bucket.png',
  'walk_under_water_one_cm': '/mc_icons/items/water_bucket.png',
  'time_since_death': '/mc_icons/blocks/special/observer^32.png',
  'time_since_rest': '/mc_icons/blocks/beds/white_bed^32.png',
  'sneak_time': '/mc_icons/blocks/dirt/dirt_path^32.png',
  'total_world_time': '/mc_icons/blocks/special/observer^32.png',
}

const getStatIconPath = (category: string, name: string) => {
  const id = name.replace('minecraft:', '')
  
  if (category === 'minecraft:custom') {
    return customStatIcons[id] || '/mc_icons/items/paper^32.png'
  }

  // Common blocks subfolder mapping
  const blockFolders: Record<string, string> = {
    'iron_ore': 'materials', 'gold_ore': 'materials', 'diamond_ore': 'materials', 'emerald_ore': 'materials',
    'lapis_ore': 'materials', 'redstone_ore': 'materials', 'coal_ore': 'materials', 'copper_ore': 'materials',
    'deepslate_iron_ore': 'materials', 'deepslate_gold_ore': 'materials', 'deepslate_diamond_ore': 'materials',
    'deepslate_emerald_ore': 'materials', 'deepslate_lapis_ore': 'materials', 'deepslate_redstone_ore': 'materials',
    'deepslate_coal_ore': 'materials', 'deepslate_copper_ore': 'materials',
    'iron_block': 'materials', 'gold_block': 'materials', 'diamond_block': 'materials', 'emerald_block': 'materials',
    'dirt': 'dirt', 'coarse_dirt': 'dirt', 'rooted_dirt': 'dirt', 'grass_block': 'dirt', 'podzol': 'dirt', 'mycelium': 'dirt',
    'stone': 'stone', 'cobblestone': 'cobblestone', 'andesite': 'andesite', 'diorite': 'diorite', 'granite': 'granite',
    'oak_log': 'wood', 'spruce_log': 'wood', 'birch_log': 'wood', 'jungle_log': 'wood', 'acacia_log': 'wood', 'dark_oak_log': 'wood',
    'oak_planks': 'wood', 'spruce_planks': 'wood', 'birch_planks': 'wood', 'jungle_planks': 'wood', 'acacia_planks': 'wood', 'dark_oak_planks': 'wood',
  }

  if (category === 'minecraft:mined' || category === 'minecraft:broken') {
    const folder = blockFolders[id]
    if (folder) return `/mc_icons/blocks/${folder}/${id}^32.png`
    // Default to a guess in blocks root if not in mapping, but since we saw blocks has no root files, 
    // we should probably try blocks/misc if we had one, but we'll try items first as a fallback.
    return `/mc_icons/items/${id}^32.png`
  }

  if (['minecraft:crafted', 'minecraft:used', 'minecraft:picked_up', 'minecraft:dropped'].includes(category)) {
      // Try items first with suffix
      return `/mc_icons/items/${id}^32.png`
  }

  if (category === 'minecraft:killed' || category === 'minecraft:killed_by') {
    return `/mc_icons/items/${id}_spawn_egg^32.png`
  }

  return null
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
            <div class="skin-wrapper">
              <SkinViewer :skin-url="getSkinUrl(info)" />
            </div>
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
          <div class="info-row linked-row" v-if="linkedAccount && info.type === 'Bedrock'">
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
              <div class="ore-list mobile-hide">
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
            <SkillItem 
              v-for="(level, skill) in filteredMcmmo" 
              :key="skill"
              :name="String(skill)"
              :level="Number(level)"
              :rank="ranks['mcmmo:' + skill.toLowerCase()]"
            />
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
                </div>
              </div>
            </div>
          </div>
        </section>

        <!-- Extended Statistics -->
        <section class="panel glass-card" v-if="stats && Object.keys(stats).length > 0">
          <div class="panel-header-simple">
            <h3><img src="/icons/all_blocks.ico" class="header-icon" /> Extended Statistics</h3>
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
            <StatBox 
              v-for="(value, name) in filteredStats" 
              :key="name"
              :name="String(name)"
              :value="value"
              :rank="ranks['stat:' + selectedCategory + ':' + name]"
              :icon="getStatIconPath(selectedCategory, String(name))"
              :formatValue="formatNumber"
            />
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
  grid-template-columns: repeat(auto-fill, minmax(110px, 1fr));
  gap: 1rem;
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
  grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
  gap: 0.5rem;
}

.adv-card {
  display: flex;
  align-items: center;
  gap: 8px;
  background: rgba(255, 255, 255, 0.03);
  padding: 6px 10px;
  border-radius: 8px;
  border: 1px solid var(--glass-border);
  transition: all 0.3s ease;
}
.adv-card:hover { border-color: var(--primary); background: rgba(255, 255, 255, 0.05); }

.adv-icon-wrap {
  width: 32px;
  height: 32px;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(0,0,0,0.3);
  border-radius: 6px;
  padding: 4px;
}
.adv-icon-wrap.goal { border: 1px solid #fcd34d; }
.adv-icon-wrap.challenge { border: 1px solid #f43f5e; }

.adv-icon { width: 20px; height: 20px; image-rendering: pixelated; }

.adv-info { display: flex; flex-direction: column; justify-content: center; min-width: 0; }
.adv-name { font-weight: 700; color: #fff; font-size: 0.8rem; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }

.stat-value { font-weight: 800; color: #fff; font-size: 1.1rem; font-family: var(--heading); }

.rank-badge.mini { font-size: 0.7rem; padding: 2px 8px; }

.adv-category { margin-bottom: 1.5rem; }
.category-name {
  font-size: 1rem;
  color: var(--text-muted);
  margin-bottom: 0.75rem;
  text-transform: uppercase;
  letter-spacing: 1px;
}

/* Duplicate class removal and consolidation */

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
