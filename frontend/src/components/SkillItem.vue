<script setup lang="ts">
const props = defineProps<{
  name: string
  level: number
  rank?: number
}>()

// McMMO max level is 1000 for most skills
const MAX_LEVEL = 1000
const progress = Math.min((props.level / MAX_LEVEL) * 100, 100)
</script>

<template>
  <div class="skill-row">
    <span class="skill-name">{{ name }}</span>
    <div class="skill-bar-wrap">
      <div class="skill-bar-fill" :style="{ width: progress + '%' }"></div>
    </div>
    <span class="skill-level">{{ level }}</span>
    <span class="skill-rank" v-if="rank">#{{ rank }}</span>
  </div>
</template>

<style scoped>
.skill-row {
  display: grid;
  grid-template-columns: 90px 1fr 38px 36px;
  align-items: center;
  gap: 8px;
  padding: 5px 10px;
  border-radius: 6px;
  background: rgba(255, 255, 255, 0.025);
  border: 1px solid transparent;
  transition: background 0.15s, border-color 0.15s;
}

.skill-row:hover {
  background: rgba(255, 255, 255, 0.05);
  border-color: var(--glass-border);
}

.skill-name {
  font-size: 0.78rem;
  color: var(--text-muted);
  line-height: 1.2;
}

.skill-bar-wrap {
  height: 4px;
  background: rgba(255, 255, 255, 0.08);
  border-radius: 2px;
  overflow: hidden;
}

.skill-bar-fill {
  height: 100%;
  background: linear-gradient(90deg, #3b82f6, #60a5fa);
  border-radius: 2px;
  transition: width 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.skill-level {
  font-size: 0.82rem;
  font-weight: 700;
  color: #fff;
  text-align: right;
  font-variant-numeric: tabular-nums;
}

.skill-rank {
  font-size: 0.65rem;
  font-weight: 700;
  color: #fcd34d;
  text-align: right;
  white-space: nowrap;
}
</style>
