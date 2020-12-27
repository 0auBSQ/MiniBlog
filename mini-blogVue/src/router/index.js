import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";

Vue.use(VueRouter);


const routes = [
    {
        path: "/",
        name: "Home",
        component: Home
    },
    {
        path: "/login",
        name: "LoginPage",
        component: () => import("../views/LoginPage.vue")
    },
    {
        path: "/register",
        name: "Register",
        component: () => import("../views/RegisterPage.vue")
    },
    {
        path: "/news",
        name: "News",
        component: () => import("../views/News.vue")
    },
    {
        path: "*",
        component: () => import("../views/Missing.vue")
    }
]

const router = new VueRouter({
    mode: "history",
    base: process.env.BASE_URL,
    routes
});

export default router;