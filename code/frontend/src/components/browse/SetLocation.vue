<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" persistent max-width="290">
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          class="ma-2"
          color="primary"
          dark
          v-bind="attrs"
          v-on="on"
          text
          icon
        >
          <v-icon>mdi-map-marker</v-icon>
        </v-btn>
      </template>

      <v-card class="mx-auto" max-width="500">
        <v-toolbar flat color="transparent">
          <v-toolbar-title>Select your location</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-btn icon @click="$refs.search.focus()">
            <v-icon>mdi-magnify</v-icon>
          </v-btn>
        </v-toolbar>

        <v-container class="py-0">
          <v-row align="center" justify="start">
            <v-col
              v-for="(selection, i) in selections"
              :key="selection.text"
              class="shrink"
            >
              <v-chip
                :disabled="loading"
                close
                @click:close="selected.splice(i, 1)"
              >
                <v-icon left v-text="selection.icon"></v-icon>
                {{ selection.text }}
              </v-chip>
            </v-col>

            <v-col v-if="!allSelected" cols="12">
              <v-text-field
                ref="search"
                v-model="search"
                full-width
                hide-details
                label="Search"
                single-line
              ></v-text-field>
            </v-col>
          </v-row>
        </v-container>

        <v-divider v-if="!allSelected"></v-divider>

        <v-list>
          <template v-for="item in categories">
            <v-list-item
              two-line
              v-if="!selected.includes(item)"
              v-show="!onlySelectOne"
              :key="item.text"
              :disabled="loading"
              @click="selected.push(item)"
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
          </template>
        </v-list>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            :disabled="!selected.length"
            :loading="loading"
            color="purple"
            text
            @click="next"
          >
            Set Location
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
export default {
  //   data() {
  //     return {
  //       dialog: false,
  //     };
  //   },

  data: () => ({
    items: [
      {
        text: "Shopperz",
        icon: "mdi-basket",
        location: "123 Bay Street",
      },
      {
        text: "NotLongos",
        icon: "mdi-basket",
        location: "420 Victor Philp Road",
      },
      {
        text: "JoeUnfresh",
        icon: "mdi-basket",
        location: "666 Allen Avenu",
      },
      {
        text: "Matt's awesome store",
        icon: "mdi-basket",
        location: "69 Troyer Avenue",
      },
      {
        i: 5,
        text: "Kinny's Konvenience",
        icon: "mdi-basket",
        location: "404 Unknown Street",
      },
    ],
    loading: false,
    search: "",
    selected: [],
    dialog: false,
  }),

  computed: {
    ...mapGetters(["getLocation", "getSelectedStore"]),
    nonEmptyStore() {
      return this.getSelectedStore !== null;
    },
    allSelected() {
      return this.selected.length === this.items.length;
    },
    onlySelectOne() {
      return this.selected.length === 1;
    },
    categories() {
      const search = this.search.toLowerCase();

      if (!search) return this.items;

      return this.items.filter((item) => {
        const text = item.text.toLowerCase();

        return text.indexOf(search) > -1;
      });
    },
    selections() {
      const selections = [];

      for (const selection of this.selected) {
        selections.push(selection);
      }

      return selections;
    },
  },

  watch: {
    selected() {
      this.search = "";
    },
  },

  methods: {
    ...mapActions(["setLocation", "setSelectedStore"]),
    next() {
      this.loading = true;

      setTimeout(() => {
        this.search = "";
        this.selected = [];
        this.loading = false;
        this.dialog = false;
      }, 2000);
      this.closeDialog();
    },

    closeDialog() {
      return {
        dialog: false,
      };
    },
  },
};
</script>

<style></style>
