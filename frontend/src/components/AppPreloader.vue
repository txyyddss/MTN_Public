<script setup lang="ts">
import { computed, onMounted, onUnmounted, shallowRef, watch } from 'vue'

import { useSiteContent } from '@/content/siteContent'

interface Props {
  ready: boolean
}

const props = defineProps<Props>()
const siteContent = useSiteContent()
const visible = shallowRef(true)
const slidingOut = shallowRef(false)
const progress = shallowRef(0)
const particleIndexes = Array.from({ length: 22 }, (_, index) => index)

let animationFrame: number | null = null
let startTime = 0
let exitTimer: ReturnType<typeof window.setTimeout> | null = null

const progressLabel = computed(() => Math.min(100, Math.round(progress.value)))
const progressText = computed(() => {
  const steps = siteContent.value.home.preloader.steps
  const percent = progressLabel.value

  if (percent < 34) {
    return steps[0]
  }

  if (percent < 72) {
    return steps[1]
  }

  if (percent < 99) {
    return steps[2]
  }

  return steps[3]
})

function getParticleStyle(index: number): Record<string, string> {
  const column = (index * 19) % 100
  const size = 7 + (index % 5) * 3
  const duration = 4.4 + (index % 6) * 0.45
  const delay = -(index % 8) * 0.38

  return {
    left: `${column}%`,
    width: `${size}px`,
    height: `${size}px`,
    animationDuration: `${duration}s, ${duration * 0.52}s`,
    animationDelay: `${delay}s, ${delay}s`
  }
}

function finish(): void {
  if (slidingOut.value || !visible.value) {
    return
  }

  progress.value = 100
  slidingOut.value = true
  exitTimer = window.setTimeout(() => {
    visible.value = false
  }, 1050)
}

function tick(now: number): void {
  if (!startTime) {
    startTime = now
  }

  const elapsed = now - startTime
  const target = props.ready && elapsed > 850
    ? 100
    : Math.min(94, 8 + elapsed * 0.042)

  progress.value += (target - progress.value) * 0.075

  if (target === 100 && progress.value > 99.4) {
    finish()
    return
  }

  animationFrame = window.requestAnimationFrame(tick)
}

watch(
  () => props.ready,
  (ready) => {
    if (ready && progress.value > 99) {
      finish()
    }
  }
)

onMounted(() => {
  if (window.matchMedia('(prefers-reduced-motion: reduce)').matches) {
    visible.value = false
    return
  }

  animationFrame = window.requestAnimationFrame(tick)
})

onUnmounted(() => {
  if (animationFrame !== null) {
    window.cancelAnimationFrame(animationFrame)
  }

  if (exitTimer) {
    window.clearTimeout(exitTimer)
  }
})
</script>

<template>
  <div v-if="visible" :class="['app-preloader', { 'is-sliding-out': slidingOut }]" role="status" aria-live="polite">
    <div class="preloader-panel preloader-panel-accent" aria-hidden="true">
      <span v-for="index in particleIndexes" :key="index" class="preloader-particle" :style="getParticleStyle(index)"></span>
      <div class="preloader-marquee">
        <span v-for="index in 7" :key="index">MTN POTATO SERVER</span>
      </div>
    </div>

    <div class="preloader-panel preloader-panel-main">
      <div class="preloader-progress">
        <div class="preloader-progress-fill" :style="{ height: `${progressLabel}%` }">
          <span></span>
        </div>
      </div>

      <div class="preloader-copy">
        <div>
          <strong>{{ progressLabel }}</strong>
          <span>%</span>
        </div>
        <p>{{ progressText }}</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.app-preloader {
  position: fixed;
  inset: 0;
  z-index: 999;
  display: grid;
  overflow: hidden;
  color: var(--text-strong);
  pointer-events: auto;
}

.preloader-panel {
  position: absolute;
  inset: 0;
  transition: transform 0.82s cubic-bezier(0.85, 0, 0.15, 1);
  will-change: transform;
}

.preloader-panel-accent {
  z-index: 1;
  display: grid;
  align-items: center;
  overflow: hidden;
  background:
    radial-gradient(circle at 18% 22%, rgba(255, 255, 255, 0.18), transparent 28%),
    linear-gradient(135deg, #000000, var(--primary) 48%, #05102a);
}

.preloader-panel-main {
  z-index: 2;
  display: grid;
  grid-template-columns: auto minmax(0, 1fr);
  gap: clamp(1.3rem, 4vw, 3rem);
  align-items: center;
  padding: clamp(2rem, 6vw, 5rem);
  background:
    radial-gradient(circle at 56% 50%, rgba(var(--secondary-rgb), 0.18), transparent 38%),
    linear-gradient(135deg, #020713 0%, #000000 100%);
  color: var(--text-main);
}

.app-preloader.is-sliding-out {
  pointer-events: none;
}

.app-preloader.is-sliding-out .preloader-panel-main {
  transform: translateX(100%);
}

.app-preloader.is-sliding-out .preloader-panel-accent {
  transform: translateX(100%);
  transition-delay: 0.34s;
}

.preloader-marquee {
  display: flex;
  gap: 5.5rem;
  padding-left: 6rem;
  color: rgba(255, 255, 255, 0.18);
  font-family: var(--mono);
  font-size: clamp(2.8rem, 8vw, 8rem);
  font-weight: 700;
  white-space: nowrap;
}

.preloader-progress {
  width: 2px;
  height: min(52vh, 440px);
  overflow: visible;
  border-radius: 999px;
  background: rgba(147, 197, 253, 0.16);
}

.preloader-progress-fill {
  position: relative;
  width: 100%;
  background: linear-gradient(180deg, var(--primary), var(--secondary));
  box-shadow: 0 0 18px rgba(var(--secondary-rgb), 0.5);
  transition: height 0.08s linear;
}

.preloader-progress-fill span {
  position: absolute;
  left: 50%;
  bottom: -6px;
  width: 12px;
  height: 12px;
  border: 2px solid var(--primary);
  background: var(--text-strong);
  transform: translateX(-50%) rotate(45deg);
}

.preloader-copy {
  display: grid;
  gap: 0.8rem;
}

.preloader-copy div {
  display: flex;
  align-items: baseline;
  gap: 0.35rem;
}

.preloader-copy strong {
  color: var(--text-strong);
  font-family: var(--mono);
  font-size: clamp(5rem, 12vw, 10rem);
  font-weight: 600;
  line-height: 0.86;
}

.preloader-copy span {
  color: var(--primary);
  font-family: var(--mono);
  font-size: clamp(1.5rem, 3vw, 2.6rem);
  font-weight: 700;
}

.preloader-copy p {
  display: inline-flex;
  align-items: center;
  gap: 0.8rem;
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.75rem;
  letter-spacing: 0.18em;
  text-transform: uppercase;
}

.preloader-copy p::before {
  content: '';
  width: 2.4rem;
  height: 1px;
  background: currentColor;
}

.preloader-particle {
  position: absolute;
  top: -8%;
  border: 1px solid rgba(255, 255, 255, 0.26);
  background: rgba(255, 255, 255, 0.28);
  image-rendering: pixelated;
  animation:
    pixelFall linear infinite,
    pixelSway ease-in-out infinite alternate;
}

@keyframes pixelFall {
  0% {
    top: -8%;
    opacity: 0;
  }

  12%,
  82% {
    opacity: 0.95;
  }

  100% {
    top: 108%;
    opacity: 0;
  }
}

@keyframes pixelSway {
  0% {
    transform: translateX(0) rotate(0deg);
  }

  100% {
    transform: translateX(34px) rotate(90deg);
  }
}

@media (max-width: 680px) {
  .preloader-panel-main {
    grid-template-columns: 1fr;
    align-content: center;
  }

  .preloader-progress {
    height: 38vh;
  }

  .preloader-copy strong {
    font-size: 5.4rem;
  }
}
</style>
