<template>
  <v-container fluid full-width>
    <!-- <ManagementDialog
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
    </ManagementDialog> -->

    <div>
      <h3>{{ title }}</h3>
      <v-row>
        <v-col md="12" v-for="item of items" :key="item.itemID">
          <ManageStockCard
            :fields="stockFields"
            :item="item"
            :inStore="inStore"
            :stockAction="stockAction"
            v-on:toggle-stock="$emit('toggle-stock', $event)"
            v-on:update="$emit('update', $event)"
          />
        </v-col>
      </v-row>
    </div>
  </v-container>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

import ManageStockCard from "../cards/ManageStockCard";

export default {
  name: "ManageStockList",
  components: {
    ManageStockCard,
  },
  props: {
    title : String,
    items : Array,
    inStore : Boolean,
    stockAction : String
  },
  data() {
    return {
      type : "Item",
      stockFields: [
        { schemaName: "row", displayName: "Row", value: ''},
        { schemaName: "col", displayName: "Column", value: ''}
      ]
    };
  },
  computed: {
    ...mapGetters(["getItems"]),
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
