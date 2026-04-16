<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { API_BASE_URL } from '@/config'

const status = ref<any>(null)
const connInfo = ref<any>(null)
let pollInterval: any = null

const fetchStatus = async () => {
  try {
    // Note: Prepend API_BASE_URL to support environment variable configuration
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
  return `${status.value.java.players} / ${status.value.java.max_players}`
}

const getBedrockTotal = () => {
  if (!status.value?.bedrock?.online) return 'Offline'
  return `${status.value.bedrock.players} / ${status.value.bedrock.max_players}`
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
              <small><b>CPU:</b> {{ status.system.cpu_percent.toFixed(1) }}%</small>
              <small><b>RAM:</b> {{ status.system.mem_used_mb }} / {{ status.system.mem_total_mb }} MB</small>
            </div>
            <p class="refresh-note">Refreshes every 5s</p>
          </div>
          <div v-else class="loading-state">Loading status...</div>
        </div>

        <!-- Connection Info Panel -->
        <div class="glass-card connection-panel">
          <h3>🔗 How to Connect</h3>
          <p class="suggestion"><strong>Tip:</strong> Choose IPv6 addresses for the best connection latency and stability if supported by your network.</p>
          
          <div v-if="connInfo?.connection" class="conn-grid">
            <div class="conn-box">
              <h4>Java Edition</h4>
              <ul>
                <li v-if="connInfo.connection.java_ipv6">
                  <span class="type-badge ipv6">IPv6</span> 
                  <code>{{ connInfo.connection.java_ipv6.domain || connInfo.connection.java_ipv6.ip }}<span v-if="connInfo.connection.java_ipv6.port">...</span></code>
                </li>
                <li v-if="connInfo.connection.java_ipv4">
                  <span class="type-badge ipv4">IPv4</span> 
                  <code>{{ connInfo.connection.java_ipv4.domain || connInfo.connection.java_ipv4.ip }}:{{ connInfo.connection.java_ipv4.port }}</code>
                </li>
                <li v-if="connInfo.connection.java_srv">
                  <span class="type-badge srv">SRV</span>
                  <code>{{ connInfo.connection.java_srv.domain }}</code>
                </li>
              </ul>
            </div>
            
            <div class="conn-box">
              <h4>Bedrock Edition</h4>
              <ul>
                <li v-if="connInfo.connection.bedrock_ipv6">
                  <span class="type-badge ipv6">IPv6</span> 
                  <code>{{ connInfo.connection.bedrock_ipv6.domain || connInfo.connection.bedrock_ipv6.ip }}:19132</code>
                </li>
                <li v-if="connInfo.connection.bedrock_ipv4">
                  <span class="type-badge ipv4">IPv4</span> 
                  <code>{{ connInfo.connection.bedrock_ipv4.domain || connInfo.connection.bedrock_ipv4.ip }}:{{ connInfo.connection.bedrock_ipv4.port }}</code>
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

h3 {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 1.5rem;
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
  0% { transform: scale(0.95); box-shadow: 0 0 0 0 rgba(0, 230, 118, 0.7); }
  70% { transform: scale(1); box-shadow: 0 0 0 10px rgba(0, 230, 118, 0); }
  100% { transform: scale(0.95); box-shadow: 0 0 0 0 rgba(0, 230, 118, 0); }
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
.java-badge { background: rgba(244, 67, 54, 0.2); color: #ff8a80; }
.bedrock-badge { background: rgba(33, 150, 243, 0.2); color: #82b1ff; }

.count {
  font-weight: 700;
  font-size: 1.1rem;
}

.system-stats {
  display: flex;
  justify-content: space-between;
  margin-top: 16px;
  padding-top: 16px;
  border-top: 1px solid var(--glass-border);
  color: var(--text-muted);
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
  color: var(--primary);
  background: rgba(0, 230, 118, 0.1);
  padding: 12px;
  border-radius: 8px;
  font-size: 0.9rem;
  margin-bottom: 1.5rem;
  border: 1px solid rgba(0, 230, 118, 0.3);
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
  align-items: center;
  gap: 12px;
  font-family: var(--mono);
}

.type-badge {
  font-size: 0.75rem;
  padding: 4px 8px;
  border-radius: 4px;
  font-weight: bold;
}
.ipv6 { background: rgba(156, 39, 176, 0.2); color: #e1bee7; }
.ipv4 { background: rgba(255, 152, 0, 0.2); color: #ffe082; }
.srv  { background: rgba(76, 175, 80, 0.2); color: #a5d6a7; }

code {
  background: rgba(0,0,0,0.3);
  padding: 4px 8px;
  border-radius: 4px;
  color: var(--text-main);
}
</style>
