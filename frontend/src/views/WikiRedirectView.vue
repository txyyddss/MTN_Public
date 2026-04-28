<script setup lang="ts">
import { onMounted } from 'vue'

import RouteHeroHeader from '@/components/common/RouteHeroHeader.vue'
import ThemedPanelFrame from '@/components/common/ThemedPanelFrame.vue'
import { useSiteContent } from '@/content/siteContent'

const wikiUrl = 'https://docs.mcmtn.net/'
const siteContent = useSiteContent()

onMounted(() => {
  window.location.replace(wikiUrl)
})
</script>

<template>
  <div class="wiki-redirect container page-shell route-page-shell">
    <RouteHeroHeader
      :kicker="siteContent.wikiRedirect.kicker"
      :title="siteContent.wikiRedirect.title"
      :body="siteContent.wikiRedirect.body"
      align="center"
    />

    <ThemedPanelFrame tag="section" class="redirect-card animate-entry delay-100" compact variant="archive">
      <div class="redirect-actions">
        <a class="btn-primary" :href="wikiUrl">{{ siteContent.wikiRedirect.cta }}</a>
      </div>
    </ThemedPanelFrame>
  </div>
</template>

<style scoped>
.wiki-redirect {
  display: grid;
  gap: 1rem;
}

.redirect-card {
  width: min(760px, 100%);
  display: grid;
}

.redirect-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 0.8rem;
}

@media (max-width: 640px) {
  .redirect-actions {
    display: grid;
  }

  .redirect-actions .btn-primary {
    width: 100%;
  }
}
</style>
