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
  <div :class="['dossier-loading-panel', 'glass-card', { 'dossier-loading-panel--compact': compact }]" :data-label="label">
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
  background:
    linear-gradient(180deg, rgba(255, 255, 255, 0.05), rgba(255, 255, 255, 0.015)),
    rgba(10, 12, 18, 0.88);
}

.dossier-loading-panel::after {
  content: attr(data-label);
  position: absolute;
  top: 1rem;
  right: 1.15rem;
  color: rgba(141, 184, 255, 0.16);
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
  opacity: 0.3;
  background-image:
    linear-gradient(rgba(255, 255, 255, 0.04) 1px, transparent 1px),
    linear-gradient(90deg, rgba(255, 255, 255, 0.04) 1px, transparent 1px);
  background-size: 32px 32px;
  mask-image: linear-gradient(180deg, transparent, #000 12%, #000 88%, transparent);
}

.dossier-loading-panel__scan {
  inset: -35% -45%;
  background: linear-gradient(118deg, transparent 36%, rgba(76, 147, 251, 0.16) 50%, transparent 64%);
  transform: translateX(-55%);
  animation: dossierSweep 1.9s linear infinite;
}

.dossier-loading-panel__content {
  position: relative;
  z-index: 1;
  display: grid;
  gap: 0.85rem;
}

.dossier-loading-panel--compact::after {
  top: 0.8rem;
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
