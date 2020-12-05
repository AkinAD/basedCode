<template>
  <v-app-bar elevate-on-scroll app>
    <span class="hidden-sm-and-up">
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"> </v-app-bar-nav-icon>
    </span>

    <v-avatar>
      <v-img :src="img"></v-img>
    </v-avatar>

    <v-tabs class="hidden-xs-only" grow>
      <!--Tabs that don't require authentication-->
      <v-tab flat v-for="item in menuItems" :key="item.title" :to="item.path">
        <v-icon left dark>{{ item.icon }}</v-icon>
        {{ item.title }}
      </v-tab>

      <!--Tabs that do require authentication-->
      <v-tab v-show="signedIn" to="/cart">
        <v-icon left dark> mdi-cart </v-icon>
        Cart
      </v-tab>

      <v-tab v-show="!signedIn" to="/login">
        <v-icon left dark> mdi-lock </v-icon>
        Login
      </v-tab>

      <!--Managers/employees tab-->
      <v-tab v-show="isManagerOrEmployee" to="/employee">
        <v-icon left dark> mdi-briefcase </v-icon>
        Manage
      </v-tab>

      <v-tooltip bottom>
        <template v-slot:activator="{ on }">
          <v-btn class="ma-2" text icon v-on="on">
            <OptimumRoute />
          </v-btn>
        </template>
        <span>Click to see optimum route.</span>
      </v-tooltip>

      <v-tooltip bottom>
        <template v-slot:activator="{ on }">
          <v-btn class="ma-2" text icon v-on="on">
            <StoreSelector />
          </v-btn>
        </template>
        <span>Click to select a store.</span>
      </v-tooltip>

      <v-tooltip bottom>
        <template v-slot:activator="{ on }">
          <v-btn
            class="ma-2"
            color="red"
            v-show="signedIn"
            icon
            v-on="on"
            @click="signOut"
          >
            <v-icon>mdi-logout</v-icon>
          </v-btn>
        </template>
        <span>Click to logout.</span>
      </v-tooltip>
    </v-tabs>
    <!-- Show only on small devices -->
    <v-navigation-drawer
      class="v-navigation-drawer"
      v-model="drawer"
      disable-resize-watcher
      absolute
      temporary
      left
      elevate-on-scroll
    >
      <v-list>
        <v-list-item
          v-for="item in menuItems"
          :key="item.title"
          :to="item.path"
          link
        >
          <v-list-item-icon
            ><v-icon left dark>{{ item.icon }}</v-icon></v-list-item-icon
          >
          {{ item.title }}
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </v-app-bar>
</template>

<script>
import StoreSelector from "../../components/browse/StoreSelector";
import { mapGetters } from "vuex";
import OptimumRoute from "../../components/OptimumRoute";
import { Auth } from "aws-amplify";

export default {
  name: "Header",
  components: {
    StoreSelector,
    OptimumRoute,
  },
  data() {
    //needs to be modified to be dynamic based on router links maybe
    return {
      drawer: false,
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
  methods: {
    signOut() {
      Auth.signOut();
    },
  },
  watch: {
    signedIn() {},
  },
};
</script>

<style scoped></style>
