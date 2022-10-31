<template>
  <div class="flex flex-col h-full">
    <Header></Header>
    <div class="flex">
      <div class="w-50">
        <Nav></Nav>
      </div>
      <router-view class="h-full" :key="$route.fullPath"></router-view>
    </div>
    <Footer></Footer>
  </div>

</template>

<script setup lang="ts">
import Nav from '@/components/admin/Nav.vue'
import Footer from '@/components/admin/Footer.vue'
import Header from '@/components/admin/Header.vue'

</script>

<style lang="scss" scoped>
.el-container {
  height: 100%;
}
</style>

