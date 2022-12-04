<template>
  <div class="flex h-15 shadow-light-200 justify-around items-center">
    <div class="flex items-center justify-center">
      <img class="rounded-1/2 h-10 w-10 " :src="userStore?.avatar" />
      <div class="h-5 pl-5">{{ userStore?.name }} 的博客</div>
    </div>
    <el-button type="danger" @click="loginOut">退出</el-button>
  </div>
</template>

<script setup lang='ts'>
import router from '/@/router'
import { useUserStore } from '@/store';
let userStore = useUserStore()
const loginOut = () => {
  userStore.resetInfo()
}
</script>
