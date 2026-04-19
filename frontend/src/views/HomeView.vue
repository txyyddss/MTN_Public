<script setup lang="ts">
import HeroSection from '@/components/home/HeroSection.vue'
import FeatureGrid from '@/components/home/FeatureGrid.vue'
import CTASection from '@/components/home/CTASection.vue'
import ServerStatusAndConnection from '@/components/ServerStatusAndConnection.vue'
</script>

<template>
  <div class="home-view">
    <HeroSection />
    
    <div class="container status-reveal">
      <ServerStatusAndConnection />
    </div>

    <FeatureGrid />

    <CTASection />
  </div>
</template>

<style scoped>
.home-view {
  display: flex;
  flex-direction: column;
  background-color: var(--bg-dark);
  gap: 0; /* Components have their own padding/margins */
  padding-bottom: 0;
}

.status-reveal {
  margin-top: -100px;
  position: relative;
  z-index: 20;
}

@media (max-width: 768px) {
  .status-reveal {
    margin-top: -40px;
  }
}
</style>
