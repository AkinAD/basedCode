<template>
  <v-container v-show="storeIsSelected" fluid>
    <Banner :text="msg" size="100px" />
    <v-row>
      <v-col md="3" offset-lg="1">
        <ProductFilter />
        <Recommendations />
      </v-col>
      <v-col md="9" lg="7">
        <v-card elevation="3">
          <ProductDisplay />
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import ProductFilter from "../components/ProductFilter";
import ProductDisplay from "../components/browse/ProductDisplay";
import Banner from "../components/layout/Banner";
import Recommendations from "../components/Recommendations";
import { mapGetters, mapMutations } from "vuex";

export default {
  name: "browse",
  components: {
    ProductFilter,
    ProductDisplay,
    Recommendations,
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
        return "Products from ".concat(`${this.getSelectedStore.address}`);
      } catch (err) {
        return "";
      }
    },
    storeIsSelected() {
      return this.getSelectedStore !== null;
    },
  },
  methods: {
    ...mapMutations(["setDialog"]),
  },
};
</script>

<style></style>
