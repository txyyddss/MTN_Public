<script setup lang="ts">
import { ref } from 'vue'
import { RouterLink, RouterView } from 'vue-router'

const menuOpen = ref(false)
const toggleMenu = () => menuOpen.value = !menuOpen.value
</script>

<template>
  <div class="app-shell">
    <nav class="desktop-nav">
      <div class="logo">
        <RouterLink to="/">MTN</RouterLink>
      </div>

      <!-- Hamburger Menu for Mobile -->
      <button class="hamburger" @click="toggleMenu" aria-label="Toggle Navigation">
        ☰
      </button>

      <!-- Nav Links -->
      <div class="nav-links" :class="{ 'open': menuOpen }">
        <RouterLink to="/" @click="menuOpen = false">Home</RouterLink>
        <RouterLink to="/players" @click="menuOpen = false">Players</RouterLink>
        <RouterLink to="/leaderboards" @click="menuOpen = false">Leaderboards</RouterLink>
        <RouterLink to="/server-intro" @click="menuOpen = false">Intro</RouterLink>
        <RouterLink to="/core-members" @click="menuOpen = false">Core Members</RouterLink>
        <RouterLink to="/gallery" @click="menuOpen = false">Gallery</RouterLink>
        <a href="/wiki">Wiki</a>
      </div>
    </nav>
    <main class="main-content">
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
.app-shell {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: #121212;
  color: #fff;
  font-family: 'Inter', system-ui, sans-serif;
}
.desktop-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 2rem;
  background-color: #1a1a1a;
  border-bottom: 2px solid #333;
  position: sticky;
  top: 0;
  z-index: 50;
}
.logo a {
  font-size: 1.5rem;
  font-weight: 800;
  letter-spacing: 2px;
  color: #00ff88;
  text-decoration: none;
}
.nav-links {
  display: flex;
  gap: 1.5rem;
  align-items: center;
}
.nav-links a {
  color: #ccc;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.2s ease;
}
.nav-links a:hover, .nav-links a.router-link-active {
  color: #00ff88;
}
.hamburger {
  display: none;
  background: none;
  border: none;
  color: #fff;
  font-size: 1.5rem;
  cursor: pointer;
}

/* Tablet & Mobile Nav */
@media (max-width: 768px) {
  .hamburger {
    display: block;
  }
  .nav-links {
    display: none;
    flex-direction: column;
    position: absolute;
    top: 100%;
    left: 0;
    width: 100%;
    background-color: #1a1a1a;
    padding: 1rem 0;
    border-bottom: 2px solid #333;
  }
  .nav-links.open {
    display: flex;
  }
  .nav-links a {
    width: 100%;
    text-align: center;
    padding: 0.5rem 0;
  }
}

.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
}
</style>
