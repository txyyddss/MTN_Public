<script setup lang="ts">
import { useRevealOnScroll } from '@/composables/useRevealOnScroll'
import ServerStatusAndConnection from '@/components/ServerStatusAndConnection.vue'
import { useSiteContent } from '@/content/siteContent'

const siteContent = useSiteContent()
const { revealed } = useRevealOnScroll<HTMLElement>('detailsSection', { rootMargin: '0px 0px -14% 0px' })
</script>

<template>
  <section id="details" ref="detailsSection" :class="['home-details', { 'is-revealed': revealed }]">
    <div class="details-stars" aria-hidden="true"></div>
    <div class="container details-shell">
      <div class="details-grid">
        <article
          v-for="(group, groupIndex) in siteContent.home.details.groups"
          :key="group.title"
          class="details-group"
          :style="{ '--group-delay': `${groupIndex * 0.12}s` }"
        >
          <h3>{{ group.title }}</h3>
          <dl>
            <div
              v-for="(row, rowIndex) in group.rows"
              :key="row.label"
              class="details-row"
              :style="{ '--row-delay': `${groupIndex * 0.12 + rowIndex * 0.06}s` }"
            >
              <dt>{{ row.label }}</dt>
              <dd>{{ row.value }}</dd>
            </div>
          </dl>
        </article>
      </div>

      <div class="details-slogan">
        <span>{{ siteContent.home.details.slogan }}</span>
        <strong>{{ siteContent.home.details.subSlogan }}</strong>
      </div>
    </div>

    <ServerStatusAndConnection class="details-status" />
  </section>
</template>

<style scoped>
.home-details {
  position: relative;
  overflow: hidden;
  padding: clamp(5rem, 8vw, 7rem) 0;
  background:
    linear-gradient(180deg, #111629, #060812 72%),
    #060812;
}

.details-stars {
  position: absolute;
  inset: 0;
  opacity: 0.28;
  background-image:
    radial-gradient(circle, rgba(255, 255, 255, 0.8) 1px, transparent 1px),
    radial-gradient(circle, rgba(141, 184, 255, 0.8) 1px, transparent 1px);
  background-position: 0 0, 32px 48px;
  background-size: 86px 86px, 132px 132px;
}

.details-shell {
  position: relative;
  display: grid;
  gap: 3rem;
}

.details-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: clamp(1rem, 3vw, 2rem);
}

.details-group {
  position: relative;
  overflow: hidden;
  display: grid;
  gap: 1.4rem;
  padding: clamp(1.25rem, 3vw, 1.7rem);
  border: 1px solid rgba(255, 255, 255, 0.12);
  border-radius: 28px;
  background:
    linear-gradient(180deg, rgba(255, 255, 255, 0.045), rgba(255, 255, 255, 0.012)),
    rgba(9, 12, 20, 0.52);
  box-shadow: 0 22px 56px rgba(0, 0, 0, 0.24);
  color: #ffffff;
  opacity: 0;
  transform: translateY(22px);
  transition:
    opacity 0.72s var(--transition-slow),
    transform 0.72s var(--transition-slow);
  transition-delay: var(--group-delay);
}

.details-group::before {
  content: '';
  position: absolute;
  inset: 0;
  z-index: 2;
  background:
    linear-gradient(118deg, rgba(76, 147, 251, 0.3), rgba(10, 12, 20, 0.96) 46%),
    rgba(10, 12, 20, 0.96);
  transform: translateX(0);
  transition: transform 0.9s cubic-bezier(0.83, 0, 0.17, 1);
  transition-delay: calc(var(--group-delay) + 0.05s);
}

.home-details.is-revealed .details-group {
  opacity: 1;
  transform: translateY(0);
}

.home-details.is-revealed .details-group::before {
  transform: translateX(102%);
}

.details-group h3 {
  position: relative;
  z-index: 1;
  color: #ffffff;
  font-family: var(--mono);
  font-size: 0.96rem;
  letter-spacing: 0.22em;
  text-transform: uppercase;
}

.details-group dl {
  display: grid;
  gap: 0.9rem;
  margin: 0;
}

.details-row {
  position: relative;
  z-index: 1;
  display: grid;
  grid-template-columns: 8rem minmax(0, 1fr);
  gap: 1rem;
  align-items: baseline;
  opacity: 0;
  transform: translateY(14px);
  transition:
    opacity 0.52s var(--transition-slow),
    transform 0.52s var(--transition-slow);
  transition-delay: calc(0.18s + var(--row-delay));
}

.home-details.is-revealed .details-row {
  opacity: 1;
  transform: translateY(0);
}

.details-row dt {
  color: rgba(255, 255, 255, 0.52);
  font-weight: 700;
}

.details-row dd {
  margin: 0;
  color: #ffffff;
  font-weight: 700;
}

.details-slogan {
  display: grid;
  justify-items: center;
  gap: 0.5rem;
  margin: 1.5rem 0 0;
  color: #ffffff;
  text-align: center;
}

.details-slogan span {
  font-size: clamp(2rem, 5vw, 4rem);
  font-weight: 800;
}

.details-slogan strong {
  color: rgba(255, 255, 255, 0.48);
  font-family: var(--mono);
  letter-spacing: 0.18em;
}

.details-status {
  position: relative;
  z-index: 1;
  margin-top: 2rem;
}

@media (max-width: 860px) {
  .details-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 560px) {
  .details-row {
    grid-template-columns: 1fr;
    gap: 0.15rem;
  }
}
</style>
