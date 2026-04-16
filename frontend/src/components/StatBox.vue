<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  name: string
  value: number | string
  rank?: number
  icon?: string
  formatValue?: (val: any) => string
}>()

const displayValue = computed(() => {
  if (props.formatValue) return props.formatValue(props.value)
  return props.value.toLocaleString()
})

const formatName = (name: string) => {
  return name.replace('minecraft:', '').replace(/_/g, ' ')
}
</script>

<template>
  <div class="stat-box">
    <div class="stat-icon-wrap" v-if="icon">
        <img :src="icon" class="stat-icon" @error="(e: any) => e.target.style.display='none'" />
    </div>
    <div class="stat-main">
      <span class="stat-name">{{ formatName(name) }}</span>
      <span class="stat-value">{{ displayValue }}</span>
    </div>
    <div class="stat-rank" v-if="rank">
      #{{ rank }}
    </div>
  </div>
</template>

<style scoped>
.stat-box {
  background: rgba(255, 255, 255, 0.03);
  border-radius: 6px;
  padding: 6px 10px;
  display: flex;
  gap: 10px;
  align-items: center;
  border: 1px solid rgba(255,255,255,0.05);
  transition: all 0.2s;
}

.stat-box:hover {
  border-color: var(--primary);
  background: rgba(255, 255, 255, 0.08);
}

.stat-icon-wrap {
    width: 24px;
    height: 24px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
}

.stat-icon {
    width: 20px;
    height: 20px;
    image-rendering: pixelated;
}

.stat-main {
  display: flex;
  flex-direction: column;
  flex: 1;
  min-width: 0;
}

.stat-name {
  font-size: 0.65rem;
  color: var(--text-muted);
  text-transform: uppercase;
  letter-spacing: 0.5px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.stat-value {
  font-weight: 800;
  color: #fff;
  font-size: 0.95rem;
  font-family: var(--heading);
  line-height: 1;
}

.stat-rank {
  font-size: 0.65rem;
  font-weight: 800;
  color: #fcd34d;
  background: rgba(245, 158, 11, 0.1);
  padding: 1px 4px;
  border-radius: 4px;
  white-space: nowrap;
}
</style>
