<template>
  <v-container fluid>
    <v-container>
      <v-row>
        <v-col>
          <h3>Products</h3>
          <v-text-field
            ref="search"
            v-model="search"
            hide-details
            label="Search"
            single-line
          ></v-text-field>
        </v-col>
      </v-row>
    </v-container>

    <v-row>
      <v-col
        sm="6"
        md="4"
        v-for="item in searchFilteredItems"
        :key="item.itemID"
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
  data: function () {
    return {
      search: "",
      searchBoxVisible: false,
    };
  },
  computed: {
    ...mapGetters([
      "getItems",
      "getSelectedFilter",
      "getSelectedCategory",
      "getCategories",
    ]),

    filteredItems() {
      const filter = this.getSelectedFilter;
      const category = this.getSelectedCategory;
      if (filter === null) {
        if (category === 0) return this.getItems;
        else
          return this.getItems.filter((item) => {
            return item.categoryID === category;
          });
      } else
        return this.getItems
          .filter((item) => {
            return item.price <= filter[1] && item.price >= filter[0];
          })
          .filter((item) => {
            return item.categoryID === category;
          });
    },

    searchFilteredItems() {
      const search = this.search.toLowerCase();
      return this.filteredItems.filter((item) => {
        return item.name.toLowerCase().includes(search);
      });
    },
  },
  methods: {
    setSearchBoxVisibility(bool) {
      this.searchBoxVisible = bool;
    },
  },
};
</script>

<style>
.smallFont {
  font-size: 0.8em;
}
.stayBinch {
  position: relative;
  bottom: 1.6em;
  left: -1.6em;
}
</style>
