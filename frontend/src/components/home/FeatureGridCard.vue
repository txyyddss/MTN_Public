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
    <div class="feature-topline">
      <div class="feature-index">{{ props.feature.icon }}</div>
      <span class="feature-mode">System</span>
    </div>
    <div class="feature-copy">
      <h3 class="feature-title">{{ props.feature.title }}</h3>
      <p class="feature-description">{{ props.feature.description }}</p>
    </div>
  </article>
</template>

<style scoped>
.feature-card {
  display: grid;
  gap: 1.2rem;
  min-height: 220px;
}

.feature-card.compact {
  min-height: auto;
  gap: 0.9rem;
}

.feature-topline {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
}

.feature-index {
  width: 2.7rem;
  height: 2.7rem;
  display: inline-grid;
  place-items: center;
  border-radius: 999px;
  font-family: var(--mono);
  font-size: 0.82rem;
  border: 1px solid rgba(255, 255, 255, 0.12);
  background: rgba(255, 255, 255, 0.05);
}

.feature-mode {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.68rem;
  letter-spacing: 0.14em;
  text-transform: uppercase;
}

.feature-copy {
  display: grid;
  gap: 0.8rem;
}

.feature-title {
  font-size: 1.45rem;
}

.feature-description {
  color: var(--text-muted);
}

.accent-copper {
  border-color: rgba(255, 255, 255, 0.1);
}

.accent-moss {
  border-color: rgba(255, 255, 255, 0.12);
}

.accent-redstone {
  border-color: rgba(76, 147, 251, 0.22);
}
</style>
