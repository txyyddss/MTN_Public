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
  padding: 1rem 0 6rem;
}

.cta-box {
  display: grid;
  grid-template-columns: 1.4fr auto;
  gap: 2rem;
  align-items: end;
}

.cta-copy {
  display: grid;
  gap: 1rem;
}

.cta-copy h2 {
  font-size: clamp(2.2rem, 4vw, 4rem);
  max-width: 14ch;
}

.cta-copy p {
  max-width: 58ch;
  color: var(--text-muted);
}

.cta-actions {
  display: grid;
  justify-items: start;
  gap: 0.9rem;
}

.cta-note {
  color: var(--text-dim);
  font-size: 0.86rem;
}

@media (max-width: 840px) {
  .cta-box {
    grid-template-columns: 1fr;
  }
}
</style>
