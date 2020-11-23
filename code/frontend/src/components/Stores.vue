<template>
  <v-container fluid>
    <v-card :loading="loading" class="mx-auto my-12" max-width="80%">
      <v-card-title>Select Shop</v-card-title>
      <v-card-text> Please select a shop to browse products from </v-card-text>
      <v-form>
        <v-container>
          <v-text-field label="search by name" outlined></v-text-field>
        </v-container>
      </v-form>
      <v-list>
        <!--show when form is submitted?-->
        <v-subheader>select a store</v-subheader>
        <v-list-item-group v-model="selectedItem" color="primary">
          <v-list-item v-for="(item, i) in items" :key="i">
            <v-list-item-content>
              <v-list-item-title v-text="item.text"></v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list-item-group>
      </v-list>

    <ProductDisplay/>

    </v-card>
  </v-container>
</template>

<script>
import ProductDisplay from '../components/browse/ProductDisplay';

export default {
  name: "stores",
  components: {
      ProductDisplay
  },
  data() {
    return {
      items: [
          { i: 1, text: "shoppers" },
          { i: 2, text: "longos" },
          { i: 3, text: "joefresh" },
          { i: 4, text: "matt's awesome store" }
          ],
    };
  },
};
</script>

<style>
</style>