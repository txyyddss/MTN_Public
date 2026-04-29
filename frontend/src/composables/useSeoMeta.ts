import { watchEffect } from 'vue'
import { useRoute } from 'vue-router'

import { useCurrentLocale, useSiteContent } from '@/content/siteContent'
import { buildSiteUrl, normalizeSitePath } from '@/utils/siteUrl'

type SeoRouteKey = keyof ReturnType<typeof useSiteContent>['value']['seo']['routes']

interface SeoRouteMeta {
  seoKey?: SeoRouteKey
  canonicalPath?: string
  noindex?: boolean
}

function ensureMeta(attribute: 'name' | 'property', key: string): HTMLMetaElement {
  const selector = `meta[${attribute}="${key}"]`
  const existing = document.head.querySelector<HTMLMetaElement>(selector)

  if (existing) {
    return existing
  }

  const meta = document.createElement('meta')
  meta.setAttribute(attribute, key)
  document.head.appendChild(meta)
  return meta
}

function setMeta(attribute: 'name' | 'property', key: string, content: string): void {
  ensureMeta(attribute, key).content = content
}

function setCanonical(url: string | null): void {
  const existing = document.head.querySelector<HTMLLinkElement>('link[rel="canonical"]')

  if (!url) {
    existing?.remove()
    return
  }

  const link = existing ?? document.createElement('link')
  link.rel = 'canonical'
  link.href = url

  if (!existing) {
    document.head.appendChild(link)
  }
}

function getCanonicalPath(meta: SeoRouteMeta, currentPath: string): string {
  return normalizeSitePath(meta.canonicalPath ?? currentPath)
}

export function useSeoMeta(): void {
  const route = useRoute()
  const siteContent = useSiteContent()
  const locale = useCurrentLocale()

  watchEffect(() => {
    const content = siteContent.value.seo
    const meta = route.meta as SeoRouteMeta
    const routeContent = content.routes[meta.seoKey ?? 'home'] ?? content.default
    const title = routeContent.title || content.default.title
    const description = routeContent.description || content.default.description
    const robots = meta.noindex ? 'noindex, nofollow' : 'index, follow'
    const canonicalUrl = buildSiteUrl(getCanonicalPath(meta, route.path))
    const ogLocale = locale.value === 'zh-CN' ? 'zh_CN' : 'en_US'

    document.title = title
    document.documentElement.lang = locale.value

    setMeta('name', 'description', description)
    setMeta('name', 'robots', robots)
    setMeta('property', 'og:site_name', content.siteName)
    setMeta('property', 'og:type', 'website')
    setMeta('property', 'og:title', title)
    setMeta('property', 'og:description', description)
    setMeta('property', 'og:locale', ogLocale)
    setMeta('name', 'twitter:card', 'summary')
    setMeta('name', 'twitter:title', title)
    setMeta('name', 'twitter:description', description)
    setCanonical(canonicalUrl)

    if (canonicalUrl) {
      setMeta('property', 'og:url', canonicalUrl)
    }
  })
}
