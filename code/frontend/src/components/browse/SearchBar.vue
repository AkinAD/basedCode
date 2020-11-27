<template>
  <v-container>
    <v-toolbar flat color="transparent">
      <v-toolbar-title class="smallFont"
        >Search for your favorite items</v-toolbar-title
      >
      <v-btn icon @click="setSearchBoxVisibility(!searchBoxVisible)">
        <v-icon>mdi-magnify</v-icon>
      </v-btn>
    </v-toolbar>

    <v-container class="py-0">
      <v-text-field
        ref="search"
        v-model="search"
        full-width
        hide-details
        label="Search"
        single-line
        v-show="searchBoxVisible"
      ></v-text-field>
    </v-container>
  </v-container>
</template>

<script>
import { mapGetters, mapMutations } from "vuex";

export default {
  data: function () {
    return {
      loading: false,
      search: "",
      searchBoxVisible: false,
    };
  },

  computed: {
    ...mapGetters(["getSelectedFilter", "fetchRecommendations"]),

    items() {
      const search = this.search.toLowerCase();

      if (!search) return this.fetchRecommendations;

      return this.fetchRecommendations.filter((item) => {
        const text = item.text.toLowerCase();
        return text.indexOf(search) > -1;
      });
    },
  },

  methods: {
    ...mapMutations(["setSearchFilter"]),

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
</style>
