import { createPinia } from 'pinia';
import { useAppStore } from './modules/app';
import { useUserStore } from './modules/user';
import { useArticleStore } from './modules/article';
import { createPersistedState } from "pinia-persistedstate-plugin";

const pinia = createPinia();
pinia.use(createPersistedState());

export { useAppStore, useUserStore,useArticleStore };
export default pinia;
