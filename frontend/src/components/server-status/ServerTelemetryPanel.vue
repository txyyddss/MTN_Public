<script setup lang="ts">
import { computed } from 'vue'

import { siteContent } from '@/content/siteContent'
import type { StatusResponse } from '@/types/api'

interface Props {
  status: StatusResponse | null
}

const props = defineProps<Props>()

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

const editionRows = computed(() => [
  {
    label: 'Java Edition',
    value: formatEditionTotal(props.status?.java?.players, props.status?.java?.online)
  },
  {
    label: 'Bedrock Edition',
    value: formatEditionTotal(props.status?.bedrock?.players, props.status?.bedrock?.online)
  }
])

const systemRows = computed(() => {
  const system = props.status?.system
  if (!system) {
    return []
  }

  return [
    { label: siteContent.serverPanels.labels.platform, value: system.platform },
    { label: siteContent.serverPanels.labels.cpu, value: system.cpu_model },
    {
      label: siteContent.serverPanels.labels.load,
      value: `${system.load_1.toFixed(2)} / ${system.load_5.toFixed(2)} / ${system.load_15.toFixed(2)}`
    },
    {
      label: siteContent.serverPanels.labels.ram,
      value: `${formatMemory(system.mem_used)} / ${system.mem_percent.toFixed(0)}%`
    },
    {
      label: siteContent.serverPanels.labels.network,
      value: formatBytes(system.net_sent_rate + system.net_recv_rate)
    }
  ]
})

const updatedLabel = computed(() => {
  if (!props.status?.updated) {
    return null
  }

  return new Date(props.status.updated).toLocaleString()
})
</script>

<template>
  <article class="glass-card panel-block">
    <div class="panel-head">
      <div class="panel-heading">
        <span class="section-kicker">{{ siteContent.serverPanels.liveStatusTitle }}</span>
        <h3 class="panel-title">Telemetry and system load</h3>
      </div>
    </div>

    <div v-if="props.status" class="status-stack">
      <div v-for="row in editionRows" :key="row.label" class="status-row">
        <span class="status-label">{{ row.label }}</span>
        <strong class="status-value">{{ row.value }}</strong>
      </div>

      <div v-if="systemRows.length" class="system-grid">
        <div v-for="row in systemRows" :key="row.label" class="system-card">
          <small class="system-label">{{ row.label }}</small>
          <strong class="system-value">{{ row.value }}</strong>
        </div>
      </div>

      <p class="panel-foot">
        {{ siteContent.serverPanels.liveStatusRefresh }}
        <span v-if="updatedLabel"> / Last update {{ updatedLabel }}</span>
      </p>
    </div>

    <p v-else class="loading-copy">{{ siteContent.serverPanels.liveStatusFallback }}</p>
  </article>
</template>

<style scoped>
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

.panel-heading {
  display: grid;
  gap: 0.4rem;
}

.panel-title {
  font-size: 2rem;
}

.status-stack {
  display: grid;
  gap: 0.9rem;
}

.status-row {
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

.status-label {
  color: var(--text-muted);
}

.status-value {
  color: var(--text-strong);
  font-weight: 600;
  text-align: right;
}

.system-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 0.85rem;
}

.system-card {
  display: grid;
  gap: 0.25rem;
  padding: 0.9rem 1rem;
  border-radius: 18px;
  background: rgba(255, 248, 234, 0.03);
  border: 1px solid rgba(255, 248, 234, 0.05);
}

.system-label {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.72rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.system-value {
  color: var(--text-strong);
  font-size: 0.96rem;
  font-weight: 600;
}

.panel-foot,
.loading-copy {
  color: var(--text-dim);
  font-size: 0.92rem;
}

@media (max-width: 980px) {
  .system-grid {
    grid-template-columns: 1fr;
  }
}
</style>
