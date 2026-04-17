import { fetchWithCache } from './dataCache';

const preloadedUrls = new Set<string>();

export const PreloadPriority = {
    URGENT: 20,       // Scripts, critical core assets
    DATA: 15,         // API Data (Leaderboards, specific player details)
    HIGH: 10,       // Player skins
    MEDIUM_HIGH: 7, // Advancements
    MEDIUM: 5,      // Other images (heads/avatars)
    LOW: 3,         // Stats icons
    BACKGROUND: 1   // Full icon pool
} as const;

interface PreloadItem {
    url: string;
    priority: number;
    type: 'image' | 'script' | 'data';
}

const queue: PreloadItem[] = [];
let activeWorkers = 0;
const MAX_WORKERS = 6;

/**
 * Preloads images in the background with priority support.
 */
export function preloadImages(urls: (string | null | undefined)[], priority: number = PreloadPriority.MEDIUM) {
    const validUrls = urls.filter((url): url is string => !!url && !preloadedUrls.has(url));

    if (validUrls.length === 0) return;

    validUrls.forEach(url => {
        preloadedUrls.add(url);
        queue.push({ url, priority, type: 'image' });
    });

    reschedule();
}

/**
 * Preloads API data in the background with priority support.
 */
export function preloadData(urls: string[], priority: number = PreloadPriority.DATA) {
    const validUrls = urls.filter(url => !preloadedUrls.has(url));
    if (validUrls.length === 0) return;

    validUrls.forEach(url => {
        preloadedUrls.add(url);
        queue.push({ url, priority, type: 'data' });
    });

    reschedule();
}

function reschedule() {
    // Sort queue so highest priority items are at the front
    queue.sort((a, b) => b.priority - a.priority);

    if (typeof window !== 'undefined' && (window as any).requestIdleCallback) {
        (window as any).requestIdleCallback(() => processQueue());
    } else {
        setTimeout(() => processQueue(), 1);
    }
}

async function processQueue() {
    while (activeWorkers < MAX_WORKERS && queue.length > 0) {
        const item = queue.shift();
        if (!item) break;

        activeWorkers++;

        if (item.type === 'data') {
            fetchWithCache(item.url)
                .catch(() => { })
                .finally(() => {
                    activeWorkers--;
                    processQueue();
                });
        } else if (item.type === 'image') {
            const img = new Image();
            img.crossOrigin = 'anonymous'; // Ensure preloaded images are stored with CORS headers for WebGL/Canvas use
            img.onload = img.onerror = () => {
                activeWorkers--;
                processQueue();
            };
            img.src = item.url;
        } else {
            // Script case (already handled mostly by preloadScripts, but here for completeness)
            activeWorkers--;
            processQueue();
        }
    }
}

/**
 * Higher-level utility to preload a specific batch and return when done (optional)
 * Note: Uses high priority by default for explicit async waits.
 */
export async function preloadImagesAsync(urls: (string | null | undefined)[], _priority: number = PreloadPriority.HIGH): Promise<void> {
    return new Promise((resolve) => {
        const validUrls = urls.filter((url): url is string => !!url && !preloadedUrls.has(url));
        if (validUrls.length === 0) {
            resolve();
            return;
        }

        let loadedCount = 0;
        validUrls.forEach(url => {
            preloadedUrls.add(url);
            const img = new Image();
            img.onload = img.onerror = () => {
                loadedCount++;
                if (loadedCount === validUrls.length) {
                    resolve();
                }
            };
            img.src = url;
        });
    });
}

/**
 * Preloads JS modules using <link rel="modulepreload">.
 */
export function preloadScripts(urls: string[]) {
    if (typeof document === 'undefined') return;

    urls.forEach(url => {
        if (preloadedUrls.has(url)) return;
        preloadedUrls.add(url);

        const link = document.createElement('link');
        link.rel = 'modulepreload';
        link.href = url;
        document.head.appendChild(link);
    });
}
