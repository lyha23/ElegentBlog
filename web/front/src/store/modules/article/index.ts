import { defineStore } from 'pinia';


export const useArticleStore = defineStore('article', {
  state: () => ({
    list:undefined,
    pagination: undefined,
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
