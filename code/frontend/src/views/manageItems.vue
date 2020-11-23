<template>
  <v-container fluid transition="slide-x-transition">
    <ProductDisplayManager />
  </v-container>
</template>

<script>

import ProductDisplayManager from "../components/browse/ProductDisplayManager";

export default {
  name: "manageItems",
  components: {
    ProductDisplayManager,
  },
};
</script>

<style></style>