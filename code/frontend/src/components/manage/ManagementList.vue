<template>
  <v-container fluid full-width>
    <div>
      <h3>{{ title }}</h3>
      <v-btn fab fixed right bottom large dark color="blue" elevation="6">
        <!--{{btnText}}-->
        <v-icon> mdi-plus </v-icon>
      </v-btn>
      <v-row>
        <v-col
          md="12"
          v-for="recommendation in allRecommendations"
          :key="recommendation.id"
        >
          <ManagementCard
            :product="recommendation"
            :itemType="type"
            :visibleImage="showImage"
          />
        </v-col>
      </v-row>
    </div>
  </v-container>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

import ManagementCard from "../cards/ManagementCard";

export default {
  name: "ManagementList",
  components: {
    ManagementCard,
  },
  props: {
    type: String,
    showImage: Boolean,
    title: String,
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
