<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { API_BASE_URL } from '@/config'

const types = ['skills', 'playtime', 'mining', 'killing', 'deaths', 'walking', 'pvp']
const currentType = ref('mining')
const entries = ref<any[]>([])
const onlinePlayers = ref<string[]>([])
const loading = ref(false)

const fetchOnline = async () => {
  try {
    const res = await fetch(`${API_BASE_URL}/api/status`)
    if (res.ok) {
        const json = await res.json()
        onlinePlayers.value = json.online_players || []
    }
  } catch (e) {}
}

const isOnline = (uuid: string) => {
    return onlinePlayers.value.includes(uuid)
}

const fetchLeaderboard = async (type: string) => {
  loading.value = true
  currentType.value = type
  try {
    const res = await fetch(`${API_BASE_URL}/api/leaderboards/${type}`)
    const json = await res.json()
    entries.value = json.entries || []
  } catch(e) {
    console.error('Failed', e)
  }
  loading.value = false
}

onMounted(() => {
    fetchLeaderboard(currentType.value)
    fetchOnline()
})

const formatValue = (val: number, type: string) => {
  if (type === 'playtime') return (val / 20 / 3600).toFixed(1) + ' hrs'
  if (type === 'walking') {
    const km = val / 100000;
    if (km > 1) return km.toFixed(1) + ' km'
    return (val / 100).toFixed(0) + ' m'
  }
  return val.toLocaleString()
}

const getAvatarUrl = (name: string) => {
  if (!name) return 'https://mineskin.eu/helm/Steve'
  const cleanName = name.startsWith('be_') ? name.substring(3) : name
  return `https://mineskin.eu/helm/${cleanName}`
}

const getRankClass = (rank: number) => {
  if (rank === 1) return 'rank-first'
  if (rank === 2) return 'rank-second'
  if (rank === 3) return 'rank-third'
  return ''
}
</script>

<template>
  <div class="leaderboards-view container animate-entry">
    <div class="header glass-card">
      <div class="header-content">
        <h1>Global Leaderboards</h1>
        <p class="text-muted">See who tops the server across various statistical categories.</p>
      </div>
      
      <div class="type-selector">
        <button 
          v-for="t in types" 
          :key="t" 
          @click="fetchLeaderboard(t)"
          :class="['tab-btn', { active: currentType === t }]"
        >
          {{ t }}
        </button>
      </div>
    </div>

    <div v-if="loading" class="loading-state">
      <div class="spinner"></div>
      <p>Fetching rankings...</p>
    </div>
    
    <div v-else-if="entries.length === 0" class="empty-state glass-card">
      <h3>No data available</h3>
      <p>No players have ranked in this category yet.</p>
    </div>
    
    <div v-else class="lb-table-wrapper glass-card animate-entry delay-100">
      <table class="lb-table">
        <thead>
          <tr>
            <th class="rank-col">Rank</th>
            <th>Player</th>
            <th class="val-col">Score</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="entry in entries" :key="entry.uuid" :class="getRankClass(entry.rank)">
            <td class="rank-col">
              <span class="rank-badge">#{{ entry.rank }}</span>
            </td>
            <td>
              <RouterLink :to="`/player/${entry.uuid}`" class="player-link">
                <img :src="getAvatarUrl(entry.name)" alt="avatar" class="avatar-small" loading="lazy" />
                <span :class="['player-name', 'minecraft-font', { 'online-name': isOnline(entry.uuid) }]">{{ entry.name }}</span>
              </RouterLink>
            </td>
            <td class="val-col">
              <span class="value-badge">{{ formatValue(entry.value, currentType) }}</span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped>
.leaderboards-view {
  padding-top: 2rem;
  padding-bottom: 4rem;
  max-width: 1000px;
}

.header {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  margin-bottom: 2rem;
  padding: 2rem;
}

@media (min-width: 768px) {
  .header {
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
  }
}

.header-content h1 {
  font-size: 2.2rem;
  color: var(--primary);
  margin-bottom: 0.5rem;
}

.type-selector {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.tab-btn {
  background: rgba(0,0,0,0.2);
  color: var(--text-muted);
  border: 1px solid var(--glass-border);
  padding: 8px 16px;
  border-radius: 20px;
  cursor: pointer;
  text-transform: capitalize;
  font-size: 0.95rem;
  font-family: var(--heading);
  transition: all 0.3s;
}

.tab-btn:hover {
  background: rgba(255,255,255,0.05);
  color: var(--text-main);
  border-color: rgba(255,255,255,0.2);
}

.tab-btn.active {
  background: var(--primary);
  color: #000;
  border-color: var(--primary);
  font-weight: 700;
}

.lb-table-wrapper {
  padding: 0;
  overflow-x: auto;
}

.lb-table {
  width: 100%;
  border-collapse: collapse;
}

.lb-table th {
  background: rgba(0, 0, 0, 0.4);
  text-align: left;
  padding: 18px 24px;
  color: var(--text-muted);
  font-weight: 600;
  border-bottom: 1px solid var(--glass-border);
  font-size: 0.95rem;
  text-transform: uppercase;
  letter-spacing: 1px;
}

.lb-table td {
  padding: 16px 24px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
  vertical-align: middle;
}

.lb-table tbody tr {
  transition: background 0.3s ease;
}

.lb-table tbody tr:hover {
  background: rgba(255, 255, 255, 0.05);
}

.lb-table tr:last-child td {
  border-bottom: none;
}

.rank-col {
  width: 100px;
}

.rank-badge {
  font-weight: 700;
  font-family: var(--mono);
  font-size: 1.1rem;
  color: var(--text-muted);
}

/* Trophy highlights */
.rank-first .rank-badge {
  color: #ffd700;
  font-size: 1.3rem;
  text-shadow: 0 0 10px rgba(255, 215, 0, 0.5);
}
.rank-second .rank-badge {
  color: #c0c0c0;
  font-size: 1.2rem;
  text-shadow: 0 0 10px rgba(192, 192, 192, 0.5);
}
.rank-third .rank-badge {
  color: #cd7f32;
  font-size: 1.15rem;
  text-shadow: 0 0 10px rgba(205, 127, 50, 0.5);
}

.val-col {
  text-align: right;
}

.value-badge {
  font-family: var(--mono);
  font-weight: 700;
  font-size: 1.05rem;
  color: var(--primary);
  background: rgba(59, 130, 246, 0.1);
  padding: 6px 12px;
  border-radius: 12px;
}

.player-link {
  display: inline-flex;
  align-items: center;
  gap: 14px;
  color: var(--text-main);
  text-decoration: none;
  font-weight: 600;
  font-size: 1.1rem;
  transition: color 0.2s;
}

.player-link:hover {
  color: var(--primary);
}

.player-link .player-name.online-name {
    color: #10b981;
    text-shadow: 0 0 10px rgba(16, 185, 129, 0.4);
}

.avatar-small {
  width: 32px;
  height: 32px;
  border-radius: 4px;
  image-rendering: pixelated;
  background: rgba(0,0,0,0.3);
  box-shadow: 0 2px 5px rgba(0,0,0,0.3);
}

.loading-state, .empty-state {
  text-align: center;
  padding: 5rem 2rem;
}

.spinner {
  width: 40px;
  height: 40px;
  border: 4px solid rgba(59, 130, 246, 0.2);
  border-left-color: var(--primary);
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 1.5rem;
}

@keyframes spin {
  100% { transform: rotate(360deg); }
}
</style>
