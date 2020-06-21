<template>
  <div class="workspace">
    <div ref="navbar">
      <AuthNavbar />
    </div>
    <div>
      <Codebox :path="current_path" :height="height" class="w-75 float-right" />
      <FileSidebar @load_path="load_path($event)" class="w-25" />
    </div>
  </div>
</template>

<script>
import AuthNavbar from "@/components/AuthNavbar.vue";
import FileSidebar from "@/components/FileSidebar.vue";
import Codebox from "@/components/Codebox.vue";

export default {
  name: "workspace",
  components: {
    AuthNavbar,
    FileSidebar,
    Codebox
  },
  data() {
    return {
      current_path: null,
      height: 0
    };
  },
  methods: {
    load_path(path) {
      this.current_path = path;
    },
    resize() {
      this.height = window.innerHeight - this.$refs.navbar.clientHeight - 86;
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
