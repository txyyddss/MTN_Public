<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { API_BASE_URL } from '@/config'
import { usePlayers, type PlayerInfo } from '@/composables/usePlayers'

const router = useRouter()
const {
  players, count, activeDays, searchQuery, showAll, loading, onlinePlayers, sortedPlayers,
  fetchPlayers, fetchOnline
} = usePlayers()


const handleRandom = async () => {
  try {
    const res = await fetch(`${API_BASE_URL}/api/players/random`)
    if (!res.ok) return
    const json = await res.json()
    if (json.uuid) {
      router.push(`/player/${json.uuid}`)
    }
  } catch (e) {
    console.error('Failed to fetch random player', e)
  }
}

const isOnline = (uuid: string) => {
    return onlinePlayers.value.includes(uuid)
}

const getAvatarUrl = (p: PlayerInfo) => {
  if (!p.last_known_name) return 'https://mineskin.eu/helm/Steve'
  let cleanName = p.last_known_name
  if (p.type === 'Bedrock' || cleanName.startsWith('.')) {
      cleanName = cleanName.replace(/^\./, '').replace(/^BE_/, '')
  }
  return `https://mineskin.eu/helm/${cleanName}`
}

onMounted(() => {
  fetchPlayers()
  fetchOnline()
})
</script>

<template>
  <div class="players-view container animate-entry">
    <div class="header">
      <div class="title-area">
        <h1>Players Directory</h1>
        <p class="subtitle" v-if="loading">Searching...</p>
        <p class="subtitle" v-else-if="!searchQuery && !showAll">
          Showing {{ count }} {{ count === 1 ? 'player' : 'players' }} active in last {{ activeDays }} days
        </p>
        <p class="subtitle" v-else-if="!searchQuery && showAll">
          Showing all {{ count }} unique {{ count === 1 ? 'player' : 'players' }}
        </p>
        <p class="subtitle" v-else>
          Found {{ count }} {{ count === 1 ? 'player' : 'players' }} matching "{{ searchQuery }}"
        </p>
      </div>

      <div class="controls glass-card">
        <input 
          v-model="searchQuery" 
          placeholder="Type to search..." 
          class="search-box"
        />
        
        <div class="filter-group">
            <button @click="showAll = !showAll" :class="['btn-filter', { active: showAll }]">
                {{ showAll ? 'Showing All' : 'Showing Recent' }}
            </button>
            <button @click="handleRandom" class="btn-icon" title="Random Player">🎲</button>
        </div>
      </div>
    </div>

    <div v-if="loading && players.length === 0" class="loading-state">
      <div class="spinner"></div>
      <p>Loading players...</p>
    </div>

    <div v-else-if="players.length === 0" class="empty-state glass-card">
      <h3>No players found</h3>
      <p>Try searching for a different username or UUID.</p>
      <button @click="() => { searchQuery = ''; showAll = false; }" class="btn-primary mt-4">Reset Search</button>
    </div>

    <div v-else class="player-grid">
      <RouterLink 
        v-for="(p, index) in sortedPlayers" 
        :key="p.uuid" 
        :to="`/player/${p.uuid}`"
        class="player-card glass-card animate-entry"
        :style="{ animationDelay: `${(index % 20) * 30}ms` }"
      >
        <div class="avatar-wrap">
          <img :src="getAvatarUrl(p)" :alt="p.last_known_name" class="avatar" loading="lazy" />
          <div v-if="isOnline(p.uuid)" class="online-indicator" title="Online now"></div>
        </div>
        
        <div class="info">
          <div class="name-row">
            <span :class="['type-tag', p.type?.toLowerCase()]">{{ p.type === 'Bedrock' ? 'BE' : 'JE' }}</span>
            <h3 :class="['player-name', { 'online-name': isOnline(p.uuid) }]">{{ p.last_known_name || 'Unknown' }}</h3>
          </div>
          <span class="last-seen">Seen {{ new Date(p.last_seen).toLocaleDateString() }}</span>
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
  padding: 1rem;
  align-items: center;
  flex-wrap: wrap;
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
}

.search-box:focus {
  outline: none;
  border-color: var(--primary);
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.2);
}

.filter-group {
    display: flex;
    gap: 0.5rem;
}

.btn-filter {
    background: rgba(255,255,255,0.05);
    border: 1px solid var(--glass-border);
    color: var(--text-muted);
    padding: 8px 16px;
    border-radius: var(--radius-sm);
    cursor: pointer;
    font-weight: 600;
    transition: all 0.2s;
}
.btn-filter.active {
    background: var(--primary);
    color: #000;
    border-color: var(--primary);
}

.btn-icon {
    background: rgba(255,255,255,0.05);
    border: 1px solid var(--glass-border);
    padding: 8px;
    border-radius: var(--radius-sm);
    cursor: pointer;
    filter: grayscale(1);
    transition: all 0.2s;
}
.btn-icon:hover { filter: grayscale(0); transform: scale(1.1); }

.player-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 1.25rem;
}

.player-card {
  display: flex;
  align-items: center;
  gap: 1rem;
  text-decoration: none;
  color: inherit;
  padding: 1rem;
  position: relative;
}

.player-card:hover {
  border-color: var(--primary);
  transform: translateY(-2px);
}

.avatar-wrap {
  width: 50px;
  height: 50px;
  position: relative;
  flex-shrink: 0;
}

.avatar {
  width: 100%;
  height: 100%;
  border-radius: 6px;
  image-rendering: pixelated;
  background: #111;
}

.online-indicator {
    position: absolute;
    bottom: -2px;
    right: -2px;
    width: 14px;
    height: 14px;
    background: #10b981;
    border: 2px solid var(--bg-dark);
    border-radius: 50%;
    box-shadow: 0 0 8px #10b981;
}

.info { overflow: hidden; }

.name-row {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 2px;
}

.player-name {
  margin: 0;
  font-size: 1.05rem;
  font-weight: 700;
  color: #fff;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  font-family: var(--minecraft);
  transition: color 0.3s;
}

.player-name.online-name {
    color: #10b981;
    text-shadow: 0 0 10px rgba(16, 185, 129, 0.4);
}

.type-tag {
    font-size: 0.65rem;
    padding: 1px 4px;
    border-radius: 3px;
    font-weight: 800;
}
.type-tag.java { background: rgba(59, 130, 246, 0.2); color: #93c5fd; }
.type-tag.bedrock { background: rgba(79, 70, 229, 0.2); color: #c7d2fe; }

.last-seen {
  color: var(--text-muted);
  font-size: 0.75rem;
}

.card-arrow {
  margin-left: auto;
  font-size: 1.2rem;
  color: var(--text-muted);
  opacity: 0.3;
  transition: 0.3s;
}
.player-card:hover .card-arrow { opacity: 1; color: var(--primary); transform: translateX(3px); }

.spinner {
  width: 40px;
  height: 40px;
  border: 3px solid rgba(59, 130, 246, 0.1);
  border-left-color: var(--primary);
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 1rem;
}

@keyframes spin { 100% { transform: rotate(360deg); } }
.mt-4 { margin-top: 1.5rem; }
</style>
