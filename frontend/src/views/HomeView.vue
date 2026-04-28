<script setup lang="ts">
import CTASection from '@/components/home/CTASection.vue'
import FeatureGrid from '@/components/home/FeatureGrid.vue'
import HeroSection from '@/components/home/HeroSection.vue'
import HomeDetailsSection from '@/components/home/HomeDetailsSection.vue'
import ServerStorySection from '@/components/home/ServerStorySection.vue'
</script>

<template>
  <div class="home-view">
    <HeroSection />
    <ServerStorySection />

    <FeatureGrid />
    <HomeDetailsSection />
    <CTASection />
  </div>
</template>

<style scoped>
.home-view {
  display: flex;
  flex-direction: column;
  min-width: 0;
  background: #ffffff;
}
</style>
