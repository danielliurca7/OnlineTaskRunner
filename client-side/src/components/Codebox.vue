<template>
  <div class="codebox">
    <codemirror
      class="text-left"
      ref="myCm"
      v-model="code"
      :options="cmOptions"
      @input="onCmCodeChange"
      @changes="onCursorActivity"
    >
    </codemirror>
  </div>
</template>

<script>
//import axios from "axios";

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
      }
    };
  },
  computed: {
    ...mapGetters(["getUsername"]),
    codemirror() {
      return this.$refs.myCm.codemirror;
    }
  },
  mounted() {
    this.codemirror.setSize("100%", "51.2em");

    this.$options.sockets.connect = () => {
      this.$socket.emit("subscribe", "test.c");
      console.log("Subscribing to test.c");
    };

    this.$options.sockets.change = ch => {
      if (ch.Fileinfo.Owner != this.getUsername) {
        var last = [];

        var start = ch.Change.Position;
        var end = ch.Change.Position + ch.Change.Previous.length;

        if (this.code.length >= end) {
          last = this.code.slice(end, this.code.length);
        }

        this.code = this.code.slice(0, start) + ch.Change.Current + last;
        this.last_code = this.code;
      }
    };

    this.$socket.emit("subscribe", "test.c");
    console.log("Subscribing to test.c");
  },
  beforeDestroy() {
    this.$socket.emit("unsubscribe", "test.c");
    console.log("Unsubscribing from test.c");
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

      this.$socket.emit("change", {
        fileinfo: {
          owner: this.getUsername,
          subject: "APD",
          assignmentname: "Tema de smecherie",
          name: "test.c",
          year: 2019
        },
        change: {
          position: start,
          current: a,
          previous: b
        }
      });
    },
    onCursorActivity() {
      console.log("cursor");
    }
  }
};
</script>
