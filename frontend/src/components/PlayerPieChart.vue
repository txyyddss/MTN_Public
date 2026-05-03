<script setup lang="ts">
import { computed, nextTick, onUnmounted, ref, watch } from 'vue'
import type { Chart, ChartData, ChartOptions } from 'chart.js'

import PlayerCollapsiblePanel from '@/components/player/PlayerCollapsiblePanel.vue'
import { getLocaleValue, getOreLabel, useCurrentLocale, useSiteContent } from '@/content/siteContent'
import type { OreStat } from '@/types/api'

const props = defineProps<{
  oreStats: OreStat[]
}>()

const pieChartCanvas = ref<HTMLCanvasElement | null>(null)
type PieChart = Chart<'pie', number[], string>
type ChartConstructor = typeof import('chart.js/auto')['default']

let chartInstance: PieChart | null = null
let chartLoader: Promise<ChartConstructor> | null = null
let chartRenderToken = 0
const siteContent = useSiteContent()
const currentLocale = useCurrentLocale()
const visibleOreStats = computed(() => props.oreStats.slice(0, 8))

function loadChart(): Promise<ChartConstructor> {
  chartLoader ??= import('chart.js/auto').then((module) => module.default)
  return chartLoader
}

function getCssVar(name: string, fallback: string): string {
  if (typeof window === 'undefined') {
    return fallback
  }

  return getComputedStyle(document.documentElement).getPropertyValue(name).trim() || fallback
}

function getChartColors(): string[] {
  const primaryRgb = getCssVar('--primary-rgb', '91, 113, 246')
  const secondaryRgb = getCssVar('--secondary-rgb', '59, 130, 246')

  return [
    getCssVar('--primary', '#5b71f6'),
    getCssVar('--secondary', '#3b82f6'),
    getCssVar('--accent', '#93c5fd'),
    getCssVar('--accent-soft', '#c9dcff'),
    `rgba(${primaryRgb}, 0.74)`,
    `rgba(${secondaryRgb}, 0.48)`
  ]
}

function createChartData(): ChartData<'pie', number[], string> {
  return {
    labels: props.oreStats.map((ore) => getOreLabel(ore.name)),
    datasets: [
      {
        data: props.oreStats.map((ore) => ore.mined),
        backgroundColor: getChartColors(),
        borderWidth: 0
      }
    ]
  }
}

function createChartOptions(): ChartOptions<'pie'> {
  return {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        position: 'right',
        labels: {
          color: getCssVar('--text-main', '#edf4ff'),
          font: { family: getCssVar('--sans', 'system-ui') }
        }
      }
    }
  }
}

async function syncPieChart(): Promise<void> {
  const token = ++chartRenderToken

  if (!pieChartCanvas.value || props.oreStats.length === 0) {
    chartInstance?.destroy()
    chartInstance = null
    return
  }

  const Chart = await loadChart()

  if (token !== chartRenderToken || !pieChartCanvas.value) {
    return
  }

  if (!chartInstance) {
    chartInstance = new Chart(pieChartCanvas.value, {
      type: 'pie',
      data: createChartData(),
      options: createChartOptions()
    })
    return
  }

  chartInstance.data = createChartData()
  chartInstance.options = createChartOptions()
  chartInstance.update()
}

watch(
  () => props.oreStats,
  async () => {
    await nextTick()
    await syncPieChart()
  },
  { deep: true, immediate: true }
)

watch(currentLocale, async () => {
  await nextTick()
  await syncPieChart()
})

onUnmounted(() => {
  chartRenderToken++
  chartInstance?.destroy()
  chartInstance = null
})

function formatNumber(value: number): string {
  void currentLocale.value

  return new Intl.NumberFormat(getLocaleValue(), {
    notation: value >= 10_000 ? 'compact' : 'standard',
    maximumFractionDigits: value >= 10_000 ? 1 : 0
  }).format(value)
}
</script>

<template>
  <PlayerCollapsiblePanel v-if="oreStats.length > 0" class="panel-card" :title="siteContent.playerDetail.sections.ores">
    <template #summary>
      <span class="meta-chip">{{ siteContent.playerDetail.sections.top.replace('{count}', String(Math.min(oreStats.length, 8))) }}</span>
    </template>

    <div class="ore-layout">
      <div class="chart-wrap">
        <canvas ref="pieChartCanvas"></canvas>
      </div>
      <div class="ore-list">
        <div v-for="ore in visibleOreStats" :key="ore.name" class="ore-row">
          <span>{{ getOreLabel(ore.name) }}</span>
          <strong>{{ formatNumber(ore.mined) }}</strong>
        </div>
      </div>
    </div>
  </PlayerCollapsiblePanel>
</template>

<style scoped>
.panel-card {
  display: grid;
  gap: 0.85rem;
}

.ore-layout {
  display: grid;
  grid-template-columns: 220px 1fr;
  gap: 0.85rem;
  align-items: center;
}

.chart-wrap {
  height: 220px;
}

.ore-list {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 0.6rem;
}

.ore-row {
  display: flex;
  justify-content: space-between;
  gap: 1rem;
  padding: 0.75rem 0.85rem;
  border-radius: 16px;
  background: var(--player-glass-tile-bg);
  border: 1px solid var(--player-glass-border-soft);
  box-shadow: var(--glass-inset);
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
  font-size: 0.68rem;
  border: 1px solid var(--player-glass-border-soft);
  background: rgba(var(--secondary-rgb), 0.1);
}

@media (max-width: 860px) {
  .ore-layout,
  .ore-list {
    grid-template-columns: 1fr;
  }
}
</style>
