<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  name: string
  value: number | string
  rank?: number
  icon?: string
  formatValue?: (value: number | string) => string
}>()

const displayValue = computed(() => {
  if (props.formatValue) {
    return props.formatValue(props.value)
  }

  return typeof props.value === 'number' ? props.value.toLocaleString() : props.value
})

function handleImageError(event: Event): void {
  const image = event.target as HTMLImageElement | null
  if (image) {
    image.style.display = 'none'
  }
}
</script>

<template>
  <div class="stat-box">
    <div v-if="icon" class="stat-icon-wrap">
      <img :src="icon" class="stat-icon" @error="handleImageError" />
    </div>
    <div class="stat-main">
      <span class="stat-name">{{ name }}</span>
      <span class="stat-value">{{ displayValue }}</span>
    </div>
    <div v-if="rank" class="stat-rank">#{{ rank }}</div>
  </div>
</template>

<style scoped>
.stat-box {
  display: flex;
  gap: 0.8rem;
  align-items: center;
  padding: 0.85rem 0.95rem;
  border-radius: 18px;
  border: 1px solid rgba(255, 248, 234, 0.06);
  background: rgba(255, 248, 234, 0.04);
}

.stat-icon-wrap {
  width: 28px;
  height: 28px;
  display: grid;
  place-items: center;
  flex-shrink: 0;
}

.stat-icon {
  width: 24px;
  height: 24px;
  image-rendering: pixelated;
}

.stat-main {
  display: flex;
  flex-direction: column;
  gap: 0.2rem;
  min-width: 0;
}

.stat-name {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.68rem;
  text-transform: uppercase;
  letter-spacing: 0.1em;
}

.stat-value {
  color: var(--text-strong);
  font-weight: 600;
}

.stat-rank {
  flex-shrink: 0;
  color: var(--primary);
  font-family: var(--mono);
  font-size: 0.76rem;
}
</style>
