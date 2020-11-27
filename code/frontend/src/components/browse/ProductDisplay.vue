<template>
  <v-container fluid>
    <h3>Products</h3>
    <v-row>
      <v-col
        sm="6"
        md="4"
        v-for="item in filteredItems"
        :key="item.id"
      >
        <VerticalProduct :product="item" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { mapGetters } from "vuex";

import VerticalProduct from "../cards/VerticalProduct";

export default {
  name: "ProductDisplay",
  components: {
    VerticalProduct,
  },
  computed: {
    ...mapGetters([
      "getItems",
      "getSelectedFilter",
    ]),

    filteredItems() {
      const filter = this.getSelectedFilter;
      if (filter === null) {//any is selected
        return this.getItems;
      } else
        return this.getItems.filter((item) => {
          return item.price <= filter[1] && item.price >= filter[0];
        });
    },
  }
};
</script>

<style></style>
