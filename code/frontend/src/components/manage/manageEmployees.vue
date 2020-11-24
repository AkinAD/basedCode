<template>
  <v-card fluid transition="slide-x-transition">
    <ManagementList title="All Employees" type="Employee" :showImage="false" />
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
