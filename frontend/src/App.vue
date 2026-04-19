<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { RouterLink, RouterView } from 'vue-router'
import { storeToRefs } from 'pinia'

import { usePreloader } from '@/composables/usePreloader'
import { siteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'

const menuOpen = ref(false)
const isLoading = ref(true)

const { initPreloading } = usePreloader()
const serverStatus = useServerStatusStore()
const { status } = storeToRefs(serverStatus)

const shellStatus = computed(() => (status.value?.java?.online ? 'Live' : 'Standby'))
const shellPlayers = computed(() => status.value?.java?.players ?? 0)

function closeMenu(): void {
  menuOpen.value = false
}

function toggleMenu(): void {
  menuOpen.value = !menuOpen.value
}

onMounted(async () => {
  await Promise.allSettled([serverStatus.refresh(), initPreloading()])
  isLoading.value = false
  serverStatus.startPolling()
})

onUnmounted(() => {
  serverStatus.stopPolling()
})
</script>

<template>
  <div class="app-shell">
    <div :class="['loader-overlay', { hidden: !isLoading }]">
      <div class="loader-mark glass-card">
        <div class="loader-topline">
          <span class="hud-kicker">{{ siteContent.app.loader.kicker }}</span>
          <span class="loader-status">Boot sequence</span>
        </div>
        <div class="loader-grid">
          <span></span>
          <span></span>
          <span></span>
          <span></span>
        </div>
        <h2 class="loader-title">{{ siteContent.app.loader.title }}</h2>
        <div class="loader-scanline"></div>
      </div>
    </div>

    <header class="top-bar">
      <nav class="container nav-frame">
        <div class="brand-cluster">
          <RouterLink class="brand-lockup" to="/" @click="closeMenu">
            <span class="brand-text">
              <span class="brand-text-mtn">MTN</span><span class="brand-text-etwork">etwork</span>
            </span>
          </RouterLink>

          <div class="shell-readout">
            <span class="hud-chip">
              <span class="status-dot" :class="{ active: shellStatus === 'Live' }"></span>
              {{ shellStatus }}
            </span>
            <span class="hud-chip">{{ shellPlayers }} online</span>
          </div>
        </div>

        <button class="menu-toggle" type="button" aria-label="Toggle navigation" @click="toggleMenu">
          <span :class="{ active: menuOpen }"></span>
          <span :class="{ active: menuOpen }"></span>
        </button>

        <div :class="['nav-links', { open: menuOpen }]">
          <template v-for="item in siteContent.app.nav" :key="item.label">
            <a
              v-if="item.external"
              :href="item.to"
              class="nav-link nav-link-external"
              @click="closeMenu"
            >
              {{ item.label }}
            </a>
            <RouterLink v-else :to="item.to" class="nav-link" @click="closeMenu">
              {{ item.label }}
            </RouterLink>
          </template>
        </div>
      </nav>
    </header>

    <main class="main-content">
      <RouterView v-slot="{ Component, route }">
        <Transition name="route-swap" mode="out-in">
          <component :is="Component" :key="route.fullPath" class="route-layer" />
        </Transition>
      </RouterView>
    </main>

  </div>
</template>

<style scoped>
.app-shell {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.loader-overlay {
  position: fixed;
  inset: 0;
  z-index: 200;
  display: grid;
  place-items: center;
  background:
    radial-gradient(circle at top, rgba(91, 113, 246, 0.22), transparent 30%),
    linear-gradient(180deg, rgba(5, 7, 13, 0.98), rgba(5, 7, 13, 1));
  transition: opacity 0.45s ease;
}

.loader-overlay.hidden {
  opacity: 0;
  pointer-events: none;
}

.loader-mark {
  width: min(420px, calc(100vw - 2rem));
  display: grid;
  gap: 1rem;
  text-align: center;
}

.loader-topline {
  display: flex;
  justify-content: space-between;
  gap: 1rem;
  align-items: center;
}

.loader-status {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.72rem;
  letter-spacing: 0.18em;
  text-transform: uppercase;
}

.loader-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 0.75rem;
}

.loader-grid span {
  height: 24px;
  border-radius: 8px;
  background: rgba(255, 248, 234, 0.08);
  border: 1px solid rgba(255, 248, 234, 0.12);
  animation: pulse-map 1.2s ease-in-out infinite alternate;
}

.loader-grid span:nth-child(2) {
  animation-delay: 0.15s;
}

.loader-grid span:nth-child(3) {
  animation-delay: 0.3s;
}

.loader-grid span:nth-child(4) {
  animation-delay: 0.45s;
}

.loader-title {
  font-size: clamp(2rem, 5vw, 3.2rem);
}

.loader-scanline {
  height: 2px;
  background: linear-gradient(90deg, transparent, rgba(121, 183, 255, 0.8), transparent);
  animation: loader-sweep 1.8s linear infinite;
}

.top-bar {
  position: sticky;
  top: 0;
  z-index: 80;
  padding: 0.65rem 0 0;
  backdrop-filter: blur(14px);
}

.nav-frame {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  padding: 0.88rem 1.05rem;
  border: 1px solid var(--glass-border);
  border-radius: 28px;
  background:
    linear-gradient(180deg, rgba(255, 255, 255, 0.04), transparent 22%),
    rgba(7, 12, 22, 0.88);
  box-shadow: 0 20px 48px rgba(0, 0, 0, 0.34);
}

.brand-cluster {
  display: flex;
  align-items: center;
  gap: 0.8rem;
  min-width: 0;
}

.brand-lockup {
  display: inline-flex;
  align-items: center;
  min-width: 0;
}

.brand-text {
  display: inline-flex;
  align-items: center;
  font-family: var(--display);
  font-size: clamp(1.2rem, 2vw, 1.42rem);
  font-weight: 700;
  letter-spacing: -0.04em;
  line-height: 1;
}

.brand-text-mtn {
  color: var(--primary);
}

.brand-text-etwork {
  color: var(--text-strong);
}

.shell-readout {
  display: flex;
  flex-wrap: wrap;
  gap: 0.45rem;
}

.status-dot {
  width: 0.55rem;
  height: 0.55rem;
  border-radius: 50%;
  background: rgba(154, 174, 203, 0.6);
}

.status-dot.active {
  background: var(--success);
  box-shadow: 0 0 0 6px rgba(93, 226, 164, 0.14);
}

.nav-links {
  display: flex;
  align-items: center;
  gap: 0.2rem;
}

.nav-link {
  padding: 0.56rem 0.82rem;
  border-radius: 999px;
  color: var(--text-muted);
  font-size: 0.78rem;
  font-family: var(--mono);
  letter-spacing: 0.08em;
  text-transform: uppercase;
  transition:
    background var(--transition-fast),
    border-color var(--transition-fast),
    color var(--transition-fast);
  border: 1px solid transparent;
}

.nav-link:hover,
.nav-link.router-link-active {
  color: var(--text-strong);
  border-color: var(--glass-border-strong);
  background: rgba(248, 251, 255, 0.06);
}

.nav-link-external {
  border-color: var(--glass-border);
}

.menu-toggle {
  display: none;
  width: 44px;
  height: 44px;
  border: 1px solid var(--glass-border);
  border-radius: 999px;
  background: transparent;
  position: relative;
}

.menu-toggle span {
  position: absolute;
  left: 12px;
  right: 12px;
  height: 2px;
  background: var(--text-main);
  transition: transform var(--transition-fast), top var(--transition-fast);
}

.menu-toggle span:first-child {
  top: 17px;
}

.menu-toggle span:last-child {
  top: 25px;
}

.menu-toggle span.active:first-child {
  top: 21px;
  transform: rotate(45deg);
}

.menu-toggle span.active:last-child {
  top: 21px;
  transform: rotate(-45deg);
}

.main-content {
  flex: 1;
  position: relative;
}

@keyframes pulse-map {
  from {
    transform: translateY(0);
    opacity: 0.65;
  }

  to {
    transform: translateY(-6px);
    opacity: 1;
    box-shadow: 0 10px 18px rgba(91, 113, 246, 0.24);
  }
}

@keyframes loader-sweep {
  from {
    transform: translateX(-25%);
    opacity: 0.35;
  }

  50% {
    opacity: 1;
  }

  to {
    transform: translateX(25%);
    opacity: 0.35;
  }
}

@media (max-width: 980px) {
  .menu-toggle {
    display: inline-block;
  }

  .shell-readout {
    display: none;
  }

  .nav-links {
    position: fixed;
    top: 5.5rem;
    right: 1rem;
    width: min(340px, calc(100vw - 2rem));
    flex-direction: column;
    align-items: stretch;
    padding: 1rem;
    border: 1px solid var(--glass-border);
    border-radius: 24px;
    background: rgba(7, 12, 22, 0.96);
    box-shadow: 0 30px 60px rgba(0, 0, 0, 0.34);
    opacity: 0;
    pointer-events: none;
    transform: translateY(-10px);
    transition: opacity var(--transition-fast), transform var(--transition-fast);
  }

  .nav-links.open {
    opacity: 1;
    pointer-events: auto;
    transform: translateY(0);
  }

  .nav-link {
    width: 100%;
  }
}

@media (max-width: 640px) {
  .loader-topline {
    flex-direction: column;
    align-items: flex-start;
  }

  .brand-cluster {
    min-width: 0;
  }
}
</style>
