<template>
  <v-container fluid>
    <v-card class="card-outter" elevation="3">
      <v-row>
        <!--Info-->
        <v-col>
          <v-card-title>{{ product.name }} </v-card-title>
          <v-card-subtitle>${{ product.price }} <br/> {{product.category}} </v-card-subtitle>
          <v-card-actions class="card-actions">
            <v-btn
              color="red"
              outlined
              v-on:click="$emit('remove-from-cart', product)"
            >
              <v-icon left>mdi-delete</v-icon>
              Remove from Cart</v-btn
            >
          </v-card-actions>
        </v-col>
      </v-row>
    </v-card>
  </v-container>
</template>

<script>
export default {
  name: "CartCard",
  props: {
    product: Object,
    visibleImage: Boolean,
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
