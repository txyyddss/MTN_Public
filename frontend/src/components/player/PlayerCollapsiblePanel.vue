<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    title: string
    sticky?: boolean
  }>(),
  {
    sticky: false
  }
)
</script>

<template>
  <section :class="['glass-card', 'collapsible-panel', 'player-glass-card', 'player-glass-reveal-soft', { sticky: props.sticky }]">
    <header class="panel-header">
      <div class="panel-heading">
        <h3 class="panel-title">{{ props.title }}</h3>
        <div v-if="$slots.summary" class="panel-summary">
          <slot name="summary" />
        </div>
      </div>

    </header>

    <div class="panel-body">
      <slot />
    </div>
  </section>
</template>

<style scoped>
.collapsible-panel {
  display: grid;
  gap: 0.85rem;
  border-color: var(--player-glass-border);
  background: var(--player-glass-bg);
  box-shadow: var(--player-glass-shadow), var(--glass-inset);
}

.collapsible-panel.sticky {
  position: sticky;
  top: 6rem;
}

.panel-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 0.85rem;
}

.panel-heading {
  display: grid;
  gap: 0.45rem;
}

.panel-title {
  font-size: 1.32rem;
}

.panel-summary {
  display: flex;
  flex-wrap: wrap;
  gap: 0.45rem;
}

.panel-body {
  display: grid;
  gap: 0.85rem;
}

@media (max-width: 1024px) {
  .collapsible-panel.sticky {
    position: static;
  }
}

@media (max-width: 720px) {
  .panel-header {
    align-items: flex-start;
    flex-direction: column;
  }
}
</style>
