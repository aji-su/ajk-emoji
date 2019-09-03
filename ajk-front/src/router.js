import Vue from "vue";
import Router from "vue-router";
import Upload from "./views/Upload.vue";
import Show from "./views/Show.vue";

Vue.use(Router);

export default new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "upload",
      component: Upload
    },
    {
      path: "/show/:requestId",
      name: "show",
      component: Show
    }
  ]
});
