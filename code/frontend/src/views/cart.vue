<template>
  <v-container fluid>
    <Banner text="Your Shopping Cart" size="100px" />
    <v-row align="center" justify="center">
      <h2>{{ subtotal }}</h2>
    </v-row>
    <v-container>
      <v-row align="center" justify="center">
        <v-btn color="red" outlined :disabled="noItems" v-on:click="emptyCart"
          >Empty Cart</v-btn
        >
      </v-row>
    </v-container>

    <v-row>
      <OptimumRoute />
    </v-row>
    <v-row>
      <CartList />
    </v-row>
  </v-container>
</template>

<script>
import CartList from "../components/cart/CartList.vue";
import Banner from "../components/layout/Banner.vue";
import OptimumRoute from "../components/OptimumRoute";
import { mapGetters, mapMutations } from "vuex";

export default {
  components: {
    CartList,
    Banner,
    OptimumRoute,
  },
  methods: {
    ...mapMutations(["emptyCart"]),
  },
  computed: {
    ...mapGetters(["getCart"]),

    noItems() {
      return this.getCart.length === 0;
    },

    subtotal() {
      var sum = 0;
      this.getCart.forEach((element) => (sum += element.price));
      return "Your Subtotal is: $" + sum.toFixed(2);
    },
  },
};
</script>

<style></style>
