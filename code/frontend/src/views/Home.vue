<template>
  <v-container fluid>
    <v-parallax src="https://cdn.vuetifyjs.com/images/backgrounds/vbanner.jpg">
      <Landing />
    </v-parallax>
    <About />
  </v-container>
</template>

<script>
import Landing from "../components/home/Landing";
import About from "../components/home/About";

export default {
  name: "Home",
  components: {
    Landing,
    About,
  },
};
</script>

<style>
</style>
