<script setup lang="ts">
import FeatureGrid from '@/components/home/FeatureGrid.vue'
import HeroSection from '@/components/home/HeroSection.vue'
import ServerStatusAndConnection from '@/components/ServerStatusAndConnection.vue'
</script>

<template>
  <div class="home-view page-shell">
    <HeroSection />
    <ServerStatusAndConnection class="status-reveal" />
    <FeatureGrid />
  </div>
</template>

<style scoped>
.home-view {
  display: flex;
  flex-direction: column;
  gap: 0;
  padding-bottom: 0;
}

.status-reveal {
  margin-top: -0.4rem;
  position: relative;
  z-index: 20;
}

@media (max-width: 768px) {
  .status-reveal {
    margin-top: 0.2rem;
  }
}
</style>
