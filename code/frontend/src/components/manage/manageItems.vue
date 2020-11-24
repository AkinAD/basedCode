<template>
  <v-card fluid transition="slide-x-transition">
    <ManagementList title="All Items " type="Item" :showImage="true" />
  </v-card>
</template>

<script>
import ManagementList from "../browse/ManagementList";

export default {
  name: "manageItems",
  components: {
    ManagementList,
  },
};
</script>

<style></style>
