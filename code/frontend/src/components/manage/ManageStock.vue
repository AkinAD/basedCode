<template>
  <v-container fluid transition="slide-x-transition">
    <v-row>
      <v-col cols= "6" md="6">
        <ManageStockList 
          title="Items in stock at"
          :items="this.getItems"
          :inStore="true"
          stockAction="Remove from Stock"
          v-on:toggle-stock="_deleteStock($event)"
          v-on:update="_updateStock($event)"/>
      </v-col>
      <v-col cols= "6" md="6">
        <ManageStockList 
          title="All other Items"
          :items="this.allItemsNotInStock"
          :inStore="false"
          stockAction="Add to Stock"
          v-on:toggle-stock="engageForm($event)"/>
      </v-col>
    </v-row>

    <v-dialog v-model="dialog" persistent max-width="600px">
    <template>   
    </template>
    <v-card>
      <v-card-title>
        <span class="headline">Fill out Location in Store</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col
              md="12"
            >
              <v-text-field label="Row" v-model="row" required></v-text-field>
            </v-col>
            <v-col
              md="12"
            >
              <v-text-field label="column" v-model="col" required></v-text-field>
            </v-col>
          </v-row>
        </v-container>
        <small>All fields required</small>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue darken-1" text @click="formClosed()">
          Close
        </v-btn>
        <v-btn color="blue darken-1" text @click="formSaved()"> Save </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
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
      dialog : false,
      row : "",
      col : "",
      localItemID : Number
    }
  },
  methods : {
    ...mapActions(["addStock", "deleteStock", "updateStock"]),
    engageForm(itemid) {
      this.dialog = true;
      this.localItemID = itemid;
    },
    formSaved() {
      this.dialog = false;
      let stockInfo = {
        itemID : this.localItemID,
        storeID : this.getSelectedStore.storeID,
        row : this.row,
        col : this.col
      }
      this._addStock(stockInfo);
    },
    formClosed() {
      this.dialog = false;
    },
    _deleteStock(iID) {
      let stockInfo = {
        itemID : iID,
        storeID : this.getSelectedStore.storeID
      }
      this.deleteStock(stockInfo);

    },
    _addStock(stockInfo) {
      
      this.addStock(stockInfo);

    },
    _updateStock(stockInfo) {
      stockInfo = {
      itemID : Number(stockInfo.itemID),
      storeID : Number(this.getSelectedStore.storeID),
      row : Number(stockInfo.row),
      col : Number(stockInfo.col)
    };
    
    this.updateStock(stockInfo);

    },

   
    
  },
  computed: {
    ...mapGetters([ "getItems","getAllItems", "getSelectedStore"]),

     allItemsNotInStock () {
      
      let allNonStockItems = [];
      const allItems = this.getAllItems;
      const stockedItems = this.getItems;
      for(var item of allItems) {
        
        if(!stockedItems.some((sItem) => (sItem.itemID === item.itemID)))
        {
          allNonStockItems.push(item);
        }
      }

      return allNonStockItems;
    },
  },
 
};
</script>

<style></style>
