<script setup lang="ts">
withDefaults(
  defineProps<{
    label?: string
    compact?: boolean
  }>(),
  {
    label: 'MTN DOSSIER',
    compact: false
  }
)
</script>

<template>
  <div :class="['dossier-loading-panel', 'glass-card', { 'dossier-loading-panel--compact': compact }]">
    <span class="dossier-loading-panel__label" aria-hidden="true">{{ label }}</span>
    <div class="dossier-loading-panel__grid" aria-hidden="true"></div>
    <div class="dossier-loading-panel__scan" aria-hidden="true"></div>
    <div class="dossier-loading-panel__content">
      <slot />
    </div>
  </div>
</template>

<style scoped>
.dossier-loading-panel {
  position: relative;
  isolation: isolate;
  display: grid;
  gap: 0.85rem;
  min-height: 100%;
  border-color: var(--glass-border);
  background:
    radial-gradient(circle at 18% 0%, rgba(var(--secondary-rgb), 0.18), transparent 32%),
    linear-gradient(180deg, rgba(255, 255, 255, 0.05), rgba(255, 255, 255, 0.015)),
    rgba(3, 8, 24, 0.82);
}

.dossier-loading-panel::before {
  background:
    linear-gradient(180deg, rgba(255, 255, 255, 0.055), transparent 34%),
    linear-gradient(120deg, rgba(var(--secondary-rgb), 0.16), transparent 46%);
}

.dossier-loading-panel__label {
  position: absolute;
  inset: 1rem 1.15rem auto auto;
  z-index: 1;
  color: rgba(141, 184, 255, 0.22);
  font-family: var(--mono);
  font-size: 0.62rem;
  font-weight: 700;
  letter-spacing: 0.22em;
  text-transform: uppercase;
  pointer-events: none;
}

.dossier-loading-panel__grid,
.dossier-loading-panel__scan {
  position: absolute;
  inset: 0;
  pointer-events: none;
}

.dossier-loading-panel__grid {
  opacity: 0.36;
  background-image:
    linear-gradient(rgba(255, 255, 255, 0.04) 1px, transparent 1px),
    linear-gradient(90deg, rgba(255, 255, 255, 0.04) 1px, transparent 1px),
    radial-gradient(circle at center, rgba(var(--secondary-rgb), 0.14) 1px, transparent 1px);
  background-size: 32px 32px;
  mask-image: linear-gradient(180deg, transparent, #000 12%, #000 88%, transparent);
}

.dossier-loading-panel__scan {
  inset: -35% -45%;
  background:
    linear-gradient(118deg, transparent 34%, rgba(var(--secondary-rgb), 0.22) 50%, transparent 66%),
    linear-gradient(118deg, transparent 42%, rgba(255, 255, 255, 0.1) 50%, transparent 58%);
  transform: translateX(-55%);
  animation: dossierSweep 1.75s linear infinite;
}

.dossier-loading-panel__content {
  position: relative;
  z-index: 1;
  display: grid;
  gap: 0.85rem;
}

.dossier-loading-panel--compact .dossier-loading-panel__label {
  inset: 0.8rem 1.15rem auto auto;
}

@keyframes dossierSweep {
  from {
    transform: translateX(-55%);
  }

  to {
    transform: translateX(55%);
  }
}

@media (prefers-reduced-motion: reduce) {
  .dossier-loading-panel__scan {
    animation: none;
    opacity: 0.3;
    transform: none;
  }
}
</style>
