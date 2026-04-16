<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  name: string
  value: number | string
  isEnlarged: boolean
  rank?: number
  icon?: string
  formatValue?: (val: any) => string
}>()

const emit = defineEmits(['toggle'])

const toggleEnlarge = (e: Event) => {
  e.stopPropagation()
  emit('toggle')
}

const displayValue = computed(() => {
  if (props.formatValue) return props.formatValue(props.value)
  return typeof props.value === 'number' ? props.value.toLocaleString() : props.value
})

const formatName = (name: string) => {
  return name.replace('minecraft:', '').replace(/_/g, ' ')
}
</script>

<template>
  <div :class="['stat-box', { enlarged: isEnlarged }]" @click="toggleEnlarge">
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
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
  position: relative;
  z-index: 1;
}

.stat-box:hover {
  border-color: var(--primary);
  background: rgba(255, 255, 255, 0.08);
}

.stat-box.enlarged {
  transform: scale(1.1);
  z-index: 10;
  background: rgba(30, 41, 59, 0.95);
  border-color: var(--primary);
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.4), 0 8px 10px -6px rgba(0, 0, 0, 0.4);
  backdrop-filter: blur(8px);
}

.stat-box.enlarged .stat-name {
  white-space: normal;
  overflow: visible;
  text-overflow: clip;
}

.stat-icon-wrap {
    width: 24px;
    height: 24px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
    transition: transform 0.3s ease;
}

.stat-box.enlarged .stat-icon-wrap {
  transform: scale(1.2);
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
  transition: all 0.3s ease;
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
