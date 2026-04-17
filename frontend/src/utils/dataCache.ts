interface CacheEntry {
    data: any;
    expiry: number;
    ttl: number;
}

const cache = new Map<string, CacheEntry>();
const refreshing = new Set<string>();

/**
 * Fetches data with an in-memory cache and TTL.
 * @param url The URL to fetch
 * @param ttl Time to live in milliseconds (default 60,000ms / 1 minute)
 */
export async function fetchWithCache<T = any>(url: string, ttl: number = 60000): Promise<T> {
    const now = Date.now();
    const entry = cache.get(url);

    if (entry) {
        const remaining = entry.expiry - now;

        // If still valid
        if (remaining > 0) {
            // Proactive refresh: if less than 20% of TTL remains, refresh in background
            if (remaining < ttl * 0.2 && !refreshing.has(url)) {
                console.log(`[DataCache] Proactive refresh triggered for: ${url}`);
                triggerBackgroundRefresh(url, ttl);
            }

            console.log(`[DataCache] Cache hit: ${url}`);
            return entry.data;
        }
    }

    console.log(`[DataCache] Cache miss/expire: ${url}`);
    return performFetch(url, ttl);
}

async function performFetch(url: string, ttl: number) {
    refreshing.add(url);
    try {
        const response = await fetch(url);
        if (!response.ok) {
            throw new Error(`Failed to fetch ${url}: ${response.statusText}`);
        }

        const data = await response.json();
        cache.set(url, {
            data,
            expiry: Date.now() + ttl,
            ttl
        });
        return data;
    } finally {
        refreshing.delete(url);
    }
}

function triggerBackgroundRefresh(url: string, ttl: number) {
    performFetch(url, ttl).catch(err => {
        console.warn(`[DataCache] Background refresh failed for ${url}:`, err);
    });
}

/**
 * Force clear the cache for a specific URL or all URLs.
 */
export function clearDataCache(url?: string) {
    if (url) {
        cache.delete(url);
    } else {
        cache.clear();
    }
}
