import { createRouter, createWebHistory } from "vue-router";

// import your page components
import AzureDevopsSetup from "../pages/AzureDevopsSetup.vue";
import Setup from "@/pages/Setup.vue";
import Dashboard from "@/pages/Dashboard.vue";
import Home from "../pages/Home.vue";

// define all your routes
const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/setup",
    name: "Setup",
    component: Setup,
  },
  {
    path: "/azure",
    name: "azure",
    component: AzureDevopsSetup,
  },
  {
    path: "/dashboard",
    name: "dashboard",
    component: Dashboard,
  },
];

// create router instance
const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
