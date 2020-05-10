<template>
  <div class="profile">
    <AuthNavbar />
    <h1>This is the profile and preferences page</h1>
  </div>
</template>

<script>
import AuthNavbar from "@/components/AuthNavbar.vue";

export default {
  name: "profile",
  components: {
    AuthNavbar
  }
};
</script>
