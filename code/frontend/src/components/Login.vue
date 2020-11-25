<template>
  <v-container fluid class="pa-6">
    <v-row align="center" justify="center">
      <amplify-authenticator>
        <amplify-sign-in
          header-text="My Custom Sign In Text"
          slot="sign-in"
          :formFields="formFields"
        ></amplify-sign-in>
      </amplify-authenticator>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "login",
  data() {
    return {
      formFields: [ //THIS BULLSHIT SHOULD WORK BUT IT DOESNT >:(
        {
          type: "email",
          label: "Custom email Label",
          placeholder: "custom email placeholder",
          required: true,
        },
        {
          type: "password",
          label: "Custom Password Label",
          placeholder: "custom password placeholder",
          required: true,
        }
      ],
    };
  },
};
</script>

<style>
</style>