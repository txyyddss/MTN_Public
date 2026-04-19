<script setup lang="ts">
import { onMounted, onUnmounted, useTemplateRef, watch } from 'vue'

const props = defineProps<{
  skinUrl: string
  width?: number
  height?: number
}>()

const canvasRef = useTemplateRef<HTMLCanvasElement>('viewerCanvas')
let viewer: import('skinview3d').SkinViewer | null = null

async function initViewer(): Promise<void> {
  if (!canvasRef.value) {
    return
  }

  const skinview3d = await import('skinview3d')

  if (viewer) {
    viewer.dispose()
  }

  viewer = new skinview3d.SkinViewer({
    canvas: canvasRef.value,
    width: props.width ?? 250,
    height: props.height ?? 300,
    skin: props.skinUrl
  })

  viewer.controls.enableZoom = false
  viewer.controls.enablePan = true
  viewer.animation = new skinview3d.IdleAnimation()
  viewer.background = null
}

onMounted(() => {
  void initViewer()
})

onUnmounted(() => {
  viewer?.dispose()
})

watch(
  () => props.skinUrl,
  async (skinUrl) => {
    if (viewer) {
      await viewer.loadSkin(skinUrl)
      return
    }

    await initViewer()
  }
)

watch(
  () => [props.width, props.height],
  () => {
    viewer?.setSize(props.width ?? 250, props.height ?? 300)
  }
)
</script>

<template>
  <div class="skin-viewer-container">
    <canvas ref="viewerCanvas"></canvas>
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
