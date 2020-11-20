<template>
  <v-app>
    <!-- main app tag -->
    <Header />
    <v-content><!--where pages are to be loaded-->
      <router-view/>
    </v-content>
    <Footer />
  </v-app>
</template>

<script>
import Header from "./components/layout/header";
import Footer from "./components/layout/Footer";

export default {
  name: "app",
  components: {
    Header,
    Footer
  },
};
</script>

<style></style>
