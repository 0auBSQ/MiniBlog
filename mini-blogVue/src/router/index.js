import Vue from "vue";
import VueRouter from "vue-router";
//import Home from "../views/Home.vue";
import AccountInfo from "../views/AccountInfo.vue";
import Admin from "../views/Admin.vue";

Vue.use(VueRouter);


const routes = [
    {
        path: "/",
        name: "Home",
        component: () => import("../views/Account"), //Home,
        alias: '/home'
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
        path: '/account',
        component : () => import("../views/Account.vue"),
        children: [{
            path: 'userinfo',
            components: {
                info: AccountInfo
            }
        },{
            path: 'admin', 
            components: {
                info: Admin
            }
        }]
    },
    {
        path: "*",
        component: () => import("../views/Missing.vue")
    },
    
]

const router = new VueRouter({
    mode: "history",
    base: process.env.BASE_URL,
    routes
});

router.beforeEach((to,from,next) => {
    if(to.matched.some(route => route.meta.requiresAuth)){
      if(Vue.$store.state.Auth.is_auth_token) return next();
  
      return next('/non-auth-required-route');
    }
  
    next();
  });

export default router;