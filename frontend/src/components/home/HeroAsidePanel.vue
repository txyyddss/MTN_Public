<script setup lang="ts">
import { computed } from 'vue'
import { RouterLink } from 'vue-router'

import { siteContent } from '@/content/siteContent'

interface Props {
  isLive: boolean
  onlineCount: number
  updatedLabel: string
}

const props = defineProps<Props>()

const linkDescriptions: Record<string, string> = {
  Players: 'Browse the current archive and jump into individual records.',
  'Core Members': 'See who runs operations and keeps the server moving.',
  Gallery: 'Open the visual world log.',
  Wiki: 'Read the server handbook and gameplay details.'
}

const quickLinks = computed(() =>
  siteContent.app.nav
    .filter((item) => item.label !== 'Home')
    .map((item, index) => ({
      ...item,
      index: `0${index + 1}`.slice(-2),
      description: linkDescriptions[item.label] ?? 'Open this route.'
    }))
)
</script>

<template>
  <aside class="hero-aside">
    <article class="glass-card hero-aside-card animate-entry delay-100">
      <div class="hero-aside-head">
        <div class="hero-aside-copy">
          <span class="section-kicker">Live surface</span>
          <h2 class="hero-aside-title">Current server signal.</h2>
        </div>
        <span :class="['hero-state-pill', { live: props.isLive }]">
          {{ props.isLive ? 'Operational' : 'Standby' }}
        </span>
      </div>

      <div class="hero-aside-metrics">
        <div class="hero-metric-card">
          <span class="hud-metric-label">Players online</span>
          <strong class="hero-metric-value">{{ props.onlineCount }}</strong>
        </div>
        <div class="hero-metric-card">
          <span class="hud-metric-label">Telemetry</span>
          <strong class="hero-metric-value">{{ props.updatedLabel }}</strong>
        </div>
      </div>

      <p class="hero-aside-note">
        Live status, connection addresses, and public records are surfaced directly below the fold.
      </p>
    </article>

    <article class="glass-card hero-aside-card animate-entry delay-200">
      <div class="hero-aside-copy">
        <span class="section-kicker">Quick routes</span>
        <h2 class="hero-aside-title">Move through the archive.</h2>
      </div>

      <div class="hero-link-list">
        <template v-for="link in quickLinks" :key="link.label">
          <a v-if="link.external" :href="link.to" class="hero-link-row">
            <span class="hero-link-index">{{ link.index }}</span>
            <span class="hero-link-copy">
              <strong>{{ link.label }}</strong>
              <span>{{ link.description }}</span>
            </span>
          </a>
          <RouterLink v-else :to="link.to" class="hero-link-row">
            <span class="hero-link-index">{{ link.index }}</span>
            <span class="hero-link-copy">
              <strong>{{ link.label }}</strong>
              <span>{{ link.description }}</span>
            </span>
          </RouterLink>
        </template>
      </div>
    </article>
  </aside>
</template>

<style scoped>
.hero-aside {
  display: grid;
  gap: 1rem;
}

.hero-aside-card {
  display: grid;
  gap: 1rem;
  min-height: 0;
}

.hero-aside-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 1rem;
}

.hero-aside-copy {
  display: grid;
  gap: 0.55rem;
}

.hero-aside-title {
  font-size: clamp(1.5rem, 3vw, 2.2rem);
  max-width: 12ch;
}

.hero-state-pill {
  display: inline-flex;
  align-items: center;
  min-height: 2.1rem;
  padding: 0.42rem 0.72rem;
  border-radius: 999px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.68rem;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.hero-state-pill.live {
  color: var(--success);
}

.hero-aside-metrics {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 0.7rem;
}

.hero-metric-card {
  display: grid;
  gap: 0.3rem;
  padding: 0.95rem;
  border-radius: 18px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  background: rgba(255, 255, 255, 0.028);
}

.hero-metric-value {
  color: var(--text-strong);
  font-size: 1rem;
  font-weight: 600;
}

.hero-aside-note {
  color: var(--text-muted);
}

.hero-link-list {
  display: grid;
  gap: 0.65rem;
}

.hero-link-row {
  display: grid;
  grid-template-columns: auto minmax(0, 1fr);
  gap: 0.85rem;
  align-items: flex-start;
  padding: 0.9rem 0.95rem;
  border-radius: 18px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  background: rgba(255, 255, 255, 0.028);
  transition:
    transform var(--transition-fast),
    border-color var(--transition-fast),
    background var(--transition-fast);
}

.hero-link-row:hover,
.hero-link-row.router-link-active {
  transform: translateY(-1px);
  border-color: rgba(76, 147, 251, 0.24);
  background: rgba(255, 255, 255, 0.04);
}

.hero-link-index {
  display: inline-grid;
  place-items: center;
  width: 2.1rem;
  min-height: 2.1rem;
  border-radius: 999px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: var(--accent);
  font-family: var(--mono);
  font-size: 0.7rem;
}

.hero-link-copy {
  display: grid;
  gap: 0.28rem;
}

.hero-link-copy strong {
  color: var(--text-strong);
  font-size: 0.98rem;
}

.hero-link-copy span {
  color: var(--text-muted);
  font-size: 0.9rem;
}

@media (max-width: 640px) {
  .hero-aside-metrics {
    grid-template-columns: 1fr;
  }

  .hero-aside-head {
    flex-direction: column;
  }
}
</style>
