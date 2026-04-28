<script setup lang="ts">
import { useSiteContent } from '@/content/siteContent'

const siteContent = useSiteContent()
</script>

<template>
  <div class="stub-view container page-shell route-page-shell">
    <div class="glass-card placeholder-card animate-entry">
      <span class="placeholder-mark" aria-hidden="true">MTN</span>
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
  border-color: rgba(76, 147, 251, 0.22);
}

.placeholder-mark {
  color: rgba(76, 147, 251, 0.9);
  font-family: var(--display);
  font-size: clamp(3.5rem, 12vw, 7.5rem);
  font-weight: 900;
  letter-spacing: 0.22em;
  line-height: 0.75;
  text-indent: 0.22em;
  text-shadow: 0 22px 50px rgba(76, 147, 251, 0.22);
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
