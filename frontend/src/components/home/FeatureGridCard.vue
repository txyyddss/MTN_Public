<script setup lang="ts">
import { computed } from 'vue'

import { siteContent } from '@/content/siteContent'

type HomeFeature = (typeof siteContent.home.features)[number]

interface Props {
  feature: HomeFeature
  index: number
}

const props = defineProps<Props>()

const cardClasses = computed(() => [
  'feature-card',
  'glass-card',
  `accent-${props.feature.accent}`,
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

.feature-index {
  width: 3rem;
  height: 3rem;
  display: inline-grid;
  place-items: center;
  border-radius: 999px;
  font-family: var(--mono);
  font-size: 0.88rem;
  border: 1px solid rgba(255, 248, 234, 0.14);
  background: rgba(255, 248, 234, 0.04);
}

.feature-copy {
  display: grid;
  gap: 0.8rem;
}

.feature-title {
  font-size: 1.8rem;
}

.feature-description {
  color: var(--text-muted);
}

.accent-copper {
  border-color: rgba(91, 113, 246, 0.34);
}

.accent-moss {
  border-color: rgba(59, 130, 246, 0.32);
}

.accent-redstone {
  border-color: rgba(125, 211, 252, 0.3);
}
</style>
