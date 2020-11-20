<template>
  <v-container>
    <Banner />
    <Recommendations />
  </v-container>
</template>

<script>
import Recommendations from "../components/Recommendations";
import Banner from "../components/Banner";

export default {
  name: "Home",
  components: {
    Banner,
    Recommendations,
  },
};
</script>

<style>
</style>
