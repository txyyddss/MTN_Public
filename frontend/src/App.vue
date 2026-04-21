<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref, watch } from 'vue'
import { RouterLink, RouterView, useRoute } from 'vue-router'
import { storeToRefs } from 'pinia'

import { usePreloader } from '@/composables/usePreloader'
import { siteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'

const menuOpen = ref(false)
const isLoading = ref(true)
const route = useRoute()

const { initPreloading } = usePreloader()
const serverStatus = useServerStatusStore()
const { status } = storeToRefs(serverStatus)

const shellStatus = computed(() => (status.value?.java?.online ? 'Live' : 'Standby'))
const shellPlayers = computed(() => status.value?.java?.players ?? 0)
const shellStatusCopy = computed(() =>
  status.value?.java?.online ? 'Telemetry refreshing every 5 seconds.' : 'Awaiting the next live heartbeat.'
)
const footerYear = new Date().getFullYear()

function closeMenu(): void {
  menuOpen.value = false
}

function toggleMenu(): void {
  menuOpen.value = !menuOpen.value
}

watch(
  () => route.fullPath,
  () => {
    closeMenu()
  }
)

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
      <div class="top-bar-glow"></div>
      <nav class="container nav-frame">
        <div class="nav-brand-block">
          <RouterLink class="brand-lockup" to="/" @click="closeMenu">
            <span class="brand-text">
              <span class="brand-text-mtn">MTN</span><span class="brand-text-etwork">etwork</span>
            </span>
          </RouterLink>
          <span class="brand-subline">Survival atlas</span>
        </div>

        <div class="nav-center">
          <div :class="['nav-links', { open: menuOpen }]">
            <template v-for="item in siteContent.app.nav" :key="item.label">
              <a
                v-if="item.external"
                :href="item.to"
                :class="['nav-link', 'nav-link-external', { 'nav-link-emphasis': item.emphasize }]"
                @click="closeMenu"
              >
                {{ item.label }}
              </a>
              <RouterLink v-else :to="item.to" class="nav-link" @click="closeMenu">
                {{ item.label }}
              </RouterLink>
            </template>
          </div>
        </div>

        <div class="shell-readout">
          <span class="hud-chip">
            <span class="status-dot" :class="{ active: shellStatus === 'Live' }"></span>
            {{ shellStatus }}
          </span>
          <span class="hud-chip">{{ shellPlayers }} online</span>
        </div>

        <button class="menu-toggle" type="button" aria-label="Toggle navigation" @click="toggleMenu">
          <span :class="{ active: menuOpen }"></span>
          <span :class="{ active: menuOpen }"></span>
        </button>
      </nav>
    </header>

    <main class="main-content">
      <RouterView v-slot="{ Component, route }">
        <Transition name="route-swap" mode="out-in">
          <component :is="Component" :key="route.fullPath" class="route-layer" />
        </Transition>
      </RouterView>
    </main>

    <footer class="site-footer">
      <div class="container footer-frame">
        <div class="footer-copy-block">
          <span class="hud-kicker">Community-run survival</span>
          <h2 class="footer-title">A long-term world with visible public history.</h2>
          <p class="footer-copy">
            Java and Bedrock share one world, one record, and one operating surface.
          </p>
        </div>

        <div class="footer-nav-block">
          <span class="footer-label">Navigate</span>
          <div class="footer-links">
            <template v-for="item in siteContent.app.nav" :key="item.label">
              <a v-if="item.external" :href="item.to" class="footer-link">
                {{ item.label }}
              </a>
              <RouterLink v-else :to="item.to" class="footer-link">
                {{ item.label }}
              </RouterLink>
            </template>
          </div>
        </div>

        <div class="footer-status-card glass-card">
          <span class="hud-kicker">Live surface</span>
          <div class="footer-status-line">
            <strong>{{ shellStatus }}</strong>
            <span>{{ shellPlayers }} online</span>
          </div>
          <p class="footer-status-copy">{{ shellStatusCopy }}</p>
        </div>
      </div>

      <div class="container footer-base">
        <span>© {{ footerYear }} MTNetwork</span>
        <span>Built for the public archive.</span>
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
    radial-gradient(circle at 50% 18%, rgba(76, 147, 251, 0.18), transparent 30%),
    linear-gradient(180deg, rgba(10, 11, 14, 0.96), rgba(0, 0, 0, 1));
  transition: opacity 0.4s ease;
}

.loader-overlay.hidden {
  opacity: 0;
  pointer-events: none;
}

.loader-mark {
  width: min(430px, calc(100vw - 2rem));
  display: grid;
  gap: 1.1rem;
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
  font-size: 0.7rem;
  letter-spacing: 0.16em;
  text-transform: uppercase;
}

.loader-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 0.6rem;
}

.loader-grid span {
  height: 18px;
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.06);
  border: 1px solid rgba(255, 255, 255, 0.08);
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
  font-size: clamp(2rem, 5vw, 3rem);
  font-weight: 600;
}

.loader-scanline {
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(141, 184, 255, 0.85), transparent);
  animation: loader-sweep 1.8s linear infinite;
}

.top-bar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 80;
  background: rgba(0, 0, 0, 0.58);
  backdrop-filter: saturate(180%) blur(18px);
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}

.top-bar-glow {
  position: absolute;
  inset: auto 0 0;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(76, 147, 251, 0.5), transparent);
  opacity: 0.8;
  pointer-events: none;
}

.nav-frame {
  display: grid;
  grid-template-columns: auto minmax(0, 1fr) auto auto;
  align-items: center;
  gap: 1rem;
  min-height: 62px;
  padding: 0.55rem 0;
}

.nav-brand-block {
  display: grid;
  gap: 0.18rem;
  min-width: 0;
}

.brand-lockup {
  display: inline-flex;
  align-items: center;
  min-width: 0;
  flex-shrink: 0;
}

.brand-text {
  display: inline-flex;
  align-items: center;
  font-family: var(--display);
  font-size: clamp(1.1rem, 2vw, 1.32rem);
  font-weight: 600;
  letter-spacing: -0.04em;
  line-height: 1;
}

.brand-text-mtn {
  color: var(--primary);
}

.brand-text-etwork {
  color: var(--text-strong);
}

.brand-subline {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.63rem;
  letter-spacing: 0.16em;
  text-transform: uppercase;
}

.nav-center {
  display: flex;
  justify-content: center;
  min-width: 0;
}

.shell-readout {
  display: flex;
  flex-wrap: wrap;
  gap: 0.45rem;
  justify-content: flex-end;
  flex-shrink: 0;
}

.status-dot {
  width: 0.5rem;
  height: 0.5rem;
  border-radius: 50%;
  background: rgba(182, 190, 203, 0.8);
}

.status-dot.active {
  background: var(--success);
  box-shadow: 0 0 0 5px rgba(131, 211, 167, 0.12);
}

.nav-links {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.15rem;
  margin: 0;
  padding: 0.28rem;
  border-radius: 999px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  background: rgba(255, 255, 255, 0.025);
  box-shadow: 0 14px 30px rgba(0, 0, 0, 0.18);
}

.nav-link {
  padding: 0.55rem 0.85rem;
  border-radius: 10px;
  color: var(--text-muted);
  font-size: 0.88rem;
  font-family: var(--sans);
  font-weight: 500;
  letter-spacing: -0.01em;
  transition:
    background var(--transition-fast),
    border-color var(--transition-fast),
    color var(--transition-fast);
  border: 1px solid transparent;
}

.nav-link:hover,
.nav-link.router-link-active {
  color: var(--text-strong);
  border-color: rgba(255, 255, 255, 0.08);
  background: rgba(255, 255, 255, 0.055);
}

.nav-link-external {
  color: var(--accent);
}

.nav-link-emphasis {
  color: #fff4c2;
  background: linear-gradient(135deg, rgba(182, 126, 23, 0.34), rgba(255, 214, 102, 0.18));
  border-color: rgba(255, 214, 102, 0.28);
  box-shadow: 0 10px 28px rgba(182, 126, 23, 0.2);
}

.nav-link-emphasis:hover {
  color: #fff9de;
  border-color: rgba(255, 230, 148, 0.42);
  background: linear-gradient(135deg, rgba(204, 144, 29, 0.42), rgba(255, 214, 102, 0.22));
}

.menu-toggle {
  display: none;
  width: 42px;
  height: 42px;
  border: 1px solid rgba(255, 255, 255, 0.12);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.04);
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
  padding-top: 4.35rem;
}

.site-footer {
  position: relative;
  padding: 1.4rem 0 2.4rem;
  border-top: 1px solid rgba(255, 255, 255, 0.08);
  background:
    linear-gradient(180deg, rgba(255, 255, 255, 0.02), transparent 45%),
    rgba(4, 5, 7, 0.7);
}

.site-footer::before {
  content: '';
  position: absolute;
  inset: 0 0 auto;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(76, 147, 251, 0.35), transparent);
  opacity: 0.9;
}

.footer-frame {
  display: grid;
  grid-template-columns: minmax(0, 1.35fr) minmax(0, 0.85fr) minmax(280px, 0.8fr);
  gap: 1rem;
  align-items: start;
  padding-bottom: 1.1rem;
}

.footer-copy-block,
.footer-nav-block {
  display: grid;
  gap: 0.8rem;
}

.footer-title {
  font-size: clamp(1.85rem, 4vw, 3rem);
  max-width: 13ch;
}

.footer-copy,
.footer-status-copy {
  color: var(--text-muted);
}

.footer-label {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.72rem;
  letter-spacing: 0.16em;
  text-transform: uppercase;
}

.footer-links {
  display: flex;
  flex-wrap: wrap;
  gap: 0.6rem;
}

.footer-link {
  display: inline-flex;
  align-items: center;
  min-height: 2.7rem;
  padding: 0.72rem 0.95rem;
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  background: rgba(255, 255, 255, 0.03);
  color: var(--text-muted);
  transition:
    transform var(--transition-fast),
    border-color var(--transition-fast),
    background var(--transition-fast),
    color var(--transition-fast);
}

.footer-link:hover,
.footer-link.router-link-active {
  color: var(--text-strong);
  border-color: rgba(255, 255, 255, 0.16);
  background: rgba(255, 255, 255, 0.06);
  transform: translateY(-1px);
}

.footer-status-card {
  display: grid;
  gap: 0.85rem;
  min-height: 100%;
}

.footer-status-line {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  color: var(--text-muted);
}

.footer-status-line strong {
  color: var(--text-strong);
  font-size: 1.4rem;
}

.footer-base {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  padding-top: 1.1rem;
  border-top: 1px solid rgba(255, 255, 255, 0.06);
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.68rem;
  letter-spacing: 0.1em;
  text-transform: uppercase;
}

@keyframes pulse-map {
  from {
    transform: translateY(0);
    opacity: 0.65;
  }

  to {
    transform: translateY(-4px);
    opacity: 1;
    box-shadow: 0 8px 18px rgba(76, 147, 251, 0.18);
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
  .nav-frame {
    grid-template-columns: minmax(0, 1fr) auto;
  }

  .menu-toggle {
    display: inline-grid;
    place-items: center;
  }

  .shell-readout {
    display: none;
  }

  .nav-center {
    justify-content: flex-end;
  }

  .nav-links {
    position: fixed;
    top: 4.45rem;
    left: 1rem;
    right: 1rem;
    width: auto;
    flex-direction: column;
    align-items: stretch;
    margin: 0;
    padding: 0.85rem;
    border: 1px solid rgba(255, 255, 255, 0.08);
    border-radius: 20px;
    background: rgba(14, 15, 18, 0.94);
    box-shadow: 0 22px 48px rgba(0, 0, 0, 0.38);
    opacity: 0;
    pointer-events: none;
    transform: translateY(-8px);
    transition: opacity var(--transition-fast), transform var(--transition-fast);
  }

  .nav-links.open {
    opacity: 1;
    pointer-events: auto;
    transform: translateY(0);
  }

  .nav-link {
    width: 100%;
    padding: 0.8rem 0.95rem;
    border-radius: 14px;
  }

  .footer-frame {
    grid-template-columns: 1fr;
  }

  .footer-status-card {
    min-height: auto;
  }
}

@media (max-width: 640px) {
  .loader-topline {
    flex-direction: column;
    align-items: flex-start;
  }

  .nav-frame {
    min-height: 54px;
  }

  .brand-subline {
    display: none;
  }

  .footer-base {
    flex-direction: column;
    align-items: flex-start;
  }
}
</style>
