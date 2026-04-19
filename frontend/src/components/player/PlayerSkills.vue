<script setup lang="ts">
import PlayerCollapsiblePanel from '@/components/player/PlayerCollapsiblePanel.vue'
import SkillItem from '@/components/SkillItem.vue'
import { siteContent } from '@/content/siteContent'
import type { FormattedSkillEntry, McMMOSkills } from '@/types/api'

defineProps<{
  mcmmo: McMMOSkills | null
  ranks: Record<string, number>
  filteredMcmmo: FormattedSkillEntry[]
}>()
</script>

<template>
  <PlayerCollapsiblePanel v-if="mcmmo" class="panel-card" :title="siteContent.playerDetail.sections.skills">
    <template #summary>
      <div class="meta-cluster">
        <span v-if="ranks.skills" class="meta-chip">#{{ ranks.skills }}</span>
        <span class="meta-chip">{{ siteContent.playerDetail.sections.total }} {{ mcmmo.total }}</span>
      </div>
    </template>

    <div v-if="filteredMcmmo.length > 0" class="skill-grid">
      <SkillItem
        v-for="entry in filteredMcmmo"
        :key="entry.key"
        :name="entry.label"
        :level="entry.level"
        :rank="ranks[`mcmmo:${entry.key}`]"
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
  font-size: 0.74rem;
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
