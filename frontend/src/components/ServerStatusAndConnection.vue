<script setup lang="ts">
import { storeToRefs } from 'pinia'

import ServerTelemetryPanel from '@/components/server-status/ServerTelemetryPanel.vue'
import { useServerStatusStore } from '@/stores/serverStatus'

const statusStore = useServerStatusStore()
const { status, history } = storeToRefs(statusStore)
</script>

<template>
  <section class="server-status-section">
    <div class="container server-status-shell">
      <div class="server-status-intro animate-entry delay-300">
        <span class="section-kicker">Operations surface</span>
        <h2 class="server-status-title">Live infrastructure and connection routing.</h2>
        <p class="server-status-copy">
          Check the node state, recent player activity, and the right address for your client before joining.
        </p>
      </div>

      <div class="layout-grid animate-entry delay-400">
        <ServerTelemetryPanel :status="status" :history="history" />
      </div>
    </div>
  </section>
</template>

<style scoped>
.server-status-section {
  padding-bottom: 4.75rem;
}

.server-status-shell {
  display: grid;
  gap: 1.25rem;
}

.server-status-intro {
  display: grid;
  gap: 0.75rem;
  max-width: 760px;
}

.server-status-title {
  font-size: clamp(2.1rem, 4.6vw, 3.6rem);
  max-width: 13ch;
}

.server-status-copy {
  color: var(--text-muted);
  max-width: 56ch;
}

.layout-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1.1rem;
}

@media (max-width: 980px) {
  .layout-grid {
    grid-template-columns: 1fr;
  }
}
</style>
