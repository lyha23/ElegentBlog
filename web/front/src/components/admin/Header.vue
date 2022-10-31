<template>
  <div class="flex justify-around">
    <div class="log">
      <span>{{ collapsed ? 'Blog' : 'My Blog' }}</span>
    </div>
    <el-button type="danger" @click="loginOut">退出</el-button>
  </div>
</template>

<script>
export default {
  methods: {
    loginOut () {
      window.sessionStorage.clear('token')
      this.$router.push('/login')
    }
  }
}
</script>
