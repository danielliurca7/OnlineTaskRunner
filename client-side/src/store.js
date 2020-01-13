import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    username: "",
    token: "",
    type: "",
    open_tabs: [],

    // should pull from server
    student_subjects: [
      {
        name: "Algoritmi Paraleli si Distribuiti",
        labs: [
          {
            name: "Laboratorul 1",
            total: 0.25,
            deadline: new Date("03-19-2020"),
            support: "Acesta este laboratorul 1"
          },
          {
            name: "Laboratorul 2",
            total: 0.25,
            deadline: new Date("03-26-2020"),
            support: "Acesta este laboratorul 2"
          }
        ],
        homeworks: [
          {
            name: "Tema 1",
            total: 1,
            deadline: new Date("04-10-2020"),
            support: "Aceasta este tema 1"
          }
        ]
      },
      {
        name: "Programarea Calculatoarelor",
        labs: [
          {
            name: "Laboratorul 1",
            result: 80,
            total: 0.25,
            deadline: new Date("10-16-2018"),
            support: "Acesta este laboratorul 1"
          }
        ],
        homeworks: [
          {
            name: "Tema 1",
            result: 90,
            total: 1,
            deadline: new Date("11-09-2018"),
            support: "Aceasta este tema 1"
          }
        ]
      }
    ],
    assistent_subjects: [
      {
        name: "Programarea Calculatoarelor",
        assignments: [
          {
            name: "Tema 1",
            students: [
              {
                name: "stud1",
                grade: "",
                gradetime: "",
                graded_by: ""
              },
              {
                name: "stud2",
                grade: "",
                gradetime: "",
                graded_by: ""
              },
              {
                name: "stud3",
                grade: "",
                gradetime: "",
                graded_by: ""
              }
            ]
          }
        ]
      }
    ]
  },
  mutations: {
    setUsername: (state, username) => (state.username = username),
    setToken: (state, token) => (state.token = token),
    setType: (state, type) => (state.type = type),
    openTab: (state, tabName) => state.open_tabs.push({ name: tabName }),
    closeTab: (state, tabName) =>
      (state.open_tabs = state.open_tabs.filter(el => el.name !== tabName)),
    grade: (state, payload) => {
      state.assistent_subjects
        .find(subject => subject.name === payload.subject_name)
        .assignments.find(
          assignment => assignment.name === payload.assignment_name
        )
        .students.find(student => student.name === payload.student_name).grade =
        payload.grade;

      state.assistent_subjects
        .find(subject => subject.name === payload.subject_name)
        .assignments.find(
          assignment => assignment.name === payload.assignment_name
        )
        .students.find(
          student => student.name === payload.student_name
        ).gradetime = payload.gradetime;

      state.assistent_subjects
        .find(subject => subject.name === payload.subject_name)
        .assignments.find(
          assignment => assignment.name === payload.assignment_name
        )
        .students.find(
          student => student.name === payload.student_name
        ).graded_by = payload.graded_by;
    }
  },
  actions: {
    setUsername({ commit }, username) {
      commit("setUsername", username);
    },
    setToken({ commit }, token) {
      commit("setToken", token);
    },
    setType({ commit }, type) {
      commit("setType", type);
    },
    openTab({ commit }, tabName) {
      commit("openTab", tabName);
    },
    closeTab({ commit }, tabName) {
      commit("closeTab", tabName);
    },
    gradeAssignment({ commit }, payload) {
      commit("grade", payload);
    }
  },
  getters: {
    getUsername: state => state.username,
    getToken: state => state.token,
    getType: state => state.type,
    getStudentSubjects: state => state.student_subjects,
    getAssistentSubjects: state => state.assistent_subjects,
    getTabs: state => state.open_tabs,
    getLabs: state => name =>
      state.student_subjects
        .find(subject => subject.name === name)
        .labs.filter(lab => Date.now() < lab.deadline),
    getHomeworks: state => name =>
      state.student_subjects
        .find(subject => subject.name === name)
        .homeworks.filter(homework => Date.now() < homework.deadline),
    getAssignments: state =>
      state.student_subjects
        .flatMap(subject =>
          subject.labs
            .map(lab => ({
              type: "lab",
              subjectName: subject.name,
              name: lab.name,
              points: lab.points,
              total: lab.total,
              deadline: lab.deadline
            }))
            .concat(
              subject.homeworks.map(homework => ({
                type: "homework",
                subjectName: subject.name,
                name: homework.name,
                points: homework.points,
                total: homework.total,
                deadline: homework.deadline
              }))
            )
        )
        .filter(assignment => Date.now() < assignment.deadline),
    getAllResults: state =>
      state.student_subjects
        .flatMap(subject =>
          subject.labs
            .map(lab => ({
              subjectName: subject.name,
              name: lab.name,
              result: lab.result,
              total: lab.total,
              deadline: lab.deadline
            }))
            .concat(
              subject.homeworks.map(homework => ({
                subjectName: subject.name,
                name: homework.name,
                result: homework.result,
                total: homework.total,
                deadline: homework.deadline
              }))
            )
        )
        .filter(assignment => Date.now() > assignment.deadline),
    getResults: state => name =>
      state.student_subjects
        .find(subject => subject.name === name)
        .homeworks.filter(homework => Date.now() > homework.deadline)
        .concat(
          state.student_subjects
            .find(subject => subject.name === name)
            .labs.filter(lab => Date.now() > lab.deadline)
        ),
    getAssignmentsToGrade: state => name =>
      state.assistent_subjects
        .find(subject => subject.name === name)
        .assignments.flatMap(assignment => ({
          name: assignment.name
        })),
    getAssignmentStudents: state => (subject_name, assignment_name) =>
      state.assistent_subjects
        .find(subject => subject.name === subject_name)
        .assignments.find(assignment => assignment.name === assignment_name)
        .students
  }
});
