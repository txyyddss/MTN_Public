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
  <section :class="['glass-card', 'collapsible-panel', 'animate-entry', { sticky: props.sticky, collapsed: !expanded }]">
    <header class="panel-header">
      <div class="panel-heading">
        <h3 class="panel-title">{{ props.title }}</h3>
        <div v-if="$slots.summary" class="panel-summary">
          <slot name="summary" />
        </div>
      </div>

      <button type="button" class="panel-toggle" @click="expanded = !expanded">
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
  gap: 1rem;
}

.collapsible-panel.sticky {
  position: sticky;
  top: 6rem;
}

.panel-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
}

.panel-heading {
  display: grid;
  gap: 0.55rem;
}

.panel-title {
  font-size: 1.4rem;
}

.panel-summary {
  display: flex;
  flex-wrap: wrap;
  gap: 0.45rem;
}

.panel-toggle {
  border: 1px solid var(--glass-border);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.04);
  color: var(--text-main);
  padding: 0.55rem 0.8rem;
  font-family: var(--mono);
  font-size: 0.72rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  transition:
    border-color var(--transition-fast),
    transform var(--transition-fast);
}

.panel-toggle:hover {
  border-color: var(--glass-border-bright);
  transform: translateY(-1px);
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

@media (max-width: 720px) {
  .panel-header {
    align-items: flex-start;
  }
}
</style>
