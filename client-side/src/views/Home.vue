<template>
  <div class="home">
    <Navbar />
    <h1>This is the homepage</h1>
  </div>
</template>

<script>
import Navbar from "@/components/Navbar.vue";

export default {
  name: "home",
  components: {
    Navbar
  },
  created() {
    // if there is a cookie for username, jump to dashboard
    if (
      this.$cookie.get("username") !== undefined &&
      this.$cookie.get("username") !== null
    ) {
      this.$router.push("/dashboard");
    }
  }
};
</script>
