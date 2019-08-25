<template>
  <div class="lab">
    <AuthNavbar />
    <div>
      <Codebox :file="file" class="w-75 float-right" />
      <FileSidebar @load_file="update_code($event)" class="w-25" />
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
  },
  data() {
    return {
      file: null
    }
  },
  methods: {
    update_code(file) {
      this.file = file;
    }
  }
};
</script>
