import { createRouter, createWebHistory } from 'vue-router'

const routes = [
    { path: '/', name: 'home', component: () => import('@/views/HomeView.vue') },
    { path: '/server-intro', name: 'intro', component: () => import('@/views/IntroView.vue') },
    { path: '/core-members', name: 'members', component: () => import('@/views/StubView.vue') },
    { path: '/gallery', name: 'gallery', component: () => import('@/views/GalleryView.vue') },
    { path: '/players', name: 'players', component: () => import('@/views/PlayersView.vue') },
    { path: '/player/:uuid', name: 'playerDetail', component: () => import('@/views/PlayerDetailView.vue') },
    { path: '/leaderboards', name: 'leaderboards', component: () => import('@/views/LeaderboardsView.vue') }
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes
})

export default router
