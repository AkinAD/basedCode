<template>
  <v-card fluid transition="slide-x-transition">
    <ManagementList title="All Items " type="Item" :showImage="true" />
  </v-card>
</template>

<script>
import ManagementList from "./ManagementList";

export default {
  name: "ManageItems",
  components: {
    ManagementList,
  },
};
</script>

<style></style>
