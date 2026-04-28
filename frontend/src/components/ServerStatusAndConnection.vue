<script setup lang="ts">
import { storeToRefs } from 'pinia'

import ConnectionPanel from '@/components/server-status/ConnectionPanel.vue'
import ServerTelemetryPanel from '@/components/server-status/ServerTelemetryPanel.vue'
import { useSiteContent } from '@/content/siteContent'
import { useServerStatusStore } from '@/stores/serverStatus'

const statusStore = useServerStatusStore()
const { status, connection, history } = storeToRefs(statusStore)
const siteContent = useSiteContent()
</script>

<template>
  <section class="server-status-section">
    <div class="container server-status-shell">
      <div class="server-status-intro animate-entry delay-300">
        <span class="section-kicker">{{ siteContent.home.serverIntro.kicker }}</span>
        <h2 class="server-status-title">{{ siteContent.home.serverIntro.title }}</h2>
        <p class="server-status-copy">{{ siteContent.home.serverIntro.body }}</p>
      </div>

      <div class="layout-grid animate-entry delay-400">
        <ServerTelemetryPanel :status="status" :history="history" />
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
  grid-template-columns: repeat(2, 1fr);
  gap: 1.2rem;
}

@media (max-width: 980px) {
  .layout-grid {
    grid-template-columns: 1fr;
  }
}
</style>
