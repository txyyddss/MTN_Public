<script setup lang="ts">
import { ref, onUnmounted } from 'vue'

const imageModules = import.meta.glob('/public/gallary/*.{png,jpg,jpeg,webp}', { eager: true })
const images = Object.keys(imageModules)
  .map(path => path.split('/').pop() || '')
  .filter(img => img !== '')
  .sort((a, b) => b.localeCompare(a)) // Sort reversed to show newest (by timestamp filename) first

const selectedImage = ref<string | null>(null)

const openLightbox = (img: string) => {
  selectedImage.value = img
  document.body.style.overflow = 'hidden'
}

const closeLightbox = () => {
  selectedImage.value = null
  document.body.style.overflow = ''
}

onUnmounted(() => {
  document.body.style.overflow = ''
})
</script>

<template>
  <div class="gallery-view container animate-entry">
    <div class="header">
      <h1 class="title">Gallery</h1>
      <p class="subtitle">A visual journey through the landscapes and builds of MT Network.</p>
    </div>

    <div class="waterfall">
      <div 
        v-for="(img, index) in images" 
        :key="img" 
        class="gallery-item glass-card"
        :style="{ animationDelay: `${index * 50}ms` }"
        @click="openLightbox(img)"
      >
        <img 
          :src="`/gallary/${img}`" 
          :alt="`Gallery image ${index + 1}`" 
          loading="lazy" 
          class="gallery-img"
        />
        <div class="overlay">
          <span>View Detailed</span>
        </div>
      </div>
    </div>

    <Transition name="fade">
      <div v-if="selectedImage" class="lightbox" @click="closeLightbox">
        <button class="close-btn" @click="closeLightbox">&times;</button>
        <img :src="`/gallary/${selectedImage}`" class="lightbox-img" @click.stop />
      </div>
    </Transition>
  </div>
</template>

<style scoped>
.gallery-view {
  padding-top: 3rem;
  padding-bottom: 6rem;
}

.header {
  margin-bottom: 4rem;
  text-align: center;
}

.title {
  font-size: 3.5rem;
  margin-bottom: 1rem;
  background: linear-gradient(to right, #fff, var(--primary));
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
}

.subtitle {
  font-size: 1.1rem;
  color: var(--text-muted);
  max-width: 600px;
  margin: 0 auto;
}

.waterfall {
  column-count: 1;
  column-gap: 1.5rem;
  width: 100%;
}

@media (min-width: 640px) {
  .waterfall {
    column-count: 2;
  }
}

@media (min-width: 1024px) {
  .waterfall {
    column-count: 3;
  }
}

.gallery-item {
  display: inline-block;
  width: 100%;
  margin-bottom: 1.5rem;
  padding: 0;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  border-radius: var(--radius-lg);
  border: 1px solid var(--glass-border);
  background: var(--glass-bg);
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}

.gallery-item:hover {
  transform: translateY(-8px) scale(1.02);
  border-color: var(--primary);
  box-shadow: 0 30px 60px rgba(0, 0, 0, 0.5);
}

.gallery-img {
  width: 100%;
  height: auto;
  display: block;
  transition: transform 0.6s cubic-bezier(0.16, 1, 0.3, 1);
}

.gallery-item:hover .gallery-img {
  transform: scale(1.05);
}

.overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(to top, rgba(0,0,0,0.8), transparent 40%);
  display: flex;
  align-items: flex-end;
  justify-content: center;
  padding-bottom: 1.5rem;
  opacity: 0;
  transition: opacity 0.3s;
}

.gallery-item:hover .overlay {
  opacity: 1;
}

.overlay span {
  color: #fff;
  font-family: var(--heading);
  font-weight: 700;
  letter-spacing: 1px;
  text-transform: uppercase;
  font-size: 0.8rem;
  background: var(--primary);
  padding: 6px 16px;
  border-radius: 20px;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
}

/* Lightbox */
.lightbox {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.95);
  backdrop-filter: blur(10px);
  z-index: 1000;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 2rem;
  overflow-y: auto;
}

.lightbox-img {
  max-width: 95vw;
  max-height: 90vh;
  object-fit: contain;
  box-shadow: 0 0 50px rgba(0, 0, 0, 0.5);
  border-radius: var(--radius-md);
  margin: auto;
}

.close-btn {
  position: fixed;
  top: 2rem;
  right: 2rem;
  background: none;
  border: none;
  color: #fff;
  font-size: 3rem;
  cursor: pointer;
  line-height: 1;
  transition: 0.3s;
}

.close-btn:hover {
  color: var(--primary);
  transform: scale(1.1);
}

/* Transitions */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
