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
        <h3 class="panel-title">Node load</h3>
      </div>
      <span class="panel-state" :class="{ live: props.status?.java?.online }">
        {{ props.status?.java?.online ? 'Operational' : 'Standby' }}
      </span>
    </div>

    <div v-if="props.status" class="status-stack">
      <div class="edition-grid">
        <div v-for="row in editionRows" :key="row.label" class="status-row">
          <span class="status-label">{{ row.label }}</span>
          <strong class="status-value">{{ row.value }}</strong>
        </div>
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
  gap: 0.95rem;
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
  font-size: 1.6rem;
}

.panel-state {
  padding: 0.45rem 0.72rem;
  border-radius: 999px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.68rem;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.panel-state.live {
  color: var(--success);
}

.status-stack {
  display: grid;
  gap: 0.75rem;
}

.edition-grid {
  display: grid;
  gap: 0.75rem;
}

.status-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  padding: 0.9rem 0.95rem;
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 16px;
  background: rgba(255, 255, 255, 0.03);
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
  gap: 0.65rem;
}

.system-card {
  display: grid;
  gap: 0.25rem;
  padding: 0.9rem;
  border-radius: 16px;
  background: rgba(255, 255, 255, 0.028);
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.system-label {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.68rem;
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
  font-size: 0.9rem;
}

@media (max-width: 980px) {
  .system-grid {
    grid-template-columns: 1fr;
  }
}
</style>
