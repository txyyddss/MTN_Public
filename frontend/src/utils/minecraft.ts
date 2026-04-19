export function getSkinUrl(name: string, type?: string) {
    if (!name) return 'https://mineskin.eu/skin/Steve'
    let cleanName = name
    if (type === 'Bedrock' || cleanName.startsWith('.')) {
        cleanName = cleanName.replace(/^\./, '').replace(/^BE_/, '')
    }
    return `https://mineskin.eu/skin/${cleanName}`
}

export function getAvatarUrl(name: string, type?: string) {
    if (!name) return 'https://mineskin.eu/helm/Steve'
    let cleanName = name
    if (type === 'Bedrock' || cleanName.startsWith('.')) {
        cleanName = cleanName.replace(/^\./, '').replace(/^BE_/, '')
    }
    return `https://mineskin.eu/helm/${cleanName}`
}
