const fs = require('node:fs')
const path = require('node:path')

const projectRoot = path.resolve(__dirname, '..')
const publicDir = path.join(projectRoot, 'public')
const seoRoutes = ['/', '/players']

function readEnvFile(filePath) {
  if (!fs.existsSync(filePath)) {
    return
  }

  const lines = fs.readFileSync(filePath, 'utf8').split(/\r?\n/)

  for (const line of lines) {
    const trimmed = line.trim()

    if (!trimmed || trimmed.startsWith('#')) {
      continue
    }

    const match = trimmed.match(/^([A-Za-z_][A-Za-z0-9_]*)=(.*)$/)

    if (!match) {
      continue
    }

    const [, key, rawValue] = match

    if (process.env[key]) {
      continue
    }

    process.env[key] = rawValue.replace(/^['"]|['"]$/g, '').trim()
  }
}

function normalizeOrigin(value) {
  const trimmed = (value || '').trim()

  if (!trimmed) {
    return ''
  }

  try {
    return new URL(trimmed).origin
  } catch {
    return ''
  }
}

function escapeXml(value) {
  return value
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
    .replace(/"/g, '&quot;')
    .replace(/'/g, '&apos;')
}

function buildSitemap(siteOrigin) {
  const today = new Date().toISOString().slice(0, 10)
  const urls = siteOrigin
    ? seoRoutes.map((route) => {
        const loc = new URL(route, siteOrigin).toString()

        return [
          '  <url>',
          `    <loc>${escapeXml(loc)}</loc>`,
          `    <lastmod>${today}</lastmod>`,
          '    <changefreq>weekly</changefreq>',
          route === '/' ? '    <priority>1.0</priority>' : '    <priority>0.7</priority>',
          '  </url>'
        ].join('\n')
      })
    : []

  return [
    '<?xml version="1.0" encoding="UTF-8"?>',
    '<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">',
    ...urls,
    '</urlset>',
    ''
  ].join('\n')
}

readEnvFile(path.join(projectRoot, '.env.local'))
readEnvFile(path.join(projectRoot, '.env'))

const siteOrigin = normalizeOrigin(process.env.VITE_SITE_URL)

if (!siteOrigin) {
  console.warn('[seo] VITE_SITE_URL is not set; sitemap will contain no absolute URLs.')
}

fs.mkdirSync(publicDir, { recursive: true })

const robotsLines = ['User-agent: *', 'Allow: /', 'Disallow: /whitelist', 'Disallow: /wiki', '']

if (siteOrigin) {
  robotsLines.push(`Sitemap: ${new URL('/sitemap.xml', siteOrigin).toString()}`, '')
}

fs.writeFileSync(path.join(publicDir, 'robots.txt'), robotsLines.join('\n'))
fs.writeFileSync(path.join(publicDir, 'sitemap.xml'), buildSitemap(siteOrigin))

console.log(`[seo] Generated robots.txt and sitemap.xml${siteOrigin ? ` for ${siteOrigin}` : ''}.`)
