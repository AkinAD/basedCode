<template>
    <v-container id="signinup-form" class="fill-height">
      <Notification
        :message="snackbarMessage"
        :snackbar="snackbar"
        :type="snackbarType"
      />
      <v-row align="center" justify="center" no-gutters>
        <v-col cols="12" sm="8" md="8" class="">
          <v-card class="evelation-12 card">
            <v-window v-model="step">
              <!--SignIn-->
              <v-window-item :value="1">
                <v-row class="">
                  <v-col cols="12" md="8" class="pt-6 pb-6">
                    <v-card-text>
                      <v-form class="signup-form-form" @submit.prevent="signin">
                        <h1
                          class="text-center display-1 mb-10"
                          :class="`${bgColor}--text`"
                        >
                          Welcome Back
                        </h1>
                        <v-text-field
                          id="username"
                          v-model="login"
                          label="Username"
                          name="Username"
                          append-icon="person"
                          type="text"
                          :color="bgColor"
                        />
                        <v-text-field
                          id="password"
                          v-model="password"
                          label="Password"
                          name="Password"
                          append-icon="lock"
                          type="password"
                          :color="bgColor"
                        />
                        <div class="text-center">
                          <a
                            href="#"
                            class="mt-3 overline no-text-decoration"
                            :class="`${bgColor}--text`"
                            @click="step = 3"
                          >
                            Forgot your password?
                          </a>
                        </div>
                        <div class="text-center mt-6">
                          <v-btn type="submit" large :color="bgColor" dark
                            >Sign In</v-btn
                          >
                        </div>
                      </v-form>
                    </v-card-text>
                  </v-col>
                  <v-col
                    cols="12"
                    md="4"
                    class="darken-2 vcenter"
                    :class="`${bgColor}`"
                  >
                    <div>
                      <v-card-text :class="`${fgColor}--text`">
                        <h1 class="text-center headline mb-3">First time?</h1>
                        <h5 class="text-center overline mb-3">
                          Sign Up to continue
                        </h5>
                      </v-card-text>
                      <div class="text-center mb-6">
                        <v-btn dark outlined @click="step = 2">Sign Up</v-btn>
                      </div>
                    </div>
                  </v-col>
                </v-row>
              </v-window-item>
              <!--SignUp-->
              <v-window-item :value="2">
                <v-row class="fill-height">
                  <v-col
                    cols="12"
                    md="4"
                    class="darken-2 vcenter"
                    :class="`${bgColor}`"
                  >
                    <div>
                      <v-card-text :class="`${fgColor}--text`">
                        <h1 class="text-center headline mb-3">
                          Already a user?
                        </h1>
                        <h5 class="text-center overline mb-3">
                          Please Sign In
                        </h5>
                      </v-card-text>
                      <div class="text-center mb-6">
                        <v-btn dark outlined @click="step = 1">Sign In</v-btn>
                      </div>
                    </div>
                  </v-col>
                  <v-col cols="12" md="8" class="pt-6 pb-6">
                    <v-card-text>
                      <h1
                        class="text-center display-1 mb-10"
                        :class="`${bgColor}--text`"
                      >
                        Sign Up
                      </h1>
                      <v-form class="signup-form-form" @submit.prevent="signup">
                        <v-text-field
                          id="username"
                          v-model="username"
                          label="Username"
                          name="username"
                          append-icon="person"
                          type="text"
                        />
                        <v-text-field
                          id="email"
                          v-model="email"
                          label="eMail"
                          name="email"
                          append-icon="email"
                          type="email"
                        />
                        <v-text-field
                          id="password"
                          v-model="password"
                          label="Password"
                          name="password"
                          append-icon="lock"
                          type="password"
                        />
                        <div class="text-center mt-6">
                          <v-btn type="submit" large :color="bgColor" dark>
                            Sign Up</v-btn
                          >
                        </div>
                      </v-form>
                    </v-card-text>
                  </v-col>
                </v-row>
              </v-window-item>
              <!--PW Rest-->
              <v-window-item :value="3">
                <v-row class="fill-height">
                  <v-col
                    cols="12"
                    md="4"
                    class="darken-2 vcenter"
                    :class="`${bgColor}`"
                  >
                    <div>
                      <v-card-text :class="`${fgColor}--text`">
                        <h1 class="text-center headline mb-3">
                          Already a user?
                        </h1>
                        <h5 class="text-center overline mb-3">
                          Please Sign In
                        </h5>
                      </v-card-text>
                      <div class="text-center mb-6">
                        <v-btn dark outlined @click="step = 1">Sign In</v-btn>
                      </div>
                    </div>
                  </v-col>
                  <v-col cols="12" md="8" class="pt-6 pb-6">
                    <v-card-text>
                      <v-form class="signup-form-form">
                        <h1
                          class="text-center display-1 mb-10"
                          :class="`${bgColor}--text`"
                        >
                          Reset Password
                        </h1>
                        <v-text-field
                          id="login"
                          v-model="login"
                          label="Username / eMail"
                          name="login"
                          append-icon="person / email"
                          type="text"
                          :color="bgColor"
                          class="v-input__icon--double"
                        />
                        <div class="text-center mt-6">
                          <v-btn large :color="bgColor" dark
                            >Reset Password</v-btn
                          >
                        </div>
                      </v-form>
                    </v-card-text>
                  </v-col>
                </v-row>
              </v-window-item>
            </v-window>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
</template>

<script>
import Notification from "../components/Notification";

export default {
  name: "Signupform",
  components: {
    Notification,
  },
  props: {
    source: {
      type: String,
      default: "",
    },
    bgColor: {
      type: String,
      default: "blue",
    },
    fgColor: {
      type: String,
      default: "white",
    },
  },
  async fetch({ store, error }, user) {
    try {
      await store.dispatch("users/signupUser", user);
    } catch (e) {
      error({
        statusCode: 503,
        message: "Unable to sign up user. Please try again later.",
      });
    }
  },
  data: () => ({
    step: 1,
    username: "",
    email: "",
    password: "",
    login: "",
    snackbarType: "success",
    snackbarMessage: "",
    snackbar: false,
  }),
  methods: {
    signup() {
      this.$auth
        .signup({
          data: {
            user: {
              username: this.username,
              email: this.email,
              password: this.password,
            },
          },
        })
        .catch((e) => {
          this.error = e + "";
        });
    },
    async signin() {
      try {
        const response = await this.$auth.loginWith("local", {
          data: {
            login: this.login,
            password: this.password,
          },
        });
        this.snackbarType = response.data.type;
        this.snackbarMessage = response.data.message;
        this.snackbar = true;
      } catch (err) {
        this.snackbarType = "error";
        this.snackbarMessage = "Error signing you in";
        this.snackbar = true;
      }
    },
  },
};
</script>

<style scoped lang="scss">
.v-input__icon--double .v-input__icon {
  margin-left: -4.25rem !important;
}
a.no-text-decoration {
  text-decoration: none;
}
#signinup-form {
  max-width: 75rem;
}
.signup-form-form {
  max-width: 23rem;
  margin: 0 auto;
}
.card {
  overflow: hidden;
}
.vcenter {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>