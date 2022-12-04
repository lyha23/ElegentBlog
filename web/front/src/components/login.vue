<template>
  <div class="flex h-full w-full justify-center items-center">
    <el-card class="w-100">
      <div class="flex py-3 items-center">
        <div class="w-15
        ">账户:</div> <el-input v-model="loginForm.username" placeholder="账户" />
      </div>
      <div class="flex py-3 items-center">
        <div class="w-15
        ">密码:</div> <el-input v-model="loginForm.password" placeholder="密码" />
      </div>
      <div class="flex justify-center">
        <el-button @click="login">登录</el-button>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import router from '@/router'
import { useUserStore } from '@/store'
import { getCurrentInstance } from "vue"

const Route = useRoute()
const userStore = useUserStore()
console.log(Route.fullPath)
const islogin = computed(() => {
  if (!userStore.isLogin) {
    router
  }
})
const loginForm = reactive({
  username: '',
  password: '',
})
const ctx = getCurrentInstance()?.proxy
const login = async () => {
  let res = await ctx?.$http.post('login', loginForm)
  userStore.token = res.data.token
  router.push('admin')
}

</script>

<style lang="scss" scoped>
.el-container {
  height: 100%;
}
</style>

