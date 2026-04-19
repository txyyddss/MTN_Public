<script setup lang="ts">
defineProps<{
  advancements: any[] | null
  completedAdvancements: number
  totalAdvancements: number
  categorizedAdvancements: Record<string, any[]>
  getAdvancementMetadata: (key: string) => any
  getAdvIconPath: (key: string) => string
}>()
</script>

<template>
  <section class="panel glass-card" v-if="advancements && advancements.length > 0">
    <h3><img src="/icons/all_advancements.ico" class="header-icon" /> Advancements <small class="text-muted">({{ completedAdvancements }}/{{ totalAdvancements }})</small></h3>
    
    <div class="adv-category" v-for="(items, category) in categorizedAdvancements" :key="category">
      <h4 class="category-name">{{ category }}</h4>
      <div class="advancements-grid">
        <div class="adv-card" v-for="adv in items" :key="adv.key">
          <div class="adv-icon-wrap" :class="getAdvancementMetadata(adv.key).type">
            <img 
               :src="getAdvIconPath(adv.key)" 
               :alt="getAdvancementMetadata(adv.key).name" 
               class="adv-icon" 
               @error="($event: Event) => ($event.target as HTMLImageElement).style.display='none'" 
            />
          </div>
          <div class="adv-info">
            <span class="adv-name">{{ getAdvancementMetadata(adv.key).name }}</span>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.panel h3 {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 2rem;
}

.header-icon {
  width: 24px;
  height: 24px;
}

.adv-category {
  margin-bottom: 2.5rem;
}

.category-name {
  font-size: 1rem;
  letter-spacing: 2px;
  text-transform: uppercase;
  color: var(--text-muted);
  margin-bottom: 1.5rem;
  padding-bottom: 0.5rem;
  border-bottom: 1px solid var(--glass-border);
}

.advancements-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(130px, 1fr));
  gap: 1.2rem;
}

.adv-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  gap: 10px;
}

.adv-icon-wrap {
  width: 54px;
  height: 54px;
  background: rgba(255, 255, 255, 0.03);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid var(--glass-border);
  transition: transform 0.3s cubic-bezier(0.16, 1, 0.3, 1);
}

.adv-card:hover .adv-icon-wrap {
  transform: scale(1.1) rotate(5deg);
  border-color: var(--primary);
}

.adv-icon {
  width: 32px;
  height: 32px;
  image-rendering: pixelated;
}

.adv-name {
  font-size: 0.8rem;
  font-weight: 500;
  color: var(--text-main);
}

/* Advancement types colors */
.goal { border-color: #fbbf24; }
.challenge { border-color: #8b5cf6; }
</style>
