<template>
  <div>
    <v-app>
      <v-sheet v-if="$vuetify.breakpoint.mdAndUp" :elevation="4" class="pa-8"
        ><h3>Filters</h3>
        <v-radial-group v-model="priceFilter" :mandatory="true">
          <v-radio
            v-for="(price, i) in prices"
            :key="i"
            :label="price"
            :value="i"
          />
        </v-radial-group>
      </v-sheet>
      <v-expansion-panels v-else>
        <v-expansion-panel>
          <v-expansion-panel-header expand-icon="arrow_drop_down"
            >Filters
          </v-expansion-panel-header>
          <v-expansion-panel-content>
            <v-radial-group v-model="priceFilter" :mandatory="true">
              <v-radio
                v-for="(price, i) in prices"
                :key="i"
                :label="price"
                :value="i"
              />
            </v-radial-group>
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-app>
  </div>
</template>

<script>
export default {
  data() {
    return {
      prices: ["Any", "Under $25", "$25 to $100", "$100 to $500", "Over $500"],
      priceFilter: 0,
      checkbox: true,
    };
  },
};
</script>

<style></style>
