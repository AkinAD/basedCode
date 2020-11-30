<template>
  <v-container fluid>
    <Banner text="Welcome to Smart Shopping" resize/>
    <Recommendations />
  </v-container>
</template>

<script>
import Recommendations from "../components/Recommendations";
import Banner from "../components/layout/Banner";

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
