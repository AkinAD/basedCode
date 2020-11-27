<template>
  <v-container fluid full-width>
    <v-row>
      <v-col
        md="12"
        v-for="recommendation in allRecommendations"
        :key="recommendation.id"
      >
        <CartCard :product="recommendation" visibleImage/>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

import CartCard from "../cards/CartCard";

export default {
  name : "CartList",
  components: {
    CartCard,
  },
  methods: {
    removeFromCart() {},
    ...mapActions(["fetchRecommendations"]),
  },
  computed: mapGetters(["allRecommendations"]),
  created() {
    this.fetchRecommendations();
  }
};
</script>

<style></style>
