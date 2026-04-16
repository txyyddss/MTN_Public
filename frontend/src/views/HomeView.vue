<script setup lang="ts">
import { ref } from 'vue'
import ServerStatusAndConnection from '@/components/ServerStatusAndConnection.vue'

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
    <!-- Hero Section -->
    <header class="hero">
      <div class="container hero-content animate-entry delay-100">
        <div class="hero-badge">Welcome to MTNetwork</div>
        <h1 class="hero-title">The Smoothest <span class="text-gradient">Vanilla Experience</span></h1>
        <p class="hero-subtitle">
          Join our cross-play supported, low-latency, and purely vanilla Minecraft server 
          built on zero-cost infrastructure without compromising quality.
        </p>
        <div class="hero-actions">
          <router-link to="/players" class="btn-primary">View Players</router-link>
          <a href="/wiki" class="btn-secondary">Read Wiki</a>
        </div>
      </div>
      
      <!-- Background Decorative Elements -->
      <div class="glow-orb orb-1"></div>
      <div class="glow-orb orb-2"></div>
    </header>

    <!-- Server Status & Connection -->
    <ServerStatusAndConnection />

    <!-- Features Section -->
    <section class="features-section">
      <div class="container">
        <div class="section-header animate-entry delay-200">
          <h2 class="section-title">Server Philosophy</h2>
          <p class="section-subtitle">The principles that drive MTNetwork forward.</p>
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
      <div class="container cta-container glass-card">
        <h2>Ready to join the adventure?</h2>
        <p>Connect using Java or Bedrock and link your accounts effortlessly.</p>
        <router-link to="/server-intro" class="btn-primary mt-4">Learn How to Connect</router-link>
      </div>
    </section>
  </div>
</template>

<style scoped>
.home-view {
  display: flex;
  flex-direction: column;
  gap: 4rem;
  padding-bottom: 4rem;
}

/* Hero Section */
.hero {
  position: relative;
  padding: 8rem 0 4rem;
  text-align: center;
  overflow: hidden;
}

.hero-content {
  position: relative;
  z-index: 10;
  max-width: 800px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.hero-badge {
  display: inline-block;
  padding: 6px 16px;
  background: var(--glass-bg);
  border: 1px solid var(--primary);
  color: var(--primary);
  border-radius: 20px;
  font-size: 0.9rem;
  font-weight: 600;
  margin-bottom: 24px;
}

.hero-title {
  font-size: clamp(2.5rem, 5vw, 4.5rem);
  font-weight: 800;
  line-height: 1.1;
  margin-bottom: 1.5rem;
}

.text-gradient {
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text; /* fixed lint */
}

.hero-subtitle {
  font-size: 1.1rem;
  color: var(--text-muted);
  max-width: 600px;
  margin-bottom: 2.5rem;
}

.hero-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
}

.btn-secondary {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 12px 24px;
  background: var(--glass-bg);
  color: #fff;
  font-family: var(--heading);
  font-weight: 600;
  border: 1px solid var(--glass-border);
  border-radius: var(--radius-sm);
  transition: all 0.3s ease;
}

.btn-secondary:hover {
  background: rgba(255, 255, 255, 0.1);
  border-color: rgba(255, 255, 255, 0.3);
}

/* Glow effects */
.glow-orb {
  position: absolute;
  width: 400px;
  height: 400px;
  border-radius: 50%;
  filter: blur(100px);
  z-index: 1;
  opacity: 0.4;
  animation: float 10s infinite alternate ease-in-out;
}
.orb-1 {
  background: var(--primary);
  top: -100px;
  left: -100px;
}
.orb-2 {
  background: var(--secondary);
  bottom: -100px;
  right: -100px;
  animation-delay: -5s;
}

@keyframes float {
  0% { transform: translate(0, 0); }
  100% { transform: translate(50px, 50px); }
}

/* Features Section */
.features-section {
  position: relative;
  z-index: 10;
  padding: 4rem 0;
}

.section-header {
  text-align: center;
  margin-bottom: 4rem;
}

.section-title {
  font-size: 3rem;
  font-weight: 800;
  margin-bottom: 1rem;
  background: linear-gradient(135deg, #fff 0%, var(--text-muted) 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.section-subtitle {
  color: var(--text-muted);
  font-size: 1.1rem;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-auto-rows: minmax(200px, auto);
  gap: 1.5rem;
}

.feature-card {
  position: relative;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  border-radius: var(--radius-xl);
  height: 100%;
}

.feature-card.feature-large {
  grid-column: span 2;
  grid-row: span 1;
}

.feature-card.feature-medium {
  grid-column: span 1;
  grid-row: span 1;
}

.feature-content {
  height: 100%;
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  padding: 8px;
}

.feature-icon-wrapper {
  position: relative;
  width: 56px;
  height: 56px;
  flex-shrink: 0;
}

.feature-icon {
  position: relative;
  z-index: 2;
  font-size: 2rem;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(255, 255, 255, 0.03);
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  transition: all 0.3s ease;
}

.icon-glow {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 40px;
  height: 40px;
  background: var(--primary);
  filter: blur(20px);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.feature-card:hover .icon-glow {
  opacity: 0.4;
}

.feature-card:hover .feature-icon {
  transform: scale(1.1);
  border-color: var(--primary);
  background: rgba(59, 130, 246, 0.05);
}

.feature-text h3 {
  font-size: 1.4rem;
  margin-bottom: 0.75rem;
  transition: color 0.3s ease;
}

.feature-card:hover h3 {
  color: var(--primary);
}

.feature-text p {
  color: var(--text-muted);
  font-size: 0.95rem;
  line-height: 1.6;
}

@media (max-width: 1024px) {
  .features-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 768px) {
  .features-grid {
    grid-template-columns: 1fr;
  }
  .feature-card.feature-large {
    grid-column: span 1;
  }
  .section-title {
    font-size: 2.5rem;
  }
}

/* CTA */
.cta-container {
  text-align: center;
  padding: 4rem 2rem;
  background: linear-gradient(135deg, rgba(21, 27, 43, 0.8), rgba(99, 102, 241, 0.1));
}

.cta-container h2 {
  font-size: 2rem;
  margin-bottom: 1rem;
}

.cta-container p {
  color: var(--text-muted);
  margin-bottom: 2rem;
}

@media (max-width: 768px) {
  .hero {
    padding: 6rem 0 3rem;
  }
  .hero-actions {
    flex-direction: column;
    width: 100%;
    max-width: 300px;
  }
}
</style>
