<script setup lang="ts">
import { ref } from 'vue'
import { useServerStatus } from '@/composables/useServerStatus'

const { status, connInfo } = useServerStatus()
const copyFeedback = ref<string | null>(null)
const showAllJava = ref(false)

const copyToClipboard = (text: string, label: string) => {
  navigator.clipboard.writeText(text).then(() => {
    copyFeedback.value = `Copied ${label}!`
    setTimeout(() => {
      copyFeedback.value = null
    }, 2000)
  })
}


const getJavaTotal = () => {
  if (!status.value?.java?.online) return 'Offline'
  const count = status.value.java.players
  return `${count} ${count <= 1 ? 'Player' : 'Players'}`
}

const getBedrockTotal = () => {
  if (!status.value?.bedrock?.online) return 'Offline'
  const count = status.value.bedrock.players
  return `${count} ${count <= 1 ? 'Player' : 'Players'}`
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

const formatDate = (dateStr: string) => {
  if (!dateStr) return 'N/A'
  return new Date(dateStr).toLocaleString()
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
                <small><b>Platform:</b> {{ status.system.platform }}</small>
              </div>
              <div class="stat-row">
                <small><b>CPU:</b> {{ status.system.cpu_model }} ({{ status.system.cpu_percent.toFixed(1) }}%)</small>
              </div>
              <div class="stat-row">
                <small><b>Load:</b> {{ status.system.load_1.toFixed(2) }} / {{ status.system.load_5.toFixed(2) }} / {{ status.system.load_15.toFixed(2) }}</small>
              </div>
              <div class="stat-row mt-2">
                <small><b>RAM:</b> {{ formatMem(status.system.mem_used) }} ({{ status.system.mem_percent.toFixed(0) }}%)</small>
                <small><b>NET:</b> {{ formatBytes(status.system.net_sent_rate + status.system.net_recv_rate) }}</small>
              </div>
            </div>
            <p class="refresh-note">Refreshes every 5s | Last: {{ formatDate(status.updated) }}</p>
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
          <div class="panel-subheader">
            <p class="suggestion">Recommended: Use IPv6 for better stability if supported by your hardware.</p>
            <button v-if="connInfo?.connection" class="toggle-view-btn" @click="showAllJava = !showAllJava">
              {{ showAllJava ? 'Simple View' : 'All Addresses' }}
            </button>
          </div>
          
          <div v-if="connInfo?.connection" class="conn-grid">
            <div class="conn-box">
              <h4>Java Edition</h4>
              <ul v-if="showAllJava">
                <li v-if="connInfo.connection.java_ipv4" @click="copyToClipboard(connInfo.connection.java_ipv4.domain || connInfo.connection.java_ipv4.ip + ':' + connInfo.connection.java_ipv4.port, 'Java IPv4')">
                   <div class="addr-row">
                    <span class="type-badge ipv4">IPv4</span> 
                    <code>{{ connInfo.connection.java_ipv4.domain || connInfo.connection.java_ipv4.ip }}:{{ connInfo.connection.java_ipv4.port }}</code>
                  </div>

                </li>
                <li v-if="connInfo.connection.java_ipv6" @click="copyToClipboard(connInfo.connection.java_ipv6.domain || connInfo.connection.java_ipv6.ip, 'Java IPv6')">
                  <div class="addr-row">
                    <span class="type-badge ipv6">IPv6</span> 
                    <code>{{ connInfo.connection.java_ipv6.domain || connInfo.connection.java_ipv6.ip }}</code>
                  </div>

                </li>
              </ul>
              <ul v-else>
                <li v-if="connInfo.addresses.java_ipv4_srv" @click="copyToClipboard(connInfo.addresses.java_ipv4_srv, 'Java SRV IPv4')">
                  <div class="addr-row">
                    <span class="type-badge srv">SRV v4</span>
                    <code>{{ connInfo.addresses.java_ipv4_srv }}</code>
                  </div>

                </li>
                <li v-if="connInfo.addresses.java_ipv6_srv" @click="copyToClipboard(connInfo.addresses.java_ipv6_srv, 'Java SRV IPv6')">
                  <div class="addr-row">
                    <span class="type-badge srv">SRV v6</span>
                    <code>{{ connInfo.addresses.java_ipv6_srv }}</code>
                  </div>

                </li>
              </ul>
            </div>
            
            <div class="conn-box">
              <h4>Bedrock Edition</h4>
              <ul>
                <li v-if="connInfo.connection.bedrock_ipv6">
                  <div class="addr-row flex-grow">
                    <span class="type-badge ipv6">IPv6</span> 
                    <code @click="copyToClipboard(connInfo.connection.bedrock_ipv6.ip, 'Bedrock IPv6 Addr')">{{ connInfo.connection.bedrock_ipv6.ip }}</code>
                    <code v-if="connInfo.connection.bedrock_ipv6.port" class="port-code" @click="copyToClipboard(connInfo.connection.bedrock_ipv6.port.toString(), 'Bedrock IPv6 Port')">{{ connInfo.connection.bedrock_ipv6.port }}</code>
                  </div>

                </li>
                <li v-if="connInfo.connection.bedrock_ipv4">
                  <div class="addr-row flex-grow">
                    <span class="type-badge ipv4">IPv4</span> 
                    <code @click="copyToClipboard(connInfo.connection.bedrock_ipv4.domain || connInfo.connection.bedrock_ipv4.ip, 'Bedrock Addr')">{{ connInfo.connection.bedrock_ipv4.domain || connInfo.connection.bedrock_ipv4.ip }}</code>
                    <code v-if="connInfo.connection.bedrock_ipv4.port" class="port-code" @click="copyToClipboard(connInfo.connection.bedrock_ipv4.port.toString(), 'Bedrock Port')">{{ connInfo.connection.bedrock_ipv4.port }}</code>
                  </div>

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
  margin-bottom: 4rem;
}

.layout-grid {
  display: grid;
  grid-template-columns: 1fr 1.8fr;
  gap: 2.5rem;
}

@media (max-width: 1024px) {
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

.panel-subheader {
  margin-bottom: 2rem;
}

.toggle-view-btn {
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid var(--glass-border-bright);
  color: #fff;
  padding: 8px 16px;
  border-radius: 8px;
  font-size: 0.8rem;
  font-weight: 600;
  cursor: pointer;
  white-space: nowrap;
  transition: all 0.3s var(--transition-fast);
  font-family: var(--heading);
}
.toggle-view-btn:hover {
  background: rgba(255, 255, 255, 0.08);
  border-color: #fff;
}

h3 {
  display: flex;
  align-items: center;
  gap: 16px;
  margin: 0;
  font-size: 1.6rem;
  font-family: var(--heading);
  letter-spacing: -0.02em;
}

.pulse-dot {
  width: 12px;
  height: 12px;
  background: var(--accent);
  border-radius: 50%;
  box-shadow: 0 0 15px var(--accent);
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0% { transform: scale(0.95); box-shadow: 0 0 0 0 hsla(180, 100%, 50%, 0.4); }
  70% { transform: scale(1); box-shadow: 0 0 0 12px hsla(180, 100%, 50%, 0); }
  100% { transform: scale(0.95); box-shadow: 0 0 0 0 hsla(180, 100%, 50%, 0); }
}

.status-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid var(--glass-border);
  border-radius: 12px;
  margin-bottom: 12px;
}

.badge {
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 0.75rem;
  font-weight: 800;
  letter-spacing: 1px;
  text-transform: uppercase;
  font-family: var(--display);
}
.java-badge { background: rgba(59, 130, 246, 0.1); color: var(--primary); border: 1px solid hsla(210, 100%, 55%, 0.2); }
.bedrock-badge { background: rgba(139, 92, 246, 0.1); color: var(--secondary); border: 1px solid hsla(260, 90%, 65%, 0.2); }

.count {
  font-weight: 700;
  font-size: 1.1rem;
  font-family: var(--heading);
}

.system-stats {
  margin-top: 2rem;
  padding-top: 1.5rem;
  border-top: 1px solid var(--glass-border);
  color: var(--text-muted);
}

.stat-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
}

.stat-row b {
  color: var(--text-dim);
  font-weight: 700;
  letter-spacing: 0.05em;
  text-transform: uppercase;
  font-size: 0.7rem;
}

.refresh-note {
  text-align: right;
  font-size: 0.7rem;
  color: var(--text-dim);
  margin-top: 12px;
  font-weight: 500;
}

.suggestion {
  color: var(--accent);
  background: hsla(180, 100%, 50%, 0.05);
  padding: 16px;
  border-radius: 12px;
  font-size: 0.95rem;
  margin-bottom: 2rem;
  border: 1px solid hsla(180, 100%, 50%, 0.15);
  line-height: 1.5;
}

.conn-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
}
@media (max-width: 600px) {
  .conn-grid { grid-template-columns: 1fr; }
}

.conn-box h4 {
  font-size: 1.2rem;
  margin-bottom: 1.25rem;
  color: #fff;
  font-family: var(--heading);
}

.conn-box li {
  padding: 12px 16px !important;
  background: rgba(255, 255, 255, 0.02) !important;
  border: 1px solid var(--glass-border) !important;
  transition: all 0.3s var(--transition-fast) !important;
}
.conn-box li:hover {
  background: rgba(255, 255, 255, 0.05) !important;
  border-color: var(--glass-border-bright) !important;
  transform: translateX(4px);
}

.type-badge {
  font-family: var(--display);
  font-weight: 800;
}

code {
  background: rgba(0,0,0,0.4);
  font-family: 'JetBrains Mono', monospace;
  font-size: 0.8rem;
}

.copy-toast {
  font-family: var(--heading);
  font-weight: 700;
}
</style>
