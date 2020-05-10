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
    ...mapActions([
      "setUserinfo",
      "setUsername",
      "setToken",
      "setStudentCourses",
      "setAssistantCourses"
    ]),
    onSubmit: function() {
      axios
        .post("http://localhost:3000/api/authenticate", {
          username: this.username,
          password: this.password
        })
        .then(response => {
          if (response.status !== 400) {
            let base64Url = response.data.split(".")[1];
            let base64 = base64Url.replace("-", "+").replace("_", "/");
            let decodedToken = JSON.parse(
              Buffer.from(base64, "base64").toString("binary")
            );

            this.setUsername(this.username);
            this.setUserinfo(decodedToken.payload);
            this.setToken(response.data);

            // if the "Remember me" box is active, create a new cookie for username
            if (this.remember) {
              localStorage.setItem("username", this.username);
              localStorage.setItem(
                "userinfo",
                JSON.stringify(decodedToken.payload)
              );
              localStorage.setItem("token", response.data);
            }

            axios
              .get("http://localhost:3000/api/student/" + this.username, {
                headers: { Authorization: response.data }
              })
              .then(response => {
                var data = response.data;

                if (data != null) {
                  data.courses.map(course => {
                    course.labs.map(
                      lab => (lab.deadline = new Date(lab.deadline))
                    );

                    course.homeworks.map(
                      homework =>
                        (homework.deadline = new Date(homework.deadline))
                    );
                  });

                  this.setStudentCourses(data.courses);
                }
              });

            axios
              .get("http://localhost:3000/api/assistant/" + this.username, {
                headers: { Authorization: response.data }
              })
              .then(response => {
                var data = response.data;

                if (data != null) {
                  this.setAssistantCourses(data.courses);
                }
              });

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
