<template>
  <div id="wrapperprojects">
    <router-view></router-view>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class Projects extends Vue{
  @Prop(Function)
  public routerFunc!: Function;
  
  @Prop(Array)
  public items!: Array<any>;

  mounted(){
    console.log('Projects mounted')
  }
}
</script>


<style scoped>
#wrapperprojects {
  height: 100vh;
  width: 100%;
}
</style>