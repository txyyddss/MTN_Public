<script setup lang="ts">
import { computed, onUnmounted, shallowRef, watch } from 'vue'

import galleryOne from '@/assets/home-remake/home-gallery-1.webp'
import galleryThree from '@/assets/home-remake/home-gallery-3.webp'
import galleryTwo from '@/assets/home-remake/home-gallery-2.webp'
import { useRevealOnScroll } from '@/composables/useRevealOnScroll'
import { useSiteContent } from '@/content/siteContent'

const siteContent = useSiteContent()
const { revealed } = useRevealOnScroll<HTMLElement>('homeGallerySection', { rootMargin: '0px 0px -16% 0px' })
const activeIndex = shallowRef(0)
const galleryImages = [galleryOne, galleryTwo, galleryThree]
const slides = computed(() =>
  siteContent.value.home.homeGallery.items.map((item, index) => ({
    ...item,
    image: galleryImages[index] ?? galleryImages[0]
  }))
)

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
    activeIndex.value = (activeIndex.value + 1) % slides.value.length
  }, 3200)
}

function setActive(index: number): void {
  activeIndex.value = index
  startRotation()
}

watch(revealed, startRotation)

onUnmounted(() => {
  stopRotation()
})
</script>

<template>
  <section id="gallery" ref="homeGallerySection" :class="['home-gallery', { 'is-revealed': revealed }]">
    <div class="container gallery-shell">
      <header class="gallery-heading">
        <h2>{{ siteContent.home.homeGallery.title }}</h2>
      </header>

      <div class="gallery-stage">
        <button
          v-for="(slide, index) in slides"
          :key="slide.caption"
          :class="['gallery-slide', { active: index === activeIndex }]"
          type="button"
          @click="setActive(index)"
        >
          <span class="gallery-mask" aria-hidden="true"></span>
          <img :src="slide.image" :alt="slide.caption" />
          <strong>{{ slide.caption }}</strong>
        </button>
      </div>

      <p class="gallery-note">{{ siteContent.home.homeGallery.note }}</p>
    </div>
  </section>
</template>

<style scoped>
.home-gallery {
  position: relative;
  overflow: hidden;
  padding: clamp(5rem, 9vw, 8rem) 0;
  background:
    radial-gradient(circle at 50% 0%, rgba(76, 147, 251, 0.12), transparent 32%),
    #252130;
  color: #ffffff;
}

.home-gallery::before {
  content: '';
  position: absolute;
  inset: 0;
  opacity: 0.22;
  background-image: radial-gradient(circle, rgba(255, 255, 255, 0.6) 1px, transparent 1px);
  background-size: 52px 52px;
}

.gallery-shell {
  position: relative;
  display: grid;
  gap: 3rem;
}

.gallery-heading {
  display: grid;
  justify-items: center;
}

.gallery-heading h2 {
  color: #ffffff;
  font-size: clamp(2rem, 4vw, 3.4rem);
  letter-spacing: 0.28em;
}

.gallery-stage {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: clamp(1rem, 3vw, 2rem);
  align-items: start;
}

.gallery-slide {
  position: relative;
  display: grid;
  gap: 1.1rem;
  padding: 0;
  border: 0;
  background: transparent;
  color: #ffffff;
  cursor: pointer;
  opacity: 0;
  transform: translateY(24px);
  transition:
    opacity 0.58s var(--transition-slow),
    transform 0.58s var(--transition-slow);
}

.home-gallery.is-revealed .gallery-slide {
  opacity: 1;
  transform: translateY(0);
}

.gallery-slide:nth-child(2) {
  transition-delay: 0.08s;
}

.gallery-slide:nth-child(3) {
  transition-delay: 0.16s;
}

.gallery-slide img {
  width: 100%;
  aspect-ratio: 16 / 10;
  object-fit: cover;
  border: 6px solid #ffffff;
  background: rgba(255, 255, 255, 0.1);
  filter: saturate(0.78) brightness(0.78);
  transition:
    filter var(--transition-fast),
    transform var(--transition-panel);
}

.gallery-slide.active img,
.gallery-slide:hover img {
  filter: saturate(0.95) brightness(0.98);
  transform: translateY(-6px);
}

.gallery-mask {
  position: absolute;
  left: 6px;
  right: 6px;
  top: 6px;
  z-index: 2;
  height: calc((100% - 2.2rem) * 0.78);
  background: #252130;
  transform: scaleY(1);
  transform-origin: bottom;
  transition: transform 0.72s cubic-bezier(0.85, 0, 0.15, 1);
}

.home-gallery.is-revealed .gallery-mask {
  transform: scaleY(0);
}

.gallery-slide strong {
  color: #ffffff;
  font-family: var(--heading);
  font-size: clamp(1rem, 1.6vw, 1.25rem);
  letter-spacing: 0.16em;
  text-align: center;
}

.gallery-note {
  color: rgba(255, 255, 255, 0.56);
  text-align: center;
}

@media (max-width: 860px) {
  .gallery-stage {
    grid-template-columns: 1fr;
  }
}
</style>
