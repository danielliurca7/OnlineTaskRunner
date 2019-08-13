import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import BootstrapVue from "bootstrap-vue";
import VueCookie from "vue-cookie";
import BootstrapVueTreeview from "bootstrap-vue-treeview";
import axios from "axios";
import VueAxios from "vue-axios";
import VueSocketio from "vue-socket.io-extended";
import io from "socket.io-client";

import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";

Vue.use(BootstrapVue);
Vue.use(VueCookie);
Vue.use(BootstrapVueTreeview);
Vue.use(VueAxios, axios);
Vue.use(
  VueSocketio,
  io("ws://localhost:8000", {
    transports: ["websocket"]
  })
);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  sockets: {
    connect() {
      console.log("socket connected");
    },
    disconnect() {
      console.log("socket disconnected");
    }
  },
  render: h => h(App)
}).$mount("#app");
