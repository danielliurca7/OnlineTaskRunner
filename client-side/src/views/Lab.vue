<template>
  <div class="lab">
    <AuthNavbar />
    <FileSidebar />
  </div>
</template>

<script>
import AuthNavbar from "@/components/AuthNavbar.vue";
import FileSidebar from "@/components/FileSidebar.vue";

export default {
  name: "lab",
  components: {
    AuthNavbar,
    FileSidebar
  }
};
</script>
