<template>
  <v-container fluid full-width>
    <ManagementDialog
      v-on:form-saved="_addCategory($event)"
      :type="type"
      :fields="categoryFields"
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
        <v-col md="12" v-for="category of getCategories" :key="category.categoryID">
          <ManageCategoryCard
            :fields="categoryFields"
            :category="category"
          />
        </v-col>
      </v-row>
    </div>
  </v-container>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

import ManageCategoryCard from "../cards/ManageCategoryCard";
import ManagementDialog from "./ManagementDialog";

export default {
  name: "ManageCategoriesList",
  components: {
    ManageCategoryCard,
    ManagementDialog,
  },
  props: {
    
  },
  data() {
    return {
      headline: "Add Categories",
      type : "Category",
      title: "All Categories",
      categoryFields: [
        { schemaName: "category", displayName: "Name", value: ''},
      ]
    };
  },
  computed: {
    ...mapGetters(["getCategories"]),
  },
  methods: {
    ...mapActions(["addCategory"]),
    mapCategoryToSchema(fieldArray) {
      let mappedCategory = {};
      for (var attribute of fieldArray) {
        mappedCategory[attribute.schemaName] = attribute.value;
      }

      return mappedCategory;
    },
    _addCategory(fieldArray) {
      this.addCategory(this.mapCategoryToSchema(fieldArray));
    },
  },
};
</script>

<style></style>
