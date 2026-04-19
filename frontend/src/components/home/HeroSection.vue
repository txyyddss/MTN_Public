<script setup lang="ts">
import { useServerStatus } from '@/composables/useServerStatus'

const { status } = useServerStatus()
</script>

<template>
  <header class="hero">
    <div class="hero-visual-bg"></div>
    
    <div class="container hero-content animate-entry delay-100">
      <div class="hero-badge-wrapper">
        <div class="hero-badge glass-card">
          <span class="pulse-icon" v-if="status?.online"></span>
          <span class="badge-text">{{ status?.players?.online || 0 }} Players Forge Legacies</span>
        </div>
      </div>
      
      <h1 class="hero-title">
        <span class="line-1">Redefining</span>
        <span class="text-gradient display-font">Pure Survival</span>
      </h1>
      
      <p class="hero-subtitle">
        MTNetwork delivers a masterclass in Minecraft survival. High-fidelity performance, 
        radical fairness, and an uncompromising cross-play architecture.
      </p>
      
      <div class="hero-actions">
        <router-link to="/players" class="btn-primary">
          Explore the Realm
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="m12 5 7 7-7 7"/></svg>
        </router-link>
        <router-link to="/server-intro" class="btn-secondary">Technical Blueprint</router-link>
      </div>
      
      <div class="quick-status glass-card animate-entry delay-300">
         <div class="status-item">
           <span class="status-label">Legacy Version</span>
           <span class="status-value">1.21.11</span>
         </div>
         <div class="status-divider"></div>
         <div class="status-item">
           <span class="status-label">Difficulty</span>
           <span class="status-value">Hardcore Instinct</span>
         </div>
         <div class="status-divider"></div>
         <div class="status-item">
           <span class="status-label">Access</span>
           <span class="status-value highlight">Cross-Protocol</span>
         </div>
      </div>
    </div>
    
    <!-- Immersive Background Elements -->
    <div class="hero-decor decor-1"></div>
    <div class="hero-decor decor-2"></div>
    <div class="hero-decor decor-3"></div>
  </header>
</template>

<style scoped>
.hero {
  position: relative;
  min-height: 95vh;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  overflow: hidden;
  padding: 8rem 0;
  background: radial-gradient(circle at 50% -20%, hsla(210, 100%, 55%, 0.05), transparent 70%);
}

.hero-visual-bg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: 
    radial-gradient(circle at 50% 10%, rgba(59, 130, 246, 0.08), transparent 60%),
    radial-gradient(circle at 10% 90%, rgba(91, 113, 246, 0.05), transparent 40%);
  z-index: 1;
}

.hero-content {
  position: relative;
  z-index: 10;
  max-width: 1000px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.hero-badge-wrapper {
  margin-bottom: 2.5rem;
}

.hero-badge {
  padding: 8px 18px !important;
  border-radius: 100px !important;
  font-size: 0.85rem;
  font-weight: 700;
  letter-spacing: 0.05em;
  text-transform: uppercase;
  color: var(--text-main);
  display: flex;
  align-items: center;
  gap: 12px;
  border-color: var(--glass-border-bright) !important;
  background: rgba(255, 255, 255, 0.02) !important;
}

.pulse-icon {
  width: 10px;
  height: 10px;
  background: var(--accent);
  border-radius: 50%;
  box-shadow: 0 0 15px var(--accent);
  animation: pulse 2.5s infinite cubic-bezier(0.4, 0, 0.6, 1);
}

@keyframes pulse {
  0%, 100% { transform: scale(1); opacity: 1; }
  50% { transform: scale(1.4); opacity: 0.4; }
}

.hero-title {
  font-size: clamp(3.5rem, 9vw, 7rem);
  font-weight: 800;
  line-height: 1.05;
  margin-bottom: 2.5rem;
  letter-spacing: -0.05em;
  font-family: var(--heading);
}

.line-1 {
  display: block;
  font-size: 0.45em;
  letter-spacing: 0.2em;
  text-transform: uppercase;
  color: var(--text-muted);
  font-family: var(--display);
  font-weight: 500;
  margin-bottom: 0.2em;
}

.display-font {
  font-family: var(--display);
  font-weight: 800;
  font-style: italic;
}

.text-gradient {
  background: linear-gradient(135deg, #fff 20%, var(--primary) 60%, var(--secondary) 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  filter: drop-shadow(0 0 30px var(--primary-glow));
}

.hero-subtitle {
  font-size: 1.25rem;
  color: var(--text-muted);
  max-width: 720px;
  margin-bottom: 4rem;
  font-weight: 400;
  line-height: 1.7;
}

.hero-actions {
  display: flex;
  gap: 2rem;
  margin-bottom: 5rem;
}

.btn-secondary {
  padding: 16px 36px;
  background: rgba(255, 255, 255, 0.02);
  color: var(--text-main);
  font-family: var(--heading);
  font-weight: 700;
  border: 1px solid var(--glass-border-bright);
  border-radius: var(--radius-sm);
  transition: all 0.4s var(--transition-fast);
  backdrop-filter: blur(10px);
}

.btn-secondary:hover {
  background: rgba(255, 255, 255, 0.06);
  border-color: #fff;
  transform: translateY(-2px);
}

.quick-status {
  display: flex;
  padding: 20px 48px !important;
  gap: 40px;
  border-radius: 100px !important;
  background: rgba(5, 5, 5, 0.8) !important;
}

.status-item {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.status-label {
  font-size: 0.7rem;
  text-transform: uppercase;
  color: var(--text-dim);
  letter-spacing: 2px;
  font-weight: 700;
  margin-bottom: 4px;
}

.status-value {
  font-size: 1.1rem;
  font-weight: 600;
  color: #fff;
  font-family: var(--heading);
}

.status-value.highlight {
  color: var(--accent);
}

.status-divider {
  width: 1px;
  height: 36px;
  background: var(--glass-border-bright);
  align-self: center;
}

.hero-decor {
  position: absolute;
  filter: blur(140px);
  opacity: 0.15;
  z-index: 1;
  pointer-events: none;
}

.decor-1 { width: 600px; height: 600px; background: var(--primary); top: -200px; right: -100px; }
.decor-2 { width: 500px; height: 500px; background: var(--secondary); bottom: -150px; left: -150px; }
.decor-3 { width: 400px; height: 400px; background: var(--accent); top: 30%; left: 10%; }

@media (max-width: 768px) {
  .hero-title { font-size: 3.5rem; }
  .quick-status { flex-direction: column; gap: 1.5rem; border-radius: 24px !important; padding: 32px !important; }
  .status-divider { display: none; }
  .hero-actions { flex-direction: column; width: 100%; gap: 1rem; }
}
</style>
