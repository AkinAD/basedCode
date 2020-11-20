<template>
  <v-card>
    <v-card-title class="text-center justify-center">
      <h1 class="text-center display-3 hidden-xs-only">Smart Shopping</h1>
      <h1 class="text-center display-1 hidden-sm-and-up">Smart Shopping</h1>
    </v-card-title>
    <v-tabs grow>
      <v-tabs-slider></v-tabs-slider>
      <v-tab flat v-for="item in menuItems" :key="item.title" :to="item.path">
        <v-icon left dark>{{ item.icon }}</v-icon>
        {{ item.title }}
      </v-tab>
    </v-tabs>
  </v-card>
</template>

<script>
export default {
  name: "Header",
  data() {
    //needs to be modified to be dynamic based on router links maybe
    return {
      menuItems: [
        { title: "Home", path: "/home", icon: "mdi-home" },
        { title: "Browse", path: "/browse", icon: "mdi-book" },
        { title: "Account", path: "/account", icon: "mdi-account" }, //should be if logged in only
        { title: "login", path: "/login", icon: "mdi-lock" },
      ],
    };
  },
};
</script>

<style scoped>
</style>