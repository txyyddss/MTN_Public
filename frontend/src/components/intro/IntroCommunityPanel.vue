<script setup lang="ts">
import { computed } from 'vue'
import { storeToRefs } from 'pinia'

import { siteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'

const statusStore = useServerStatusStore()
const { status } = storeToRefs(statusStore)

const liveState = computed(() => (status.value?.java?.online ? 'Operational' : 'Offline'))
const updatedLabel = computed(() => {
  if (!status.value?.updated) {
    return 'Waiting for telemetry'
  }

  return new Date(status.value.updated).toLocaleString()
})
</script>

<template>
  <section id="community-panel" class="community-grid">
    <article class="glass-card community-card animate-entry delay-300">
      <span class="section-kicker">Community</span>
      <h2>{{ siteContent.intro.community.title }}</h2>
      <p>{{ siteContent.intro.community.body }}</p>
      <div class="community-foot">{{ siteContent.intro.community.caption }}</div>
    </article>

    <article class="glass-card status-card animate-entry delay-400">
      <div class="status-copy">
        <span class="section-kicker">Live node facts</span>
        <div class="status-list">
          <div class="status-row">
            <span>Location</span>
            <strong>Shenzhen, CN</strong>
          </div>
          <div class="status-row">
            <span>Network</span>
            <strong>China Mobile</strong>
          </div>
          <div class="status-row">
            <span>Hardware</span>
            <strong>Dedicated node</strong>
          </div>
          <div class="status-row">
            <span>Status</span>
            <strong>{{ liveState }}</strong>
          </div>
        </div>
        <small>Last status refresh: {{ updatedLabel }}</small>
      </div>
      <div class="qr-frame">
        <img src="/qrcode/qqgroup.jpg" alt="Official MT Network QQ group QR code" />
      </div>
    </article>
  </section>
</template>

<style scoped>
.community-grid {
  display: grid;
  grid-template-columns: 0.95fr 1.05fr;
  gap: 1.25rem;
  margin-top: 1.25rem;
}

.community-card,
.status-card {
  display: grid;
  gap: 1rem;
}

.community-card h2 {
  font-size: clamp(2rem, 4vw, 3.3rem);
  max-width: 12ch;
}

.community-card p,
.community-foot,
.status-copy small {
  color: var(--text-muted);
}

.status-card {
  grid-template-columns: 1fr 200px;
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
  background: rgba(255, 248, 234, 0.04);
  border: 1px solid rgba(255, 248, 234, 0.06);
}

.status-row span {
  color: var(--text-dim);
}

.qr-frame {
  padding: 0.7rem;
  border-radius: 22px;
  background: rgba(255, 248, 234, 0.05);
  border: 1px solid rgba(255, 248, 234, 0.08);
}

.qr-frame img {
  width: 100%;
  border-radius: 16px;
}

@media (max-width: 980px) {
  .community-grid,
  .status-card {
    grid-template-columns: 1fr;
  }
}
</style>
