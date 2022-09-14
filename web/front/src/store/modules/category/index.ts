import { get } from '@/utils/http/axios';
import { defineStore } from 'pinia';


export const useCategoryStore = defineStore('list', {
  state: () => ({
    list:[],
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
    async getInfo() {
      let res = await get({
        url: '/category',
      })
      if (res.code === 200) {
        return res
      }
    },
    async updateInfo(){
      let res = await this.getInfo()
      this.setInfo(res)
    }
  },
});
