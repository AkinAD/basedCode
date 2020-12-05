<template>
  <v-card fluid transition="slide-x-transition">
    <ManageItemsList />
  </v-card>
  
</template>

<script>
import ManageItemsList from "./ManageItemsList";

// category: String
// categoryID: Int
// col: Int
// description: String
// itemID: Int
// name: String
// price: Double
// row: Int

export default {
  name: "ManageItems",
  components: {
    ManageItemsList,
  },
  data () {
    return {
      
    }
  },
  methods : {
  },
  computed: {
    
  }
};
</script>

<style></style>
