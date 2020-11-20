<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png" />
    <Banner />
    <Recommendations />
    <Footer />
  </div>
</template>

<script>
// @ is an alias to /src

import Recommendations from "../components/Recommendations";
import Banner from "../components/Banner";
import Footer from "../components/layout/Footer";
export default {
  name: "Home",
  components: {
    Banner,
    Recommendations,
    Footer,
  },
};
</script>

<style>
.home {
  text-align: center;
}
</style>
