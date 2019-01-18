<template>
  <div>    
    <iframe
      :src="projectSrc"
      title="project iframe"
      :width="width"
      :height="height"
      scrolling="no"
      frameborder="0"
      class="iframe-custom"
    >
      <p>Your browser does not support iframes.</p>
    </iframe>
    <div class="desc-text" :style="style">
      <p>{{desc_text}}</p>
    </div>
  </div>

</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component({
  name: 'ProjectContainer',
})
export default class ProjectContainer extends Vue {
  @Prop(String)
  public projectSrc!:string;

  @Prop(Number)
  public width!:number;

  @Prop(Number)
  public height!:number;

  @Prop(Number)
  public desc_width!:number;

  @Prop(String)
  public desc_text!:String;


  get style(){
    return `z-index:500; text-align:left; position:absolute;top:1vh;right:1vw;width:${this.desc_width}vw; `;
  }
}
</script>

<style>
p {
  color: aqua;
  padding-left: 1vw;
  padding-right: 1vw;
}

.iframe-custom {
  float:left;
  margin-left: 15vw;
}
</style>



