<script setup lang="ts">
import { ref, onMounted, onUnmounted, nextTick } from 'vue'
import { useRoute } from 'vue-router'
import { SkinViewer, IdleAnimation } from 'skinview3d'

const route = useRoute()
const uuid = route.params.uuid as string
const loading = ref(true)

const info = ref<any>(null)
const stats = ref<any>(null)
const advancements = ref<any>(null)
const mcmmo = ref<any>(null)
const linkedAccount = ref<any>(null)
const oreStats = ref<any>([])

const viewerCanvas = ref<HTMLCanvasElement | null>(null)
let viewer: SkinViewer | null = null

const fetchDetail = async () => {
  try {
    const res = await fetch(`/api/players/${uuid}`)
    const json = await res.json()
    info.value = json.info
    stats.value = json.stats
    advancements.value = json.advancements
    mcmmo.value = json.mcmmo
    linkedAccount.value = json.linked_account
    oreStats.value = json.ore_stats
    
    // Setup SkinViewer
    nextTick(() => {
      if (viewerCanvas.value) {
        viewer = new SkinViewer({
          canvas: viewerCanvas.value,
          width: 200,
          height: 300,
          skin: `https://crafatar.com/skins/${uuid}`
        })
        viewer.animation = new IdleAnimation()
      }
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
  if (viewer) {
    viewer.dispose()
  }
})

const formatDate = (ms: number) => new Date(ms).toLocaleString()
const formatPlaytime = (ticks: number) => {
  const hours = (ticks / 20 / 3600).toFixed(1)
  return `${hours}h`
}
</script>

<template>
  <div class="player-detail">
    <div v-if="loading" class="loading">Loading details...</div>
    
    <div v-else-if="!info" class="empty-state">
      Player not found.
    </div>
    
    <div v-else class="content-grid">
      <!-- Sidebar profile -->
      <aside class="profile-card">
        <h2>{{ info.last_known_name }}</h2>
        <span class="uuid">{{ info.uuid }}</span>
        
        <div class="canvas-wrapper">
          <canvas ref="viewerCanvas"></canvas>
        </div>
        
        <div class="basic-info">
          <p><strong>First Join:</strong> {{ formatDate(info.first_played) }}</p>
          <p><strong>Last Activity:</strong> {{ formatDate(info.last_seen) }}</p>
          <p><strong>Playtime:</strong> {{ formatPlaytime(info.ticks_lived) }}</p>
          <p v-if="linkedAccount">
            <strong>Linked Account:</strong> {{ linkedAccount.discord_id || 'Yes' }}
          </p>
        </div>
      </aside>

      <!-- Main content -->
      <div class="details">
        <section class="panel" v-if="oreStats && oreStats.length > 0">
          <h3>Ore Statistics</h3>
          <div class="ore-grid">
            <div class="ore-item" v-for="ore in oreStats" :key="ore.name">
              <span class="ore-name">{{ ore.name }}</span>
              <div class="ore-bars">
                <div class="stat">M: {{ ore.mined }}</div>
                <div class="stat">U: {{ ore.used }}</div>
              </div>
            </div>
          </div>
        </section>

        <section class="panel" v-if="mcmmo">
          <h3>Skills</h3>
          <div class="skill-grid">
            <div class="skill-item">Level: {{ mcmmo.level }}</div>
            <!-- Populate other skills if they exist -->
          </div>
        </section>
      </div>
    </div>
  </div>
</template>

<style scoped>
.player-detail {
  max-width: 1200px;
  margin: 0 auto;
}
.loading, .empty-state {
  text-align: center;
  padding: 4rem;
  color: #888;
}
.content-grid {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}
@media (min-width: 768px) {
  .content-grid {
    flex-direction: row;
  }
  .profile-card {
    width: 300px;
    flex-shrink: 0;
  }
}
.profile-card, .panel {
  background: #1a1a1a;
  border: 1px solid #333;
  border-radius: 8px;
  padding: 1.5rem;
}
.profile-card h2 {
  margin: 0;
  color: #00ff88;
}
.uuid {
  font-family: monospace;
  color: #666;
  font-size: 0.875rem;
}
.canvas-wrapper {
  margin: 1.5rem 0;
  display: flex;
  justify-content: center;
  background: #111;
  border-radius: 8px;
  padding: 1rem;
}
.basic-info p {
  margin: 0.5rem 0;
  font-size: 0.9rem;
  color: #ccc;
}
.basic-info strong {
  color: #fff;
}
.details {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}
.panel h3 {
  margin-top: 0;
  border-bottom: 2px solid #333;
  padding-bottom: 0.5rem;
  color: #fff;
}
.ore-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 1rem;
}
.ore-item {
  background: #222;
  border-radius: 4px;
  padding: 0.75rem;
}
.ore-name {
  font-weight: bold;
  color: #00ff88;
  display: block;
  margin-bottom: 0.5rem;
}
.ore-bars {
  display: flex;
  justify-content: space-between;
  font-size: 0.875rem;
  color: #aaa;
}
</style>
