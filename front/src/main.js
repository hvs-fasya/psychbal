import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import axiosCookieJarSupport from 'axios-cookiejar-support'
import VueAxios from 'vue-axios'
import router from './router'
import store from './store'
import './registerServiceWorker'

import './styles/quasar.styl'
import 'quasar-framework/dist/quasar.ie.polyfills'
import 'quasar-extras/animate'
import 'quasar-extras/material-icons'
import Quasar from 'quasar'

Vue.use(Quasar, {
  config: {}
 });

Vue.config.productionTip = false;

Vue.use(VueAxios, axios);
Vue.axios.defaults.baseURL = process.env.VUE_APP_ROOT_API;
axiosCookieJarSupport(axios);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app');
