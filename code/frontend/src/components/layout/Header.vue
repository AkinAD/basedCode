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
      <v-tab v-show="this.signedIn" to="/account">
        <v-icon left dark> mdi-account </v-icon>
        Account
      </v-tab>

      <v-tab v-show="!this.signedIn" to="/login">
        <v-icon left dark> mdi-lock </v-icon>
        Login
      </v-tab>

      <!--Managers/employees tab-->
      <v-tab v-show="this.signedIn" to="/employee">
        <v-icon left dark> mdi-briefcase </v-icon>
        Manage
      </v-tab>

      <v-btn v-show="this.signedIn" class="ma-2" text icon color="grey">
        <StoreSelector />
      </v-btn>

      <v-btn class="ma-2" text icon>
        <v-icon>mdi-cart</v-icon>
      </v-btn>
    </v-tabs>
  </v-card>
</template>

<script>
import { mapGetters } from "vuex";
import StoreSelector from "../../components/browse/StoreSelector";

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
    ...mapGetters(["signedIn"]),
  },
};
</script>

<style scoped></style>
