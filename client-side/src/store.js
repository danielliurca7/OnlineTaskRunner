import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    username: "",
    open_tabs: [],

    // should pull from server
    subjects: [
      { name: "PC" },
      { name: "APD" },
      { name: "PA" },
      { name: "EGC" }
    ]
  },
  mutations: {
    setUsername: (state, username) => (state.username = username),
    openTab: (state, tabName) => state.open_tabs.push({name: tabName})
  },
  actions: {
    setUsername({ commit }, username) {
      commit("setUsername", username);
    },
    openTab({ commit }, tabName) {
      commit("openTab", tabName);
    }
  },
  getters: {
    getUsername: state => state.username,
    getSubjects: state => state.subjects,
    getTabs: state => state.open_tabs
  }
});
