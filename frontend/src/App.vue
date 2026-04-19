<script setup lang="ts">
import { onMounted, onUnmounted, ref } from 'vue'
import { RouterLink, RouterView } from 'vue-router'

import { usePreloader } from '@/composables/usePreloader'
import { siteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'

const menuOpen = ref(false)
const isLoading = ref(true)

const { initPreloading } = usePreloader()
const serverStatus = useServerStatusStore()

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
      <div class="loader-mark">
        <div class="loader-grid">
          <span></span>
          <span></span>
          <span></span>
          <span></span>
        </div>
        <p class="loader-kicker">{{ siteContent.app.loader.kicker }}</p>
        <h2 class="loader-title">{{ siteContent.app.loader.title }}</h2>
      </div>
    </div>

    <header class="top-bar">
      <nav class="container nav-frame">
        <RouterLink class="brand-lockup" to="/" @click="closeMenu">
          <span class="brand-text">
            <span class="brand-text-mtn">MTN</span><span class="brand-text-etwork">etwork</span>
          </span>
        </RouterLink>

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
      <RouterView />
    </main>

    <footer class="app-footer">
      <div class="container footer-frame">
        <div>
          <p class="footer-brand">{{ siteContent.app.brand }}</p>
          <p class="footer-copy">{{ siteContent.app.footer }}</p>
        </div>
        <p class="footer-meta">(c) 2026 MTNetwork</p>
      </div>
    </footer>
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
  display: grid;
  justify-items: center;
  gap: 1rem;
  text-align: center;
}

.loader-grid {
  display: grid;
  grid-template-columns: repeat(2, 22px);
  gap: 0.75rem;
}

.loader-grid span {
  width: 22px;
  height: 22px;
  border-radius: 6px;
  background: rgba(255, 248, 234, 0.08);
  border: 1px solid rgba(255, 248, 234, 0.12);
  animation: pulse-map 1.2s ease-in-out infinite alternate;
}

.loader-grid span:nth-child(2) {
  animation-delay: 0.2s;
}

.loader-grid span:nth-child(3) {
  animation-delay: 0.35s;
}

.loader-grid span:nth-child(4) {
  animation-delay: 0.5s;
}

.loader-kicker {
  color: var(--primary);
  font-family: var(--mono);
  font-size: 0.76rem;
  letter-spacing: 0.2em;
  text-transform: uppercase;
}

.loader-title {
  font-size: clamp(2rem, 5vw, 3.2rem);
}

.top-bar {
  position: sticky;
  top: 0;
  z-index: 80;
  padding: 1rem 0;
  backdrop-filter: blur(14px);
}

.nav-frame {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  padding: 1rem 1.25rem;
  border: 1px solid var(--glass-border);
  border-radius: 999px;
  background: rgba(7, 10, 18, 0.86);
  box-shadow: 0 20px 48px rgba(0, 0, 0, 0.34);
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

.nav-links {
  display: flex;
  align-items: center;
  gap: 0.35rem;
}

.nav-link {
  padding: 0.55rem 0.9rem;
  border-radius: 999px;
  color: var(--text-muted);
  font-size: 0.9rem;
  transition: background var(--transition-fast), color var(--transition-fast);
}

.nav-link:hover,
.nav-link.router-link-active {
  color: var(--text-strong);
  background: rgba(248, 251, 255, 0.06);
}

.nav-link-external {
  border: 1px solid var(--glass-border);
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
}

.app-footer {
  padding: 0 0 2rem;
}

.footer-frame {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  gap: 1.5rem;
  padding: 1.5rem 0 0;
  border-top: 1px solid var(--glass-border);
}

.footer-brand {
  margin-bottom: 0.35rem;
  color: var(--text-strong);
  font-family: var(--display);
  font-size: 1.2rem;
}

.footer-copy,
.footer-meta {
  color: var(--text-dim);
  font-size: 0.88rem;
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

@media (max-width: 980px) {
  .menu-toggle {
    display: inline-block;
  }

  .nav-links {
    position: fixed;
    top: 5.5rem;
    right: 1rem;
    width: min(320px, calc(100vw - 2rem));
    flex-direction: column;
    align-items: stretch;
    padding: 1rem;
    border: 1px solid var(--glass-border);
    border-radius: 24px;
    background: rgba(7, 10, 18, 0.96);
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

  .footer-frame {
    flex-direction: column;
    align-items: flex-start;
  }
}
</style>
