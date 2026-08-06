<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
  <header class="sticky top-0 z-40 border-b border-(--ui-border) bg-(--ui-bg)/80 backdrop-blur-md">
    <div class="mx-auto flex h-16 w-full max-w-6xl items-center justify-between px-6">
      <RouterLink to="/" class="text-lg font-semibold tracking-tight text-(--ui-text-highlighted)">
        App
      </RouterLink>
      <UColorModeButton />
    </div>
  </header>
</template>
