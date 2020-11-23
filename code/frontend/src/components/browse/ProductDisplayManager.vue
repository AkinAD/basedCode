<template>
  <v-container fluid>
    <div>
      <h3>Products</h3>
      <v-row
        sm="6"
        md="4"
        v-for="recommendation in allRecommendations"
        :key="recommendation.id"
    >
        <VerticalProductManager :product="recommendation" />
        
      </v-row>
    </div>
  </v-container>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

import VerticalProductManager from "../cards/VerticalProductManager";

export default {
  name: "ProductDisplay",
  components: {
    VerticalProductManager,
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
