<script setup lang="ts">
import serverMascot from '@/assets/server-mascot.png'
import { useSiteContent } from '@/content/siteContent'

const siteContent = useSiteContent()
</script>

<template>
  <aside class="hero-aside">
    <article class="glass-card hero-aside-card animate-entry delay-200">
      <div class="hero-aside-copy">
        <span class="section-kicker">{{ siteContent.home.heroAside.kicker }}</span>
        <h2 class="hero-aside-title">{{ siteContent.home.heroAside.title }}</h2>
      </div>

      <div class="mascot-stage">
        <div class="voxel-field" aria-hidden="true">
          <span></span>
          <span></span>
          <span></span>
          <span></span>
          <span></span>
        </div>
        <img :src="serverMascot" :alt="siteContent.home.heroAside.imageAlt" class="hero-mascot" />
      </div>

      <div class="hero-aside-status">
        <span>{{ siteContent.home.heroAside.statusLabel }}</span>
        <strong>{{ siteContent.home.heroAside.statusValue }}</strong>
      </div>

      <div class="hero-chip-row">
        <span v-for="chip in siteContent.home.heroAside.chips" :key="chip" class="hud-chip">
          {{ chip }}
        </span>
      </div>
    </article>
  </aside>
</template>

<style scoped>
.hero-aside {
  display: grid;
  min-height: 100%;
}

.hero-aside-card {
  position: relative;
  display: grid;
  grid-template-rows: auto minmax(0, 1fr) auto auto;
  gap: 1rem;
  min-height: 0;
  isolation: isolate;
}

.hero-aside-title {
  font-size: clamp(1.55rem, 3vw, 2.35rem);
  max-width: 9ch;
}

.mascot-stage {
  position: relative;
  display: grid;
  place-items: end center;
  min-height: 460px;
  overflow: hidden;
}

.mascot-stage::before {
  content: '';
  position: absolute;
  inset: 12% 8% 8%;
  z-index: -1;
  border-radius: 34px;
  background:
    radial-gradient(circle at 48% 28%, rgba(76, 147, 251, 0.28), transparent 42%),
    linear-gradient(180deg, rgba(76, 147, 251, 0.16), rgba(255, 255, 255, 0.02));
  filter: blur(1px);
}

.hero-mascot {
  width: min(100%, 350px);
  max-height: 450px;
  object-fit: contain;
  filter: drop-shadow(0 28px 44px rgba(0, 0, 0, 0.46));
  animation: mascotHover 5.6s ease-in-out infinite;
}

.voxel-field {
  position: absolute;
  inset: 0;
  pointer-events: none;
}

.voxel-field span {
  position: absolute;
  width: 0.72rem;
  height: 0.72rem;
  border: 1px solid rgba(141, 184, 255, 0.32);
  background: rgba(76, 147, 251, 0.14);
  box-shadow: 0 0 24px rgba(76, 147, 251, 0.24);
  animation: voxelDrift 7s ease-in-out infinite;
}

.voxel-field span:nth-child(1) {
  top: 14%;
  left: 18%;
}

.voxel-field span:nth-child(2) {
  top: 28%;
  right: 14%;
  animation-delay: -1.3s;
}

.voxel-field span:nth-child(3) {
  bottom: 24%;
  left: 11%;
  animation-delay: -2.2s;
}

.voxel-field span:nth-child(4) {
  right: 24%;
  bottom: 12%;
  animation-delay: -3.1s;
}

.voxel-field span:nth-child(5) {
  left: 46%;
  top: 8%;
  animation-delay: -4s;
}

.hero-aside-status {
  display: grid;
  gap: 0.35rem;
  padding: 0.95rem 1rem;
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 16px;
  background: rgba(255, 255, 255, 0.032);
}

.hero-aside-status span {
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.68rem;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.hero-aside-status strong {
  color: var(--text-strong);
  font-size: 1rem;
}

.hero-chip-row {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

@keyframes mascotHover {
  0%,
  100% {
    transform: translateY(0);
  }

  50% {
    transform: translateY(-10px);
  }
}

@keyframes voxelDrift {
  0%,
  100% {
    opacity: 0.36;
    transform: translate3d(0, 0, 0) rotate(0deg);
  }

  50% {
    opacity: 0.85;
    transform: translate3d(8px, -12px, 0) rotate(45deg);
  }
}

@media (max-width: 640px) {
  .mascot-stage {
    min-height: 370px;
  }

  .hero-mascot {
    max-height: 370px;
  }
}
</style>
