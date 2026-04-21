<script setup lang="ts">
import type { LinkedAccount, PlayerInfo } from '@/types/api'

defineProps<{
  info: PlayerInfo
  isOnline: boolean
  linkedAccount: LinkedAccount | null
}>()
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
  max-width: 56ch;
}
</style>
