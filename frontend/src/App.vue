<script setup lang="ts">
import { computed, onMounted, onUnmounted, shallowRef, watch } from 'vue'
import { RouterLink, RouterView, useRoute } from 'vue-router'
import { storeToRefs } from 'pinia'

import AppPreloader from '@/components/AppPreloader.vue'
import LocaleSwitcher from '@/components/LocaleSwitcher.vue'
import ShellSideControls from '@/components/ShellSideControls.vue'
import { usePreloader } from '@/composables/usePreloader'
import { formatPlayerCount, useSiteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'

const menuOpen = shallowRef(false)
const route = useRoute()
const siteContent = useSiteContent()

const { initPreloading, isReady } = usePreloader()
const serverStatus = useServerStatusStore()
const { status } = storeToRefs(serverStatus)

const isLive = computed(() => Boolean(status.value?.java?.online))
const shellStatus = computed(() =>
  isLive.value ? siteContent.value.serverPanels.operational : siteContent.value.serverPanels.standby
)
const shellPlayers = computed(() => status.value?.java?.players ?? 0)
const shellPlayersLabel = computed(() => formatPlayerCount(shellPlayers.value))

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

watch(menuOpen, (open) => {
  document.body.style.overflow = open ? 'hidden' : ''
})

onMounted(() => {
  void serverStatus.refresh()
  void initPreloading()
  serverStatus.startPolling()
})

onUnmounted(() => {
  document.body.style.overflow = ''
  serverStatus.stopPolling()
})
</script>

<template>
  <div class="app-shell">
    <AppPreloader :ready="isReady" />

    <button
      :class="['menu-trigger', 'action-inline', 'action-press', { active: menuOpen }]"
      type="button"
      :aria-label="siteContent.app.menuToggleAria"
      :aria-expanded="menuOpen"
      @click="toggleMenu"
    >
      <span></span>
      <span></span>
      <span></span>
    </button>

    <Transition name="overlay-fade">
      <button v-if="menuOpen" class="nav-scrim" type="button" :aria-label="siteContent.app.menuToggleAria" @click="closeMenu"></button>
    </Transition>

    <nav :class="['nav-overlay', { active: menuOpen }]" :aria-hidden="!menuOpen">
      <div class="nav-overlay-inner">
        <RouterLink class="brand-lockup" to="/" @click="closeMenu">
          <span class="brand-mark" aria-hidden="true">
            <span></span>
            <span></span>
            <span></span>
            <span></span>
          </span>
          <span>
            <strong>MTNetwork</strong>
            <small>{{ siteContent.app.brandSubline }}</small>
          </span>
        </RouterLink>

        <div class="shell-readout">
          <span class="shell-chip">
            <span class="status-dot" :class="{ active: isLive }"></span>
            {{ shellStatus }}
          </span>
          <span class="shell-chip">{{ shellPlayersLabel }}</span>
        </div>

        <div class="nav-links">
          <template v-for="item in siteContent.app.nav" :key="item.id">
            <a
              v-if="item.external"
              :href="item.to"
              :class="['nav-link', { 'nav-link-emphasis': item.emphasize }]"
              target="_blank"
              rel="noopener noreferrer"
              @click="closeMenu"
            >
              <span>{{ item.label }}</span>
            </a>
            <RouterLink v-else :to="item.to" class="nav-link" @click="closeMenu">
              <span>{{ item.label }}</span>
            </RouterLink>
          </template>
        </div>

        <LocaleSwitcher class="nav-locale-switch" />
      </div>
    </nav>

    <ShellSideControls />

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

.menu-trigger {
  position: fixed;
  top: 1.75rem;
  right: 1.75rem;
  z-index: 120;
  width: 44px;
  height: 34px;
  border: 0;
  background: transparent;
  mix-blend-mode: difference;
}

.menu-trigger span {
  position: absolute;
  left: 0;
  right: 0;
  height: 3px;
  background: #ffffff;
  transition:
    top var(--transition-fast),
    opacity var(--transition-fast),
    transform var(--transition-fast);
}

.menu-trigger span:nth-child(1) {
  top: 3px;
}

.menu-trigger span:nth-child(2) {
  top: 15px;
}

.menu-trigger span:nth-child(3) {
  top: 27px;
}

.menu-trigger.active {
  mix-blend-mode: normal;
}

.menu-trigger.active span {
  background: #ffffff;
}

.menu-trigger.active span:nth-child(1) {
  top: 15px;
  transform: rotate(-45deg);
}

.menu-trigger.active span:nth-child(2) {
  opacity: 0;
}

.menu-trigger.active span:nth-child(3) {
  top: 15px;
  transform: rotate(45deg);
}

.nav-scrim {
  position: fixed;
  inset: 0;
  z-index: 89;
  border: 0;
  background: rgba(0, 0, 0, 0.34);
  backdrop-filter: blur(2px);
}

.nav-overlay {
  position: fixed;
  top: 0;
  right: 0;
  z-index: 100;
  width: min(100vw, 430px);
  height: 100vh;
  padding: 5.5rem 2.2rem 2rem;
  background:
    linear-gradient(180deg, rgba(255, 255, 255, 0.05), transparent 38%),
    #101827;
  color: #ffffff;
  box-shadow: -28px 0 64px rgba(0, 0, 0, 0.34);
  transform: translateX(100%);
  transition: transform 0.5s cubic-bezier(0.075, 0.82, 0.165, 1);
}

.nav-overlay.active {
  transform: translateX(0);
}

.nav-overlay-inner {
  display: grid;
  gap: 2rem;
  min-height: 100%;
  align-content: center;
}

.brand-lockup {
  display: inline-flex;
  align-items: center;
  gap: 0.85rem;
  width: fit-content;
}

.brand-mark {
  display: grid;
  grid-template-columns: repeat(2, 13px);
  gap: 5px;
  transform: rotate(45deg);
}

.brand-mark span {
  width: 13px;
  aspect-ratio: 1;
  background: var(--primary);
}

.brand-mark span:nth-child(3),
.brand-mark span:nth-child(4) {
  background: #83d3a7;
}

.brand-lockup strong,
.brand-lockup small {
  display: block;
}

.brand-lockup strong {
  color: #ffffff;
  font-family: var(--display);
  font-size: 1.35rem;
  font-weight: 700;
}

.brand-lockup small {
  color: rgba(255, 255, 255, 0.58);
  font-family: var(--mono);
  font-size: 0.7rem;
  letter-spacing: 0.16em;
  text-transform: uppercase;
}

.shell-readout {
  display: flex;
  flex-wrap: wrap;
  gap: 0.55rem;
}

.shell-chip {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.72rem;
  border: 1px solid rgba(255, 255, 255, 0.12);
  color: rgba(255, 255, 255, 0.74);
  font-family: var(--mono);
  font-size: 0.72rem;
}

.status-dot {
  width: 0.5rem;
  height: 0.5rem;
  border-radius: 50%;
  background: rgba(182, 190, 203, 0.8);
}

.status-dot.active {
  background: var(--success);
  box-shadow: 0 0 0 5px rgba(131, 211, 167, 0.14);
}

.nav-links {
  display: grid;
  gap: 0.2rem;
}

.nav-link {
  position: relative;
  display: flex;
  align-items: center;
  min-height: 3.5rem;
  color: rgba(255, 255, 255, 0.78);
  font-family: var(--display);
  font-size: clamp(1.35rem, 4vw, 2rem);
  font-weight: 600;
  transition:
    color var(--transition-fast),
    transform var(--transition-fast);
}

.nav-link::before {
  content: '';
  width: 0;
  height: 2px;
  margin-right: 0;
  background: var(--primary);
  transition:
    width var(--transition-fast),
    margin-right var(--transition-fast);
}

.nav-link:hover,
.nav-link.router-link-active {
  color: #ffffff;
  transform: translateX(3px);
}

.nav-link:hover::before,
.nav-link.router-link-active::before {
  width: 2.2rem;
  margin-right: 0.85rem;
}

.nav-link-emphasis {
  color: #d7e8ff;
}

.nav-locale-switch {
  justify-self: start;
}

.main-content {
  flex: 1;
  position: relative;
}

.overlay-fade-enter-active,
.overlay-fade-leave-active {
  transition: opacity 0.2s ease;
}

.overlay-fade-enter-from,
.overlay-fade-leave-to {
  opacity: 0;
}

@media (max-width: 680px) {
  .menu-trigger {
    top: 1rem;
    right: 1rem;
    width: 38px;
  }

  .nav-overlay {
    width: 100vw;
    padding: 4.5rem 1.3rem 1.5rem;
  }
}
</style>
