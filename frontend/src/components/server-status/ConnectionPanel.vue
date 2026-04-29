<script setup lang="ts">
import { computed, ref } from 'vue'

import { useSiteContent } from '@/content/siteContent'
import type { ConnectionResponse } from '@/types/api'

interface Props {
  connection: ConnectionResponse | null
}

const props = defineProps<Props>()
const copiedId = ref<string | null>(null)
const siteContent = useSiteContent()

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

  const formatAddress = (address?: { domain?: string; ip: string; port?: string } | string) => {
    if (!address) return ''
    if (typeof address === 'string') return address

    const host = address.domain || address.ip
    const port = address.port

    if (!port) return host

    if (host.includes(`:${port}`)) return host

    const isIPv6Literal = host.includes(':') && !host.includes('.')
    
    if (isIPv6Literal) {
      return `[${host}]:${port}`
    }

    return `${host}:${port}`
  }

  if (conn?.java_srv) {
    list.push({ label: 'SRV', value: formatAddress(conn.java_srv), id: 'java-srv' })
  } else if (addrs?.java_ipv4_srv) {
    list.push({ label: 'SRV v4', value: addrs.java_ipv4_srv, id: 'java-srv-static' })
  }

  if (addrs?.java_ipv6_srv) {
    list.push({ label: 'SRV v6', value: addrs.java_ipv6_srv, id: 'java-srv-v6' })
  }

  if (conn?.java_ipv4) {
    list.push({ label: 'IPv4', value: formatAddress(conn.java_ipv4), id: 'java-ipv4' })
  }

  if (conn?.java_ipv6) {
    list.push({ label: 'IPv6', value: formatAddress(conn.java_ipv6), id: 'java-ipv6' })
  } else if (addrs?.java_ipv6) {
    list.push({ label: 'IPv6', value: addrs.java_ipv6, id: 'java-ipv6-static' })
  }

  return list
})

const bedrockAddresses = computed(() => {
  const conn = props.connection?.connection
  const addrs = props.connection?.addresses
  const list: AddressDisplay[] = []

  const formatAddress = (address?: { domain?: string; ip: string; port?: string } | string) => {
    if (!address) return ''
    if (typeof address === 'string') return address
    const host = address.domain || address.ip
    const port = address.port
    if (!port) return host
    if (host.includes(`:${port}`)) return host
    const isIPv6Literal = host.includes(':') && !host.includes('.')
    return isIPv6Literal ? `[${host}]:${port}` : `${host}:${port}`
  }

  if (conn?.bedrock_ipv4) {
    list.push({ label: 'IPv4', value: formatAddress(conn.bedrock_ipv4), id: 'bedrock-ipv4' })
  }

  if (conn?.bedrock_ipv6) {
    list.push({ label: 'IPv6', value: formatAddress(conn.bedrock_ipv6), id: 'bedrock-ipv6' })
  } else if (addrs?.bedrock_ipv6) {
    list.push({ label: 'IPv6', value: addrs.bedrock_ipv6, id: 'bedrock-ipv6-static' })
  }

  return list
})

</script>

<template>
  <article class="glass-card panel-block">
    <span class="panel-watermark" aria-hidden="true">MTN</span>
    <div class="panel-head">
      <div class="panel-heading">
        <span class="section-kicker">{{ siteContent.serverPanels.connectionTitle }}</span>
        <h3 class="panel-title">{{ siteContent.serverPanels.connectionPanelTitle }}</h3>
      </div>
    </div>

    <div v-if="props.connection" class="connection-stack">
      <div class="edition-section">
        <h4 class="edition-title">{{ siteContent.serverPanels.javaTitle }}</h4>
        <div class="address-grid">
          <div v-for="addr in javaAddresses" :key="addr.id" class="address-row">
            <div class="address-info">
              <span class="address-label">{{ addr.label }}</span>
              <code class="address-value" :title="addr.value">{{ addr.value }}</code>
            </div>
            <button 
              class="copy-btn" 
              type="button" 
              @click="copyToClipboard(addr.value, addr.id)"
              :class="{ copied: copiedId === addr.id }"
            >
              {{ copiedId === addr.id ? siteContent.serverPanels.copied : siteContent.serverPanels.copy }}
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
              <code class="address-value" :title="addr.value">{{ addr.value }}</code>
            </div>
            <button 
              class="copy-btn" 
              type="button" 
              @click="copyToClipboard(addr.value, addr.id)"
              :class="{ copied: copiedId === addr.id }"
            >
              {{ copiedId === addr.id ? siteContent.serverPanels.copied : siteContent.serverPanels.copy }}
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
  isolation: isolate;
  display: grid;
  gap: 1.2rem;
  min-height: 100%;
  padding: 1.5rem;
  border-color: var(--glass-border);
  background:
    radial-gradient(circle at 88% 0%, rgba(var(--secondary-rgb), 0.16), transparent 34%),
    linear-gradient(180deg, rgba(255, 255, 255, 0.06), rgba(255, 255, 255, 0.018)),
    rgba(3, 8, 24, 0.78);
}

.panel-watermark {
  position: absolute;
  right: 1rem;
  bottom: 0.7rem;
  z-index: 0;
  color: rgba(141, 184, 255, 0.07);
  font-family: var(--mono);
  font-size: clamp(3rem, 7vw, 5rem);
  font-weight: 900;
  letter-spacing: 0.26em;
  pointer-events: none;
}

.panel-head,
.connection-stack {
  position: relative;
  z-index: 1;
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
  font-size: clamp(1.95rem, 3.5vw, 2.55rem);
}

.connection-stack {
  display: grid;
  gap: 1.3rem;
}

.edition-section {
  display: grid;
  gap: 0.85rem;
}

.edition-title {
  font-size: 0.9rem;
  color: var(--accent);
  font-family: var(--mono);
  font-weight: 700;
  letter-spacing: 0.12em;
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
  padding: 0.95rem 1rem;
  border-radius: 18px;
  background:
    linear-gradient(135deg, rgba(var(--secondary-rgb), 0.1), transparent 38%),
    rgba(255, 255, 255, 0.035);
  transition: border-color var(--transition-fast), background var(--transition-fast);
}

.address-row:hover {
  border-color: rgba(var(--secondary-rgb), 0.22);
  background: rgba(var(--secondary-rgb), 0.08);
}

.address-info {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  min-width: 0;
  flex: 1;
}

.address-label {
  font-family: var(--mono);
  font-size: 0.65rem;
  color: var(--primary);
  background: rgba(var(--secondary-rgb), 0.12);
  padding: 0.24rem 0.46rem;
  border: 1px solid rgba(var(--secondary-rgb), 0.18);
  border-radius: 999px;
  flex-shrink: 0;
}

.address-value {
  font-family: var(--mono);
  font-size: 0.9rem;
  color: var(--text-strong);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  flex: 1;
}

.copy-btn {
  padding: 0.48rem 0.88rem;
  font-size: 0.75rem;
  font-weight: 600;
  border-radius: 999px;
  border: 1px solid rgba(var(--secondary-rgb), 0.2);
  background: rgba(var(--secondary-rgb), 0.08);
  color: var(--text-main);
  transition: all var(--transition-fast);
  flex-shrink: 0;
}

.copy-btn:hover {
  background: rgba(var(--secondary-rgb), 0.16);
  color: var(--text-strong);
  border-color: rgba(var(--secondary-rgb), 0.34);
  transform: translateY(-1px);
}

.copy-btn.copied {
  background: rgba(var(--secondary-rgb), 0.18);
  color: var(--success);
  border-color: var(--success);
}

.panel-foot {
  color: var(--text-dim);
  font-size: 0.85rem;
  border-top: 1px solid var(--glass-border-soft);
  padding-top: 0.75rem;
}

.title-skeleton { width: 100px; height: 1.2rem; }
.skeleton-row { border-style: dashed; background: transparent; }
.label-skeleton { width: 40px; height: 1rem; }
.value-skeleton { width: 150px; height: 1rem; }

@media (max-width: 480px) {
  .address-row {
    flex-direction: column;
    align-items: stretch;
    padding: 1rem;
    gap: 0.85rem;
  }
  
  .address-info {
    width: 100%;
  }

  .address-value {
    font-size: 0.85rem;
  }

  .copy-btn {
    width: 100%;
    padding: 0.6rem;
  }
}

@keyframes animate-pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.5; }
}
.animate-pulse {
  animation: animate-pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}
</style>
