<script setup lang="ts">
import PlayerCollapsiblePanel from '@/components/player/PlayerCollapsiblePanel.vue'
import SkinViewer from '@/components/SkinViewer.vue'
import { useSiteContent } from '@/content/siteContent'
import type { LinkedAccount, PlayerInfo } from '@/types/api'
import { getSkinUrl } from '@/utils/minecraft'

defineProps<{
  info: PlayerInfo
  isOnline: boolean
  linkedAccount: LinkedAccount | null
  formatPlaytime: (value: number) => string
}>()

const siteContent = useSiteContent()
</script>

<template>
  <PlayerCollapsiblePanel class="profile-card" sticky :title="siteContent.playerDetail.profileCardTitle">
    <div class="avatar-header">
      <div class="skin-wrapper">
        <SkinViewer :skin-url="getSkinUrl(info.last_known_name, info.type)" :width="220" :height="260" />
      </div>
      <h2 :class="['profile-name', 'minecraft-font', { online: isOnline }]">
        {{ info.last_known_name }}
      </h2>
    </div>

    <div class="basic-info">
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

</style>
