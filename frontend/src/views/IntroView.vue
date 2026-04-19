<script setup lang="ts">
import { useServerStatus } from '@/composables/useServerStatus'
const { status } = useServerStatus()
</script>

<template>
  <div class="intro-view container animate-entry">
    <header class="editorial-header">
      <div class="header-meta">EST. 2024</div>
      <h1 class="main-title">MT NETWORK</h1>
      <p class="tagline">A New Standard for Minecraft Survival</p>
      <div class="header-line"></div>
    </header>

    <div class="editorial-grid">
      <!-- Story Section -->
      <section class="editorial-section story">
        <div class="section-label">THE ORIGIN STORY</div>
        <h2 class="section-title-large">从私人角落到<br><span class="text-gradient">全服公测</span></h2>
        
        <div class="story-columns">
           <p class="lead">MTN的起点极其平凡。2024年10月，三位核心成员——硬件专家 <strong>mc9957</strong>、资深玩家 <strong>wor114514</strong> 以及随后加入的 <strong>txrog</strong>，在深圳的一台Homeserver上种下了这颗种子。</p>
           <p>最初，它仅仅是一个为了三位好友间游玩而设立的小型私人空间。随着 <strong>txrog</strong> 对软件架构的全面优化，服务器成功从封闭的基岩版转型为高性能的Java版架构，并毅然决定面向公众开放。</p>
           <p>在此过程中，铁路建设大师 <strong>HEHAWellgood</strong> 的加入为服务器注入了基建灵魂。经过无数个昼夜的打磨与版本迭代，MTN终于从一个私人存档演变成了现在的形态。</p>
        </div>
      </section>

      <!-- Philosophy / Preface -->
      <section class="glass-card philosophy-card">
        <div class="card-tag">PHILOSOPHY</div>
        <h2>官方群序言</h2>
        <div class="philosophy-content">
          <p>Minecraft 并不只是一个方块游戏，它是开放性与创造力的代名词。</p>
          <p>在 Minecraft 的基石之上，MTMCer 团队创建了 <strong>MT Network (MTN)</strong>。这是一个打破设备隔阂，支持 Java 与 基岩 互通的生存圣地。</p>
          <p>我们追求纯粹。这意味着在这里，你可以体验到最原始的生存乐趣，并深度参与多个现代化城市的宏伟建设。</p>
        </div>
        <div class="philosophy-footer">
          <p>“让我们共同维护这个和谐、有趣的社区，让 MTN 变得更加美好。”</p>
        </div>
      </section>

      <!-- Features (The Blueprint) -->
      <section class="blueprint-section">
        <div class="section-label">TECHNICAL SPECIFICATIONS</div>
        <div class="blueprint-grid">
          <div class="blueprint-item">
            <div class="bp-icon">⚡</div>
            <div class="bp-info">
              <h4>纯公益运营</h4>
              <p>零氪金环境，拒绝任何付费特权或 administrative bias。</p>
            </div>
          </div>
          <div class="blueprint-item">
            <div class="bp-icon">🌲</div>
            <div class="bp-info">
              <h4>原汁原味</h4>
              <p>坚守原版生存核心逻辑，仅进行极致的公平性微调。</p>
            </div>
          </div>
          <div class="blueprint-item">
            <div class="bp-icon">🎮</div>
            <div class="bp-info">
              <h4>全平台互通</h4>
              <p>Java/基岩无缝连接，通过绑定系统实现资产数据共享。</p>
            </div>
          </div>
          <div class="blueprint-item">
            <div class="bp-icon">⚙️</div>
            <div class="bp-info">
              <h4>生电友好</h4>
              <p>不限制高频率红线装置，为生电玩家提供坚实底层支撑。</p>
            </div>
          </div>
        </div>
      </section>

      <!-- Hardware & Community -->
      <div class="community-hardware-wrap">
        <aside class="glass-card hardware-card">
          <h3>NODE STATUS</h3>
          <div class="hw-stats">
             <div class="hw-row">
               <span>LOCATION</span>
               <strong>Shenzhen, CN</strong>
             </div>
             <div class="hw-row">
               <span>NETWORK</span>
               <strong>China Mobile</strong>
             </div>
             <div class="hw-row">
               <span>HARDWARE</span>
               <strong>Dedicated Node</strong>
             </div>
             <div class="hw-row">
               <span>STATUS</span>
               <strong :class="{ online: status?.online }">{{ status?.online ? 'OPERATIONAL' : 'OFFLINE' }}</strong>
             </div>
          </div>
        </aside>

        <section class="glass-card community-card">
          <div class="qq-group-brand">
            <h3>JOIN THE COMMUNTIY</h3>
            <p>与数千位MTN玩家交流心得</p>
          </div>
          <div class="qr-container">
            <div class="qr-glow"></div>
            <img src="/qrcode/qqgroup.jpg" alt="QQ Group" class="qr-code" />
          </div>
          <div class="join-info">
             <span class="group-id">Official Group</span>
             <p>扫码即可开启你的MTN之旅</p>
          </div>
        </section>
      </div>
    </div>
  </div>
</template>

<style scoped>
.intro-view {
  padding-top: 6rem;
  padding-bottom: 10rem;
}

/* Editorial Header */
.editorial-header {
  text-align: center;
  margin-bottom: 8rem;
}

.header-meta {
  font-size: 0.8rem;
  font-weight: 800;
  letter-spacing: 0.4em;
  color: var(--primary);
  margin-bottom: 1rem;
}

.main-title {
  font-family: var(--display);
  font-size: clamp(4rem, 12vw, 10rem);
  font-weight: 800;
  margin-bottom: 0.5rem;
  letter-spacing: -0.05em;
  line-height: 1;
}

.tagline {
  font-size: 1.5rem;
  color: var(--text-muted);
  font-weight: 400;
  letter-spacing: 0.1em;
}

.header-line {
  width: 60px;
  height: 4px;
  background: var(--primary);
  margin: 3rem auto 0;
}

/* Content Grid */
.editorial-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 4rem;
}

@media (min-width: 1100px) {
  .editorial-grid {
    grid-template-columns: 1.6fr 1fr;
  }
}

.section-label {
  font-size: 0.75rem;
  font-weight: 800;
  letter-spacing: 0.3em;
  color: var(--primary);
  margin-bottom: 1.5rem;
  display: flex;
  align-items: center;
  gap: 15px;
}

.section-label::after {
  content: '';
  flex: 1;
  height: 1px;
  background: var(--glass-border);
}

.section-title-large {
  font-size: 3.5rem;
  line-height: 1.1;
  margin-bottom: 2.5rem;
}

/* Story Section */
.story-columns {
  columns: 1;
  column-gap: 3rem;
  color: var(--text-muted);
  font-size: 1.05rem;
  line-height: 1.8;
}

@media (min-width: 768px) {
  .story-columns { columns: 2; }
}

.lead {
  font-size: 1.25rem;
  color: #fff;
  font-weight: 500;
  margin-bottom: 1.5rem;
  break-inside: avoid;
}

/* Philosophy Card */
.philosophy-card {
  height: fit-content;
  border-left: 4px solid var(--primary);
}

.glitch-title {
  font-family: var(--display);
  font-size: clamp(3.5rem, 8vw, 6rem);
  font-weight: 800;
  letter-spacing: -0.05em;
  line-height: 0.9;
  margin-bottom: 1.5rem;
}

.card-tag {
  font-size: 0.7rem;
  font-weight: 800;
  letter-spacing: 2px;
  color: var(--primary);
  margin-bottom: 1rem;
}

.philosophy-content p {
  margin-bottom: 1.5rem;
  font-size: 1.1rem;
  line-height: 1.7;
}

.philosophy-footer {
  margin-top: 2rem;
  padding-top: 2rem;
  border-top: 1px solid var(--glass-border);
  font-style: italic;
  color: var(--text-muted);
}

/* Blueprint Section */
.blueprint-section {
  grid-column: 1 / -1;
  background: rgba(255, 255, 255, 0.02);
  padding: 4rem;
  border-radius: var(--radius-xl);
  border: 1px solid var(--glass-border);
}

.blueprint-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 3rem;
}

.blueprint-item {
  display: flex;
  gap: 1.5rem;
}

.bp-icon {
  font-size: 2rem;
  width: 50px;
  height: 50px;
  background: rgba(255, 255, 255, 0.05);
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: var(--radius-md);
  flex-shrink: 0;
}

.bp-info h4 {
  font-size: 1.1rem;
  margin-bottom: 0.5rem;
  color: var(--primary);
}

.bp-info p {
  font-size: 0.9rem;
  color: var(--text-muted);
}

/* Community & Hardware */
.community-hardware-wrap {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.hardware-card h3 {
  font-size: 1.1rem;
  margin-bottom: 2rem;
  letter-spacing: 2px;
}

.hw-row {
  display: flex;
  justify-content: space-between;
  padding: 15px 0;
  border-bottom: 1px solid var(--glass-border);
}

.hw-row span {
  font-size: 0.8rem;
  color: var(--text-muted);
  font-weight: 700;
}

.hw-row strong.online {
  color: #10B981;
}

.community-card {
  text-align: center;
}

.qq-group-brand h3 {
  font-size: 1.4rem;
  margin-bottom: 0.5rem;
}

.qr-container {
  position: relative;
  width: 220px;
  height: 220px;
  margin: 2rem auto;
  background: #fff;
  padding: 10px;
  border-radius: var(--radius-lg);
  box-shadow: 0 0 50px rgba(59, 130, 246, 0.2);
}

.qr-code {
  width: 100%;
  height: 100%;
}

.qr-glow {
  position: absolute;
  inset: -20px;
  background: var(--primary);
  filter: blur(40px);
  opacity: 0.2;
  z-index: -1;
}

@media (max-width: 768px) {
  .intro-view { padding-top: 4rem; }
  .section-title-large { font-size: 2.5rem; }
  .blueprint-section { padding: 2rem; }
}
</style>
