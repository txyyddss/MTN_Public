<script setup lang="ts">
import { storeToRefs } from 'pinia'

import ConnectionPanel from '@/components/server-status/ConnectionPanel.vue'
import ServerTelemetryPanel from '@/components/server-status/ServerTelemetryPanel.vue'
import { useServerStatusStore } from '@/stores/serverStatus'

const statusStore = useServerStatusStore()
const { status, connection } = storeToRefs(statusStore)
</script>

<template>
  <section class="server-status-section">
    <div class="container server-status-shell">
      <div class="layout-grid animate-entry delay-400">
        <ServerTelemetryPanel :status="status" />
        <ConnectionPanel :connection="connection" />
      </div>
    </div>
  </section>
</template>

<style scoped>
.server-status-section {
  position: relative;
  padding-bottom: 4.75rem;
}

.server-status-shell {
  display: grid;
  gap: 1.4rem;
}

.layout-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 1.2rem;
}

@media (max-width: 980px) {
  .layout-grid {
    grid-template-columns: 1fr;
  }
}
</style>
