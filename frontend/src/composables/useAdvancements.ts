import { computed, type Ref } from 'vue'
import advancementData from '@/assets/advancements.json'

export function useAdvancements(advancements: Ref<any[] | null>) {
    const totalAdvancements = computed(() => advancements.value?.length || 0)
    const completedAdvancements = computed(() => advancements.value?.filter((a: any) => a.done).length || 0)

    const categorizedAdvancements = computed(() => {
        if (!advancements.value) return {}
        const result: Record<string, any[]> = {}
        for (const adv of advancements.value) {
            if (!adv.done) continue
            let category = 'Others'
            const parts = adv.key.split('/')
            if (parts.length > 1) {
                category = parts[0].replace('minecraft:', '').replace(/[_-]/g, ' ')
            }
            category = category.charAt(0).toUpperCase() + category.slice(1)
            if (!result[category]) result[category] = []
            result[category].push(adv)
        }
        return result
    })

    const getAdvancementMetadata = (key: string) => {
        return (advancementData as any)[key] || {
            name: key.split('/').pop(),
            icon: key.split('/').pop(),
            description: '',
            type: 'task'
        }
    }

    const getAdvIconPath = (advKey: string) => {
        const meta = getAdvancementMetadata(advKey)
        let category = advKey.split(':')[1]?.split('/')[0] || 'minecraft'
        if (category === 'story') category = 'minecraft'
        const iconName = meta.icon
        return `/mc_icons/advancements/${category}/${iconName}.png`
    }

    return {
        totalAdvancements,
        completedAdvancements,
        categorizedAdvancements,
        getAdvancementMetadata,
        getAdvIconPath
    }
}
