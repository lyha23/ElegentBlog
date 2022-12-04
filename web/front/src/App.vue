<template>
  <div class="h-full">
    <Admin v-if="isAdmin" />
    <Front v-else />
  </div>
</template>
<script setup lang="ts">
import Front from '@/router/front.vue'
import Admin from '@/router/admin.vue'
const Route = useRoute()
// console.log(Route)
const isAdmin = computed(() => {
  // return true
  return Route.name?.toString().startsWith('admin')
})

</script>

<style lang="scss">
body,
html {
  height: 100%
}
</style>
