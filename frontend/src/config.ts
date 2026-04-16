/**
 * Central configuration for the application.
 * Uses environment variables where available.
 */

// If VITE_BACKEND_URL is set, use it. 
// Otherwise, use an empty string which implies the origin (relative paths).
// Vite exposes env vars starting with VITE_ to the client.
export const API_BASE_URL = import.meta.env.VITE_BACKEND_URL || '';

console.log('API Base URL:', API_BASE_URL || '(current origin)');
