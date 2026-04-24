<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref, watch } from 'vue'
import { RouterLink, RouterView, useRoute } from 'vue-router'
import { storeToRefs } from 'pinia'

import LocaleSwitcher from '@/components/LocaleSwitcher.vue'
import { usePreloader } from '@/composables/usePreloader'
import { formatPlayerCount, useSiteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'

const menuOpen = ref(false)
const route = useRoute()
const siteContent = useSiteContent()

const { initPreloading } = usePreloader()
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

onMounted(() => {
  void serverStatus.refresh()
  void initPreloading()
  serverStatus.startPolling()
})

onUnmounted(() => {
  serverStatus.stopPolling()
})
</script>

<template>
  <div class="app-shell">
    <header class="top-bar">
      <div class="top-bar-glow"></div>
      <nav class="container nav-frame">
        <div class="nav-brand-block">
          <RouterLink class="brand-lockup" to="/" @click="closeMenu">
            <span class="brand-text">
              <span class="brand-text-mtn">MTN</span><span class="brand-text-etwork">etwork</span>
            </span>
          </RouterLink>
          <span class="brand-subline">{{ siteContent.app.brandSubline }}</span>
        </div>

        <div class="nav-center">
          <div :class="['nav-links', { open: menuOpen }]">
            <template v-for="item in siteContent.app.nav" :key="item.id">
              <a
                v-if="item.external"
                :href="item.to"
                :class="['nav-link', 'nav-link-external', 'action-inline', 'action-press', { 'nav-link-emphasis': item.emphasize }]"
                @click="closeMenu"
              >
                {{ item.label }}
              </a>
              <RouterLink v-else :to="item.to" class="nav-link action-inline action-press" @click="closeMenu">
                {{ item.label }}
              </RouterLink>
            </template>

            <LocaleSwitcher mobile class="nav-locale-switch" />
          </div>
        </div>

        <div class="shell-readout">
          <span class="hud-chip">
            <span class="status-dot" :class="{ active: isLive }"></span>
            {{ shellStatus }}
          </span>
          <span class="hud-chip">{{ shellPlayersLabel }}</span>
          <LocaleSwitcher />
        </div>

        <button class="menu-toggle action-inline action-press" type="button" :aria-label="siteContent.app.menuToggleAria" @click="toggleMenu">
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


  </div>
</template>

<style scoped>
.app-shell {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
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
  border: 1px solid var(--control-border);
  background: var(--control-surface);
  box-shadow: 0 14px 30px rgba(0, 0, 0, 0.18);
}

.nav-link {
  padding: 0.55rem 0.85rem;
  border-radius: 10px;
  color: var(--control-text);
  font-size: 0.88rem;
  font-family: var(--sans);
  font-weight: 500;
  letter-spacing: -0.01em;
  transition:
    transform var(--transition-fast),
    background var(--transition-fast),
    border-color var(--transition-fast),
    color var(--transition-fast),
    box-shadow var(--transition-fast);
  border: 1px solid transparent;
}

.nav-link:hover,
.nav-link.router-link-active {
  color: var(--control-text-active);
  border-color: var(--control-border-active);
  background: var(--control-bg-active);
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
  border: 1px solid var(--control-border-hover);
  border-radius: 999px;
  background: var(--control-surface);
  position: relative;
  transition:
    transform var(--transition-fast),
    border-color var(--transition-fast),
    background var(--transition-fast),
    box-shadow var(--transition-fast);
}

.menu-toggle span {
  position: absolute;
  left: 12px;
  right: 12px;
  height: 2px;
  background: var(--text-main);
  transition:
    transform var(--transition-fast),
    top var(--transition-fast),
    opacity var(--transition-fast);
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

.nav-locale-switch {
  display: none;
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
    gap: 0.35rem;
    margin: 0;
    padding: 0.85rem;
    border: 1px solid var(--control-border);
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

  .nav-locale-switch {
    display: flex;
  }

}

@media (max-width: 640px) {
  .nav-frame {
    min-height: 54px;
  }

  .brand-subline {
    display: none;
  }


}
</style>
