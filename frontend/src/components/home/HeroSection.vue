<script setup lang="ts">
import HeroImageCarousel from '@/components/home/HeroImageCarousel.vue'
import { useSiteContent } from '@/content/siteContent'

const siteContent = useSiteContent()
</script>

<template>
  <section class="home-hero">
    <HeroImageCarousel
      :image-alt="siteContent.home.hero.imageAlt"
      :previous-label="siteContent.gallery.previousImage"
      :next-label="siteContent.gallery.nextImage"
      :frame-label="siteContent.gallery.frameLabel"
    />
    <div class="home-hero-shade" aria-hidden="true"></div>

    <div class="home-hero-center">
      <div class="home-hero-mark" aria-hidden="true">MTN</div>
      <p class="home-hero-tagline">{{ siteContent.home.hero.tagline }}</p>
      <p class="home-hero-project">{{ siteContent.home.hero.projectLine }}</p>
      <div class="hero-jump-actions">
        <a class="hero-jump-btn" href="#details">{{ siteContent.home.hero.jumpDetailsCta }}</a>
        <a class="hero-jump-btn hero-jump-btn--outline" href="https://docs.mcmtn.net/" target="_blank" rel="noopener noreferrer">{{ siteContent.home.hero.jumpDocsCta }}</a>
        <a class="hero-jump-btn hero-jump-btn--outline" href="#join-mtn">{{ siteContent.home.hero.jumpJoinCta }}</a>
      </div>
    </div>

    <a class="home-hero-cue" href="#intro-section" :aria-label="siteContent.home.hero.scrollCta">
      <span></span>
    </a>
  </section>
</template>

<style scoped>
.home-hero {
  position: relative;
  display: grid;
  place-items: center;
  min-height: 100svh;
  overflow: hidden;
  background: #000000;
  color: #ffffff;
}

.home-hero-shade {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
}

.home-hero-shade {
  z-index: 1;
  pointer-events: none;
  background:
    linear-gradient(90deg, rgba(0, 2, 8, 0.8), rgba(2, 8, 20, 0.24) 46%, rgba(0, 2, 8, 0.76)),
    radial-gradient(circle at 50% 42%, rgba(var(--secondary-rgb), 0.18), transparent 30%);
}

.home-hero-center {
  position: relative;
  z-index: 2;
  display: grid;
  align-content: center;
  justify-items: center;
  gap: 1.05rem;
  width: min(1120px, calc(100vw - 2rem));
  min-height: 100svh;
  margin: 0 auto;
  text-align: center;
  place-self: center;
  animation: heroReveal 1.1s cubic-bezier(0.16, 1, 0.3, 1) both;
}

.home-hero-mark {
  font-family: var(--display);
  font-size: clamp(4.6rem, 14vw, 10rem);
  font-weight: 900;
  letter-spacing: 0.22em;
  line-height: 0.72;
  text-indent: 0.22em;
  color: rgba(var(--primary-rgb), 0.98);
  text-shadow:
    0 24px 56px rgba(var(--secondary-rgb), 0.32),
    0 0 44px rgba(var(--primary-rgb), 0.24),
    0 2px 0 rgba(255, 255, 255, 0.08);
  text-transform: uppercase;
  margin-bottom: 0.05rem;
  animation: heroReveal 1.1s cubic-bezier(0.16, 1, 0.3, 1) 0.08s both;
}

.home-hero-project {
  font-family: var(--mono);
  text-transform: uppercase;
}

.home-hero-tagline {
  display: flex;
  align-items: center;
  gap: 1rem;
  color: #ffffff;
  font-size: clamp(0.92rem, 2vw, 1.25rem);
  font-weight: 700;
  letter-spacing: 0.3em;
  animation: heroReveal 0.9s cubic-bezier(0.16, 1, 0.3, 1) 0.26s both;
}

.home-hero-tagline::before,
.home-hero-tagline::after {
  content: '';
  width: min(14vw, 8rem);
  height: 1px;
  background: rgba(255, 255, 255, 0.42);
}

.home-hero-project {
  color: var(--accent);
  font-size: 0.82rem;
  letter-spacing: 0.28em;
  animation: heroReveal 0.9s cubic-bezier(0.16, 1, 0.3, 1) 0.3s both;
}

.hero-jump-actions {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 0.75rem;
  margin-top: 0.5rem;
  animation: heroReveal 0.9s cubic-bezier(0.16, 1, 0.3, 1) 0.38s both;
}

.hero-jump-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-height: 2.6rem;
  padding: 0.6rem 1.5rem;
  border-radius: 999px;
  border: 1px solid rgba(147, 197, 253, 0.5);
  background:
    linear-gradient(135deg, rgba(var(--primary-rgb), 0.92), rgba(var(--secondary-rgb), 0.9));
  color: #ffffff;
  font-family: var(--mono);
  font-size: 0.8rem;
  font-weight: 700;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  box-shadow: 0 12px 32px rgba(var(--secondary-rgb), 0.34);
  backdrop-filter: blur(16px);
  transition:
    transform 0.18s cubic-bezier(0.2, 0, 0.2, 1),
    background 0.18s cubic-bezier(0.2, 0, 0.2, 1),
    box-shadow 0.18s cubic-bezier(0.2, 0, 0.2, 1);
}

.hero-jump-btn:hover {
  transform: translateY(-2px);
  background:
    linear-gradient(135deg, rgba(var(--primary-rgb), 0.98), rgba(var(--secondary-rgb), 0.96));
  box-shadow: 0 16px 40px rgba(var(--secondary-rgb), 0.42);
}

.hero-jump-btn:active {
  transform: translateY(0) scale(0.985);
}

.hero-jump-btn--outline {
  background: rgba(2, 8, 20, 0.42);
  border-color: rgba(147, 197, 253, 0.28);
  box-shadow: 0 12px 28px rgba(0, 0, 0, 0.18);
}

.hero-jump-btn--outline:hover {
  background: rgba(255, 255, 255, 0.12);
  border-color: rgba(255, 255, 255, 0.46);
  box-shadow: none;
}

.home-hero-cue {
  position: absolute;
  left: 50%;
  bottom: 2rem;
  z-index: 3;
  width: 3.6rem;
  height: 3.6rem;
  transform: translateX(-50%);
}

.home-hero-cue span {
  display: block;
  width: 2.2rem;
  height: 2.2rem;
  margin: 0 auto;
  border-right: 3px solid rgba(255, 255, 255, 0.88);
  border-bottom: 3px solid rgba(255, 255, 255, 0.88);
  transform: rotate(45deg);
  animation: cueBounce 1.8s ease-in-out infinite;
}

@keyframes cueBounce {
  0%,
  100% {
    transform: translateY(-4px) rotate(45deg);
  }

  50% {
    transform: translateY(8px) rotate(45deg);
  }
}

@keyframes heroReveal {
  from {
    opacity: 0;
    transform: translateY(22px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@media (max-width: 680px) {
  .home-hero-center {
    width: min(100%, calc(100vw - 2.5rem));
  }

  .home-hero-mark {
    font-size: clamp(3.1rem, 19vw, 5.2rem);
    letter-spacing: 0.18em;
    text-indent: 0.18em;
  }

  .home-hero-tagline {
    gap: 0.6rem;
    font-size: 0.88rem;
    letter-spacing: 0.1em;
  }

  .home-hero-tagline::before,
  .home-hero-tagline::after {
    width: 2.4rem;
  }

  .home-hero-project {
    font-size: 0.68rem;
    letter-spacing: 0.2em;
  }

  .hero-jump-actions {
    flex-direction: column;
    width: 100%;
  }

  .hero-jump-btn {
    width: 100%;
  }
}
</style>
