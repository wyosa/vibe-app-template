<script setup lang="ts">
import { usePageTitle } from '@/composables/usePageTitle'

usePageTitle('App')
</script>

<template>
  <div class="mx-auto w-full max-w-4xl px-6 py-16">
    <h1 class="text-3xl font-bold tracking-tight text-(--ui-text-highlighted)">App</h1>
    <p class="mt-2 text-(--ui-text-muted)">Чистый стартовый шаблон.</p>
  </div>
</template>
