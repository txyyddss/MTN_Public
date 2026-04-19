<script setup lang="ts">
import { siteContent } from '@/content/siteContent'
</script>

<template>
  <section class="intro-hero glass-card animate-entry">
    <div class="intro-hero-copy">
      <span class="section-kicker">{{ siteContent.intro.hero.eyebrow }}</span>
      <h1>{{ siteContent.intro.hero.title }}</h1>
      <p>{{ siteContent.intro.hero.body }}</p>
    </div>
    <div class="intro-hero-aside">
      <div class="hero-aside-card">
        <span class="hud-kicker">Mission log</span>
        <p>Private origin, public world, stability-first operations.</p>
      </div>
    </div>
  </section>
</template>

<style scoped>
.intro-hero {
  display: grid;
  grid-template-columns: minmax(0, 1.35fr) minmax(240px, 0.7fr);
  gap: 1.25rem;
  margin-bottom: 0;
}

.intro-hero-copy {
  display: grid;
  gap: 1rem;
}

.intro-hero h1 {
  font-size: clamp(3rem, 7vw, 5.8rem);
  max-width: 12ch;
}

.intro-hero p,
.hero-aside-card p {
  max-width: 62ch;
  color: var(--text-muted);
}

.hero-aside-card {
  display: grid;
  gap: 0.75rem;
  padding: 1rem;
  height: 100%;
  border-radius: 24px;
  background: rgba(7, 16, 30, 0.88);
  border: 1px solid rgba(121, 183, 255, 0.14);
}

@media (max-width: 900px) {
  .intro-hero {
    grid-template-columns: 1fr;
  }
}
</style>
