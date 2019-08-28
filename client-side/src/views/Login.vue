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
// @ is an alias to /src
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
    ...mapActions(["setUsername"]),
    onSubmit: function() {
      // authentication with cs.curs api
      // check credentials
      var year = 2018;
      var base_url = "https://acs.curs.pub.ro/" + year + "/login/token.php";
      var params = {
        username: this.username,
        password: this.password,
        service: "moodle_mobile_app"
      };

      var queryString = Object.keys(params)
        .map(key => key + "=" + params[key])
        .join("&");

      var url = base_url + "?" + queryString;

      axios
        .post(url)
        .then(response => {
          console.log(response.data);
        })
        .catch(error => {
          console.log(error);
        });

      this.setUsername(this.username);

      // if the "Remember me" box is active, create a new cookie for username
      if (this.remember) {
        this.$cookie.set("username", this.username, 1);
      }

      // go to dashboard
      this.$router.push("/dashboard");
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
