import Vue from "vue";
import router from './router';
import App from "./App.vue";
import store from "./store";
import axios from "./backend";

Vue.config.productionTip = false;

new Vue({
  router,
  axios,
  store,
  render: h => h(App)
}).$mount("#app");
