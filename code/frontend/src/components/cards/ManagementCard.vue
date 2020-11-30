<template>
  <v-container outlined full-width>
    <v-row>
      <!--Image-->
      <v-col md="4">
        <v-img
          v-show="visibleImage"
          :src="product.image"
          contain
          height="200px"
        />
      </v-col>
      <!--Info-->
      <v-col md="8">
        <v-card-title>{{ product.title }} </v-card-title>
        <v-card-subtitle>$ {{ product.price }} </v-card-subtitle>
        <v-card-actions>
          <v-btn color="green" outlined>
            <v-icon left>mdi-pencil</v-icon>
            Edit {{ itemType }}</v-btn
          >
          <v-btn color="red" outlined>
            <v-icon left>mdi-delete</v-icon>
            Delete {{ itemType }}</v-btn
          >
        </v-card-actions>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "ManagementCard",
  props: {
    product: Object,
    itemType: String,
    visibleImage: Boolean,
  },
};
</script>

<style></style>
