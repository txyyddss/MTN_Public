<script setup lang="ts">
import { useSiteContent } from '@/content/siteContent'
import type { LinkedAccount, PlayerInfo } from '@/types/api'

defineProps<{
  info: PlayerInfo
  isOnline: boolean
  linkedAccount: LinkedAccount | null
}>()

const siteContent = useSiteContent()
</script>

<template>
  <section class="glass-card detail-header animate-entry">
    <span class="detail-header-mark" aria-hidden="true">MTN</span>
    <div class="detail-header-copy">
      <span class="page-kicker">{{ siteContent.playerDetail.header.kicker }}</span>
      <h1 :class="['detail-player-name', 'minecraft-font', { online: isOnline }]">{{ info.last_known_name }}</h1>
      <p class="detail-player-copy">
        {{ isOnline ? siteContent.playerDetail.header.liveBody : siteContent.playerDetail.header.archiveBody }}
      </p>
    </div>
  </section>
</template>

<style scoped>
.detail-header {
  position: relative;
  isolation: isolate;
  display: grid;
  gap: 0.9rem;
  border-color: rgba(76, 147, 251, 0.22);
}

.detail-header-copy {
  position: relative;
  z-index: 1;
  display: grid;
  gap: 0.55rem;
}

.detail-header-mark {
  position: absolute;
  right: 1rem;
  top: 1rem;
  color: rgba(141, 184, 255, 0.08);
  font-family: var(--display);
  font-size: clamp(3rem, 10vw, 6rem);
  font-weight: 900;
  letter-spacing: 0.24em;
  line-height: 0.75;
  pointer-events: none;
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
  max-width: 56ch;
}
</style>
