<script setup lang="ts">
import { onMounted } from 'vue'

import { useSiteContent } from '@/content/siteContent'

const wikiUrl = 'https://mtn.1919801.xyz/'
const siteContent = useSiteContent()

onMounted(() => {
  window.location.replace(wikiUrl)
})
</script>

<template>
  <div class="wiki-redirect container page-shell">
    <div class="glass-card redirect-card">
      <span class="section-kicker">{{ siteContent.wikiRedirect.kicker }}</span>
      <h1>{{ siteContent.wikiRedirect.title }}</h1>
      <p>{{ siteContent.wikiRedirect.body }}</p>
      <a class="btn-primary" :href="wikiUrl">{{ siteContent.wikiRedirect.cta }}</a>
    </div>
  </div>
</template>

<style scoped>
.wiki-redirect {
  display: grid;
  place-items: center;
}

.redirect-card {
  width: min(720px, 100%);
  display: grid;
  gap: 1rem;
}

.redirect-card p {
  color: var(--text-muted);
}
</style>
