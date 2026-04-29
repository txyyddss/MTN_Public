<script setup lang="ts">
import ThemedPanelFrame from '@/components/common/ThemedPanelFrame.vue'

const token = defineModel<string>('token', { required: true })

defineProps<{
  loading: boolean
  count: number
  javaCount: number
  bedrockCount: number
  error: string
  notice: string
}>()

const emit = defineEmits<{
  refresh: []
}>()
</script>

<template>
  <ThemedPanelFrame tag="section" class="token-panel" kicker="Access">
    <template #actions>
      <button class="btn-primary" type="button" :disabled="loading || !token.trim()" @click="emit('refresh')">
        {{ loading ? 'Refreshing' : 'Refresh' }}
      </button>
    </template>

    <div class="token-main">
      <label class="token-field">
        <span>API Token</span>
        <input
          v-model="token"
          class="action-field"
          type="password"
          autocomplete="off"
          placeholder="Bearer token"
        />
      </label>
    </div>

    <div class="token-metrics">
      <div class="metric-tile">
        <span>Total</span>
        <strong>{{ count }}</strong>
      </div>
      <div class="metric-tile">
        <span>Java</span>
        <strong>{{ javaCount }}</strong>
      </div>
      <div class="metric-tile">
        <span>Bedrock</span>
        <strong>{{ bedrockCount }}</strong>
      </div>
    </div>

    <p v-if="error" class="state-line error">{{ error }}</p>
    <p v-else-if="notice" class="state-line success">{{ notice }}</p>
  </ThemedPanelFrame>
</template>

<style scoped>
.token-panel {
  display: grid;
  gap: 1rem;
}

.token-main,
.token-field {
  display: grid;
  gap: 0.65rem;
  min-width: 0;
}

.token-field span {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.72rem;
  letter-spacing: 0;
  text-transform: uppercase;
}

.token-field input {
  width: 100%;
  min-height: 3.05rem;
  padding: 0 1rem;
  border: 1px solid var(--control-border);
  border-radius: var(--radius-sm);
  background: var(--control-bg);
  color: var(--text-main);
}

.token-field input:focus {
  outline: none;
  border-color: var(--control-border-active);
  background: var(--control-bg-hover);
}

.token-metrics {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 0.75rem;
}

.metric-tile {
  display: grid;
  gap: 0.25rem;
  min-height: 4.25rem;
  padding: 0.8rem 0.9rem;
  border: 1px solid var(--glass-border-soft);
  border-radius: var(--radius-md);
  background:
    linear-gradient(135deg, rgba(var(--secondary-rgb), 0.08), transparent 42%),
    rgba(255, 255, 255, 0.035);
  box-shadow: var(--glass-inset);
}

.metric-tile span {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.68rem;
  letter-spacing: 0;
  text-transform: uppercase;
}

.metric-tile strong {
  color: var(--text-strong);
  font-size: 1.4rem;
  line-height: 1;
}

.state-line {
  padding: 0.75rem 0.85rem;
  border-radius: var(--radius-sm);
  font-size: 0.9rem;
}

.state-line.error {
  border: 1px solid rgba(var(--primary-rgb), 0.28);
  background: rgba(var(--primary-rgb), 0.1);
  color: var(--danger);
}

.state-line.success {
  border: 1px solid rgba(var(--secondary-rgb), 0.28);
  background: rgba(var(--secondary-rgb), 0.1);
  color: var(--success);
}

@media (max-width: 720px) {
  .token-metrics {
    grid-template-columns: 1fr;
  }
}
</style>
