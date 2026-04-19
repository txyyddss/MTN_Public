<script setup lang="ts">
import { siteContent } from '@/content/siteContent'
</script>

<template>
  <section class="intro-hero glass-card animate-entry">
    <span class="section-kicker">{{ siteContent.intro.hero.eyebrow }}</span>
    <h1>{{ siteContent.intro.hero.title }}</h1>
    <p>{{ siteContent.intro.hero.body }}</p>
  </section>
</template>

<style scoped>
.intro-hero {
  display: grid;
  gap: 1rem;
  margin-bottom: 1.25rem;
}

.intro-hero h1 {
  font-size: clamp(3rem, 7vw, 5.8rem);
  max-width: 12ch;
}

.intro-hero p {
  max-width: 62ch;
  color: var(--text-muted);
}
</style>
