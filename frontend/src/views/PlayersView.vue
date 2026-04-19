<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { API_BASE_URL } from '@/config'
import { usePlayers } from '@/composables/usePlayers'

const router = useRouter()
const {
  players, count, activeDays, searchQuery, showAll, loading, onlinePlayers, sortedPlayers,
  fetchPlayers, fetchOnline, startAutoRefresh, stopAutoRefresh
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

import { getAvatarUrl } from '@/utils/minecraft'

onMounted(() => {
  fetchPlayers()
  fetchOnline()
  startAutoRefresh(5000)
})

onUnmounted(() => {
  stopAutoRefresh()
})
</script>

<template>
  <div class="players-view container">
    <!-- Sophisticated Header -->
    <header class="view-header animate-entry">
      <div class="title-area">
        <h1>DIRECTORY</h1>
        <div class="stats-pills">
           <span class="pill" v-if="loading">Searching...</span>
           <span class="pill" v-else>
             {{ count }} {{ count === 1 ? 'Player' : 'Players' }} 
             <span v-if="!searchQuery && !showAll">in last {{ activeDays }} days</span>
             <span v-else-if="searchQuery">matching "{{ searchQuery }}"</span>
           </span>
        </div>
      </div>

      <div class="controls-panel glass-card">
        <div class="search-input-wrap">
          <svg class="search-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
          <input 
            v-model="searchQuery" 
            placeholder="Search by username or UUID..." 
            class="prime-search"
          />
        </div>
        
        <div class="action-group">
            <button @click="showAll = !showAll" :class="['btn-toggle', { active: showAll }]">
                {{ showAll ? 'ALL' : 'ACTIVE' }}
            </button>
            <button @click="handleRandom" class="btn-hero-action" title="Random Player">
              <span class="icon">🎲</span>
            </button>
        </div>
      </div>
    </header>

    <!-- Content States -->
    <div v-if="loading && players.length === 0" class="loading-state">
      <div class="loader"></div>
      <p>Consulting the archives...</p>
    </div>

    <div v-else-if="players.length === 0" class="empty-state glass-card animate-entry">
      <div class="empty-icon">📂</div>
      <h3>No matches found</h3>
      <p>We couldn't find any players matching your criteria.</p>
      <button @click="() => { searchQuery = ''; showAll = false; }" class="btn-primary mt-4">Reset Search</button>
    </div>

    <!-- Players Grid -->
    <div v-else class="player-grid">
      <RouterLink 
        v-for="(p, index) in sortedPlayers" 
        :key="p.uuid" 
        :to="`/player/${p.uuid}`"
        class="player-card glass-card animate-entry"
        :style="{ animationDelay: `${(index % 30) * 40}ms` }"
      >
        <div class="card-bg-glow"></div>
        <div class="avatar-container">
          <img :src="getAvatarUrl(p.last_known_name, p.type)" :alt="p.last_known_name" class="avatar" loading="lazy" />
          <div v-if="isOnline(p.uuid)" class="status-indicator online" title="Online now"></div>
        </div>
        
        <div class="player-info">
          <div class="name-box">
            <h3 :class="['player-name', 'minecraft-font', { online: isOnline(p.uuid) }]">
              {{ p.last_known_name || 'Anonymous' }}
            </h3>
            <span :class="['platform-badge', p.type?.toLowerCase()]">{{ p.type === 'Bedrock' ? 'BE' : 'JE' }}</span>
          </div>
          <div class="metadata">
            <span class="label">LAST SEEN</span>
            <span class="value">{{ new Date(p.last_seen).toLocaleDateString() }}</span>
          </div>
        </div>
        
        <div class="card-action">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="m12 5 7 7-7 7"/></svg>
        </div>
      </RouterLink>
    </div>
  </div>
</template>

<style scoped>
.players-view {
  padding-top: 4rem;
  padding-bottom: 8rem;
}

/* Header Refinement */
.view-header {
  display: flex;
  flex-direction: column;
  gap: 3rem;
  margin-bottom: 5rem;
}

@media (min-width: 1024px) {
  .view-header {
    flex-direction: row;
    justify-content: space-between;
    align-items: flex-end;
  }
}

.title-area h1 {
  font-size: 4rem;
  font-weight: 800;
  letter-spacing: -2px;
  line-height: 0.9;
  margin-bottom: 1rem;
}

.stats-pills {
  display: flex;
  gap: 8px;
}

.pill {
  font-size: 0.75rem;
  font-weight: 800;
  text-transform: uppercase;
  letter-spacing: 1px;
  color: var(--primary);
  background: rgba(59, 130, 246, 0.1);
  padding: 4px 12px;
  border-radius: 100px;
  border: 1px solid rgba(59, 130, 246, 0.2);
}

/* Controls Panel */
.controls-panel {
  display: flex;
  padding: 12px 12px 12px 24px !important;
  border-radius: 100px !important;
  background: rgba(10, 10, 10, 0.6) !important;
  align-items: center;
  gap: 1.5rem;
  min-width: 450px;
}

@media (max-width: 600px) {
  .controls-panel {
    min-width: 100%;
    border-radius: 20px !important;
    padding: 16px !important;
    flex-direction: column;
  }
}

.search-input-wrap {
  position: relative;
  flex: 1;
  display: flex;
  align-items: center;
}

.search-icon {
  width: 18px;
  height: 18px;
  color: var(--text-muted);
  margin-right: 12px;
}

.prime-search {
  background: transparent;
  border: none;
  color: #fff;
  font-size: 1rem;
  outline: none;
  width: 100%;
}

.action-group {
  display: flex;
  gap: 8px;
}

.btn-toggle {
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid var(--glass-border);
  color: #fff;
  font-family: var(--heading);
  font-weight: 800;
  font-size: 0.75rem;
  padding: 8px 20px;
  border-radius: 100px;
  cursor: pointer;
  transition: all 0.3s;
}

.btn-toggle.active {
  background: var(--primary);
  color: #000;
  border-color: var(--primary);
}

.btn-hero-action {
  width: 38px;
  height: 38px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid var(--glass-border);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.2rem;
  transition: all 0.3s;
}

.btn-hero-action:hover {
  transform: rotate(15deg) scale(1.1);
  background: rgba(255, 255, 255, 0.08);
  border-color: var(--primary);
}

/* Player Grid */
.player-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 1.5rem;
}

.player-card {
  padding: 24px !important;
  display: flex;
  align-items: center;
  gap: 1.5rem;
  transition: all 0.4s var(--transition-slow);
}

.card-bg-glow {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: radial-gradient(circle at 100% 0%, var(--primary-glow), transparent 40%);
  opacity: 0;
  transition: opacity 0.4s;
}

.player-card:hover .card-bg-glow {
  opacity: 1;
}

.avatar-container {
  width: 64px;
  height: 64px;
  position: relative;
  flex-shrink: 0;
}

.avatar {
  width: 100%;
  height: 100%;
  background: #111;
  border-radius: 12px;
  image-rendering: pixelated;
  border: 1px solid var(--glass-border);
}

.status-indicator {
  position: absolute;
  bottom: -4px;
  right: -4px;
  width: 16px;
  height: 16px;
  border-radius: 50%;
  border: 3px solid #050505;
}

.status-indicator.online {
  background: #10B981;
  box-shadow: 0 0 10px #10B981;
  animation: pulse-simple 2s infinite;
}

@keyframes pulse-simple {
  0% { transform: scale(1); }
  50% { transform: scale(1.2); }
  100% { transform: scale(1); }
}

.player-info {
  flex: 1;
  min-width: 0;
}

.name-box {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 8px;
}

.player-name {
  font-size: 1.25rem;
  color: #fff;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  line-height: 1;
}

.player-name.online {
  color: #10B981;
}

.platform-badge {
  font-size: 0.65rem;
  font-weight: 800;
  padding: 2px 6px;
  border-radius: 4px;
  background: rgba(255, 255, 255, 0.05);
}

.platform-badge.java { color: #93c5fd; }
.platform-badge.bedrock { color: #c4b5fd; }

.metadata {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.metadata .label {
  font-size: 0.65rem;
  font-weight: 800;
  color: var(--text-muted);
  letter-spacing: 1px;
}

.metadata .value {
  font-size: 0.85rem;
  color: #fff;
  font-weight: 500;
}

.card-action {
  color: var(--text-muted);
  opacity: 0.2;
  transition: all 0.3s;
}

.player-card:hover .card-action {
  opacity: 1;
  color: var(--primary);
  transform: translateX(5px);
}

.empty-state {
  text-align: center;
  padding: 6rem 2rem !important;
}

.empty-icon {
  font-size: 4rem;
  margin-bottom: 2rem;
  opacity: 0.5;
}
</style>
