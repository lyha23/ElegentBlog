import { createRouter, createWebHashHistory } from 'vue-router';
import routes from 'virtual:generated-pages';
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';

routes.push({
  path: '/',
  redirect: `/page?page=1`,
  name:'index'
})

routes.push(
  {
    path: '/admin/login',
    name: 'login',
    meta: {
      title: '请登录'
    },
  })

//导入生成的路由数据
const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

console.log(router)

router.beforeEach(async (_to, _from, next) => {
  NProgress.start();
  next();
});

router.afterEach((_to) => {
  NProgress.done();
});

export default router;
