import 'virtual:windi-base.css';
import 'virtual:windi-components.css';
import 'virtual:windi-utilities.css';

import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import piniaStore from './store';
import './index.css';
//改为element-plus默认UI
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
// 注册 Icon
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
// 支持SVG
import 'virtual:svg-icons-register';
import plugins from '@/plugins'
// import "vis/dist/vis.css";

const app = createApp(App);

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}

app.use(router).use(plugins).use(ElementPlus).use(piniaStore).mount('#app');
