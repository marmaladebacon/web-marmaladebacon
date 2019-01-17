<template>
  <div
    id="wrapperhome"
    class="home"
    @mousemove="mouseMoveHome"
    ref="containerWindow"
  >
    <div></div>
    <img
      src="../assets/pmcatEyes.png"
      ref="catEyes"
      style="position:absolute; top:0.2vh; right:1.4vw;z-index:1000; height: 17vh;"
    />
    <img
      src="../assets/pmcatBody.png"
      ref="catBody"
      style="position:absolute; top:0.2vh; right:1.4vw; height: 17vh;"
    />
    <img
      src="../assets/pmcatWindow.png"
      ref="catWindow"
      style="position:absolute; top:0.2vh; right:1vw; height: 20vh;"
    />

    <MouseTrackingLink
      v-for="(item, index) in links"
      :key="index"
      :component_index="index"
      :container_width="containerWidth"
      :container_height="containerHeight"
      :mouseX="mouseX"
      :mouseY="mouseY"
    >
      <router-link :to="item.link">{{ item.urlText }}</router-link>
    </MouseTrackingLink>
    <!-- </div> -->
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import { default as getLinks } from "@/utility/staticLinkData";
import { TweenMax } from "gsap";
import { isUndefined } from "lodash-es";
import MouseTrackingLink from "@/components/MouseTrackingLink.vue"; // @ is an alias to /src

@Component({
  name: "home",
  components: {
    MouseTrackingLink
  }
})
export default class Home extends Vue {
  page: number = 0;
  pageCount: number = 8;
  mouseX: number = 0;
  mouseY: number = 0;
  containerWidth: number = 0;
  containerHeight: number = 0;

  @Prop()
  public links!: { link: string; urlText: string }[];

  created() {
    this.links = getLinks();
  }

  mouseMoveHome(e: any) {
    this.mouseX = e.pageX;
    this.mouseY = e.pageY;
    try {
      this.containerWidth = (this.$refs.containerWindow as any).offsetWidth;
      this.containerHeight = (this.$refs.containerWindow as any).offsetHeight;

      this.parallax(e, this.$refs.catEyes, 28);
      this.parallax(e, this.$refs.catBody, 20);
      this.parallax(e, this.$refs.catWindow, 10);
    } catch (e) {
      console.error(e);
    }
  }
  parallax(e: any, target: any, movement: number) {
    var relX = e.pageX - (this.$refs.containerWindow as any).offsetLeft;
    var relY = e.pageY - (this.$refs.containerWindow as any).offsetTop;

    TweenMax.to(target, 1, {
      x:
        ((relX - (this.$refs.containerWindow as any).offsetWidth / 2) /
          (this.$refs.containerWindow as any).offsetWidth) *
        movement,
      y:
        ((relY - (this.$refs.containerWindow as any).offsetHeight / 2) /
          (this.$refs.containerWindow as any).offsetHeight) *
        movement
    });
  }
  linkStyle(index: number) {
    let frac = 0;
    if (!isUndefined(this.$refs.containerWindow)) {
      const width = (this.$refs.containerWindow as any).offsetLeft;
      frac = (this.mouseX - width) / width;
    }
    console.log(frac);
    return `top:${5 * index + 20}vh; left:${1.4 + frac * 20}vw;`;
  }
}
</script>
<style>
a:visited {
  color: #cccccc;
}

a:link {
  color: #ffffff;
}

a:hover {
  color: black !important;
}
#wrapperhome {
  padding-top: 2vh;
  position: relative;
  height: 100vh;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}
.home-links {
  position: absolute;
}
.wrapper {
  display: grid;
  grid-template-columns: 2.5% 30% 2.5% 30% 2.5% 30% 2.5%;
  grid-template-rows: auto 10px auto;
}
.a {
  grid-column: 2 / 3;
  grid-row: 1 / 4;
  border: 1px solid black;
}
.b {
  grid-column: 4 / 5;
  grid-row: 1 / 2;
}
.c {
  grid-column: 6 / 7;
  grid-row: 1 / 2;
}
.d {
  grid-column: 4 / 6;
  grid-row: 3 / 4;
}
</style>
