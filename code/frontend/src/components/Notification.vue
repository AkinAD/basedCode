<template>
  <div class="text-center">
    <v-snackbar
      v-model="notification.show"
      :timeout="notification.timeout"
      :color="notification.type"
    >
      {{ notification.message }}

      <template v-slot:action="{ attrs }">
        <v-btn dark text v-bind="attrs" @click="notification.show = false">
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </div>
</template>

<script>
import { mapState } from "vuex";
export default {
  name: "Notification",
  computed: {
    ...mapState(["notification"])
  }
};
</script>