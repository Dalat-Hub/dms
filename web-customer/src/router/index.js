import { createRouter, createWebHashHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";

const routes = [
  {
    path: "/",
    name: "home",
    component: HomeView,
  },
  {
    path: "/about",
    name: "about",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/AboutView.vue"),
  },
  {
    path: "/van-ban",
    name: "document",
    // route level code-splitting
    // this generates a separate chunk (document.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "document" */ "../views/HomeView.vue"),
  },
  {
    path: "/van-ban/:id",
    name: "document-detail",
    // route level code-splitting
    // this generates a separate chunk (document-detail.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(
        /* webpackChunkName: "document-detail" */ "../views/DocumentView.vue"
      ),
  },
  {
    path: "/dang-nhap",
    name: "dms-login",
    component: () =>
      import(/* webpackChunkName: "dms-login" */ "../views/LoginView.vue"),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
