<template>
  <v-container>
    <v-row>
      <v-col sm="8" md="5" offset-md="2" lg="4" offset-lg="3">
        <ProductList />
      </v-col>
      <v-col sm="4" md="3">
        <OptimumRoute />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import ProductList from "../components/cart/ProductList.vue";
export default {
  components: { ProductList },
};
</script>

<style></style>
