<template>
  <v-container>
    <div v-if="resize">
      <v-sheet :height="sheetHeight">
        <v-container class="fill-height">
          <v-row align="center" justify="center">
            <h1 class="text-center display-3 hidden-xs-only">
              {{ text }}
            </h1>
            <h1 class="text-center display-1 hidden-sm-and-up">
              {{ text }}
            </h1>
          </v-row>
        </v-container>
      </v-sheet>
    </div>
    <div v-else>
      <v-sheet :height="size">
        <v-container class="fill-height">
          <v-row align="center" justify="center">
            <h1 class="text-center display-2 hidden-xs-only">
              {{ text }}
            </h1>
            <h1 class="text-center display-1 hidden-sm-and-up">
              {{ text }}
            </h1>
          </v-row>
        </v-container>
      </v-sheet>
    </div>
  </v-container>
</template>

<script>
export default {
  name: "Banner",
  props: {
    text: String,
    resize: Boolean,
    size: String,
    display: String
  },
  computed: {
    sheetHeight() {
      switch (this.$vuetify.breakpoint.name) {
        case "xs":
          return "200px";
        case "sm":
          return "300px";
        case "md":
          return "300px";
        case "lg":
          return "300px";
        case "xl":
          return "350px";
      }
      return "300px";
    },
  },
};
</script>

<style></style>
