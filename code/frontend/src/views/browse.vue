<template>
  <v-app>
    <v-container>
      <v-row>
        <v-col md="3" offset-lg="1"><Sidebar /></v-col>
        <v-col md="9" lg="7"> <ProductDisplay /> </v-col>
      </v-row>
    </v-container>
  </v-app>
</template>

<script>
import Sidebar from "../components/Sidebar";
import ProductDisplay from "../components/browse/ProductDisplay";

export default {
  name: "browse",
  components: {
    Sidebar,
    ProductDisplay,
  },
};
</script>

<style></style>
