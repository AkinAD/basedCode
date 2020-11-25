<template>
  <v-container fluid>
    <h3>Products</h3>
    <v-row>
      <v-col
        sm="6"
        md="4"
        v-for="recommendation in filteredRecommendations"
        :key="recommendation.id"
      >
        <VerticalProduct :product="recommendation" />
      </v-col>
    </v-row>
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
  computed: {
    ...mapGetters([
      "allRecommendations",
      "getSelectedStore",
      "getSelectedFilter",
    ]),

    filteredRecommendations() {
      const filter = this.getSelectedFilter;
      if (filter === null) {
        return this.allRecommendations;
      } else
        return this.allRecommendations.filter((product) => {
          return product.price <= filter[1] && product.price >= filter[0];
        });
    },
  },
  created() {
    this.fetchRecommendations();
  },
};
</script>

<style></style>
