<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch } from 'vue'
import * as skinview3d from 'skinview3d'

const props = defineProps<{
  skinUrl: string
  width?: number
  height?: number
}>()

const canvasRef = ref<HTMLCanvasElement | null>(null)
let viewer: skinview3d.SkinViewer | null = null

const initViewer = () => {
  if (!canvasRef.value) return

  if (viewer) {
    viewer.dispose()
  }

  viewer = new skinview3d.SkinViewer({
    canvas: canvasRef.value,
    width: props.width || 250,
    height: props.height || 300,
    skin: props.skinUrl
  })

  // Configure controls
  viewer.controls.enableZoom = false
  viewer.controls.enablePan = true // Allow panning if needed, or keep false

  // Set idle animation
  viewer.animation = new skinview3d.IdleAnimation()
  
  // Set background to transparent (default is transparent, but explicitly setting it)
  viewer.background = null
}

onMounted(() => {
  initViewer()
})

onUnmounted(() => {
  if (viewer) {
    viewer.dispose()
  }
})

watch(() => props.skinUrl, (newUrl) => {
  if (viewer) {
    viewer.loadSkin(newUrl)
  }
})

watch(() => [props.width, props.height], () => {
  if (viewer) {
    viewer.setSize(props.width || 250, props.height || 300)
  }
})
</script>

<template>
  <div class="skin-viewer-container">
    <canvas ref="canvasRef"></canvas>
  </div>
</template>

<style scoped>
.skin-viewer-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  min-height: 300px;
}

canvas {
  cursor: grab;
  outline: none;
}

canvas:active {
  cursor: grabbing;
}
</style>
