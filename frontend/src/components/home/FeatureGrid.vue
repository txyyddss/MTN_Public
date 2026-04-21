<script setup lang="ts">
import { onUnmounted, shallowRef, useTemplateRef, watch } from 'vue'

import FeatureGridCard from '@/components/home/FeatureGridCard.vue'
import { useMediaQuery } from '@/composables/useMediaQuery'
import { useRevealOnScroll } from '@/composables/useRevealOnScroll'
import { siteContent } from '@/content/siteContent'

const features = siteContent.home.features
const { matches: isPhone } = useMediaQuery('(max-width: 680px)')
const currentIndex = shallowRef(0)
const carouselTrack = useTemplateRef<HTMLDivElement>('carouselTrack')
const { revealed } = useRevealOnScroll<HTMLDivElement>('featureCarousel', { rootMargin: '0px 0px -16% 0px' })

let rotationTimer: ReturnType<typeof window.setInterval> | null = null
let resumeTimer: ReturnType<typeof window.setTimeout> | null = null
let scrollFrame: number | null = null

function stopRotation(): void {
  if (rotationTimer) {
    window.clearInterval(rotationTimer)
    rotationTimer = null
  }
}

function clearResumeTimer(): void {
  if (resumeTimer) {
    window.clearTimeout(resumeTimer)
    resumeTimer = null
  }
}

function getStepSize(): number {
  const track = carouselTrack.value
  const firstSlide = track?.querySelector<HTMLElement>('.feature-slide')
  if (!track || !firstSlide) {
    return 1
  }

  const styles = window.getComputedStyle(track)
  const gap = Number.parseFloat(styles.columnGap || styles.gap || '0')
  return firstSlide.offsetWidth + gap
}

function syncCurrentIndex(): void {
  const track = carouselTrack.value
  if (!track) {
    return
  }

  currentIndex.value = Math.max(0, Math.min(features.length - 1, Math.round(track.scrollLeft / getStepSize())))
}

function scrollToIndex(index: number, behavior: ScrollBehavior = 'smooth'): void {
  const track = carouselTrack.value
  if (!track) {
    currentIndex.value = index
    return
  }

  currentIndex.value = index
  track.scrollTo({
    left: getStepSize() * index,
    behavior
  })
}

function scheduleRotation(): void {
  stopRotation()
  clearResumeTimer()

  if (!isPhone.value || !revealed.value || features.length <= 1) {
    return
  }

  rotationTimer = window.setInterval(() => {
    scrollToIndex((currentIndex.value + 1) % features.length)
  }, 3600)
}

function pauseRotationAfterInteraction(): void {
  stopRotation()
  clearResumeTimer()

  if (!isPhone.value || !revealed.value) {
    return
  }

  resumeTimer = window.setTimeout(() => {
    scheduleRotation()
  }, 4200)
}

function handleTrackScroll(): void {
  pauseRotationAfterInteraction()

  if (scrollFrame !== null) {
    window.cancelAnimationFrame(scrollFrame)
  }

  scrollFrame = window.requestAnimationFrame(() => {
    syncCurrentIndex()
    scrollFrame = null
  })
}

watch([isPhone, revealed], () => {
  scheduleRotation()
})

onUnmounted(() => {
  stopRotation()
  clearResumeTimer()
  if (scrollFrame !== null) {
    window.cancelAnimationFrame(scrollFrame)
  }
})
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

      <div v-else ref="featureCarousel" class="feature-carousel glass-card animate-entry delay-100">
        <div ref="carouselTrack" class="carousel-track" @scroll.passive="handleTrackScroll">
          <div v-for="(feature, index) in features" :key="feature.title" class="feature-slide">
            <FeatureGridCard :feature="feature" :index="index" compact />
          </div>
        </div>

        <div class="carousel-dots" aria-hidden="true">
          <span
            v-for="(_, index) in features"
            :key="index"
            :class="['carousel-dot', { active: index === currentIndex }]"
          ></span>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.features-section {
  padding: 0.45rem 0 3.8rem;
}

.section-head {
  display: grid;
  gap: 0.9rem;
  margin-bottom: 1.5rem;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 0.8rem;
}

.feature-carousel {
  display: grid;
  gap: 0.7rem;
}

.carousel-track {
  display: grid;
  grid-auto-flow: column;
  grid-auto-columns: 100%;
  gap: 0.75rem;
  overflow-x: auto;
  scroll-snap-type: x mandatory;
  scrollbar-width: none;
}

.carousel-track::-webkit-scrollbar {
  display: none;
}

.feature-slide {
  scroll-snap-align: start;
}

.carousel-dots {
  display: flex;
  justify-content: center;
  gap: 0.45rem;
}

.carousel-dot {
  width: 0.45rem;
  height: 0.45rem;
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.14);
  transition:
    transform var(--transition-fast),
    background var(--transition-fast);
}

.carousel-dot.active {
  background: var(--primary);
  transform: scale(1.1);
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
