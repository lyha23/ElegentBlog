import { get } from '@/utils/http/axios';
import { defineStore } from 'pinia';
import {
  login as userLogin,
  logout as userLogout,
  getUserProfile,
} from '/@/api/user/index';
import { setToken, clearToken, clearAll } from '/@/utils/auth';

export const useUserStore = defineStore('user', {
  state: () => ({
    ID: undefined,
    avatar: undefined,
    CreatedAt: undefined,
    DeletedAt: undefined,
    role:undefined,
    name:undefined,
    desc:undefined,
    qq_chat:undefined,
    wechat:undefined,
    bili:undefined,
    email:undefined,
    github:undefined,
    img:undefined,
    aboutMe:undefined,
    token:undefined
  }),
  getters: {
    userProfile(state) {
      return { ...state };
    },
    isLogin() {
      if (typeof this.token === 'undefined') {
      return false
      } else {
      return true
      }
    }
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
        url: `profile/1`
      })
      return res
    },
    async updateInfo(){
      if(!this.ID){
        let res = await this.getInfo()
        this.setInfo(res.data)
      }
    },
    // 获取用户信息
    async info() {
      if(!this.ID){
      const result = await getUserProfile();
      this.setInfo(result);
      }
    },
    // 异步登录并存储token
    async login(loginForm) {
      const result = await userLogin();
      const token = result?.token;
      if (token) {
        setToken(token);
      }
      return result;
    },
    async setToken(token: string) {
      if(token){
      this.setInfo({'token':token})
      }
    },
    // Logout
    async logout() {
      await userLogout();
      this.resetInfo();
      clearToken();
      clearAll();
      // 路由表重置
      location.reload();
    },
  },
});
