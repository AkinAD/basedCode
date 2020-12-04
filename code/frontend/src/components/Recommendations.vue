<template>
  <v-container fluid>
    <v-row>
      <h2>We think you'll like these products</h2>
    </v-row>
    <v-row
      v-for="recommendation in allRecommendations"
      :key="recommendation.id"
    >
      <v-container fluid>
        <VerticalProduct :product="recommendation" />
      </v-container>
    </v-row>
    <v-container fluid>
      <v-card v-if="allRecommendations.length === 0">
        <v-card-title
          >Oh no! It appears we have no recommendations for you :(
        </v-card-title>
      </v-card>
    </v-container>
  </v-container>
</template>

<script>
import { mapGetters } from "vuex";

import VerticalProduct from "../components/cards/VerticalProduct";

export default {
  name: "Recommendations",
  components: { VerticalProduct },
  computed: {
    ...mapGetters(["getCart", "getItems"]),
    allRecommendations() {
      return this.getItems.filter((item) => !this.getCart.includes(item));
    },
  },
};
</script>

<style></style>
