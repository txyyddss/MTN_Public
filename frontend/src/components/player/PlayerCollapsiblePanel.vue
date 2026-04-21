<script setup lang="ts">
import { shallowRef } from 'vue'

const props = withDefaults(
  defineProps<{
    title: string
    defaultExpanded?: boolean
    sticky?: boolean
  }>(),
  {
    defaultExpanded: true,
    sticky: false
  }
)

const expanded = shallowRef(props.defaultExpanded)
</script>

<template>
  <section :class="['glass-card', 'collapsible-panel', 'animate-entry-soft', { sticky: props.sticky, collapsed: !expanded }]">
    <header class="panel-header">
      <div class="panel-heading">
        <h3 class="panel-title">{{ props.title }}</h3>
        <div v-if="$slots.summary" class="panel-summary">
          <slot name="summary" />
        </div>
      </div>

      <button type="button" class="panel-toggle" :aria-expanded="expanded" @click="expanded = !expanded">
        {{ expanded ? 'Collapse' : 'Expand' }}
      </button>
    </header>

    <div :class="['panel-body-shell', { collapsed: !expanded }]">
      <div class="panel-body">
        <slot />
      </div>
    </div>
  </section>
</template>

<style scoped>
.collapsible-panel {
  display: grid;
  gap: 0.85rem;
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

.panel-toggle {
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.04);
  color: var(--text-main);
  padding: 0.48rem 0.72rem;
  font-family: var(--sans);
  font-size: 0.82rem;
  font-weight: 500;
  letter-spacing: -0.01em;
  transition:
    border-color var(--transition-fast),
    transform var(--transition-fast),
    background var(--transition-fast);
}

.panel-toggle:hover {
  border-color: rgba(76, 147, 251, 0.28);
  transform: translateY(-1px);
  background: rgba(255, 255, 255, 0.07);
}

.panel-body-shell {
  display: grid;
  grid-template-rows: 1fr;
  opacity: 1;
  transition:
    grid-template-rows 0.3s ease,
    opacity 0.2s ease;
}

.panel-body-shell.collapsed {
  grid-template-rows: 0fr;
  opacity: 0.2;
}

.panel-body {
  min-height: 0;
  overflow: hidden;
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
