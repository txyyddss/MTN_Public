interface CacheEntry {
    data: any;
    expiry: number;
    ttl: number;
    lastAccessed: number;
}

const cache = new Map<string, CacheEntry>();
// Maps in-flight URLs to their Promise so concurrent callers can await the same fetch
const inFlight = new Map<string, Promise<any>>();

// Maximum idle time before we stop auto-refreshing (5 minutes)
const MAX_IDLE_TIME = 300000;

/**
 * Fetches data with an in-memory cache and TTL.
 */
export async function fetchWithCache<T = any>(url: string, ttl: number = 60000): Promise<T> {
    const now = Date.now();
    const entry = cache.get(url);

    if (entry) {
        entry.lastAccessed = now;
        const remaining = entry.expiry - now;

        if (remaining > 0) {
            // Proactive refresh: if less than 30% of TTL remains, refresh in background
            if (remaining < ttl * 0.3 && !inFlight.has(url)) {
                console.log(`[DataCache] Proactive refresh: ${url}`);
                triggerBackgroundRefresh(url, ttl);
            }
            return entry.data;
        } else {
            // Serve expired cache but trigger background refresh (stale-while-revalidate)
            if (!inFlight.has(url)) {
                console.log(`[DataCache] Serving expired data and refreshing: ${url}`);
                triggerBackgroundRefresh(url, ttl);
            }
            return entry.data;
        }
    }

    // If a fetch is already in flight for this URL, await the same promise
    if (inFlight.has(url)) {
        console.log(`[DataCache] Awaiting in-flight: ${url}`);
        return inFlight.get(url)!;
    }

    console.log(`[DataCache] Cache miss/expire: ${url}`);
    return performFetch(url, ttl);
}

async function performFetch(url: string, ttl: number): Promise<any> {
    const promise = (async () => {
        try {
            const response = await fetch(url);
            if (!response.ok) throw new Error(`Fetch failed: ${response.statusText}`);
            const data = await response.json();

            cache.set(url, {
                data,
                expiry: Date.now() + ttl,
                ttl,
                lastAccessed: Date.now()
            });
            return data;
        } catch (error) {
            console.error(`[DataCache] Fetch error for ${url}:`, error);
            throw error;
        } finally {
            inFlight.delete(url);
        }
    })();

    inFlight.set(url, promise);
    return promise;
}

function triggerBackgroundRefresh(url: string, ttl: number) {
    performFetch(url, ttl).catch(() => { });
}

// Background worker to auto-refresh active cache entries
setInterval(() => {
    const now = Date.now();
    for (const [url, entry] of cache.entries()) {
        const idleTime = now - entry.lastAccessed;

        // Only auto-refresh if the data was recently accessed
        if (idleTime < MAX_IDLE_TIME) {
            const remaining = entry.expiry - now;
            // Refresh if less than 15 seconds remain or 25% of TTL
            const threshold = Math.min(15000, entry.ttl * 0.25);

            if (remaining < threshold && !inFlight.has(url)) {
                console.log(`[DataCache] Background auto-refresh: ${url}`);
                triggerBackgroundRefresh(url, entry.ttl);
            }
        } else {
            // Cleanup stale cache
            if (now > entry.expiry + MAX_IDLE_TIME) {
                cache.delete(url);
            }
        }
    }
}, 10000); // Check every 10 seconds

export function clearDataCache(url?: string) {
    if (url) cache.delete(url);
    else cache.clear();
}

