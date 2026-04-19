<script setup lang="ts">
import { nextTick, onUnmounted, ref, watch } from 'vue'

import PlayerCollapsiblePanel from '@/components/player/PlayerCollapsiblePanel.vue'
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
          backgroundColor: ['#3d78ff', '#1d4ed8', '#5b8cff', '#102d7a', '#6cb4ff', '#0a1637'],
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
            color: '#d6e4ff',
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
  <PlayerCollapsiblePanel v-if="oreStats.length > 0" class="panel-card" :title="siteContent.playerDetail.sections.ores">
    <template #summary>
      <span class="meta-chip">Top {{ Math.min(oreStats.length, 8) }}</span>
    </template>

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
  </PlayerCollapsiblePanel>
</template>

<style scoped>
.panel-card {
  display: grid;
  gap: 1rem;
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
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid rgba(255, 255, 255, 0.06);
}

.ore-row span {
  color: var(--text-muted);
}

.ore-row strong {
  color: var(--text-strong);
}

.meta-chip {
  padding: 0.45rem 0.7rem;
  border-radius: 999px;
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.74rem;
  border: 1px solid rgba(255, 255, 255, 0.08);
}

@media (max-width: 860px) {
  .ore-layout,
  .ore-list {
    grid-template-columns: 1fr;
  }
}
</style>
