import { onUnmounted, shallowRef, watch, type Ref } from 'vue'

interface AutoRotatingIndexOptions {
  intervalMs?: number
  initialIndex?: number
}

export function useAutoRotatingIndex(
  length: Readonly<Ref<number>>,
  enabled: Readonly<Ref<boolean>>,
  options: AutoRotatingIndexOptions | number = {}
) {
  const resolvedOptions = typeof options === 'number' ? { intervalMs: options } : options
  const intervalMs = resolvedOptions.intervalMs ?? 4200
  const currentIndex = shallowRef(normalizeIndex(resolvedOptions.initialIndex ?? 0))
  let timer: ReturnType<typeof window.setInterval> | null = null

  function normalizeIndex(index: number): number {
    if (length.value <= 0) {
      return 0
    }

    return ((index % length.value) + length.value) % length.value
  }

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

    currentIndex.value = normalizeIndex(currentIndex.value + 1)
    startRotation()
  }

  function previous(): void {
    if (length.value <= 1) {
      return
    }

    currentIndex.value = normalizeIndex(currentIndex.value - 1)
    startRotation()
  }

  function setIndex(index: number): void {
    if (length.value <= 0) {
      currentIndex.value = 0
      stopRotation()
      return
    }

    currentIndex.value = normalizeIndex(index)
    startRotation()
  }

  function rotateNext(): void {
    if (length.value <= 1) {
      return
    }

    currentIndex.value = normalizeIndex(currentIndex.value + 1)
  }

  function startRotation(): void {
    stopRotation()

    if (!enabled.value || length.value <= 1) {
      return
    }

    timer = window.setInterval(() => {
      rotateNext()
    }, intervalMs)
  }

  watch(
    [length, enabled],
    () => {
      if (length.value <= 0) {
        currentIndex.value = 0
      } else if (currentIndex.value >= length.value) {
        currentIndex.value = normalizeIndex(currentIndex.value)
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
    previous,
    setIndex
  }
}
