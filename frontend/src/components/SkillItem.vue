<script setup lang="ts">
const props = defineProps<{
  name: string
  level: number
  rank?: number
}>()

const maxLevel = 1000
const progress = Math.min((props.level / maxLevel) * 100, 100)
</script>

<template>
  <div class="skill-row">
    <span class="skill-name">{{ name }}</span>
    <div class="skill-bar-wrap">
      <div class="skill-bar-fill" :style="{ width: `${progress}%` }"></div>
    </div>
    <span class="skill-level">{{ level }}</span>
    <span v-if="rank" class="skill-rank">#{{ rank }}</span>
  </div>
</template>

<style scoped>
.skill-row {
  display: grid;
  grid-template-columns: 90px 1fr 42px 42px;
  gap: 0.6rem;
  align-items: center;
  padding: 0.72rem 0.85rem;
  border-radius: 16px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.skill-name {
  color: var(--text-muted);
  font-size: 0.76rem;
}

.skill-bar-wrap {
  height: 6px;
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.08);
  overflow: hidden;
}

.skill-bar-fill {
  height: 100%;
  border-radius: 999px;
  background: linear-gradient(90deg, var(--secondary), var(--primary));
}

.skill-level,
.skill-rank {
  text-align: right;
  font-family: var(--mono);
}

.skill-level {
  color: var(--text-strong);
}

.skill-rank {
  color: var(--primary);
  font-size: 0.72rem;
}
</style>
