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
</script>

<template>
  <section ref="quickNavSection" :class="['quick-nav-section', { 'is-revealed': revealed }]">
    <div class="container quick-nav-shell">
      <h2>{{ siteContent.home.quickCards.title }}</h2>

      <div class="quick-card-row">
        <template v-for="card in cards" :key="card.title">
          <a v-if="card.external" class="quick-card" :href="card.href" target="_blank" rel="noopener noreferrer">
            <img :src="card.image" :alt="card.title" />
            <span>{{ card.index }}</span>
            <strong>{{ card.title }}</strong>
          </a>
          <RouterLink v-else class="quick-card" :to="card.href">
            <img :src="card.image" :alt="card.title" />
            <span>{{ card.index }}</span>
            <strong>{{ card.title }}</strong>
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
    #f5f8fc;
}

.quick-nav-shell {
  display: grid;
  gap: 2rem;
}

.quick-nav-shell h2 {
  color: #2a3040;
  font-size: clamp(1.1rem, 2.6vw, 1.6rem);
  letter-spacing: 0.22em;
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
  min-height: 22rem;
  align-content: end;
  overflow: visible;
  color: #ffffff;
  opacity: 0;
  transform: translateY(22px);
  transition:
    opacity 0.58s var(--transition-slow),
    transform 0.58s var(--transition-slow);
}

.quick-nav-section.is-revealed .quick-card {
  opacity: 1;
  transform: translateY(0);
}

.quick-card:nth-child(2) {
  transition-delay: 0.1s;
}

.quick-card:nth-child(3) {
  transition-delay: 0.2s;
}

.quick-card img {
  position: absolute;
  inset: 0;
  width: calc(100% - 1.8rem);
  height: 100%;
  object-fit: cover;
  border: 7px solid #ffffff;
  box-shadow: 0 22px 54px rgba(16, 24, 39, 0.18);
  filter: brightness(0.72) saturate(0.85);
  transition:
    transform var(--transition-panel),
    filter var(--transition-fast);
}

.quick-card:hover img {
  transform: translateY(-8px);
  filter: brightness(0.82) saturate(0.95);
}

.quick-card span,
.quick-card strong {
  position: relative;
  z-index: 1;
  justify-self: end;
  margin-right: -0.2rem;
  writing-mode: vertical-rl;
  text-orientation: mixed;
}

.quick-card span {
  margin-bottom: 0.8rem;
  color: var(--accent);
  font-family: var(--mono);
  font-weight: 700;
}

.quick-card strong {
  padding: 1rem 0.55rem;
  background: #ffffff;
  color: #2a3040;
  box-shadow: 0 12px 30px rgba(16, 24, 39, 0.14);
  font-size: 1.1rem;
  letter-spacing: 0.16em;
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
