<script setup lang="ts">
import { computed } from 'vue'
import { storeToRefs } from 'pinia'

import heroArt from '@/assets/hero.png'
import HeroAtlasCard from '@/components/home/HeroAtlasCard.vue'
import HeroPrimaryPanel from '@/components/home/HeroPrimaryPanel.vue'
import { siteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'

const statusStore = useServerStatusStore()
const { status } = storeToRefs(statusStore)

const onlineCount = computed(() => status.value?.java?.players ?? 0)
const javaVersion = computed(() => status.value?.java?.version || 'Fetching live version')
const updatedLabel = computed(() => {
  if (!status.value?.updated) {
    return 'Waiting for telemetry'
  }

  return new Date(status.value.updated).toLocaleTimeString()
})
</script>

<template>
  <section class="hero">
    <div class="container hero-grid">
      <HeroPrimaryPanel
        class="animate-entry"
        :content="siteContent.home.hero"
        :online-count="onlineCount"
        :java-version="javaVersion"
        :updated-label="updatedLabel"
      />

      <HeroAtlasCard
        class="animate-entry delay-200"
        :facts="siteContent.home.hero.facts"
        image-alt="MTNetwork world capture"
        :image-src="heroArt"
      />
    </div>
  </section>
</template>

<style scoped>
.hero {
  padding: 4rem 0 3rem;
}

.hero-grid {
  display: grid;
  grid-template-columns: 1.2fr 0.9fr;
  gap: 2rem;
  align-items: stretch;
}

@media (max-width: 960px) {
  .hero-grid {
    grid-template-columns: 1fr;
  }
}
</style>
