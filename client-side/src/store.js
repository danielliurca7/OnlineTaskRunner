import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    username: "",
    userinfo: {},
    token: "",
    open_tabs: [],
    workspaces: [],

    student_courses: [],
    assistant_courses: []
  },
  mutations: {
    setUsername: (state, username) => (state.username = username),
    setUserinfo: (state, userinfo) => (state.userinfo = userinfo),
    setToken: (state, token) => (state.token = token),
    addWorkspace: (state, workspace) => state.workspaces.push(workspace),
    changeFile: (state, payload) =>
      (state.workspaces
        .find(
          ws =>
            ws.workspace.course === payload.workspace.course &&
            ws.workspace.series === payload.workspace.series &&
            ws.workspace.year === parseInt(payload.workspace.year) &&
            ws.workspace.assignmentname === payload.workspace.assignmentname &&
            ws.workspace.owner === payload.workspace.owner
        )
        .files.find(
          file => JSON.stringify(file.path) === JSON.stringify(payload.path)
        ).data = payload.newVal),
    setStudentCourses: (state, student_courses) =>
      (state.student_courses = student_courses),
    setAssistantCourses: (state, assistant_courses) =>
      (state.assistant_courses = assistant_courses),
    openTab: (state, assignment) => state.open_tabs.push(assignment),
    closeTab: (state, assignment) =>
      (state.open_tabs = state.open_tabs.filter(
        el => el.name !== assignment.name
      )),
    grade: (state, payload) => {
      state.assistant_courses
        .find(course => course.name === payload.course_name)
        .assignments.find(
          assignment => assignment.name === payload.assignment_name
        )
        .students.find(student => student.name === payload.student_name).grade =
        payload.grade;

      state.assistant_courses
        .find(course => course.name === payload.course_name)
        .assignments.find(
          assignment => assignment.name === payload.assignment_name
        )
        .students.find(
          student => student.name === payload.student_name
        ).gradetime = payload.gradetime;

      state.assistant_courses
        .find(course => course.name === payload.course_name)
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
    setUserinfo({ commit }, userinfo) {
      commit("setUserinfo", userinfo);
    },
    setToken({ commit }, token) {
      commit("setToken", token);
    },
    addWorkspace({ commit }, workspace) {
      commit("addWorkspace", workspace);
    },
    changeFile({ commit }, payload) {
      commit("changeFile", payload);
    },
    setStudentCourses({ commit }, student_courses) {
      commit("setStudentCourses", student_courses);
    },
    setAssistantCourses({ commit }, assistant_courses) {
      commit("setAssistantCourses", assistant_courses);
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
    getUserinfo: state => state.userinfo,
    getToken: state => state.token,
    getWorkspaceFile: state => (
      course,
      series,
      year,
      assignment,
      owner,
      path
    ) =>
      state.workspaces
        .find(
          workspace =>
            workspace.workspace.course === course &&
            workspace.workspace.series === series &&
            workspace.workspace.year === parseInt(year) &&
            workspace.workspace.assignmentname === assignment &&
            workspace.workspace.owner === owner
        )
        .files.find(file => JSON.stringify(file.path) === JSON.stringify(path)),
    getFilesByStudent: state => (course, year, series, assignment, owner) =>
      state.files.find(
        file =>
          file.course === course &&
          file.year === year &&
          file.series === series &&
          file.assignment === assignment &&
          file.owner === owner
      ).files,
    getStudentCourses: state => state.student_courses,
    getAssistantCourses: state => state.assistant_courses,
    getTabs: state => state.open_tabs,
    getLabs: state => name =>
      state.student_courses
        .filter(course => course.name === name)
        .flatMap(course =>
          course.labs.map(lab => ({
            course: course.name,
            year: course.year,
            series: course.series,
            abbreviation: course.abbreviation,
            name: lab.name,
            points: lab.points,
            total: lab.total,
            deadline: lab.deadline
          }))
        )
        .filter(assignment => Date.now() < assignment.deadline),
    getHomeworks: state => name =>
      state.student_courses
        .filter(course => course.name === name)
        .flatMap(course =>
          course.homeworks.map(homework => ({
            course: course.name,
            year: course.year,
            series: course.series,
            abbreviation: course.abbreviation,
            name: homework.name,
            points: homework.points,
            total: homework.total,
            deadline: homework.deadline
          }))
        )
        .filter(assignment => Date.now() < assignment.deadline),
    getAssignments: state =>
      state.student_courses
        .flatMap(course =>
          course.labs
            .map(lab => ({
              course: course.name,
              year: course.year,
              series: course.series,
              abbreviation: course.abbreviation,
              name: lab.name,
              points: lab.points,
              total: lab.total,
              deadline: lab.deadline
            }))
            .concat(
              course.homeworks.map(homework => ({
                course: course.name,
                year: course.year,
                series: course.series,
                abbreviation: course.abbreviation,
                name: homework.name,
                points: homework.points,
                total: homework.total,
                deadline: homework.deadline
              }))
            )
        )
        .filter(assignment => Date.now() < assignment.deadline),
    getAllResults: state =>
      state.student_courses
        .flatMap(course =>
          course.labs
            .map(lab => ({
              course: course.name,
              year: course.year,
              series: course.series,
              abbreviation: course.abbreviation,
              name: lab.name,
              result: lab.result,
              total: lab.total,
              deadline: lab.deadline
            }))
            .concat(
              course.homeworks.map(homework => ({
                course: course.name,
                year: course.year,
                series: course.series,
                abbreviation: course.abbreviation,
                name: homework.name,
                result: homework.result,
                total: homework.total,
                deadline: homework.deadline
              }))
            )
        )
        .filter(assignment => Date.now() > assignment.deadline),
    getResults: state => name =>
      state.student_courses
        .find(course => course.name === name)
        .homeworks.filter(homework => Date.now() > homework.deadline)
        .concat(
          state.student_courses
            .find(course => course.name === name)
            .labs.filter(lab => Date.now() > lab.deadline)
        ),
    getAssignmentsToGrade: state => name =>
      state.assistant_courses
        .find(course => course.name === name)
        .assignments.flatMap(assignment => ({
          name: assignment.name
        })),
    getAssignmentStudents: state => (course_name, assignment_name) =>
      state.assistant_courses
        .find(course => course.name === course_name)
        .assignments.find(assignment => assignment.name === assignment_name)
        .students
  }
});
