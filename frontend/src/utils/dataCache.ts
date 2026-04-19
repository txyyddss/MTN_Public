interface CacheEntry {
  data: unknown
  expiry: number
  ttl: number
  lastAccessed: number
}

const cache = new Map<string, CacheEntry>()
const inFlight = new Map<string, Promise<unknown>>()
const maxIdleTime = 300000
const isDev = import.meta.env.DEV

function debugLog(message: string, url: string): void {
  if (isDev) {
    console.debug(`[DataCache] ${message}: ${url}`)
  }
}

export async function fetchWithCache<T>(url: string, ttl = 60000): Promise<T> {
  const now = Date.now()
  const entry = cache.get(url)

  if (entry) {
    entry.lastAccessed = now
    const remaining = entry.expiry - now

    if (remaining > 0) {
      if (remaining < ttl * 0.3 && !inFlight.has(url)) {
        debugLog('Proactive refresh', url)
        triggerBackgroundRefresh(url, ttl)
      }

      return entry.data as T
    }

    if (!inFlight.has(url)) {
      debugLog('Serving stale data and refreshing', url)
      triggerBackgroundRefresh(url, ttl)
    }

    return entry.data as T
  }

  const existingRequest = inFlight.get(url)
  if (existingRequest) {
    debugLog('Awaiting in-flight request', url)
    return (await existingRequest) as T
  }

  debugLog('Cache miss', url)
  return (await performFetch(url, ttl)) as T
}

async function performFetch(url: string, ttl: number): Promise<unknown> {
  const promise = (async () => {
    try {
      const response = await fetch(url)
      if (!response.ok) {
        throw new Error(`Fetch failed: ${response.statusText}`)
      }

      const data = (await response.json()) as unknown
      cache.set(url, {
        data,
        expiry: Date.now() + ttl,
        ttl,
        lastAccessed: Date.now()
      })

      return data
    } catch (error) {
      console.error(`Data cache fetch failed for ${url}`, error)
      throw error
    } finally {
      inFlight.delete(url)
    }
  })()

  inFlight.set(url, promise)
  return promise
}

function triggerBackgroundRefresh(url: string, ttl: number): void {
  void performFetch(url, ttl).catch(() => undefined)
}

setInterval(() => {
  const now = Date.now()

  for (const [url, entry] of cache.entries()) {
    const idleTime = now - entry.lastAccessed

    if (idleTime < maxIdleTime) {
      const remaining = entry.expiry - now
      const threshold = Math.min(15000, entry.ttl * 0.25)

      if (remaining < threshold && !inFlight.has(url)) {
        debugLog('Background auto-refresh', url)
        triggerBackgroundRefresh(url, entry.ttl)
      }

      continue
    }

    if (now > entry.expiry + maxIdleTime) {
      cache.delete(url)
    }
  }
}, 10000)

export function clearDataCache(url?: string): void {
  if (url) {
    cache.delete(url)
    return
  }

  cache.clear()
}

