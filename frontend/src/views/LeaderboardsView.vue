<script setup lang="ts">
import { ref, onMounted } from 'vue'

const types = ['skills', 'playtime', 'mining', 'killing', 'deaths', 'walking', 'pvp']
const currentType = ref('mining')
const entries = ref<any[]>([])
const loading = ref(false)

const fetchLeaderboard = async (type: string) => {
  loading.value = true
  currentType.value = type
  try {
    const res = await fetch(`/api/leaderboards/${type}`)
    const json = await res.json()
    entries.value = json.entries || []
  } catch(e) {
    console.error('Failed', e)
  }
  loading.value = false
}

onMounted(() => fetchLeaderboard(currentType.value))

const formatValue = (val: number, type: string) => {
  if (type === 'playtime') return (val / 20 / 3600).toFixed(1) + ' hours'
  if (type === 'walking') return (val / 100).toFixed(0) + ' m'
  return val.toLocaleString()
}
</script>

<template>
  <div class="leaderboards">
    <div class="header">
      <h1>Leaderboards</h1>
      <div class="type-selector">
        <button 
          v-for="t in types" 
          :key="t" 
          @click="fetchLeaderboard(t)"
          :class="['btn', { active: currentType === t }]"
        >
          {{ t.charAt(0).toUpperCase() + t.slice(1) }}
        </button>
      </div>
    </div>

    <div v-if="loading" class="loading">Loading...</div>
    <div v-else-if="entries.length === 0" class="empty-state">No data available</div>
    
    <div v-else class="lb-table-wrapper">
      <table class="lb-table">
        <thead>
          <tr>
            <th class="rank-col">Rank</th>
            <th>Player</th>
            <th class="val-col">Score</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="entry in entries" :key="entry.uuid">
            <td class="rank-col">#{{ entry.rank }}</td>
            <td>
              <RouterLink :to="`/player/${entry.uuid}`" class="player-link">
                {{ entry.name }}
              </RouterLink>
            </td>
            <td class="val-col">{{ formatValue(entry.value, currentType) }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped>
.leaderboards {
  max-width: 900px;
  margin: 0 auto;
}
.header {
  margin-bottom: 2rem;
}
h1 {
  color: #00ff88;
}
.type-selector {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin-top: 1rem;
}
.btn {
  background: #222;
  color: #ccc;
  border: 1px solid #333;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.2s;
}
.btn:hover {
  border-color: #00ff88;
  color: #fff;
}
.btn.active {
  background: #00ff88;
  color: #000;
  border-color: #00ff88;
  font-weight: bold;
}
.lb-table-wrapper {
  background: #1a1a1a;
  border-radius: 8px;
  border: 1px solid #333;
  overflow: hidden;
}
.lb-table {
  width: 100%;
  border-collapse: collapse;
}
.lb-table th {
  background: #222;
  text-align: left;
  padding: 1rem;
  color: #888;
  font-weight: normal;
  border-bottom: 1px solid #333;
}
.lb-table td {
  padding: 1rem;
  border-bottom: 1px solid #333;
}
.lb-table tr:last-child td {
  border-bottom: none;
}
.rank-col {
  width: 80px;
  color: #aaa;
}
.val-col {
  text-align: right;
  font-family: monospace;
  font-size: 1.1em;
}
.player-link {
  color: #fff;
  text-decoration: none;
  font-weight: bold;
}
.player-link:hover {
  color: #00ff88;
  text-decoration: underline;
}
.loading, .empty-state {
  padding: 3rem;
  text-align: center;
  color: #888;
}
</style>
