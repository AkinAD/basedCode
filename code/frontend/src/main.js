import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import router from './router'
import store from './store'
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import "vuetify/dist/vuetify.min.css";

//AWS
import Amplify, * as AmplifyModules from 'aws-amplify';
import {AmplifyPlugin} from 'aws-amplify-vue'
import '@aws-amplify/ui-vue';
import aws_exports from './aws-exports';
Amplify.configure(aws_exports);
Vue.use(AmplifyPlugin, AmplifyModules)

Vue.config.productionTip = false

new Vue({
  el: "#app",
  vuetify,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
