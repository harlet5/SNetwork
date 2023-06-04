import { createRouter, createWebHistory } from "vue-router";
import { useAuthStore } from "@/stores/auth";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Home",
      component: () => import("../views/HomeView.vue"),
    },
    {
      path: "/posts",
      name: "Posts",
      component: () => import("../views/PostsView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/posts/:id",
      name: "Postdetails",
      component: () => import("../views/PostView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/register",
      name: "Register",
      component: () => import("../views/RegisterView.vue" as string),
      meta: { guest: true },
    },
    {
      path: "/login",
      name: "Login",
      component: () => import("../views/LoginView.vue"),
      meta: { guest: true },
    },
    {
      path: "/logout",
      name: "Logout",
      component: () => import("../views/LogoutView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/chat",
      name: "Chat",
      component: () => import("../views/ChatView.vue"),
      meta: { requiresAuth: true },
      //props: (route) => ({ foo: route.query.foo }),
    },
    {
      path: "/profile",
      name: "Profile",
      component: () => import("../views/ProfileView.vue"),
      meta: { requiresAuth: true },
      //props: (route) => ({ foo: route.query.foo }),
    },
    {
      path: "/user/:id",
      name: "UserProfile",
      component: () => import("../views/ProfileView.vue"),
      meta: { requiresAuth: true },
      //props: (route) => ({ foo: route.query.foo }),
    },
    {
      path: "/groups",
      name: "Groups",
      component: () => import("../views/GroupsView.vue"),
      meta: { requiresAuth: true },
      //props: (route) => ({ foo: route.query.foo }),
    },
    {
      path: "/groups/:id",
      name: "Groupdetails",
      component: () => import("../views/GroupView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/groups/posts/:id",
      name: "GroupPostdetails",
      component: () => import("../views/PostView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/user/posts/:id",
      name: "UserPostDetails",
      component: () => import("../views/PostView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/:catchAll(.*)",
      name: "notFound",
      component: () => import("../views/ErrorView.vue"),
    },
  ],
});

router.beforeResolve((to) => {
  if (to.meta.requiresAuth && !useAuthStore().loggedIn) {
    return { name: "Login" };
  } else if (to.meta.guest && useAuthStore().loggedIn) {
    return { name: "Home" };
  } else {
    return true;
  }
});

export default router;
