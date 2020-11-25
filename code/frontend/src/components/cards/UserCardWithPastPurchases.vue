<template>
  <v-container>
    <v-row justify="space-around">
      <v-card width="400">
        <v-img
          height="200px"
          src="https://cdn.pixabay.com/photo/2020/07/12/07/47/bee-5396362_1280.jpg"
        >
          <v-app-bar flat color="rgba(0, 0, 0, 0)">
            <v-toolbar-title class="title white--text pl-0">
              Overview
            </v-toolbar-title>

            <v-spacer></v-spacer>

            <v-menu transition="slide-y-transition" bottom>
              <template v-slot:activator="{ on, attrs }">
                <v-btn color="white" icon v-bind="attrs" v-on="on">
                  <v-icon>mdi-dots-vertical</v-icon>
                </v-btn>
              </template>
              <v-list>
                <v-list-item>
                  <v-list-item-title>Edit Profile</v-list-item-title>
                </v-list-item>
                <v-list-item>
                  <v-list-item-title>Change Defaut Store</v-list-item-title>
                </v-list-item>
              </v-list>
            </v-menu>
          </v-app-bar>

          <v-card-title class="white--text mt-8">
            <v-avatar size="56">
              <img
                alt="user"
                src="https://cdn.pixabay.com/photo/2020/06/24/19/12/cabbage-5337431_1280.jpg"
              />
            </v-avatar>
            <p class="ml-3">{{userName}}</p>
          </v-card-title>
        </v-img>

        <v-card-text>
          <div class="font-weight-bold ml-8 mb-2">Your Purchases</div>

          <v-timeline align-top dense>
            <v-timeline-item
              v-for="purchase in purchases"
              :key="purchase.address"
              :color="purchase.color"
              small
            >
              <div>
                <div class="font-weight-normal">
                  <strong>{{ purchase.store }}</strong> @{{ purchase.address }}
                </div>
                <div>{{ purchase.purchaseCount }}</div>
              </div>
            </v-timeline-item>
          </v-timeline>
        </v-card-text>
      </v-card>
    </v-row>
  </v-container>
</template>

<script>
import {mapGetters} from 'vuex'

export default {
  data: () => ({
    purchases: [ //needs to be converted to vuex
      {
        store: "Sloblaws",
        purchaseCount: `Purchased 9 items.`,
        address: "1738 Jane Street",
        color: "deep-purple lighten-1",
      },
      {
        store: "Waltmart",
        purchaseCount: "Purchases 16 items",
        address: "420 Albert Road",
        color: "green",
      },
      {
        store: "No Skills",
        purchaseCount: "Purchased 12 items",
        address: "69 Tomahawk Place",
        color: "deep-purple lighten-1",
      },
    ],
  }),

  computed: {
    ...mapGetters(["getUser"]),
    userName() {
      return this.getUser.attributes.email.split("@")[0]; //could be changed
    }
  }
};
</script>

<style></style>
