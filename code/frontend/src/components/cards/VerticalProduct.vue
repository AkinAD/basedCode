<template>
  <v-card outlined height="100%" class="card-outter" elevation="1">
    <v-img :src="product.image" contain height="200px" />
    <v-card-title>{{ product.title }} </v-card-title>
    <v-card-subtitle>$ {{ product.price }} </v-card-subtitle>
    <v-spacer></v-spacer>
    <v-card-actions class="card-actions" v-show="signedIn">
      <v-btn
        color="success"
        :disabled="!notAdded"
        outlined
        @click="addToCart(product)"
        ><v-icon left v-show="notAdded">add</v-icon> {{ buttonText }}</v-btn
      >
    </v-card-actions>
  </v-card>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "VerticalProduct",
  props: {
    product: Object, //this cause big is shoe
  },

  computed: {
    ...mapGetters(["signedIn", "getCart"]),

    notAdded() {
      return !(
        this.getCart.filter((product) => this.product.id === product.id)
          .length === 1
      );
    },

    buttonText() {
      return this.notAdded ? "add to cart" : "added";
    },
  },

  methods: {
    ...mapActions(["addToCart"]),
  },
};
</script>

<style scoped>
.card-outter {
  position: relative;
  padding-bottom: 50px;
}
.card-actions {
  position: absolute;
  bottom: 0;
}
</style>
