<template>
  <div class="flex justify-between xl:hidden">
    <button @click="display = true">
      <i style="font-size:24px" class=" p-5 iconfont icon-caidan" />
    </button>
    <button class="center" @click="toggleTheme()">
      <i style="font-size:28px" class=" p-5 iconfont "
        :class="theme === 'light' ?  'icon-taiyang':'icon-yueliang' "></i>
    </button>
  </div>


  <drawer title="菜单" v-model:display.sync="display" :inner="true">
    <theMenu />
  </drawer>
  <div class="flex justify-center ">
    <div class="flex justify-center items-center xl:overflow-hidden xl:justify-start <xl:flex-col">
      <div v-if="notMobblePage"
        class="flex flex-col w-screen-xs p-4 z-3 items-center hideBar xl:h-200 xl:overflow-y-auto ">
        <leftBar />
      </div>
      <div class="h-auto  shadow-md p-5  z-3 hideBar xl:h-200 xl:w-screen-lg xl:overflow-y-auto">
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
import { useCategoryStore, useUserStore, useAppStore } from '@/store'
import theMenu from '@/components/theMenu.vue';

///
import { useDark, useToggle } from '@vueuse/core';

const appStore = useAppStore();

const isDark = useDark({
  selector: 'body',
  attribute: 'arco-theme',
  valueDark: 'dark',
  valueLight: 'light',
  storageKey: 'arco-theme',
  onChanged(dark: boolean) {
    appStore.toggleTheme(dark);
  },
});
const toggleTheme = useToggle(isDark);
const theme = computed(() => {
  return appStore.theme;
});
///


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
  @apply font-medium bg-light-50 dark: bg-dark-400 dark:text-gray-400 ;
  font-family: Roboto, Noto, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

.hideBar::-webkit-scrollbar {
  width: 0;
}

.hideBar {
  -ms-overflow-style: none;
  overflow: -moz-scrollbars-none;
}
</style>
