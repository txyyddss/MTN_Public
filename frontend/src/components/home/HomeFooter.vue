<script setup lang="ts">
import { RouterLink } from 'vue-router'

import { useSiteContent } from '@/content/siteContent'

const siteContent = useSiteContent()
</script>

<template>
  <footer class="home-footer">
    <div class="container footer-shell">
      <div class="footer-brand">
        <h2>MTNetwork</h2>
        <p>{{ siteContent.home.footer.tagline }}</p>
        <div class="footer-socials">
          <a
            v-for="link in siteContent.home.follow.links"
            :key="link.href"
            :href="link.href"
            target="_blank"
            rel="noopener noreferrer"
          >
            {{ link.short }}
          </a>
        </div>
      </div>

      <div class="footer-links">
        <h3>{{ siteContent.home.footer.navTitle }}</h3>
        <template v-for="link in siteContent.home.footer.links" :key="link.label">
          <a v-if="link.external" :href="link.href" target="_blank" rel="noopener noreferrer">{{ link.label }}</a>
          <RouterLink v-else :to="link.href">{{ link.label }}</RouterLink>
        </template>
      </div>

      <div class="footer-links">
        <h3>{{ siteContent.home.footer.supportTitle }}</h3>
        <a :href="siteContent.home.cta.siteUrl" target="_blank" rel="noopener noreferrer">{{ siteContent.home.cta.siteCta }}</a>
        <a :href="siteContent.home.floatingJoin.href" target="_blank" rel="noopener noreferrer">{{ siteContent.home.floatingJoin.label }}</a>
      </div>

      <p class="footer-legal">{{ siteContent.home.footer.legal }}</p>
    </div>
  </footer>
</template>

<style scoped>
.home-footer {
  padding: clamp(4rem, 7vw, 6rem) 0 2rem;
  background: #ffffff;
  color: #303747;
}

.footer-shell {
  display: grid;
  grid-template-columns: minmax(0, 1.35fr) minmax(180px, 0.42fr) minmax(180px, 0.42fr);
  gap: clamp(2rem, 5vw, 4rem);
}

.footer-brand {
  display: grid;
  align-content: start;
  gap: 1rem;
}

.footer-brand h2 {
  color: #2a3040;
  font-size: clamp(2.4rem, 5vw, 4.4rem);
  letter-spacing: 0;
}

.footer-brand p {
  max-width: 38rem;
  color: #667084;
}

.footer-socials {
  display: flex;
  gap: 0.65rem;
}

.footer-socials a {
  display: grid;
  place-items: center;
  width: 2.2rem;
  height: 2.2rem;
  border: 1px solid rgba(16, 24, 39, 0.14);
  color: var(--primary);
  font-family: var(--mono);
  font-size: 0.72rem;
  font-weight: 800;
}

.footer-links {
  display: grid;
  align-content: start;
  gap: 0.7rem;
}

.footer-links h3 {
  margin-bottom: 0.35rem;
  color: #2a3040;
  font-size: 0.98rem;
  letter-spacing: 0.12em;
}

.footer-links a {
  color: #667084;
  transition: color var(--transition-fast);
}

.footer-links a:hover {
  color: var(--primary);
}

.footer-legal {
  grid-column: 1 / -1;
  padding-top: 2rem;
  border-top: 1px solid rgba(16, 24, 39, 0.1);
  color: #8791a3;
  font-size: 0.82rem;
}

@media (max-width: 780px) {
  .footer-shell {
    grid-template-columns: 1fr;
  }
}
</style>
