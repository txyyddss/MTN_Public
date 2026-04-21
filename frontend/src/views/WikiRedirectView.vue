<script setup lang="ts">
import { onMounted } from 'vue'

const wikiUrl = 'https://mtn.1919801.xyz/'

onMounted(() => {
  window.location.replace(wikiUrl)
})
</script>

<template>
  <div class="wiki-redirect container page-shell">
    <div class="glass-card redirect-card">
      <span class="section-kicker">Wiki</span>
      <h1>Redirecting to the MTNetwork wiki.</h1>
      <p>If the redirect does not start automatically, use the link below.</p>
      <a class="btn-primary" :href="wikiUrl">Open Wiki</a>
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
