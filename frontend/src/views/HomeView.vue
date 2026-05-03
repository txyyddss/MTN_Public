<script setup lang="ts">
import { defineAsyncComponent, nextTick, onMounted, onUnmounted, shallowRef } from 'vue'

import HeroSection from '@/components/home/HeroSection.vue'

const CTASection = defineAsyncComponent(() => import('@/components/home/CTASection.vue'))
const FeatureGrid = defineAsyncComponent(() => import('@/components/home/FeatureGrid.vue'))
const HomeDetailsSection = defineAsyncComponent(() => import('@/components/home/HomeDetailsSection.vue'))
const ServerStorySection = defineAsyncComponent(() => import('@/components/home/ServerStorySection.vue'))
const showDeferredSections = shallowRef(false)
const deferredHashes = new Set(['#intro-section', '#features', '#details', '#join-mtn'])

let deferredTimer: ReturnType<typeof window.setTimeout> | null = null

function revealDeferredSections(): void {
  if (showDeferredSections.value) {
    return
  }

  showDeferredSections.value = true
}

function revealHashTarget(): void {
  const hash = window.location.hash

  if (!deferredHashes.has(hash)) {
    return
  }

  revealDeferredSections()

  void nextTick(() => {
    document.getElementById(hash.slice(1))?.scrollIntoView({
      behavior: 'smooth',
      block: 'start'
    })
  })
}

onMounted(() => {
  deferredTimer = window.setTimeout(() => {
    revealDeferredSections()
  }, 900)

  window.addEventListener('hashchange', revealHashTarget)
  revealHashTarget()
})

onUnmounted(() => {
  if (deferredTimer) {
    window.clearTimeout(deferredTimer)
  }

  window.removeEventListener('hashchange', revealHashTarget)
})
</script>

<template>
  <div class="home-view">
    <HeroSection />

    <template v-if="showDeferredSections">
      <ServerStorySection />
      <FeatureGrid />
      <HomeDetailsSection />
      <CTASection />
    </template>
  </div>
</template>

<style scoped>
.home-view {
  display: flex;
  flex-direction: column;
  min-width: 0;
  background: #000000;
}
</style>
