<template>
  <v-container fluid>
    <Landing />
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
