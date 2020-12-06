<template>
  <v-container fluid full-width>
    <ManagementDialog
      v-on:form-saved="_addEmployee($event)"
      :type="type"
      :fields="employeeFields"
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
        <v-col md="12" v-for="employee in getEmployees" :key="employee.employeeID">
          <ManageEmployeeCard
            :fields="employeeFields"
            :employee="employee"
          />
        </v-col>
      </v-row>
    </div>
  </v-container>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

import ManageEmployeeCard from "../cards/ManageEmployeeCard";
import ManagementDialog from "./ManagementDialog";

export default {
  name: "ManageEmployeesList",
  components: {
    ManageEmployeeCard,
    ManagementDialog,
  },
  props: {
    
  },
  data() {
    return {
      headline: "Add Employees",
      type : "Employee",
      title: "All Employees",
      employeeFields: [
        { schemaName: "username", displayName: "Username", value: ''},
        { schemaName: "email", displayName: "Email", value: ''},
        { schemaName: "storeid", displayName: "Store ID", value: ''},
        { schemaName: "firstname", displayName: "First Name", value: ''},
        { schemaName: "lastname", displayName: "Last Name", value: ''},
      ]
    };
  },
  computed: {
    ...mapGetters(["getEmployees"]),
  },
  methods: {
    ...mapActions(["addEmployee"]),
    mapEmployeeToSchema(fieldArray) {
      let mappedEmployee = {};
      for (var attribute of fieldArray) {
        mappedEmployee[attribute.schemaName] = attribute.value;
      }

      return mappedEmployee;
    },
    _addEmployee(fieldArray) {
      this.addEmployee(this.mapEmployeeToSchema(fieldArray));
    },
  },
};
</script>

<style></style>
