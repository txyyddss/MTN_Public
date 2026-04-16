const fs = require('fs');
const path = require('path');

const iconsDir = path.join(__dirname, 'public', 'mc_icons');
const outputDir = path.join(__dirname, 'src', 'assets');
const outputFile = path.join(outputDir, 'icon_map.json');

const iconMap = {};

function walk(dir) {
    const files = fs.readdirSync(dir);
    for (const file of files) {
        const fullPath = path.join(dir, file);
        if (fs.statSync(fullPath).isDirectory()) {
            walk(fullPath);
        } else {
            if (!file.endsWith('.png')) continue;

            // Get relative path from public folder
            const relPath = '/' + path.relative(path.join(__dirname, 'public'), fullPath).replace(/\\/g, '/');

            // Extract base ID
            // Examples: 
            // oak_log^32.png -> oak_log
            // zombie.png -> zombie
            // silverfish^16.png -> silverfish
            let id = file.replace('.png', '');
            id = id.split('^')[0]; // Remove ^32, ^16 etc
            id = id.split('$')[0]; // Remove animation suffixes if any

            // Store mapping
            // If multiple files map to same ID, we might want to prioritize ^32 or just use the first found
            // For now, let's keep the first one found, but prefer ^32 if we see it later
            if (!iconMap[id] || relPath.includes('^32')) {
                iconMap[id] = relPath;
            }
        }
    }
}

console.log('Scanning icons in:', iconsDir);
if (fs.existsSync(iconsDir)) {
    walk(iconsDir);

    if (!fs.existsSync(outputDir)) {
        fs.mkdirSync(outputDir, { recursive: true });
    }

    fs.writeFileSync(outputFile, JSON.stringify(iconMap, null, 2));
    console.log(`Generated icon map with ${Object.keys(iconMap).length} icons at ${outputFile}`);
} else {
    console.error('Icons directory not found:', iconsDir);
}
