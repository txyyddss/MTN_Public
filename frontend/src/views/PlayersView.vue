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

onMounted(() => {
  fetchPlayers()
})
</script>

<template>
  <div class="players-view">
    <div class="header">
      <div class="title-area">
        <h1>Players</h1>
        <p class="subtitle" v-if="!searchQuery && activeDays > 0">
          Showing {{ count }} players active in the last {{ activeDays }} days
        </p>
        <p class="subtitle" v-else-if="searchQuery">
          Found {{ count }} players matching "{{ searchQuery }}"
        </p>
      </div>

      <div class="controls">
        <input 
          v-model="searchQuery" 
          @keyup.enter="handleSearch"
          placeholder="Search by name or UUID..." 
          class="search-box"
        />
        <button @click="handleSearch" class="btn primary">Search</button>
        <button @click="handleRandom" class="btn secondary">Random Player</button>
      </div>
    </div>

    <div v-if="loading" class="loading">Loading players...</div>

    <div v-else-if="players.length === 0" class="empty-state">
      No players found.
    </div>

    <div v-else class="grid">
      <RouterLink 
        v-for="p in players" 
        :key="p.uuid" 
        :to="`/player/${p.uuid}`"
        class="player-card"
      >
        <img :src="`https://crafatar.com/avatars/${p.uuid}?size=64&overlay`" :alt="p.last_known_name" class="avatar" />
        <div class="info">
          <h3>{{ p.last_known_name || 'Unknown' }}</h3>
          <span class="uuid">{{ p.uuid.split('-')[0] }}...</span>
        </div>
      </RouterLink>
    </div>
  </div>
</template>

<style scoped>
.players-view {
  max-width: 1200px;
  margin: 0 auto;
}
.header {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  margin-bottom: 2rem;
}
@media (min-width: 768px) {
  .header {
    flex-direction: row;
    justify-content: space-between;
    align-items: flex-end;
  }
}
.title-area h1 {
  font-size: 2.5rem;
  color: #00ff88;
  margin: 0 0 0.5rem 0;
}
.subtitle {
  color: #aaa;
  margin: 0;
}
.controls {
  display: flex;
  gap: 0.5rem;
  flex-wrap: wrap;
}
.search-box {
  padding: 0.75rem 1rem;
  border-radius: 4px;
  border: 1px solid #444;
  background: #222;
  color: white;
  flex: 1;
  min-width: 200px;
}
.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
}
.btn.primary {
  background-color: #00ff88;
  color: #000;
}
.btn.secondary {
  background-color: #444;
  color: #fff;
}
.grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 1.5rem;
}
.player-card {
  background: #1a1a1a;
  border: 1px solid #333;
  border-radius: 8px;
  padding: 1.5rem;
  display: flex;
  align-items: center;
  gap: 1rem;
  text-decoration: none;
  color: inherit;
  transition: transform 0.2s, border-color 0.2s;
}
.player-card:hover {
  transform: translateY(-2px);
  border-color: #00ff88;
}
.avatar {
  width: 64px;
  height: 64px;
  border-radius: 8px;
  image-rendering: pixelated;
}
.info h3 {
  margin: 0 0 0.25rem 0;
  color: #fff;
}
.uuid {
  font-family: monospace;
  color: #666;
  font-size: 0.875rem;
}
.loading, .empty-state {
  text-align: center;
  padding: 4rem;
  color: #888;
}
</style>
