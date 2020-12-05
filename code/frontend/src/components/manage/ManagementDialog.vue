<template>
  <v-dialog v-model="dialog" persistent max-width="600px">
    <template v-slot:activator="{ on, attrs }">
      <slot name="button" v-bind="attrs" v-bind:on="on">
        You forgot to slot in a button!
      </slot>
    </template>

    <v-card>
      <v-card-title>
        <span class="headline">{{headline}}</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col
              md="12"
              v-for="field in localFields"
              :key="field.schemaName"
            >
              <v-text-field :label="field.displayName" v-model="field.value" required></v-text-field>
            </v-col>
          </v-row>
        </v-container>
        <small>All fields required</small>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue darken-1" text @click="formClosed()">
          Close
        </v-btn>
        <v-btn color="blue darken-1" text @click="formSaved()"> Save </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "ManagementDialog",
  props: {
    isShown: Boolean,
    type: String,
    fields: Array,
    headline: String
  },
  data () {
      return {
        dialog: false,
        localFields : JSON.parse(JSON.stringify(this.fields)),

      } 
  },
  methods: {
    formSaved() {
      this.dialog = false;
      this.$emit("form-saved", this.localFields);
      //deep copy -__-
      this.localFields = JSON.parse(JSON.stringify(this.fields));
    },
    formClosed() {
      this.dialog = false;
      this.localFields = JSON.parse(JSON.stringify(this.fields));
    }
    

  },
  computed : {
      
  }
};
</script>
