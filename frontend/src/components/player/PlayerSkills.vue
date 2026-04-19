<script setup lang="ts">
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
  <section v-if="mcmmo" class="glass-card panel-card">
    <div class="panel-head">
      <h3>{{ siteContent.playerDetail.sections.skills }}</h3>
      <div class="meta-cluster">
        <span v-if="ranks.skills" class="meta-chip">#{{ ranks.skills }}</span>
        <span class="meta-chip">{{ siteContent.playerDetail.sections.total }} {{ mcmmo.total }}</span>
      </div>
    </div>

    <div class="skill-grid">
      <SkillItem
        v-for="entry in filteredMcmmo"
        :key="entry.key"
        :name="entry.label"
        :level="entry.level"
        :rank="ranks[`mcmmo:${entry.key}`]"
      />
    </div>
  </section>
</template>

<style scoped>
.panel-card {
  display: grid;
  gap: 1rem;
}

.panel-head {
  display: flex;
  justify-content: space-between;
  gap: 1rem;
  align-items: center;
}

.panel-head h3 {
  font-size: 1.8rem;
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
  border: 1px solid rgba(255, 248, 234, 0.08);
}

.skill-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 0.75rem;
}

@media (max-width: 720px) {
  .skill-grid {
    grid-template-columns: 1fr;
  }
}
</style>
