<script setup lang="ts">
import { computed } from 'vue'

import { useRevealOnScroll } from '@/composables/useRevealOnScroll'
import { siteContent } from '@/content/siteContent'
import type { LinkedAccount, PlayerInfo } from '@/types/api'

const props = defineProps<{
  info: PlayerInfo
  isOnline: boolean
  linkedAccount: LinkedAccount | null
  completedAdvancements: number
  totalAdvancements: number
  skillTotal: number
  rankHighlights: Array<{ label: string; value: string }>
  formatPlaytime: (value: number) => string
}>()

const { revealed } = useRevealOnScroll<HTMLElement>('heroPanel')

const linkedAccountName = computed(() => {
  if (!props.linkedAccount) {
    return ''
  }

  if (props.info.type === 'Bedrock') {
    return props.linkedAccount.java_username
  }

  return props.linkedAccount.bedrock_username || props.linkedAccount.java_username
})

const summaryMetrics = computed(() => [
  { label: siteContent.playerDetail.summary.playtime, value: props.formatPlaytime(props.info.ticks_lived) },
  {
    label: siteContent.playerDetail.summary.advancements,
    value: `${props.completedAdvancements}/${props.totalAdvancements || 0}`
  },
  {
    label: siteContent.playerDetail.summary.skillTotal,
    value: props.skillTotal > 0 ? props.skillTotal.toLocaleString() : siteContent.playerDetail.summary.noSkillData
  },
  { label: siteContent.playerDetail.summary.xpLevel, value: props.info.xp_level.toLocaleString() }
])
</script>

<template>
  <section
    ref="heroPanel"
    :class="['glass-card', 'player-hero', 'scroll-reveal', 'hover-rise', { 'is-revealed': revealed }]"
  >
    <div class="player-hero-copy">
      <span class="page-kicker">Public player dossier</span>
      <div class="player-title-row">
        <h1 :class="['player-title', 'minecraft-font', { online: isOnline }]">{{ info.last_known_name }}</h1>
      </div>

      <p class="page-lede">
        {{ isOnline ? 'Live presence detected in the server status feed.' : 'Historical record and progression summary.' }}
      </p>

      <div class="player-badge-row">
        <span :class="['badge-pill', 'status-badge', { online: isOnline }]">
          <strong>{{ isOnline ? siteContent.playerDetail.summary.onlineNow : siteContent.playerDetail.summary.archiveRecord }}</strong>
        </span>
        <span class="badge-pill"><strong>{{ info.type }}</strong></span>
        <span v-if="linkedAccountName" class="badge-pill linked-badge">
          {{ siteContent.playerDetail.summary.linkedTo }}
          <strong>{{ linkedAccountName }}</strong>
        </span>
      </div>
    </div>

    <div class="player-hero-summary">
      <div class="hero-metric-grid">
        <article v-for="metric in summaryMetrics" :key="metric.label" class="hud-metric-card hero-metric-card">
          <span class="hud-metric-label">{{ metric.label }}</span>
          <strong class="hud-metric-value">{{ metric.value }}</strong>
        </article>
      </div>

      <div v-if="rankHighlights.length > 0" class="hero-rank-row">
        <span class="hud-caption">{{ siteContent.playerDetail.summary.ranks }}</span>
        <div class="hero-rank-chips">
          <span v-for="highlight in rankHighlights" :key="highlight.label" class="badge-pill hero-rank-chip">
            {{ highlight.label }}
            <strong>{{ highlight.value }}</strong>
          </span>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.player-hero {
  display: grid;
  grid-template-columns: minmax(0, 1.2fr) minmax(320px, 0.9fr);
  gap: 1rem;
  align-items: start;
}

.player-hero-copy,
.player-hero-summary {
  display: grid;
  gap: 0.8rem;
}

.player-title-row {
  display: flex;
  flex-wrap: wrap;
  gap: 0.65rem;
  align-items: flex-end;
}

.player-title {
  font-size: clamp(2.2rem, 5vw, 4rem);
}

.player-title.online {
  color: var(--success);
}

.player-badge-row {
  display: flex;
  flex-wrap: wrap;
  gap: 0.55rem;
}

.status-badge.online strong {
  color: var(--success);
}

.linked-badge {
  max-width: 100%;
}

.hero-metric-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 0.75rem;
}

.hero-metric-card {
  min-height: 100%;
}

.hero-rank-row {
  display: grid;
  gap: 0.55rem;
  padding-top: 0.15rem;
}

.hero-rank-chips {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.hero-rank-chip strong {
  color: var(--accent);
}

@media (max-width: 980px) {
  .player-hero {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 640px) {
  .player-hero {
    gap: 0.75rem;
  }

  .player-hero-copy,
  .player-hero-summary {
    gap: 0.65rem;
  }

  .player-title {
    font-size: clamp(1.85rem, 9vw, 2.8rem);
  }

  .player-badge-row {
    gap: 0.45rem;
  }

  .hero-metric-grid {
    grid-template-columns: 1fr 1fr;
    gap: 0.55rem;
  }

  .hero-metric-card {
    padding: 0.75rem;
  }

  .hero-metric-card :deep(.hud-metric-label),
  .hero-metric-card .hud-metric-label {
    font-size: 0.64rem;
    letter-spacing: 0.12em;
  }

  .hero-metric-card :deep(.hud-metric-value),
  .hero-metric-card .hud-metric-value {
    font-size: 0.98rem;
  }

  .hero-rank-row {
    display: none;
  }
}
</style>
