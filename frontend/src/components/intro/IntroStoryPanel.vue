<script setup lang="ts">
import { siteContent } from '@/content/siteContent'
</script>

<template>
  <section class="story-grid">
    <article v-for="block in siteContent.intro.story" :key="block.heading" class="glass-card story-card animate-entry">
      <span class="section-kicker">{{ block.heading }}</span>
      <div class="story-copy">
        <p v-for="paragraph in block.paragraphs" :key="paragraph">{{ paragraph }}</p>
      </div>
    </article>
  </section>
</template>

<style scoped>
.story-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 1.15rem;
}

.story-card {
  display: grid;
  gap: 1rem;
}

.story-copy {
  display: grid;
  gap: 1rem;
  color: var(--text-muted);
}

@media (max-width: 900px) {
  .story-grid {
    grid-template-columns: 1fr;
  }
}
</style>
