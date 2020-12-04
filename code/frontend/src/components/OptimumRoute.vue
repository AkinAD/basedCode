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
              v-for="(item, i) in orderedItems"
              two-line
              :key="item.title"
              :value="i"
            >
              <v-list-item-content>
                <h3>Visit Order: {{ i + 1 }}</h3>
                <v-list-item-title v-text="item.name"></v-list-item-title>
                <v-list-item-subtitle>
                  ${{ item.price }} <br />
                  {{ item.category }}
                </v-list-item-subtitle>
                <v-list-item-subtitle>
                  Row: {{ item.row }} Column: {{ item.col }}
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

    orderedItems() {
      let temp = this.getCart;
      return temp.sort((a, b) => {
        if (a.row == b.row) {
          return a.col > b.col ? 1 : a.col < b.col ? -1 : 0;
        }
        return a.row > b.row ? 1 : -1;
      });
    },

    isEmpty() {
      return this.getCart.length === 0;
    },

    buttonColor() {
      if (this.isEmpty) return "grey";
      return "primary";
    },
    cartCount() {
      return this.getCart.length;
    },
  },

  methods: {},
};
</script>

<style></style>
