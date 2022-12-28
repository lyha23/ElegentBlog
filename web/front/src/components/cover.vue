<script setup lang="ts">
  import { PropType } from 'vue';
  import day from 'dayjs';

  function dateformate(indate) {
    return day(indate).format(`YYYY-MM-DD HH:mm`);
  }

  function xsdateformate(indate) {
    return day(indate).format(`MM月DD日`);
  }

  type articleType = {
    Category?: Object;
    CreatedAt?: string;
    DeletedAt?: null;
    ID?: Number;
    UpdatedAt?: string;
    cid?: Number;
    comment_count?: Number;
    content?: string;
    desc?: string;
    img?: string;
    read_count?: Number;
    title?: string;
  };

  defineProps({
    item: {
      type: Object as PropType<articleType>,
      default: () => {
        content: '';
      },
    },
  });
</script>

<template>
  <div class="block xl:w-100">
    <router-link :to="`/detail?id=${item?.ID}`" class="w-full btn">
      <div
        class="rounded-xl h-30 w-full ceshi relative article title-bg xl:h-screen-xs xl:w-100"
        :style="{ background: `url(${item?.img}) no-repeat center center` }"
      >
        <div
          class="p-10 bottom-0 left-0 text-light-800 text-3xl absolute title <xl:text-xl <xl:p-3 dark:text-gray-200"
        >
          {{ item?.title }}</div
        >
      </div>
    </router-link>
    <div class="flex mt-3">
      <div class="bg-[rgba(65,105,225,0.15)] rounded-3xl p-1 text-[#4169e1] xl:px-4 xl:hidden">
        {{ xsdateformate(item?.UpdatedAt) }}
      </div>
      <div class="bg-[rgba(65,105,225,0.15)] rounded-3xl p-2 px-4 text-[#4169e1] <xl:hidden">
        {{ dateformate(item?.UpdatedAt) }}
      </div>
      <div class="bg-[rgba(255,118,30,0.15)] rounded-3xl ml-3 p-1 text-[#ff761e] xl:p-2 xl:px-4">
        {{ item?.content?.length || '0' }} 字
      </div>
      <div class="bg-[rgba(255,185,0,0.15)] rounded-3xl ml-3 p-1 text-[#ffb900] xl:p-2 xl:px-4"
        >大概 {{ Math.ceil(item?.content?.length / 300) || '0' }} 分钟
      </div>
      <!-- <div class="bg-[rgba(51,213,122,0.15)] rounded-3xl ml-3 p-2 px-4 text-[#33d57a]">置顶</div> -->
    </div>
    <div class="mt-3 mb-5">{{ item?.desc }}</div>
  </div>
</template>

<style type="scss">
  .title-bg {
    background-size: cover !important;
  }

  .title {
    @apply rounded-b-xl w-full;
    background: -webkit-linear-gradient(bottom, gray, transparent);
  }
</style>
