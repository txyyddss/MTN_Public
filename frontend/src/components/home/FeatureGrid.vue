<script setup lang="ts">
import FeatureGridCard from '@/components/home/FeatureGridCard.vue'
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
        <FeatureGridCard
          v-for="(feature, index) in siteContent.home.features"
          :key="feature.title"
          :feature="feature"
          :index="index"
        />
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
