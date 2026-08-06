<script setup lang="ts">
import { RouterView } from 'vue-router'

import DefaultLayout from '@/layouts/DefaultLayout.vue'
</script>

<template>
  <UApp>
    <DefaultLayout>
      <RouterView />
    </DefaultLayout>
  </UApp>
</template>
