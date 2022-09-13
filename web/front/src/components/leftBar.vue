<script setup>
import router from '@/router';
import { useUserStore, useArticleStore } from '@/store';
import { get, post } from '/@/utils/http/axios';

const userStore = useUserStore();
const articleStore = useArticleStore()

const Route = useRoute()
defineProps({
  articleNum: {
    type: Number,
    default: 0
  }
})


const list = ref([])
onBeforeMount(async () => {
  let res = await get({
    url: '/category',
  })
  if (res.code === 200) {
    list.value = res.list
  }
})

let type = ['success', 'info', 'warning', 'danger']

</script>

<template>
  <router-link to="/" class="rounded-1/2 h-40 shadow-lg mt-10 w-40">
    <img class="rounded-1/2 h-40 shadow-lg w-40 avatar" :src="userStore?.avatar" />
  </router-link>
  <div class="flex mt-12 w-full justify-evenly">
    <div class="flex flex-col items-center">
      <div class="font-medium text-base">文章</div>
      <div class="text-gray-500">{{articleStore?.pagination.total}}</div>
    </div>
    <div class="flex flex-col items-center">
      <div class="font-medium text-base">标签</div>
      <div class="text-gray-500">{{list?.length}}</div>
    </div>
  </div>
  <router-link :to="`/`"
    class="rounded-xl flex m-1  mt-8 w-full py-2 text-[#9ca2a8] duration-300 justify-center hover:bg-gray-200 hover:shadow-xl"
    :class="Route.fullPath==`/page?page=1`?'selected':''">
    <div class="flex  w-20 justify-between items-center"><i style="font-size:22px"
        class="text-xl iconfont icon-Leftbar_index_unselected"></i>首页
    </div>
  </router-link>
  <router-link :to="`/tidy`"
    class="rounded-xl flex m-1 w-full py-2 text-[#9ca2a8] duration-300 justify-center hover:bg-gray-200 hover:shadow-xl"
    :class="Route.name==`tidy`?'selected':''">
    <div class="flex  w-20 justify-between items-center"><i style="font-size:22px"
        class="text-xl iconfont icon-Calander"></i>归档
    </div>
  </router-link>
  <router-link :to="`about`" :class="Route.name==`about`?'selected':''"
    class="rounded-xl  flex  m-1 w-full py-2 text-[#9ca2a8] duration-300 justify-center hover:bg-gray-200 hover:shadow-xl">
    <div class="flex  w-20 justify-between items-center"><i style="font-size:22px"
        class="text-xl icon-Leftbar_mine_unselected iconfont"></i>关于
    </div>
  </router-link>

  <input class="rounded-xl shadow-md mt-5 w-full p-4" placeholder="搜全站！" />
  <div class="bg-white rounded-xl flex flex-wrap shadow-md my-1 mt-5 w-full p-3">
    <router-link :to="`/page?category=${item.id}&page=1`" v-for="(item,index) in list">
      <el-tag :type="type[index%4]" class="my-1 mr-2" round>
        {{item.name}}</el-tag>
    </router-link>
  </div>
  <div class="rounded-xl flex bg-light-50 shadow-md mt-5 w-full p-3 justify-evenly">
    <!-- <div class="font-medium text-[#6495ed]">文章归档</div> -->
    <a :href="userStore.bili">
      <el-tooltip effect="dark" content="bilibili" placement="top-start">
        <div
          class="flex bg-[rgb(231,106,141,.15)] rounded-1/2 h-9 text-[rgb(231,106,141)] w-9 justify-center items-center">
          <i style="font-size:22px" class="iconfont icon-xiankuangjiazaishibaidaitiicon" />
        </div>
      </el-tooltip>
    </a>
    <el-tooltip effect="dark" :content="`wechat: ${userStore.wechat}`" placement="top-start">
      <div class="flex bg-[rgb(7,193,96,.15)] rounded-1/2 h-9 text-[#07c160] w-9 justify-center items-center">
        <i style="font-size:22px" class="iconfont icon-weixin" />
      </div>
    </el-tooltip>
    <el-tooltip effect="dark" :content="`qq: ${userStore.qq_chat}`" placement="top-start">
      <div class="flex bg-[rgb(18,183,245,.15)] rounded-1/2 h-9 text-[rgb(18,183,245)] w-9 justify-center items-center">
        <i style="font-size:22px" class="iconfont icon-qq" />
      </div>
    </el-tooltip>
    <a :href="userStore.github">
      <el-tooltip effect="dark" content="github" placement="top-start">
        <div class="flex bg-[rgb(25,23,23,.15)] rounded-1/2 h-9 text-[rgb(25,23,23)] w-9 justify-center items-center">
          <i style="font-size:22px" class="iconfont icon-github-fill" />
        </div>
      </el-tooltip>
    </a>
  </div>
  <view class="flex w-full justify-center">
    <view class="flex my-3 text-sm text-[#767676] w-78">
      <text class="text-sm text-center text-[#767676] hyphens-auto" style="color: #868686">&copy; 2022 li.
        All&nbsp;Rights&nbsp;Reserved.
      </text>
    </view>
  </view>
  <div class="mt-2 text-sm text-gray-400">
    <div>地址：上海市闵行区剑川路600号G202</div>
    <!-- <div>Tel:15900733604</div> -->
    <a target="_blank" class="flex items-center" href="http://www.beian.gov.cn/"><img
        src="@/assets/images/copyright.jpg" class="w-4" />
      <div class="pl-1">沪ICP备2022004472号</div>
    </a>
  </div>
</template>

<style setup lang="scss">
.avatar {
  box-shadow: 0 0.3rem 2rem rgb(161 177 204 / 60%);
}

.selected {
  @apply bg-[#6495ed] shadow-xl text-light-50;
  box-shadow: 0 2px 12px #6495ed
}
</style>
