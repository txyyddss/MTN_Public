const crypto = require('node:crypto')
const fs = require('node:fs')
const path = require('node:path')
const sharp = require('sharp')

const projectRoot = path.resolve(__dirname, '..')
const sourceDir = path.join(projectRoot, 'src', 'assets', 'gallery-source')
const generatedRoot = path.join(projectRoot, 'public', 'generated', 'gallery')
const thumbDir = path.join(generatedRoot, 'thumbs')
const largeDir = path.join(generatedRoot, 'large')
const manifestDir = path.join(projectRoot, 'src', 'assets', 'generated')
const manifestPath = path.join(manifestDir, 'gallery-manifest.json')
const imagePattern = /\.(png|jpe?g|webp)$/i
const transformVersion = 'gallery-webp-v1-thumb640-large1800'

function assertInside(parent, target) {
  const relative = path.relative(parent, target)

  if (relative.startsWith('..') || path.isAbsolute(relative)) {
    throw new Error(`Refusing to write outside ${parent}: ${target}`)
  }
}

function slugify(value) {
  return value
    .toLowerCase()
    .replace(/[^a-z0-9._-]+/g, '-')
    .replace(/^-+|-+$/g, '')
}

function getHash(filePath) {
  const hash = crypto.createHash('sha256')
  hash.update(fs.readFileSync(filePath))
  hash.update(transformVersion)
  return hash.digest('hex').slice(0, 10)
}

function getLabel(id) {
  return id.replace(/_hdr$/i, '').replace('_', ' ').replace(/\.(?=\d{2}\b)/g, ':')
}

async function generateImage(fileName) {
  const sourcePath = path.join(sourceDir, fileName)
  const extension = path.extname(fileName)
  const id = path.basename(fileName, extension)
  const slug = slugify(id)
  const hash = getHash(sourcePath)
  const thumbFileName = `${slug}-${hash}-thumb.webp`
  const largeFileName = `${slug}-${hash}-large.webp`
  const thumbPath = path.join(thumbDir, thumbFileName)
  const largePath = path.join(largeDir, largeFileName)

  const thumbInfo = await sharp(sourcePath)
    .rotate()
    .resize({ width: 640, withoutEnlargement: true })
    .webp({ quality: 72 })
    .toFile(thumbPath)

  const largeInfo = await sharp(sourcePath)
    .rotate()
    .resize({ width: 1800, withoutEnlargement: true })
    .webp({ quality: 82 })
    .toFile(largePath)

  return {
    id,
    label: getLabel(id),
    thumb: `/generated/gallery/thumbs/${thumbFileName}`,
    large: `/generated/gallery/large/${largeFileName}`,
    thumbWidth: thumbInfo.width,
    thumbHeight: thumbInfo.height,
    width: largeInfo.width,
    height: largeInfo.height
  }
}

async function main() {
  assertInside(projectRoot, generatedRoot)
  assertInside(projectRoot, manifestPath)

  fs.rmSync(generatedRoot, { recursive: true, force: true })
  fs.mkdirSync(thumbDir, { recursive: true })
  fs.mkdirSync(largeDir, { recursive: true })
  fs.mkdirSync(manifestDir, { recursive: true })

  if (!fs.existsSync(sourceDir)) {
    fs.writeFileSync(manifestPath, '[]\n')
    console.warn(`[gallery] Source directory missing: ${sourceDir}`)
    return
  }

  const files = fs.readdirSync(sourceDir)
    .filter((fileName) => imagePattern.test(fileName))
    .sort((left, right) => right.localeCompare(left))

  const manifest = []

  for (const fileName of files) {
    manifest.push(await generateImage(fileName))
  }

  fs.writeFileSync(manifestPath, `${JSON.stringify(manifest, null, 2)}\n`)
  console.log(`[gallery] Generated ${manifest.length} optimized gallery entries.`)
}

main().catch((error) => {
  console.error(error)
  process.exitCode = 1
})
