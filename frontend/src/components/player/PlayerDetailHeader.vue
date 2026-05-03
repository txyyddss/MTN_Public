<script setup lang="ts">
import { computed } from 'vue'

import SkinViewer from '@/components/SkinViewer.vue'
import { getLocaleValue, getPlatformLabel, useSiteContent } from '@/content/siteContent'
import type { LeaderboardTarget, LinkedAccount, PlayerInfo, WhitelistAccount } from '@/types/api'
import type { PlayerRankHighlight } from '@/types/playerDetail'
import { getSkinUrl } from '@/utils/minecraft'

const props = defineProps<{
  info: PlayerInfo
  isOnline: boolean
  linkedAccount: LinkedAccount | null
  whitelistAccount: WhitelistAccount | null
  completedAdvancements: number
  totalAdvancements: number
  rankHighlights: PlayerRankHighlight[]
  topRankHighlight: PlayerRankHighlight | null
  formatDate: (value: number) => string
  formatDateTime: (value: number) => string
  formatPlaytime: (value: number) => string
}>()

const emit = defineEmits<{
  (event: 'selectLeaderboard', value: LeaderboardTarget): void
}>()

interface HeroMetric {
  label: string
  value: string
}

interface IdentityRow {
  label: string
  value: string
  compact?: boolean
}

const siteContent = useSiteContent()

const platformLabel = computed(() => getPlatformLabel(props.info.type))

const statusLabel = computed(() =>
  props.isOnline ? siteContent.value.playerDetail.summary.onlineNow : siteContent.value.playerDetail.summary.archiveRecord
)

const linkedAccountName = computed(() => {
  if (!props.linkedAccount) {
    return ''
  }

  if (props.info.type === 'Bedrock') {
    return props.linkedAccount.java_username || props.linkedAccount.bedrock_username || props.linkedAccount.java_uuid
  }

  return props.linkedAccount.bedrock_username || props.linkedAccount.java_username || props.linkedAccount.bedrock_uuid
})

const heroMetrics = computed<HeroMetric[]>(() => [
  {
    label: siteContent.value.playerDetail.profile.playtime,
    value: props.formatPlaytime(props.info.ticks_lived)
  },
  {
    label: siteContent.value.playerDetail.profile.xpLevel,
    value: props.info.xp_level.toLocaleString(getLocaleValue())
  },
  {
    label: siteContent.value.playerDetail.profile.firstJoin,
    value: props.formatDate(props.info.first_played)
  },
  {
    label: siteContent.value.playerDetail.profile.lastSeen,
    value: props.formatDateTime(props.info.last_seen)
  },
  {
    label: siteContent.value.playerDetail.summary.advancements,
    value: `${props.completedAdvancements}/${props.totalAdvancements || 0}`
  }
])

const identityRows = computed<IdentityRow[]>(() => {
  const rows: IdentityRow[] = [
    {
      label: siteContent.value.playerDetail.header.uuid,
      value: props.info.uuid,
      compact: true
    },
    {
      label: siteContent.value.playerDetail.header.edition,
      value: platformLabel.value
    }
  ]

  if (linkedAccountName.value) {
    rows.push({
      label: siteContent.value.playerDetail.profile.linkedTo,
      value: linkedAccountName.value
    })
  }

  if (props.whitelistAccount) {
    rows.push({
      label: siteContent.value.playerDetail.profile.qqOwner,
      value: props.whitelistAccount.qq_id_masked
    })
  }

  return rows
})

const secondaryRankHighlights = computed(() => props.rankHighlights.slice(1))

function selectLeaderboard(target: LeaderboardTarget): void {
  emit('selectLeaderboard', target)
}
</script>

<template>
  <section class="glass-card detail-hero player-glass-card player-glass-reveal">
    <span class="detail-header-mark" aria-hidden="true">MTN</span>

    <div class="detail-hero-layout">
      <div class="skin-frame" aria-hidden="true">
        <SkinViewer :skin-url="getSkinUrl(info.last_known_name, info.type)" :width="230" :height="292" />
      </div>

      <div class="detail-hero-main">
        <div class="detail-kicker-row">
          <span class="page-kicker">{{ siteContent.playerDetail.header.kicker }}</span>
          <span :class="['hero-chip', 'status-chip', { online: isOnline }]">{{ statusLabel }}</span>
          <span class="hero-chip">{{ platformLabel }}</span>
        </div>

        <div class="detail-title-block">
          <h1 :class="['detail-player-name', 'minecraft-font', { online: isOnline }]">{{ info.last_known_name }}</h1>
          <p class="detail-player-copy">
            {{ isOnline ? siteContent.playerDetail.header.liveBody : siteContent.playerDetail.header.archiveBody }}
          </p>
        </div>

        <div class="hero-metric-strip" :aria-label="siteContent.playerDetail.header.metricsAria">
          <div v-for="metric in heroMetrics" :key="metric.label" class="hero-metric">
            <span class="hero-metric-label">{{ metric.label }}</span>
            <strong class="hero-metric-value">{{ metric.value }}</strong>
          </div>

          <button
            v-if="topRankHighlight"
            type="button"
            class="hero-metric hero-metric-action"
            @click="selectLeaderboard(topRankHighlight.target)"
          >
            <span class="hero-metric-label">{{ siteContent.playerDetail.summary.bestRank }}</span>
            <strong class="hero-metric-value">{{ topRankHighlight.value }}</strong>
            <small class="hero-metric-note">{{ topRankHighlight.label }}</small>
          </button>
          <div v-else class="hero-metric">
            <span class="hero-metric-label">{{ siteContent.playerDetail.summary.bestRank }}</span>
            <strong class="hero-metric-value">{{ siteContent.playerDetail.summary.unranked }}</strong>
          </div>
        </div>
      </div>
    </div>

    <div class="detail-hero-footer">
      <div class="identity-strip" :aria-label="siteContent.playerDetail.header.identityAria">
        <div
          v-for="row in identityRows"
          :key="row.label"
          :class="['identity-row', { compact: row.compact }]"
        >
          <span class="identity-label">{{ row.label }}</span>
          <strong class="identity-value">{{ row.value }}</strong>
        </div>
      </div>

      <div v-if="secondaryRankHighlights.length > 0" class="rank-ribbon">
        <span class="hud-caption">{{ siteContent.playerDetail.summary.ranks }}</span>
        <div class="rank-chip-row">
          <button
            v-for="highlight in secondaryRankHighlights"
            :key="highlight.label"
            type="button"
            class="badge-pill rank-chip"
            @click="selectLeaderboard(highlight.target)"
          >
            {{ highlight.label }}
            <strong>{{ highlight.value }}</strong>
          </button>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.detail-hero {
  position: relative;
  isolation: isolate;
  display: grid;
  gap: 1rem;
  border-color: var(--player-glass-border);
  background: var(--player-glass-bg);
  box-shadow: var(--player-glass-shadow), var(--glass-inset);
  padding: 1.35rem;
}

.detail-header-mark {
  position: absolute;
  right: 1rem;
  top: 1rem;
  color: rgba(141, 184, 255, 0.07);
  font-family: var(--display);
  font-size: clamp(3rem, 10vw, 6rem);
  font-weight: 900;
  letter-spacing: 0.24em;
  line-height: 0.75;
  pointer-events: none;
}

.detail-hero-layout,
.detail-hero-footer {
  position: relative;
  z-index: 1;
}

.detail-hero-layout {
  display: grid;
  grid-template-columns: minmax(190px, 240px) minmax(0, 1fr);
  gap: 1.15rem;
  align-items: stretch;
}

.skin-frame {
  position: relative;
  display: grid;
  height: 100%;
  min-height: 292px;
  place-items: center;
  overflow: hidden;
  border: 1px solid var(--player-glass-border-soft);
  border-radius: 18px;
  background: var(--player-glass-tile-bg);
  box-shadow: var(--glass-inset);
}

.skin-frame::after {
  content: '';
  position: absolute;
  inset: auto 18% 1.4rem;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(141, 184, 255, 0.48), transparent);
}

.detail-hero-main {
  display: grid;
  gap: 0.95rem;
  align-content: center;
  min-width: 0;
}

.detail-kicker-row,
.rank-chip-row {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  align-items: center;
}

.hero-chip {
  display: inline-flex;
  align-items: center;
  min-height: 1.85rem;
  padding: 0.42rem 0.68rem;
  border: 1px solid var(--chip-border);
  border-radius: 999px;
  background: rgba(var(--secondary-rgb), 0.1);
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.68rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.status-chip.online {
  border-color: rgba(var(--secondary-rgb), 0.34);
  background: rgba(var(--secondary-rgb), 0.1);
  color: var(--success);
}

.detail-title-block {
  display: grid;
  gap: 0.55rem;
}

.detail-player-name {
  font-size: clamp(2.2rem, 5vw, 4.1rem);
  line-height: 1;
  overflow-wrap: anywhere;
}

.detail-player-name.online {
  color: var(--success);
}

.detail-player-copy {
  color: var(--text-muted);
  max-width: 62ch;
}

.hero-metric-strip {
  display: grid;
  grid-template-columns: repeat(6, minmax(0, 1fr));
  overflow: hidden;
  border: 1px solid var(--player-glass-border-soft);
  border-radius: 18px;
  background: var(--player-glass-tile-bg);
  box-shadow: var(--glass-inset);
}

.hero-metric {
  display: grid;
  gap: 0.32rem;
  min-width: 0;
  padding: 0.78rem 0.8rem;
  border: 0;
  border-right: 1px solid rgba(255, 255, 255, 0.07);
  background: transparent;
  color: inherit;
  text-align: left;
}

.hero-metric:last-child {
  border-right: 0;
}

.hero-metric-action {
  cursor: pointer;
  transition:
    background var(--transition-fast),
    color var(--transition-fast);
}

.hero-metric-action:hover {
  background: var(--player-glass-tile-bg-hover);
  color: var(--text-strong);
}

.hero-metric-label,
.identity-label {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.64rem;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.hero-metric-value {
  color: var(--text-strong);
  font-size: 0.94rem;
  font-weight: 700;
  line-height: 1.2;
  overflow-wrap: anywhere;
}

.hero-metric-note {
  color: var(--primary);
  font-family: var(--mono);
  font-size: 0.62rem;
  overflow-wrap: anywhere;
}

.detail-hero-footer,
.rank-ribbon {
  display: grid;
  gap: 0.75rem;
}

.identity-strip {
  display: grid;
  grid-template-columns: minmax(0, 1.45fr) repeat(3, minmax(0, 0.85fr));
  gap: 0.55rem;
}

.identity-row {
  display: grid;
  gap: 0.28rem;
  min-width: 0;
  padding: 0.68rem 0.78rem;
  border: 1px solid var(--player-glass-border-soft);
  border-left-color: rgba(141, 184, 255, 0.28);
  border-radius: 14px;
  background: var(--player-glass-tile-bg);
  box-shadow: var(--glass-inset);
}

.identity-value {
  color: var(--text-strong);
  font-size: 0.84rem;
  font-weight: 650;
  overflow-wrap: anywhere;
}

.identity-row.compact .identity-value {
  font-family: var(--mono);
  font-size: 0.72rem;
  font-weight: 500;
}

.rank-chip {
  cursor: pointer;
  color: inherit;
}

.rank-chip:hover {
  border-color: rgba(var(--secondary-rgb), 0.3);
  background: rgba(var(--secondary-rgb), 0.1);
  color: var(--text-strong);
  transform: translateY(-1px);
}

@media (max-width: 1120px) {
  .hero-metric-strip {
    grid-template-columns: repeat(3, minmax(0, 1fr));
  }

  .hero-metric:nth-child(3n) {
    border-right: 0;
  }

  .hero-metric:nth-child(n + 4) {
    border-top: 1px solid rgba(255, 255, 255, 0.07);
  }

  .identity-strip {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (max-width: 780px) {
  .detail-hero {
    padding: 1rem;
  }

  .detail-hero-layout,
  .identity-strip {
    grid-template-columns: 1fr;
  }

  .skin-frame {
    height: 250px;
    min-height: 0;
  }

  .skin-frame :deep(.skin-viewer-container) {
    min-height: 0;
  }

  .skin-frame :deep(canvas) {
    transform: scale(0.84);
    transform-origin: center;
  }
}

@media (max-width: 560px) {
  .hero-metric-strip {
    grid-template-columns: 1fr;
  }

  .hero-metric,
  .hero-metric:nth-child(3n) {
    border-right: 0;
  }

  .hero-metric:nth-child(n + 2) {
    border-top: 1px solid rgba(255, 255, 255, 0.07);
  }

  .detail-player-name {
    font-size: clamp(2rem, 12vw, 3rem);
  }

  .skin-frame {
    height: 220px;
  }

  .skin-frame :deep(canvas) {
    transform: scale(0.76);
  }
}
</style>
