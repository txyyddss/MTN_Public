<script setup lang="ts">
import { siteContent } from '@/content/siteContent'

type HeroFact = (typeof siteContent.home.hero.facts)[number]

interface Props {
  facts: readonly HeroFact[]
  imageAlt: string
  imageSrc: string
}

const props = defineProps<Props>()
</script>

<template>
  <aside class="hero-atlas glass-card">
    <div class="atlas-visual">
      <img class="atlas-image" :src="props.imageSrc" :alt="props.imageAlt" />
    </div>
    <div class="atlas-caption">
      <p class="atlas-kicker">Field notes</p>
      <div class="atlas-facts">
        <div v-for="fact in props.facts" :key="fact.label" class="atlas-fact">
          <span class="atlas-fact-label">{{ fact.label }}</span>
          <strong class="atlas-fact-value">{{ fact.value }}</strong>
        </div>
      </div>
    </div>
  </aside>
</template>

<style scoped>
.hero-atlas {
  display: grid;
  gap: 1.2rem;
  min-height: 100%;
}

.atlas-visual {
  overflow: hidden;
  border-radius: calc(var(--radius-lg) - 12px);
  border: 1px solid rgba(255, 248, 234, 0.08);
  min-height: 320px;
}

.atlas-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.atlas-caption {
  display: grid;
  gap: 1rem;
}

.atlas-kicker {
  color: var(--primary);
  font-family: var(--mono);
  font-size: 0.72rem;
  letter-spacing: 0.18em;
  text-transform: uppercase;
}

.atlas-facts {
  display: grid;
  gap: 0.75rem;
}

.atlas-fact {
  display: flex;
  justify-content: space-between;
  gap: 1rem;
  padding: 0.95rem 1rem;
  border-radius: 18px;
  background: rgba(255, 248, 234, 0.04);
  border: 1px solid rgba(255, 248, 234, 0.06);
}

.atlas-fact-label {
  color: var(--text-dim);
  font-size: 0.86rem;
  text-transform: uppercase;
  letter-spacing: 0.08em;
}

.atlas-fact-value {
  color: var(--text-strong);
  font-weight: 600;
  text-align: right;
}
</style>
