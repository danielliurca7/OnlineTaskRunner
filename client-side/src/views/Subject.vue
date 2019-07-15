<template>
  <div class="subject">
    <AuthNavbar />
    {{ $route.params.subject }}
  </div>
</template>

<script>
import AuthNavbar from "@/components/AuthNavbar.vue";

export default {
  name: "subject",
  components: {
    AuthNavbar
  }
};
</script>

<style scoped>
</style>
