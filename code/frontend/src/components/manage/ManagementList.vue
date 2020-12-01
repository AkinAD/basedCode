<template>
  <v-container fluid full-width>
    <ManagementDialog v-on:form-saved="$emit('add-event', $event)" :type="type" :fields="fields" :headline="headline">
      <template v-slot:button="{on : dialog}">
        <v-tooltip top>
          <template v-slot:activator="{ on : tooltip, attrs }">
            <v-btn v-bind="attrs" fab fixed right bottom large dark color="blue" elevation="6" v-on="{ ...tooltip, ...dialog }">
             <v-icon> mdi-plus </v-icon>
            </v-btn>
          </template>
          <span>Add {{ type }}</span>
        </v-tooltip>
      </template>
    </ManagementDialog>

    <div>
      <h3>{{ title }}</h3>
      <v-row>
        <v-col
          md="12"
          v-for="recommendation in allRecommendations"
          :key="recommendation.id"
        >
          <ManagementCard
            :fields="fields" 
            v-on:update-event="$emit('update-event', $event)"
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
import ManagementDialog from "./ManagementDialog";

export default {
  name: "ManagementList",
  components: {
    ManagementCard,
    ManagementDialog
  },
  props: {
    type: String,
    showImage: Boolean,
    title: String,
    fields: Array
  },
  data () {
    return {
      headline : `Add ${this.type}`
    }
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
