<script setup lang="ts">
import { onUnmounted, shallowRef } from 'vue'

import { useSiteContent } from '@/content/siteContent'

const siteContent = useSiteContent()
const copied = shallowRef(false)
let copiedTimer: ReturnType<typeof window.setTimeout> | null = null

function markCopied(): void {
  copied.value = true

  if (copiedTimer) {
    window.clearTimeout(copiedTimer)
  }

  copiedTimer = window.setTimeout(() => {
    copied.value = false
    copiedTimer = null
  }, 1800)
}

function copyWithFallback(text: string): boolean {
  const textarea = document.createElement('textarea')
  textarea.value = text
  textarea.setAttribute('readonly', '')
  textarea.style.position = 'fixed'
  textarea.style.left = '-9999px'
  document.body.appendChild(textarea)
  textarea.select()

  try {
    return document.execCommand('copy')
  } finally {
    document.body.removeChild(textarea)
  }
}

async function copyGroupNumber(): Promise<void> {
  const groupNumber = siteContent.value.home.cta.qqGroup

  try {
    if (navigator.clipboard) {
      await navigator.clipboard.writeText(groupNumber)
    } else if (!copyWithFallback(groupNumber)) {
      return
    }

    markCopied()
  } catch {
    if (copyWithFallback(groupNumber)) {
      markCopied()
    }
  }
}

onUnmounted(() => {
  if (copiedTimer) {
    window.clearTimeout(copiedTimer)
  }
})
</script>

<template>
  <section id="join-mtn" class="cta-section">
    <div class="cta-backdrop" aria-hidden="true"></div>
    <div class="container cta-shell">
      <span class="cta-kicker">{{ siteContent.home.cta.kicker }}</span>
      <h2>{{ siteContent.home.cta.title }}</h2>
      <p>{{ siteContent.home.cta.body }}</p>
      <div class="cta-actions">
        <button class="btn-primary" type="button" @click="copyGroupNumber">
          {{ copied ? siteContent.home.cta.copiedLabel : siteContent.home.cta.primaryCta }}
        </button>
        <a class="btn-secondary" :href="siteContent.home.cta.siteUrl" target="_blank" rel="noopener noreferrer">
          {{ siteContent.home.cta.siteCta }}
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
  background: var(--primary);
  color: #ffffff;
}

.cta-backdrop {
  position: absolute;
  inset: -30% auto auto 50%;
  width: min(60vw, 560px);
  aspect-ratio: 1;
  border: 1px solid rgba(255, 255, 255, 0.24);
  background: radial-gradient(circle, rgba(255, 255, 255, 0.2), transparent 62%);
  transform: rotate(45deg);
  animation: ctaPulse 4.8s ease-in-out infinite;
}

.cta-shell {
  position: relative;
  display: grid;
  justify-items: center;
  gap: 1rem;
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
  color: #ffffff;
  font-size: clamp(2.8rem, 7vw, 6rem);
  letter-spacing: 0;
}

.cta-shell p {
  max-width: 58ch;
  color: rgba(255, 255, 255, 0.86);
}

.cta-actions {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 0.8rem;
  padding-top: 0.8rem;
}

.cta-actions .btn-primary {
  border-color: rgba(255, 255, 255, 0.42);
  background: #ffffff;
  color: var(--primary);
  box-shadow: 0 20px 44px rgba(16, 24, 39, 0.18);
}

.cta-actions .btn-secondary {
  border-color: rgba(255, 255, 255, 0.46);
  background: rgba(255, 255, 255, 0.1);
  color: #ffffff;
}

.cta-shell small {
  color: rgba(255, 255, 255, 0.72);
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
