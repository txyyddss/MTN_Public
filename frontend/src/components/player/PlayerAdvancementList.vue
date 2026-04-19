<script setup lang="ts">
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
  <section v-if="advancements && advancements.length > 0" class="glass-card panel-card">
    <div class="panel-head">
      <h3>{{ siteContent.playerDetail.sections.advancements }}</h3>
      <span class="meta-chip">{{ completedAdvancements }}/{{ totalAdvancements }}</span>
    </div>

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
  </section>
</template>

<style scoped>
.panel-card {
  display: grid;
  gap: 1rem;
}

.panel-head {
  display: flex;
  justify-content: space-between;
  gap: 1rem;
  align-items: center;
}

.panel-head h3 {
  font-size: 1.8rem;
}

.meta-chip {
  padding: 0.45rem 0.7rem;
  border-radius: 999px;
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.74rem;
  border: 1px solid rgba(255, 248, 234, 0.08);
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
  gap: 0.75rem;
}

.advancement-card {
  display: flex;
  gap: 0.85rem;
  padding: 0.9rem 1rem;
  border-radius: 18px;
  background: rgba(255, 248, 234, 0.04);
  border: 1px solid rgba(255, 248, 234, 0.06);
  align-items: center;
}

.icon-wrap {
  width: 44px;
  height: 44px;
  display: grid;
  place-items: center;
  border-radius: 14px;
  background: rgba(255, 248, 234, 0.05);
  border: 1px solid rgba(255, 248, 234, 0.08);
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
  width: 28px;
  height: 28px;
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
