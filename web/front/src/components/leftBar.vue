<script setup lang="ts">
  import theMenu from '@/components/theMenu.vue';
  import { useUserStore, useArticleStore, useCategoryStore } from '@/store';

  const userStore = useUserStore();
  const articleStore = useArticleStore();
  const categoryStore = useCategoryStore();

  defineProps({
    articleNum: {
      type: Number,
      default: 0,
    },
  });
</script>

<template>
  <router-link to="/" class="rounded-1/2 h-40 shadow-lg mt-10 w-40">
    <img class="rounded-1/2 h-40 shadow-lg w-40 avatar" :src="userStore?.avatar" />
  </router-link>
  <div class="flex mt-12 w-full justify-evenly">
    <div class="flex flex-col items-center">
      <div class="font-medium text-base">文章</div>
      <div class="text-gray-500">{{ articleStore?.pagination.total || '--' }}</div>
    </div>
    <div class="flex flex-col items-center">
      <div class="font-medium text-base">标签</div>
      <div class="text-gray-500"> {{ categoryStore?.list.length }}</div>
    </div>
  </div>
  <div class="flex flex-col w-screen-xs p-4 z-3 items-center hideBar <xl:hidden">
    <theMenu class=" " />
  </div>
</template>

<style setup lang="scss">
  .avatar {
    box-shadow: 0 0.3rem 2rem rgb(161 177 204 / 60%);
  }

  .selected {
    @apply bg-[#6495ed] shadow-xl text-light-50;
    box-shadow: 0 2px 12px #6495ed;
  }
</style>
