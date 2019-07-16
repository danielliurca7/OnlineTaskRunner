import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    username: "",
    open_tabs: [],

    // should pull from server
    subjects: [
      { name: "Programarea Calculatoarelor" },
      { name: "Algoritmi Paraleli si Distribuiti" },
      { name: "Proiectarea Algoritmilor" },
      { name: "Paradigme de Programare" },
      { name: "Protocoale de Comunicatie" }
    ]
  },
  mutations: {
    setUsername: (state, username) => (state.username = username),
    openTab: (state, tabName) => state.open_tabs.push({ name: tabName }),
    closeTab: (state, tabName) =>
      (state.open_tabs = state.open_tabs.filter(el => el.name !== tabName))
  },
  actions: {
    setUsername({ commit }, username) {
      commit("setUsername", username);
    },
    openTab({ commit }, tabName) {
      commit("openTab", tabName);
    },
    closeTab({ commit }, tabName) {
      commit("closeTab", tabName);
    }
  },
  getters: {
    getUsername: state => state.username,
    getSubjects: state => state.subjects,
    getTabs: state => state.open_tabs
  }
});
