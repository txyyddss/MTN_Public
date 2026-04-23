<script setup lang="ts">
import { computed } from 'vue'

import HourlyPresenceHeatmap from '@/components/heatmap/HourlyPresenceHeatmap.vue'
import PlayerPieChart from '@/components/PlayerPieChart.vue'
import PlayerCollapsiblePanel from '@/components/player/PlayerCollapsiblePanel.vue'
import PlayerSidebar from '@/components/player/PlayerSidebar.vue'
import PlayerSkills from '@/components/player/PlayerSkills.vue'
import { useRevealOnScroll } from '@/composables/useRevealOnScroll'
import { useSiteContent } from '@/content/siteContent'
import type { LeaderboardTarget } from '@/types/api'
import type { FormattedSkillEntry, LinkedAccount, McMMOSkills, PlayerInfo, OreStat, PlayerOnlineHeatmap } from '@/types/api'
import type { PlayerRankHighlight } from '@/types/playerDetail'

const props = defineProps<{
  info: PlayerInfo
  isOnline: boolean
  linkedAccount: LinkedAccount | null
  oreStats: OreStat[]
  mcmmo: McMMOSkills | null
  ranks: Record<string, number>
  filteredMcmmo: FormattedSkillEntry[]
  completedAdvancements: number
  totalAdvancements: number
  onlineHeatmap: PlayerOnlineHeatmap | null
  rankHighlights: PlayerRankHighlight[]
  topRankHighlight: PlayerRankHighlight | null
  formatDate: (value: number) => string
  formatDateTime: (value: number) => string
  formatPlaytime: (value: number) => string
}>()

const emit = defineEmits<{
  (event: 'selectLeaderboard', value: LeaderboardTarget): void
}>()

const { revealed } = useRevealOnScroll<HTMLElement>('overviewSection')
const siteContent = useSiteContent()

const overviewMetrics = computed(() => [
  { label: siteContent.value.playerDetail.profile.firstJoin, value: props.formatDate(props.info.first_played), target: null },
  { label: siteContent.value.playerDetail.profile.lastSeen, value: props.formatDateTime(props.info.last_seen), target: null },
  {
    label: siteContent.value.playerDetail.summary.advancements,
    value: `${props.completedAdvancements}/${props.totalAdvancements || 0}`,
    target: null
  },
  {
    label: siteContent.value.playerDetail.summary.skillLeaderboard,
    value: props.topRankHighlight?.value ?? siteContent.value.playerDetail.summary.unranked,
    target: props.topRankHighlight?.target ?? null
  }
])
</script>

<template>
  <section
    ref="overviewSection"
    :class="['overview-layout', 'scroll-reveal', { 'is-revealed': revealed }]"
  >
    <div class="overview-sidebar">
      <PlayerSidebar
        class="hover-rise"
        sticky
        :info="info"
        :is-online="isOnline"
        :linked-account="linkedAccount"
        :format-playtime="formatPlaytime"
      />
    </div>

    <div class="overview-stack">
      <article class="glass-card overview-summary hover-rise">
        <div class="overview-summary-head">
          <span class="hud-kicker">{{ siteContent.playerDetail.tabs.overview }}</span>
        </div>

        <div class="hud-metric-grid overview-metric-grid">
          <component
            :is="metric.target ? 'button' : 'article'"
            v-for="metric in overviewMetrics"
            :key="metric.label"
            :class="['hud-metric-card', { 'metric-action': Boolean(metric.target) }]"
            type="button"
            @click="metric.target && emit('selectLeaderboard', metric.target)"
          >
            <span class="hud-metric-label">{{ metric.label }}</span>
            <strong class="hud-metric-value">{{ metric.value }}</strong>
          </component>
        </div>

        <div v-if="rankHighlights.length > 1" class="overview-rank-row">
          <span class="hud-caption">{{ siteContent.playerDetail.summary.ranks }}</span>
          <div class="overview-rank-chips">
            <button
              v-for="highlight in rankHighlights.slice(1)"
              :key="highlight.label"
              type="button"
              class="badge-pill overview-rank-button"
              @click="emit('selectLeaderboard', highlight.target)"
            >
              {{ highlight.label }}
              <strong>{{ highlight.value }}</strong>
            </button>
          </div>
        </div>
      </article>

      <PlayerPieChart class="hover-rise" :ore-stats="oreStats" />
      <PlayerCollapsiblePanel class="panel-card hover-rise" :title="siteContent.playerDetail.sections.onlineHistory">
        <template #summary>
          <span class="meta-chip">{{ onlineHeatmap?.timezone ?? siteContent.playerDetail.timezoneFallback }}</span>
        </template>

        <HourlyPresenceHeatmap
          v-if="onlineHeatmap"
          :days="onlineHeatmap.days"
          :hours="onlineHeatmap.hours"
          :cells="onlineHeatmap.cells"
          :timezone="onlineHeatmap.timezone"
          variant="player"
        />
        <p v-else class="empty-copy">{{ siteContent.playerDetail.summary.noHistoryData }}</p>
      </PlayerCollapsiblePanel>
      <PlayerSkills
        class="hover-rise"
        :mcmmo="mcmmo"
        :ranks="ranks"
        :filtered-mcmmo="filteredMcmmo"
        @select-leaderboard="emit('selectLeaderboard', $event)"
      />
    </div>
  </section>
</template>

<style scoped>
.overview-layout {
  display: grid;
  grid-template-columns: minmax(280px, 320px) minmax(0, 1fr);
  gap: 0.9rem;
  align-items: start;
}

.overview-sidebar,
.overview-stack {
  display: grid;
  gap: 0.9rem;
  align-items: start;
}

.overview-summary {
  display: grid;
  gap: 0.85rem;
}

.panel-card {
  display: grid;
  gap: 0.85rem;
}

.overview-summary-head {
  display: grid;
  gap: 0.6rem;
}

.overview-rank-chips {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.overview-metric-grid {
  grid-template-columns: repeat(4, minmax(0, 1fr));
}

.overview-rank-row {
  display: grid;
  gap: 0.5rem;
}

.meta-chip {
  padding: 0.45rem 0.7rem;
  border-radius: 999px;
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.68rem;
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.empty-copy {
  color: var(--text-muted);
}

.metric-action,
.overview-rank-button {
  cursor: pointer;
  color: inherit;
  text-align: left;
  transition:
    transform var(--transition-fast),
    border-color var(--transition-fast),
    background var(--transition-fast);
}

.metric-action {
  width: 100%;
}

.metric-action:hover,
.overview-rank-button:hover {
  transform: translateY(-1px);
  border-color: rgba(76, 147, 251, 0.28);
  background: rgba(255, 255, 255, 0.05);
}

@media (max-width: 1080px) {
  .overview-layout {
    grid-template-columns: 1fr;
  }

  .overview-metric-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (max-width: 640px) {
  .overview-metric-grid {
    grid-template-columns: 1fr;
  }
}
</style>
