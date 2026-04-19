<script setup lang="ts">
import IntroCommunityPanel from '@/components/intro/IntroCommunityPanel.vue'
import IntroHeroPanel from '@/components/intro/IntroHeroPanel.vue'
import IntroOperationsPanel from '@/components/intro/IntroOperationsPanel.vue'
import IntroStoryPanel from '@/components/intro/IntroStoryPanel.vue'
</script>

<template>
  <div class="intro-view container page-shell">
    <IntroHeroPanel />
    <IntroStoryPanel />
    <IntroOperationsPanel />
    <IntroCommunityPanel />
  </div>
</template>

<style scoped>
.intro-view {
  display: grid;
  gap: 0;
}
</style>
