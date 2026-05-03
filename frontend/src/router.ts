import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

import PlayersView from '@/views/PlayersView.vue'

const routes: RouteRecordRaw[] = [
    { path: '/', name: 'home', component: () => import('@/views/HomeView.vue'), meta: { seoKey: 'home', canonicalPath: '/' } },
    { path: '/server-intro', redirect: '/' },
    { path: '/wiki', name: 'wiki', component: () => import('@/views/WikiRedirectView.vue'), meta: { seoKey: 'wiki', canonicalPath: '/wiki', noindex: true } },
    { path: '/gallery', redirect: '/' },
    { path: '/players', name: 'players', component: PlayersView, meta: { seoKey: 'players', canonicalPath: '/players' } },
    { path: '/player/:uuid', name: 'playerDetail', component: () => import('@/views/PlayerDetailView.vue'), meta: { seoKey: 'playerDetail' } },
    { path: '/whitelist', name: 'whitelist', component: () => import('@/views/WhitelistView.vue'), meta: { seoKey: 'whitelist', canonicalPath: '/whitelist' } }
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
    async scrollBehavior(to, _from, savedPosition) {
        if (savedPosition) {
            return savedPosition
        }

        if (to.hash) {
            await new Promise((resolve) => window.setTimeout(resolve, 160))

            return {
                el: to.hash,
                top: 24,
                behavior: 'smooth'
            }
        }

        return { top: 0 }
    }
})

export default router
