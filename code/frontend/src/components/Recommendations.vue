<template>
  <v-app>
    <v-container mb-12>
      <v-row>
        <v-col sm="10" offset-sm="1" lg="8" offset-lg="2">
          <h2>
            We think you'll like these products
            <v-btn to="recommended" color="primary" small text>View All</v-btn>
          </h2>
        </v-col>
      </v-row>
      <v-row>
        <v-col sm="10" offset-sm="1" lg="8" offset-lg="2">
          <v-row>
            <v-col
              sm="6"
              md="4"
              v-for="recommendation in allRecommendations.slice(2, 9)"
              :key="recommendation.id"
            >
              <VerticalProduct :product="recommendation" />
            </v-col>
          </v-row>
        </v-col>
      </v-row>
    </v-container>
  </v-app>
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
