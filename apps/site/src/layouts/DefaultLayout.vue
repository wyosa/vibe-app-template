<script setup lang="ts">
import AppHeader from '@/components/AppHeader.vue'
</script>

<template>
  <div class="flex min-h-dvh flex-col bg-(--ui-bg) text-(--ui-text)">
    <AppHeader />
    <main class="flex-1">
      <slot />
    </main>
  </div>
</template>
