/**
 * Central configuration for the application.
 * Uses environment variables where available.
 */

const rawBaseUrl = import.meta.env.VITE_BACKEND_URL || ''

export const API_BASE_URL = rawBaseUrl.replace(/\/+$/, '')
