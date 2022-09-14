import { get } from '@/utils/http/axios';
import { defineStore } from 'pinia';


export const useArticleStore = defineStore('article', {
  state: () => ({
    list:undefined,
    pagination: {
      total:0,
      page:1,
      pageSize:10
    },
  }),
  getters: {
    userProfile(state) {
      return { ...state };
    },
  },
  actions: {
    // 设置用户的信息
    setInfo(partial) {
      this.$patch(partial);
    },
    // 重置用户信息
    resetInfo() {
      this.$reset();
    },
  },
});
