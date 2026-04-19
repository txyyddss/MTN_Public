<script setup lang="ts">
import { computed } from 'vue'

import PlayerPieChart from '@/components/PlayerPieChart.vue'
import PlayerSidebar from '@/components/player/PlayerSidebar.vue'
import PlayerSkills from '@/components/player/PlayerSkills.vue'
import { useRevealOnScroll } from '@/composables/useRevealOnScroll'
import { siteContent } from '@/content/siteContent'
import type { FormattedSkillEntry, LinkedAccount, McMMOSkills, PlayerInfo, OreStat } from '@/types/api'

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
  rankHighlights: Array<{ label: string; value: string }>
  formatDate: (value: number) => string
  formatPlaytime: (value: number) => string
}>()

const { revealed } = useRevealOnScroll<HTMLElement>('overviewSection')

const linkedAccountName = computed(() => {
  if (!props.linkedAccount) {
    return ''
  }

  if (props.info.type === 'Bedrock') {
    return props.linkedAccount.java_username
  }

  return props.linkedAccount.bedrock_username || props.linkedAccount.java_username
})

const overviewMetrics = computed(() => [
  { label: siteContent.playerDetail.profile.firstJoin, value: props.formatDate(props.info.first_played) },
  { label: siteContent.playerDetail.profile.lastSeen, value: props.formatDate(props.info.last_seen) },
  {
    label: siteContent.playerDetail.summary.advancements,
    value: `${props.completedAdvancements}/${props.totalAdvancements || 0}`
  },
  {
    label: siteContent.playerDetail.summary.ranks,
    value: props.rankHighlights[0]?.value ?? 'Unranked'
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
        :format-date="formatDate"
        :format-playtime="formatPlaytime"
      />
    </div>

    <div class="overview-stack">
      <article class="glass-card overview-summary hover-rise">
        <div class="overview-summary-head">
          <span class="hud-kicker">{{ siteContent.playerDetail.tabs.overview }}</span>
          <div class="overview-badges">
            <span class="badge-pill"><strong>{{ info.type }}</strong></span>
            <span :class="['badge-pill', 'overview-status', { online: isOnline }]">
              <strong>{{ isOnline ? siteContent.playerDetail.summary.onlineNow : siteContent.playerDetail.summary.archiveRecord }}</strong>
            </span>
            <span v-if="linkedAccountName" class="badge-pill">
              {{ siteContent.playerDetail.summary.linkedTo }}
              <strong>{{ linkedAccountName }}</strong>
            </span>
          </div>
        </div>

        <div class="hud-metric-grid overview-metric-grid">
          <article v-for="metric in overviewMetrics" :key="metric.label" class="hud-metric-card">
            <span class="hud-metric-label">{{ metric.label }}</span>
            <strong class="hud-metric-value">{{ metric.value }}</strong>
          </article>
        </div>

        <div v-if="rankHighlights.length > 1" class="overview-rank-row">
          <span class="hud-caption">{{ siteContent.playerDetail.summary.ranks }}</span>
          <div class="overview-rank-chips">
            <span v-for="highlight in rankHighlights.slice(1)" :key="highlight.label" class="badge-pill">
              {{ highlight.label }}
              <strong>{{ highlight.value }}</strong>
            </span>
          </div>
        </div>
      </article>

      <PlayerPieChart class="hover-rise" :ore-stats="oreStats" />
      <PlayerSkills class="hover-rise" :mcmmo="mcmmo" :ranks="ranks" :filtered-mcmmo="filteredMcmmo" />
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

.overview-summary-head {
  display: grid;
  gap: 0.6rem;
}

.overview-badges,
.overview-rank-chips {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.overview-status.online strong {
  color: #8fe3b3;
}

.overview-metric-grid {
  grid-template-columns: repeat(4, minmax(0, 1fr));
}

.overview-rank-row {
  display: grid;
  gap: 0.5rem;
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
