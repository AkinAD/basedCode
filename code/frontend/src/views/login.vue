<template>
  <v-container fluid>
    <v-row align="center" justify="center">
      <amplify-authenticator></amplify-authenticator>
    </v-row>
  </v-container>
</template>

<script>
import { Auth } from "aws-amplify";
import { AmplifyEventBus } from "aws-amplify-vue";

export default {
  name: "login",
  computed: {
    signedIn() {
      return this.$store.state.signedIn;
    },
  },
  created() {
    this.findUser();
    AmplifyEventBus.$on("authState", (info) => {
      if (info === "signedIn") {
        this.findUser();
      } else {
        this.$store.state.signedIn = false;
        this.$state.user = null;
      }
    });
  },
  components: {},
  methods: {
    async findUser() {
      try {
        const user = await Auth.currentAuthenticatedUser();
        this.$store.state.signedIn = true;
        this.$store.state.user = user;
        this.$router.push('/account')
      } catch (err) {
        this.$store.state.signedIn = false;
        this.$store.state.user = null;
      }
    },
  },
};
</script>

<style>
</style>