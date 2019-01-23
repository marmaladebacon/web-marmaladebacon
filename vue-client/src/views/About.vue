<template>
  <div class="about" id="wrapperabout">
    <h1>Stuff about me</h1>
    <div style="text-align:left">
      <p>
        Hi, I am Matthew Ko. I'm a coder that enjoys creative design work and web development. Currently working at a company that is a competitor to Bloomberg and Thomson Reuters, my most recent work in my current job role has involved JavaScript/TypeScript and NodeJs, whilst my side projects involved C# and Unity.
      </p>
      <p>
        While most of my work is within the company's enterprise GitHub, I have put up some of my libraries that my hobby projects use on GitHub:
      </p>
      <ul>
        <li><a target="_blank" href="https://github.com/marmaladebacon/mb-ai-movement">Steering behaviours 2d</a></li>
        <li><a target="_blank" href="https://github.com/marmaladebacon/mb-x-input">Cross platform touch input</a></li>
        <li><a target="_blank" href="https://github.com/marmaladebacon/yet-another-ooo">Work in progress Github App</a></li>
      </ul>
      <p>
        Some of my older side projects which I've recorded but have been unable to keep up with the unity upgrades over the years can be seen on youtube below, while the more recent ones have been made into webgl builds for this site:
      </p>
      <ul>
        <li><a target="_blank" href="https://www.youtube.com/watch?v=XKNyxVTYF0A">Boids experiment</a></li>
        <li><a target="_blank" href="https://www.youtube.com/watch?v=2SFfLb3z6-Y">Choose your army composition and let the AI battle it out</a></li>
        <li><a target="_blank" href="https://www.youtube.com/watch?v=QeOvbmYZ5lM">Hex grid A* pathing experiment</a></li>
        <li><a target="_blank" href="https://www.youtube.com/watch?v=esLAHQrWJHQ">(Semi)-radial menu for a real time/slow motion game</a></li>
        <li><a target="_blank" href="https://www.youtube.com/watch?v=2eoDZVxfm8o">Grid based dark soulsalike, Josh Holinaty lent me some of his ground textures for this. I wasn't satisfied with how I've implemented some of the code</a></li>
        <li><a href="https://www.youtube.com/watch?v=W9J2hEBPh_s">Same as the above, but trying to nail the combat mechanics</a></li>
      </ul>
      <br />
    </div>
    <div>
      <p>
        Contact me via DM initially at
        <a href="https://twitter.com/marmaladebacon">@marmaladabacon</a>
      </p>
    </div>
  </div>
</template>
<script>
import ProjectContainer from "@/components/ProjectContainer.vue";

export default {
  name: "About",
  components: {
    ProjectContainer
  }
};
</script>

<style scoped>
h1 {
  color: aqua;
}
p {
  color: aqua;
  padding-left: 1vw;
  padding-right: 1vw;
}
.wrapperabout{
  position: relative;
  height: 100vh;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}
#wrapperabout {
  padding-left: 10vw;
  padding-right: 20vw;
  padding-top: 2vh;
  position: relative;
  height: 100vh;
  width: 100%;
  display: flow-root;
  justify-content: center;
  align-items: center;
}
</style>