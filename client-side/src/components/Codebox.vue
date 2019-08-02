<template>
  <div class="codebox">
    <codemirror
      class="text-left"
      ref="myCm"
      v-model="code"
      :options="cmOptions"
    >
    </codemirror>
  </div>
</template>

<script>
import { codemirror } from "vue-codemirror";

import "codemirror/lib/codemirror.css";

import "codemirror/mode/clike/clike.js";
import "codemirror/theme/base16-dark.css";

export default {
  name: "Codebox",
  components: {
    codemirror
  },
  data() {
    return {
      code: "",
      cmOptions: {
        tabSize: 4,
        mode: "text/x-csrc",
        theme: "base16-dark",
        lineNumbers: true,
        line: true,
        lineWrapping: true,
        styleActiveLine: true
      }
    };
  },
  computed: {
    codemirror() {
      return this.$refs.myCm.codemirror;
    }
  },
  mounted() {
    this.codemirror.setSize("100%", "51.2em");
  }
};
</script>
