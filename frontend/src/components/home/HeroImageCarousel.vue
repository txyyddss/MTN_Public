<script setup lang="ts">
import { computed, nextTick, onMounted, onUnmounted, useTemplateRef, watch } from 'vue'

import galleryManifest from '@/assets/generated/gallery-manifest.json'
import { useAutoRotatingIndex } from '@/composables/useAutoRotatingIndex'
import { useMediaQuery } from '@/composables/useMediaQuery'

interface GalleryImage {
  id: string
  label: string
  thumb: string
  large: string
  thumbWidth: number
  thumbHeight: number
  width: number
  height: number
}

const props = defineProps<{
  imageAlt: string
  previousLabel: string
  nextLabel: string
  frameLabel: string
}>()

const images = galleryManifest as GalleryImage[]
const imageCount = computed(() => images.length)
const { matches: prefersReducedMotion } = useMediaQuery('(prefers-reduced-motion: reduce)')
const rotationEnabled = computed(() => !prefersReducedMotion.value && imageCount.value > 1)
const initialIndex = 0
const { currentIndex, next, previous, setIndex } = useAutoRotatingIndex(imageCount, rotationEnabled, {
  intervalMs: 3200,
  initialIndex
})
const currentImage = computed(() => images[currentIndex.value] ?? images[0] ?? null)
const currentImageSrcset = computed(() => {
  if (!currentImage.value) {
    return ''
  }

  return `${currentImage.value.thumb} ${currentImage.value.thumbWidth}w, ${currentImage.value.large} ${currentImage.value.width}w`
})
const currentFrameLabel = computed(() => `${props.frameLabel} ${currentIndex.value + 1}`)
const highPriorityImageId = computed(() => images[initialIndex]?.id ?? images[0]?.id ?? '')
const dotsRef = useTemplateRef<HTMLDivElement>('dots')
const dotRefs = useTemplateRef<HTMLButtonElement[]>('dot')
const dotScrollGutter = 10
let dotsResizeObserver: ResizeObserver | null = null

function getImageAlt(index: number): string {
  return `${props.imageAlt} - ${props.frameLabel} ${index + 1}`
}

function formatDotIndex(index: number): string {
  return String(index + 1).padStart(2, '0')
}

function keepActiveDotVisible(): void {
  void nextTick(() => {
    const dots = dotsRef.value
    const activeDot = dotRefs.value?.[currentIndex.value]

    if (!dots || !activeDot) {
      return
    }

    const maxLeft = Math.max(0, dots.scrollWidth - dots.clientWidth)
    if (maxLeft <= 0) {
      if (dots.scrollLeft !== 0) {
        dots.scrollTo({ left: 0, behavior: 'auto' })
      }
      return
    }

    const visibleLeft = dots.scrollLeft
    const visibleRight = visibleLeft + dots.clientWidth
    const dotLeft = activeDot.offsetLeft
    const dotRight = dotLeft + activeDot.offsetWidth
    let targetLeft = visibleLeft

    if (dotLeft < visibleLeft + dotScrollGutter) {
      targetLeft = dotLeft - dotScrollGutter
    } else if (dotRight > visibleRight - dotScrollGutter) {
      targetLeft = dotRight - dots.clientWidth + dotScrollGutter
    }

    const clampedLeft = Math.min(Math.max(targetLeft, 0), maxLeft)

    if (Math.abs(clampedLeft - visibleLeft) < 1) {
      return
    }

    dots.scrollTo({
      left: clampedLeft,
      behavior: 'auto'
    })
  })
}

watch(currentIndex, keepActiveDotVisible, { flush: 'post' })

onMounted(() => {
  keepActiveDotVisible()

  if (typeof ResizeObserver === 'undefined' || !dotsRef.value) {
    return
  }

  dotsResizeObserver = new ResizeObserver(() => {
    keepActiveDotVisible()
  })
  dotsResizeObserver.observe(dotsRef.value)
})

onUnmounted(() => {
  dotsResizeObserver?.disconnect()
  dotsResizeObserver = null
})
</script>

<template>
  <div class="hero-carousel" :aria-label="imageAlt" aria-roledescription="carousel">
    <Transition name="hero-carousel-fade">
      <div v-if="currentImage" :key="currentImage.id" class="hero-carousel-slide">
        <img
          :src="currentImage.large"
          :srcset="currentImageSrcset"
          sizes="100vw"
          :alt="getImageAlt(currentIndex)"
          class="hero-carousel-image"
          :class="{ 'hero-carousel-image--still': prefersReducedMotion }"
          :width="currentImage.width"
          :height="currentImage.height"
          loading="eager"
          decoding="async"
          :fetchpriority="currentImage.id === highPriorityImageId ? 'high' : 'auto'"
        />
      </div>
    </Transition>

    <div v-if="imageCount > 1" class="hero-carousel-controls" :aria-label="currentFrameLabel">
      <button class="hero-carousel-arrow hero-carousel-arrow--previous" type="button" :aria-label="previousLabel" @click="previous">
        <span aria-hidden="true"></span>
      </button>

      <div ref="dots" class="hero-carousel-dots" aria-label="Gallery images">
        <button
          v-for="(_, index) in images"
          ref="dot"
          :key="`hero-dot-${index}`"
          type="button"
          :class="['hero-carousel-dot', { active: index === currentIndex }]"
          :aria-label="`${frameLabel} ${index + 1}`"
          :aria-current="index === currentIndex ? 'true' : undefined"
          @click="setIndex(index)"
        >
          <span aria-hidden="true">{{ formatDotIndex(index) }}</span>
        </button>
      </div>

      <button class="hero-carousel-arrow hero-carousel-arrow--next" type="button" :aria-label="nextLabel" @click="next">
        <span aria-hidden="true"></span>
      </button>
    </div>
  </div>
</template>

<style scoped>
.hero-carousel {
  position: absolute;
  inset: 0;
  overflow: hidden;
}

.hero-carousel-slide {
  position: absolute;
  inset: 0;
  overflow: hidden;
}

.hero-carousel-image {
  display: block;
  width: 100%;
  height: 100%;
  object-fit: cover;
  filter: saturate(0.78) brightness(0.62) contrast(1.08);
  animation: heroImageDrift 8.8s ease-in-out both;
}

.hero-carousel-image--still {
  animation: none;
  transform: scale(1.04);
}

.hero-carousel-fade-enter-active,
.hero-carousel-fade-leave-active {
  transition:
    opacity 1.05s cubic-bezier(0.16, 1, 0.3, 1),
    transform 1.05s cubic-bezier(0.16, 1, 0.3, 1);
}

.hero-carousel-fade-enter-from {
  opacity: 0;
  transform: scale(1.018);
}

.hero-carousel-fade-leave-to {
  opacity: 0;
  transform: scale(1.034);
}

.hero-carousel-controls {
  position: absolute;
  left: 50%;
  bottom: clamp(5.3rem, 9vh, 7rem);
  z-index: 4;
  display: grid;
  grid-template-columns: 2.8rem minmax(0, 1fr) 2.8rem;
  align-items: center;
  gap: 0.55rem;
  width: min(48rem, calc(100vw - 2rem));
  transform: translateX(-50%);
}

.hero-carousel-arrow,
.hero-carousel-dot {
  display: inline-grid;
  place-items: center;
  border: 1px solid rgba(var(--primary-rgb), 0.32);
  color: #ffffff;
  background: rgba(3, 8, 18, 0.48);
  box-shadow: 0 14px 34px rgba(0, 0, 0, 0.3);
  backdrop-filter: saturate(155%) blur(18px);
  transition:
    transform 0.18s cubic-bezier(0.2, 0, 0.2, 1),
    border-color 0.18s cubic-bezier(0.2, 0, 0.2, 1),
    background 0.18s cubic-bezier(0.2, 0, 0.2, 1);
}

.hero-carousel-arrow {
  width: 2.8rem;
  height: 2.8rem;
  border-radius: 999px;
}

.hero-carousel-arrow span {
  width: 0.68rem;
  height: 0.68rem;
  border-top: 2px solid currentColor;
  border-left: 2px solid currentColor;
}

.hero-carousel-arrow--previous span {
  transform: translateX(0.12rem) rotate(-45deg);
}

.hero-carousel-arrow--next span {
  transform: translateX(-0.12rem) rotate(135deg);
}

.hero-carousel-dots {
  display: flex;
  align-items: center;
  gap: 0.4rem;
  min-width: 0;
  max-width: 100%;
  height: 2.8rem;
  padding: 0.42rem;
  overflow-x: auto;
  overflow-y: hidden;
  scrollbar-width: none;
  border: 1px solid rgba(var(--primary-rgb), 0.28);
  border-radius: 999px;
  background: rgba(3, 8, 18, 0.48);
  box-shadow: 0 14px 34px rgba(0, 0, 0, 0.28);
  backdrop-filter: saturate(155%) blur(18px);
}

.hero-carousel-dots::-webkit-scrollbar {
  display: none;
}

.hero-carousel-dot {
  flex: 0 0 auto;
  width: 2.05rem;
  height: 1.5rem;
  padding: 0;
  border-radius: 999px;
}

.hero-carousel-dot span {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
  border-radius: inherit;
  color: rgba(237, 244, 255, 0.74);
  font-family: var(--mono);
  font-size: 0.62rem;
  font-weight: 700;
  letter-spacing: 0;
}

.hero-carousel-dot.active {
  width: 2.45rem;
  border-color: rgba(147, 197, 253, 0.86);
  background: linear-gradient(135deg, rgba(var(--primary-rgb), 0.92), rgba(var(--secondary-rgb), 0.9));
  box-shadow: 0 10px 24px rgba(var(--secondary-rgb), 0.28);
}

.hero-carousel-dot.active span {
  color: #ffffff;
}

.hero-carousel-arrow:hover,
.hero-carousel-dot:hover {
  transform: translateY(-2px);
  border-color: rgba(147, 197, 253, 0.72);
  background: rgba(var(--secondary-rgb), 0.48);
}

.hero-carousel-arrow:focus-visible,
.hero-carousel-dot:focus-visible {
  outline: 2px solid rgba(255, 255, 255, 0.9);
  outline-offset: 3px;
}

@keyframes heroImageDrift {
  from {
    transform: scale(1.04) translate3d(-0.4%, -0.28%, 0);
  }

  to {
    transform: scale(1.09) translate3d(0.65%, 0.52%, 0);
  }
}

@media (max-width: 680px) {
  .hero-carousel-controls {
    grid-template-columns: 2.55rem minmax(0, 1fr) 2.55rem;
    bottom: 5rem;
    gap: 0.45rem;
    width: min(100%, calc(100vw - 1.25rem));
  }

  .hero-carousel-arrow {
    width: 2.55rem;
    height: 2.55rem;
  }

  .hero-carousel-dots {
    height: 2.55rem;
  }
}

@media (prefers-reduced-motion: reduce) {
  .hero-carousel-image {
    animation: none;
    transform: scale(1.04);
  }

  .hero-carousel-fade-enter-active,
  .hero-carousel-fade-leave-active,
  .hero-carousel-arrow,
  .hero-carousel-dot {
    transition: none;
  }
}
</style>
