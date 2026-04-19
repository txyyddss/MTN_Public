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
    <div class="hero-panel-grid">
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
          <router-link class="btn-secondary" to="/server-intro">
            {{ props.content.secondaryCta }}
          </router-link>
        </div>
      </div>

      <div class="hero-aside">
        <div class="hero-console">
          <div class="hero-console-head">
            <span class="hud-kicker">Command board</span>
            <span class="console-status">HUD active</span>
          </div>

          <div class="hud-metric-grid">
            <div v-for="fact in factRows" :key="fact.label" class="hud-metric-card">
              <span class="hud-metric-label">{{ fact.label }}</span>
              <strong class="hud-metric-value">{{ fact.value }}</strong>
            </div>
          </div>
        </div>
      </div>
    </div>
  </article>
</template>

<style scoped>
.hero-panel {
  padding: 1.8rem;
}

.hero-panel-grid {
  display: grid;
  grid-template-columns: minmax(0, 1.4fr) minmax(300px, 0.9fr);
  gap: 1.5rem;
  align-items: stretch;
}

.hero-copy {
  display: grid;
  align-content: center;
  gap: 1rem;
}

.hero-title {
  font-size: clamp(3.2rem, 7vw, 6rem);
  max-width: 11ch;
}

.hero-body {
  max-width: 56ch;
  color: var(--text-muted);
  font-size: 1rem;
}

.hero-badges {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
}

.hero-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 0.85rem;
  padding-top: 0.25rem;
}

.hero-aside {
  display: grid;
}

.hero-console {
  display: grid;
  gap: 1rem;
  padding: 1rem;
  border-radius: 24px;
  background:
    linear-gradient(180deg, rgba(255, 255, 255, 0.04), transparent 28%),
    rgba(6, 14, 26, 0.9);
  border: 1px solid rgba(121, 183, 255, 0.16);
}

.hero-console-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
}

.console-status {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.74rem;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.live-dot {
  width: 0.55rem;
  height: 0.55rem;
  border-radius: 50%;
  background: var(--success);
  box-shadow: 0 0 0 6px rgba(93, 226, 164, 0.14);
}

@media (max-width: 960px) {
  .hero-panel-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 640px) {
  .hero-title {
    font-size: 3.2rem;
  }

  .hero-actions {
    flex-direction: column;
  }
}
</style>
