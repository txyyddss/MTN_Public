<script setup lang="ts">
import { ref } from 'vue'

import { siteContent } from '@/content/siteContent'
import type { ConnectionResponse } from '@/types/api'

interface Props {
  connection: ConnectionResponse | null
}

const props = defineProps<Props>()
const copiedId = ref<string | null>(null)

async function copyToClipboard(text: string, id: string) {
  try {
    await navigator.clipboard.writeText(text)
    copiedId.value = id
    setTimeout(() => {
      if (copiedId.value === id) {
        copiedId.value = null
      }
    }, 2000)
  } catch (err) {
    console.error('Failed to copy:', err)
  }
}

interface AddressDisplay {
  label: string
  value: string
  id: string
}

const javaAddresses = computed(() => {
  const conn = props.connection?.connection
  const addrs = props.connection?.addresses
  const list: AddressDisplay[] = []

  if (conn?.java_srv) {
    list.push({ label: 'SRV', value: conn.java_srv.domain || conn.java_srv.ip, id: 'java-srv' })
  } else if (addrs?.java_ipv4_srv) {
    list.push({ label: 'SRV', value: addrs.java_ipv4_srv, id: 'java-srv-static' })
  }

  if (conn?.java_ipv4) {
    const val = conn.java_ipv4.port ? `${conn.java_ipv4.domain || conn.java_ipv4.ip}:${conn.java_ipv4.port}` : (conn.java_ipv4.domain || conn.java_ipv4.ip)
    list.push({ label: 'IPv4', value: val, id: 'java-ipv4' })
  }

  if (conn?.java_ipv6) {
    const val = conn.java_ipv6.port ? `[${conn.java_ipv6.domain || conn.java_ipv6.ip}]:${conn.java_ipv6.port}` : (conn.java_ipv6.domain || conn.java_ipv6.ip)
    list.push({ label: 'IPv6', value: val, id: 'java-ipv6' })
  } else if (addrs?.java_ipv6) {
    list.push({ label: 'IPv6', value: addrs.java_ipv6, id: 'java-ipv6-static' })
  }

  return list
})

const bedrockAddresses = computed(() => {
  const conn = props.connection?.connection
  const addrs = props.connection?.addresses
  const list: AddressDisplay[] = []

  if (conn?.bedrock_ipv4) {
    const val = conn.bedrock_ipv4.port ? `${conn.bedrock_ipv4.domain || conn.bedrock_ipv4.ip}:${conn.bedrock_ipv4.port}` : (conn.bedrock_ipv4.domain || conn.bedrock_ipv4.ip)
    list.push({ label: 'IPv4', value: val, id: 'bedrock-ipv4' })
  }

  if (conn?.bedrock_ipv6) {
    const val = conn.bedrock_ipv6.port ? `[${conn.bedrock_ipv6.domain || conn.bedrock_ipv6.ip}]:${conn.bedrock_ipv6.port}` : (conn.bedrock_ipv6.domain || conn.bedrock_ipv6.ip)
    list.push({ label: 'IPv6', value: val, id: 'bedrock-ipv6' })
  } else if (addrs?.bedrock_ipv6) {
    list.push({ label: 'IPv6', value: addrs.bedrock_ipv6, id: 'bedrock-ipv6-static' })
  }

  return list
})

import { computed } from 'vue'
</script>

<template>
  <article class="glass-card panel-block">
    <div class="panel-head">
      <div class="panel-heading">
        <span class="section-kicker">{{ siteContent.serverPanels.connectionTitle }}</span>
        <h3 class="panel-title">Connection routing</h3>
      </div>
    </div>

    <div v-if="props.connection" class="connection-stack">
      <div class="edition-section">
        <h4 class="edition-title">{{ siteContent.serverPanels.javaTitle }}</h4>
        <div class="address-grid">
          <div v-for="addr in javaAddresses" :key="addr.id" class="address-row">
            <div class="address-info">
              <span class="address-label">{{ addr.label }}</span>
              <code class="address-value">{{ addr.value }}</code>
            </div>
            <button 
              class="copy-btn" 
              type="button" 
              @click="copyToClipboard(addr.value, addr.id)"
              :class="{ copied: copiedId === addr.id }"
            >
              {{ copiedId === addr.id ? 'Copied' : 'Copy' }}
            </button>
          </div>
        </div>
      </div>

      <div class="edition-section">
        <h4 class="edition-title">{{ siteContent.serverPanels.bedrockTitle }}</h4>
        <div class="address-grid">
          <div v-for="addr in bedrockAddresses" :key="addr.id" class="address-row">
            <div class="address-info">
              <span class="address-label">{{ addr.label }}</span>
              <code class="address-value">{{ addr.value }}</code>
            </div>
            <button 
              class="copy-btn" 
              type="button" 
              @click="copyToClipboard(addr.value, addr.id)"
              :class="{ copied: copiedId === addr.id }"
            >
              {{ copiedId === addr.id ? 'Copied' : 'Copy' }}
            </button>
          </div>
        </div>
      </div>

      <p class="panel-foot">
        {{ siteContent.serverPanels.connectionHint }}
      </p>
    </div>

    <div v-else class="connection-stack animate-pulse">
      <div v-for="i in 2" :key="i" class="edition-section">
        <div class="skeleton-line title-skeleton"></div>
        <div class="address-grid">
          <div v-for="j in 2" :key="j" class="address-row skeleton-row">
            <div class="skeleton-line label-skeleton"></div>
            <div class="skeleton-line value-skeleton"></div>
          </div>
        </div>
      </div>
    </div>
  </article>
</template>

<style scoped>
.panel-block {
  display: grid;
  gap: 1.25rem;
}

.panel-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
}

.panel-heading {
  display: grid;
  gap: 0.4rem;
}

.panel-title {
  font-size: 1.6rem;
}

.connection-stack {
  display: grid;
  gap: 1.5rem;
}

.edition-section {
  display: grid;
  gap: 0.75rem;
}

.edition-title {
  font-size: 0.9rem;
  color: var(--text-muted);
  font-family: var(--mono);
  letter-spacing: 0.05em;
  text-transform: uppercase;
}

.address-grid {
  display: grid;
  gap: 0.6rem;
}

.address-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  padding: 0.75rem 1rem;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 14px;
  transition: border-color var(--transition-fast), background var(--transition-fast);
}

.address-row:hover {
  background: rgba(255, 255, 255, 0.05);
  border-color: rgba(255, 255, 255, 0.12);
}

.address-info {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  min-width: 0;
}

.address-label {
  font-family: var(--mono);
  font-size: 0.65rem;
  color: var(--primary);
  background: rgba(76, 147, 251, 0.1);
  padding: 0.2rem 0.4rem;
  border-radius: 4px;
  flex-shrink: 0;
}

.address-value {
  font-family: var(--mono);
  font-size: 0.9rem;
  color: var(--text-strong);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.copy-btn {
  padding: 0.4rem 0.8rem;
  font-size: 0.75rem;
  font-weight: 600;
  border-radius: 8px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  background: rgba(255, 255, 255, 0.05);
  color: var(--text-muted);
  transition: all var(--transition-fast);
  flex-shrink: 0;
}

.copy-btn:hover {
  background: rgba(255, 255, 255, 0.1);
  color: var(--text-strong);
  border-color: rgba(255, 255, 255, 0.2);
}

.copy-btn.copied {
  background: var(--success);
  color: #000;
  border-color: var(--success);
}

.panel-foot {
  color: var(--text-dim);
  font-size: 0.85rem;
  border-top: 1px solid rgba(255, 255, 255, 0.08);
  padding-top: 0.75rem;
}

.title-skeleton { width: 100px; height: 1.2rem; }
.skeleton-row { border-style: dashed; background: transparent; }
.label-skeleton { width: 40px; height: 1rem; }
.value-skeleton { width: 150px; height: 1rem; }

@keyframes animate-pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.5; }
}
.animate-pulse {
  animation: animate-pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}
</style>
