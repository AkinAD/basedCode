<template>
  <v-app>
    <!-- main app tag -->
    <Header />
    <v-main
      ><!--where pages are to be loaded-->
      <router-view />
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
  components: {
    Header,
    Footer,
  },
  computed: {},
  beforeCreate() {
    Hub.listen("auth", (data) => {
      const { payload } = data;
      if (payload.event === "signIn") {
        this.signIn();
        this.$router.push("/account");
      }
      if (payload.event === "signOut") {
        this.signOut();
        this.$router.push("/home");
      }
    });
    Auth.currentAuthenticatedUser()
      .then(() => {
        this.signIn();
      })
      .catch(() => {
        this.signOut();
      });
  },
  methods: {
    ...mapActions(["signIn", "signOut"]),
  },
};
</script>

<style></style>
