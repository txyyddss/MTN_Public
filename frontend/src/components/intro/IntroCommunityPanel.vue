<script setup lang="ts">
import { computed } from 'vue'
import { storeToRefs } from 'pinia'

import IntroCommunityCard from '@/components/intro/IntroCommunityCard.vue'
import IntroLiveFactsCard from '@/components/intro/IntroLiveFactsCard.vue'
import { siteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'

const statusStore = useServerStatusStore()
const { status } = storeToRefs(statusStore)

const liveState = computed(() => (status.value?.java?.online ? 'Operational' : 'Offline'))
const updatedLabel = computed(() => {
  if (!status.value?.updated) {
    return 'Waiting for telemetry'
  }

  return new Date(status.value.updated).toLocaleString()
})
</script>

<template>
  <section id="community-panel" class="community-grid">
    <IntroCommunityCard
      :body="siteContent.intro.community.body"
      :caption="siteContent.intro.community.caption"
      :title="siteContent.intro.community.title"
    />
    <IntroLiveFactsCard :live-state="liveState" :updated-label="updatedLabel" />
  </section>
</template>

<style scoped>
.community-grid {
  display: grid;
  grid-template-columns: 0.95fr 1.05fr;
  gap: 1.25rem;
  margin-top: 1.25rem;
}

@media (max-width: 980px) {
  .community-grid {
    grid-template-columns: 1fr;
  }
}
</style>
