const preloadedUrls = new Set<string>();

/**
 * Preloads images in the background to warm up the browser cache.
 * @param urls List of image URLs to preload
 * @param concurrency Maximum number of simultaneous requests
 */
export function preloadImages(urls: (string | null | undefined)[], concurrency: number = 6) {
    const validUrls = urls.filter((url): url is string => !!url && !preloadedUrls.has(url));

    if (validUrls.length === 0) return;

    // Add all to set immediately to prevent duplicate triggers from other calls
    validUrls.forEach(url => preloadedUrls.add(url));

    let index = 0;

    const loadNext = () => {
        if (index >= validUrls.length) return;

        const url = validUrls[index++];
        const img = new Image();

        // Using a promise-like structure but without async for simplicity in the loop
        img.onload = img.onerror = () => {
            // Clean up to avoid memory leaks if used in a long-running app
            // although standard <img> tags are usually fine.
            // (img as any).onload = (img as any).onerror = null;
            loadNext();
        };

        img.src = url;
    };

    const initialConcurrency = Math.min(concurrency, validUrls.length);
    for (let i = 0; i < initialConcurrency; i++) {
        loadNext();
    }
}

/**
 * Higher-level utility to preload a specific batch and return when done (optional)
 */
export async function preloadImagesAsync(urls: (string | null | undefined)[]): Promise<void> {
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
