import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    username: "",
    open_tabs: [],

    // should pull from server
    subjects: [
      {
        name: "Programarea Calculatoarelor",
        labs: [
          {
            name: "Laboratorul 1",
            support: "Acesta este laboratorul 1",
            given_code: "int main() {}",
            tests: [
              { name: "Test 1", value: "Passed" },
              { name: "Test 2", value: "Failed" },
              { name: "Test 3", value: "Passed" }
            ]
          },
          {
            name: "Laboratorul 2",
            support: "Acesta este laboratorul 2",
            given_code: "void main() {}",
            tests: [
              { name: "Test 1", value: "Passed" },
              { name: "Test 2", value: "Passed" },
              { name: "Test 3", value: "Passed" }
            ]
          }
        ],
        homeworks: [
          {
            name: "Tema 1",
            support: "Aceasta este tema 1",
            given_code: "#include <iostream>",
            tests: [
              { name: "Test 1", value: "Passed" },
              { name: "Test 2", value: "Passed" },
              { name: "Test 3", value: "Passed" }
            ]
          }
        ]
      },
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
    getTabs: state => state.open_tabs,
    getLabs: state => name =>
      state.subjects.find(subject => subject.name === name).labs,
    getHomeworks: state => name =>
      state.subjects.find(subject => subject.name === name).homeworks
  }
});
