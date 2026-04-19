<script setup lang="ts">
import { computed } from 'vue'
import { storeToRefs } from 'pinia'

import heroArt from '@/assets/hero.png'
import { siteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'

const statusStore = useServerStatusStore()
const { status } = storeToRefs(statusStore)

const onlineCount = computed(() => status.value?.java?.players ?? 0)
const javaVersion = computed(() => status.value?.java?.version || 'Fetching live version')
const updatedLabel = computed(() => {
  if (!status.value?.updated) {
    return 'Waiting for telemetry'
  }

  return new Date(status.value.updated).toLocaleTimeString()
})
</script>

<template>
  <section class="hero">
    <div class="container hero-grid">
      <div class="hero-copy animate-entry">
        <span class="page-kicker">{{ siteContent.home.hero.eyebrow }}</span>
        <h1 class="hero-title">{{ siteContent.home.hero.title }}</h1>
        <p class="hero-body">{{ siteContent.home.hero.body }}</p>

        <div class="hero-badges">
          <span class="badge-pill">
            <span class="live-dot"></span>
            {{ onlineCount }} players visible right now
          </span>
          <span class="badge-pill">Java version: {{ javaVersion }}</span>
          <span class="badge-pill">Updated {{ updatedLabel }}</span>
        </div>

        <div class="hero-actions">
          <router-link class="btn-primary" to="/players">
            {{ siteContent.home.hero.primaryCta }}
          </router-link>
          <router-link class="btn-secondary" to="/server-intro">
            {{ siteContent.home.hero.secondaryCta }}
          </router-link>
        </div>
      </div>

      <aside class="hero-atlas glass-card animate-entry delay-200">
        <div class="atlas-visual">
          <img :src="heroArt" alt="MT Network world capture" />
        </div>
        <div class="atlas-caption">
          <p class="atlas-kicker">Field notes</p>
          <div class="atlas-facts">
            <div v-for="fact in siteContent.home.hero.facts" :key="fact.label" class="atlas-fact">
              <span>{{ fact.label }}</span>
              <strong>{{ fact.value }}</strong>
            </div>
          </div>
        </div>
      </aside>
    </div>
  </section>
</template>

<style scoped>
.hero {
  padding: 4rem 0 3rem;
}

.hero-grid {
  display: grid;
  grid-template-columns: 1.2fr 0.9fr;
  gap: 2rem;
  align-items: stretch;
}

.hero-copy {
  display: grid;
  align-content: center;
  gap: 1.35rem;
  padding: 1.5rem 0;
}

.hero-title {
  font-size: clamp(3.8rem, 8vw, 7rem);
  max-width: 12ch;
}

.hero-body {
  max-width: 60ch;
  color: var(--text-muted);
  font-size: 1.08rem;
}

.hero-badges {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
}

.live-dot {
  width: 0.55rem;
  height: 0.55rem;
  border-radius: 50%;
  background: var(--success);
  box-shadow: 0 0 0 6px rgba(126, 165, 103, 0.14);
}

.hero-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
  padding-top: 0.4rem;
}

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

.atlas-visual img {
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

.atlas-fact span {
  color: var(--text-dim);
  font-size: 0.86rem;
  text-transform: uppercase;
  letter-spacing: 0.08em;
}

.atlas-fact strong {
  color: var(--text-strong);
  font-weight: 600;
  text-align: right;
}

@media (max-width: 960px) {
  .hero-grid {
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
