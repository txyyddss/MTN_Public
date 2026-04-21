import { API_BASE_URL } from '@/config'
import { preloadImages, PreloadPriority, preloadScripts } from '@/utils/preloader'
import { fetchWithCache } from '@/utils/dataCache'
import { shallowRef } from 'vue'
import type { PlayerListResponse } from '@/types/api'

export function usePreloader() {
    const isReady = shallowRef(false)
    let initPromise: Promise<void> | null = null

    const preloadGlobalAssets = async () => {
        if (import.meta.env.DEV) {
            preloadScripts([
                '/src/views/PlayersView.vue',
                '/src/views/PlayerDetailView.vue'
            ])
        }

        const fetchPlayersAndSkins = async () => {
            try {
                const defaultUrl = `${API_BASE_URL}/api/players?`

                const json = await fetchWithCache<PlayerListResponse>(defaultUrl)

                const players = json.players || []
                const avatarUrls: string[] = []

                players.forEach((p) => {
                    let cleanName = p.last_known_name || 'Steve'
                    if (p.type === 'Bedrock' || cleanName.startsWith('.')) {
                        cleanName = cleanName.replace(/^\./, '').replace(/^BE_/, '')
                    }
                    avatarUrls.push(`https://mineskin.eu/helm/${cleanName}`)
                })

                preloadImages(avatarUrls, PreloadPriority.MEDIUM)
            } catch (e) {
                console.error('Player preload failed:', e)
            }
        }

        await Promise.allSettled([fetchPlayersAndSkins()])
    }

    const waitForLoad = () =>
        new Promise<void>((resolve) => {
            if (document.readyState === 'complete') {
                resolve()
                return
            }

            const onLoad = () => {
                window.removeEventListener('load', onLoad)
                resolve()
            }

            window.addEventListener('load', onLoad, { once: true })
        })

    const initPreloading = () => {
        if (initPromise) {
            return initPromise
        }

        initPromise = waitForLoad()
            .then(() => preloadGlobalAssets())
            .finally(() => {
                isReady.value = true
            })

        return initPromise
    }

    return {
        isReady,
        initPreloading
    }
}
