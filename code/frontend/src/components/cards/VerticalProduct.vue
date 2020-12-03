<template>
  <v-card outlined height="100%" elevation="1" class="d-flex flex-column">
    <v-container fluid>
      <v-img :src="product.image" contain height="200px" />
    </v-container>

    <v-card-title>{{ product.title }} </v-card-title>
    <v-card-subtitle>$ {{ product.price }} </v-card-subtitle>

    <v-spacer></v-spacer>

    <v-card-actions v-show="signedIn">
      <v-row>
        <v-col>
          <v-btn
            color="success"
            :disabled="!notAdded"
            outlined
            @click="addToCart(product)"
            ><v-icon left v-show="notAdded">add</v-icon> {{ buttonText }}</v-btn
          >
        </v-col>
        <v-spacer></v-spacer>
        <v-col>
          <VerticalProductModal :product="this.product" />
        </v-col>
      </v-row>
    </v-card-actions>
  </v-card>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
import VerticalProductModal from "./VerticalProductModal.vue";

export default {
  components: { VerticalProductModal },
  name: "VerticalProduct",
  props: {
    product: Object, //this cause big is shoe
  },

  computed: {
    ...mapGetters(["signedIn", "getCart"]),

    notAdded() {
      return !(
        this.getCart.filter((product) => this.product.id === product.id)
          .length === 1
      );
    },

    buttonText() {
      return this.notAdded ? "add to cart" : "added";
    },
  },

  methods: {
    ...mapActions(["addToCart"]),
  },
};
</script>

<style scoped>
.card-outter {
  position: relative;
  padding-bottom: 50px;
}
.card-actions {
  position: absolute;
  bottom: 0;
}
</style>
