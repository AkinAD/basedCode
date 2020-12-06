<template>
  <v-container fluid transition="slide-x-transition">
    <v-row>
      <v-col cols= "6" md="6">
        <ManageStockList 
          title="Items in stock at"
          :items="this.getItems"
          inStore=true
          stockAction="Remove from Stock"/>
      </v-col>
      <v-col cols= "6" md="6">
        <ManageStockList 
          title="All other Items"
          :items="this.allItemsNotInStock"
          inStore=false
          stockAction="Add to Stock"/>
      </v-col>
    </v-row>
    
  </v-container>
  
</template>

<script>
import { mapGetters, mapActions } from 'vuex';
import ManageStockList from "./ManageStockList";


export default {
  name: "ManageStock",
  components: {
    ManageStockList,
  },
  data () {
    return {

    }
  },
  methods : {
    ...mapActions(["addStock", "deleteStock", "updateStock"]),
    
  },
  computed: {
    ...mapGetters([ "getItems","getAllItems", "getSelectedStore"]),

    allItemsNotInStock () {

      let allNonStockItems = [];
      allNonStockItems = this.getAllItems.filter((i) => !this.getItems.includes(i.itemID));
      return allNonStockItems;
    },
  },
 
};
</script>

<style></style>
