<template>
  <v-container fluid>
    <v-dialog v-model="dialog" max-width="500">
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          class="ma-2"
          :color="buttonColor"
          :dark="!isEmpty"
          v-bind="attrs"
          :disabled="isEmpty"
          v-on="on"
          icon
        >
          <v-badge color="primary" :content="cartCount" :value="cartCount">
            <v-icon>mdi-map</v-icon>
          </v-badge>
        </v-btn>
      </template>

      <v-card max-width="500">
        <v-card-title>Optimum Path for Your Items</v-card-title>
        <v-divider></v-divider>
        <v-card-text>
          <v-list>
            <v-list-item
              v-for="(item, i) in orderedItems(getCart)"
              two-line
              :key="item.title"
              :value="i"
            >
              <v-list-item-avatar>
                <v-img :src="item.image" contain></v-img>
              </v-list-item-avatar>
              <v-list-item-content>
                <h3>Visit Order: {{ i + 1 }}</h3>
                <v-list-item-title v-text="item.title"></v-list-item-title>
                <v-list-item-subtitle> ${{ item.price }} </v-list-item-subtitle>
                <v-list-item-subtitle>
                  Location: {{ item.location }}
                </v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-card-text>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  data() {
    return {
      dialog: false,
    };
  },
  computed: {
    ...mapGetters(["getCart"]),

    isEmpty() {
      return this.getCart.length === 0;
    },

    buttonColor() {
      if (this.isEmpty) return "grey";
      return "primary";
    },
    cartCount() {
      return this.getCart.length;
    }
  },

  methods: {
    orderedItems(items) {
      //cart
      //do math
      return items;
    },
  },
};
</script>

<style></style>
