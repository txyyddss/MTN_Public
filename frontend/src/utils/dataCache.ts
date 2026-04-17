interface CacheEntry {
    data: any;
    expiry: number;
}

const cache = new Map<string, CacheEntry>();

/**
 * Fetches data with an in-memory cache and TTL.
 * @param url The URL to fetch
 * @param ttl Time to live in milliseconds (default 60,000ms / 1 minute)
 */
export async function fetchWithCache<T = any>(url: string, ttl: number = 60000): Promise<T> {
    const now = Date.now();
    const entry = cache.get(url);

    if (entry && entry.expiry > now) {
        console.log(`[DataCache] Cache hit: ${url}`);
        return entry.data;
    }

    console.log(`[DataCache] Cache miss/expire: ${url}`);
    const response = await fetch(url);
    if (!response.ok) {
        throw new Error(`Failed to fetch ${url}: ${response.statusText}`);
    }

    const data = await response.json();
    cache.set(url, {
        data,
        expiry: now + ttl
    });

    return data;
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
