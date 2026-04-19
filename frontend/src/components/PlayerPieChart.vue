<script setup lang="ts">
import { nextTick, onUnmounted, ref, watch } from 'vue'

import { siteContent } from '@/content/siteContent'
import type { OreStat } from '@/types/api'

const props = defineProps<{
  oreStats: OreStat[]
}>()

const pieChartCanvas = ref<HTMLCanvasElement | null>(null)
let chartInstance: import('chart.js').Chart | null = null

async function initPieChart(): Promise<void> {
  if (!pieChartCanvas.value || props.oreStats.length === 0) {
    return
  }

  const { default: Chart } = await import('chart.js/auto')

  chartInstance?.destroy()

  chartInstance = new Chart(pieChartCanvas.value, {
    type: 'pie',
    data: {
      labels: props.oreStats.map((ore) => ore.name),
      datasets: [
        {
          data: props.oreStats.map((ore) => ore.mined),
          backgroundColor: ['#5b71f6', '#3b82f6', '#7dd3fc', '#8b5cf6', '#38bdf8', '#1d4ed8'],
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
          labels: {
            color: '#c4b69d',
            font: { family: 'Instrument Sans' }
          }
        }
      }
    }
  })
}

watch(
  () => props.oreStats,
  async () => {
    await nextTick()
    await initPieChart()
  },
  { deep: true, immediate: true }
)

onUnmounted(() => {
  chartInstance?.destroy()
})

function formatNumber(value: number): string {
  if (value >= 1_000_000) {
    return `${(value / 1_000_000).toFixed(1)}M`
  }

  if (value >= 1_000) {
    return `${(value / 1_000).toFixed(1)}K`
  }

  return value.toLocaleString()
}
</script>

<template>
  <section v-if="oreStats.length > 0" class="glass-card panel-card">
    <div class="panel-head">
      <h3>{{ siteContent.playerDetail.sections.ores }}</h3>
    </div>
    <div class="ore-layout">
      <div class="chart-wrap">
        <canvas ref="pieChartCanvas"></canvas>
      </div>
      <div class="ore-list">
        <div v-for="ore in oreStats.slice(0, 8)" :key="ore.name" class="ore-row">
          <span>{{ ore.name }}</span>
          <strong>{{ formatNumber(ore.mined) }}</strong>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.panel-card {
  display: grid;
  gap: 1rem;
}

.panel-head h3 {
  font-size: 1.8rem;
}

.ore-layout {
  display: grid;
  grid-template-columns: 260px 1fr;
  gap: 1rem;
  align-items: center;
}

.chart-wrap {
  height: 240px;
}

.ore-list {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 0.75rem;
}

.ore-row {
  display: flex;
  justify-content: space-between;
  gap: 1rem;
  padding: 0.9rem 1rem;
  border-radius: 16px;
  background: rgba(255, 248, 234, 0.04);
  border: 1px solid rgba(255, 248, 234, 0.06);
}

.ore-row span {
  color: var(--text-muted);
}

.ore-row strong {
  color: var(--text-strong);
}

@media (max-width: 860px) {
  .ore-layout,
  .ore-list {
    grid-template-columns: 1fr;
  }
}
</style>
