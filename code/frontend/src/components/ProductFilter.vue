<template>
  <div>
    <v-card elevation="3" class="pa-3">
      <v-card-title> Price Filter </v-card-title>
      <v-divider></v-divider>
      <v-card-text>
        <v-radio-group v-model="selectionPrice" :mandatory="true">
          <v-radio
            v-for="(price, i) in this.getFilters"
            :key="i"
            :label="price"
            :value="i"
            @click="setFilter"
          />
        </v-radio-group>
      </v-card-text>
      <v-card-title>Category Filter</v-card-title>
      <v-divider></v-divider>
      <v-card-text>
        <v-radio-group v-model="selectionCategory" :mandatory="true">
          <v-radio
            v-for="cat in this.getCategories"
            :key="cat.categoryID"
            :label="cat.category"
            :value="cat.categoryID"
            @click="setSelectedCategory(selectionCategory)"
          />
        </v-radio-group>
      </v-card-text>
    </v-card>
  </div>
</template>

<script>
import { mapGetters, mapMutations, mapActions } from "vuex";

export default {
  data() {
    return {
      checkbox: true,
      selectionPrice: 0,
      selectionCategory: 0,
    };
  },
  computed: {
    ...mapGetters([
      "getFilters",
      "getPriceFilter",
      "getCategories",
      "getSelectedCategory",
    ]),
  },

  methods: {
    ...mapMutations(["setSelectedFilter", "setSelectedCategory"]),
    ...mapActions(["fetchCategories"]),

    setFilter() {
      if (this.selectionPrice === 0) this.setSelectedFilter(null);
      if (this.selectionPrice === 1) this.setSelectedFilter([0, 25]);
      if (this.selectionPrice === 2) this.setSelectedFilter([25, 100]);
      if (this.selectionPrice === 3) this.setSelectedFilter([100, 500]);
      if (this.selectionPrice === 4) this.setSelectedFilter([500, 10000]);
    },
  },
};
</script>

<style></style>
