<template>
  <v-container fluid class="rounded-card">
    <v-parallax src="https://cdn.vuetifyjs.com/images/parallax/material.jpg">
      <v-layout justify-start align-end column pa-5>
        <div class="display-4 font-weight-black white--text text-xs-center">
          SMART SHOPPING
        </div>
        <div
          outline
          class="display-4 font-weight white--text text-xs-center mb-3"
        >
          MADE EASY
        </div>
        <div class="display-1 font-weight-bold white--text text-xs-center">
          &lt; find great products that relate to you >
        </div>
        <div class="display-1 font-weight-bold white--text text-xs-center">
          &lt; get the optimum path between products >
        </div>
        <div class="display-1 font-weight-bold white--text text-xs-center">
          &lt; streamline your shopping experience >
        </div>
      </v-layout>
    </v-parallax>
  </v-container>
</template>

<script>
export default {
  name: "landing",
};
</script>

<style scoped>
.rounded-card {
  border-radius: 50px;
}
</style>