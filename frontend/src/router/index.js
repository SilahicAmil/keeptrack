import { createRouter, createWebHistory } from "vue-router";

// import your page components
import Home from "../pages/Home.vue";
import AzureDevopsSetup from "../pages/AzureDevopsSetup.vue";
import Setup from "@/pages/Setup.vue";

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
];

// create router instance
const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
