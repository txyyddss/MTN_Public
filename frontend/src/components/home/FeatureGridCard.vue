<script setup lang="ts">
import { computed } from 'vue'

import { siteContent } from '@/content/siteContent'

type HomeFeature = (typeof siteContent.home.features)[number]

interface Props {
  feature: HomeFeature
  index: number
  compact?: boolean
}

const props = defineProps<Props>()

const cardClasses = computed(() => [
  'feature-card',
  'glass-card',
  `accent-${props.feature.accent}`,
  { compact: props.compact },
  'animate-entry'
])

const cardStyle = computed(() => ({
  animationDelay: `${props.index * 0.08}s`
}))
</script>

<template>
  <article :class="cardClasses" :style="cardStyle">
    <div class="feature-index">{{ props.feature.icon }}</div>
    <div class="feature-copy">
      <h3 class="feature-title">{{ props.feature.title }}</h3>
      <p class="feature-description">{{ props.feature.description }}</p>
    </div>
  </article>
</template>

<style scoped>
.feature-card {
  display: grid;
  gap: 1.5rem;
  min-height: 250px;
}

.feature-card.compact {
  min-height: auto;
  gap: 1rem;
}

.feature-index {
  width: 3rem;
  height: 3rem;
  display: inline-grid;
  place-items: center;
  border-radius: 999px;
  font-family: var(--mono);
  font-size: 0.88rem;
  border: 1px solid rgba(255, 255, 255, 0.14);
  background: rgba(255, 255, 255, 0.04);
}

.feature-copy {
  display: grid;
  gap: 0.8rem;
}

.feature-title {
  font-size: 1.5rem;
}

.feature-description {
  color: var(--text-muted);
}

.accent-copper {
  border-color: rgba(43, 98, 255, 0.34);
}

.accent-moss {
  border-color: rgba(30, 84, 219, 0.32);
}

.accent-redstone {
  border-color: rgba(98, 171, 255, 0.3);
}
</style>
