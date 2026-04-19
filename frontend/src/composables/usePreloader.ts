import { API_BASE_URL } from '@/config'
import { preloadImages, PreloadPriority, preloadScripts, preloadData } from '@/utils/preloader'
import { fetchWithCache } from '@/utils/dataCache'
import iconMap from '@/assets/icon_map.json'

export function usePreloader() {
    const preloadGlobalAssets = async () => {
        // 0. Preload critical JS chunks for other pages (highest priority)
        preloadScripts([
            '/src/views/PlayersView.vue',
            '/src/views/LeaderboardsView.vue',
            '/src/views/PlayerDetailView.vue'
        ])

        // 1. Preload server status (for online players list in /players)
        preloadData([`${API_BASE_URL}/api/status`], PreloadPriority.DATA)

        // 2. Preload players and skins (High Priority)
        const fetchPlayersAndSkins = async () => {
            try {
                const defaultUrl = `${API_BASE_URL}/api/players?`
                const allUrl = `${API_BASE_URL}/api/players?all=true`

                // Fetch default (recent) list first
                const json = await fetchWithCache(defaultUrl)
                // Also warm the all=true cache in background
                fetchWithCache(allUrl).catch(() => { })

                const players = json.players || []
                const avatarUrls: string[] = []
                const skinUrls: string[] = []

                players.forEach((p: any) => {
                    let cleanName = p.last_known_name || 'Steve'
                    if (p.type === 'Bedrock' || cleanName.startsWith('.')) {
                        cleanName = cleanName.replace(/^\./, '').replace(/^BE_/, '')
                    }
                    avatarUrls.push(`https://mineskin.eu/helm/${cleanName}`)
                    skinUrls.push(`https://mineskin.eu/skin/${cleanName}`)
                });

                preloadImages(avatarUrls, PreloadPriority.MEDIUM)
                preloadImages(skinUrls, PreloadPriority.HIGH)
            } catch (e) {
                console.error('Player preload failed:', e)
            }
        }

        fetchPlayersAndSkins()

        // 3. Preload all icons from the map with BACKGROUND priority
        const iconUrls = Object.values(iconMap)
        preloadImages(iconUrls, PreloadPriority.BACKGROUND)

        // 4. Preload leaderboard data
        const lbTypes = ['skills', 'playtime', 'mining', 'killing', 'deaths', 'walking', 'pvp']
        lbTypes.forEach(type => {
            preloadData([`${API_BASE_URL}/api/leaderboards/${type}`], PreloadPriority.BACKGROUND)
        })
    }

    const initPreloading = () => {
        window.addEventListener('load', () => {
            if ((window as any).requestIdleCallback) {
                (window as any).requestIdleCallback(() => {
                    setTimeout(preloadGlobalAssets, 500);
                });
            } else {
                setTimeout(preloadGlobalAssets, 2000);
            }
        });
    }

    return {
        initPreloading
    }
}
