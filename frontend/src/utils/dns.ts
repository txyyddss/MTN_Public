/**
 * DNS Utility for SRV resolution using multiple DNS-over-HTTPS (DoH) providers.
 * Races multiple providers to get the fastest response.
 */

export interface SRVRecord {
    target: string;
    port: number;
    priority: number;
    weight: number;
}

const DOH_PROVIDERS = [
    { name: 'Cloudflare', url: (domain: string) => `https://cloudflare-dns.com/dns-query?name=${encodeURIComponent(domain)}&type=SRV` },
    { name: 'AliDNS', url: (domain: string) => `https://dns.alidns.com/resolve?name=${encodeURIComponent(domain)}&type=SRV` },
    { name: 'Tencent', url: (domain: string) => `https://doh.pub/dns-query?name=${encodeURIComponent(domain)}&type=SRV` },
    { name: 'OpenDNS', url: (domain: string) => `https://doh.opendns.com/dns-query?name=${encodeURIComponent(domain)}&type=SRV` }
];

async function fetchFromProvider(provider: typeof DOH_PROVIDERS[0], domain: string): Promise<any> {
    const response = await fetch(provider.url(domain), {
        headers: {
            'Accept': 'application/dns-json'
        }
    });

    if (!response.ok) {
        throw new Error(`Provider ${provider.name} failed: ${response.statusText}`);
    }

    const data = await response.json();
    if (!data.Answer) {
        throw new Error(`Provider ${provider.name} returned no answer`);
    }
    return data;
}

export async function resolveSRV(domain: string): Promise<SRVRecord[]> {
    try {
        // Race multiple providers
        const fastestResponse = await Promise.any(
            DOH_PROVIDERS.map(p => fetchFromProvider(p, domain))
        );

        return fastestResponse.Answer.map((ans: any) => {
            // SRV data format: "priority weight port target"
            const [priority, weight, port, target] = ans.data.split(' ');
            return {
                priority: parseInt(priority),
                weight: parseInt(weight),
                port: parseInt(port),
                target: target.replace(/\.$/, '')
            };
        }).sort((a: any, b: any) => a.priority - b.priority || b.weight - a.weight);

    } catch (error) {
        console.error('All SRV Resolution providers failed or returned error:', error);
        return [];
    }
}
