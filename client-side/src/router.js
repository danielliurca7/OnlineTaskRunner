import Vue from "vue";
import Router from "vue-router";
import Home from "./views/Home.vue";

Vue.use(Router);

export default new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "home",
      component: Home
    },
    {
      path: "/about",
      name: "about",
      component: () => import("./views/About.vue")
    },
    {
      path: "/contact",
      name: "contact",
      component: () => import("./views/Contact.vue")
    },
    {
      path: "/login",
      name: "login",
      component: () => import("./views/Login.vue")
    },
    {
      path: "/dashboard",
      name: "dashboard",
      component: () => import("./views/Dashboard.vue")
    },
    {
      path: "/subjectoverview",
      name: "subjectoverview",
      component: () => import("./views/SubjectOverview.vue")
    },
    {
      path: "/subject/:name",
      name: "subject",
      component: () => import("./views/Subject.vue")
    },
    {
      path: "/results",
      name: "results",
      component: () => import("./views/Results.vue")
    },
    {
      path: "/gradeoverview",
      name: "gradeoverview",
      component: () => import("./views/GradeOverview.vue")
    },
    {
      path: "/grade/:name",
      name: "grade",
      component: () => import("./views/Grade.vue")
    },
    {
      path: "/grade/:subject/:assignment",
      name: "gradeassignment",
      component: () => import("./views/GradeAssignment.vue")
    },
    {
      path: "/profile",
      name: "profile",
      component: () => import("./views/Profile.vue")
    },
    {
      path: "/lab/:name",
      name: "lab",
      component: () => import("./views/Lab.vue")
    },
    {
      path: "/homework/:name",
      name: "homework",
      component: () => import("./views/Homework.vue")
    }
  ]
});
