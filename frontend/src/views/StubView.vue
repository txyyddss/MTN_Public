<script setup lang="ts">
import { useRoute } from 'vue-router'

const route = useRoute()
const pageTitle = route.name?.toString().replace('-', ' ') || 'Work in Progress'
</script>

<template>
  <div class="stub-view container animate-entry">
    <div class="glass-card construction-card">
      <div class="icon-spin">⚙️</div>
      <h1 class="title capitalize">{{ pageTitle }}</h1>
      <p class="subtitle">This page is currently under construction. Please check back later!</p>
      <router-link to="/" class="btn-primary mt-4">Return Home</router-link>
    </div>
  </div>
</template>

<style scoped>
.stub-view {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 60vh;
  padding-top: 4rem;
}

.construction-card {
  text-align: center;
  max-width: 500px;
  width: 100%;
  padding: 4rem 2rem;
}

.icon-spin {
  font-size: 4rem;
  margin-bottom: 1.5rem;
  display: inline-block;
  animation: spin 6s linear infinite;
}

@keyframes spin {
  100% { transform: rotate(360deg); }
}

.title {
  font-size: 2.5rem;
  margin-bottom: 1rem;
}

.capitalize {
  text-transform: capitalize;
}

.subtitle {
  color: var(--text-muted);
  margin-bottom: 2.5rem;
}

.mt-4 {
  margin-top: 1rem;
}
</style>
