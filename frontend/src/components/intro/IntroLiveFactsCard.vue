<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  liveState: string
  updatedLabel: string
}

const props = defineProps<Props>()

const factRows = computed(() => [
  { label: 'Location', value: 'Shenzhen, CN' },
  { label: 'Network', value: 'China Mobile' },
  { label: 'Hardware', value: 'Dedicated node' },
  { label: 'Status', value: props.liveState }
])
</script>

<template>
  <article class="glass-card status-card animate-entry delay-400">
    <div class="status-copy">
      <span class="section-kicker">Live node facts</span>
      <div class="status-list">
        <div v-for="row in factRows" :key="row.label" class="status-row">
          <span class="status-label">{{ row.label }}</span>
          <strong class="status-value">{{ row.value }}</strong>
        </div>
      </div>
      <small class="status-updated">Last status refresh: {{ props.updatedLabel }}</small>
    </div>
    <div class="qr-frame">
      <img class="qr-image" src="/qrcode/qqgroup.jpg" alt="Official MTNetwork QQ group QR code" />
    </div>
  </article>
</template>

<style scoped>
.status-card {
  display: grid;
  grid-template-columns: 1fr 200px;
  gap: 1rem;
  align-items: center;
}

.status-list {
  display: grid;
  gap: 0.75rem;
  margin-bottom: 1rem;
}

.status-row {
  display: flex;
  justify-content: space-between;
  gap: 1rem;
  padding: 0.9rem 1rem;
  border-radius: 18px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.status-label {
  color: var(--text-dim);
}

.status-value {
  color: var(--text-strong);
  font-weight: 600;
}

.status-updated {
  color: var(--text-muted);
}

.qr-frame {
  padding: 0.7rem;
  border-radius: 22px;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.qr-image {
  width: 100%;
  border-radius: 16px;
}

@media (max-width: 980px) {
  .status-card {
    grid-template-columns: 1fr;
  }
}
</style>
