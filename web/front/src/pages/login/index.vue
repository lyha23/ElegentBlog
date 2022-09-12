<template>
  <div class="flex justify-center">
    <wxlogin
      appid="wx4d6f7bbc7c5f3e4e"
      :scope="'snsapi_login'"
      :theme="'black'"
      :redirect_uri="encodeURIComponent('https://www.qqoc.co/#/login/next')"
      :href="'data:text/css;base64,LmltcG93ZXJCb3ggLnFyY29kZxxx='"
      rel="external nofollow"
    >
    </wxlogin>
  </div>
</template>

<script setup lang="ts">
  import wxlogin from '/@/components/wxLogin/vue-wxlogin.vue';
  import { login } from '/@/api/user/index';
  import router from '/@/router';
  import { useUserStore } from '/@/store/index';

  const userStore = useUserStore();

  onMounted(() => {
    login().then((res) => {
      if (res.data.Token) {
        console.log(res);
        localStorage.setItem('Token', res.data.Token);
        userStore.setInfo(res.data.userInfo);
        router.push('/');
      }
    });
  });
</script>

<style lang="less" scoped>
  .container {
    display: flex;
    height: 100vh;

    .banner {
      width: 550px;
      background: linear-gradient(163.85deg, #1d2129 0%, #00308f 100%);
    }

    .content {
      position: relative;
      display: flex;
      flex: 1;
      align-items: center;
      justify-content: center;
      padding-bottom: 40px;
    }

    .footer {
      position: absolute;
      right: 0;
      bottom: 0;
      width: 100%;
    }
  }

  .logo {
    position: fixed;
    top: 24px;
    left: 22px;
    z-index: 1;
    display: inline-flex;
    align-items: center;

    &-text {
      margin-right: 4px;
      margin-left: 4px;
      color: var(--color-fill-1);
      font-size: 20px;
    }
  }
</style>
