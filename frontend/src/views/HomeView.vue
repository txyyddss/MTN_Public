<script setup lang="ts">
import { ref } from 'vue'
import ServerStatusAndConnection from '@/components/ServerStatusAndConnection.vue'

const features = ref([
  {
    title: 'Zero-Cost Infrastructure',
    description: 'Goal: using rubbish hardware and relatively zero-cost homeserver to create the smoothest MC Server.',
    icon: '⚡'
  },
  {
    title: 'Vanilla Survival',
    description: 'Positioning: Vanilla survival server with strong anticheat and small modifications to enhance gameplay.',
    icon: '🌲'
  },
  {
    title: 'Pure Fairness',
    description: 'No discrimination between VIPs and normal players, zero pay-to-win, no admin privilege abuse.',
    icon: '⚖️'
  },
  {
    title: 'Cross-play Supported',
    description: 'Supports both Java and Bedrock connection with full account binding support.',
    icon: '🎮'
  },
  {
    title: 'Long-Term Stability',
    description: 'Stable and open for a long time due to low-cost maintenance requirements.',
    icon: '🛡️'
  },
  {
    title: 'Non-Profit Community',
    description: 'It is a strictly non-profit server aimed at bringing players together.',
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
        <h2 class="section-title animate-entry delay-200">Server Philosophy</h2>
        <div class="features-grid">
          <div 
            v-for="(feature, index) in features" 
            :key="feature.title"
            class="feature-card glass-card animate-entry"
            :style="{ animationDelay: `${300 + index * 100}ms` }"
          >
            <div class="feature-icon">{{ feature.icon }}</div>
            <h3>{{ feature.title }}</h3>
            <p>{{ feature.description }}</p>
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
}

.section-title {
  text-align: center;
  font-size: 2.5rem;
  margin-bottom: 3rem;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.feature-card {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.feature-icon {
  font-size: 2.5rem;
  width: 60px;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 12px;
  border: 1px solid var(--glass-border);
}

.feature-card h3 {
  font-size: 1.25rem;
}
.feature-card p {
  color: var(--text-muted);
  font-size: 0.95rem;
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
