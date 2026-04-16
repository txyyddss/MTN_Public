<script setup lang="ts">
const props = defineProps<{
  name: string
  level: number
  rank?: number
}>()
</script>

<template>
  <div class="skill-item">
    <div class="skill-info">
      <span class="skill-name">{{ name }}</span>
      <span class="skill-level">{{ level }}</span>
    </div>
    <div class="skill-rank" v-if="rank">
      #{{ rank }}
    </div>
  </div>
</template>

<style scoped>
.skill-item {
  background: rgba(255, 255, 255, 0.03);
  padding: 8px 12px;
  border-radius: 8px;
  border-left: 3px solid var(--primary);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.skill-info {
  display: flex;
  flex-direction: column;
}

.skill-name {
  color: var(--text-muted);
  font-size: 0.75rem;
  display: block;
  line-height: 1;
  margin-bottom: 2px;
}

.skill-level {
  font-weight: 700;
  color: #fff;
  font-size: 1rem;
  line-height: 1.2;
}

.skill-rank {
  font-size: 0.7rem;
  font-weight: 700;
  color: #fcd34d;
}
</style>
