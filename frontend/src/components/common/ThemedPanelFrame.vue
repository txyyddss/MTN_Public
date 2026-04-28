<script setup lang="ts">
withDefaults(
  defineProps<{
    tag?: string
    kicker?: string
    title?: string
    body?: string
    status?: string
    statusTone?: 'default' | 'live' | 'success' | 'warning'
    compact?: boolean
  }>(),
  {
    tag: 'article',
    kicker: undefined,
    title: undefined,
    body: undefined,
    status: undefined,
    statusTone: 'default',
    compact: false
  }
)
</script>

<template>
  <component :is="tag" :class="['themed-panel-frame', 'glass-card', { 'themed-panel-frame--compact': compact }]" data-panel-mark="MTN">
    <header
      v-if="kicker || title || body || status || $slots.header || $slots.actions"
      class="themed-panel-frame__header"
    >
      <div class="themed-panel-frame__copy">
        <slot name="header">
          <span v-if="kicker" class="section-kicker">{{ kicker }}</span>
          <h3 v-if="title" class="themed-panel-frame__title">{{ title }}</h3>
          <p v-if="body" class="themed-panel-frame__body">{{ body }}</p>
        </slot>
      </div>

      <div v-if="status || $slots.actions" class="themed-panel-frame__side">
        <slot name="actions" />
        <span v-if="status" :class="['themed-panel-frame__status', `is-${statusTone}`]">
          {{ status }}
        </span>
      </div>
    </header>

    <div class="themed-panel-frame__content">
      <slot />
    </div>
  </component>
</template>

<style scoped>
.themed-panel-frame {
  position: relative;
  isolation: isolate;
  display: grid;
  gap: 1rem;
  min-height: 100%;
}

.themed-panel-frame::after {
  content: attr(data-panel-mark);
  position: absolute;
  top: 1rem;
  right: 1.25rem;
  color: rgba(141, 184, 255, 0.07);
  font-family: var(--mono);
  font-size: clamp(1.9rem, 4vw, 3.1rem);
  font-weight: 700;
  letter-spacing: 0.3em;
  pointer-events: none;
}

.themed-panel-frame__header {
  position: relative;
  z-index: 1;
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 1rem;
}

.themed-panel-frame__copy {
  display: grid;
  gap: 0.42rem;
  min-width: 0;
}

.themed-panel-frame__title {
  font-size: clamp(1.7rem, 3vw, 2.15rem);
}

.themed-panel-frame__body {
  color: var(--text-muted);
  font-size: 0.94rem;
  line-height: 1.7;
}

.themed-panel-frame__side {
  display: flex;
  flex-wrap: wrap;
  justify-content: flex-end;
  align-items: flex-start;
  gap: 0.65rem;
}

.themed-panel-frame__content {
  position: relative;
  z-index: 1;
  display: grid;
  gap: 0.9rem;
}

.themed-panel-frame__status {
  display: inline-flex;
  align-items: center;
  min-height: 2.2rem;
  padding: 0.45rem 0.8rem;
  border-radius: 999px;
  border: 1px solid var(--chip-border);
  background: var(--chip-bg);
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.68rem;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.themed-panel-frame__status.is-live,
.themed-panel-frame__status.is-success {
  color: var(--success);
}

.themed-panel-frame__status.is-warning {
  color: var(--warning);
}

.themed-panel-frame--compact {
  gap: 0.8rem;
}

@media (max-width: 720px) {
  .themed-panel-frame__header {
    flex-direction: column;
  }

  .themed-panel-frame__side {
    justify-content: flex-start;
  }
}
</style>
