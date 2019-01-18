import Vue from "vue";
import Router from "vue-router";
import Home from "./views/Home.vue";
import ProjectContainer from "./components/ProjectContainer.vue";
Vue.use(Router);
Vue.config.devtools = true;
export default new Router({
  routes: [
    {
      path: "/",
      name: "home",
      component: Home
    },
    {
      path: "/about",
      name: "about",
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () =>
        import(/* webpackChunkName: "about" */ "./views/About.vue")
    },
    {
      path: "/projects/",
      name: "projects",
      component: () => import("./views/Projects.vue"),
      props: {
        routerFunc: function(item:any) {
          const projectSrc = item.path ;
          console.log('Item Path:' + projectSrc);
          return {
            path: "/" + item.path,
            query: { projectSrc }
          };
        },
        items: [
          {
            path: "arch-reborn-web",
            label: "arch-reborn"
          },
          {
            path: "tensorflow-linear-regression",
            label: "tensorflow linear regression"
          }
        ]
      },
      children: [
        {
          path: "tensorflow-linear-regression",
          component: ProjectContainer,
          props(route) {
            console.log('Mounted tf');
            return {
              projectSrc: "/tensorflow-linear-regression",
              desc_width: 20,
              desc_text: "This was a learning exercise i've learnt step by step from The Coding Train https://www.youtube.com/watch?v=dLp10CFIvxI . Would definitely invest more time in learning tensorflow js in the future."
            }
          }
        },
        {
          path: "arch-reborn-web",
          component: ProjectContainer,
          props(route) {            
            return {
              projectSrc: "/arch-reborn-web",
              width: 950,
              height: 650,
              desc_width: 17,
              desc_text: "I think I've definitely overreached on this project. Though I do think many games lack the multi-tile items from diablo 1 and 2 and inventory management, and more can be done with that mechanic similar making character builds with it. I would like to revisit this someday especially after sorting out the overly complicated turn system i implemented."
            };
          }
        },
        {
          path: "herding-cats",
          component: ProjectContainer,
          props(route) {
            return {
              //save the cats was the initial name
              projectSrc:"/save-the-cats",
              width: 370,
              height: 650,
              desc_width: 25,
              desc_text: "This is currently my ongoing project. It incorporates flocking behaviour for the cats, and allows the cats to be influenced by external forces as well. Code wise i've been keeping the scope down and concentrating on the essential player input and core puzzling mechanics instead of adding more features. The overmap is meant to elicit some decision making from players, but it is not quite done yet. I'm not much of an artist so most of it is pixel art."
            }
          }
        }
      ]
    }
  ]
});
