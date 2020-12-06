<template>
  <v-container fluid full-width>
    <ManagementDialog
      v-on:form-saved="_addStore($event)"
      :type="type"
      :fields="storeFields"
      :headline="headline"
    >
      <template v-slot:button="{ on: dialog }">
        <v-tooltip top>
          <template v-slot:activator="{ on: tooltip, attrs }">
            <v-btn
              v-bind="attrs"
              fab
              fixed
              right
              bottom
              large
              dark
              color="blue"
              elevation="6"
              v-on="{ ...tooltip, ...dialog }"
            >
              <v-icon> mdi-plus </v-icon>
            </v-btn>
          </template>
          <span>Add {{ type }}</span>
        </v-tooltip>
      </template>
    </ManagementDialog>

    <div>
      <h3>{{ title }}</h3>
      <v-row>
        <v-col md="12" v-for="store of getStores" :key="store.storeID">
          <ManageStoreCard
            :fields="storeFields"
            :store="store"
          />
        </v-col>
      </v-row>
    </div>
  </v-container>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

import ManageStoreCard from "../cards/ManageStoreCard";
import ManagementDialog from "./ManagementDialog";

export default {
  name: "ManageStoresList",
  components: {
    ManageStoreCard,
    ManagementDialog,
  },
  props: {
    
  },
  data() {
    return {
      headline: "Add Stores",
      type : "Store",
      title: "All Stores",
      storeFields: [
        { schemaName: "address", displayName: "Address", value: ''},
      ]
    };
  },
  computed: {
    ...mapGetters(["getStores"]),
  },
  methods: {
    ...mapActions(["addStore"]),
    mapStoreToSchema(fieldArray) {
      let mappedStore = {};
      for (var attribute of fieldArray) {
        mappedStore[attribute.schemaName] = attribute.value;
      }

      return mappedStore;
    },
    _addStore(fieldArray) {
      this.addStore(this.mapStoreToSchema(fieldArray));
    },
  },
};
</script>

<style></style>
