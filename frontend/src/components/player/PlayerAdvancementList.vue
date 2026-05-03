<script setup lang="ts">
import { computed } from 'vue'

import PlayerCollapsiblePanel from '@/components/player/PlayerCollapsiblePanel.vue'
import { getAdvancementTypeLabel, useSiteContent } from '@/content/siteContent'
import type { AdvancementMetadata, PlayerAdvancement } from '@/types/api'

const props = defineProps<{
  advancements: PlayerAdvancement[] | null
  completedAdvancements: number
  totalAdvancements: number
  categorizedAdvancements: Record<string, PlayerAdvancement[]>
  getAdvancementMetadata: (key: string) => AdvancementMetadata
  getAdvIconPath: (key: string) => string
}>()

const siteContent = useSiteContent()
const hasAdvancements = computed(() => (props.advancements?.length ?? 0) > 0)
const advancementSections = computed(() =>
  Object.entries(props.categorizedAdvancements).map(([category, items]) => ({
    category,
    items: items.map((advancement) => {
      const metadata = props.getAdvancementMetadata(advancement.key)

      return {
        key: advancement.key,
        metadata,
        iconPath: props.getAdvIconPath(advancement.key),
        typeLabel: getAdvancementTypeLabel(metadata.type)
      }
    })
  }))
)

function handleImageError(event: Event): void {
  const image = event.target as HTMLImageElement | null
  if (image) {
    image.style.display = 'none'
  }
}
</script>

<template>
  <PlayerCollapsiblePanel
    v-if="hasAdvancements"
    class="panel-card"
    :title="siteContent.playerDetail.sections.advancements"
  >
    <div v-for="section in advancementSections" :key="section.category" class="advancement-section">
      <h4>{{ section.category }}</h4>
      <div class="advancement-grid">
        <article v-for="advancement in section.items" :key="advancement.key" class="advancement-card">
          <div :class="['icon-wrap', advancement.metadata.type]">
            <img :src="advancement.iconPath" :alt="advancement.metadata.name" @error="handleImageError" />
          </div>
          <div class="advancement-copy">
            <strong>{{ advancement.metadata.name }}</strong>
            <small>{{ advancement.typeLabel }}</small>
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
  background: var(--player-glass-tile-bg);
  border: 1px solid var(--player-glass-border-soft);
  box-shadow: var(--glass-inset);
  align-items: center;
  transition:
    transform var(--transition-fast),
    border-color var(--transition-fast),
    background var(--transition-fast);
}

.advancement-card:hover {
  transform: translateY(-1px);
  border-color: var(--player-glass-border-strong);
  background: var(--player-glass-tile-bg-hover);
}

.icon-wrap {
  width: 40px;
  height: 40px;
  display: grid;
  place-items: center;
  border-radius: 14px;
  background: rgba(var(--secondary-rgb), 0.08);
  border: 1px solid var(--player-glass-border-soft);
  flex-shrink: 0;
}

.icon-wrap.task {
  border-color: rgba(var(--secondary-rgb), 0.35);
}

.icon-wrap.goal {
  border-color: rgba(var(--primary-rgb), 0.38);
}

.icon-wrap.challenge {
  border-color: rgba(147, 197, 253, 0.34);
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
