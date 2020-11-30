<template>
  <v-card outlined fluid>
    <v-img :src="product.image" contain height="200px" />
    <v-card-title>{{ product.title }} </v-card-title>
    <v-card-subtitle>$ {{ product.price }} </v-card-subtitle>
    <v-card-actions>
      <v-btn color="success" outlined
        ><v-icon left>add</v-icon> Add to cart</v-btn
      >
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: "VerticalProduct",
  props: {
    product: Object,
    addToCart: Function,
  },
};
</script>

<style></style>
