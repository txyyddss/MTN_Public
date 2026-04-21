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
const {
  players,
  count,
  activeDays,
  searchQuery,
  showAll,
  loading,
  sortedPlayers,
  fetchPlayers,
  startAutoRefresh,
  stopAutoRefresh
} = usePlayers(onlinePlayers)

const resultsDescriptor = computed(() => {
  if (searchQuery.value) {
    return `matching "${searchQuery.value}"`
  }

  if (!showAll.value) {
    return `active within the last ${activeDays.value} days`
  }

  return 'in the full archive'
})

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

function resetFilters(): void {
  searchQuery.value = ''
  showAll.value = false
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
      <div class="controls-copy">
        <span class="hud-kicker">Roster query</span>
        <p class="controls-note">Search the live archive or move between recent and full-history views.</p>
      </div>

      <label class="search-field">
        <span>Search</span>
        <input
          id="player-search"
          v-model="searchQuery"
          name="player-search"
          :placeholder="siteContent.players.searchPlaceholder"
        />
      </label>

      <div class="control-actions">
        <button :class="['toggle-chip', { active: !showAll }]" type="button" @click="showAll = false">
          {{ siteContent.players.recentLabel }}
          <small>last {{ activeDays }} days</small>
        </button>
        <button :class="['toggle-chip', { active: showAll }]" type="button" @click="showAll = true">
          {{ siteContent.players.allLabel }}
          <small>full log</small>
        </button>
        <button class="btn-secondary random-button" type="button" :title="siteContent.players.randomTitle" @click="handleRandom">
          Random
        </button>
      </div>
    </section>

    <p v-if="!loading || players.length > 0" class="results-line animate-entry delay-200">
      <span class="badge-pill"><strong>{{ count }}</strong> {{ count === 1 ? 'player' : 'players' }}</span>
      <span>{{ resultsDescriptor }}</span>
    </p>

    <div v-if="loading && players.length === 0" class="player-grid" aria-hidden="true">
      <article v-for="index in 6" :key="index" class="player-card glass-card player-card-skeleton">
        <div class="player-head">
          <span class="skeleton-avatar player-skeleton-avatar"></span>
          <div class="player-copy player-skeleton-copy">
            <span class="skeleton-line player-skeleton-name"></span>
            <span class="skeleton-line player-skeleton-meta"></span>
          </div>
        </div>

        <div class="player-footer">
          <span class="skeleton-chip"></span>
        </div>
      </article>
    </div>

    <div v-else-if="players.length === 0" class="glass-card state-card">
      <h2>{{ siteContent.players.emptyTitle }}</h2>
      <p>{{ siteContent.players.emptyBody }}</p>
      <button class="btn-primary" type="button" @click="resetFilters">
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
          <div class="player-copy">
            <div class="player-line">
              <h2 :class="['player-name', 'minecraft-font', { online: isOnline(player.uuid) }]">
                {{ player.last_known_name || siteContent.players.firstSeenFallback }}
              </h2>
              <span :class="['player-status', { online: isOnline(player.uuid) }]">
                {{ isOnline(player.uuid) ? siteContent.players.onlineNow : player.type }}
              </span>
            </div>

            <p class="player-meta">
              {{ isOnline(player.uuid) ? 'Visible in the live player list.' : `Last seen ${new Date(player.last_seen).toLocaleDateString()}` }}
            </p>
          </div>
        </div>

        <div class="player-footer">
          <span class="player-platform">{{ player.type }}</span>
        </div>
      </RouterLink>
    </div>
  </div>
</template>

<style scoped>
.players-view {
  display: grid;
  gap: 1rem;
}

.controls-row {
  display: grid;
  grid-template-columns: minmax(0, 0.9fr) minmax(0, 1.1fr);
  gap: 1rem;
  align-items: center;
}

.controls-copy {
  display: grid;
  gap: 0.55rem;
}

.controls-note {
  color: var(--text-muted);
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
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 18px;
  background: rgba(255, 255, 255, 0.03);
  color: var(--text-main);
}

.control-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 0.55rem;
  grid-column: 1 / -1;
}

.toggle-chip {
  display: grid;
  gap: 0.2rem;
  min-height: 2.7rem;
  padding: 0.64rem 0.84rem;
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  background: rgba(255, 255, 255, 0.03);
  color: var(--text-muted);
  text-align: left;
  transition:
    transform var(--transition-fast),
    border-color var(--transition-fast),
    background var(--transition-fast);
}

.toggle-chip:hover {
  transform: translateY(-1px);
  border-color: rgba(76, 147, 251, 0.24);
}

.toggle-chip.active {
  border-color: rgba(76, 147, 251, 0.3);
  background: rgba(76, 147, 251, 0.08);
  color: var(--text-strong);
}

.toggle-chip small {
  color: var(--text-dim);
  font-size: 0.72rem;
}

.random-button {
  min-width: 112px;
  min-height: 2.7rem;
  padding: 0.64rem 0.9rem;
}

.results-line {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
  align-items: center;
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
  gap: 1rem;
}

.player-card-skeleton {
  min-height: 172px;
}

.player-head {
  display: flex;
  gap: 1rem;
  align-items: center;
}

.player-avatar {
  width: 72px;
  height: 72px;
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  image-rendering: pixelated;
}

.player-copy {
  flex: 1;
  display: grid;
  gap: 0.45rem;
}

.player-line {
  display: flex;
  flex-wrap: wrap;
  gap: 0.6rem;
  align-items: center;
}

.player-name {
  font-size: 1.2rem;
}

.player-name.online {
  color: var(--success);
}

.player-status {
  padding: 0.3rem 0.55rem;
  border-radius: 999px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.68rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.player-status.online {
  color: var(--success);
}

.player-meta {
  color: var(--text-muted);
  font-size: 0.92rem;
}

.player-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
}

.player-platform {
  color: var(--text-muted);
  font-family: var(--mono);
  font-size: 0.72rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.player-skeleton-avatar {
  width: 72px;
  height: 72px;
  border-radius: 16px;
}

.player-skeleton-copy {
  gap: 0.55rem;
}

.player-skeleton-name {
  width: min(11rem, 70%);
}

.player-skeleton-meta {
  width: min(14rem, 90%);
}

@media (max-width: 980px) {
  .controls-row,
  .player-grid {
    grid-template-columns: 1fr;
  }
}
</style>
