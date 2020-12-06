<template>
  <v-card fluid transition="slide-x-transition">
    <ManageEmployeesList />
  </v-card>
</template>

<script>
import ManageEmployeesList from "./ManageEmployeesList";

export default {
  name: "ManageEmployees",
  components: {
    ManageEmployeesList,
  },
};
</script>

<style></style>
