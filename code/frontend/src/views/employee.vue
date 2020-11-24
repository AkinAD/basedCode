<template>
  <v-card>
    <v-tabs grow v-model="tabs">
      <!--Tabs that don't require manager authentication-->
      <v-tab>
        <v-icon left dark> mdi-briefcase </v-icon>
        Manage Items
      </v-tab>
      <v-tab>
        <v-icon left dark> mdi-briefcase </v-icon>
        Manage Categories
      </v-tab>
      <!--Tabs that do require manager authentication-->
      <v-tab>
        <v-icon left dark> mdi-briefcase </v-icon>
        Manage Employees
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
        <v-tab-item>
          <!--v-show if authenticated as manager-->
          <ManageEmployees />
        </v-tab-item>
      </keep-alive>
    </v-tabs-items>
  </v-card>
</template>

<script>
import ManageItems from "../components/manage/manageItems";
import ManageEmployees from "../components/manage/manageEmployees";
import ManageCategories from "../components/manage/manageCategories";

export default {
  name: "employee",
  components: { ManageItems, ManageEmployees, ManageCategories },
  data() {
    return {
      tabs: null,
    };
  },
};
</script>
