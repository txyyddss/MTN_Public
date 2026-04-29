<script setup lang="ts">
import { computed } from 'vue'

import { formatPlayerCount, getLocaleValue, useSiteContent } from '@/content/siteContent'
import type { StatusResponse } from '@/types/api'

interface Props {
  status: StatusResponse | null
}

const props = defineProps<Props>()
const siteContent = useSiteContent()

function formatEditionTotal(count: number | undefined, online: boolean | undefined): string {
  if (!online) {
    return siteContent.value.serverPanels.labels.offline
  }

  return formatPlayerCount(count ?? 0)
}

function formatMemory(bytes: number): string {
  return `${(bytes / (1024 * 1024 * 1024)).toFixed(1)} GB`
}

const editionRows = computed(() => [
  {
    label: siteContent.value.serverPanels.javaTitle,
    value: formatEditionTotal(props.status?.java?.players, props.status?.java?.online)
  },
  {
    label: siteContent.value.serverPanels.bedrockTitle,
    value: formatEditionTotal(props.status?.bedrock?.players, props.status?.bedrock?.online)
  }
])

const systemRows = computed(() => {
  const system = props.status?.system
  if (!system) {
    return []
  }

  return [
    { label: siteContent.value.serverPanels.labels.platform, value: system.platform },
    { label: siteContent.value.serverPanels.labels.cpu, value: system.cpu_model },
    {
      label: siteContent.value.serverPanels.labels.load,
      value: `${system.load_1.toFixed(2)} / ${system.load_5.toFixed(2)} / ${system.load_15.toFixed(2)}`
    },
    {
      label: siteContent.value.serverPanels.labels.ram,
      value: `${formatMemory(system.mem_used)} / ${system.mem_percent.toFixed(0)}%`
    }
  ]
})

const updatedLabel = computed(() => {
  if (!props.status?.updated) {
    return null
  }

  return new Intl.DateTimeFormat(getLocaleValue(), {
    year: 'numeric',
    month: 'numeric',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit'
  }).format(new Date(props.status.updated))
})

</script>

<template>
  <article class="glass-card panel-block">
    <span class="panel-watermark" aria-hidden="true">MTN</span>
    <div class="panel-head">
      <div class="panel-heading">
        <span class="section-kicker">{{ siteContent.serverPanels.liveStatusTitle }}</span>
        <h3 class="panel-title">{{ siteContent.serverPanels.liveStatusPanelTitle }}</h3>
      </div>
      <span class="panel-state" :class="{ live: props.status?.java?.online }">
        {{ props.status?.java?.online ? siteContent.serverPanels.operational : siteContent.serverPanels.standby }}
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
        <span v-if="updatedLabel"> / {{ siteContent.serverPanels.lastUpdate }} {{ updatedLabel }}</span>
      </p>
    </div>

    <div v-else class="status-stack" aria-hidden="true">
      <div class="edition-grid">
        <div v-for="index in 2" :key="index" class="status-row">
          <span class="skeleton-line status-skeleton-label"></span>
          <span class="skeleton-line status-skeleton-value"></span>
        </div>
      </div>

      <div class="system-grid">
        <div v-for="index in 4" :key="index" class="system-card">
          <span class="skeleton-line system-skeleton-label"></span>
          <span class="skeleton-line system-skeleton-value"></span>
        </div>
      </div>

      <p class="loading-copy">{{ siteContent.serverPanels.liveStatusFallback }}</p>
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
.status-stack {
  position: relative;
  z-index: 1;
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
  font-size: clamp(1.95rem, 3.5vw, 2.55rem);
}

.panel-state {
  padding: 0.5rem 0.78rem;
  border-radius: 999px;
  border: 1px solid var(--chip-border);
  background: var(--chip-bg);
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.68rem;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.panel-state.live {
  color: var(--success);
  border-color: rgba(var(--secondary-rgb), 0.3);
  background: rgba(var(--secondary-rgb), 0.1);
}

.status-stack {
  display: grid;
  gap: 0.9rem;
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
  padding: 1rem 1.05rem;
  border-radius: 18px;
  background:
    linear-gradient(135deg, rgba(var(--secondary-rgb), 0.12), transparent 42%),
    rgba(255, 255, 255, 0.035);
}

.status-label {
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.72rem;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.status-value {
  color: var(--text-strong);
  font-size: clamp(1.05rem, 2vw, 1.35rem);
  font-weight: 700;
  text-align: right;
}

.system-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 0.65rem;
}

.system-card {
  display: grid;
  gap: 0.3rem;
  padding: 1rem;
  border-radius: 16px;
  background:
    radial-gradient(circle at 100% 0%, rgba(var(--secondary-rgb), 0.1), transparent 44%),
    rgba(255, 255, 255, 0.035);
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
  padding-top: 0.25rem;
  font-family: var(--mono);
}

.status-skeleton-label {
  width: 7rem;
}

.status-skeleton-value {
  width: 5.5rem;
}

.system-skeleton-label {
  width: 4rem;
}

.system-skeleton-value {
  width: 100%;
}

@media (max-width: 980px) {
  .system-grid {
    grid-template-columns: 1fr;
  }
}
</style>
