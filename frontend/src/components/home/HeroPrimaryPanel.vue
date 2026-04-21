<script setup lang="ts">
import { computed } from 'vue'

import { siteContent } from '@/content/siteContent'

interface Props {
  content: typeof siteContent.home.hero
  onlineCount: number
  updatedLabel: string
}

const props = defineProps<Props>()

const factRows = computed(() => props.content.facts)
</script>

<template>
  <article class="hero-panel glass-card">
    <div class="hero-copy">
      <span class="page-kicker">{{ props.content.eyebrow }}</span>
      <h1 class="hero-title">{{ props.content.title }}</h1>
      <p class="hero-body">{{ props.content.body }}</p>

      <div class="hero-badges">
        <span class="hud-chip">
          <span class="live-dot"></span>
          {{ props.onlineCount }} online
        </span>
        <span class="hud-chip">Updated {{ props.updatedLabel }}</span>
      </div>

      <div class="hero-actions">
        <router-link class="btn-primary" to="/players">
          {{ props.content.primaryCta }}
        </router-link>
        <a class="btn-secondary" href="https://mtn.1919801.xyz/">
          {{ props.content.secondaryCta }}
        </a>
      </div>

      <div class="hero-facts">
        <div v-for="fact in factRows" :key="fact.label" class="hero-fact-card">
          <span class="hud-metric-label">{{ fact.label }}</span>
          <strong class="hud-metric-value">{{ fact.value }}</strong>
        </div>
      </div>
    </div>
  </article>
</template>

<style scoped>
.hero-panel {
  padding: 1.75rem;
}

.hero-copy {
  display: grid;
  gap: 0.95rem;
}

.hero-title {
  font-size: clamp(3.15rem, 7vw, 5.8rem);
  max-width: 9ch;
}

.hero-body {
  max-width: 48ch;
  color: var(--text-muted);
  font-size: 1.02rem;
}

.hero-badges {
  display: flex;
  flex-wrap: wrap;
  gap: 0.6rem;
}

.hero-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
  padding-top: 0.15rem;
}

.hero-facts {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 0.65rem;
  padding-top: 0.2rem;
}

.hero-fact-card {
  display: grid;
  gap: 0.35rem;
  padding: 0.82rem 0.88rem;
  border-radius: 18px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  background: rgba(255, 255, 255, 0.03);
}

.live-dot {
  width: 0.5rem;
  height: 0.5rem;
  border-radius: 50%;
  background: var(--success);
  box-shadow: 0 0 0 5px rgba(131, 211, 167, 0.12);
}

@media (max-width: 960px) {
  .hero-facts {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 640px) {
  .hero-title {
    font-size: 3rem;
  }

  .hero-actions {
    flex-direction: column;
  }
}
</style>
