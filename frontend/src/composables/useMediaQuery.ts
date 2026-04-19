import { onMounted, onUnmounted, shallowRef } from 'vue'

export function useMediaQuery(query: string) {
  const matches = shallowRef(false)
  let mediaQuery: MediaQueryList | null = null

  function syncMatch(event?: MediaQueryListEvent): void {
    matches.value = event?.matches ?? mediaQuery?.matches ?? false
  }

  onMounted(() => {
    mediaQuery = window.matchMedia(query)
    syncMatch()
    mediaQuery.addEventListener('change', syncMatch)
  })

  onUnmounted(() => {
    mediaQuery?.removeEventListener('change', syncMatch)
  })

  return {
    matches
  }
}
