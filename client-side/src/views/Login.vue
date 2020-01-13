<template>
  <div class="login">
    <Navbar />
    <div class="login-form">
      <b-form @submit.prevent="onSubmit">
        <b-form-group>
          <b-form-input
            v-model="username"
            type="text"
            placeholder="Username"
            required
          />
        </b-form-group>

        <b-form-group>
          <b-form-input
            v-model="password"
            type="password"
            placeholder="Password"
            required
          />
        </b-form-group>

        <div class="text-center">
          <b-button type="submit" variant="primary" block>Sign In</b-button>
        </div>

        <div class="float-left mt-3">
          <b-form-checkbox
            v-model="remember"
            value="true"
            unchecked-value="false"
          >
            Remember me
          </b-form-checkbox>
        </div>
      </b-form>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Navbar from "@/components/Navbar.vue";
import { mapActions } from "vuex";

export default {
  name: "login",
  components: {
    Navbar
  },
  data() {
    return {
      username: "",
      password: "",
      remember: false
    };
  },
  methods: {
    ...mapActions(["setUsername", "setToken", "setType"]),
    onSubmit: function() {
      if (this.$cookie.get("username") !== null) {
        this.setUsername(this.$cookie.get("username"));

        // go to dashboard
        this.$router.push("/dashboard");
      }

      axios
        .post("http://localhost:7000/api/authenticate", {
          username: this.username,
          passwordHash: this.CryptoJS.SHA256(this.password).toString()
        })
        .then(response => {
          console.log(response.data);

          if (response.data !== "Invalid password") {
            this.setUsername(this.username);
            this.setToken(response.data);

            axios
              .post("http://localhost:7000/api/type", {
                username: this.username
              })
              .then(response => {
                this.setType(response.data);
                this.$cookie.set("type", response.data, 1);
              })
              .catch(error => {
                console.log(error.message);
              });

            // if the "Remember me" box is active, create a new cookie for username
            if (this.remember) {
              this.$cookie.set("username", this.username, 1);
              this.$cookie.set("token", response.data, 1);
            }

            // go to dashboard
            this.$router.push("/dashboard");
          }
        })
        .catch(error => {
          console.log(error.message);
        });
    }
  }
};
</script>

<style scoped>
.login-form {
  margin: auto;
  margin-top: 25vh;
  width: 50%;
}
</style>
