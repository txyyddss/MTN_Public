<script setup lang="ts">
import { computed } from 'vue'

import { type FeatureCardContent, useSiteContent } from '@/content/siteContent'

interface Props {
  feature: FeatureCardContent
  index: number
  compact?: boolean
}

const props = defineProps<Props>()
const siteContent = useSiteContent()

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
      <span class="feature-mode">{{ siteContent.home.featureMode }}</span>
    </div>
    <div class="feature-copy">
      <h3 class="feature-title">{{ props.feature.title }}</h3>
      <p class="feature-description">{{ props.feature.description }}</p>
    </div>
  </article>
</template>

<style scoped>
.feature-card {
  position: relative;
  display: grid;
  gap: 0.95rem;
  min-height: 184px;
  padding: 1rem;
  isolation: isolate;
}

.feature-card::before {
  content: '';
  position: absolute;
  inset: 0;
  z-index: -1;
  background:
    linear-gradient(135deg, rgba(76, 147, 251, 0.12), transparent 48%),
    linear-gradient(180deg, rgba(255, 255, 255, 0.04), transparent 42%);
  opacity: 0;
  transform: translateY(10px);
  transition:
    opacity var(--transition-panel),
    transform var(--transition-panel);
}

.feature-card::after {
  content: '';
  position: absolute;
  inset: 0;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.06), transparent);
  opacity: 0;
  transform: translateX(-120%);
  transition:
    opacity 0.18s ease,
    transform 0.72s cubic-bezier(0.16, 1, 0.3, 1);
  pointer-events: none;
}

.feature-card:hover {
  transform: translateY(-5px);
  border-color: rgba(76, 147, 251, 0.28);
  box-shadow: 0 34px 68px rgba(0, 0, 0, 0.42);
}

.feature-card:hover::before {
  opacity: 1;
  transform: translateY(0);
}

.feature-card:hover::after {
  opacity: 1;
  transform: translateX(120%);
}

.feature-card.compact {
  min-height: 166px;
  gap: 0.78rem;
  padding: 0.95rem;
}

.feature-topline {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
}

.feature-index {
  width: 2.35rem;
  height: 2.35rem;
  display: inline-grid;
  place-items: center;
  border-radius: 999px;
  font-family: var(--mono);
  font-size: 0.76rem;
  border: 1px solid rgba(255, 255, 255, 0.12);
  background:
    linear-gradient(135deg, rgba(76, 147, 251, 0.18), rgba(255, 255, 255, 0.04)),
    rgba(255, 255, 255, 0.05);
  color: var(--text-strong);
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.06);
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
  gap: 0.58rem;
}

.feature-title {
  font-size: 1.22rem;
  letter-spacing: -0.025em;
}

.feature-description {
  color: var(--text-muted);
  font-size: 0.92rem;
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
