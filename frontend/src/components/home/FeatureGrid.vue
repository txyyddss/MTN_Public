<script setup lang="ts">
import { ref } from 'vue'

const features = ref([
  {
    title: 'Efficiency by Design',
    description: 'We leverage optimized infrastructure and advanced resource management to deliver a buttery-smooth survival experience without the premium cost.',
    icon: '⚡',
    class: 'f-large highlight-primary'
  },
  {
    title: 'Pure Vanilla',
    description: 'A genuine Minecraft experience. We focus on fair-play with robust anticheat and essential enhancements.',
    icon: '🌲',
    class: 'f-medium'
  },
  {
    title: 'Absolute Fairness',
    description: 'Strict non-P2W policy. Every player starts on equal footing, with no paid advantages or administrative bias.',
    icon: '⚖️',
    class: 'f-small'
  },
  {
    title: 'Full Cross-Play',
    description: 'Java or Bedrock. Our integrated account system ensures your progress stays with you on any device.',
    icon: '🎮',
    class: 'f-medium highlight-secondary'
  },
  {
    title: 'Built for Longevity',
    description: 'Sustainable operational model ensures stability for years. Your builds are safe in our realm.',
    icon: '🛡️',
    class: 'f-small'
  },
  {
    title: 'Community First',
    description: 'MTNetwork is a passion project dedicated to fostering a diverse and welcoming community.',
    icon: '💖',
    class: 'f-large display-accent'
  }
])
</script>

<template>
  <section class="features-section">
    <div class="container">
      <div class="section-header animate-entry delay-200">
        <div class="sub-heading">THE CORE PRINCIPLES</div>
        <h2 class="section-title display-font">Rooted in <span class="text-gradient">Excellence</span></h2>
        <p class="section-subtitle">MTNetwork is engineered for performance, fairness, and community longevity.</p>
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
              <h3 class="heading-font">{{ feature.title }}</h3>
              <p>{{ feature.description }}</p>
            </div>
          </div>
          <div class="card-accent"></div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.features-section {
  padding: 8rem 0;
  position: relative;
}

.section-header {
  margin-bottom: 6rem;
}

.sub-heading {
  font-size: 0.8rem;
  font-weight: 800;
  letter-spacing: 4px;
  color: var(--primary);
  margin-bottom: 1.5rem;
  font-family: var(--display);
}

.section-title {
  font-size: clamp(2.5rem, 5vw, 4.5rem);
  margin-bottom: 2rem;
  line-height: 1.1;
}

.section-subtitle {
  font-size: 1.1rem;
  color: var(--text-muted);
  max-width: 600px;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(12, 1fr);
  gap: 2.5rem;
}

.feature-card {
  padding: 3rem !important;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  overflow: hidden;
}

/* Asymmetrical Layout Logic */
.f-large { grid-column: span 7; }
.f-medium { grid-column: span 5; }
.f-small { grid-column: span 4; }

.highlight-primary { border-top: 2px solid var(--primary); }
.highlight-secondary { border-top: 2px solid var(--secondary); }
.display-accent { background: linear-gradient(135deg, rgba(var(--primary-rgb), 0.05), rgba(var(--accent-rgb), 0.05)); }

.feature-icon-wrapper {
  width: 64px;
  height: 64px;
  position: relative;
  margin-bottom: 2.5rem;
}

.feature-icon {
  font-size: 2.2rem;
  background: rgba(255, 255, 255, 0.03);
  width: 100%;
  height: 100%;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid var(--glass-border-bright);
}

.feature-text h3 {
  font-size: 1.75rem;
  margin-bottom: 1.25rem;
  font-family: var(--heading);
}

.feature-text p {
  color: var(--text-muted);
  line-height: 1.7;
}

.card-accent {
  position: absolute;
  top: 0;
  right: 0;
  width: 100px;
  height: 100px;
  background: radial-gradient(circle at top right, var(--primary-glow), transparent 70%);
  opacity: 0;
  transition: opacity 0.5s var(--transition-slow);
}

.feature-card:hover .card-accent {
  opacity: 1;
}

@media (max-width: 1024px) {
  .f-large, .f-medium, .f-small { grid-column: span 6; }
}

@media (max-width: 768px) {
  .features-grid { grid-template-columns: 1fr; }
  .f-large, .f-medium, .f-small { grid-column: span 1; }
  .feature-card { padding: 2.5rem !important; }
}
</style>
