<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { watch } from 'vue'
import { API_BASE_URL } from '@/config'
import advancementData from '@/assets/advancements.json'
import StatBox from '@/components/StatBox.vue'
import SkillItem from '@/components/SkillItem.vue'
import SkinViewer from '@/components/SkinViewer.vue'
import PlayerPieChart from '@/components/PlayerPieChart.vue'
import { fetchWithCache } from '@/utils/dataCache'
import { preloadImages, PreloadPriority } from '@/utils/preloader'

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

const CUSTOM_STAT_TRANSLATIONS: Record<string, string> = {
  'play_one_minute': 'Time Played',
  'play_time': 'Time Played',
  'jump': 'Jumps',
  'damage_dealt': 'Damage Dealt',
  'damage_taken': 'Damage Taken',
  'deaths': 'Deaths',
  'mob_kills': 'Mob Kills',
  'player_kills': 'Player Kills',
  'walk_one_cm': 'Distance Walked',
  'sprint_one_cm': 'Distance Sprinted',
  'fly_one_cm': 'Distance Flown',
  'climb_one_cm': 'Distance Climbed',
  'fall_one_cm': 'Distance Fallen',
  'minecart_one_cm': 'Distance by Minecart',
  'boat_one_cm': 'Distance by Boat',
  'pig_one_cm': 'Distance by Pig',
  'horse_one_cm': 'Distance by Horse',
  'strider_one_cm': 'Distance by Strider',
  'aviate_one_cm': 'Distance by Elytra',
  'swim_one_cm': 'Distance Swum',
  'walk_on_water_one_cm': 'Distance on Water',
  'walk_under_water_one_cm': 'Distance Under Water',
  'time_since_death': 'Time Since Last Death',
  'time_since_rest': 'Time Since Last Rest',
  'sneak_time': 'Sneak Time',
  'total_world_time': 'Total World Time',
  'leave_game': 'Games Quit',
  'dropped': 'Items Dropped',
  'interact_with_beacon': 'Beacon Interactions',
  'inspect_hopper': 'Hopper Inspections',
  'interact_with_blast_furnace': 'Blast Furnace Interactions',
  'interact_with_smoker': 'Smoker Interactions',
  'interact_with_camp_fire': 'Campfire Interactions',
  'interact_with_campfire': 'Campfire Interactions',
  'talked_to_villager': 'Villager Talks',
  'traded_with_villager': 'Villager Trades',
  'fish_caught': 'Fish Caught',
  'sleep_in_bed': 'Times Slept',
  'raid_win': 'Raids Won',
  'raid_trigger': 'Raids Triggered',
  'trigger_trapped_chest': 'Trapped Chests Triggered',
  'damage_absorbed': 'Damage Absorbed',
  'interact_with_furnace': 'Furnace Interactions',
  'crouch_one_cm': 'Distance Crouched',
  'interact_with_stonecutter': 'Stonecutter Interactions',
  'damage_resisted': 'Damage Resisted',
  'damage_blocked_by_shield': 'Damage Blocked by Shield',
  'interact_with_crafting_table': 'Crafting Table Interactions',
  'inspect_dropper': 'Dropper Inspections',
  'target_hit': 'Targets Hit',
  'fill_cauldron': 'Cauldrons Filled',
  'interact_with_grindstone': 'Grindstone Interactions',
  'open_shulker_box': 'Shulker Boxes Opened',
  'open_enderchest': 'Ender Chests Opened',
  'damage_dealt_absorbed': 'Damage Dealt (Absorbed)',
  'interact_with_brewingstand': 'Brewing Stand Interactions',
  'inspect_dispenser': 'Dispenser Inspections',
  'interact_with_loom': 'Loom Interactions',
  'play_noteblock': 'Note Blocks Played',
  'interact_with_lectern': 'Lectern Interactions',
  'drop': 'Items Dropped',
  'use_cauldron': 'Cauldrons Used',
  'bell_ring': 'Bells Rung',
  'open_barrel': 'Barrels Opened',
  'interact_with_cartography_table': 'Cartography Table Interactions',
  'open_chest': 'Chests Opened',
  'tune_noteblock': 'Note Blocks Tuned',
  'interact_with_anvil': 'Anvil Interactions',
  'animals_bred': 'Animals Bred',
  'play_record': 'Music Discs Played',
  'interact_with_smithing_table': 'Smithing Table Interactions',
  'pot_flower': 'Flowers Potted',
  'clean_banner': 'Banners Cleaned',
  'enchant_item': 'Items Enchanted',
  'clean_armor': 'Armor Pieces Cleaned',
  'clean_shulker_box': 'Shulker Boxes Cleaned',
  'eat_cake_slice': 'Cake Slices Eaten',
  'interact_with_beehive': 'Beehive Interactions',
  'interact_with_lodestone': 'Lodestone Interactions',
  'interact_with_respawn_anchor': 'Respawn Anchor Interactions'
}

const enlargedAdvancement = ref<string | null>(null)
const enlargedStat = ref<string | null>(null)

const toggleAdvancement = (key: string, e: Event) => {
  e.stopPropagation()
  if (enlargedAdvancement.value === key) {
    enlargedAdvancement.value = null
  } else {
    resetAllEnlarged()
    enlargedAdvancement.value = key
  }
}

const resetAllEnlarged = () => {
  enlargedAdvancement.value = null
  enlargedStat.value = null
}

const toggleStat = (key: string) => {
  if (enlargedStat.value === key) {
    enlargedStat.value = null
  } else {
    resetAllEnlarged()
    enlargedStat.value = key
  }
}

const selectedCategory = ref<string>('minecraft:custom')
const statSearch = ref('')

const fuzzyMatch = (text: string, query: string) => {
  let i = 0, j = 0;
  while (i < text.length && j < query.length) {
    if (text[i] === query[j]) j++;
    i++;
  }
  return j === query.length;
}

const filteredStats = computed(() => {
  if (!selectedCategory.value || !stats.value || !stats.value[selectedCategory.value]) return {}
  const pool = stats.value[selectedCategory.value]
  
  let entries = Object.entries(pool)

  if (selectedCategory.value === 'minecraft:custom') {
    entries = entries.filter(([name]) => {
      const id = name.replace('minecraft:', '')
      return !!CUSTOM_STAT_TRANSLATIONS[id]
    })
  }
  
  if (statSearch.value) {
    const query = statSearch.value.toLowerCase().trim()
    entries = entries.filter(([name]) => {
      // Use formatStatName if it is defined, else fallback to name
      // since formatStatName is initialized after this computed property in the file
      // we can safely call it inside the closure, but just to be safe we check
      const id = name.replace('minecraft:', '')
      const displayName = ((id === 'custom' ? 'Global' : (CUSTOM_STAT_TRANSLATIONS[id] || id.replace(/_/g, ' ')))).toLowerCase()
      const rawName = name.toLowerCase()
      return fuzzyMatch(displayName, query) || fuzzyMatch(rawName, query)
    })
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
  const id = name.replace('minecraft:', '')
  if (id === 'custom') return 'Global'
  if (CUSTOM_STAT_TRANSLATIONS[id]) return CUSTOM_STAT_TRANSLATIONS[id]
  return id.replace(/_/g, ' ')
}

const formatStatValue = (val: number, name: string) => {
    const id = name.replace('minecraft:', '')
    
    // Time stats (ticks)
    if (id.includes('time') || id.includes('one_minute') || id.includes('since')) {
        const hours = val / 20 / 3600
        if (hours > 1) return hours.toFixed(1) + 'h'
        const mins = val / 20 / 60
        return mins.toFixed(0) + 'm'
    }

    // Distance stats (cm)
    if (id.endsWith('_one_cm')) {
        const km = val / 100000;
        if (km > 1) return km.toFixed(2) + 'km'
        const m = val / 100
        return m.toFixed(1) + 'm'
    }

    return formatNumber(val)
}


const fetchDetail = async () => {
  loading.value = true
  try {
    const json = await fetchWithCache(`${API_BASE_URL}/api/players/${uuid.value}`)
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
  window.addEventListener('click', resetAllEnlarged)
  window.addEventListener('scroll', resetAllEnlarged, { capture: true, passive: true })
})

watch(uuid, (newUuid) => {
  if (newUuid) {
    fetchDetail()
  }
})

onUnmounted(() => {
  window.removeEventListener('click', resetAllEnlarged)
  window.removeEventListener('scroll', resetAllEnlarged, { capture: true })
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

import iconMap from '@/assets/icon_map.json'

const getStatIconPath = (category: string, name: string) => {
  const id = name.replace('minecraft:', '')
  
  if (category === 'minecraft:custom') return null

  // Try exact match in icon map
  if ((iconMap as any)[id]) return (iconMap as any)[id]
  
  // Fuzzy matching for mobs
  if (category === 'minecraft:killed' || category === 'minecraft:killed_by') {
      // If it's a mob, try adding _spawn_egg if not found
      const eggId = `${id}_spawn_egg`
      if ((iconMap as any)[eggId]) return (iconMap as any)[eggId]
  }

  // Try removing 'spawn_egg' if it's there
  const cleanId = id.replace('_spawn_egg', '')
  if ((iconMap as any)[cleanId]) return (iconMap as any)[cleanId]

  // Final fallbacks for items/blocks
  if (category === 'minecraft:mined' || category === 'minecraft:broken') {
    return `/mc_icons/blocks/misc/${id}^32.png` 
  }

  if (['minecraft:crafted', 'minecraft:used', 'minecraft:picked_up', 'minecraft:dropped'].includes(category)) {
      return `/mc_icons/items/${id}.png`
  }

  return null
}

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
    const cats = newGroup.categories.filter(cat => stats.value && stats.value[cat])
    if (cats.length > 0) {
        if (!cats.includes(selectedCategory.value)) {
            selectedCategory.value = cats[0]
        }
    }
})

// REALTIME PRELOAD LOGIC - Specific to current player
// Preload skin when info is fetched
watch(info, (newInfo) => {
  if (newInfo) {
    preloadImages([getSkinUrl(newInfo)], PreloadPriority.HIGH)
  }
})

// Preload advancement icons when they are categorized
watch(categorizedAdvancements, (categories) => {
  if (categories) {
    const urls: string[] = []
    Object.values(categories).forEach(items => {
      items.forEach(adv => {
        urls.push(getAdvIconPath(adv.key))
      })
    })
    preloadImages(urls, PreloadPriority.MEDIUM_HIGH)
  }
}, { immediate: true })
</script>

<template>
  <div class="player-detail container animate-entry" @click="resetAllEnlarged">
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
        <PlayerPieChart :oreStats="oreStats" @resetEnlarged="resetAllEnlarged" />

        <!-- McMMO Skills -->
        <section class="panel glass-card" v-if="mcmmo">
          <div class="panel-header-simple" @click.stop="resetAllEnlarged">
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
          <h3 @click.stop="resetAllEnlarged"><img src="/icons/all_advancements.ico" class="header-icon" /> Advancements <small class="text-muted">({{ completedAdvancements }}/{{ totalAdvancements }})</small></h3>
          
          <div class="adv-category" v-for="(items, category) in categorizedAdvancements" :key="category">
            <h4 class="category-name">{{ category }}</h4>
            <div class="advancements-grid">
              <div class="adv-card" v-for="adv in items" :key="adv.key" :class="{ enlarged: enlargedAdvancement === adv.key }" @click="toggleAdvancement(adv.key, $event)">
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
          <div class="panel-header-simple" @click.stop="resetAllEnlarged">
            <h3><img src="/icons/all_blocks.ico" class="header-icon" /> Extended Statistics</h3>
            <div class="rank-group">
                <div class="rank-badge mini" v-if="ranks.playtime">Playtime #{{ ranks.playtime }}</div>
                <div class="rank-badge mini" v-if="ranks.mining">Mining #{{ ranks.mining }}</div>
            </div>
          </div>

          <div class="group-tabs">
            <button 
              v-for="group in statGroups" 
              :key="group.name"
              :class="['group-btn', { active: activeGroup.name === group.name }]"
              @click="activeGroup = group"
            >
              {{ group.name }}
            </button>
          </div>
          
          <div class="tabs-header">
            <div class="tabs" v-if="groupCategories.length > 0">
              <button 
                v-for="category in groupCategories" 
                :key="category"
                :class="['tab-btn', { active: selectedCategory === category }]"
                @click="selectedCategory = category as string"
              >
                {{ formatStatName(category as string) }}
              </button>
            </div>
            <div class="search-wrapper">
              <svg class="search-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <circle cx="11" cy="11" r="8"></circle>
                <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
              </svg>
              <input v-model="statSearch" placeholder="Search stats..." class="stat-search-box" />
            </div>
          </div>
          
          <div :class="selectedCategory === 'minecraft:custom' ? 'custom-stat-list' : 'stat-grid'" v-if="selectedCategory && Object.keys(filteredStats).length > 0">
            <template v-if="selectedCategory === 'minecraft:custom'">
              <div v-for="(value, name) in filteredStats" :key="name" class="custom-stat-item">
                <span class="stat-label">{{ formatStatName(String(name)) }}</span>
                <div class="stat-value-wrap">
                    <span class="stat-rank" v-if="ranks['stat:' + selectedCategory + ':' + name]">#{{ ranks['stat:' + selectedCategory + ':' + name] }}</span>
                    <span class="stat-val">{{ formatStatValue(value, String(name)) }}</span>
                </div>
              </div>
            </template>
            <StatBox 
              v-else
              v-for="(value, name) in filteredStats" 
              :key="name"
              :name="String(name)"
              :value="value"
              :isEnlarged="enlargedStat === name"
              @toggle="toggleStat(String(name))"
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
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
  position: relative;
  z-index: 1;
}
.adv-card:hover { border-color: var(--primary); background: rgba(255, 255, 255, 0.05); }

.adv-card.enlarged {
  transform: scale(1.1);
  z-index: 10;
  background: rgba(30, 41, 59, 0.95);
  border-color: var(--primary);
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.4), 0 8px 10px -6px rgba(0, 0, 0, 0.4);
  backdrop-filter: blur(8px);
}

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
  transition: transform 0.3s ease;
}
.adv-card.enlarged .adv-icon-wrap { transform: scale(1.2); }
.adv-icon-wrap.goal { border: 1px solid #fcd34d; }
.adv-icon-wrap.challenge { border: 1px solid #f43f5e; }

.adv-icon { width: 20px; height: 20px; image-rendering: pixelated; }

.adv-info { display: flex; flex-direction: column; justify-content: center; min-width: 0; }
.adv-name { 
  font-weight: 700; 
  color: #fff; 
  font-size: 0.8rem; 
  white-space: nowrap; 
  overflow: hidden; 
  text-overflow: ellipsis; 
  transition: all 0.3s ease;
}

.adv-card.enlarged .adv-name {
  white-space: normal;
  overflow: visible;
  text-overflow: clip;
}

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
