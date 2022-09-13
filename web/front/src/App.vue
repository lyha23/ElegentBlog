<template>
  <div class="bg h-full w-full -z-2 fixed"></div>
  <div class="flex justify-center">
    <wavesBg class="fixed" :top="0"></wavesBg>

    <div class="flex justify-start">
      <div class="flex flex-col h-auto h-screen-lg w-screen-xs p-4 items-center ">
        <!--用来占位-->
      </div>
      <div class="flex flex-col h-auto h-screen-lg w-screen-xs p-4 items-center fixed">
        <left-bar />
      </div>
      <div class="bg-light-50 shadow-md w-screen-lg p-5 z-3 ">
        <router-view />
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import wavesBg from '@/components/wave.vue'

import router from './router';
import { get } from './utils/http/axios';
import { useUserStore } from '@/store'

const userStore = useUserStore()

async function getProfileInfo() {
  let res = await get({
    url: `profile/1`
  })
  return res
}

router.beforeEach(async (to, from) => {
  console.log("from:", from, `to: `, to)
})

onBeforeMount(async () => {
  let res = await getProfileInfo()
  console.log(res)
  userStore.setInfo(res.data)
})
</script>

<style lang="scss">
@import '@/assets/iconfont/iconfont.css';

#app {
  @apply font-medium;
  font-family: Roboto, Noto, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  background-color: var(--color-bg-1);
}

.bg {
  /* 兜底，IE和Firefox浏览器 */
  filter: blur(500px); //重点：使用的是 filter 来实现的
  --transparent: url(data:image/gif;base64,R0lGODlhAQABAIAAAP///wAAACH5BAEAAAAALAAAAAABAAEAAAICRAEAOw==);
  /* Safari最近版本已经不需要私有前缀了 */
  background: cross-fade(var(--transparent), url('@/assets/images/yokina.png'), 50%) no-repeat top center;
  /* 如使用自定义属性，-webkit-语句需要放在没有私有前缀语句的下面 */
  background: -webkit-cross-fade(var(--transparent), url('@/assets/images/yokina.png'), 50%) no-repeat top center;
  background-size: 100% auto;
  filter: Alpha(Opacity=50);
}
</style>
