<script setup lang="ts">
import { useRevealOnScroll } from '@/composables/useRevealOnScroll'
import { useSiteContent } from '@/content/siteContent'

const siteContent = useSiteContent()
const { revealed } = useRevealOnScroll<HTMLElement>('newsSection', { rootMargin: '0px 0px -16% 0px' })
</script>

<template>
  <section id="news" ref="newsSection" :class="['home-news', { 'is-revealed': revealed }]">
    <div class="container news-shell">
      <header class="news-heading">
        <span>{{ siteContent.home.news.kicker }}</span>
        <h2>{{ siteContent.home.news.title }}</h2>
      </header>

      <div class="news-list">
        <a v-for="item in siteContent.home.news.items" :key="item.date + item.title" class="news-item" :href="item.href">
          <time>{{ item.date }}</time>
          <span>{{ item.title }}</span>
          <strong v-if="item.badge">{{ item.badge }}</strong>
        </a>
      </div>
    </div>
  </section>
</template>

<style scoped>
.home-news {
  position: relative;
  padding: clamp(4.5rem, 8vw, 7rem) 0;
  background: #ffffff;
  color: #2b3040;
}

.news-shell {
  display: grid;
  gap: clamp(2rem, 5vw, 4rem);
  max-width: 980px;
}

.news-heading {
  display: grid;
  justify-items: center;
  gap: 1rem;
  text-align: center;
  opacity: 0;
  transform: translateY(22px);
  transition:
    opacity 0.58s var(--transition-slow),
    transform 0.58s var(--transition-slow);
}

.home-news.is-revealed .news-heading {
  opacity: 1;
  transform: translateY(0);
}

.news-heading span {
  display: inline-flex;
  align-items: center;
  gap: 0.85rem;
  color: var(--primary);
  font-family: var(--mono);
  font-size: 0.76rem;
  letter-spacing: 0.28em;
  text-transform: uppercase;
}

.news-heading span::before,
.news-heading span::after {
  content: '';
  width: 2.2rem;
  height: 1px;
  background: currentColor;
}

.news-heading h2 {
  color: #2a3040;
  font-size: clamp(2.35rem, 5vw, 4rem);
  letter-spacing: 0;
}

.news-list {
  display: grid;
  border-top: 1px solid rgba(16, 24, 39, 0.12);
}

.news-item {
  display: grid;
  grid-template-columns: 9rem minmax(0, 1fr) auto;
  gap: 1.3rem;
  align-items: center;
  padding: 1.05rem 1rem;
  border-bottom: 1px solid rgba(16, 24, 39, 0.12);
  color: #303747;
  opacity: 0;
  transform: translateY(16px);
  transition:
    opacity 0.5s var(--transition-slow),
    transform 0.5s var(--transition-slow),
    background var(--transition-fast),
    color var(--transition-fast);
}

.home-news.is-revealed .news-item {
  opacity: 1;
  transform: translateY(0);
}

.news-item:nth-child(2) {
  transition-delay: 0.06s;
}

.news-item:nth-child(3) {
  transition-delay: 0.12s;
}

.news-item:nth-child(4) {
  transition-delay: 0.18s;
}

.news-item:hover {
  background: rgba(76, 147, 251, 0.055);
  color: var(--primary);
}

.news-item time {
  font-family: var(--mono);
  font-weight: 700;
}

.news-item strong {
  color: #83a8e9;
  font-family: var(--mono);
  font-size: 0.7rem;
  letter-spacing: 0.14em;
}

@media (max-width: 680px) {
  .news-item {
    grid-template-columns: 1fr;
    gap: 0.3rem;
  }
}
</style>
