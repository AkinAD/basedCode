<template>
  <v-container>
    <Banner text="Your Shopping Cart" size="100px" />
    <CartList />
    <!-- <v-row>
      <v-col sm="8" md="5" offset-md="2" lg="4" offset-lg="3">
        <CartList />
      </v-col>
      <v-col sm="4" md="3">
        <OptimumRoute />
      </v-col>
    </v-row> -->
    <v-btn fab fixed right bottom large dark color="blue" elevation="6" to="/browse">
        <!--{{btnText}}-->
        <v-icon> mdi-plus </v-icon>
    </v-btn>
  </v-container>
</template>

<script>
import CartList from "../components/cart/CartList.vue";
import Banner from "../components/layout/Banner.vue";

export default {
  components: { 
    CartList,
    Banner
    },
};
</script>

<style></style>
