<template>
  <v-container outlined full-width>
    <v-row>
      <!--Info-->
      <v-col>
        <v-card outlined>
          <v-card-title>{{ category.category }} </v-card-title>
          <v-card-subtitle>Category ID: {{ category.categoryID }}</v-card-subtitle>
          <v-card-actions>
            <ManagementDialog
              v-on:form-saved="_updateCategory(category.categoryID, $event)"
              :type="categoryType"
              :fields="populateFieldsValues"
              :headline="headline"
            >
              <template v-slot:button="{ on }">
                <v-btn v-on="on" color="green" outlined>
                  <v-icon left>mdi-pencil</v-icon>
                  Edit {{ categoryType }}
                </v-btn>
                <v-btn
                  color="red"
                  outlined
                  @click="deleteCategory(category.categoryID)"
                >
                  <v-icon left>mdi-delete</v-icon>
                  Delete {{ categoryType }}</v-btn
                >
              </template>
            </ManagementDialog>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import ManagementDialog from "../manage/ManagementDialog";
import { mapActions } from "vuex";

export default {
  name: "ManageCategoryCard",
  components: {
    ManagementDialog,
  },
  props: {
    category: Object,
    fields: Array,
  },
  data() {
    return {
      headline: "Edit Category",
      categoryType: "Category",
    };
  },
  computed: {
    populateFieldsValues() {
      let filledArray = JSON.parse(JSON.stringify(this.fields));

      for (var field of filledArray) {
        field.value = this.category[`${field.schemaName}`];
      }

      return filledArray;
    },
  },
  methods: {
    ...mapActions(["deleteCategory", "updateCategory"]),
    mapCategoryToSchema(fieldArray) {
      let mappedCategory = this.category;
      for (var attribute of fieldArray) {
        mappedCategory[attribute.schemaName] = attribute.value;
      }

      return mappedCategory;
    },
    _updateCategory(id, category) {
      category = this.mapCategoryToSchema(category);
      category.categoryID = id;
      this.updateCategory(category);
    },
  },
};
</script>

<style></style>
