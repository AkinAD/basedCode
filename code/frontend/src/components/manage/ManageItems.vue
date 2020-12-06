<template>
  <v-card fluid transition="slide-x-transition">
    <ManagementList 
      v-on:add-event="addItem($event)"
      v-on:update-event="updateItem($event)" 
      v-on:delete-event="deleteItem($event)"
      :fields="itemFields" 
      title="All Items " 
      type="Item" 
      :showImage="true" />
  </v-card>
  
</template>

<script>
import { mapActions, mapGetters} from 'vuex';
import ManagementList from "./ManagementList";


export default {
  name: "ManageItems",
  components: {
    ManagementList,
  },
  data () {
    return {
      //Item fields go here whenever we find out what they are
      itemFields: [
        { name: "Item Name", value: ''},
        { name: "Price", value: ''},
        { name: "Category", value: ''}
        
      ]
    }
  },
  methods : {
    ...mapActions(["addItem", "deleteItem", "updateItem"]),
  },
  computed: {
    ...mapGetters(["getItems"]),
  }
};
</script>

<style></style>
