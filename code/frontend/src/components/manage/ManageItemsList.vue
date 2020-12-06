<template>
  <v-container fluid full-width>
    <ManagementDialog
      v-on:form-saved="_addItem($event)"
      :type="type"
      :fields="itemFields"
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
        <v-col md="12" v-for="item of getAllItems" :key="item.itemID">
          <ManageItemCard
            :fields="itemFields"
            :item="item"
          />
        </v-col>
      </v-row>
    </div>
  </v-container>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

import ManageItemCard from "../cards/ManageItemCard";
import ManagementDialog from "./ManagementDialog";

export default {
  name: "ManageItemsList",
  components: {
    ManageItemCard,
    ManagementDialog,
  },
  props: {
    
  },
  data() {
    return {
      headline: "Add Items",
      type : "Item",
      title: "All Items",
      itemFields: [
        { schemaName: "name", displayName: "Name", value: ''},
        { schemaName: "description", displayName: "Description", value: ''},
        { schemaName: "price", displayName: "Price", value: ''},
        { schemaName: "categoryID", displayName: "Category ID", value: ''},
      ]
    };
  },
  computed: {
    ...mapGetters(["getAllItems"]),
  },
  methods: {
    ...mapActions(["addItem"]),
    mapItemToSchema(fieldArray) {
      let mappedItem = {};
      for (var attribute of fieldArray) {
        mappedItem[attribute.schemaName] = attribute.value;
      }

      return mappedItem;
    },
    _addItem(fieldArray) {
      this.addItem(this.mapItemToSchema(fieldArray));
    },
  },
};
</script>

<style></style>
