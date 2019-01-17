<template>
  <div class="links" :style="style"><slot></slot></div>
</template>

<script lang="ts">
import { isNaN, isUndefined } from "lodash-es";
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class MouseTrackingLink extends Vue {
  @Prop(Number)
  public component_index!: number;

  @Prop(Number)
  public container_width!: number;

  @Prop(Number)
  public container_height!: number;

  @Prop(Number)
  public mouseX!: number;

  @Prop(Number)
  public mouseY!: number;

  get style() {
    let frac = 0;
    const baseTop = 35;
    if (!isUndefined(this.container_width)) {
      frac = (this.container_width - this.mouseX) / this.container_width;
      frac = frac < 0.4 ? 0 : frac - 0.4;
    }
    if (isNaN(frac)) {
      const val = 5 * this.component_index + baseTop;
      return `top:${val}vh; left:4vw`;
    }

    const maxHeightDiff = 15;
    const hPos =
      ((5 * this.component_index + baseTop) * this.container_height) / 100;

    const heightabsolutefrac =
      frac == 0
        ? 0
        : (this.container_height - Math.abs(hPos - this.mouseY)) /
          this.container_height;

    const direction = hPos > this.mouseY ? 1 : -1;
    let heightCtrl =
      ((hPos - this.mouseY) / this.container_height) * maxHeightDiff;
    heightCtrl = Math.abs(heightCtrl) > 1.5 ? 1.5 * direction : heightCtrl;

    return `top:${5 * this.component_index + baseTop + heightCtrl}vh; left:${2 +
      heightabsolutefrac * heightabsolutefrac * heightabsolutefrac * frac * 15 +
      frac * frac * frac * frac * 5}vw`;
  }
}
</script>

<style scoped>
a:visited {
  color: #cccccc;
}

a:link {
  color: #ffffff;
}

a:hover {
  color: black !important;
}
.links {
  position: absolute;
  font-family: "Courier New", Courier, monospace;
}
</style>
