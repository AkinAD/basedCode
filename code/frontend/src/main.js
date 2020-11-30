import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import router from './router'
import store from './store'
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import "vuetify/dist/vuetify.min.css";

Vue.config.productionTip = false

new Vue({
  el: "#app",
  vuetify,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
