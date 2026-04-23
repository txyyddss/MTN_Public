<script setup lang="ts">
import { computed } from 'vue'

import HourlyPresenceHeatmap from '@/components/heatmap/HourlyPresenceHeatmap.vue'
import { formatPlayerCount, getLocaleValue, useSiteContent } from '@/content/siteContent'
import type { ServerHistoryResponse, StatusResponse } from '@/types/api'

interface Props {
  status: StatusResponse | null
  history: ServerHistoryResponse | null
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

const historySummary = computed(() => {
  if (!props.history) {
    return siteContent.value.serverPanels.historyLoading
  }

  const peak = props.history.weekly_max_players
  return peak > 0
    ? siteContent.value.serverPanels.historyPeak.replace('{peak}', String(peak))
    : siteContent.value.serverPanels.historyEmpty
})
</script>

<template>
  <article class="glass-card panel-block">
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

      <div class="history-block">
        <div class="history-head">
          <div class="panel-heading">
            <span class="section-kicker">{{ siteContent.serverPanels.historyTitle }}</span>
            <p class="panel-foot">{{ siteContent.serverPanels.historyHint }}</p>
          </div>
          <span class="panel-state" :class="{ live: (props.history?.weekly_max_players ?? 0) > 0 }">
            {{ historySummary }}
          </span>
        </div>

        <HourlyPresenceHeatmap
          v-if="props.history"
          :days="props.history.days"
          :hours="props.history.hours"
          :cells="props.history.cells"
          :timezone="props.history.timezone"
          :weekly-max-players="props.history.weekly_max_players"
          variant="server"
        />
        <p v-else class="loading-copy">{{ siteContent.serverPanels.historyLoading }}</p>
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

.history-block {
  display: grid;
  gap: 0.75rem;
  padding-top: 0.25rem;
  border-top: 1px solid rgba(255, 255, 255, 0.08);
}

.history-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 1rem;
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
