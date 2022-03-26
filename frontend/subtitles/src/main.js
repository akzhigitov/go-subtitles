import Vue from 'vue'
import App from './App.vue'
import VueAxios from 'vue-axios'
import axios from "@/plugins/axios";
import Vuetify from 'vuetify';

Vue.use(VueAxios, axios);
Vue.use(Vuetify);

Vue.config.productionTip = false

new Vue({
  vuetify: new Vuetify(),
  render: h => h(App)
}).$mount('#app')
