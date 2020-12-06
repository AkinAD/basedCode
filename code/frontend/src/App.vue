<template>
  <v-app>
    <!-- main app tag -->
    <Header class="v-navigation-drawer" />
    <v-main
      ><!--where pages are to be loaded-->
      <v-container fluid>
        <router-view />
      </v-container>
      <v-container>
        <!-- <v-btn
          v-scroll="onScroll"
          v-show="fab"
          fab
          dark
          fixed
          bottom
          right
          color="primary"
          @click="toTop"
        >
          <v-icon>mdi-chevron-up</v-icon>
        </v-btn> -->
      </v-container>
    </v-main>
    <Footer />
  </v-app>
</template>

<script>
import Header from "./components/layout/Header";
import Footer from "./components/layout/Footer";

import { Hub } from "aws-amplify";
import { mapActions } from "vuex";
import { Auth } from "aws-amplify";

export default {
  name: "app",
  data() {
    return {
      fab: false,
    };
  },
  components: {
    Header,
    Footer,
  },
  beforeCreate() {
    Hub.listen("auth", (data) => {
      const { payload } = data;
      if (payload.event === "signIn") {
        this.signIn();
        this.$router.push("/home");
      }
      if (payload.event === "signOut") {
        this.signOut();
        this.$router.push("/home");
      }
    });
    Auth.currentAuthenticatedUser()
      .then(() => {
        this.signIn();
        this.$router.push("/browse");
      })
      .catch(() => {
        this.signOut();
        this.$router.push("/home");
      });
  },
  created() {
    this.updateStores();
    this.fetchCategories();
    this.fetchAllItems();
  },
  methods: {
    ...mapActions([
      "signIn",
      "signOut",
      "updateStores",
      "fetchCategories",
      "fetchEmployees",
      "fetchAllItems",
    ]),
    onScroll(e) {
      if (typeof window === "undefined") return;
      const top = window.pageYOffset || e.target.scrollTop || 0;
      this.fab = top > 20;
    },
    toTop() {
      this.$vuetify.goTo(0);
    },
  },
};
</script>

<style>
.v-navigation-drawer {
  z-index: 999999 !important;
}
</style>
