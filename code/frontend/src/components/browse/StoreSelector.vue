<template>
  <v-container fluid>
    <v-dialog v-model="dialog" persistent max-width="290">
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          class="ma-2"
          :color="buttonColor"
          dark
          v-bind="attrs"
          v-on="on"
          text
          icon
        >
          <v-icon>mdi-map-marker</v-icon>
        </v-btn>
      </template>

      <v-card max-width="500">
        <v-toolbar flat color="transparent">
          <v-toolbar-title>Select your location</v-toolbar-title>
          <v-btn icon @click="setSearchBoxVisibility(!searchBoxVisible)">
            <v-icon>mdi-magnify</v-icon>
          </v-btn>
        </v-toolbar>

        <v-container class="py-0">
          <v-text-field
            ref="search"
            v-model="search"
            full-width
            hide-details
            label="Search"
            single-line
            v-show="searchBoxVisible"
          ></v-text-field>
        </v-container>

        <v-list>
          <v-list-item-group color="primary">
            <v-list-item
              v-for="item in categories"
              two-line
              :key="item.text"
              :disabled="loading"
              @click="select(item)"
            >
              <v-list-item-avatar>
                <v-icon :disabled="loading" v-text="item.icon"></v-icon>
              </v-list-item-avatar>
              <v-list-item-content>
                <v-list-item-title v-text="item.text"></v-list-item-title>
                <v-list-item-subtitle
                  v-text="item.location"
                ></v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-list-item-group>
        </v-list>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            :disabled="!nonEmptyStore"
            :loading="loading"
            color="primary"
            text
            @click="next"
          >
            Set Location
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import { mapGetters, mapMutations, mapActions } from "vuex";
export default {
  data: function () {
    return {
      loading: false,
      search: "",
      searchBoxVisible: false,
      selected: null,
    };
  },

  computed: {
    ...mapGetters(["getSelectedStore", "getStores", "getDialog"]),
    nonEmptyStore() {
      return this.selected !== null;
    },
    categories() {
      const search = this.search.toLowerCase();

      if (!search) return this.getStores;

      return this.getStores.filter((store) => {
        const text = store.text.toLowerCase();
        return text.indexOf(search) > -1;
      });
    },
    dialog: {
      get() {
        return this.getDialog;
      },
      set(value) {
        this.setDialog(value);
      },
    },
    buttonColor() {
      if (this.getSelectedStore === null) return "grey";
      return "primary";
    },
  },

  methods: {
    ...mapMutations(["setDialog"]),
    ...mapActions(["setSelectedStore"]),

    next() {
      this.loading = true;
      this.setSelectedStore(this.selected);
      this.search = "";
      this.loading = false;
      this.setDialog(false);
    },

    setSearchBoxVisibility(bool) {
      this.searchBoxVisible = bool;
    },

    select(item) {
      if (item === this.selected) this.selected = null;
      else this.selected = item;
    },
  },
};
</script>

<style></style>
