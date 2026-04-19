<script setup lang="ts">
import { ref, onUnmounted, computed } from 'vue'

const imageModules = import.meta.glob('/public/gallary/*.{png,jpg,jpeg,webp}', { eager: true })
const images = Object.keys(imageModules)
  .map(path => path.split('/').pop() || '')
  .filter(img => img !== '')
  .sort((a, b) => b.localeCompare(a)) // Sort reversed to show newest (by timestamp filename) first

const selectedIndex = ref<number | null>(null)
const selectedImage = computed(() => selectedIndex.value !== null ? images[selectedIndex.value] : null)

const openLightbox = (index: number) => {
  selectedIndex.value = index
  document.body.style.overflow = 'hidden'
}

const closeLightbox = () => {
  selectedIndex.value = null
  document.body.style.overflow = ''
}

const nextImage = () => {
  if (selectedIndex.value === null) return
  selectedIndex.value = (selectedIndex.value + 1) % images.length
}

const prevImage = () => {
  if (selectedIndex.value === null) return
  selectedIndex.value = (selectedIndex.value - 1 + images.length) % images.length
}

const handleKeydown = (e: KeyboardEvent) => {
  if (selectedIndex.value === null) return
  if (e.key === 'Escape') closeLightbox()
  if (e.key === 'ArrowRight') nextImage()
  if (e.key === 'ArrowLeft') prevImage()
}

window.addEventListener('keydown', handleKeydown)

onUnmounted(() => {
  window.removeEventListener('keydown', handleKeydown)
  document.body.style.overflow = ''
})
</script>

<template>
  <div class="gallery-view container">
    <header class="gallery-header animate-entry">
      <div class="title-wrap">
        <h1 class="glitch-title">ARCHIVES</h1>
        <div class="status-line">
          <span class="count">{{ images.length }} CAPTURES</span>
          <span class="separator">/</span>
          <span class="desc">A VISUAL CHRONICLE OF EVOLUTION</span>
        </div>
      </div>
    </header>

    <div class="waterfall-container">
      <div 
        v-for="(img, index) in images" 
        :key="img" 
        class="gallery-item glass-card animate-entry"
        :style="{ animationDelay: `${(index % 20) * 40}ms` }"
        @click="openLightbox(index)"
      >
        <div class="card-glow"></div>
        <img 
          :src="`/gallary/${img}`" 
          :alt="`Archive Capture ${index + 1}`" 
          loading="lazy" 
          class="gallery-img"
        />
        <div class="item-meta">
          <span class="index">#{{ images.length - index }}</span>
          <span class="action">DECRYPT VIEW</span>
        </div>
      </div>
    </div>

    <!-- Immersive Lightbox -->
    <Transition name="lightbox-zoom">
      <div v-if="selectedImage" class="lightbox-overlay" @click="closeLightbox">
        <div class="lightbox-controls" @click.stop>
          <button class="nav-btn prev" @click="prevImage" aria-label="Previous">
             <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="m15 18-6-6 6-6"/></svg>
          </button>
          <div class="lightbox-content">
            <img :src="`/gallary/${selectedImage}`" class="main-image" />
            <div class="image-label">CAPTURE {{ images.length - (selectedIndex ?? 0) }}</div>
          </div>
          <button class="nav-btn next" @click="nextImage" aria-label="Next">
             <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><path d="m9 18 6-6-6-6"/></svg>
          </button>
        </div>
        <button class="close-lux" @click="closeLightbox">CLOSE</button>
      </div>
    </Transition>
  </div>
</template>

<style scoped>
.gallery-view {
  padding-top: 6rem;
  padding-bottom: 10rem;
}

.gallery-header {
  margin-bottom: 6rem;
}

.glitch-title {
  font-size: 5rem;
  font-weight: 800;
  letter-spacing: -3px;
  line-height: 0.8;
  margin-bottom: 1.5rem;
}

.status-line {
  display: flex;
  align-items: center;
  gap: 12px;
  font-family: var(--heading);
  font-size: 0.8rem;
  font-weight: 700;
  letter-spacing: 2px;
}

.status-line .count { color: var(--primary); }
.status-line .separator { opacity: 0.2; }
.status-line .desc { color: var(--text-muted); }

/* Waterfall Implementation */
.waterfall-container {
  columns: 1;
  column-gap: 1.5rem;
}

@media (min-width: 640px) { .waterfall-container { columns: 2; } }
@media (min-width: 1024px) { .waterfall-container { columns: 3; } }
@media (min-width: 1600px) { .waterfall-container { columns: 4; } }

.gallery-item {
  break-inside: avoid;
  display: block;
  width: 100%;
  margin-bottom: 1.5rem;
  padding: 0 !important;
  border-radius: var(--radius-lg) !important;
  overflow: hidden;
  cursor: pointer;
  transition: all 0.5s var(--transition-slow);
}

.card-glow {
  position: absolute;
  inset: 0;
  background: radial-gradient(circle at top right, var(--primary-glow), transparent 60%);
  opacity: 0;
  transition: opacity 0.5s;
  z-index: 1;
}

.gallery-item:hover {
  transform: translateY(-10px) scale(1.01);
  border-color: var(--primary);
  box-shadow: 0 40px 80px rgba(0,0,0,0.6);
}

.gallery-item:hover .card-glow { opacity: 0.4; }

.gallery-img {
  width: 100%;
  height: auto;
  display: block;
  transition: transform 0.8s var(--transition-slow);
}

.gallery-item:hover .gallery-img { transform: scale(1.08); }

.item-meta {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 20px;
  background: linear-gradient(to top, rgba(0,0,0,0.9), transparent);
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  opacity: 0;
  transform: translateY(10px);
  transition: all 0.4s;
  z-index: 2;
}

.gallery-item:hover .item-meta {
  opacity: 1;
  transform: translateY(0);
}

.index { font-family: var(--heading); font-weight: 800; font-size: 1.2rem; color: #fff; }
.action { font-size: 0.65rem; font-weight: 800; color: var(--primary); letter-spacing: 1px; }

/* Lightbox Elevation */
.lightbox-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.98);
  backdrop-filter: blur(20px);
  z-index: 2000;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.lightbox-controls {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 2rem;
  flex: 1;
}

.nav-btn {
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid var(--glass-border);
  color: #fff;
  width: 60px;
  height: 60px;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s;
}

.nav-btn:hover {
  background: var(--primary);
  color: #000;
  border-color: var(--primary);
  transform: scale(1.1);
}

.lightbox-content {
  position: relative;
  max-width: 85vw;
  max-height: 80vh;
}

.main-image {
  max-width: 100%;
  max-height: 80vh;
  object-fit: contain;
  border-radius: 12px;
  box-shadow: 0 0 100px rgba(0,0,0,0.8);
}

.image-label {
  position: absolute;
  top: -40px;
  left: 0;
  font-family: var(--heading);
  font-weight: 800;
  font-size: 0.8rem;
  letter-spacing: 4px;
  color: var(--primary);
}

.close-lux {
  position: absolute;
  top: 3rem;
  right: 3rem;
  background: transparent;
  border: none;
  color: #fff;
  font-family: var(--heading);
  font-weight: 800;
  letter-spacing: 4px;
  cursor: pointer;
  padding: 10px;
  transition: all 0.3s;
}

.close-lux:hover {
  color: var(--primary);
  text-shadow: 0 0 15px var(--primary-glow);
}

/* Transitions */
.lightbox-zoom-enter-active,
.lightbox-zoom-leave-active {
  transition: all 0.5s cubic-bezier(0.16, 1, 0.3, 1);
}

.lightbox-zoom-enter-from,
.lightbox-zoom-leave-to {
  opacity: 0;
  transform: scale(0.95);
}

@media (max-width: 768px) {
  .nav-btn { display: none; }
  .close-lux { top: 1rem; right: 1rem; }
  .glitch-title { font-size: 3rem; }
}
</style>
