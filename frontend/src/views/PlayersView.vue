<script setup lang="ts">
import { computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'

import DossierLoadingPanel from '@/components/common/DossierLoadingPanel.vue'
import RouteHeroHeader from '@/components/common/RouteHeroHeader.vue'
import ThemedPanelFrame from '@/components/common/ThemedPanelFrame.vue'
import { API_BASE_URL } from '@/config'
import { formatPlayerCount, getLocaleValue, getPlatformLabel, useSiteContent } from '@/content/siteContent'
import { usePlayers } from '@/composables/usePlayers'
import { useServerStatusStore } from '@/stores/serverStatus'
import type { RandomPlayerResponse } from '@/types/api'
import { getAvatarUrl } from '@/utils/minecraft'

const router = useRouter()
const statusStore = useServerStatusStore()
const { status } = storeToRefs(statusStore)
const siteContent = useSiteContent()

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
    return siteContent.value.players.resultMatching.replace('{query}', searchQuery.value)
  }

  if (!showAll.value) {
    return siteContent.value.players.resultActiveWindow.replace('{days}', String(activeDays.value))
  }

  return siteContent.value.players.resultArchive
})

function formatLastSeen(value: number): string {
  if (!value) {
    return siteContent.value.players.lastSeenUnavailable
  }

  const formatted = new Intl.DateTimeFormat(getLocaleValue(), {
    year: 'numeric',
    month: 'numeric',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(new Date(value))

  return siteContent.value.players.lastSeen.replace('{value}', formatted)
}

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
  <div class="players-view container page-shell route-page-shell">
    <RouteHeroHeader
      :kicker="siteContent.players.kicker"
      :title="siteContent.players.title"
      :body="siteContent.players.body"
      align="center"
    />

    <ThemedPanelFrame tag="section" class="controls-panel animate-entry delay-100" variant="archive">
      <div class="controls-row">
        <div class="controls-copy">
          <span class="hud-kicker">{{ siteContent.players.controlsKicker }}</span>
          <p class="controls-note">{{ siteContent.players.controlsNote }}</p>
        </div>

        <label class="search-field">
          <span>{{ siteContent.players.searchLabel }}</span>
          <input
            id="player-search"
            v-model="searchQuery"
            class="action-field"
            name="player-search"
            :placeholder="siteContent.players.searchPlaceholder"
          />
        </label>

        <div class="control-actions">
          <div class="segmented-control">
            <button :class="['switch-btn', 'action-inline', 'action-press', { active: !showAll }]" type="button" @click="showAll = false">
              {{ siteContent.players.recentLabel }}
              <small>{{ siteContent.players.recentDetail.replace('{days}', String(activeDays)) }}</small>
            </button>
            <button :class="['switch-btn', 'action-inline', 'action-press', { active: showAll }]" type="button" @click="showAll = true">
              {{ siteContent.players.allLabel }}
              <small>{{ siteContent.players.allDetail }}</small>
            </button>
          </div>
          <button class="btn-secondary random-button" type="button" :title="siteContent.players.randomTitle" @click="handleRandom">
            {{ siteContent.players.randomLabel }}
          </button>
        </div>
      </div>
    </ThemedPanelFrame>

    <p v-if="!loading || players.length > 0" class="results-line animate-entry delay-200">
      <span class="badge-pill">{{ formatPlayerCount(count) }}</span>
      <span>{{ resultsDescriptor }}</span>
    </p>

    <div v-if="loading && players.length === 0" class="player-grid" aria-hidden="true">
      <DossierLoadingPanel
        v-for="index in 6"
        :key="index"
        class="player-card player-card-skeleton animate-entry-soft"
        label="MTN PLAYER DOSSIER"
        :style="{ animationDelay: `${(index - 1) * 0.05}s` }"
      >
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
      </DossierLoadingPanel>
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
        class="player-card player-live-card glass-card action-card animate-entry"
        :style="{ animationDelay: `${(index % 18) * 0.03}s` }"
      >
        <span class="player-card-mark" aria-hidden="true">MTN</span>
        <div class="player-head">
          <img
            :src="getAvatarUrl(player.last_known_name, player.type)"
            :alt="player.last_known_name || siteContent.players.firstSeenFallback"
            class="player-avatar action-media"
            loading="lazy"
          />
          <div class="player-copy">
            <div class="player-line">
              <h2 :class="['player-name', 'minecraft-font', { online: isOnline(player.uuid) }]">
                {{ player.last_known_name || siteContent.players.firstSeenFallback }}
              </h2>
              <span :class="['player-status', 'action-chip', { online: isOnline(player.uuid) }]">
                {{ isOnline(player.uuid) ? siteContent.players.onlineNow : getPlatformLabel(player.type) }}
              </span>
            </div>

            <p class="player-meta">
              {{ isOnline(player.uuid) ? siteContent.players.visibleInLiveList : formatLastSeen(player.last_seen) }}
            </p>
          </div>
        </div>

        <div class="player-footer">
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

.controls-panel :deep(.themed-panel-frame__content) {
  gap: 0;
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
  border: 1px solid var(--control-border);
  border-radius: 18px;
  background: var(--control-bg);
  color: var(--text-main);
}

.search-field input:focus {
  outline: none;
  border-color: var(--control-border-active);
  background: var(--control-bg-hover);
}

.control-actions {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 0.75rem;
  grid-column: 1 / -1;
}

.segmented-control {
  display: flex;
  background: var(--control-surface);
  border-radius: 12px;
  padding: 3px;
  border: 1px solid var(--control-border);
}

.switch-btn {
  display: flex;
  align-items: baseline;
  gap: 0.4rem;
  padding: 0.35rem 0.65rem;
  border-radius: 9px;
  border: 1px solid transparent;
  background: transparent;
  color: var(--control-text);
  font-size: 0.82rem;
  cursor: pointer;
  transition: all var(--transition-fast);
}

.switch-btn:hover:not(.active) {
  color: var(--control-text-hover);
  background: var(--control-bg-hover);
}

.switch-btn.active {
  border-color: var(--control-border-active);
  background: var(--control-bg-active);
  color: var(--control-text-active);
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.18);
}

.switch-btn small {
  color: var(--text-dim);
  font-size: 0.65rem;
  font-weight: normal;
}

.random-button {
  min-width: 72px;
  padding: 0.35rem 0.65rem;
  font-size: 0.82rem;
  height: max-content;
}

.results-line {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
  align-items: center;
  padding: 0 0.25rem;
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
  padding: 1.2rem;
  min-height: 100%;
  border-color: rgba(76, 147, 251, 0.14);
}

.player-card-mark {
  position: absolute;
  inset: auto 1rem 0.65rem auto;
  z-index: 0;
  color: rgba(141, 184, 255, 0.08);
  font-family: var(--mono);
  font-size: 2.35rem;
  font-weight: 900;
  letter-spacing: 0.24em;
  pointer-events: none;
}

.player-head,
.player-footer {
  position: relative;
  z-index: 1;
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
  line-height: 1.7;
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
