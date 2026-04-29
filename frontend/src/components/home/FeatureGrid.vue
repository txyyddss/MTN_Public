<script setup lang="ts">
import { computed, onUnmounted, shallowRef, watch } from 'vue'

import { useRevealOnScroll } from '@/composables/useRevealOnScroll'
import { useSiteContent } from '@/content/siteContent'

const siteContent = useSiteContent()
const activeIndex = shallowRef(0)
const { revealed } = useRevealOnScroll<HTMLElement>('featuresSection', { rootMargin: '0px 0px -16% 0px' })
const featureTabs = computed(() => siteContent.value.home.features.slice(0, 4))
const activeFeature = computed(() => featureTabs.value[activeIndex.value] ?? featureTabs.value[0])
const activeFeatureBackdrop = computed(() => siteContent.value.home.featureBackdrops[activeIndex.value] ?? 'MTN')

let rotationTimer: ReturnType<typeof window.setInterval> | null = null

function stopRotation(): void {
  if (rotationTimer) {
    window.clearInterval(rotationTimer)
    rotationTimer = null
  }
}

function startRotation(): void {
  stopRotation()

  if (!revealed.value || window.matchMedia('(prefers-reduced-motion: reduce)').matches) {
    return
  }

  rotationTimer = window.setInterval(() => {
    activeIndex.value = (activeIndex.value + 1) % featureTabs.value.length
  }, 4200)
}

function selectFeature(index: number): void {
  activeIndex.value = index
  startRotation()
}

watch(revealed, startRotation)

onUnmounted(() => {
  stopRotation()
})
</script>

<template>
  <section id="features" ref="featuresSection" :class="['features-section', { 'is-revealed': revealed }]">
    <div class="container features-shell">
      <header class="features-heading">
        <span>{{ siteContent.home.featuresIntro.kicker }}</span>
        <h2>{{ siteContent.home.featuresIntro.title }}</h2>
      </header>

      <div class="feature-diamond-row" role="tablist" :aria-label="siteContent.home.featuresIntro.title">
        <button
          v-for="(feature, index) in featureTabs"
          :key="feature.title"
          :class="['feature-diamond', { active: index === activeIndex }]"
          type="button"
          role="tab"
          :aria-selected="index === activeIndex"
          @click="selectFeature(index)"
        >
          <span>{{ feature.icon }}</span>
        </button>
      </div>

      <Transition name="feature-fade" mode="out-in">
        <article :key="activeFeature.title" class="feature-panel glass-card action-card">
          <span class="feature-backdrop" aria-hidden="true">{{ activeFeatureBackdrop }}</span>
          <div class="feature-number">// {{ activeFeature.icon }}</div>
          <h3>{{ activeFeature.title }}</h3>
          <p>{{ activeFeature.description }}</p>
          <blockquote>{{ siteContent.home.featureQuotes[activeIndex] }}</blockquote>
        </article>
      </Transition>
    </div>
  </section>
</template>

<style scoped>
.features-section {
  position: relative;
  overflow: hidden;
  padding: clamp(5rem, 9vw, 8rem) 0;
  background:
    radial-gradient(circle at 12% 0%, rgba(var(--primary-rgb), 0.2), transparent 30%),
    radial-gradient(circle at 84% 22%, rgba(var(--secondary-rgb), 0.14), transparent 28%),
    linear-gradient(180deg, #000000, #030916 72%, #000000);
  color: var(--text-main);
}

.features-shell {
  display: grid;
  gap: clamp(2rem, 5vw, 4rem);
  justify-items: center;
}

.features-heading {
  position: relative;
  display: grid;
  justify-items: center;
  gap: 1rem;
  text-align: center;
}

.features-heading::before {
  content: 'SYSTEM';
  position: absolute;
  top: -3.4rem;
  color: rgba(147, 197, 253, 0.075);
  font-family: var(--heading);
  font-size: clamp(4.6rem, 11vw, 9rem);
  font-weight: 900;
  letter-spacing: 0;
}

.features-heading span,
.feature-number {
  color: var(--primary);
  font-family: var(--mono);
  font-weight: 700;
  letter-spacing: 0.18em;
  text-transform: uppercase;
}

.features-heading span {
  position: relative;
  z-index: 1;
  font-size: 0.75rem;
}

.features-heading h2 {
  position: relative;
  z-index: 1;
  color: var(--text-strong);
  font-size: clamp(2.2rem, 5vw, 4rem);
  letter-spacing: 0;
}

.features-heading h2::after {
  content: '';
  display: block;
  width: 4rem;
  height: 3px;
  margin: 1rem auto 0;
  background: var(--primary);
}

.feature-diamond-row {
  position: relative;
  display: flex;
  justify-content: center;
  gap: clamp(1.4rem, 5vw, 4rem);
  width: min(760px, 100%);
}

.feature-diamond-row::before {
  content: '';
  position: absolute;
  left: 0;
  right: 0;
  top: 50%;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(var(--secondary-rgb), 0.45), transparent);
}

.feature-diamond {
  position: relative;
  z-index: 1;
  display: grid;
  place-items: center;
  width: clamp(4.6rem, 8vw, 6.2rem);
  aspect-ratio: 1;
  border: 1px solid var(--glass-border-soft);
  border-radius: var(--radius-md);
  background: var(--glass-bg);
  color: var(--text-main);
  transform: translateY(0);
  box-shadow: var(--glass-shadow), var(--glass-inset);
  backdrop-filter: saturate(155%) blur(18px);
  transition:
    background var(--transition-fast),
    color var(--transition-fast),
    border-color var(--transition-fast),
    transform var(--transition-panel);
}

.feature-diamond span {
  font-family: var(--mono);
  font-size: 1rem;
  font-weight: 800;
}

.feature-diamond.active,
.feature-diamond:hover {
  border-color: var(--glass-border-strong);
  background: var(--glass-bg-hover);
  color: var(--text-strong);
  transform: translateY(-4px);
}

.feature-panel {
  position: relative;
  display: grid;
  justify-items: center;
  gap: 1rem;
  width: min(900px, 100%);
  min-height: 300px;
  padding: clamp(2rem, 6vw, 4rem) 1rem;
  border-radius: var(--radius-xl);
  text-align: center;
}

.feature-backdrop {
  position: absolute;
  inset: 50% auto auto 50%;
  color: rgba(147, 197, 253, 0.07);
  font-family: var(--heading);
  font-size: clamp(4rem, 14vw, 9rem);
  font-weight: 900;
  line-height: 1;
  transform: translate(-50%, -50%);
  white-space: nowrap;
}

.feature-number {
  position: relative;
  z-index: 1;
  font-size: 1rem;
}

.feature-panel h3,
.feature-panel p,
.feature-panel blockquote {
  position: relative;
  z-index: 1;
}

.feature-panel h3 {
  color: var(--text-strong);
  font-size: clamp(2.5rem, 6vw, 4.7rem);
  letter-spacing: 0;
}

.feature-panel p {
  max-width: 62ch;
  color: var(--text-muted);
  font-size: 1.06rem;
}

.feature-panel blockquote {
  margin: 0.5rem 0 0;
  color: var(--accent-soft);
  font-family: var(--heading);
  font-size: 1.1rem;
  font-style: italic;
  letter-spacing: 0.08em;
}

.feature-fade-enter-active,
.feature-fade-leave-active {
  transition:
    opacity 0.26s ease,
    transform 0.34s var(--transition-slow);
}

.feature-fade-enter-from,
.feature-fade-leave-to {
  opacity: 0;
  transform: translateY(12px);
}

@media (max-width: 680px) {
  .feature-diamond-row {
    gap: 0.9rem;
  }

  .feature-diamond {
    width: 4.2rem;
  }

  .feature-panel h3 {
    font-size: 2.5rem;
  }
}
</style>
