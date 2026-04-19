<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref } from 'vue'

import { useMediaQuery } from '@/composables/useMediaQuery'
import { siteContent } from '@/content/siteContent'

const imageModules = import.meta.glob('/public/gallery-images/*.{png,jpg,jpeg,webp}', { eager: true })
const images = Object.keys(imageModules)
  .map((path) => path.split('/').pop() || '')
  .filter((image): image is string => image.length > 0)
  .sort((left, right) => right.localeCompare(left))

const { matches: isPhone } = useMediaQuery('(max-width: 720px)')
const selectedIndex = ref<number | null>(null)
const selectedImage = computed(() => (selectedIndex.value === null ? null : images[selectedIndex.value]))
const selectedLabel = computed(() =>
  selectedIndex.value === null ? '' : `${siteContent.gallery.frameLabel} ${images.length - selectedIndex.value}`
)

function openLightbox(index: number): void {
  if (isPhone.value) {
    return
  }

  selectedIndex.value = index
  document.body.style.overflow = 'hidden'
}

function closeLightbox(): void {
  selectedIndex.value = null
  document.body.style.overflow = ''
}

function nextImage(): void {
  if (selectedIndex.value === null) {
    return
  }

  selectedIndex.value = (selectedIndex.value + 1) % images.length
}

function previousImage(): void {
  if (selectedIndex.value === null) {
    return
  }

  selectedIndex.value = (selectedIndex.value - 1 + images.length) % images.length
}

function handleKeydown(event: KeyboardEvent): void {
  if (selectedIndex.value === null) {
    return
  }

  if (event.key === 'Escape') {
    closeLightbox()
  } else if (event.key === 'ArrowRight') {
    nextImage()
  } else if (event.key === 'ArrowLeft') {
    previousImage()
  }
}

onMounted(() => {
  window.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  window.removeEventListener('keydown', handleKeydown)
  document.body.style.overflow = ''
})
</script>

<template>
  <div class="gallery-view container page-shell">
    <header class="page-header animate-entry">
      <span class="page-kicker">{{ siteContent.gallery.eyebrow }}</span>
      <h1>{{ siteContent.gallery.title }}</h1>
      <p class="page-lede">{{ siteContent.gallery.body }}</p>
      <div class="badge-pill"><strong>{{ images.length }}</strong> {{ siteContent.gallery.capturesLabel }}</div>
    </header>

    <div class="waterfall-grid">
      <button
        v-for="(image, index) in images"
        :key="image"
        type="button"
        :class="['gallery-card', 'glass-card', 'animate-entry', { passive: isPhone }]"
        :aria-disabled="isPhone"
        :tabindex="isPhone ? -1 : 0"
        :style="{ animationDelay: `${(index % 18) * 0.04}s` }"
        @click="openLightbox(index)"
      >
        <img :src="`/gallery-images/${image}`" :alt="`${siteContent.gallery.frameLabel} ${index + 1}`" loading="lazy" />
        <div class="gallery-meta">
          <strong>{{ images.length - index }}</strong>
          <span>{{ siteContent.gallery.action }}</span>
        </div>
      </button>
    </div>

    <Transition name="lightbox-fade">
      <div v-if="selectedImage" class="lightbox-overlay" @click="closeLightbox">
        <div class="lightbox-shell" @click.stop>
          <button class="lightbox-nav" type="button" aria-label="Previous image" @click="previousImage">&lt;</button>
          <div class="lightbox-image-wrap">
            <p class="lightbox-label">{{ selectedLabel }}</p>
            <img :src="`/gallery-images/${selectedImage}`" :alt="selectedLabel" class="lightbox-image" />
          </div>
          <button class="lightbox-nav" type="button" aria-label="Next image" @click="nextImage">&gt;</button>
        </div>
        <button class="lightbox-close" type="button" @click="closeLightbox">
          {{ siteContent.gallery.close }}
        </button>
      </div>
    </Transition>
  </div>
</template>

<style scoped>
.gallery-view {
  display: grid;
  gap: 1rem;
}

.waterfall-grid {
  columns: 1;
  column-gap: 1rem;
}

@media (min-width: 720px) {
  .waterfall-grid {
    columns: 2;
  }
}

@media (min-width: 1100px) {
  .waterfall-grid {
    columns: 3;
  }
}

.gallery-card {
  break-inside: avoid;
  display: block;
  width: 100%;
  padding: 0.8rem;
  margin-bottom: 1rem;
  cursor: pointer;
}

.gallery-card.passive {
  cursor: default;
}

.gallery-card img {
  width: 100%;
  border-radius: calc(var(--radius-lg) - 12px);
}

.gallery-meta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  margin-top: 0.9rem;
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.74rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.gallery-meta strong {
  color: var(--text-strong);
}

.lightbox-overlay {
  position: fixed;
  inset: 0;
  z-index: 120;
  display: grid;
  place-items: center;
  background: rgba(8, 7, 5, 0.94);
  backdrop-filter: blur(16px);
}

.lightbox-shell {
  display: grid;
  grid-template-columns: auto minmax(0, 1fr) auto;
  gap: 1rem;
  width: min(1200px, calc(100vw - 2rem));
  align-items: center;
}

.lightbox-image-wrap {
  display: grid;
  gap: 0.75rem;
}

.lightbox-label {
  color: var(--primary);
  font-family: var(--mono);
  font-size: 0.8rem;
  text-transform: uppercase;
  letter-spacing: 0.16em;
}

.lightbox-image {
  width: 100%;
  max-height: 82vh;
  object-fit: contain;
  border-radius: 24px;
}

.lightbox-nav,
.lightbox-close {
  border: 1px solid var(--glass-border-strong);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.05);
  color: var(--text-main);
}

.lightbox-nav {
  width: 54px;
  height: 54px;
  font-size: 2rem;
}

.lightbox-close {
  position: absolute;
  top: 1.25rem;
  right: 1.25rem;
  padding: 0.7rem 1rem;
  text-transform: uppercase;
  letter-spacing: 0.08em;
}

.lightbox-fade-enter-active,
.lightbox-fade-leave-active {
  transition: opacity 0.2s ease;
}

.lightbox-fade-enter-from,
.lightbox-fade-leave-to {
  opacity: 0;
}

@media (max-width: 720px) {
  .lightbox-shell {
    grid-template-columns: 1fr;
  }

  .lightbox-nav {
    display: none;
  }

  .gallery-card:hover {
    transform: none;
  }
}
</style>
