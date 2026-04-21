<script setup lang="ts">
import { computed } from 'vue'
import { RouterLink } from 'vue-router'

import { siteContent } from '@/content/siteContent'

interface Props {
  // Props for the quick routes section could be added here if needed in the future
}

// props are currently unused as the live signal card was removed
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

}
</style>
