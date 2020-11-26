<template>
  <v-container>
    <v-app-bar elevate-on-scroll app>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-app-bar-title>
        <v-container> Smart Shopping </v-container>
      </v-app-bar-title>

      <v-tabs grow @change="toTop">
        <!--Tabs that don't require authentication-->
        <v-tab v-for="item in menuItems" :key="item.title" :to="item.path">
          <v-icon left dark>{{ item.icon }}</v-icon>
          {{ item.title }}
        </v-tab>

        <!--Tabs that do require authentication-->
        <v-tab v-show="this.signedIn" to="/account">
          <v-icon left dark> mdi-account </v-icon>
          Account
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

        <v-btn class="ma-2" icon>
          <StoreSelector />
        </v-btn>

        <v-btn class="ma-2" icon v-show="signedIn">
          <v-icon>mdi-cart</v-icon>
        </v-btn>

        <amplify-sign-out class="ma-2" v-show="signedIn"></amplify-sign-out>
      </v-tabs>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" app temporary>
      <!--  -->
    </v-navigation-drawer>
  </v-container>
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
      drawer: null,
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
  methods: {
    toTop() {
      this.$vuetify.goTo(0);
    },
  },
};
</script>

<style scoped>
</style>
