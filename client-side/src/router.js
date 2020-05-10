import Vue from "vue";
import Router from "vue-router";
import Home from "./views/frontpage/Home.vue";

Vue.use(Router);

export default new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      component: Home
    },
    {
      path: "/about",
      component: () => import("./views/frontpage/About.vue")
    },
    {
      path: "/contact",
      component: () => import("./views/frontpage/Contact.vue")
    },
    {
      path: "/login",
      component: () => import("./views/frontpage/Login.vue")
    },
    {
      path: "/dashboard",
      component: () => import("./views/subject/Dashboard.vue")
    },
    {
      path: "/subjectoverview",
      component: () => import("./views/subject/SubjectOverview.vue")
    },
    {
      path: "/subject/:name",
      component: () => import("./views/subject/Subject.vue")
    },
    {
      path: "/results",
      component: () => import("./views/results/Results.vue")
    },
    {
      path: "/gradeoverview",
      component: () => import("./views/grade/GradeOverview.vue")
    },
    {
      path: "/grade/:course/:series/:year",
      component: () => import("./views/grade/Grade.vue")
    },
    {
      path: "/grade/:course/:series/:year/:assignment",
      component: () => import("./views/grade/GradeAssignment.vue")
    },
    {
      path: "/profile",
      component: () => import("./views/profile/Profile.vue")
    },
    {
      path: "/workspace/:course/:series/:year/:assignmentname/:owner",
      component: () => import("./views/workspace/Workspace.vue")
    }
  ]
});
