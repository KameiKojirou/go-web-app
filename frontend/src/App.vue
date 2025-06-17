<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="p-4">
    <nav class="flex flex-row gap-4 [&>a]:btn [&>a]:btn-ghost">
      <RouterLink to="/">Home</RouterLink>
      <RouterLink to="/about">About</RouterLink>
    </nav>
  </div>
  <main class="p-4">
    <RouterView />
  </main>
</template>
