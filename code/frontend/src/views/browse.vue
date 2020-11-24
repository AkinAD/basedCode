<template>
  <v-container fluid transition="slide-x-transition">
    <v-row>
      <v-col md="3" offset-lg="1"><Sidebar /></v-col>
      <v-col md="9" lg="7"> <ProductDisplay /> </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Sidebar from "../components/Sidebar";
import ProductDisplay from "../components/browse/ProductDisplay";
import { mapGetters } from "vuex";

export default {
  name: "browse",
  components: {
    Sidebar,
    ProductDisplay,
  },
  created() {
    if (this.getLocation === null) {
      const trigerLocationSetBtn = this.$refs.setLocationBtn;
      console.log("what it do: ", trigerLocationSetBtn);
      trigerLocationSetBtn.click();
    }
  },
  computed: {
    ...mapGetters(["getLocation"]),
  },
};
</script>

<style></style>
