import Vue from "vue";
import VueRouter from "vue-router";
//import store from "../store";
import Home from "../views/Home.vue";
import AccountInfo from "../views/AccountInfo.vue";
import Admin from "../views/Admin.vue";

Vue.use(VueRouter);


/*const isNotAuth = (to, from, next) => {
    if(store.getters.isAuth){
        next()
        return
    }
    next("/account")
}*/

/*const isAuth = (to, from, next) => {
    if(store.getters.isAuth) {
        next()
        return
    }
    next("/login")
}*/

const routes = [
    {
        path: "/",
        name: "Home",
        component: Home,
        alias: '/home'
    },
    {
        path: "/login",
        name: "LoginPage",
        component: () => import("../views/LoginPage.vue"),
        
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
        path: "/tmp",
        name: "Tmp",
        component: () => import("../views/Tmp.vue")
    },
    {
        path: '/account',
        component : () => import("../views/Account.vue"),
        meta: { requiresAuth: true },
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
        path: '/article/:id',
        name: 'article',
        component: () => import("../views/Post.vue"),
        props: true
    },
    {
        path: '/addarticle',
        name: 'addarticle',
        component: () => import("../views/AddArticle.vue"),
        meta: { requiresAdmin: true },
        
    },
    {
        path: '/editArticle/:id',
        name: 'editArticle',
        component: () => import("../views/EditArticle.vue"),
        meta: { requiresAdmin: true },
        props: true


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
/*
router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAuth)) {
     
      if (store.getters.isAuth != true) {
        next({
          path: '/login',
          query: { redirect: to.fullPath }
        })
      } else {
        next()
      }
    } else {
      next() 
    }
})

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAdmin)) {
     
      if (store.getters.isAdmin != true) {
        next({
          path: '/',
          query: { redirect: to.fullPath }
        })
      } else {
        next()
      }
    } else {
      next() 
    }
})*/



export default router;