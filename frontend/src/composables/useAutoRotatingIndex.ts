import { onUnmounted, shallowRef, watch, type Ref } from 'vue'

export function useAutoRotatingIndex(
  length: Readonly<Ref<number>>,
  enabled: Readonly<Ref<boolean>>,
  intervalMs = 4200
) {
  const currentIndex = shallowRef(0)
  let timer: ReturnType<typeof window.setInterval> | null = null

  function stopRotation(): void {
    if (!timer) {
      return
    }

    window.clearInterval(timer)
    timer = null
  }

  function next(): void {
    if (length.value <= 1) {
      return
    }

    currentIndex.value = (currentIndex.value + 1) % length.value
  }

  function previous(): void {
    if (length.value <= 1) {
      return
    }

    currentIndex.value = (currentIndex.value - 1 + length.value) % length.value
  }

  function startRotation(): void {
    stopRotation()

    if (!enabled.value || length.value <= 1) {
      return
    }

    timer = window.setInterval(() => {
      next()
    }, intervalMs)
  }

  watch(
    [length, enabled],
    () => {
      if (currentIndex.value >= length.value) {
        currentIndex.value = 0
      }

      startRotation()
    },
    { immediate: true }
  )

  onUnmounted(() => {
    stopRotation()
  })

  return {
    currentIndex,
    next,
    previous
  }
}
