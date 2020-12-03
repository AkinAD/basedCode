<template>
  <div class="text-center">
    <v-dialog v-model="dialog" width="500">
      <template v-slot:activator="{ on, attrs }">
        <v-btn color="primary" v-bind="attrs" outlined v-on="on"> info </v-btn>
      </template>

      <v-card>
        <v-container fluid mx-auto>
          <v-card-title>
            {{ product.name }}
          </v-card-title>

          <v-card-subtitle>
            Price: ${{ product.price }} <br />
            Category: {{ product.category }}
          </v-card-subtitle>

          <v-divider></v-divider>

          <v-card-text> <h4>Product description:</h4> </v-card-text>

          <v-card-text> {{ product.description }}</v-card-text>

          <v-divider></v-divider>

          <v-card-text> <h4>Location in store:</h4> </v-card-text>

          <v-card-text>
            Row: {{ product.row }} Column: {{ product.col }}</v-card-text
          >
        </v-container>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
export default {
  props: {
    product: Object, //this cause big is shoe
  },
  data() {
    return {
      dialog: false,
    };
  },
};
</script>

<style>
</style>