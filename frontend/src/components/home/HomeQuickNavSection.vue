<script setup lang="ts">
import { computed } from 'vue'
import { RouterLink } from 'vue-router'

import quickFeatures from '@/assets/home-remake/quick-features.webp'
import quickGallery from '@/assets/home-remake/quick-gallery.webp'
import quickWiki from '@/assets/home-remake/quick-wiki.webp'
import { useRevealOnScroll } from '@/composables/useRevealOnScroll'
import { useSiteContent } from '@/content/siteContent'

const siteContent = useSiteContent()
const { revealed } = useRevealOnScroll<HTMLElement>('quickNavSection', { rootMargin: '0px 0px -16% 0px' })
const cardImages = [quickFeatures, quickGallery, quickWiki]
const cards = computed(() =>
  siteContent.value.home.quickCards.items.map((item, index) => ({
    ...item,
    image: cardImages[index] ?? cardImages[0],
    index: `0${index + 1}`.slice(-2)
  }))
)

function isAnchorLink(href: string): boolean {
  return href.startsWith('#')
}

function handleQuickCardClick(event: MouseEvent, href: string): void {
  if (!isAnchorLink(href)) {
    return
  }

  const target = document.querySelector(href)
  if (!(target instanceof HTMLElement)) {
    return
  }

  event.preventDefault()
  target.scrollIntoView({
    behavior: window.matchMedia('(prefers-reduced-motion: reduce)').matches ? 'auto' : 'smooth',
    block: 'start'
  })
  window.history.replaceState(null, '', href)
}
</script>

<template>
  <section ref="quickNavSection" :class="['quick-nav-section', { 'is-revealed': revealed }]">
    <div class="container quick-nav-shell">
      <div class="quick-nav-heading">
        <span class="section-kicker">{{ siteContent.home.quickCards.title }}</span>
        <h2>MTN ACCESS GRID</h2>
      </div>

      <div class="quick-card-row">
        <template v-for="(card, index) in cards" :key="card.title">
          <a
            v-if="card.external || isAnchorLink(card.href)"
            class="quick-card action-card"
            :href="card.href"
            :target="card.external ? '_blank' : undefined"
            :rel="card.external ? 'noopener noreferrer' : undefined"
            :style="{ '--card-delay': `${index * 0.08}s` }"
            @click="handleQuickCardClick($event, card.href)"
          >
            <div class="quick-card-backdrop" aria-hidden="true">MTN</div>
            <img :src="card.image" :alt="card.title" class="quick-card-media action-media" />
            <div class="quick-card-copy">
              <span>{{ card.index }}</span>
              <strong>{{ card.title }}</strong>
            </div>
          </a>
          <RouterLink
            v-else
            class="quick-card action-card"
            :to="card.href"
            :style="{ '--card-delay': `${index * 0.08}s` }"
          >
            <div class="quick-card-backdrop" aria-hidden="true">MTN</div>
            <img :src="card.image" :alt="card.title" class="quick-card-media action-media" />
            <div class="quick-card-copy">
              <span>{{ card.index }}</span>
              <strong>{{ card.title }}</strong>
            </div>
          </RouterLink>
        </template>
      </div>
    </div>
  </section>
</template>

<style scoped>
.quick-nav-section {
  position: relative;
  overflow: hidden;
  padding: clamp(4rem, 8vw, 6rem) 0;
  background:
    linear-gradient(180deg, rgba(76, 147, 251, 0.08), transparent 22%),
    linear-gradient(180deg, #f7fbff, #eef4fb 72%);
}

.quick-nav-shell {
  display: grid;
  gap: 2rem;
}

.quick-nav-heading {
  display: grid;
  justify-items: center;
  gap: 0.7rem;
}

.quick-nav-heading h2 {
  color: #2a3040;
  font-size: clamp(1.8rem, 4vw, 3.2rem);
  letter-spacing: 0.16em;
  text-align: center;
  text-transform: uppercase;
}

.quick-card-row {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: clamp(1rem, 3vw, 2.2rem);
}

.quick-card {
  position: relative;
  display: grid;
  min-height: 20rem;
  overflow: hidden;
  border: 1px solid rgba(16, 24, 39, 0.12);
  border-radius: 30px;
  background: #08111d;
  color: #ffffff;
  opacity: 0;
  transform: translateY(22px) scale(0.985);
  transition:
    opacity 0.58s var(--transition-slow),
    transform 0.58s var(--transition-slow),
    box-shadow var(--transition-panel),
    border-color var(--transition-fast);
}

.quick-nav-section.is-revealed .quick-card {
  opacity: 1;
  transform: translateY(0) scale(1);
  transition-delay: var(--card-delay);
}

.quick-card:hover,
.quick-card:focus-visible {
  border-color: rgba(76, 147, 251, 0.34);
  box-shadow: 0 28px 64px rgba(10, 20, 36, 0.24);
  transition-delay: 0s;
}

.quick-card-backdrop,
.quick-card-media,
.quick-card-copy {
  position: absolute;
}

.quick-card-backdrop {
  right: 1rem;
  bottom: 0.8rem;
  z-index: 1;
  color: rgba(141, 184, 255, 0.14);
  font-family: var(--mono);
  font-size: clamp(2.8rem, 6vw, 4.2rem);
  font-weight: 700;
  letter-spacing: 0.3em;
}

.quick-card-media {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  filter: brightness(0.64) saturate(0.88);
  transition:
    transform var(--transition-panel),
    filter var(--transition-fast);
}

.quick-card::before {
  content: '';
  position: absolute;
  inset: 0;
  z-index: 1;
  background:
    linear-gradient(180deg, rgba(4, 10, 18, 0.02), rgba(4, 10, 18, 0.78) 72%),
    linear-gradient(130deg, rgba(76, 147, 251, 0.14), transparent 40%);
}

.quick-card:hover .quick-card-media {
  transform: scale(1.05);
  filter: brightness(0.78) saturate(1);
}

.quick-card-copy {
  inset: auto 1.25rem 1.2rem 1.25rem;
  z-index: 2;
  display: grid;
  gap: 0.45rem;
}

.quick-card-copy span {
  display: inline-flex;
  align-items: center;
  width: fit-content;
  padding: 0.38rem 0.65rem;
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.08);
  color: var(--accent);
  font-family: var(--mono);
  font-weight: 700;
  font-size: 0.74rem;
  letter-spacing: 0.14em;
}

.quick-card-copy strong {
  max-width: 11ch;
  color: #ffffff;
  font-size: clamp(1.35rem, 2.3vw, 2rem);
  line-height: 1;
  letter-spacing: 0.04em;
}

@media (max-width: 840px) {
  .quick-card-row {
    grid-template-columns: 1fr;
  }

  .quick-card {
    min-height: 17rem;
  }
}
</style>
