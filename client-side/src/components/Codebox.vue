<template>
  <div class="codebox">
    <codemirror
      class="text-left"
      ref="myCm"
      v-model="code"
      :options="cmOptions"
      @input="onCmCodeChange"
      @cursorActivity="onCursorActivity"
    >
    </codemirror>
  </div>
</template>

<script>
import axios from "axios";

import { mapGetters } from "vuex";

import { codemirror } from "vue-codemirror";

import "codemirror/lib/codemirror.css";

import "codemirror/mode/clike/clike.js";
import "codemirror/theme/base16-dark.css";

export default {
  name: "Codebox",
  components: {
    codemirror
  },
  props: ["file", "height"],
  watch: {
    file: function(newVal) {
      this.code = newVal.data;
      this.last_code = this.code;
    },
    height: function(height) {
      this.codemirror.setSize("100%", height + "px");
    }
  },
  data() {
    return {
      code: "",
      last_code: "",
      cmOptions: {
        tabSize: 4,
        indentUnit: 4,
        mode: "text/x-csrc",
        theme: "base16-dark",
        lineNumbers: true,
        line: true,
        lineWrapping: true,
        styleActiveLine: true
      },
      cursor_pos: {
        line: 0,
        ch: 0
      },
      received: false,
      code_loaded: false
    };
  },
  computed: {
    ...mapGetters(["getUsername"]),
    codemirror() {
      return this.$refs.myCm.codemirror;
    }
  },
  mounted() {
    axios
      .post("http://localhost:8000/api/get", {
        owner: this.getUsername,
        subject: "APD",
        assignmentname: "Tema de smecherie",
        year: 2019
      })
      .then(response => {
        console.log(response.data);
      })
      .catch(error => {
        console.log(error);
      });
  },
  beforeDestroy() {
    axios
      .post("http://localhost:8000/api/clear", {
        owner: this.getUsername,
        subject: "APD",
        assignmentname: "Tema de smecherie",
        year: 2019
      })
      .then(response => {
        console.log(response.data);
      })
      .catch(error => {
        console.log(error);
      });
  },
  methods: {
    onCmCodeChange() {
      var a = this.code;
      var b = this.last_code;

      if (a === b) {
        return;
      }

      var i, j, start;

      for (i = 0; i <= Math.min(a.length, b.length); ) {
        if (a[i] === b[i]) {
          i++;
        } else {
          break;
        }
      }

      start = i;

      a = a.slice(i, a.length);
      b = b.slice(i, b.length);

      for (i = a.length - 1, j = b.length - 1; i >= 0 && j >= 0; ) {
        if (a[i] === b[j]) {
          i--;
          j--;
        } else {
          break;
        }
      }

      a = a.slice(0, i + 1);
      b = b.slice(0, j + 1);

      this.last_code = this.code;

      axios
        .post("http://localhost:8000/api/change", {
          fileinfo: this.file.info,
          change: {
            position: start,
            current: a,
            previous: b
          }
        })
        .then(response => {
          //console.log(response)
        })
        .catch(error => {
          console.log(error);
        });
    },
    onCursorActivity() {
      if (!this.received) {
        this.cursor_pos = {
          line: this.codemirror.doc.getCursor().line,
          ch: this.codemirror.doc.getCursor().ch
        };
      } else {
        this.codemirror.doc.setCursor(this.cursor_pos);
        this.received = false;
      }
    }
  }
};
</script>
