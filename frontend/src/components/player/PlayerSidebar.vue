<script setup lang="ts">
import PlayerCollapsiblePanel from '@/components/player/PlayerCollapsiblePanel.vue'
import SkinViewer from '@/components/SkinViewer.vue'
import { siteContent } from '@/content/siteContent'
import type { LinkedAccount, PlayerInfo } from '@/types/api'
import { getSkinUrl } from '@/utils/minecraft'

defineProps<{
  info: PlayerInfo
  isOnline: boolean
  linkedAccount: LinkedAccount | null
  formatDate: (value: number) => string
  formatPlaytime: (value: number) => string
}>()
</script>

<template>
  <PlayerCollapsiblePanel class="profile-card" sticky title="Profile">
    <template #summary>
      <span class="meta-cluster">
        <span :class="['meta-chip', { live: isOnline }]">{{ isOnline ? 'Online' : info.type }}</span>
      </span>
    </template>

    <div class="avatar-header">
      <div class="skin-wrapper">
        <SkinViewer :skin-url="getSkinUrl(info.last_known_name, info.type)" :width="220" :height="260" />
      </div>
      <h2 :class="['profile-name', 'minecraft-font', { online: isOnline }]">
        {{ info.last_known_name }}
      </h2>
      <span class="type-tag">{{ info.type }}</span>
    </div>

    <div class="basic-info">
      <div class="info-row">
        <span class="label">{{ siteContent.playerDetail.profile.firstJoin }}</span>
        <span class="value">{{ formatDate(info.first_played) }}</span>
      </div>
      <div class="info-row">
        <span class="label">{{ siteContent.playerDetail.profile.lastSeen }}</span>
        <span class="value">{{ formatDate(info.last_seen) }}</span>
      </div>
      <div class="info-row">
        <span class="label">{{ siteContent.playerDetail.profile.playtime }}</span>
        <span class="value">{{ formatPlaytime(info.ticks_lived) }}</span>
      </div>
      <div class="info-row">
        <span class="label">{{ siteContent.playerDetail.profile.xpLevel }}</span>
        <span class="value">{{ info.xp_level }}</span>
      </div>
      <div v-if="linkedAccount && info.type === 'Bedrock'" class="info-row">
        <span class="label">{{ siteContent.playerDetail.profile.linkedTo }}</span>
        <span class="value linked">{{ linkedAccount.bedrock_username || linkedAccount.java_username }}</span>
      </div>
    </div>
  </PlayerCollapsiblePanel>
</template>

<style scoped>
.profile-card {
  display: grid;
  gap: 0.95rem;
  align-content: start;
}

.avatar-header {
  display: grid;
  justify-items: center;
  gap: 0.65rem;
}

.skin-wrapper {
  width: 100%;
  min-height: 250px;
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.08);
  overflow: hidden;
}

.profile-name {
  font-size: 1.48rem;
  text-align: center;
}

.profile-name.online {
  color: var(--success);
}

.type-tag {
  padding: 0.35rem 0.68rem;
  border-radius: 999px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.7rem;
  text-transform: uppercase;
}

.basic-info {
  display: grid;
  gap: 0.6rem;
}

.info-row {
  display: flex;
  justify-content: space-between;
  gap: 1rem;
  padding: 0.74rem 0.88rem;
  border-radius: 18px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.label {
  color: var(--text-dim);
}

.value {
  color: var(--text-strong);
  font-weight: 600;
  text-align: right;
}

.linked {
  color: var(--primary);
}

.meta-cluster {
  display: inline-flex;
}

.meta-chip {
  padding: 0.38rem 0.62rem;
  border-radius: 999px;
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.68rem;
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.meta-chip.live {
  color: var(--success);
}
</style>
