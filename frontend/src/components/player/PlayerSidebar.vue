<script setup lang="ts">
import SkinViewer from '@/components/SkinViewer.vue'
import { getSkinUrl } from '@/utils/minecraft'

defineProps<{
  info: any
  isOnline: boolean
  linkedAccount: any
  formatDate: (ms: number) => string
  formatPlaytime: (ticks: number) => string
}>()
</script>

<template>
  <aside class="profile-card glass-card">
    <div class="avatar-header">
      <div class="skin-wrapper">
        <SkinViewer :skin-url="getSkinUrl(info.last_known_name, info.type)" />
      </div>
      <h2 :class="['profile-name', 'minecraft-font', { online: isOnline }]">
        {{ info.last_known_name }}
      </h2>
      <span :class="['type-tag', info.type?.toLowerCase()]">{{ info.type === 'Bedrock' ? 'Bedrock' : 'Java' }}</span>
    </div>

    <div class="basic-info">
      <div class="info-row">
        <span class="label">First Join</span>
        <span class="val">{{ formatDate(info.first_played) }}</span>
      </div>
      <div class="info-row">
        <span class="label">Last Seen</span>
        <span class="val">{{ formatDate(info.last_seen) }}</span>
      </div>
      <div class="info-row">
        <span class="label">Playtime</span>
        <span class="val">{{ formatPlaytime(info.ticks_lived) }}</span>
      </div>
      <div class="info-row">
        <span class="label">XP Level</span>
        <span class="val badge lvl-badge">{{ info.xp_level }}</span>
      </div>
      <div class="info-row linked-row" v-if="linkedAccount && info.type === 'Bedrock'">
        <span class="label">Linked to</span>
        <span class="val account-link">
          {{ linkedAccount.bedrock_username || linkedAccount.java_username }}
        </span>
      </div>
    </div>
  </aside>
</template>

<style scoped>
.profile-card {
  height: fit-content;
  position: sticky;
  top: 100px;
}

.avatar-header {
  text-align: center;
  margin-bottom: 2rem;
}

.skin-wrapper {
  width: 100%;
  aspect-ratio: 1;
  background: rgba(255, 255, 255, 0.03);
  border-radius: var(--radius-lg);
  margin-bottom: 1.5rem;
  overflow: hidden;
  border: 1px solid var(--glass-border);
}

.profile-name {
  font-size: 1.8rem;
  margin-bottom: 0.5rem;
  transition: color 0.3s;
}

.profile-name.online {
  color: #10B981;
  text-shadow: 0 0 15px rgba(16, 185, 129, 0.4);
}

.type-tag {
  display: inline-block;
  padding: 4px 12px;
  border-radius: 100px;
  font-size: 0.75rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 1px;
}

.type-tag.java {
  background: rgba(59, 130, 246, 0.1);
  color: var(--primary);
  border: 1px solid var(--primary);
}

.type-tag.bedrock {
  background: rgba(16, 185, 129, 0.1);
  color: #10B981;
  border: 1px solid #10B981;
}

.basic-info {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.info-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px;
  background: rgba(255, 255, 255, 0.03);
  border-radius: var(--radius-md);
  border: 1px solid rgba(255, 255, 255, 0.05);
}

.label {
  color: var(--text-muted);
  font-size: 0.85rem;
  font-weight: 500;
}

.val {
  font-weight: 700;
  color: #fff;
}

.badge {
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 0.8rem;
}

.lvl-badge {
  background: var(--primary);
  color: #000;
}

.account-link {
  color: var(--primary);
}
</style>
