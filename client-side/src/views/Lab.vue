<template>
  <div class="lab">
    <AuthNavbar />
    <div>
      <FileSidebar />
      <Codebox class="w-75 float-right" />
    </div>
  </div>
</template>

<script>
import AuthNavbar from "@/components/AuthNavbar.vue";
import FileSidebar from "@/components/FileSidebar.vue";
import Codebox from "@/components/Codebox.vue";

export default {
  name: "lab",
  components: {
    AuthNavbar,
    FileSidebar,
    Codebox
  }
};
</script>
