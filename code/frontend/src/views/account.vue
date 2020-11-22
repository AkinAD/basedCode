<template>
  <v-container>
    <Banner text="Hello Akin, here's your account" size="150px"/>
    <v-row align="center" justify="center">
       <amplify-sign-out></amplify-sign-out>
    </v-row>
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
import Banner from "../components/layout/Banner"
import { mapGetters, mapActions } from "vuex";

export default {
  name: "account",
  components: {
    UserCardWithPastPurchases,
    Recommendations,
    Banner
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
