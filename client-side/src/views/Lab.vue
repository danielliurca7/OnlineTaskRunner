<template>
  <div class="lab">
    <div ref="navbar">
      <AuthNavbar />
    </div>
    <div>
      <Codebox :file="file" :height="height" class="w-75 float-right" />
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
      file: null,
      height: 0
    };
  },
  methods: {
    update_code(file) {
      this.file = file;
    },
    resize() {
      this.height = window.innerHeight - this.$refs.navbar.clientHeight;
    }
  },
  mounted() {
    window.addEventListener("resize", this.resize);
    this.resize();
  },
  beforeDestroy() {
    window.removeEventListener("resize", this.resize);
  }
};
</script>
