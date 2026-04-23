<script setup lang="ts">
import PlayerCollapsiblePanel from '@/components/player/PlayerCollapsiblePanel.vue'
import SkillItem from '@/components/SkillItem.vue'
import { useSiteContent } from '@/content/siteContent'
import { createMcmmoLeaderboardTarget } from '@/utils/leaderboards'
import type { FormattedSkillEntry, McMMOSkills } from '@/types/api'

const props = defineProps<{
  mcmmo: McMMOSkills | null
  ranks: Record<string, number>
  filteredMcmmo: FormattedSkillEntry[]
}>()

const emit = defineEmits<{
  (event: 'selectLeaderboard', value: ReturnType<typeof createMcmmoLeaderboardTarget>): void
}>()

const siteContent = useSiteContent()
</script>

<template>
  <PlayerCollapsiblePanel v-if="mcmmo" class="panel-card" :title="siteContent.playerDetail.sections.skills">
    <template #summary>
      <div class="meta-cluster">
        <span class="meta-chip">{{ siteContent.playerDetail.sections.total }} {{ mcmmo.total }}</span>
      </div>
    </template>

    <div v-if="filteredMcmmo.length > 0" class="skill-grid">
      <SkillItem
        v-for="entry in filteredMcmmo"
        :key="entry.key"
        :name="entry.label"
        :level="entry.level"
        :clickable="true"
        :rank="props.ranks[`mcmmo:${entry.key}`]"
        @select="emit('selectLeaderboard', createMcmmoLeaderboardTarget(entry.key, entry.label))"
      />
    </div>
    <p v-else class="empty-copy">{{ siteContent.playerDetail.summary.noSkillData }}</p>
  </PlayerCollapsiblePanel>
</template>

<style scoped>
.panel-card {
  display: grid;
  gap: 0.85rem;
}

.meta-cluster {
  display: flex;
  flex-wrap: wrap;
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

.skill-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 0.65rem;
}

.empty-copy {
  color: var(--text-muted);
}

@media (max-width: 720px) {
  .skill-grid {
    grid-template-columns: 1fr;
  }
}
</style>
