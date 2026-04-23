<script setup lang="ts">
import { useSiteContent } from '@/content/siteContent'

const siteContent = useSiteContent()
</script>

<template>
  <div class="stub-view container page-shell">
    <div class="glass-card placeholder-card animate-entry">
      <div class="placeholder-copy">
        <span class="section-kicker">{{ siteContent.teamPlaceholder.eyebrow }}</span>
        <h1>{{ siteContent.teamPlaceholder.title }}</h1>
        <p>{{ siteContent.teamPlaceholder.body }}</p>
      </div>

      <div class="placeholder-actions">
        <a class="btn-primary" href="https://mtn.1919801.xyz/">{{ siteContent.teamPlaceholder.primaryCta }}</a>
        <RouterLink class="btn-secondary" to="/">
          {{ siteContent.teamPlaceholder.secondaryCta }}
        </RouterLink>
      </div>

      <small>{{ siteContent.teamPlaceholder.note }}</small>
    </div>
  </div>
</template>

<style scoped>
.stub-view {
  display: grid;
  place-items: center;
}

.placeholder-card {
  width: min(880px, 100%);
  display: grid;
  gap: 1.2rem;
}

.placeholder-copy {
  display: grid;
  gap: 1rem;
}

.placeholder-card h1 {
  font-size: clamp(2.6rem, 5vw, 4.4rem);
  max-width: 12ch;
}

.placeholder-card p,
.placeholder-card small {
  color: var(--text-muted);
}

.placeholder-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
}
</style>
