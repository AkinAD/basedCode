<template>
  <v-autocomplete
    v-model="model"
    :items="allRecommendations"
    :loading="isLoading"
    :search-input.sync="search"
    chips
    clearable
    hide-details
    auto-select-first
    item-text="title"
    label="Search for a your favorite items ..."
    solo
  >
    <template v-slot:no-data>
      <v-list-item>
        <v-list-item-title>
          Search for all your Smart<strong>Shopping</strong> items
        </v-list-item-title>
      </v-list-item>
    </template>
    <template v-slot:selection="{ attr, on, item, selected }">
      <v-chip
        v-bind="attr"
        :input-value="selected"
        color="blue-grey"
        class="white--text"
        v-on="on"
      >
        <v-icon left> mdi-shopping </v-icon>
        <span v-text="item.title"></span>
      </v-chip>
    </template>
    <template v-slot:item="{ item }">
      <v-list-item-avatar
        color="indigo"
        class="headline font-weight-light white--text"
      >
        <v-icon dark> mdi-tag-text </v-icon>
      </v-list-item-avatar>
      <v-list-item-content>
        <v-list-item-title v-text="item.title"></v-list-item-title>
        <v-list-item-subtitle v-text="item.category"></v-list-item-subtitle>
      </v-list-item-content>
      <v-list-item-action>
        <v-btn color="success" outlined
          ><v-icon left>add</v-icon> Add to cart</v-btn
        >
        <!-- <v-icon>mdi-shopping</v-icon> -->
      </v-list-item-action>
    </template>
  </v-autocomplete>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "SearchBar",
  components: {},
  methods: {
    ...mapActions(["fetchRecommendations"]),
  },
  computed: {
    ...mapGetters([
      "allRecommendations",
      "getSelectedStore",
      "getSelectedFilter",
    ]),

    filteredRecommendations() {
      const filter = this.getSelectedFilter;
      //TODO: not sure if i wanna be using the filters here tbh
      if (filter === null) {
        return this.allRecommendations;
      } else
        return this.allRecommendations.filter((product) => {
          return product.price <= filter[1] && product.price >= filter[0];
        });
    },
  },
  watch: {
    // model(val) {
    //   console.log("This is executing: ", val);
    // },
    // search(val) {
    //   console.log(val);
    //   if (this.items.length > 0) return;
    // },
  },

  created() {
    this.fetchRecommendations();
  },
};
</script>

<style></style>
