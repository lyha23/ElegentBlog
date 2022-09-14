<template>
  <div class="bg h-full w-full -z-2 fixed"></div>

  <button class="xl:hidden" @click="display = true">
    <i style="font-size:24px" class=" p-5 iconfont icon-caidan" />
  </button>

  <drawer title="菜单" v-model:display.sync="display" :inner="true">
    <theMenu />
  </drawer>
  <div class="flex justify-center">
    <div class="flex justify-center items-center xl:overflow-hidden xl:justify-start <xl:flex-col">
      <div v-if="notMobblePage"
        class="flex flex-col w-screen-xs p-4 z-3 items-center hideBar xl:h-200 xl:overflow-y-auto ">
        <leftBar />
      </div>
      <div class="h-auto bg-light-50 shadow-md p-5  z-3 hideBar xl:h-200 xl:w-screen-lg xl:overflow-y-auto">
        <router-view />
      </div>
    </div>
    <wavesBg class="fixed " :top="0"></wavesBg>
  </div>
</template>
<script setup lang="ts">
import wavesBg from '@/components/wave.vue'
import drawer from '@/components/sideDrawer.vue'
import leftBar from '@/components/leftBar.vue'

import router from './router';
import { useCategoryStore, useUserStore } from '@/store'
import theMenu from '@/components/theMenu.vue';

const userStore = useUserStore()
const categoryStore = useCategoryStore()
const Route = useRoute()
const notMobblePage = computed(() => {
  if (document.body.clientWidth < 1200 && (Route.name != "page" && Route.name != "tidy" && Route.name != "about"))
    return false
  else {
    return true
  }
})

const display = ref(false)
router.beforeEach(async (to, from) => {
  console.log("from:", from, `to: `, to)
})

onBeforeMount(async () => {
  await categoryStore.updateInfo()
  await userStore.updateInfo()
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

.hideBar::-webkit-scrollbar {
  width: 0;
}

.hideBar {
  -ms-overflow-style: none;
  overflow: -moz-scrollbars-none;
}
</style>
