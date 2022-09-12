import Vue from 'vue'
import App from './App.vue'
import router from './router'

import './plugin/http'
import './plugin/antui'
import './assets/css/style.css'
import VMdEditor from './plugin/mdEditor'

Vue.config.productionTip = false
Vue.use(VMdEditor)

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
