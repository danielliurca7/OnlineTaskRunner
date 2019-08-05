<template>
  <div class="homework">
    <AuthNavbar />
    <div>
      <Codebox class="w-75 float-right" />
      <FileSidebar class="w-25" />
    </div>
  </div>
</template>

<script>
import AuthNavbar from "@/components/AuthNavbar.vue";
import FileSidebar from "@/components/FileSidebar.vue";
import Codebox from "@/components/Codebox.vue";

export default {
  name: "homework",
  components: {
    AuthNavbar,
    FileSidebar,
    Codebox
  }
};
</script>
