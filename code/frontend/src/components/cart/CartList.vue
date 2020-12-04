<template>
  <v-container>
    <v-row md="12" v-for="item in getCart" :key="item.id">
      <CartCard
        :product="item"
        visibleImage
        v-on:remove-from-cart="removeFromCart($event)"
      />
    </v-row>
  </v-container>
</template>

<script>
import { mapGetters, mapMutations } from "vuex";

import CartCard from "../cards/CartCard";

export default {
  name: "CartList",
  components: {
    CartCard,
  },
  methods: {
    ...mapMutations(["removeFromCart"]),
  },
  computed: {
    ...mapGetters(["getCart"]),
  },
};
</script>

<style></style>
