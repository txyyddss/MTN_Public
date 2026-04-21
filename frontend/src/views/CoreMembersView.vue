<script setup lang="ts">
import { coreMembers } from '@/content/coreMembers'
import { siteContent } from '@/content/siteContent'
</script>

<template>
  <div class="core-members-view container page-shell">
    <header class="page-header animate-entry">
      <span class="page-kicker">{{ siteContent.coreMembers.eyebrow }}</span>
      <h1>{{ siteContent.coreMembers.title }}</h1>
      <p class="page-lede">{{ siteContent.coreMembers.body }}</p>
    </header>

    <section class="member-grid">
      <article
        v-for="(member, index) in coreMembers"
        :key="member.name"
        class="glass-card member-card animate-entry"
        :style="{ animationDelay: `${index * 0.05}s` }"
      >
        <div class="member-head">
          <div class="member-copy">
            <span class="hud-kicker">{{ siteContent.coreMembers.roleLabel }}</span>
            <h2 class="minecraft-font">{{ member.name }}</h2>
          </div>
          <span class="badge-pill"><strong>{{ member.role }}</strong></span>
        </div>

        <div class="member-section">
          <span class="member-label">{{ siteContent.coreMembers.nicknamesLabel }}</span>
          <div class="member-chip-row">
            <span v-if="member.nicknames.length === 0" class="meta-chip">None</span>
            <span v-for="nickname in member.nicknames" :key="nickname" class="meta-chip">
              {{ nickname }}
            </span>
          </div>
        </div>

        <div class="member-section">
          <span class="member-label">{{ siteContent.coreMembers.introLabel }}</span>
          <p class="member-introduction">{{ member.introduction }}</p>
        </div>
      </article>
    </section>
  </div>
</template>

<style scoped>
.core-members-view {
  display: grid;
  gap: 1rem;
}

.member-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 0.95rem;
}

.member-card,
.member-head,
.member-copy,
.member-section {
  display: grid;
  gap: 0.55rem;
}

.member-card {
  gap: 1rem;
}

.member-card h2 {
  font-size: clamp(2rem, 3.8vw, 2.8rem);
}

.member-label {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.72rem;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.member-chip-row {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.meta-chip {
  padding: 0.45rem 0.7rem;
  border-radius: 999px;
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.68rem;
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.member-introduction {
  color: var(--text-muted);
}

@media (max-width: 980px) {
  .member-grid {
    grid-template-columns: 1fr;
  }
}
</style>
