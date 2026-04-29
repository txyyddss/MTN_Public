<script setup lang="ts">
import { useSiteContent } from '@/content/siteContent'

const siteContent = useSiteContent()
</script>

<template>
  <section id="join-mtn" class="cta-section">
    <div class="cta-backdrop" aria-hidden="true"></div>
    <div class="container cta-shell glass-card">
      <span class="cta-kicker">{{ siteContent.home.cta.kicker }}</span>
      <h2>{{ siteContent.home.cta.title }}</h2>
      <p>{{ siteContent.home.cta.body }}</p>
      <div class="cta-actions">
        <a class="btn-primary" :href="siteContent.home.cta.qqGroupUrl" target="_blank" rel="noopener noreferrer">
          {{ siteContent.home.cta.primaryCta }}
        </a>
        <a class="btn-secondary" :href="siteContent.home.cta.siteUrl" target="_blank" rel="noopener noreferrer">
          {{ siteContent.home.cta.siteCta }}
        </a>
        <a class="btn-secondary" :href="siteContent.home.cta.bilibiliUrl" target="_blank" rel="noopener noreferrer">
          {{ siteContent.home.cta.bilibiliCta }}
        </a>
      </div>
      <small>{{ siteContent.home.cta.note }}</small>
    </div>
  </section>
</template>

<style scoped>
.cta-section {
  position: relative;
  overflow: hidden;
  padding: clamp(5rem, 8vw, 7rem) 0;
  background:
    radial-gradient(circle at 50% 0%, rgba(var(--primary-rgb), 0.32), transparent 38%),
    linear-gradient(180deg, #020713, #000000);
  color: var(--text-strong);
}

.cta-backdrop {
  position: absolute;
  inset: -30% auto auto 50%;
  width: min(60vw, 560px);
  aspect-ratio: 1;
  border: 1px solid rgba(147, 197, 253, 0.2);
  background: radial-gradient(circle, rgba(var(--secondary-rgb), 0.22), transparent 62%);
  transform: rotate(45deg);
  animation: ctaPulse 4.8s ease-in-out infinite;
}

.cta-shell {
  position: relative;
  display: grid;
  justify-items: center;
  gap: 1rem;
  width: min(980px, calc(100vw - 2rem));
  padding: clamp(1.6rem, 4vw, 3rem);
  text-align: center;
}

.cta-kicker {
  font-family: var(--mono);
  font-size: 0.78rem;
  font-weight: 700;
  letter-spacing: 0.28em;
  text-transform: uppercase;
}

.cta-shell h2 {
  max-width: 13ch;
  color: var(--text-strong);
  font-size: clamp(2.8rem, 7vw, 6rem);
  letter-spacing: 0;
}

.cta-shell p {
  max-width: 58ch;
  color: var(--text-muted);
}

.cta-actions {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 0.8rem;
  padding-top: 0.8rem;
}

.cta-actions .btn-primary {
  border-color: rgba(147, 197, 253, 0.5);
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  color: #ffffff;
  box-shadow: 0 12px 28px rgba(var(--secondary-rgb), 0.28);
}

.cta-actions .btn-secondary {
  border-color: var(--control-border-hover);
  background: var(--control-surface);
  color: var(--text-main);
}

.cta-shell small {
  color: var(--text-dim);
  font-family: var(--mono);
}

@keyframes ctaPulse {
  0%,
  100% {
    opacity: 0.52;
    transform: rotate(45deg) scale(0.94);
  }

  50% {
    opacity: 0.85;
    transform: rotate(45deg) scale(1.06);
  }
}

@media (max-width: 620px) {
  .cta-actions {
    display: grid;
    width: 100%;
  }

  .cta-actions .btn-primary,
  .cta-actions .btn-secondary {
    width: 100%;
  }
}
</style>
