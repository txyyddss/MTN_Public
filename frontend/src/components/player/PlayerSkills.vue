<script setup lang="ts">
import SkillItem from '@/components/SkillItem.vue'

defineProps<{
  mcmmo: any
  ranks: any
  filteredMcmmo: Record<string, any>
}>()
</script>

<template>
  <section class="panel glass-card" v-if="mcmmo">
    <div class="panel-header-simple">
      <h3><img src="/icons/monsters_hunted.ico" class="header-icon" /> McMMO Skills</h3>
      <div class="rank-badge" v-if="ranks.skills">Rank #{{ ranks.skills }}</div>
      <div class="total-badge">Total {{ mcmmo.total }}</div>
    </div>
    <div class="skill-grid">
      <SkillItem 
        v-for="(level, skill) in filteredMcmmo" 
        :key="skill"
        :name="String(skill)"
        :level="Number(level)"
        :rank="ranks['mcmmo:' + skill.toLowerCase()]"
      />
    </div>
  </section>
</template>

<style scoped>
.panel h3 {
  display: flex;
  align-items: center;
  gap: 12px;
  margin: 0;
}

.header-icon {
  width: 24px;
  height: 24px;
}

.panel-header-simple {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.rank-badge {
  padding: 4px 12px;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid var(--primary);
  border-radius: 4px;
  font-size: 0.8rem;
  font-weight: 700;
  color: var(--primary);
}

.total-badge {
  font-weight: 800;
  color: var(--text-muted);
  font-size: 0.9rem;
}

.skill-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 1.5rem;
}

@media (max-width: 600px) {
  .skill-grid {
    grid-template-columns: 1fr 1fr;
  }
}
</style>
