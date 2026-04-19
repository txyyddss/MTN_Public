<script setup lang="ts">
import { computed, ref } from 'vue'
import { storeToRefs } from 'pinia'

import { siteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'
import type { ConnectionAddress } from '@/types/api'

interface AddressRow {
  label: string
  badge: string
  value: string
  copyValue: string
}

const showAllJava = ref(false)
const copyFeedback = ref<string | null>(null)

const statusStore = useServerStatusStore()
const { status, connection } = storeToRefs(statusStore)

function copyToClipboard(text: string, label: string): void {
  void navigator.clipboard.writeText(text).then(() => {
    copyFeedback.value = `${siteContent.serverPanels.copyPrefix} ${label}`
    window.setTimeout(() => {
      copyFeedback.value = null
    }, 1800)
  })
}

function formatEditionTotal(count: number | undefined, online: boolean | undefined): string {
  if (!online) {
    return siteContent.serverPanels.labels.offline
  }

  return `${count ?? 0} ${count === 1 ? 'player' : 'players'}`
}

function formatBytes(bytes: number): string {
  if (bytes <= 0) {
    return '0 B/s'
  }

  const units = ['B/s', 'KB/s', 'MB/s', 'GB/s']
  const index = Math.min(Math.floor(Math.log(bytes) / Math.log(1024)), units.length - 1)
  return `${(bytes / Math.pow(1024, index)).toFixed(1)} ${units[index]}`
}

function formatMemory(bytes: number): string {
  return `${(bytes / (1024 * 1024 * 1024)).toFixed(1)} GB`
}

function buildDirectRow(label: string, badge: string, entry?: ConnectionAddress): AddressRow | null {
  if (!entry) {
    return null
  }

  const host = entry.domain || entry.ip
  const value = entry.port ? `${host}:${entry.port}` : host

  return {
    label,
    badge,
    value,
    copyValue: value
  }
}

const javaRows = computed<AddressRow[]>(() => {
  if (!connection.value) {
    return []
  }

  if (showAllJava.value) {
    return [
      buildDirectRow('Java IPv4', 'IPv4', connection.value.connection?.java_ipv4),
      buildDirectRow('Java IPv6', 'IPv6', connection.value.connection?.java_ipv6)
    ].filter((row): row is AddressRow => Boolean(row))
  }

  return [
    connection.value.addresses.java_ipv4_srv
      ? {
          label: 'Java SRV IPv4',
          badge: 'SRV v4',
          value: connection.value.addresses.java_ipv4_srv,
          copyValue: connection.value.addresses.java_ipv4_srv
        }
      : null,
    connection.value.addresses.java_ipv6_srv
      ? {
          label: 'Java SRV IPv6',
          badge: 'SRV v6',
          value: connection.value.addresses.java_ipv6_srv,
          copyValue: connection.value.addresses.java_ipv6_srv
        }
      : null
  ].filter((row): row is AddressRow => Boolean(row))
})

const bedrockRows = computed<AddressRow[]>(() =>
  [
    buildDirectRow('Bedrock IPv6', 'IPv6', connection.value?.connection?.bedrock_ipv6),
    buildDirectRow('Bedrock IPv4', 'IPv4', connection.value?.connection?.bedrock_ipv4)
  ].filter((row): row is AddressRow => Boolean(row))
)

const updatedLabel = computed(() => {
  if (!status.value?.updated) {
    return null
  }

  return new Date(status.value.updated).toLocaleString()
})
</script>

<template>
  <section class="server-status-section">
    <div class="container layout-grid animate-entry delay-300">
      <article class="glass-card panel-block">
        <div class="panel-head">
          <div>
            <span class="section-kicker">{{ siteContent.serverPanels.liveStatusTitle }}</span>
            <h3>Telemetry and system load</h3>
          </div>
        </div>

        <div v-if="status" class="status-stack">
          <div class="status-row">
            <span>Java Edition</span>
            <strong>{{ formatEditionTotal(status.java?.players, status.java?.online) }}</strong>
          </div>
          <div class="status-row">
            <span>Bedrock Edition</span>
            <strong>{{ formatEditionTotal(status.bedrock?.players, status.bedrock?.online) }}</strong>
          </div>

          <div v-if="status.system" class="system-grid">
            <div>
              <small>{{ siteContent.serverPanels.labels.platform }}</small>
              <strong>{{ status.system.platform }}</strong>
            </div>
            <div>
              <small>{{ siteContent.serverPanels.labels.cpu }}</small>
              <strong>{{ status.system.cpu_model }}</strong>
            </div>
            <div>
              <small>{{ siteContent.serverPanels.labels.load }}</small>
              <strong>
                {{ status.system.load_1.toFixed(2) }} / {{ status.system.load_5.toFixed(2) }} /
                {{ status.system.load_15.toFixed(2) }}
              </strong>
            </div>
            <div>
              <small>{{ siteContent.serverPanels.labels.ram }}</small>
              <strong>{{ formatMemory(status.system.mem_used) }} · {{ status.system.mem_percent.toFixed(0) }}%</strong>
            </div>
            <div>
              <small>{{ siteContent.serverPanels.labels.network }}</small>
              <strong>{{ formatBytes(status.system.net_sent_rate + status.system.net_recv_rate) }}</strong>
            </div>
          </div>

          <p class="panel-foot">
            {{ siteContent.serverPanels.liveStatusRefresh }}
            <span v-if="updatedLabel"> · Last update {{ updatedLabel }}</span>
          </p>
        </div>
        <p v-else class="loading-copy">{{ siteContent.serverPanels.liveStatusFallback }}</p>
      </article>

      <article class="glass-card panel-block">
        <div class="panel-head panel-head-spread">
          <div>
            <span class="section-kicker">{{ siteContent.serverPanels.connectionTitle }}</span>
            <h3>Verified addresses</h3>
          </div>
          <button v-if="connection?.connection" class="toggle-view-btn" @click="showAllJava = !showAllJava">
            {{ showAllJava ? siteContent.serverPanels.simpleView : siteContent.serverPanels.fullView }}
          </button>
        </div>

        <p class="connection-hint">{{ siteContent.serverPanels.connectionHint }}</p>
        <transition name="fade">
          <p v-if="copyFeedback" class="copy-toast">{{ copyFeedback }}</p>
        </transition>

        <div v-if="connection" class="connection-grid">
          <section class="connection-block">
            <h4>{{ siteContent.serverPanels.javaTitle }}</h4>
            <button
              v-for="row in javaRows"
              :key="row.label"
              type="button"
              class="address-row"
              @click="copyToClipboard(row.copyValue, row.label)"
            >
              <span class="address-badge">{{ row.badge }}</span>
              <span class="address-value">{{ row.value }}</span>
            </button>
          </section>

          <section class="connection-block">
            <h4>{{ siteContent.serverPanels.bedrockTitle }}</h4>
            <button
              v-for="row in bedrockRows"
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
    </div>
  </section>
</template>

<style scoped>
.server-status-section {
  padding-bottom: 5rem;
}

.layout-grid {
  display: grid;
  grid-template-columns: 0.95fr 1.15fr;
  gap: 1.25rem;
}

.panel-block {
  display: grid;
  gap: 1.35rem;
}

.panel-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 1rem;
}

.panel-head h3 {
  font-size: 2rem;
}

.panel-head-spread {
  align-items: center;
}

.status-stack {
  display: grid;
  gap: 0.9rem;
}

.status-row,
.address-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  width: 100%;
  padding: 0.95rem 1rem;
  border: 1px solid rgba(255, 248, 234, 0.08);
  border-radius: 18px;
  background: rgba(255, 248, 234, 0.035);
}

.status-row span,
.address-value {
  color: var(--text-muted);
}

.status-row strong {
  color: var(--text-strong);
  font-weight: 600;
  text-align: right;
}

.system-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 0.85rem;
}

.system-grid > div {
  display: grid;
  gap: 0.25rem;
  padding: 0.9rem 1rem;
  border-radius: 18px;
  background: rgba(255, 248, 234, 0.03);
  border: 1px solid rgba(255, 248, 234, 0.05);
}

.system-grid small {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.72rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.system-grid strong {
  color: var(--text-strong);
  font-size: 0.96rem;
  font-weight: 600;
}

.panel-foot,
.connection-hint,
.loading-copy {
  color: var(--text-dim);
  font-size: 0.92rem;
}

.toggle-view-btn {
  border: 1px solid var(--glass-border-bright);
  border-radius: 999px;
  background: rgba(255, 248, 234, 0.03);
  color: var(--text-main);
  padding: 0.65rem 0.95rem;
}

.connection-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 1rem;
}

.connection-block {
  display: grid;
  gap: 0.75rem;
}

.connection-block h4 {
  font-size: 1.45rem;
}

.address-row {
  cursor: pointer;
  text-align: left;
}

.address-row:hover {
  border-color: var(--glass-border-bright);
  transform: translateY(-1px);
}

.address-badge {
  flex-shrink: 0;
  color: var(--primary);
  font-family: var(--mono);
  font-size: 0.74rem;
  text-transform: uppercase;
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
  .layout-grid,
  .connection-grid,
  .system-grid {
    grid-template-columns: 1fr;
  }
}
</style>
