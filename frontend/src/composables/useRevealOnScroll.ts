import { onMounted, onUnmounted, shallowRef, useTemplateRef } from 'vue'

interface RevealOptions {
  rootMargin?: string
  threshold?: number
}

export function useRevealOnScroll<T extends Element>(refName: string, options: RevealOptions = {}) {
  const elementRef = useTemplateRef<T>(refName)
  const revealed = shallowRef(false)
  let observer: IntersectionObserver | null = null

  function reveal(): void {
    revealed.value = true
    observer?.disconnect()
    observer = null
  }

  onMounted(() => {
    if (typeof window === 'undefined' || window.matchMedia('(prefers-reduced-motion: reduce)').matches) {
      reveal()
      return
    }

    if (!elementRef.value || typeof IntersectionObserver === 'undefined') {
      reveal()
      return
    }

    observer = new IntersectionObserver(
      (entries) => {
        if (entries.some((entry) => entry.isIntersecting)) {
          reveal()
        }
      },
      {
        rootMargin: options.rootMargin ?? '0px 0px -12% 0px',
        threshold: options.threshold ?? 0.12
      }
    )

    observer.observe(elementRef.value)
  })

  onUnmounted(() => {
    observer?.disconnect()
  })

  return {
    elementRef,
    revealed
  }
}
