<script setup lang="ts">
import { siteContent } from '@/content/siteContent'

interface Props {
  content: typeof siteContent.home.hero
  onlineCount: number
  javaVersion: string
  updatedLabel: string
}

const props = defineProps<Props>()
</script>

<template>
  <div class="hero-copy">
    <span class="page-kicker">{{ props.content.eyebrow }}</span>
    <h1 class="hero-title">{{ props.content.title }}</h1>
    <p class="hero-body">{{ props.content.body }}</p>

    <div class="hero-badges">
      <span class="badge-pill">
        <span class="live-dot"></span>
        {{ props.onlineCount }} players visible right now
      </span>
      <span class="badge-pill">Java version: {{ props.javaVersion }}</span>
      <span class="badge-pill">Updated {{ props.updatedLabel }}</span>
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
</template>

<style scoped>
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

@media (max-width: 640px) {
  .hero-title {
    font-size: 3.2rem;
  }

  .hero-actions {
    flex-direction: column;
  }
}
</style>
