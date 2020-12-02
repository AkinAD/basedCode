<template>
  <v-app-bar elevate-on-scroll app>
    <v-app-bar-nav-icon></v-app-bar-nav-icon>

    <v-avatar>
      <v-img :src="img"></v-img>
    </v-avatar>

    <v-tabs grow>
      <!--Tabs that don't require authentication-->
      <v-tab flat v-for="item in menuItems" :key="item.title" :to="item.path">
        <v-icon left dark>{{ item.icon }}</v-icon>
        {{ item.title }}
      </v-tab>

      <!--Tabs that do require authentication-->
      <v-tab v-show="this.signedIn" to="/cart">
        <v-icon left dark> mdi-cart </v-icon>
        Cart
      </v-tab>

      <v-tab v-show="!this.signedIn" to="/login">
        <v-icon left dark> mdi-lock </v-icon>
        Login
      </v-tab>

      <!--Managers/employees tab-->
      <v-tab v-show="isManagerOrEmployee" to="/employee">
        <v-icon left dark> mdi-briefcase </v-icon>
        Manage
      </v-tab>

      <v-btn class="ma-2" text icon>
        <StoreSelector />
      </v-btn>

      <amplify-sign-out class="ma-2" v-show="signedIn"></amplify-sign-out>
    </v-tabs>
  </v-app-bar>
</template>

<script>
import StoreSelector from "../../components/browse/StoreSelector";
import { mapGetters } from "vuex";

export default {
  name: "Header",
  components: {
    StoreSelector,
  },
  data() {
    //needs to be modified to be dynamic based on router links maybe
    return {
      menuItems: [
        { title: "Home", path: "/home", icon: "mdi-home" },
        { title: "Browse", path: "/browse", icon: "mdi-book" },
      ],
      img: require("./braincart.png"),
    };
  },
  computed: {
    ...mapGetters(["signedIn", "getUserGroups"]),
    isManagerOrEmployee() {
      return (
        this.getUserGroups.includes("manager") ||
        this.getUserGroups.includes("employee")
      );
    },
  },
};
</script>

<style scoped>
</style>
