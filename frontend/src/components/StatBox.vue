<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  name: string
  value: number | string
  rank?: number
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
  background: rgba(255, 255, 255, 0.02);
  border-radius: 8px;
  padding: 10px 14px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-left: 2px solid transparent;
  transition: all 0.2s;
}

.stat-box:hover {
  border-left-color: var(--primary);
  background: rgba(255, 255, 255, 0.05);
}

.stat-main {
  display: flex;
  flex-direction: column;
}

.stat-name {
  font-size: 0.75rem;
  color: var(--text-muted);
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.stat-value {
  font-weight: 800;
  color: #fff;
  font-size: 1.1rem;
  font-family: var(--heading);
}

.stat-rank {
  font-size: 0.8rem;
  font-weight: 800;
  color: #fcd34d;
  background: rgba(245, 158, 11, 0.1);
  padding: 2px 8px;
  border-radius: 4px;
}
</style>
