<template>
  <v-container>
    <Banner :text="msg" size="100px" />
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
            <!--past purchases to be displayed here -> get across all stores -->
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
import Banner from "../components/layout/Banner";
import { mapGetters, mapActions } from "vuex";

export default {
  name: "account",
  components: {
    UserCardWithPastPurchases,
    Recommendations,
    Banner,
  },
  methods: {
    ...mapActions(["fetchPurchasedItems"]),
  },
  computed: {
    ...mapGetters(["allPurchasedItems"]),
    ...mapGetters(["getUser", "getUserGroups"]),
    msg() {
      try {
        return "Hello, ".concat(`${this.getUser.attributes.email}`)
      } catch (err) {
        return ""
      }
    }
  },
};
</script>

<style></style>
