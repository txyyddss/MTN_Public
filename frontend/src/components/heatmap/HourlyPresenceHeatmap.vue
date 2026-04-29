<script setup lang="ts">
import { computed } from 'vue'

import { formatPlayerCount, useSiteContent } from '@/content/siteContent'
import type { HeatmapDay } from '@/types/api'

const props = defineProps<{
  days: HeatmapDay[]
  hours: number[]
  cells: Array<Array<number | boolean>>
  timezone: string
  variant: 'server' | 'player'
  weeklyMaxPlayers?: number
}>()

const siteContent = useSiteContent()
const isEmpty = computed(() => props.days.length === 0 || props.hours.length === 0 || props.cells.length === 0)
const emptyCellStyle = {
  backgroundColor: 'var(--heatmap-cell-empty-bg)',
  borderColor: 'var(--heatmap-cell-empty-border)'
}

function interpolateChannel(start: number, end: number, ratio: number): number {
  return Math.round(start + (end - start) * ratio)
}

function getWeekdayLabel(weekday: string): string {
  const key = weekday.toLowerCase()
  return siteContent.value.heatmap.weekdays[key as keyof typeof siteContent.value.heatmap.weekdays] ?? weekday
}

function getCellStyle(value: number | boolean): { backgroundColor: string; borderColor: string } {
  if (props.variant === 'player') {
    return value
      ? {
          backgroundColor: 'var(--heatmap-player-active)',
          borderColor: 'var(--heatmap-player-active-border)'
        }
      : emptyCellStyle
  }

  const numericValue = typeof value === 'number' ? value : 0
  if (numericValue <= 0 || !props.weeklyMaxPlayers) {
    return emptyCellStyle
  }

  const ratio = Math.min(numericValue / props.weeklyMaxPlayers, 1)
  const red = interpolateChannel(18, 59, ratio)
  const green = interpolateChannel(32, 130, ratio)
  const blue = interpolateChannel(74, 246, ratio)
  const alpha = 0.28 + ratio * 0.5

  return {
    backgroundColor: `rgb(${red}, ${green}, ${blue})`,
    borderColor: `rgba(${red}, ${green}, ${blue}, ${alpha.toFixed(2)})`
  }
}

function getTooltip(day: HeatmapDay, hourLabel: number, value: number | boolean): string {
  const hourStart = String(hourLabel - 1).padStart(2, '0')
  const hourEnd = String(hourLabel % 24).padStart(2, '0')
  const range = `${hourStart}:00-${hourEnd}:00`

  if (props.variant === 'player') {
    const statusLabel = value ? siteContent.value.heatmap.online : siteContent.value.heatmap.offline
    return `${day.date} ${range} (${props.timezone}) · ${statusLabel}`
  }

  const numericValue = typeof value === 'number' ? value : 0
  return `${day.date} ${range} (${props.timezone}) · ${formatPlayerCount(numericValue)}`
}
</script>

<template>
  <div v-if="!isEmpty" class="heatmap-shell">
    <div class="heatmap-scroll">
      <div class="heatmap-grid" :style="{ gridTemplateColumns: `4.6rem repeat(${hours.length}, minmax(1.2rem, 1fr))` }">
        <div class="heatmap-corner"></div>
        <span v-for="hour in hours" :key="hour" class="heatmap-hour">{{ hour }}</span>

        <template v-for="(day, rowIndex) in days" :key="day.date">
          <div class="heatmap-day">
            <strong>{{ getWeekdayLabel(day.weekday) }}</strong>
            <small>{{ day.date.slice(5) }}</small>
          </div>

          <div
            v-for="(cell, columnIndex) in cells[rowIndex] ?? []"
            :key="`${day.date}-${hours[columnIndex]}`"
            class="heatmap-cell"
            :style="getCellStyle(cell)"
            :title="getTooltip(day, hours[columnIndex], cell)"
            :aria-label="getTooltip(day, hours[columnIndex], cell)"
          ></div>
        </template>
      </div>
    </div>
  </div>
</template>

<style scoped>
.heatmap-shell {
  display: grid;
  gap: 0.75rem;
}

.heatmap-scroll {
  overflow-x: auto;
  padding-bottom: 0.2rem;
}

.heatmap-grid {
  display: grid;
  gap: 0.35rem;
  min-width: max-content;
  align-items: center;
}

.heatmap-corner {
  min-height: 1.1rem;
}

.heatmap-hour,
.heatmap-day small {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.64rem;
  letter-spacing: 0.04em;
}

.heatmap-hour {
  text-align: center;
}

.heatmap-day {
  display: grid;
  gap: 0.05rem;
}

.heatmap-day strong {
  color: var(--text-main);
  font-size: 0.78rem;
  font-weight: 600;
}

.heatmap-cell {
  width: 1.2rem;
  height: 1.2rem;
  border-radius: 6px;
  border: 1px solid rgba(255, 255, 255, 0.2);
  box-shadow: inset 0 0 0 1px rgba(0, 0, 0, 0.08);
}
</style>
