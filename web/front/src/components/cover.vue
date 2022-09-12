<script setup lang="ts">
import { PropType } from 'vue';
import day from 'dayjs'

function dateformate(indate) {
  return day(indate).format(`YYYY-MM-DD HH:mm`)
}

type articleType = {
  Category?: Object
  CreatedAt?: string,
  DeletedAt?: null
  ID?: Number,
  UpdatedAt?: string,
  cid?: Number,
  comment_count?: Number,
  content?: string,
  desc?: string,
  img?: string,
  read_count?: Number,
  title?: string,
}

defineProps({
  item: Object as PropType<articleType>,
})
</script>


<template>
  <router-link :to="`/detail?id=${item?.ID}`" class="h-max w-full btn">
    <div class="rounded-xl h-screen-sm w-full ceshi relative article " :style="{background:`url(${item?.img})`}">
      <div class="text-white p-10 bottom-0 left-0 text-5xl absolute title">{{item?.title}}</div>
    </div>
  </router-link>
  <div class="flex mt-3">
    <div class="bg-[rgba(65,105,225,0.15)] rounded-3xl p-2 px-4 text-[#4169e1]">{{dateformate(item?.UpdatedAt)}}</div>
    <div class="bg-[rgba(255,118,30,0.15)] rounded-3xl ml-3 p-2 px-4 text-[#ff761e]">{{item?.content.length||'0'}} 字
    </div>
    <div class="bg-[rgba(255,185,0,0.15)] rounded-3xl ml-3 p-2 px-4 text-[#ffb900]">大概
      {{Math.ceil(item?.content.length/300)||'0'}} 分钟
    </div>
    <!-- <div class="bg-[rgba(51,213,122,0.15)] rounded-3xl ml-3 p-2 px-4 text-[#33d57a]">置顶</div> -->
  </div>
  <div class="mt-3 mb-5">{{item?.desc}}</div>
</template>

<style type="scss">
.title {
  @apply rounded-b-xl w-full;
  background: -webkit-linear-gradient(bottom, gray, transparent)
}
</style>
