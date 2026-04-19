<script setup lang="ts">
import { ref } from 'vue'
import { useServerStatus } from '@/composables/useServerStatus'
import ServerStatusAndConnection from '@/components/ServerStatusAndConnection.vue'

const { status } = useServerStatus()

const features = ref([
  {
    title: 'Efficiency by Design',
    description: 'We leverage optimized infrastructure and advanced resource management to deliver a buttery-smooth survival experience without the premium cost.',
    icon: '⚡',
    class: 'feature-large'
  },
  {
    title: 'Pure Vanilla',
    description: 'A genuine Minecraft experience. We focus on fair-play with robust anticheat and only the most essential quality-of-life enhancements.',
    icon: '🌲'
  },
  {
    title: 'Absolute Fairness',
    description: 'Strict non-P2W policy. Every player starts on equal footing, with no paid advantages, rank-based perks, or administrative bias.',
    icon: '⚖️'
  },
  {
    title: 'Full Cross-Play',
    description: 'Seamlessly connect from Java or Bedrock Editions. Our integrated account system ensures your progress stays with you on any device.',
    icon: '🎮',
    class: 'feature-medium'
  },
  {
    title: 'Built for Longevity',
    description: 'Our sustainable operational model ensures the server remains online and stable for years to come. Your builds are safe here.',
    icon: '🛡️'
  },
  {
    title: 'Community First',
    description: 'MTNetwork is a non-profit, passion project dedicated to fostering a diverse and welcoming community for all Minecraft enthusiasts.',
    icon: '💖'
  }
])
</script>

<template>
  <div class="home-view">
    <!-- Immersive Hero Section -->
    <header class="hero">
      <div class="hero-visual-bg"></div>
      
      <div class="container hero-content animate-entry delay-100">
        <div class="hero-badge">
          <span class="pulse-icon" v-if="status?.online"></span>
          {{ status?.players?.online || 0 }} Players Online Now
        </div>
        
        <h1 class="hero-title">
          Build Your Legacy on <br/>
          <span class="text-gradient">Pure MTNetwork</span>
        </h1>
        
        <p class="hero-subtitle">
          Experience Minecraft as it was meant to be. High performance, 
          strictly non-P2W, and fully cross-play optimized for every explorer.
        </p>
        
        <div class="hero-actions">
          <router-link to="/players" class="btn-primary">
            Explore Players
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="m12 5 7 7-7 7"/></svg>
          </router-link>
          <router-link to="/server-intro" class="btn-secondary">How to Join</router-link>
        </div>
        
        <div class="quick-status glass-card animate-entry delay-300">
           <div class="status-item">
             <span class="status-label">Version</span>
             <span class="status-value">1.19.x - 1.21.x</span>
           </div>
           <div class="status-divider"></div>
           <div class="status-item">
             <span class="status-label">Mode</span>
             <span class="status-value">Hard Survival</span>
           </div>
           <div class="status-divider"></div>
           <div class="status-item">
             <span class="status-label">Address</span>
             <span class="status-value copy-address">play.mtnetwork.top</span>
           </div>
        </div>
      </div>
      
      <!-- Interactive Background Decorations -->
      <div class="hero-decor decor-top"></div>
      <div class="hero-decor decor-bottom"></div>
    </header>

    <!-- Detailed Status Component -->
    <ServerStatusAndConnection />

    <!-- Enhanced Features Section -->
    <section class="features-section">
      <div class="container">
        <div class="section-header animate-entry delay-200">
          <div class="sub-heading">THE CORE PRINCIPLES</div>
          <h2 class="section-title">Rooted in <span class="text-gradient">Excellence</span></h2>
          <p class="section-subtitle">MTNetwork is built on a foundation of performance, fairness, and community.</p>
        </div>
        
        <div class="features-grid">
          <div 
            v-for="(feature, index) in features" 
            :key="feature.title"
            class="feature-card glass-card animate-entry"
            :class="feature.class"
            :style="{ animationDelay: `${300 + index * 100}ms` }"
          >
            <div class="feature-content">
              <div class="feature-icon-wrapper">
                <div class="feature-icon">{{ feature.icon }}</div>
                <div class="icon-glow"></div>
              </div>
              <div class="feature-text">
                <h3>{{ feature.title }}</h3>
                <p>{{ feature.description }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Call to Action -->
    <section class="cta-section animate-entry delay-400">
      <div class="container">
        <div class="cta-box glass-card">
          <div class="cta-content">
            <h2>Ready to start your adventure?</h2>
            <p>Join thousands of players in our thriving cross-play community.</p>
          </div>
          <router-link to="/server-intro" class="btn-primary">Connect Now</router-link>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
.home-view {
  display: flex;
  flex-direction: column;
  gap: 6rem;
  padding-bottom: 6rem;
}

/* Hero Section */
.hero {
  position: relative;
  min-height: 90vh;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  overflow: hidden;
  padding: 4rem 0;
}

.hero-visual-bg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: 
    radial-gradient(circle at 50% 10%, rgba(59, 130, 246, 0.15), transparent 60%),
    radial-gradient(circle at 5% 90%, rgba(91, 113, 246, 0.1), transparent 40%);
  z-index: 1;
}

.hero-content {
  position: relative;
  z-index: 10;
  max-width: 900px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.hero-badge {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  padding: 8px 20px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid var(--glass-border-bright);
  backdrop-filter: blur(10px);
  color: #fff;
  border-radius: 100px;
  font-size: 0.9rem;
  font-weight: 600;
  margin-bottom: 32px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.3);
}

.pulse-icon {
  width: 8px;
  height: 8px;
  background: #10B981;
  border-radius: 50%;
  box-shadow: 0 0 10px #10B981;
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0% { transform: scale(1); opacity: 1; }
  50% { transform: scale(1.5); opacity: 0.5; }
  100% { transform: scale(1); opacity: 1; }
}

.hero-title {
  font-size: clamp(3rem, 7vw, 5.5rem);
  font-weight: 800;
  line-height: 1;
  margin-bottom: 2rem;
  letter-spacing: -0.04em;
}

.text-gradient {
  background: linear-gradient(135deg, #fff 0%, var(--primary) 50%, var(--accent) 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.hero-subtitle {
  font-size: 1.25rem;
  color: var(--text-muted);
  max-width: 650px;
  margin-bottom: 3.5rem;
  font-weight: 500;
}

.hero-actions {
  display: flex;
  gap: 1.5rem;
  margin-bottom: 4rem;
}

.hero-actions .btn-primary svg {
  margin-left: 10px;
  transition: transform 0.3s var(--transition-fast);
}

.hero-actions .btn-primary:hover svg {
  transform: translateX(5px);
}

.btn-secondary {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 16px 36px;
  background: rgba(255, 255, 255, 0.03);
  color: #fff;
  font-family: var(--heading);
  font-weight: 700;
  border: 1px solid var(--glass-border-bright);
  border-radius: var(--radius-sm);
  transition: all 0.3s var(--transition-fast);
}

.btn-secondary:hover {
  background: rgba(255, 255, 255, 0.08);
  border-color: #fff;
  transform: translateY(-2px);
}

/* Quick Status Bar */
.quick-status {
  display: flex;
  padding: 16px 40px !important;
  gap: 32px;
  border-radius: 100px !important;
  background: rgba(10, 10, 10, 0.6) !important;
}

.status-item {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.status-label {
  font-size: 0.75rem;
  text-transform: uppercase;
  color: var(--text-muted);
  letter-spacing: 1px;
  font-weight: 700;
}

.status-value {
  font-size: 1rem;
  font-weight: 600;
  color: #fff;
}

.status-divider {
  width: 1px;
  height: 30px;
  background: var(--glass-border);
  align-self: center;
}

.copy-address {
  color: var(--primary);
  cursor: pointer;
  transition: opacity 0.2s;
}

.copy-address:hover {
  opacity: 0.8;
}

/* Background decorations */
.hero-decor {
  position: absolute;
  filter: blur(120px);
  opacity: 0.2;
  z-index: 1;
  pointer-events: none;
}

.decor-top {
  width: 500px;
  height: 500px;
  background: var(--primary);
  top: -250px;
  right: -100px;
}

.decor-bottom {
  width: 400px;
  height: 400px;
  background: var(--secondary);
  bottom: -200px;
  left: -100px;
}

/* Section Header Refinement */
.sub-heading {
  font-size: 0.8rem;
  font-weight: 800;
  letter-spacing: 3px;
  color: var(--primary);
  margin-bottom: 12px;
}

.section-title {
  font-size: 3.5rem;
  margin-bottom: 1.5rem;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 2rem;
}

.feature-card {
  padding: 40px !important;
}

.feature-icon-wrapper {
  width: 72px;
  height: 72px;
  margin-bottom: 2rem;
}

.feature-icon {
  font-size: 2.5rem;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 20px;
}

.feature-text h3 {
  font-size: 1.6rem;
  margin-bottom: 1rem;
}

/* CTA Box */
.cta-box {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 60px 80px !important;
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.1), rgba(0,0,0,0.4)) !important;
  border-radius: var(--radius-xl) !important;
}

.cta-content h2 {
  font-size: 2.5rem;
  margin-bottom: 1rem;
}

@media (max-width: 1024px) {
  .features-grid { grid-template-columns: repeat(2, 1fr); }
  .cta-box { flex-direction: column; text-align: center; gap: 2rem; padding: 40px !important; }
}

@media (max-width: 768px) {
  .hero-title { font-size: 3.5rem; }
  .features-grid { grid-template-columns: 1fr; }
  .quick-status { flex-direction: column; gap: 1rem; border-radius: 24px !important; }
  .status-divider { display: none; }
  .hero-actions { flex-direction: column; width: 100%; }
}
</style>
