<template>
  <v-app> <!-- main app tag -->
    <Header />
    <v-content> <!--where pages are to be loaded-->
      <router-view></router-view>
    </v-content>
  </v-app>
</template>

<script>
import Header from "./components/layout/header";
export default {
  name: "app",
  components: {
    Header,
  },
};
</script>

<style>

</style>