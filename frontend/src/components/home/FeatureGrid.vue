<script setup lang="ts">
import { computed } from 'vue'

import FeatureGridCard from '@/components/home/FeatureGridCard.vue'
import { useAutoRotatingIndex } from '@/composables/useAutoRotatingIndex'
import { useMediaQuery } from '@/composables/useMediaQuery'
import { siteContent } from '@/content/siteContent'

const features = siteContent.home.features
const { matches: isPhone } = useMediaQuery('(max-width: 680px)')
const { currentIndex, next, previous } = useAutoRotatingIndex(computed(() => features.length), isPhone, 3600)
const visibleFeature = computed(() => features[currentIndex.value] ?? features[0])
</script>

<template>
  <section class="features-section">
    <div class="container">
      <div class="section-head animate-entry">
        <span class="section-kicker">Operating principles</span>
        <h2 class="section-title">What this world is built for.</h2>
        <p class="section-copy">Fair rules, stable survival, shared history made visible.</p>
      </div>

      <div v-if="!isPhone" class="features-grid">
        <FeatureGridCard
          v-for="(feature, index) in features"
          :key="feature.title"
          :feature="feature"
          :index="index"
        />
      </div>

      <div v-else class="feature-carousel glass-card animate-entry delay-100">
        <Transition name="carousel-fade" mode="out-in">
          <FeatureGridCard :key="visibleFeature.title" :feature="visibleFeature" :index="currentIndex" compact />
        </Transition>

        <div class="carousel-controls">
          <button type="button" class="carousel-btn" @click="previous">Prev</button>
          <span class="carousel-status">{{ currentIndex + 1 }} / {{ features.length }}</span>
          <button type="button" class="carousel-btn" @click="next">Next</button>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.features-section {
  padding: 1rem 0 4.5rem;
}

.section-head {
  display: grid;
  gap: 1rem;
  margin-bottom: 1.75rem;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 1.1rem;
}

.feature-carousel {
  display: grid;
  gap: 1rem;
}

.carousel-controls {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 0.6rem;
}

.carousel-status {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.76rem;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.carousel-btn {
  border: 1px solid var(--glass-border);
  border-radius: 999px;
  background: rgba(9, 18, 33, 0.92);
  color: var(--text-main);
  padding: 0.58rem 0.8rem;
}

.carousel-fade-enter-active,
.carousel-fade-leave-active {
  transition:
    opacity 0.25s ease,
    transform 0.25s ease;
}

.carousel-fade-enter-from,
.carousel-fade-leave-to {
  opacity: 0;
  transform: translateY(10px);
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
