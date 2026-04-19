<script setup lang="ts">
import { siteContent } from '@/content/siteContent'
</script>

<template>
  <section class="features-section">
    <div class="container">
      <div class="section-head animate-entry">
        <span class="section-kicker">Operating principles</span>
        <h2 class="section-title">What the project actually optimizes for.</h2>
        <p class="section-copy">
          The frontend no longer sells an abstract luxury server. It describes the verified shape of this one.
        </p>
      </div>

      <div class="features-grid">
        <article
          v-for="(feature, index) in siteContent.home.features"
          :key="feature.title"
          :class="['feature-card', 'glass-card', `accent-${feature.accent}`, 'animate-entry']"
          :style="{ animationDelay: `${index * 0.08}s` }"
        >
          <div class="feature-index">{{ feature.icon }}</div>
          <div class="feature-copy">
            <h3>{{ feature.title }}</h3>
            <p>{{ feature.description }}</p>
          </div>
        </article>
      </div>
    </div>
  </section>
</template>

<style scoped>
.features-section {
  padding: 2rem 0 5rem;
}

.section-head {
  display: grid;
  gap: 1rem;
  margin-bottom: 2rem;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 1.25rem;
}

.feature-card {
  display: grid;
  gap: 1.5rem;
  min-height: 250px;
}

.feature-index {
  width: 3rem;
  height: 3rem;
  display: inline-grid;
  place-items: center;
  border-radius: 999px;
  font-family: var(--mono);
  font-size: 0.88rem;
  border: 1px solid rgba(255, 248, 234, 0.14);
  background: rgba(255, 248, 234, 0.04);
}

.feature-copy {
  display: grid;
  gap: 0.8rem;
}

.feature-copy h3 {
  font-size: 1.8rem;
}

.feature-copy p {
  color: var(--text-muted);
}

.accent-copper {
  border-color: rgba(196, 122, 66, 0.28);
}

.accent-moss {
  border-color: rgba(120, 130, 91, 0.28);
}

.accent-redstone {
  border-color: rgba(186, 77, 55, 0.28);
}

@media (max-width: 980px) {
  .features-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (max-width: 680px) {
  .features-grid {
    grid-template-columns: 1fr;
  }
}
</style>
