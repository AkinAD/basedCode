<template>
  <v-container outlined full-width>
    <v-row>
      <!--Info-->
      <v-col md="8">
        <v-card-title>{{ store.address }} </v-card-title>
        <!-- <v-card-subtitle>$ {{ store.price }} <br/> {{store.category}} </v-card-subtitle> -->
        <v-card-actions>
          <ManagementDialog 
            v-on:form-saved="_updateStore(store.storeID, $event)" 
            :type="type"
            :fields="populateFieldsValues" 
            :headline="headline">
            <template v-slot:button="{ on }">
              <v-btn v-on="on" color="green" outlined>
                <v-icon left>mdi-pencil</v-icon>
                Edit {{ type }}
              </v-btn>
              <v-btn color="red" outlined @click="deleteStore(store.storeID)">
                <v-icon left>mdi-delete</v-icon>
                Delete {{ type }}</v-btn
              >
            </template>
          </ManagementDialog>
        </v-card-actions>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import ManagementDialog from "../manage/ManagementDialog";
import { mapActions} from 'vuex';

export default {
  name: "ManageStoreCard",
  components: {
    ManagementDialog,
  },
  props: {
    store: Object,
    fields: Array,
  },
  data() {
    return {
      headline: "Edit Store",
      type: "Store",
    };
  },
  computed: {
    populateFieldsValues() {
      let filledArray = JSON.parse(JSON.stringify(this.fields));
      
      for(var field of filledArray)
      {
        field.value = this.store[`${field.schemaName}`];
      }

      return filledArray;
    }
  },
   methods : {
    ...mapActions(["deleteStore", "updateStore"]),
    mapStoreToSchema(fieldArray) {
      let mappedStore = this.store;
      for (var attribute of fieldArray) {
        mappedStore[attribute.schemaName] = attribute.value;
      }

      return mappedStore;
    },
    _updateStore(id, store) {
      store = this.mapStoreToSchema(store);
      store.storeID = id;
      this.updateStore(store);
    }
  },
};
</script>

<style></style>
