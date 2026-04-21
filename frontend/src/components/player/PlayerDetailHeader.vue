<script setup lang="ts">
import { computed } from 'vue'

import { siteContent } from '@/content/siteContent'
import type { LinkedAccount, PlayerInfo } from '@/types/api'

const props = defineProps<{
  info: PlayerInfo
  isOnline: boolean
  linkedAccount: LinkedAccount | null
}>()

const linkedAccountName = computed(() => {
  if (!props.linkedAccount) {
    return ''
  }

  if (props.info.type === 'Bedrock') {
    return props.linkedAccount.java_username
  }

  return props.linkedAccount.bedrock_username || props.linkedAccount.java_username
})
</script>

<template>
  <section class="glass-card detail-header animate-entry">
    <div class="detail-header-copy">
      <span class="page-kicker">Public player dossier</span>
      <h1 :class="['detail-player-name', 'minecraft-font', { online: isOnline }]">{{ info.last_known_name }}</h1>
      <p class="detail-player-copy">
        {{ isOnline ? 'Live presence detected in the server status feed.' : 'Historical record and progression summary.' }}
      </p>
    </div>

    <div class="detail-header-badges">
      <span :class="['badge-pill', 'detail-status-pill', { online: isOnline }]">
        <strong>{{ isOnline ? siteContent.playerDetail.summary.onlineNow : siteContent.playerDetail.summary.archiveRecord }}</strong>
      </span>
      <span class="badge-pill"><strong>{{ info.type }}</strong></span>
      <span v-if="linkedAccountName" class="badge-pill">
        {{ siteContent.playerDetail.summary.linkedTo }}
        <strong>{{ linkedAccountName }}</strong>
      </span>
    </div>
  </section>
</template>

<style scoped>
.detail-header {
  display: grid;
  gap: 0.9rem;
}

.detail-header-copy {
  display: grid;
  gap: 0.55rem;
}

.detail-player-name {
  font-size: clamp(2rem, 5vw, 3.6rem);
}

.detail-player-name.online,
.detail-status-pill.online strong {
  color: var(--success);
}

.detail-player-copy {
  color: var(--text-muted);
  max-width: 48ch;
}

.detail-header-badges {
  display: flex;
  flex-wrap: wrap;
  gap: 0.55rem;
}
</style>
