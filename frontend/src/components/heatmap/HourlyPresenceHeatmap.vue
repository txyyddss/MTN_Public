<script setup lang="ts">
import { computed } from 'vue'

import type { HeatmapDay } from '@/types/api'

const props = defineProps<{
  days: HeatmapDay[]
  hours: number[]
  cells: Array<Array<number | boolean>>
  timezone: string
  variant: 'server' | 'player'
  weeklyMaxPlayers?: number
}>()

const isEmpty = computed(() => props.days.length === 0 || props.hours.length === 0 || props.cells.length === 0)

function getCellStyle(value: number | boolean): { backgroundColor: string; borderColor: string } {
  if (props.variant === 'player') {
    return value
      ? { backgroundColor: 'rgb(76, 147, 251)', borderColor: 'rgba(76, 147, 251, 0.68)' }
      : { backgroundColor: '#ffffff', borderColor: 'rgba(255, 255, 255, 0.5)' }
  }

  const numericValue = typeof value === 'number' ? value : 0
  if (numericValue <= 0 || !props.weeklyMaxPlayers) {
    return { backgroundColor: '#ffffff', borderColor: 'rgba(255, 255, 255, 0.5)' }
  }

  const ratio = Math.min(numericValue / props.weeklyMaxPlayers, 1)
  const red = Math.round(255 + (76 - 255) * ratio)
  const green = Math.round(255 + (147 - 255) * ratio)
  const blue = Math.round(255 + (251 - 255) * ratio)

  return {
    backgroundColor: `rgb(${red}, ${green}, ${blue})`,
    borderColor: `rgba(${red}, ${green}, ${blue}, 0.85)`
  }
}

function getTooltip(day: HeatmapDay, hourLabel: number, value: number | boolean): string {
  const hourStart = String(hourLabel - 1).padStart(2, '0')
  const hourEnd = String(hourLabel % 24).padStart(2, '0')
  const range = `${hourStart}:00-${hourEnd}:00`

  if (props.variant === 'player') {
    return `${day.date} ${range} (${props.timezone}) · ${value ? 'Online' : 'Offline'}`
  }

  const numericValue = typeof value === 'number' ? value : 0
  const playerLabel = numericValue === 1 ? 'player' : 'players'
  return `${day.date} ${range} (${props.timezone}) · ${numericValue} ${playerLabel}`
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
            <strong>{{ day.weekday }}</strong>
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
