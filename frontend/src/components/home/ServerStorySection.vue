<script setup lang="ts">
import serverMascot from '@/assets/server-mascot.png'
import { useRevealOnScroll } from '@/composables/useRevealOnScroll'
import { useSiteContent } from '@/content/siteContent'

const siteContent = useSiteContent()
const { revealed } = useRevealOnScroll<HTMLElement>('introSection', { rootMargin: '0px 0px -18% 0px' })
</script>

<template>
  <section id="intro-section" ref="introSection" :class="['home-intro', { 'is-revealed': revealed }]">
    <div class="intro-texture" aria-hidden="true"></div>
    <div class="container intro-layout">
      <article class="intro-copy">
        <p class="intro-kicker">{{ siteContent.home.story.kicker }}</p>
        <div class="intro-columns">
          <h2>{{ siteContent.home.story.title }}</h2>
          <div class="intro-body">
            <p v-for="paragraph in siteContent.home.story.paragraphs.slice(0, 4)" :key="paragraph">
              {{ paragraph }}
            </p>
          </div>
        </div>
      </article>

      <div class="intro-mascot-wrap" aria-hidden="true">
        <img :src="serverMascot" :alt="siteContent.home.heroAside.imageAlt" class="intro-mascot" />
      </div>
    </div>
  </section>
</template>

<style scoped>
.home-intro {
  position: relative;
  overflow: hidden;
  padding: clamp(5rem, 10vw, 8rem) 0;
  background:
    radial-gradient(circle at 4% 10%, rgba(16, 24, 39, 0.035), transparent 16%),
    linear-gradient(180deg, #ffffff, #f7fbff);
  color: #2a3040;
}

.intro-texture {
  position: absolute;
  inset: 0;
  opacity: 0.42;
  background-image:
    radial-gradient(circle, rgba(16, 24, 39, 0.32) 1px, transparent 1px),
    radial-gradient(circle, rgba(76, 147, 251, 0.28) 1px, transparent 1px);
  background-position: 0 0, 22px 38px;
  background-size: 46px 46px, 70px 70px;
  mask-image: linear-gradient(180deg, transparent, #000 8%, #000 92%, transparent);
}

.intro-layout {
  position: relative;
  display: grid;
  grid-template-columns: 1fr;
  gap: clamp(2rem, 6vw, 4rem);
  align-items: start;
  max-width: 860px;
}

/* ── Copy section ── */
.intro-copy {
  display: grid;
  gap: 2.4rem;
  order: 1;
  opacity: 0;
  transform: translateY(28px);
  transition:
    opacity 0.7s var(--transition-slow) 0.08s,
    transform 0.7s var(--transition-slow) 0.08s;
}

.home-intro.is-revealed .intro-copy {
  opacity: 1;
  transform: translateY(0);
}

.intro-kicker {
  color: var(--primary);
  font-family: var(--mono);
  font-size: 0.84rem;
  letter-spacing: 0.5em;
  text-transform: uppercase;
}

.intro-columns {
  display: grid;
  gap: 1.6rem;
  writing-mode: horizontal-tb;
}

.intro-columns h2 {
  padding: 0 0 1.2rem;
  border-bottom: 1px solid rgba(16, 24, 39, 0.18);
  color: #2d3342;
  font-family: var(--heading);
  font-size: clamp(1.8rem, 3.6vw, 2.9rem);
  font-weight: 800;
  line-height: 1.22;
  letter-spacing: 0;
}

.intro-body {
  display: grid;
  gap: 1.1rem;
  color: #667084;
  font-size: clamp(0.95rem, 1.4vw, 1.08rem);
  line-height: 1.9;
}

/* ── Mascot ── */
.intro-mascot-wrap {
  position: relative;
  display: grid;
  place-items: center;
  order: 2;
  min-height: 320px;
  opacity: 0;
  transform: translateY(28px);
  transition:
    opacity 0.7s var(--transition-slow) 0.22s,
    transform 0.7s var(--transition-slow) 0.22s;
}

.home-intro.is-revealed .intro-mascot-wrap {
  opacity: 1;
  transform: translateY(0);
  animation: mascotStep 3.4s steps(2, end) infinite 0.9s;
}

.intro-mascot {
  max-height: min(55vh, 520px);
  object-fit: contain;
  image-rendering: pixelated;
  filter: drop-shadow(0 28px 0 rgba(16, 24, 39, 0.08));
}

@keyframes mascotStep {
  0%,
  100% {
    transform: translateY(0);
  }

  50% {
    transform: translateY(-10px);
  }
}

@media (min-width: 860px) {
  .intro-layout {
    grid-template-columns: minmax(0, 1fr) minmax(220px, 340px);
    max-width: none;
  }

  .intro-copy {
    order: 1;
  }

  .intro-mascot-wrap {
    order: 2;
    position: sticky;
    top: clamp(6rem, 12vh, 8rem);
    align-self: start;
  }
}
</style>
