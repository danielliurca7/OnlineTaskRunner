<template>
  <div class="homework">
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
  name: "homework",
  components: {
    AuthNavbar,
    FileSidebar,
    Codebox
  },
  data() {
    return {
      code: "",
      cmOptions: {
        tabSize: 4,
        mode: "text/javascript",
        theme: "base16-dark",
        lineNumbers: true,
        line: true,
        lineWrapping: true,
        styleActiveLine: true
      }
    };
  }
};
</script>
