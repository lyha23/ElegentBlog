<template>
  <div v-if="!islogin">
    <Login />
  </div>
  <div v-else class="flex flex-col h-full items-center">
    <div class="w-full">
      <Header></Header>
      <div class="flex h-full">
        <div class="w-50">
          <Nav></Nav>
        </div>
        <div class="h-full w-full">
          <router-view class="h-full" :key="$route.fullPath"></router-view>
        </div>
      </div>
      <Footer></Footer>
    </div>
  </div>
</template>

<script setup lang="ts">
  import Login from './admin-components/login.vue';
  import Nav from './admin-components/Nav.vue';
  import Footer from './admin-components/Footer.vue';
  import Header from './admin-components/Header.vue';
  import router from '@/router';
  import { useUserStore } from '@/store';
  const Route = useRoute();
  const userStore = useUserStore();
  console.log(Route.fullPath);
  const islogin = computed(() => {
    return userStore.isLogin ? true : false;
  });

  onBeforeMount(async () => {
    await userStore.updateInfo();
  });
</script>

<style lang="scss" scoped>
  .el-container {
    height: 100%;
  }
</style>
