<script setup lang="ts">
import { ref, onUnmounted, nextTick, watch } from 'vue'
import Chart from 'chart.js/auto'

const props = defineProps<{
  oreStats: any[]
}>()

const emit = defineEmits(['resetEnlarged'])

const pieChartCanvas = ref<HTMLCanvasElement | null>(null)
let chartInstance: Chart | null = null

const initPieChart = () => {
  if (!pieChartCanvas.value || !props.oreStats || props.oreStats.length === 0) {
    return
  }
  
  if (chartInstance) chartInstance.destroy()

  const labels = props.oreStats.map((o: any) => o.name)
  const minedData = props.oreStats.map((o: any) => o.mined)

  chartInstance = new Chart(pieChartCanvas.value, {
    type: 'pie',
    data: {
      labels,
      datasets: [
        {
          data: minedData,
          backgroundColor: [
            '#c9ff00', '#2563EB', '#8b5cf6', '#ec4899', '#f43f5e', 
            '#f59e0b', '#10b981', '#14b8a6', '#64748b', '#ffffff'
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
          labels: { color: '#8a8a8a', font: { family: 'Manrope' } }
        }
      }
    }
  })
}

watch(() => props.oreStats, () => {
  nextTick(initPieChart)
}, { deep: true, immediate: true })

onUnmounted(() => {
  if (chartInstance) chartInstance.destroy()
})

const formatNumber = (num: number) => {
  if (num >= 1000000) return (num / 1000000).toFixed(1) + 'M'
  if (num >= 1000) return (num / 1000).toFixed(1) + 'K'
  return num.toLocaleString()
}
</script>

<template>
  <section class="panel glass-card" v-if="oreStats && oreStats.length > 0">
    <h3 @click.stop="emit('resetEnlarged')">
      <img src="/icons/all_blocks.ico" class="header-icon" /> Ore Mined Statistics
    </h3>
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
</template>

<style scoped>
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
</style>
