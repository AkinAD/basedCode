<template>
  <div>
    <v-card v-if="$vuetify.breakpoint.mdAndUp" elevation="3" class="pa-8">
      <h3>Price Filter</h3>
      <v-radio-group v-model="selection" :mandatory="true">
        <v-radio
          v-for="(price, i) in this.getFilters"
          :key="i"
          :label="price"
          :value="i"
          @click="setFilter"
        />
      </v-radio-group>
    </v-card>

    <v-expansion-panels v-else>
      <v-expansion-panel>
        <v-expansion-panel-header expand-icon="arrow_drop_down"
          >Filters
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          <v-radio-group v-model="selection" :mandatory="true">
            <v-radio
              v-for="(price, i) in this.getFilters"
              :key="i"
              :label="price"
              :value="i"
              @click="setFilter"
            />
          </v-radio-group>
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
  </div>
</template>

<script>
import { mapGetters, mapMutations } from "vuex";

export default {
  data() {
    return {
      checkbox: true,
      selection: 0,
    };
  },
  computed: {
    ...mapGetters(["getFilters", "getPriceFilter"]),
  },

  methods: {
    ...mapMutations(["setSelectedFilter"]),

    setFilter() {
      if (this.selection === 0) this.setSelectedFilter(null);
      if (this.selection === 1) this.setSelectedFilter([0, 25]);
      if (this.selection === 2) this.setSelectedFilter([25, 100]);
      if (this.selection === 3) this.setSelectedFilter([100, 500]);
      if (this.selection === 4) this.setSelectedFilter([500, 10000]);
    },
  },
};
</script>

<style></style>
