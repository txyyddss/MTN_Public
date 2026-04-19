<script setup lang="ts">
interface Props {
  body: string
  caption: string
  title: string
}

const props = defineProps<Props>()
</script>

<template>
  <article class="glass-card community-card animate-entry delay-300">
    <span class="section-kicker">Community</span>
    <h2 class="community-title">{{ props.title }}</h2>
    <p class="community-body">{{ props.body }}</p>
    <div class="community-foot">{{ props.caption }}</div>
  </article>
</template>

<style scoped>
.community-card {
  display: grid;
  gap: 1rem;
}

.community-title {
  font-size: clamp(2rem, 4vw, 3.3rem);
  max-width: 12ch;
}

.community-body,
.community-foot {
  color: var(--text-muted);
}
</style>
