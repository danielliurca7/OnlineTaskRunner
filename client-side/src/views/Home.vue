<template>
  <div class="home">
    <Navbar />
    <h1>This is the homepage</h1>
  </div>
</template>

<script>
// @ is an alias to /src
import Navbar from "@/components/Navbar.vue";

export default {
  name: "home",
  components: {
    Navbar
  }
};
</script>
