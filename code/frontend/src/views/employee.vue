<template>
  <v-card>
    <v-tabs grow v-model="tabs">
      <!--Tabs that don't require manager authentication-->
      <v-tab>
        <v-icon left dark> mdi-barcode </v-icon>
        Manage Items
      </v-tab>
      <v-tab>
        <v-icon left dark> mdi-shape </v-icon>
        Manage Categories
      </v-tab>
      <!--Tabs that do require manager authentication-->
      <v-tab v-show="isManager">
        <v-icon left dark> mdi-account-group </v-icon>
        Manage Employees
      </v-tab>
      <v-tab v-show="isManager">
        <v-icon left dark> mdi-city </v-icon>
        Manage Stores
      </v-tab>
       <v-tab v-show="isManager">
        <v-icon left dark> mdi-stocking </v-icon>
        Manage Stock
      </v-tab>
    </v-tabs>

    <v-tabs-items v-model="tabs">
      <keep-alive>
        <v-tab-item>
          <ManageItems />
        </v-tab-item>
      </keep-alive>
      <keep-alive>
        <v-tab-item>
          <ManageCategories />
        </v-tab-item>
      </keep-alive>
      <keep-alive>
        <v-tab-item v-show="isManager">
          <ManageEmployees />
        </v-tab-item>
      </keep-alive>
      <keep-alive>
        <v-tab-item v-show="isManager">
          <ManageStores />
        </v-tab-item>
      </keep-alive>
    </v-tabs-items>
  </v-card>
</template>

<script>
import ManageItems from "../components/manage/ManageItems";
import ManageEmployees from "../components/manage/ManageEmployees";
import ManageCategories from "../components/manage/ManageCategories";
import ManageStores from "../components/manage/ManageStores";
import { mapGetters } from "vuex";

export default {
  name: "employee",
  components: { ManageItems, ManageEmployees, ManageCategories, ManageStores },
  data() {
    return {
      tabs: null,
    };
  },
  computed: {
    ...mapGetters(["getUserGroups"]),
    isManager() {
      return this.getUserGroups.includes("manager");
    },
  },
};
</script>
