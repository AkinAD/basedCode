<template>
  <v-card fluid transition="slide-x-transition">
    <ManagementList title="All Categories" type="Category" :showImage="false" />
  </v-card>
</template>

<script>
import ManagementList from "../browse/ManagementList";

export default {
  name: "manageEmployees",
  components: {
    ManagementList,
  },
};
</script>

<style></style>
