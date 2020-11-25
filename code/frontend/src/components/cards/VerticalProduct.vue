<template>
  <v-card outlined fluid>
    <v-img :src="product.image" contain height="200px" />
    <v-card-title>{{ product.title }} </v-card-title>
    <v-card-subtitle>$ {{ product.price }} </v-card-subtitle>
    <v-card-actions>
      <v-btn v-show="signedIn" color="success" outlined
        ><v-icon left>add</v-icon> Add to cart</v-btn
      >
    </v-card-actions>
  </v-card>
</template>

<script>
import {mapGetters} from 'vuex'

export default {
  name: "VerticalProduct",
  props: {
    product: Object,
    addToCart: Function,
  },

  computed: {
    ...mapGetters(["signedIn"])
  }
};
</script>

<style></style>
