<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { RouterLink, RouterView } from 'vue-router'

const menuOpen = ref(false)
const toggleMenu = () => menuOpen.value = !menuOpen.value

const isLoading = ref(true)

onMounted(() => {
  // Simulate app initialization / asset loading
  setTimeout(() => {
    isLoading.value = false
  }, 1500)
})
</script>

<template>
  <div class="app-shell">
    <!-- Loading Overlay -->
    <div :class="['loader-overlay', { hidden: !isLoading }]">
      <div class="mtn-blocks">
        <div class="mtn-block"></div>
        <div class="mtn-block"></div>
        <div class="mtn-block"></div>
      </div>
      <h2 style="position: absolute; margin-top: 120px; font-family: 'Outfit', sans-serif; letter-spacing: 4px;">
        <span class="logo-accent">MTN</span><span class="logo-white">etwork</span>
      </h2>
    </div>

    <!-- Main Navigation -->
    <nav class="desktop-nav glass-nav">
      <div class="container nav-content">
        <div class="logo">
          <RouterLink to="/">
            <span class="logo-accent">MTN</span><span class="logo-white">etwork</span>
          </RouterLink>
        </div>

        <!-- Hamburger -->
        <button class="hamburger" @click="toggleMenu" aria-label="Toggle Menu">
          <span class="bar" :class="{ 'open': menuOpen }"></span>
          <span class="bar" :class="{ 'open': menuOpen }"></span>
          <span class="bar" :class="{ 'open': menuOpen }"></span>
        </button>

        <!-- Links -->
        <div class="nav-links" :class="{ 'open': menuOpen }">
          <RouterLink to="/" @click="menuOpen = false">Home</RouterLink>
          <RouterLink to="/players" @click="menuOpen = false">Players</RouterLink>
          <RouterLink to="/leaderboards" @click="menuOpen = false">Leaderboards</RouterLink>
          <RouterLink to="/server-intro" @click="menuOpen = false">Intro</RouterLink>
          <RouterLink to="/core-members" @click="menuOpen = false">Team</RouterLink>
          <RouterLink to="/gallery" @click="menuOpen = false">Gallery</RouterLink>
          <a href="/wiki" class="wiki-link">Wiki</a>
        </div>
      </div>
    </nav>

    <!-- Main Content Application -->
    <main class="main-content">
      <RouterView />
    </main>

    <!-- Footer -->
    <footer class="app-footer">
      <div class="container">
        <p>&copy; 2026 MTNetwork. All Rights Reserved. A non-profit community server.</p>
      </div>
    </footer>
  </div>
</template>

<style scoped>
.app-shell {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.glass-nav {
  position: sticky;
  top: 0;
  z-index: 50;
  background: var(--glass-bg);
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  border-bottom: 1px solid var(--glass-border);
}

.nav-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 0.75rem;
  padding-bottom: 0.75rem;
}

.logo a {
  font-family: var(--heading);
  font-size: 1.5rem;
  font-weight: 800;
  letter-spacing: 2px;
  text-decoration: none;
  display: flex;
  align-items: center;
}

.logo-accent {
  color: var(--primary);
}

.logo-white {
  color: #fff;
}

.nav-links {
  display: flex;
  gap: 1.5rem;
  align-items: center;
}

.nav-links a {
  color: var(--text-muted);
  text-decoration: none;
  font-weight: 600;
  font-family: var(--heading);
  letter-spacing: 0.5px;
  font-size: 0.95rem;
  position: relative;
}

.nav-links a:not(.wiki-link)::after {
  content: '';
  position: absolute;
  bottom: -4px;
  left: 0;
  width: 0;
  height: 2px;
  background: var(--primary);
  transition: width 0.3s ease;
  border-radius: 2px;
}

.nav-links a:hover, .nav-links a.router-link-active {
  color: #fff;
}

.nav-links a:not(.wiki-link):hover::after, .nav-links a.router-link-active::after {
  width: 100%;
}

.wiki-link {
  padding: 8px 16px;
  border: 1px solid var(--glass-border);
  border-radius: var(--radius-sm);
  background: rgba(255,255,255,0.05);
  transition: all 0.3s;
}
.wiki-link:hover {
  background: rgba(255,255,255,0.1);
  border-color: rgba(255,255,255,0.3);
}

/* Hamburger */
.hamburger {
  display: none;
  background: none;
  border: none;
  cursor: pointer;
  width: 30px;
  height: 20px;
  position: relative;
  z-index: 60;
}
.hamburger .bar {
  position: absolute;
  height: 2px;
  width: 100%;
  background: #fff;
  transition: all 0.3s ease;
  left: 0;
}
.hamburger .bar:nth-child(1) { top: 0; }
.hamburger .bar:nth-child(2) { top: 9px; }
.hamburger .bar:nth-child(3) { top: 18px; }

.hamburger .bar.open:nth-child(1) {
  top: 9px;
  transform: rotate(45deg);
}
.hamburger .bar.open:nth-child(2) {
  opacity: 0;
}
.hamburger .bar.open:nth-child(3) {
  top: 9px;
  transform: rotate(-45deg);
}

@media (max-width: 860px) {
  .hamburger {
    display: block;
  }
  .nav-links {
    position: fixed;
    top: 0;
    right: -100%;
    width: 250px;
    height: 100vh;
    background: var(--bg-lighter);
    flex-direction: column;
    padding-top: 80px;
    transition: right 0.4s cubic-bezier(0.68, -0.55, 0.265, 1.55);
    border-left: 1px solid var(--glass-border);
    box-shadow: -10px 0 30px rgba(0,0,0,0.5);
  }
  .nav-links.open {
    right: 0;
  }
}

.main-content {
  flex: 1;
}

.app-footer {
  text-align: center;
  padding: 2rem 0;
  border-top: 1px solid var(--glass-border);
  color: var(--text-muted);
  font-size: 0.9rem;
  margin-top: auto;
}
</style>
