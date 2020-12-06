<template>
  <v-card fluid transition="slide-x-transition">
    <ManageCategoriesList />
  </v-card>
</template>

<script>
import ManageCategoriesList from "./ManageCategoriesList";

export default {
  name: "ManageCategories",
  components: {
   ManageCategoriesList,
  },
  computed : {
  }
};
</script>

<style></style>
