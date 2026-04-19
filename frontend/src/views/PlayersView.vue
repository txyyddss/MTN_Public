<script setup lang="ts">
import { computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'

import { API_BASE_URL } from '@/config'
import { siteContent } from '@/content/siteContent'
import { usePlayers } from '@/composables/usePlayers'
import { useServerStatusStore } from '@/stores/serverStatus'
import type { RandomPlayerResponse } from '@/types/api'
import { getAvatarUrl } from '@/utils/minecraft'

const router = useRouter()
const statusStore = useServerStatusStore()
const { status } = storeToRefs(statusStore)

const onlinePlayers = computed(() => status.value?.online_players ?? [])
const { players, count, activeDays, searchQuery, showAll, loading, sortedPlayers, fetchPlayers, startAutoRefresh, stopAutoRefresh } =
  usePlayers(onlinePlayers)

async function handleRandom(): Promise<void> {
  try {
    const response = await fetch(`${API_BASE_URL}/api/players/random`)
    if (!response.ok) {
      return
    }

    const json = (await response.json()) as RandomPlayerResponse
    if (json.uuid) {
      await router.push(`/player/${json.uuid}`)
    }
  } catch (error) {
    console.error('Failed to open a random player', error)
  }
}

function isOnline(uuid: string): boolean {
  return onlinePlayers.value.includes(uuid)
}

onMounted(() => {
  void fetchPlayers()
  startAutoRefresh()
})

onUnmounted(() => {
  stopAutoRefresh()
})
</script>

<template>
  <div class="players-view container page-shell">
    <header class="page-header animate-entry">
      <span class="page-kicker">Public player records</span>
      <h1>{{ siteContent.players.title }}</h1>
      <p class="page-lede">{{ siteContent.players.body }}</p>
    </header>

    <section class="controls-row glass-card animate-entry delay-100">
      <label class="search-field">
        <span>Search</span>
        <input v-model="searchQuery" :placeholder="siteContent.players.searchPlaceholder" />
      </label>

      <div class="control-actions">
        <button :class="['toggle-chip', { active: !showAll }]" type="button" @click="showAll = false">
          {{ siteContent.players.recentLabel }}
          <small v-if="!showAll">last {{ activeDays }} days</small>
        </button>
        <button :class="['toggle-chip', { active: showAll }]" type="button" @click="showAll = true">
          {{ siteContent.players.allLabel }}
        </button>
        <button class="btn-secondary random-button" type="button" :title="siteContent.players.randomTitle" @click="handleRandom">
          Random
        </button>
      </div>
    </section>

    <p class="results-line animate-entry delay-200">
      <span class="badge-pill">{{ count }} {{ count === 1 ? 'player' : 'players' }}</span>
      <span v-if="searchQuery">matching “{{ searchQuery }}”</span>
      <span v-else-if="!showAll">active within the last {{ activeDays }} days</span>
      <span v-else>in the full archive</span>
    </p>

    <div v-if="loading && players.length === 0" class="glass-card state-card">
      {{ siteContent.players.loading }}
    </div>

    <div v-else-if="players.length === 0" class="glass-card state-card">
      <h2>{{ siteContent.players.emptyTitle }}</h2>
      <p>{{ siteContent.players.emptyBody }}</p>
      <button class="btn-primary" type="button" @click="() => { searchQuery = ''; showAll = false }">
        {{ siteContent.players.reset }}
      </button>
    </div>

    <div v-else class="player-grid">
      <RouterLink
        v-for="(player, index) in sortedPlayers"
        :key="player.uuid"
        :to="`/player/${player.uuid}`"
        class="player-card glass-card animate-entry"
        :style="{ animationDelay: `${(index % 18) * 0.03}s` }"
      >
        <div class="player-head">
          <img
            :src="getAvatarUrl(player.last_known_name, player.type)"
            :alt="player.last_known_name || siteContent.players.firstSeenFallback"
            class="player-avatar"
            loading="lazy"
          />
          <div>
            <h2 :class="['player-name', 'minecraft-font', { online: isOnline(player.uuid) }]">
              {{ player.last_known_name || siteContent.players.firstSeenFallback }}
            </h2>
            <p class="player-meta">
              {{ player.type }} · {{ isOnline(player.uuid) ? siteContent.players.onlineNow : 'Last seen ' + new Date(player.last_seen).toLocaleDateString() }}
            </p>
          </div>
        </div>
        <span class="player-arrow">Open record</span>
      </RouterLink>
    </div>
  </div>
</template>

<style scoped>
.controls-row {
  display: grid;
  grid-template-columns: 1fr auto;
  gap: 1rem;
  margin-bottom: 1rem;
}

.search-field {
  display: grid;
  gap: 0.55rem;
}

.search-field span {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.75rem;
  text-transform: uppercase;
  letter-spacing: 0.12em;
}

.search-field input {
  width: 100%;
  min-height: 3.2rem;
  padding: 0 1rem;
  border: 1px solid var(--glass-border);
  border-radius: 999px;
  background: rgba(255, 248, 234, 0.03);
  color: var(--text-main);
}

.control-actions {
  display: flex;
  flex-wrap: wrap;
  justify-content: flex-end;
  gap: 0.75rem;
}

.toggle-chip {
  display: grid;
  gap: 0.2rem;
  padding: 0.7rem 1rem;
  border-radius: 18px;
  border: 1px solid var(--glass-border);
  background: rgba(255, 248, 234, 0.03);
  color: var(--text-muted);
}

.toggle-chip.active {
  border-color: rgba(196, 122, 66, 0.4);
  color: var(--text-strong);
}

.toggle-chip small {
  color: var(--text-dim);
  font-size: 0.72rem;
}

.random-button {
  min-width: 120px;
}

.results-line {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
  align-items: center;
  margin-bottom: 1rem;
  color: var(--text-muted);
}

.state-card {
  display: grid;
  gap: 1rem;
  justify-items: start;
}

.player-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 1rem;
}

.player-card {
  display: grid;
  gap: 1.4rem;
}

.player-head {
  display: flex;
  gap: 1rem;
  align-items: center;
}

.player-avatar {
  width: 72px;
  height: 72px;
  border-radius: 18px;
  border: 1px solid rgba(255, 248, 234, 0.08);
  image-rendering: pixelated;
}

.player-name {
  font-size: 1.2rem;
}

.player-name.online {
  color: #a9d08e;
}

.player-meta,
.player-arrow {
  color: var(--text-muted);
  font-size: 0.92rem;
}

.player-arrow {
  font-family: var(--mono);
  text-transform: uppercase;
  letter-spacing: 0.08em;
}

@media (max-width: 980px) {
  .controls-row,
  .player-grid {
    grid-template-columns: 1fr;
  }

  .control-actions {
    justify-content: flex-start;
  }
}
</style>
