<template>
  <v-container outlined full-width>
    <v-row>
      <!--Info-->
      <v-col md="8">
        <v-card-title>{{ product.name }} </v-card-title>
        <v-card-subtitle>$ {{ product.price }} <br/> {{product.category}} </v-card-subtitle>
        <v-card-actions>
          <ManagementDialog 
            v-on:form-saved="$emit('update-event', $event)" 
            :type="itemType"
            :fields="fields" 
            :headline="headline">
            <template v-slot:button="{ on }">
              <v-btn v-on="on" color="green" outlined>
                <v-icon left>mdi-pencil</v-icon>
                Edit {{ itemType }}
              </v-btn>
              <v-btn color="red" outlined>
                <v-icon left>mdi-delete</v-icon>
                Delete {{ itemType }}</v-btn
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

export default {
  name: "ManagementCard",
  components: {
    ManagementDialog,
  },
  props: {
    product: Object,
    itemType: String,
    visibleImage: Boolean,
    fields: Array,
  },
  data() {
    return {
      headline: `Edit ${this.itemType}`,
    };
  },
};
</script>

<style></style>
