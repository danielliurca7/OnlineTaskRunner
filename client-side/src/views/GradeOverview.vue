<template>
  <div class="grade">
    <AuthNavbar />
    <GradeSidebar />
    <h1>Insert grade overview</h1>
  </div>
</template>

<script>
import AuthNavbar from "@/components/AuthNavbar.vue";
import GradeSidebar from "@/components/GradeSidebar.vue";

export default {
  name: "gradeoverview",
  components: {
    AuthNavbar,
    GradeSidebar
  }
};
</script>
