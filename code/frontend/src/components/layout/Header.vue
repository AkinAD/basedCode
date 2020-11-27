<template>
  <v-card>
    <v-card-title class="text-center justify-center">
      <h1 class="text-center display-3 hidden-xs-only" style="color: grey">
        Smart
      </h1>
      <h1 class="text-center display-3 hidden-xs-only">Shopping</h1>
      <h1 class="text-center display-1 hidden-sm-and-up" style="color: grey">
        Smart
      </h1>
      <h1 class="text-center display-1 hidden-sm-and-up">Shopping</h1>
      <v-icon large color="black"> mdi-cart-outline </v-icon>
    </v-card-title>
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

      <v-btn class="ma-2" text icon>
        <StoreSelector />
      </v-btn>

      <amplify-sign-out v-show="signedIn"></amplify-sign-out>
    </v-tabs>
  </v-card>
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
