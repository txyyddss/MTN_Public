<script setup lang="ts">
import { computed } from 'vue'

import WhitelistCommandForm from '@/components/whitelist/WhitelistCommandForm.vue'
import WhitelistEntriesTable from '@/components/whitelist/WhitelistEntriesTable.vue'
import WhitelistTokenPanel from '@/components/whitelist/WhitelistTokenPanel.vue'
import { useWhitelist } from '@/composables/useWhitelist'
import type { WhitelistAction, WhitelistEditionFilter, WhitelistMutationInput } from '@/types/api'

const {
  token,
  editionFilter,
  entries,
  loading,
  mutating,
  error,
  notice,
  hasToken,
  activeCount,
  javaCount,
  bedrockCount,
  loadEntries,
  addEntry,
  removeEntry
} = useWhitelist()

const busy = computed(() => loading.value || mutating.value)

function setFilter(value: WhitelistEditionFilter): void {
  editionFilter.value = value
}

async function handleSubmit(action: WhitelistAction, input: WhitelistMutationInput): Promise<void> {
  if (action === 'add') {
    await addEntry(input)
    return
  }

  await removeEntry(input)
}
</script>

<template>
  <div class="whitelist-view container page-shell">
    <header class="page-header animate-entry">
      <span class="page-kicker">Whitelist Admin</span>
      <h1>Whitelist Console</h1>
      <p class="page-lede">RCON-backed Java and Bedrock access control.</p>
    </header>

    <WhitelistTokenPanel
      v-model:token="token"
      class="animate-entry delay-100"
      :loading="loading"
      :count="activeCount"
      :java-count="javaCount"
      :bedrock-count="bedrockCount"
      :error="error"
      :notice="notice"
      @refresh="loadEntries"
    />

    <section class="admin-layout">
      <WhitelistCommandForm
        class="animate-entry delay-200"
        :busy="mutating"
        :disabled="!hasToken"
        @submit="handleSubmit"
      />

      <div class="list-column animate-entry delay-300">
        <div class="filter-bar hud-panel">
          <span class="hud-kicker">Filter</span>
          <div class="filter-controls segmented-control">
            <button
              :class="['filter-btn', { active: editionFilter === 'all' }]"
              type="button"
              :disabled="busy"
              @click="setFilter('all')"
            >
              All
            </button>
            <button
              :class="['filter-btn', { active: editionFilter === 'java' }]"
              type="button"
              :disabled="busy"
              @click="setFilter('java')"
            >
              Java
            </button>
            <button
              :class="['filter-btn', { active: editionFilter === 'bedrock' }]"
              type="button"
              :disabled="busy"
              @click="setFilter('bedrock')"
            >
              Bedrock
            </button>
          </div>
        </div>

        <WhitelistEntriesTable :entries="entries" :busy="mutating" @remove="removeEntry" />
      </div>
    </section>
  </div>
</template>

<style scoped>
.whitelist-view {
  display: grid;
  gap: 1rem;
}

.page-header h1 {
  max-width: 12ch;
}

.admin-layout {
  display: grid;
  grid-template-columns: minmax(280px, 0.75fr) minmax(0, 1.25fr);
  gap: 1rem;
  align-items: start;
}

.list-column {
  display: grid;
  gap: 1rem;
  min-width: 0;
}

.filter-bar {
  display: grid;
  grid-template-columns: auto minmax(0, 1fr);
  gap: 1rem;
  align-items: center;
}

.filter-controls {
  justify-self: end;
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  width: min(100%, 360px);
  min-height: 2.7rem;
  padding: 3px;
  border: 1px solid var(--control-border);
  border-radius: 8px;
  background: var(--control-surface);
}

.filter-btn {
  border: 1px solid transparent;
  border-radius: 6px;
  background: transparent;
  color: var(--control-text);
  font-size: 0.84rem;
  font-weight: 600;
  transition:
    border-color var(--transition-fast),
    background var(--transition-fast),
    color var(--transition-fast);
}

.filter-btn:hover:not(:disabled) {
  color: var(--control-text-hover);
  background: var(--control-bg-hover);
}

.filter-btn.active {
  border-color: var(--control-border-active);
  background: var(--control-bg-active);
  color: var(--control-text-active);
}

@media (max-width: 980px) {
  .admin-layout {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 680px) {
  .filter-bar {
    grid-template-columns: 1fr;
  }

  .filter-controls {
    justify-self: stretch;
    width: 100%;
  }
}
</style>
