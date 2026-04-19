<script setup lang="ts">
import { computed } from 'vue'

import { useAutoRotatingIndex } from '@/composables/useAutoRotatingIndex'
import { useMediaQuery } from '@/composables/useMediaQuery'
import { siteContent } from '@/content/siteContent'

const principles = siteContent.intro.principles
const { matches: isPhone } = useMediaQuery('(max-width: 680px)')
const { currentIndex, next, previous } = useAutoRotatingIndex(computed(() => principles.length), isPhone, 4200)
const visiblePrinciple = computed(() => principles[currentIndex.value] ?? principles[0])
</script>

<template>
  <section class="operations-grid">
    <article class="glass-card principles-card animate-entry delay-100">
      <span class="section-kicker">Operating stance</span>
      <h2>How the server runs.</h2>
      <div v-if="!isPhone" class="principle-list">
        <div v-for="principle in principles" :key="principle.title" class="principle-row">
          <h3>{{ principle.title }}</h3>
          <p>{{ principle.description }}</p>
        </div>
      </div>

      <div v-else class="principle-carousel">
        <Transition name="principle-fade" mode="out-in">
          <div :key="visiblePrinciple.title" class="principle-row">
            <h3>{{ visiblePrinciple.title }}</h3>
            <p>{{ visiblePrinciple.description }}</p>
          </div>
        </Transition>

        <div class="carousel-controls">
          <button type="button" class="carousel-btn" @click="previous">Prev</button>
          <span class="carousel-status">{{ currentIndex + 1 }} / {{ principles.length }}</span>
          <button type="button" class="carousel-btn" @click="next">Next</button>
        </div>
      </div>
    </article>

    <article class="glass-card infrastructure-card animate-entry delay-200">
      <span class="section-kicker">Visibility</span>
      <h2>{{ siteContent.intro.infrastructure.title }}</h2>
      <p>{{ siteContent.intro.infrastructure.body }}</p>
    </article>
  </section>
</template>

<style scoped>
.operations-grid {
  display: grid;
  grid-template-columns: 1.3fr 0.9fr;
  gap: 1.25rem;
  margin-top: 1.25rem;
}

.principles-card,
.infrastructure-card {
  display: grid;
  gap: 1rem;
}

.principles-card h2,
.infrastructure-card h2 {
  font-size: clamp(2rem, 4vw, 3.4rem);
}

.principle-list {
  display: grid;
  gap: 0.9rem;
}

.principle-carousel {
  display: grid;
  gap: 1rem;
}

.principle-row {
  display: grid;
  gap: 0.45rem;
  padding: 1rem;
  border-radius: 18px;
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid rgba(255, 255, 255, 0.06);
}

.principle-row h3 {
  font-size: 1.3rem;
}

.principle-row p,
.infrastructure-card p {
  color: var(--text-muted);
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
  background: rgba(255, 255, 255, 0.03);
  color: var(--text-main);
  padding: 0.58rem 0.8rem;
}

.principle-fade-enter-active,
.principle-fade-leave-active {
  transition:
    opacity 0.26s ease,
    transform 0.26s ease;
}

.principle-fade-enter-from,
.principle-fade-leave-to {
  opacity: 0;
  transform: translateY(10px);
}

@media (max-width: 980px) {
  .operations-grid {
    grid-template-columns: 1fr;
  }
}
</style>
