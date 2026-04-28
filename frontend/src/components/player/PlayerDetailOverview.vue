<script setup lang="ts">
import HourlyPresenceHeatmap from '@/components/heatmap/HourlyPresenceHeatmap.vue'
import PlayerPieChart from '@/components/PlayerPieChart.vue'
import PlayerCollapsiblePanel from '@/components/player/PlayerCollapsiblePanel.vue'
import PlayerSkills from '@/components/player/PlayerSkills.vue'
import { useRevealOnScroll } from '@/composables/useRevealOnScroll'
import { useSiteContent } from '@/content/siteContent'
import type { FormattedSkillEntry, LeaderboardTarget, McMMOSkills, OreStat, PlayerOnlineHeatmap } from '@/types/api'

defineProps<{
  oreStats: OreStat[]
  mcmmo: McMMOSkills | null
  ranks: Record<string, number>
  filteredMcmmo: FormattedSkillEntry[]
  onlineHeatmap: PlayerOnlineHeatmap | null
}>()

const emit = defineEmits<{
  (event: 'selectLeaderboard', value: LeaderboardTarget): void
}>()

const { revealed } = useRevealOnScroll<HTMLElement>('overviewSection')
const siteContent = useSiteContent()
</script>

<template>
  <section
    ref="overviewSection"
    :class="['overview-dashboard', 'scroll-reveal', { 'is-revealed': revealed }]"
  >
    <PlayerPieChart class="overview-panel overview-panel-ores hover-rise" :ore-stats="oreStats" />

    <PlayerCollapsiblePanel
      class="overview-panel overview-panel-history hover-rise"
      :title="siteContent.playerDetail.sections.onlineHistory"
    >
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
      class="overview-panel overview-panel-skills hover-rise"
      :mcmmo="mcmmo"
      :ranks="ranks"
      :filtered-mcmmo="filteredMcmmo"
      @select-leaderboard="emit('selectLeaderboard', $event)"
    />
  </section>
</template>

<style scoped>
.overview-dashboard {
  display: grid;
  grid-template-columns: minmax(0, 1fr) minmax(320px, 0.82fr);
  gap: 0.9rem;
  align-items: start;
}

.overview-panel {
  min-width: 0;
}

.overview-panel-skills {
  grid-column: 1 / -1;
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

@media (max-width: 1080px) {
  .overview-dashboard {
    grid-template-columns: 1fr;
  }

  .overview-panel-skills {
    grid-column: auto;
  }
}
</style>
