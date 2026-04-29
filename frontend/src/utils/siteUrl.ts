export function normalizeSiteOrigin(value: string | undefined): string {
  const trimmed = value?.trim() ?? ''

  if (!trimmed) {
    return ''
  }

  try {
    const url = new URL(trimmed)
    return url.origin
  } catch {
    return ''
  }
}

export function normalizeSitePath(value: string): string {
  const [pathOnly] = value.split(/[?#]/)
  const normalized = pathOnly.startsWith('/') ? pathOnly : `/${pathOnly}`
  return normalized === '/index.html' ? '/' : normalized
}

export const SITE_ORIGIN = normalizeSiteOrigin(import.meta.env.VITE_SITE_URL)

export function buildSiteUrl(path: string): string | null {
  if (!SITE_ORIGIN) {
    return null
  }

  return new URL(normalizeSitePath(path), SITE_ORIGIN).toString()
}
