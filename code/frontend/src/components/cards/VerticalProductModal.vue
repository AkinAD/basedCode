<template>
  <div class="text-center">
    <v-dialog v-model="dialog" width="500">
      <template v-slot:activator="{ on, attrs }">
        <v-btn color="primary" v-bind="attrs" outlined v-on="on"> info </v-btn>
      </template>

      <v-card>
        <v-container fluid mx-auto>
          <v-card-title class="headline grey lighten-2">
            {{ product.title }}
          </v-card-title>
          <v-card-subtitle> $ {{ product.price }} </v-card-subtitle>

          <v-img :src="product.image" contain height="200px"> </v-img>

          <v-card-text> Item description here </v-card-text>

          <v-divider></v-divider>

          <v-card-text> Location in store </v-card-text>

          <!--in store location component here | pass product as prop-->

          <v-divider></v-divider>

          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" text @click="dialog = false"> Close </v-btn>
          </v-card-actions>
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