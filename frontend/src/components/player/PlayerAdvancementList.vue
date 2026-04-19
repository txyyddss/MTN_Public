<script setup lang="ts">
import PlayerCollapsiblePanel from '@/components/player/PlayerCollapsiblePanel.vue'
import { siteContent } from '@/content/siteContent'
import type { AdvancementMetadata, PlayerAdvancement } from '@/types/api'

defineProps<{
  advancements: PlayerAdvancement[] | null
  completedAdvancements: number
  totalAdvancements: number
  categorizedAdvancements: Record<string, PlayerAdvancement[]>
  getAdvancementMetadata: (key: string) => AdvancementMetadata
  getAdvIconPath: (key: string) => string
}>()

function handleImageError(event: Event): void {
  const image = event.target as HTMLImageElement | null
  if (image) {
    image.style.display = 'none'
  }
}
</script>

<template>
  <PlayerCollapsiblePanel
    v-if="advancements && advancements.length > 0"
    class="panel-card"
    :title="siteContent.playerDetail.sections.advancements"
  >
    <template #summary>
      <span class="meta-chip">{{ completedAdvancements }}/{{ totalAdvancements }}</span>
    </template>

    <div v-for="(items, category) in categorizedAdvancements" :key="category" class="advancement-section">
      <h4>{{ category }}</h4>
      <div class="advancement-grid">
        <article v-for="advancement in items" :key="advancement.key" class="advancement-card">
          <div :class="['icon-wrap', getAdvancementMetadata(advancement.key).type]">
            <img :src="getAdvIconPath(advancement.key)" :alt="getAdvancementMetadata(advancement.key).name" @error="handleImageError" />
          </div>
          <div class="advancement-copy">
            <strong>{{ getAdvancementMetadata(advancement.key).name }}</strong>
            <small>{{ getAdvancementMetadata(advancement.key).type }}</small>
          </div>
        </article>
      </div>
    </div>
  </PlayerCollapsiblePanel>
</template>

<style scoped>
.panel-card {
  display: grid;
  gap: 0.85rem;
}

.meta-chip {
  padding: 0.38rem 0.62rem;
  border-radius: 999px;
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.68rem;
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.advancement-section {
  display: grid;
  gap: 0.8rem;
}

.advancement-section h4 {
  font-size: 1.2rem;
}

.advancement-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 0.65rem;
}

.advancement-card {
  display: flex;
  gap: 0.7rem;
  padding: 0.76rem 0.86rem;
  border-radius: 18px;
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid rgba(255, 255, 255, 0.06);
  align-items: center;
}

.icon-wrap {
  width: 40px;
  height: 40px;
  display: grid;
  place-items: center;
  border-radius: 14px;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.08);
  flex-shrink: 0;
}

.icon-wrap.task {
  border-color: rgba(59, 130, 246, 0.35);
}

.icon-wrap.goal {
  border-color: rgba(91, 113, 246, 0.38);
}

.icon-wrap.challenge {
  border-color: rgba(125, 211, 252, 0.34);
}

.icon-wrap img {
  width: 24px;
  height: 24px;
  image-rendering: pixelated;
}

.advancement-copy {
  display: grid;
  gap: 0.2rem;
  min-width: 0;
}

.advancement-copy strong {
  color: var(--text-strong);
}

.advancement-copy small {
  color: var(--text-dim);
  text-transform: capitalize;
  font-size: 0.72rem;
}

@media (max-width: 980px) {
  .advancement-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (max-width: 640px) {
  .advancement-grid {
    grid-template-columns: 1fr;
  }
}
</style>
