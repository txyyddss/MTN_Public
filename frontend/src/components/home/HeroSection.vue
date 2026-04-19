<script setup lang="ts">
import { computed } from 'vue'
import { storeToRefs } from 'pinia'

import HeroPrimaryPanel from '@/components/home/HeroPrimaryPanel.vue'
import { siteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'

const statusStore = useServerStatusStore()
const { status } = storeToRefs(statusStore)

const onlineCount = computed(() => status.value?.java?.players ?? 0)
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
        :updated-label="updatedLabel"
      />
    </div>
  </section>
</template>

<style scoped>
.hero {
  padding: 2.4rem 0 2rem;
}

.hero-grid {
  display: block;
}
</style>
