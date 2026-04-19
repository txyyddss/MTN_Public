<script setup lang="ts">
import { siteContent } from '@/content/siteContent'

interface Props {
  content: typeof siteContent.home.hero
  onlineCount: number
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
        {{ props.onlineCount }} online
      </span>
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
  gap: 1rem;
  padding: 0.75rem 0;
}

.hero-title {
  font-size: clamp(3.2rem, 7vw, 6rem);
  max-width: 11ch;
}

.hero-body {
  max-width: 54ch;
  color: var(--text-muted);
  font-size: 0.98rem;
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
  gap: 0.85rem;
  padding-top: 0.25rem;
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
