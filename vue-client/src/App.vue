<template>
  <div
    id="app"
    class="app-class"
    :style="style"
    @mousemove="mouseMoveApp"
    ref="appWindow"
  >
    <div v-show="showBack" class="back-link">
      <router-link to="/">&lt;</router-link>
    </div>
    <transition name="slide-fade" mode="out-in"> <router-view /> </transition>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Watch } from "vue-property-decorator";
import {isUndefined} from 'lodash-es';
import RGB from "./types/RGB";

@Component
export default class App extends Vue {
  start: RGB = { r: 0, g: 0, b: 0 };
  end: RGB = { r: 255, g: 255, b: 255 };
  mod: RGB = { r: 0, g: 0, b: 0 };
  transitionName: string = "";

  created() {
    this.start = {
      r: this.rndInt(255),
      g: this.rndInt(255),
      b: this.rndInt(255)
    };
    this.end = {
      r: this.modInt(this.start.r, 125, 255),
      g: this.modInt(this.start.g, 125, 255),
      b: this.modInt(this.start.b, 125, 255)
    };
    this.mod = {
      r: 0,
      g: 0,
      b: 0
    };
  }

  rndInt(max: number) {
    return Math.floor(Math.random() * max);
  }

  modInt(start: number, toAdd: number, modBy: number) {
    return (start + toAdd) % modBy;
  }

  mouseMoveApp(e: any) {
    const offsetLeft = (this.$refs.appWindow as any).clientWidth;
    const offsetTop = (this.$refs.appWindow as any).clientHeight;
    
    const relX = Math.floor(((e.pageX - offsetLeft) / offsetLeft) * 80);
    const relY = Math.floor(((e.pageY - offsetTop) / offsetTop) * 80);

    this.mod.r = relX;
    this.mod.g = relY;
    this.mod.g = Math.floor((relX + relY) / 4);    
  }

  norm(val: number, min: number, max: number) {
    return val > max ? max : val < min ? min : val;
  }
  
  @Watch("$route", { immediate: true, deep: true })
  onRouteChange(val: any, oldVal: any) {
    if(isUndefined(oldVal)){
      return;
    }
    const toDepth = oldVal.path.split("/").length;
    const fromDepth = val.path.split("/").length;
    this.transitionName = toDepth < fromDepth ? "slide-right" : "slide-left";
  }
  
  get showBack() {
    if (this.$route.path == "/home" || this.$route.path == "/") {
      return false;
    } else {
      return true;
    }
  }

  get style() {
    const startR = this.norm(this.start.r + this.mod.r, 0, 255);
    const startG = this.norm(this.start.g + this.mod.g, 0, 255);
    const startB = this.norm(this.start.b + this.mod.b, 0, 255);

    const endR = this.norm(this.end.r + this.mod.r, 0, 255);
    const endG = this.norm(this.end.g + this.mod.g, 0, 255);
    const endB = this.norm(this.end.b + this.mod.b, 0, 255);
    const result =
      "background: linear-gradient(140deg, rgb(" +
      startR +
      "," +
      startG +
      "," +
      startB +
      "), rgb(" +
      endR +
      "," +
      endG +
      "," +
      endB +
      "))";
    
    return result;
  }
}
</script>

<style>
.container {
  position: fixed;
  left: 0;
  top: 0;
  height: 100vh;
  width: 100vw;
  background-color: #820263;
}

.app-class {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.back-link {
  z-index: 1000;
  font-size: 48px;
  font-family: "Courier New", Courier, monospace;
  float: left;
  padding-left: 5vw;
  position:fixed
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}

.slide-fade-enter-active {
  transition: all .3s ease;
}
.slide-fade-leave-active {
  transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}
.slide-fade-enter, .slide-fade-leave-to
/* .slide-fade-leave-active below version 2.1.8 */ {
  transform: translateX(10px);
  opacity: 0;
}
</style>
