<template>
  <v-container outlined full-width>
    <v-row>
      <!--Info-->
      <v-col>
        <v-card outlined>
          <v-card-title
            >{{ employee.firstname }} {{ employee.lastname }}</v-card-title
          >
          <v-card-subtitle
            >{{ employee.username }} <br />
            {{ employee.email }} <br />
            {{ employee.storeid }}
          </v-card-subtitle>
          <v-card-actions>
            <v-btn
              color="red"
              outlined
              v-on:click="deleteEmployee(employee.username)"
            >
              <v-icon left>mdi-delete</v-icon>
              Delete {{ type }}</v-btn
            >
            <v-btn
              color="blue"
              v-show="isManager || isAdmin"
              outlined
              v-on:click="promoteToManager(employee)"
            >
              <v-icon left>mdi-plus</v-icon>
              Promote to Manager</v-btn
            >
            <v-btn
              color="blue"
              v-show="isAdmin"
              outlined
              v-on:click="promoteToAdmin(employee)"
            >
              <v-icon left>mdi-plus</v-icon>
              Promote to Admin</v-btn
            >
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
//import ManagementDialog from "../manage/ManagementDialog";
import { mapActions, mapGetters } from "vuex";

export default {
  name: "ManageEmployeeCard",
  components: {
    //ManagementDialog,
  },
  props: {
    employee: Object,
    fields: Array,
  },
  data() {
    return {
      headline: "Edit Employee",
      type: "Employee",
    };
  },
  computed: {
    ...mapGetters(["getUserGroups"]),
    populateFieldsValues() {
      let filledArray = JSON.parse(JSON.stringify(this.fields));

      for (var field of filledArray) {
        field.value = this.employee[`${field.schemaName}`];
      }

      return filledArray;
    },
    isManager() {
      const groups = this.getUserGroups;
      return groups.includes("manager");
    },
    isAdmin() {
      const groups = this.getUserGroups;
      return groups.includes("admin");
    },
  },
  methods: {
    ...mapActions([
      "deleteEmployee",
      "updateEmployee",
      "promoteToAdmin",
      "promoteToManager",
    ]),
    mapEmployeeToSchema(fieldArray) {
      let mappedEmployee = {};
      for (var attribute of fieldArray) {
        mappedEmployee[attribute.schemaName] = attribute.value;
      }

      return mappedEmployee;
    },
    _updateEmployee(id, employee) {
      employee.employeeID = id;
      this.updateEmployee(this.mapEmployeeToSchema(employee));
    },
  },
};
</script>

<style></style>
