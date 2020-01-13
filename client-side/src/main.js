import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import BootstrapVue from "bootstrap-vue";
import VueCookie from "vue-cookie";
import BootstrapVueTreeview from "bootstrap-vue-treeview";
import axios from "axios";
import VueAxios from "vue-axios";
import VueCryptojs from "vue-cryptojs";

import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";

Vue.use(BootstrapVue);
Vue.use(VueCookie);
Vue.use(BootstrapVueTreeview);
Vue.use(VueAxios, axios);
Vue.use(VueCryptojs);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
