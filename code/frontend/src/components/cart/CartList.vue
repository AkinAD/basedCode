<template>
  <v-container fluid full-width>
    <v-row>
      <v-col
        md="12"
        v-for="item in cartItems"
        :key="item.id"
      >
        <CartCard :product="item" visibleImage v-on:remove-from-cart="removeFromCart($event)"/>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { mapGetters, mapActions, mapMutations } from "vuex";

import CartCard from "../cards/CartCard";

export default {
  name : "CartList",
  components: {
    CartCard,
  },
  methods: {
    ...mapActions(["fetchSavedCart",]),
    ...mapMutations(["removeFromCart"]),
  },
  computed: 
    mapGetters(["cartItems"]),
    created() {
      this.fetchSavedCart();
    },
  
};
</script>

<style></style>
