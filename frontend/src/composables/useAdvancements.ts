import { computed, type Ref } from 'vue'

import advancementData from '@/assets/advancements.json'
import type { AdvancementMetadata, PlayerAdvancement } from '@/types/api'

const typedAdvancementData = advancementData as Record<string, AdvancementMetadata>

const fallbackMetadata = (key: string): AdvancementMetadata => ({
  name: key.split('/').pop() ?? key,
  icon: key.split('/').pop() ?? 'barrier',
  description: '',
  type: 'task'
})

export function useAdvancements(advancements: Ref<PlayerAdvancement[] | null>) {
  const totalAdvancements = computed(() => advancements.value?.length ?? 0)
  const completedAdvancements = computed(() => advancements.value?.filter((advancement) => advancement.done).length ?? 0)

  const categorizedAdvancements = computed((): Record<string, PlayerAdvancement[]> => {
    if (!advancements.value) {
      return {}
    }

    const result: Record<string, PlayerAdvancement[]> = {}

    for (const advancement of advancements.value) {
      if (!advancement.done) {
        continue
      }

      let category = 'Others'
      const parts = advancement.key.split('/')

      if (parts.length > 1) {
        category = parts[0].replace('minecraft:', '').replace(/[_-]/g, ' ')
      }

      const label = category.charAt(0).toUpperCase() + category.slice(1)
      if (!result[label]) {
        result[label] = []
      }

      result[label].push(advancement)
    }

    return result
  })

  const getAdvancementMetadata = (key: string): AdvancementMetadata => {
    const metadata = typedAdvancementData[key]
    if (!metadata) {
      return fallbackMetadata(key)
    }

    return {
      ...metadata,
      name: metadata.name.replace(/\u807d/g, ' ').replace(/\s+/g, ' ').trim()
    }
  }

  const getAdvIconPath = (advancementKey: string): string => {
    const metadata = getAdvancementMetadata(advancementKey)
    let category = advancementKey.split(':')[1]?.split('/')[0] ?? 'minecraft'

    if (category === 'story') {
      category = 'minecraft'
    }

    return `/mc_icons/advancements/${category}/${metadata.icon}.png`
  }

  return {
    totalAdvancements,
    completedAdvancements,
    categorizedAdvancements,
    getAdvancementMetadata,
    getAdvIconPath
  }
}
