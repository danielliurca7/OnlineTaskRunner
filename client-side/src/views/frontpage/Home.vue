<template>
  <div class="home">
    <Navbar />
    <h1>This is the homepage</h1>
  </div>
</template>

<script>
import { mapActions } from "vuex";
import Navbar from "@/components/Navbar.vue";

export default {
  name: "home",
  components: {
    Navbar
  },
  created() {
    if (
      localStorage.getItem("username") &&
      localStorage.getItem("userinfo") &&
      localStorage.getItem("token")
    ) {
      this.setUsername(localStorage.getItem("username"));
      this.setUserinfo(JSON.parse(localStorage.getItem("userinfo")));
      this.setToken(localStorage.getItem("token"));

      this.$router.push("/dashboard");
    }
  },
  methods: {
    ...mapActions(["setUserinfo", "setUsername", "setToken"])
  }
};
</script>
