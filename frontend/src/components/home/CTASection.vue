<script setup lang="ts">
import { siteContent } from '@/content/siteContent'
</script>

<template>
  <section class="cta-section animate-entry delay-200">
    <div class="container">
      <div class="cta-box glass-card">
        <div class="cta-copy">
          <span class="section-kicker">Next step</span>
          <h2>{{ siteContent.home.cta.title }}</h2>
          <p>{{ siteContent.home.cta.body }}</p>
        </div>
        <div class="cta-actions">
          <router-link class="btn-primary" to="/server-intro">{{ siteContent.home.cta.primaryCta }}</router-link>
          <p class="cta-note">{{ siteContent.home.cta.note }}</p>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.cta-section {
  padding: 0.4rem 0 5rem;
}

.cta-box {
  display: grid;
  grid-template-columns: minmax(0, 1.4fr) auto;
  gap: 1.6rem;
  align-items: center;
}

.cta-copy {
  display: grid;
  gap: 0.9rem;
}

.cta-copy h2 {
  font-size: clamp(2.1rem, 4vw, 3.8rem);
  max-width: 12ch;
}

.cta-copy p {
  max-width: 46ch;
  color: var(--text-muted);
}

.cta-actions {
  display: grid;
  justify-items: start;
  gap: 0.9rem;
}

.cta-note {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.72rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

@media (max-width: 840px) {
  .cta-box {
    grid-template-columns: 1fr;
  }
}
</style>
