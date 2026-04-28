<script setup lang="ts">
withDefaults(
  defineProps<{
    kicker: string
    title: string
    body: string
    showMark?: boolean
  }>(),
  {
    showMark: true
  }
)
</script>

<template>
  <header :class="['route-hero-header', 'glass-card', 'animate-entry', { 'route-hero-header--no-mark': !showMark }]">
    <div class="route-hero-header__grid">
      <div class="route-hero-header__copy">
        <span class="page-kicker">{{ kicker }}</span>
        <div v-if="showMark" class="route-hero-header__mark" aria-hidden="true">MTN</div>
        <slot name="title">
          <h1 class="route-hero-header__title">{{ title }}</h1>
        </slot>
        <slot name="body">
          <p class="route-hero-header__body">{{ body }}</p>
        </slot>
      </div>

      <div v-if="$slots.actions" class="route-hero-header__actions">
        <slot name="actions" />
      </div>
    </div>
  </header>
</template>

<style scoped>
.route-hero-header {
  position: relative;
  overflow: hidden;
  padding: clamp(1.55rem, 3vw, 2.4rem);
}

.route-hero-header::before {
  content: '';
  position: absolute;
  inset: 0;
  background:
    radial-gradient(circle at 82% 18%, rgba(76, 147, 251, 0.16), transparent 26%),
    linear-gradient(125deg, rgba(76, 147, 251, 0.08), transparent 34%);
  pointer-events: none;
}

.route-hero-header::after {
  content: '';
  position: absolute;
  inset: auto 1.5rem 1.2rem;
  height: 1px;
  background: linear-gradient(90deg, rgba(141, 184, 255, 0.34), rgba(141, 184, 255, 0));
  pointer-events: none;
}

.route-hero-header__grid {
  position: relative;
  z-index: 1;
  display: grid;
  grid-template-columns: minmax(0, 1fr) auto;
  gap: 1.5rem;
  align-items: start;
}

.route-hero-header__copy {
  display: grid;
  gap: 0.65rem;
  max-width: min(100%, 54rem);
}

.route-hero-header__mark {
  color: rgba(76, 147, 251, 0.92);
  font-family: var(--display);
  font-size: clamp(3.2rem, 10vw, 7.4rem);
  font-weight: 900;
  letter-spacing: 0.28em;
  line-height: 0.78;
  text-indent: 0.28em;
  text-shadow: 0 20px 48px rgba(76, 147, 251, 0.18);
}

.route-hero-header__title {
  max-width: 11ch;
  font-size: clamp(3.35rem, 8vw, 6.2rem);
  line-height: 0.9;
}

.route-hero-header__body {
  max-width: 60ch;
  color: var(--text-muted);
  font-size: 1rem;
  line-height: 1.75;
}

.route-hero-header__actions {
  display: flex;
  flex-wrap: wrap;
  justify-content: flex-end;
  align-items: flex-start;
  gap: 0.65rem;
}

.route-hero-header--no-mark .route-hero-header__title {
  max-width: 12ch;
  font-size: clamp(3rem, 6vw, 5.2rem);
}

@media (max-width: 860px) {
  .route-hero-header__grid {
    grid-template-columns: 1fr;
  }

  .route-hero-header__actions {
    justify-content: flex-start;
  }

  .route-hero-header__mark {
    font-size: clamp(2.6rem, 16vw, 4.5rem);
  }

  .route-hero-header__title,
  .route-hero-header--no-mark .route-hero-header__title {
    max-width: 100%;
    font-size: clamp(2.7rem, 12vw, 4.6rem);
  }
}
</style>
