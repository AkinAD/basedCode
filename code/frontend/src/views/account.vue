<template>
  <v-container>
    <v-container class="fill-height">
      <v-row align="center" justify="center">
        <h1 class="text-center display-1 hidden-xs-only">
          How's it going Akin?
        </h1>
        <h1 class="text-center display-2 hidden-sm-and-up">
          How's it going Akin?
        </h1>
      </v-row>
    </v-container>
    <v-row>
      <v-col md="3" offset-lg="1"><UserCardWithPastPurchases /></v-col>
      <v-col md="9" lg="7">
        <v-col>
          <v-card
            ><v-col sm="10" lg="8">
              <h2>Frequent Purchases</h2>
            </v-col>
            <!-- <v-col
              sm="6"
              md="4"
              v-for="purchasedItem in allPurchasedItems.slice(2, 9)"
              :key="purchasedItem.id"
            >
            
            </v-col> -->

            <v-card class="" max-width="344" outlined elevation="2">
              <v-list-item three-line>
                <v-list-item-content>
                  <div class="overline mb-4">Banana</div>
                </v-list-item-content>
                <v-list-item-avatar
                  tile
                  size="80"
                  color="yellow"
                ></v-list-item-avatar>
              </v-list-item>
              <v-card-actions>
                <v-btn color="success" outlined
                  ><v-icon left>add</v-icon> Add to cart</v-btn
                >
              </v-card-actions>
            </v-card>
            <v-card class="" max-width="344" outlined elevation="2">
              <v-list-item three-line>
                <v-list-item-content>
                  <div class="overline mb-4">Banana</div>
                </v-list-item-content>
                <v-list-item-avatar
                  tile
                  size="80"
                  color="yellow"
                ></v-list-item-avatar>
              </v-list-item>
              <v-card-actions>
                <v-btn color="success" outlined
                  ><v-icon left>add</v-icon> Add to cart</v-btn
                >
              </v-card-actions>
            </v-card>
          </v-card>
        </v-col>
      </v-col>
    </v-row>

    <v-row>
      <Recommendations />
    </v-row>
  </v-container>
</template>

<script>
import UserCardWithPastPurchases from "../components/cards/UserCardWithPastPurchases";
import Recommendations from "../components/Recommendations";
import { mapGetters, mapActions } from "vuex";

export default {
  name: "account",
  components: {
    UserCardWithPastPurchases,
    Recommendations,
  },
  methods: {
    ...mapActions(["fetchPurchasedItems"]),
  },
  computed: mapGetters(["allPurchasedItems"]),
  created() {
    this.fetchPurchasedItems();
  },
};
</script>

<style></style>
