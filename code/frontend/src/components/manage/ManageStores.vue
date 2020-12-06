<template>
  <v-card fluid transition="slide-x-transition">
    <ManageStoresList />
  </v-card>
  
</template>

<script>
import ManageStoresList from "./ManageStoresList";

export default {
  name: "ManageStores",
  components: {
    ManageStoresList,
  },
  data () {
    return {
      
    }
  },
  methods : {
  },
  computed: {
    
  },
};
</script>

<style></style>
