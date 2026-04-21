<script setup lang="ts">
import { computed, onUnmounted, shallowRef } from 'vue'

import { useServerConnectionRows } from '@/composables/useServerConnectionRows'
import { siteContent } from '@/content/siteContent'
import type { ConnectionResponse } from '@/types/api'

interface Props {
  connection: ConnectionResponse | null
}

const props = defineProps<Props>()

const showAllJava = shallowRef(false)
const copyFeedback = shallowRef<string | null>(null)

const connectionRef = computed(() => props.connection)
const { javaRows, bedrockRows } = useServerConnectionRows(connectionRef, showAllJava)

const connectionSections = computed(() => [
  { title: siteContent.serverPanels.javaTitle, rows: javaRows.value },
  { title: siteContent.serverPanels.bedrockTitle, rows: bedrockRows.value }
])

const canToggleDetailedView = computed(() => Boolean(props.connection?.connection))

let feedbackTimer: ReturnType<typeof window.setTimeout> | null = null

async function copyToClipboard(text: string, label: string): Promise<void> {
  if (!navigator.clipboard) {
    return
  }

  try {
    await navigator.clipboard.writeText(text)
    copyFeedback.value = `${siteContent.serverPanels.copyPrefix} ${label}`

    if (feedbackTimer) {
      window.clearTimeout(feedbackTimer)
    }

    feedbackTimer = window.setTimeout(() => {
      copyFeedback.value = null
      feedbackTimer = null
    }, 1800)
  } catch (error) {
    console.error('Failed to copy connection row', error)
  }
}

onUnmounted(() => {
  if (!feedbackTimer) {
    return
  }

  window.clearTimeout(feedbackTimer)
})
</script>

<template>
  <article class="glass-card panel-block">
    <div class="panel-head panel-head-spread">
      <div class="panel-heading">
        <span class="section-kicker">{{ siteContent.serverPanels.connectionTitle }}</span>
        <h3 class="panel-title">Addresses</h3>
      </div>

      <button v-if="canToggleDetailedView" type="button" class="toggle-view-btn" @click="showAllJava = !showAllJava">
        {{ showAllJava ? siteContent.serverPanels.simpleView : siteContent.serverPanels.fullView }}
      </button>
    </div>

    <p class="connection-hint">{{ siteContent.serverPanels.connectionHint }}</p>
    <transition name="fade">
      <p v-if="copyFeedback" class="copy-toast">{{ copyFeedback }}</p>
    </transition>

    <div v-if="props.connection" class="connection-grid">
      <section v-for="section in connectionSections" :key="section.title" class="connection-block">
        <h4 class="connection-block-title">{{ section.title }}</h4>

        <button
          v-for="row in section.rows"
          :key="row.label"
          type="button"
          class="address-row"
          @click="copyToClipboard(row.copyValue, row.label)"
        >
          <span class="address-badge">{{ row.badge }}</span>
          <span class="address-value">{{ row.value }}</span>
        </button>
      </section>
    </div>

    <p v-else class="loading-copy">{{ siteContent.serverPanels.connectionLoading }}</p>
  </article>
</template>

<style scoped>
.panel-block {
  display: grid;
  gap: 0.95rem;
}

.panel-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 1rem;
}

.panel-head-spread {
  align-items: center;
}

.panel-heading {
  display: grid;
  gap: 0.4rem;
}

.panel-title {
  font-size: 1.6rem;
}

.toggle-view-btn {
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.04);
  color: var(--text-main);
  padding: 0.58rem 0.86rem;
  font-family: var(--sans);
  font-size: 0.84rem;
  letter-spacing: -0.01em;
}

.connection-hint,
.loading-copy {
  color: var(--text-dim);
  font-size: 0.92rem;
}

.connection-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 0.75rem;
}

.connection-block {
  display: grid;
  gap: 0.55rem;
}

.connection-block-title {
  font-size: 1.08rem;
}

.address-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  width: 100%;
  padding: 0.9rem 0.92rem;
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 16px;
  background: rgba(255, 255, 255, 0.03);
  cursor: pointer;
  text-align: left;
  transition: border-color var(--transition-fast), transform var(--transition-fast);
}

.address-row:hover {
  border-color: rgba(76, 147, 251, 0.28);
  transform: translateY(-1px);
}

.address-badge {
  flex-shrink: 0;
  color: var(--primary);
  font-family: var(--mono);
  font-size: 0.68rem;
  text-transform: uppercase;
}

.address-value {
  color: var(--text-muted);
  text-align: right;
}

.copy-toast {
  color: var(--primary);
  font-family: var(--mono);
  font-size: 0.82rem;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

@media (max-width: 980px) {
  .connection-grid {
    grid-template-columns: 1fr;
  }
}
</style>
