<template>
  <v-container>
    <Banner :text="msg" size="100px" />
    <v-row> </v-row>
  </v-container>
</template>

<script>
import Banner from "../components/layout/Banner";
import { mapGetters } from "vuex";

export default {
  name: "account",
  components: {
    Banner,
  },
  computed: {
    ...mapGetters(["getUser"]),
    msg() {
      try {
        return "Hello, ".concat(`${this.getUser.attributes.email}`);
      } catch (err) {
        return "";
      }
    },
  },
};
</script>

<style></style>
