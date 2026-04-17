/**
 * Ping Utility for measuring TCP and UDP (approximation) latency in the browser.
 */

/**
 * Measures TCP latency by attempting a connection to the host/port.
 * Uses a 'no-cors' fetch which will likely fail, but the timing is a good proxy for RTT.
 */
export async function tcpPing(host: string, port: number, timeoutMs: number = 2000): Promise<number | null> {
    const start = performance.now();
    const controller = new AbortController();
    const timeout = setTimeout(() => controller.abort(), timeoutMs);

    try {
        // We use http scheme but any port. The browser will likely block common ports or 
        // fail with a connection error. The timing of this error is what we measure.
        await fetch(`http://${host}:${port}`, {
            mode: 'no-cors',
            signal: controller.signal,
            cache: 'no-cache',
            referrerPolicy: 'no-referrer'
        });
    } catch (error: any) {
        if (error.name === 'AbortError') {
            return null; // Timeout
        }
    } finally {
        clearTimeout(timeout);
    }

    const latency = performance.now() - start;
    return Math.round(latency);
}

/**
 * Measures UDP latency. Since browsers cannot do raw UDP, this is an approximation.
 * For Minecraft Bedrock (UDP), we measure connectivity to the host.
 */
export async function udpPing(host: string, port: number, timeoutMs: number = 2000): Promise<number | null> {
    // Since we can't do real UDP, we perform a TCP ping to the same host 
    // on a likely-open port, or just use the TCP ping logic if port is provided.
    // Many cloud hosts allow some basic TCP connectivity even if it's a UDP service.
    return tcpPing(host, port, timeoutMs);
}
