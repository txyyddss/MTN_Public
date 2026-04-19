<script setup lang="ts">
import { siteContent } from '@/content/siteContent'
</script>

<template>
  <section class="operations-grid">
    <article class="glass-card principles-card animate-entry delay-100">
      <span class="section-kicker">Operating stance</span>
      <h2>How the server describes itself.</h2>
      <div class="principle-list">
        <div v-for="principle in siteContent.intro.principles" :key="principle.title" class="principle-row">
          <h3>{{ principle.title }}</h3>
          <p>{{ principle.description }}</p>
        </div>
      </div>
    </article>

    <article class="glass-card infrastructure-card animate-entry delay-200">
      <span class="section-kicker">Visibility</span>
      <h2>{{ siteContent.intro.infrastructure.title }}</h2>
      <p>{{ siteContent.intro.infrastructure.body }}</p>
    </article>
  </section>
</template>

<style scoped>
.operations-grid {
  display: grid;
  grid-template-columns: 1.3fr 0.9fr;
  gap: 1.25rem;
  margin-top: 1.25rem;
}

.principles-card,
.infrastructure-card {
  display: grid;
  gap: 1rem;
}

.principles-card h2,
.infrastructure-card h2 {
  font-size: clamp(2rem, 4vw, 3.4rem);
}

.principle-list {
  display: grid;
  gap: 0.9rem;
}

.principle-row {
  display: grid;
  gap: 0.45rem;
  padding: 1rem;
  border-radius: 18px;
  background: rgba(255, 248, 234, 0.04);
  border: 1px solid rgba(255, 248, 234, 0.06);
}

.principle-row h3 {
  font-size: 1.3rem;
}

.principle-row p,
.infrastructure-card p {
  color: var(--text-muted);
}

@media (max-width: 980px) {
  .operations-grid {
    grid-template-columns: 1fr;
  }
}
</style>
