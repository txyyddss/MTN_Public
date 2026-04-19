<script setup lang="ts">
import { siteContent } from '@/content/siteContent'
</script>

<template>
  <div class="stub-view container page-shell">
    <div class="glass-card placeholder-card animate-entry">
      <span class="section-kicker">{{ siteContent.teamPlaceholder.eyebrow }}</span>
      <h1>{{ siteContent.teamPlaceholder.title }}</h1>
      <p>{{ siteContent.teamPlaceholder.body }}</p>
      <div class="placeholder-actions">
        <RouterLink class="btn-primary" to="/server-intro">{{ siteContent.teamPlaceholder.primaryCta }}</RouterLink>
        <RouterLink class="btn-secondary" to="/server-intro#community-panel">
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
  width: min(760px, 100%);
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
