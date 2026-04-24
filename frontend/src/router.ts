import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
    { path: '/', name: 'home', component: () => import('@/views/HomeView.vue') },
    { path: '/server-intro', redirect: '/' },
    { path: '/wiki', name: 'wiki', component: () => import('@/views/WikiRedirectView.vue') },
    { path: '/gallery', name: 'gallery', component: () => import('@/views/GalleryView.vue') },
    { path: '/players', name: 'players', component: () => import('@/views/PlayersView.vue') },
    { path: '/player/:uuid', name: 'playerDetail', component: () => import('@/views/PlayerDetailView.vue') }
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes
})

export default router
