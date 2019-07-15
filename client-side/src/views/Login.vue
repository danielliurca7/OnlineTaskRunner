<template>
  <div class="login">
    <Navbar />
    <div class="login-form">
      <b-form @submit="onSubmit">
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
// @ is an alias to /src
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
    ...mapActions(["setUsername"]),

    onSubmit: function(e) {
      e.preventDefault();

      // Authentication with cs.curs api
      // Check credentials

      this.setUsername(this.username);

      this.$router.push("dashboard");
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
