<template>
  <v-container fluid>
    <div>
      <h3>Products</h3>
      <v-row>
        <v-col
          sm="6"
          md="4"
          v-for="recommendation in allRecommendations"
          :key="recommendation.id"
        >
          <VerticalProduct :product="recommendation" />
        </v-col>
      </v-row>
    </div>
  </v-container>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

import VerticalProduct from "../cards/VerticalProduct";

export default {
  name: "ProductDisplay",
  components: {
    VerticalProduct,
  },
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
