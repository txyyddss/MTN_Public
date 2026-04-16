<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { API_BASE_URL } from '@/config'

const status = ref<any>(null)
const connInfo = ref<any>(null)
const copyFeedback = ref<string | null>(null)
let pollInterval: any = null

const fetchStatus = async () => {
  try {
    const res = await fetch(`${API_BASE_URL}/api/status`)
    if (res.ok) {
      status.value = await res.json()
    }
  } catch (err) {
    console.error('Failed to fetch status', err)
  }
}

const fetchConnection = async () => {
  try {
    const res = await fetch(`${API_BASE_URL}/api/connection`)
    if (res.ok) {
      connInfo.value = await res.json()
    }
  } catch (err) {
    console.error('Failed to fetch connection info', err)
  }
}

const copyToClipboard = (text: string, label: string) => {
  navigator.clipboard.writeText(text).then(() => {
    copyFeedback.value = `Copied ${label}!`
    setTimeout(() => {
      copyFeedback.value = null
    }, 2000)
  })
}

onMounted(() => {
  fetchStatus()
  fetchConnection()
  pollInterval = setInterval(fetchStatus, 5000)
})

onUnmounted(() => {
  if (pollInterval) clearInterval(pollInterval)
})

const getJavaTotal = () => {
  if (!status.value?.java?.online) return 'Offline'
  return `${status.value.java.players} Players`
}

const getBedrockTotal = () => {
  if (!status.value?.bedrock?.online) return 'Offline'
  return `${status.value.bedrock.players} Players`
}

const formatBytes = (bytes: number) => {
  if (bytes === 0) return '0 B/s'
  const k = 1024
  const sizes = ['B/s', 'KB/s', 'MB/s', 'GB/s']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(1)) + ' ' + sizes[i]
}

const formatMem = (bytes: number) => {
  return (bytes / (1024 * 1024 * 1024)).toFixed(1) + ' GB'
}
</script>

<template>
  <section class="server-status-section">
    <div class="container animate-entry delay-300">
      
      <div class="grid layout-grid">
        <!-- Live Status Panel -->
        <div class="glass-card status-panel">
          <h3><span class="pulse-dot"></span> Live Server Status</h3>
          <div v-if="status" class="status-content">
            <div class="status-item">
              <span class="badge java-badge">Java Edition</span>
              <span class="count">{{ getJavaTotal() }}</span>
            </div>
            <div class="status-item">
              <span class="badge bedrock-badge">Bedrock Edition</span>
              <span class="count">{{ getBedrockTotal() }}</span>
            </div>
            
            <div class="system-stats" v-if="status.system">
              <div class="stat-row">
                <small><b>CPU:</b> {{ status.system.cpu_percent.toFixed(1) }}%</small>
                <small><b>RAM:</b> {{ formatMem(status.system.mem_used) }} ({{ status.system.mem_percent.toFixed(0) }}%)</small>
              </div>
              <div class="stat-row">
                <small><b>DISK:</b> {{ formatBytes(status.system.disk_read_rate + status.system.disk_write_rate) }}</small>
                <small><b>NET:</b> {{ formatBytes(status.system.net_sent_rate + status.system.net_recv_rate) }}</small>
              </div>
            </div>
            <p class="refresh-note">Refreshes every 5s</p>
          </div>
          <div v-else class="loading-state">Loading status...</div>
        </div>

        <!-- Connection Info Panel -->
        <div class="glass-card connection-panel">
          <div class="panel-header">
            <h3><span>⬡</span> How to Connect</h3>
            <transition name="fade">
              <span v-if="copyFeedback" class="copy-toast">{{ copyFeedback }}</span>
            </transition>
          </div>
          <p class="suggestion">Recommended: Use IPv6 for better stability if supported by your hardware.</p>
          
          <div v-if="connInfo?.connection" class="conn-grid">
            <div class="conn-box">
              <h4>Java Edition</h4>
              <ul>
                <li v-if="connInfo.connection.java_ipv6" @click="copyToClipboard(connInfo.connection.java_ipv6.domain || connInfo.connection.java_ipv6.ip, 'Java IPv6')">
                  <div class="addr-row">
                    <span class="type-badge ipv6">IPv6</span> 
                    <code>{{ connInfo.connection.java_ipv6.domain || connInfo.connection.java_ipv6.ip }}</code>
                  </div>
                  <span v-if="status?.connections?.java_ipv6" :class="['latency-dot', status.connections.java_ipv6.online ? 'online' : 'offline']">
                    {{ status.connections.java_ipv6.latency_ms || '...' }}ms
                  </span>
                </li>
                <li v-if="connInfo.connection.java_ipv4" @click="copyToClipboard(connInfo.connection.java_ipv4.domain || connInfo.connection.java_ipv4.ip + ':' + connInfo.connection.java_ipv4.port, 'Java IPv4')">
                   <div class="addr-row">
                    <span class="type-badge ipv4">IPv4</span> 
                    <code>{{ connInfo.connection.java_ipv4.domain || connInfo.connection.java_ipv4.ip }}:{{ connInfo.connection.java_ipv4.port }}</code>
                  </div>
                  <span v-if="status?.connections?.java_ipv4" :class="['latency-dot', status.connections.java_ipv4.online ? 'online' : 'offline']">
                    {{ status.connections.java_ipv4.latency_ms || '...' }}ms
                  </span>
                </li>
                <li v-if="connInfo.connection.java_srv" @click="copyToClipboard(connInfo.connection.java_srv.domain, 'Java SRV')">
                  <div class="addr-row">
                    <span class="type-badge srv">SRV</span>
                    <code>{{ connInfo.connection.java_srv.domain }}</code>
                  </div>
                  <span v-if="status?.connections?.java_ipv4_srv" :class="['latency-dot', status.connections.java_ipv4_srv.online ? 'online' : 'offline']">
                    {{ status.connections.java_ipv4_srv.latency_ms || '...' }}ms
                  </span>
                </li>
              </ul>
            </div>
            
            <div class="conn-box">
              <h4>Bedrock Edition</h4>
              <ul>
                <li v-if="connInfo.connection.bedrock_ipv6" @click="copyToClipboard(connInfo.connection.bedrock_ipv6.domain || connInfo.connection.bedrock_ipv6.ip, 'Bedrock IPv6')">
                  <div class="addr-row">
                    <span class="type-badge ipv6">IPv6</span> 
                    <code>{{ connInfo.connection.bedrock_ipv6.domain || connInfo.connection.bedrock_ipv6.ip }}</code>
                  </div>
                   <span v-if="status?.connections?.bedrock_ipv6" :class="['latency-dot', status.connections.bedrock_ipv6.online ? 'online' : 'offline']">
                    {{ status.connections.bedrock_ipv6.latency_ms || '...' }}ms
                  </span>
                </li>
                <li v-if="connInfo.connection.bedrock_ipv4" @click="copyToClipboard((connInfo.connection.bedrock_ipv4.domain || connInfo.connection.bedrock_ipv4.ip) + ' ' + connInfo.connection.bedrock_ipv4.port, 'Bedrock Addr + Port')">
                  <div class="addr-row">
                    <span class="type-badge ipv4">IPv4</span> 
                    <code>{{ connInfo.connection.bedrock_ipv4.domain || connInfo.connection.bedrock_ipv4.ip }}</code>
                    <code class="port-code">{{ connInfo.connection.bedrock_ipv4.port }}</code>
                  </div>
                  <span v-if="status?.connections?.bedrock_ipv4" :class="['latency-dot', status.connections.bedrock_ipv4.online ? 'online' : 'offline']">
                    {{ status.connections.bedrock_ipv4.latency_ms || '...' }}ms
                  </span>
                </li>
              </ul>
            </div>
          </div>
          <div v-else class="loading-state">Loading addresses...</div>
        </div>
      </div>

    </div>
  </section>
</template>

<style scoped>
.server-status-section {
  position: relative;
  z-index: 10;
  margin-top: -2rem;
  margin-bottom: 2rem;
}

.layout-grid {
  display: grid;
  grid-template-columns: 1fr 2fr;
  gap: 2rem;
}

@media (max-width: 900px) {
  .layout-grid {
    grid-template-columns: 1fr;
  }
}

.panel-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

h3 {
  display: flex;
  align-items: center;
  gap: 12px;
  margin: 0;
  font-size: 1.5rem;
}

.pulse-dot {
  width: 12px;
  height: 12px;
  background: var(--primary);
  border-radius: 50%;
  box-shadow: 0 0 10px var(--primary);
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0% { transform: scale(0.95); box-shadow: 0 0 0 0 rgba(59, 130, 246, 0.7); }
  70% { transform: scale(1); box-shadow: 0 0 0 10px rgba(59, 130, 246, 0); }
  100% { transform: scale(0.95); box-shadow: 0 0 0 0 rgba(59, 130, 246, 0); }
}

.status-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: rgba(255, 255, 255, 0.03);
  border-radius: 8px;
  margin-bottom: 12px;
}

.badge {
  padding: 4px 10px;
  border-radius: 4px;
  font-size: 0.8rem;
  font-weight: 700;
  letter-spacing: 0.5px;
}
.java-badge { background: rgba(59, 130, 246, 0.2); color: #93c5fd; }
.bedrock-badge { background: rgba(79, 70, 229, 0.2); color: #c7d2fe; }

.count {
  font-weight: 700;
  font-size: 1.1rem;
}

.system-stats {
  margin-top: 16px;
  padding-top: 16px;
  border-top: 1px solid var(--glass-border);
  color: var(--text-muted);
}

.stat-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 4px;
}

.refresh-note {
  text-align: right;
  font-size: 0.75rem;
  color: var(--text-muted);
  margin-top: 8px;
  font-style: italic;
}

.loading-state {
  color: var(--text-muted);
  font-style: italic;
  padding: 2rem 0;
  text-align: center;
}

.suggestion {
  color: #93c5fd;
  background: rgba(59, 130, 246, 0.1);
  padding: 12px;
  border-radius: 8px;
  font-size: 0.9rem;
  margin-bottom: 1.5rem;
  border: 1px solid rgba(59, 130, 246, 0.2);
}

.conn-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1.5rem;
}
@media (max-width: 600px) {
  .conn-grid { grid-template-columns: 1fr; }
}

.conn-box h4 {
  font-size: 1.1rem;
  margin-bottom: 1rem;
  color: #fff;
}

.conn-box ul {
  list-style: none;
  display: flex;
  flex-direction: column;
  gap: 12px;
}
.conn-box li {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 12px;
  padding: 8px;
  background: rgba(255, 255, 255, 0.02);
  border-radius: 6px;
  cursor: pointer;
  transition: background 0.2s;
}
.conn-box li:hover {
  background: rgba(255, 255, 255, 0.05);
}

.addr-row {
  display: flex;
  align-items: center;
  gap: 8px;
  overflow: hidden;
}

.type-badge {
  font-size: 0.65rem;
  padding: 2px 6px;
  border-radius: 4px;
  font-weight: bold;
  flex-shrink: 0;
}
.ipv6 { background: rgba(139, 92, 246, 0.2); color: #c4b5fd; }
.ipv4 { background: rgba(245, 158, 11, 0.2); color: #fcd34d; }
.srv  { background: rgba(16, 185, 129, 0.2); color: #6ee7b7; }

code {
  background: rgba(0,0,0,0.3);
  padding: 2px 6px;
  border-radius: 4px;
  color: var(--text-main);
  font-size: 0.85rem;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.port-code {
  color: var(--primary);
  border: 1px solid rgba(59, 130, 246, 0.3);
}

.latency-dot {
  font-size: 0.7rem;
  padding: 2px 6px;
  border-radius: 10px;
  background: rgba(0,0,0,0.4);
}
.latency-dot.online { border-left: 3px solid #10b981; color: #6ee7b7; }
.latency-dot.offline { border-left: 3px solid #ef4444; color: #fca5a5; }

.copy-toast {
  font-size: 0.8rem;
  color: var(--primary);
  background: rgba(59, 130, 246, 0.1);
  padding: 4px 12px;
  border-radius: 20px;
  border: 1px solid var(--primary);
}

.fade-enter-active, .fade-leave-active { transition: opacity 0.5s; }
.fade-enter-from, .fade-leave-to { opacity: 0; }
</style>
