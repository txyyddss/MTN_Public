import { shallowRef } from 'vue'

export function usePreloader() {
    const isReady = shallowRef(false)
    let initPromise: Promise<void> | null = null

    const waitForLoad = () =>
        new Promise<void>((resolve) => {
            if (document.readyState === 'complete') {
                resolve()
                return
            }

            const onLoad = () => {
                window.removeEventListener('load', onLoad)
                resolve()
            }

            window.addEventListener('load', onLoad, { once: true })
        })

    const initPreloading = () => {
        if (initPromise) {
            return initPromise
        }

        initPromise = waitForLoad()
            .finally(() => {
                isReady.value = true
            })

        return initPromise
    }

    return {
        isReady,
        initPreloading
    }
}
