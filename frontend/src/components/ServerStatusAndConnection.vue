<script setup lang="ts">
import { storeToRefs } from 'pinia'

import ServerConnectionPanel from '@/components/server-status/ServerConnectionPanel.vue'
import ServerTelemetryPanel from '@/components/server-status/ServerTelemetryPanel.vue'
import { useServerStatusStore } from '@/stores/serverStatus'

const statusStore = useServerStatusStore()
const { status, connection } = storeToRefs(statusStore)
</script>

<template>
  <section class="server-status-section">
    <div class="container layout-grid animate-entry delay-300">
      <ServerTelemetryPanel :status="status" />
      <ServerConnectionPanel :connection="connection" />
    </div>
  </section>
</template>

<style scoped>
.server-status-section {
  padding-bottom: 5rem;
}

.layout-grid {
  display: grid;
  grid-template-columns: 0.95fr 1.15fr;
  gap: 1.25rem;
}

@media (max-width: 980px) {
  .layout-grid {
    grid-template-columns: 1fr;
  }
}
</style>
