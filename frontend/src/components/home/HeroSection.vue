<script setup lang="ts">
import { computed } from 'vue'
import { storeToRefs } from 'pinia'

import HeroAsidePanel from '@/components/home/HeroAsidePanel.vue'
import HeroPrimaryPanel from '@/components/home/HeroPrimaryPanel.vue'
import { siteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'

const statusStore = useServerStatusStore()

// computed properties for isLive and onlineCount removed as they were only used for HeroAsidePanel
</script>

<template>
  <section class="hero">
    <div class="container hero-grid">
      <HeroPrimaryPanel
        class="animate-entry"
        :content="siteContent.home.hero"
      />
      <HeroAsidePanel />
    </div>
  </section>
</template>

<style scoped>
.hero {
  position: relative;
  padding: 1.8rem 0 2.4rem;
}

.hero::before {
  content: '';
  position: absolute;
  inset: 0 0 auto;
  height: min(32vw, 280px);
  background:
    radial-gradient(circle at 18% 28%, rgba(76, 147, 251, 0.18), transparent 34%),
    radial-gradient(circle at 78% 0%, rgba(141, 184, 255, 0.08), transparent 30%);
  pointer-events: none;
}

.hero-grid {
  position: relative;
  display: grid;
  grid-template-columns: minmax(0, 1.18fr) minmax(320px, 0.82fr);
  gap: 1rem;
  align-items: stretch;
}

@media (max-width: 1040px) {
  .hero-grid {
    grid-template-columns: 1fr;
  }
}
</style>
