<script setup lang="ts">
// No complex logic needed for intro
</script>

<template>
  <div class="intro-view container animate-entry">
    <header class="hero">
      <h1 class="main-title">MT Network</h1>
      <p class="tagline">Explore. Create. Build. Together.</p>
    </header>

    <div class="grid-content">
      <!-- History Section -->
      <section class="glass-card section-card history">
        <div class="card-icon">⏳</div>
        <h2>MTN 历史</h2>
        <div class="content">
          <p>MTN于2024年10月由三位核心成员创建，包括硬件负责人 <strong>mc9957</strong> 和骨灰级玩家 <strong>wor114514</strong>。起初，这只是一个小型的私人服务器，仅供三位好友间游玩。</p>
          <p>2025年5月，软件负责人 <strong>txrog</strong> 加入团队，致力于打破局限，将原有的封闭基岩版服务器转型为高性能的Java版服务器，并向公众开放。在此过程中，铁路大亨 <strong>HEHAWellgood</strong> 也随后加入。经过数个版本的精心打磨与迭代，MTN正式开启了它的公测之旅。</p>
        </div>
      </section>

      <!-- Preface Section -->
      <section class="glass-card section-card preface">
        <div class="card-icon">📜</div>
        <h2>官方群序言</h2>
        <div class="content">
          <p>Minecraft 已经成为了开放性与创造力的代名词。我们由衷地感谢 Mojang Studios 为全球玩家带来的快乐。</p>
          <p>在 Minecraft 的基石之上，MTMCer 团队创建了 <strong>MT Network (MTN)</strong>。这是一个支持 Java 版与基岩版互通的生存服务器。在这里，你可以体验最纯粹的生存乐趣，并深度参与多个城市的现代化建设。</p>
          <p>我们希望每位玩家在进入前都能阅读服务器规章。让我们共同维护这个和谐、有趣的社区，让 MTN 变得更加美好！</p>
        </div>
      </section>

      <!-- Stats Section -->
      <section class="glass-card section-card stats">
        <div class="card-icon">⚡</div>
        <h2>服务器特性</h2>
        <ul class="feature-list">
          <li><strong>纯公益</strong> - 零氪金，真正的平等竞技</li>
          <li><strong>原汁原味</strong> - 仅在原版基础上进行微调</li>
          <li><strong>无边界</strong> - 无地图大小限制，自由探索</li>
          <li><strong>生电友好</strong> - 不限制生电玩法</li>
          <li><strong>硬核挑战</strong> - 困难模式 + 死亡惩罚</li>
          <li><strong>全平台</strong> - Java/基岩互通 & 账号绑定</li>
          <li><strong>正版支持</strong> - 支持正版与离线模式登录</li>
          <li><strong>零容忍</strong> - 强力反作弊，守护公平环境</li>
        </ul>
      </section>

      <!-- Hardware/Contact Section -->
      <section class="glass-card section-card hardware">
        <div class="card-icon">🖥️</div>
        <h2>硬件与联系</h2>
        <div class="hardware-info">
          <div class="info-item">
            <span class="label">服务器位置</span>
            <span class="value">中国 · 深圳</span>
          </div>
          <div class="info-item">
            <span class="label">网络环境</span>
            <span class="value">中国移动</span>
          </div>
          <div class="info-item">
            <span class="label">硬件类型</span>
            <span class="value">Homeserver</span>
          </div>
        </div>
        
        <div class="qq-group">
          <h3>加入官方群</h3>
          <div class="qr-wrap">
            <img src="/qrcode/qqgroup.jpg" alt="QQ Group QR Code" class="qr-code" />
          </div>
          <p class="qr-hint">扫码加入社区</p>
        </div>
      </section>
    </div>
  </div>
</template>

<style scoped>
.intro-view {
  padding-top: 4rem;
  padding-bottom: 8rem;
}

.hero {
  text-align: center;
  margin-bottom: 6rem;
}

.main-title {
  font-size: 5rem;
  margin-bottom: 1rem;
  font-weight: 800;
  text-transform: uppercase;
  letter-spacing: -2px;
  background: linear-gradient(135deg, #fff 30%, var(--primary) 100%);
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
  filter: drop-shadow(0 10px 20px rgba(59, 130, 246, 0.2));
}

.tagline {
  font-size: 1.5rem;
  color: var(--text-muted);
  font-family: var(--heading);
  font-weight: 500;
}

.grid-content {
  display: grid;
  grid-template-columns: 1fr;
  gap: 2rem;
}

@media (min-width: 1024px) {
  .grid-content {
    grid-template-columns: 1.5fr 1fr;
  }
}

.section-card {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.section-card h2 {
  font-size: 1.8rem;
  margin-bottom: 1.5rem;
  display: flex;
  align-items: center;
  gap: 12px;
}

.card-icon {
  font-size: 2.5rem;
  margin-bottom: 1rem;
  opacity: 0.8;
}

.content p {
  margin-bottom: 1.2rem;
  color: var(--text-main);
  opacity: 0.9;
  line-height: 1.8;
}

.feature-list {
  list-style: none;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 1rem;
}

.feature-list li {
  padding: 12px;
  background: rgba(255, 255, 255, 0.03);
  border-radius: var(--radius-md);
  border: 1px solid rgba(255, 255, 255, 0.05);
  font-size: 0.95rem;
}

.feature-list li strong {
  color: var(--primary);
  display: block;
  margin-bottom: 4px;
}

.hardware-info {
  margin-bottom: 2rem;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.info-item {
  display: flex;
  justify-content: space-between;
  padding: 12px;
  background: rgba(255, 255, 255, 0.03);
  border-radius: var(--radius-md);
}

.info-item .label {
  color: var(--text-muted);
  font-size: 0.9rem;
}

.info-item .value {
  font-weight: 700;
  color: #fff;
}

.qq-group {
  margin-top: auto;
  text-align: center;
  padding-top: 2rem;
  border-top: 1px solid var(--glass-border);
}

.qr-wrap {
  margin: 1.5rem auto;
  width: 280px;
  height: 280px;
  background: #fff;
  padding: 10px;
  border-radius: var(--radius-md);
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);
}

.qr-code {
  width: 100%;
  height: 100%;
}

.qr-hint {
  font-size: 0.85rem;
  color: var(--text-muted);
}

@media (max-width: 768px) {
  .main-title {
    font-size: 3rem;
  }
}
</style>
