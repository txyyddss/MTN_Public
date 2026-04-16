const preloadedUrls = new Set<string>();

export const PreloadPriority = {
    URGENT: 20,       // Scripts, critical core assets
    HIGH: 10,       // Player skins
    MEDIUM_HIGH: 7, // Advancements
    MEDIUM: 5,      // Other images (heads/avatars)
    LOW: 3,         // Stats icons
    BACKGROUND: 1   // Full icon pool
} as const;

interface PreloadItem {
    url: string;
    priority: number;
    type?: 'image' | 'script';
}

const queue: PreloadItem[] = [];
let activeWorkers = 0;
const MAX_WORKERS = 6;

/**
 * Preloads images in the background with priority support.
 * @param urls List of image URLs to preload
 * @param priority Priority of this batch (higher = sooner)
 */
export function preloadImages(urls: (string | null | undefined)[], priority: number = PreloadPriority.MEDIUM) {
    const validUrls = urls.filter((url): url is string => !!url && !preloadedUrls.has(url));

    if (validUrls.length === 0) return;

    // Add items to queue and sort by priority descending
    validUrls.forEach(url => {
        preloadedUrls.add(url);
        queue.push({ url, priority });
    });

    // Sort queue so highest priority items are at the front
    queue.sort((a, b) => b.priority - a.priority);

    // Use requestIdleCallback to start processing if we aren't already processing
    // This helps ensure preloading doesn't interfere with main thread critical tasks
    if (typeof window !== 'undefined' && (window as any).requestIdleCallback) {
        (window as any).requestIdleCallback(() => processQueue());
    } else {
        setTimeout(() => processQueue(), 1);
    }
}

function processQueue() {
    while (activeWorkers < MAX_WORKERS && queue.length > 0) {
        const item = queue.shift();
        if (!item) break;

        activeWorkers++;
        const img = new Image();

        img.onload = img.onerror = () => {
            activeWorkers--;
            processQueue();
        };

        img.src = item.url;
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
 * These are injected immediately as they are typically top-priority for navigation.
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
