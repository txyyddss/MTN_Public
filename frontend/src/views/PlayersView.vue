<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

interface PlayerInfo {
  uuid: string
  last_known_name: string
  last_seen: number
}

const router = useRouter()
const players = ref<PlayerInfo[]>([])
const count = ref(0)
const activeDays = ref(0)
const searchQuery = ref('')
const loading = ref(true)

const fetchPlayers = async (search = '') => {
  loading.value = true
  try {
    const url = search ? `/api/players?search=${encodeURIComponent(search)}` : '/api/players'
    const res = await fetch(url)
    const json = await res.json()
    players.value = json.players || []
    count.value = json.count || 0
    if (json.active_days !== undefined) {
      activeDays.value = json.active_days
    }
  } catch (e) {
    console.error('Failed to fetch players:', e)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  fetchPlayers(searchQuery.value)
}

const handleRandom = async () => {
  try {
    const res = await fetch('/api/players/random')
    if (!res.ok) return
    const json = await res.json()
    if (json.uuid) {
      router.push(`/player/${json.uuid}`)
    }
  } catch (e) {
    console.error('Failed to fetch random player', e)
  }
}

const getAvatarUrl = (p: PlayerInfo) => {
  if (!p.last_known_name) return 'https://mineskin.eu/helm/Steve'
  const cleanName = p.last_known_name.startsWith('be_') 
    ? p.last_known_name.substring(3) 
    : p.last_known_name
  return `https://mineskin.eu/helm/${cleanName}`
}

onMounted(() => {
  fetchPlayers()
})
</script>

<template>
  <div class="players-view container animate-entry">
    <div class="header">
      <div class="title-area">
        <h1>Players Directory</h1>
        <p class="subtitle" v-if="!searchQuery && activeDays > 0">
          Showing {{ count }} players active in the last {{ activeDays }} days
        </p>
        <p class="subtitle" v-else-if="searchQuery">
          Found {{ count }} players matching "{{ searchQuery }}"
        </p>
      </div>

      <div class="controls glass-card">
        <input 
          v-model="searchQuery" 
          @keyup.enter="handleSearch"
          placeholder="Search by name or UUID..." 
          class="search-box"
        />
        <button @click="handleSearch" class="btn-primary">Search</button>
        <button @click="handleRandom" class="btn-secondary">Random Player</button>
      </div>
    </div>

    <div v-if="loading" class="loading-state">
      <div class="spinner"></div>
      <p>Loading players...</p>
    </div>

    <div v-else-if="players.length === 0" class="empty-state glass-card">
      <h3>No players found</h3>
      <p>Try searching for a different username or UUID.</p>
      <button @click="() => { searchQuery = ''; fetchPlayers(); }" class="btn-primary mt-4">Reset Search</button>
    </div>

    <div v-else class="grid">
      <RouterLink 
        v-for="(p, index) in players" 
        :key="p.uuid" 
        :to="`/player/${p.uuid}`"
        class="player-card glass-card animate-entry"
        :style="{ animationDelay: `${(index % 20) * 50}ms` }"
      >
        <div class="avatar-container">
          <img :src="getAvatarUrl(p)" :alt="p.last_known_name" class="avatar" loading="lazy" />
        </div>
        <div class="info">
          <h3>{{ p.last_known_name || 'Unknown' }}</h3>
          <span class="uuid">{{ p.uuid.split('-')[0] }}...</span>
        </div>
        <div class="card-arrow">→</div>
      </RouterLink>
    </div>
  </div>
</template>

<style scoped>
.players-view {
  padding-top: 2rem;
  padding-bottom: 4rem;
}

.header {
  display: flex;
  flex-direction: column;
  gap: 2rem;
  margin-bottom: 3rem;
}

@media (min-width: 768px) {
  .header {
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
  }
}

.title-area h1 {
  font-size: 2.5rem;
  margin: 0 0 0.5rem 0;
}

.subtitle {
  color: var(--text-muted);
  margin: 0;
}

.controls {
  display: flex;
  gap: 1rem;
  flex-wrap: wrap;
  padding: 1.5rem;
  align-items: center;
}

.search-box {
  padding: 12px 16px;
  border-radius: var(--radius-sm);
  border: 1px solid var(--glass-border);
  background: rgba(0, 0, 0, 0.2);
  color: white;
  flex: 1;
  min-width: 200px;
  font-family: var(--sans);
  font-size: 1rem;
  transition: border-color 0.3s, box-shadow 0.3s;
}

.search-box:focus {
  outline: none;
  border-color: var(--primary);
  box-shadow: 0 0 0 2px rgba(0, 230, 118, 0.2);
}

.btn-secondary {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 12px 24px;
  background: rgba(255, 255, 255, 0.1);
  color: #fff;
  font-family: var(--heading);
  font-weight: 700;
  border: 1px solid var(--glass-border);
  border-radius: var(--radius-sm);
  cursor: pointer;
  transition: all 0.3s ease;
}

.btn-secondary:hover {
  background: rgba(255, 255, 255, 0.2);
  border-color: rgba(255, 255, 255, 0.4);
}

.grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 1.5rem;
}

.player-card {
  display: flex;
  align-items: center;
  gap: 1.5rem;
  text-decoration: none;
  color: inherit;
  padding: 1.5rem;
  position: relative;
  overflow: hidden;
}

.player-card:hover {
  border-color: var(--primary);
}

.player-card:hover .card-arrow {
  transform: translateX(0);
  opacity: 1;
  color: var(--primary);
}

.avatar-container {
  width: 70px;
  height: 70px;
  border-radius: var(--radius-sm);
  background: rgba(0, 0, 0, 0.3);
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden;
}

.avatar {
  width: 100%;
  height: 100%;
  object-fit: cover;
  image-rendering: pixelated;
  transform: scale(1.1); /* To crop out any edge artifacts from skins */
}

.info {
  flex: 1;
}

.info h3 {
  margin: 0 0 0.25rem 0;
  font-size: 1.2rem;
  font-weight: 700;
  color: #fff;
}

.uuid {
  font-family: var(--mono);
  color: var(--text-muted);
  font-size: 0.85rem;
}

.card-arrow {
  font-size: 1.5rem;
  font-weight: bold;
  color: var(--text-muted);
  transform: translateX(-10px);
  opacity: 0.5;
  transition: all 0.3s ease;
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

@keyframes spin {
  100% { transform: rotate(360deg); }
}

.mt-4 { margin-top: 1.5rem; }
</style>
