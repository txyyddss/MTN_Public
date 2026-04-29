<script setup lang="ts">
import { computed, shallowRef } from 'vue'

import ThemedPanelFrame from '@/components/common/ThemedPanelFrame.vue'
import type { WhitelistAction, WhitelistEdition, WhitelistMutationInput } from '@/types/api'

defineProps<{
  busy: boolean
  disabled: boolean
}>()

const emit = defineEmits<{
  submit: [action: WhitelistAction, input: WhitelistMutationInput]
}>()

const edition = shallowRef<WhitelistEdition>('java')
const action = shallowRef<WhitelistAction>('add')
const playerName = shallowRef('')
const qqId = shallowRef('')

const canSubmit = computed(() => playerName.value.trim().length > 0)

function submitForm(): void {
  if (!canSubmit.value) {
    return
  }

  emit('submit', action.value, {
    edition: edition.value,
    playerName: playerName.value.trim(),
    qqId: qqId.value.trim() || undefined
  })
}
</script>

<template>
  <ThemedPanelFrame tag="form" class="command-form" kicker="Command" title="Whitelist Control" @submit.prevent="submitForm">
    <div class="control-grid">
      <div class="field-block">
        <span>Edition</span>
        <div class="segmented-control">
          <button
            :class="['segment-btn', { active: edition === 'java' }]"
            type="button"
            :disabled="busy || disabled"
            @click="edition = 'java'"
          >
            Java
          </button>
          <button
            :class="['segment-btn', { active: edition === 'bedrock' }]"
            type="button"
            :disabled="busy || disabled"
            @click="edition = 'bedrock'"
          >
            Bedrock
          </button>
        </div>
      </div>

      <div class="field-block">
        <span>Action</span>
        <div class="segmented-control">
          <button
            :class="['segment-btn', { active: action === 'add' }]"
            type="button"
            :disabled="busy || disabled"
            @click="action = 'add'"
          >
            Add
          </button>
          <button
            :class="['segment-btn', { active: action === 'remove' }]"
            type="button"
            :disabled="busy || disabled"
            @click="action = 'remove'"
          >
            Remove
          </button>
        </div>
      </div>

      <label class="field-block">
        <span>Player</span>
        <input
          v-model="playerName"
          class="action-field"
          type="text"
          autocomplete="off"
          placeholder="playername"
          :disabled="busy || disabled"
        />
      </label>

      <label class="field-block">
        <span>QQ Owner</span>
        <input
          v-model="qqId"
          class="action-field"
          type="text"
          inputmode="numeric"
          autocomplete="off"
          placeholder="optional"
          :disabled="busy || disabled"
        />
      </label>
    </div>

    <div class="form-actions">
      <button class="btn-primary" type="submit" :disabled="busy || disabled || !canSubmit">
        {{ busy ? 'Working' : action === 'add' ? 'Add Entry' : 'Remove Entry' }}
      </button>
    </div>
  </ThemedPanelFrame>
</template>

<style scoped>
.command-form {
  display: grid;
  gap: 1rem;
}

.control-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 0.9rem;
}

.field-block {
  display: grid;
  gap: 0.55rem;
  min-width: 0;
}

.field-block > span {
  color: var(--text-dim);
  font-family: var(--mono);
  font-size: 0.7rem;
  letter-spacing: 0;
  text-transform: uppercase;
}

.field-block input {
  width: 100%;
  min-height: 2.9rem;
  padding: 0 0.9rem;
  border: 1px solid var(--control-border);
  border-radius: var(--radius-sm);
  background: var(--control-bg);
  color: var(--text-main);
}

.field-block input:focus {
  outline: none;
  border-color: var(--control-border-active);
  background: var(--control-bg-hover);
}

.segmented-control {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  min-height: 2.9rem;
  padding: 3px;
  border: 1px solid var(--control-border);
  border-radius: var(--radius-sm);
  background: var(--control-surface);
}

.segment-btn {
  border: 1px solid transparent;
  border-radius: var(--radius-xs);
  background: transparent;
  color: var(--control-text);
  font-size: 0.86rem;
  font-weight: 600;
  transition:
    border-color var(--transition-fast),
    background var(--transition-fast),
    color var(--transition-fast),
    transform var(--transition-fast);
}

.segment-btn:hover:not(:disabled) {
  color: var(--control-text-hover);
  background: var(--control-bg-hover);
}

.segment-btn.active {
  border-color: var(--control-border-active);
  background: var(--control-bg-active);
  color: var(--control-text-active);
}

.form-actions {
  display: flex;
  justify-content: flex-end;
}

@media (max-width: 760px) {
  .control-grid {
    grid-template-columns: 1fr;
  }

  .form-actions button {
    width: 100%;
  }
}
</style>
