<script setup lang="ts"></script>

<template>
  <section class="cta-section animate-entry delay-400">
    <div class="container">
      <div class="cta-box glass-card">
        <div class="cta-visual-bg"></div>
        <div class="cta-content">
          <h2 class="display-font">Ready to Forge Your <span class="text-gradient">Empire?</span></h2>
          <p>Join a community of elite explorers in the most refined survival environment ever engineered.</p>
        </div>
        <div class="cta-actions">
           <router-link to="/server-intro" class="btn-primary large-btn">Start Your Adventure</router-link>
           <p class="cta-note">Java & Bedrock Fully Supported</p>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.cta-section {
  padding: 8rem 0;
}

.cta-box {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 80px 100px !important;
  background: linear-gradient(135deg, rgba(var(--primary-rgb), 0.1), rgba(0,0,0,0.6)) !important;
  border-radius: var(--radius-xl) !important;
  position: relative;
  overflow: hidden;
  border: 1px solid var(--glass-border-bright) !important;
}

.cta-visual-bg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: 
    radial-gradient(circle at 100% 0%, hsla(210, 100%, 55%, 0.15), transparent 50%),
    radial-gradient(circle at 0% 100%, hsla(260, 90%, 65%, 0.1), transparent 50%);
  pointer-events: none;
}

.cta-content {
  position: relative;
  z-index: 2;
  max-width: 500px;
}

.cta-content h2 {
  font-size: clamp(2rem, 4vw, 3.5rem);
  margin-bottom: 1.5rem;
  line-height: 1.1;
}

.cta-content p {
  font-size: 1.2rem;
  color: var(--text-muted);
  line-height: 1.6;
}

.cta-actions {
  position: relative;
  z-index: 2;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1.5rem;
}

.large-btn {
  padding: 20px 48px;
  font-size: 1.1rem;
}

.cta-note {
  font-size: 0.8rem;
  color: var(--text-dim);
  font-weight: 700;
  letter-spacing: 1px;
  text-transform: uppercase;
}

.text-gradient {
  background: linear-gradient(135deg, var(--primary), var(--accent));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

@media (max-width: 1024px) {
  .cta-box { flex-direction: column; text-align: center; padding: 60px 40px !important; gap: 3rem; }
  .cta-content { max-width: 100%; }
}
</style>
