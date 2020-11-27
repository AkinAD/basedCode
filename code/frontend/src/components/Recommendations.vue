<template>
  <v-container fluid>
    <v-row>
      <v-card elevation="3" class="pa-8">
        <h2>We think you'll like these products</h2>
      </v-card>
    </v-row>
    <v-row
      v-for="recommendation in allRecommendations"
      :key="recommendation.id"
      sm="6"
      md="4"
    >
      <v-card elevation="2">
        <VerticalProduct :product="recommendation" />
      </v-card>
    </v-row>
  </v-container>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

import VerticalProduct from "../components/cards/VerticalProduct";

export default {
  name: "Recommendations",
  components: { VerticalProduct },
  methods: {
    ...mapActions(["fetchRecommendations"]),
  },
  computed: mapGetters(["allRecommendations"]),
  created() {
    this.fetchRecommendations();
  },
};
</script>

<style></style>
