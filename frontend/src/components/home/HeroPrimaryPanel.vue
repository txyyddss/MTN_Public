<script setup lang="ts">
import { computed } from 'vue'

import type { HeroContent } from '@/content/siteContent'

interface Props {
  content: HeroContent
}

const props = defineProps<Props>()

const factRows = computed(() => props.content.facts)
</script>

<template>
  <article class="hero-panel glass-card">
    <div class="hero-copy">
      <span class="page-kicker">{{ props.content.eyebrow }}</span>
      <h1 class="hero-title">{{ props.content.title }}</h1>
      <p class="hero-tagline">{{ props.content.tagline }}</p>
      <p class="hero-accent">{{ props.content.accentLine }}</p>
      <p class="hero-body">{{ props.content.body }}</p>

      <div class="hero-actions">
        <a class="btn-primary" :href="props.content.primaryHref">
          {{ props.content.primaryCta }}
        </a>
        <a class="btn-secondary" :href="props.content.secondaryHref" target="_blank" rel="noopener noreferrer">
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
  padding: clamp(1.35rem, 3vw, 2rem);
  isolation: isolate;
}

.hero-copy {
  display: grid;
  align-content: start;
  gap: 0.95rem;
  min-height: 100%;
}

.hero-title {
  max-width: 100%;
  font-size: 5.35rem;
  letter-spacing: -0.055em;
  word-break: keep-all;
  overflow-wrap: normal;
}

.hero-tagline {
  max-width: 34rem;
  color: var(--text-strong);
  font-size: 1.24rem;
  font-weight: 700;
  line-height: 1.45;
  text-wrap: balance;
}

.hero-accent {
  width: fit-content;
  padding: 0.42rem 0.72rem;
  border: 1px solid rgba(76, 147, 251, 0.26);
  border-radius: 999px;
  background: rgba(76, 147, 251, 0.1);
  color: var(--accent);
  font-family: var(--mono);
  font-size: 0.78rem;
  font-weight: 700;
  letter-spacing: 0.08em;
  box-shadow: 0 12px 28px rgba(76, 147, 251, 0.12);
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
  padding-top: 0.35rem;
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
  position: relative;
  display: grid;
  gap: 0.35rem;
  align-content: start;
  min-height: 100%;
  padding: 0.92rem 0.95rem;
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  background: rgba(255, 255, 255, 0.03);
  overflow: hidden;
}

.hero-fact-card::before {
  content: '';
  position: absolute;
  inset: 0;
  background: linear-gradient(90deg, transparent, rgba(76, 147, 251, 0.1), transparent);
  opacity: 0;
  transform: translateX(-120%);
  transition:
    opacity 0.2s ease,
    transform 0.7s cubic-bezier(0.16, 1, 0.3, 1);
}

.hero-fact-card:hover::before {
  opacity: 1;
  transform: translateX(120%);
}

@media (max-width: 960px) {
  .hero-title {
    font-size: 4.45rem;
  }

  .hero-facts {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (max-width: 640px) {
  .hero-title {
    font-size: 2.75rem;
  }

  .hero-tagline {
    font-size: 1.05rem;
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

@media (max-width: 420px) {
  .hero-title {
    font-size: 2.36rem;
  }
}
</style>
