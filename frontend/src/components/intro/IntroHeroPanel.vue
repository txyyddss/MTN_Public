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
  gap: 1rem;
  margin-bottom: 0;
}

.intro-hero-copy {
  display: grid;
  gap: 0.95rem;
}

.intro-hero h1 {
  font-size: clamp(3rem, 7vw, 5.2rem);
  max-width: 10ch;
}

.intro-hero p,
.hero-aside-card p {
  max-width: 62ch;
  color: var(--text-muted);
}

.hero-aside-card {
  display: grid;
  gap: 0.7rem;
  padding: 1.05rem;
  height: 100%;
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.08);
}

@media (max-width: 900px) {
  .intro-hero {
    grid-template-columns: 1fr;
  }
}
</style>
