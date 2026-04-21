<script setup lang="ts">
import { computed } from 'vue'

import { siteContent } from '@/content/siteContent'

interface Props {
  content: typeof siteContent.home.hero
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
  display: grid;
  min-height: 100%;
  padding: 1.7rem;
}

.hero-copy {
  display: grid;
  align-content: start;
  gap: 1rem;
  min-height: 100%;
}

.hero-title {
  font-size: clamp(3.3rem, 8vw, 6.35rem);
  max-width: 10ch;
}

.hero-body {
  max-width: 52ch;
  color: var(--text-muted);
  font-size: 1.06rem;
}

.hero-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
  padding-top: 0.25rem;
}

.hero-actions :deep(.btn-primary),
.hero-actions :deep(.btn-secondary) {
  min-width: 11.5rem;
}

.hero-facts {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(165px, 1fr));
  gap: 0.7rem;
  margin-top: auto;
  padding-top: 0.4rem;
}

.hero-fact-card {
  display: grid;
  gap: 0.35rem;
  align-content: start;
  min-height: 100%;
  padding: 0.92rem 0.95rem;
  border-radius: 18px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  background: rgba(255, 255, 255, 0.03);
}

@media (max-width: 960px) {
  .hero-facts {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (max-width: 640px) {
  .hero-title {
    font-size: 3rem;
  }

  .hero-actions {
    flex-direction: column;
  }

  .hero-actions :deep(.btn-primary),
  .hero-actions :deep(.btn-secondary) {
    width: 100%;
    min-width: 0;
  }

  .hero-facts {
    grid-template-columns: 1fr;
  }
}
</style>
