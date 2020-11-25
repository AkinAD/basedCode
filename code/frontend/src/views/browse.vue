<template>
  <v-container v-show="nonNull" fluid transition="slide-x-transition">
    <Banner :text="msg" size="100px" />
    <v-row>
      <v-col md="3" offset-lg="1"><Sidebar /></v-col>
      <v-col md="9" lg="7"> <ProductDisplay /> </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Sidebar from "../components/Sidebar";
import ProductDisplay from "../components/browse/ProductDisplay";
import Banner from "../components/layout/Banner";
import { mapGetters, mapMutations } from "vuex";

export default {
  name: "browse",
  components: {
    Sidebar,
    ProductDisplay,
    Banner,
  },
  mounted() {
    if (this.getSelectedStore === null) {
      this.setDialog(true);
    }
  },
  computed: {
    ...mapGetters(["getSelectedStore"]),
    msg() {
      try {
        return "Products from ".concat(`${this.getSelectedStore.text}`);
      } catch (err) {
        return "";
      }
    },
    nonNull() {
      return this.getSelectedStore !== null;
    },
  },
  methods: {
    ...mapMutations(["setDialog"]),
  },
};
</script>

<style></style>
