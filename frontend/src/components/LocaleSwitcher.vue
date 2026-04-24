<script setup lang="ts">
import { setLocale, useCurrentLocale, useLocaleOptions, useSiteContent } from '@/content/siteContent'

const props = withDefaults(
  defineProps<{
    mobile?: boolean
  }>(),
  {
    mobile: false
  }
)

const currentLocale = useCurrentLocale()
const localeOptions = useLocaleOptions()
const siteContent = useSiteContent()
</script>

<template>
  <div :class="['locale-switcher', { mobile: props.mobile }]" :aria-label="siteContent.app.locale.label">
    <span v-if="props.mobile" class="locale-label">{{ siteContent.app.locale.label }}</span>

    <div class="locale-track">
      <button
        v-for="option in localeOptions"
        :key="option.value"
        type="button"
        :class="['locale-button', 'action-inline', 'action-press', { active: currentLocale === option.value }]"
        @click="setLocale(option.value)"
      >
        {{ option.label }}
      </button>
    </div>
  </div>
</template>

<style scoped>
.locale-switcher {
  display: inline-flex;
  align-items: center;
  gap: 0.6rem;
}

.locale-switcher.mobile {
  width: 100%;
  justify-content: space-between;
  padding-top: 0.4rem;
}

.locale-label {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.68rem;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.locale-track {
  display: inline-flex;
  align-items: center;
  gap: 0.18rem;
  padding: 0.22rem;
  border-radius: 999px;
  border: 1px solid var(--control-border);
  background: var(--control-surface);
}

.locale-button {
  min-width: 3rem;
  padding: 0.42rem 0.7rem;
  border: 1px solid transparent;
  border-radius: 999px;
  background: transparent;
  color: var(--control-text);
  font-family: var(--mono);
  font-size: 0.72rem;
  font-weight: 500;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  transition:
    color var(--transition-fast),
    background var(--transition-fast),
    border-color var(--transition-fast),
    transform var(--transition-fast);
}

.locale-button:hover {
  color: var(--control-text-hover);
  background: var(--control-bg-hover);
  transform: translateY(-1px);
}

.locale-button.active {
  color: var(--control-text-active);
  border-color: var(--control-border-active);
  background: var(--control-bg-active);
}
</style>
