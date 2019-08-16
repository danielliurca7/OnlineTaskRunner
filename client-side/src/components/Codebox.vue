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
      },
      received: false
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

        var cursor_pos = {
            line: this.codemirror.doc.getCursor().line,
            ch: this.codemirror.doc.getCursor().ch
          },
          i = 0,
          start_ch = 0,
          end_ch = 0;

        while (cursor_pos.ch >= 0) {
          i++;

          if (cursor_pos.line === 0) {
            cursor_pos.ch--;
          }

          if (this.code[i] === "\n") {
            cursor_pos.line--;
          }
        }

        for (var j = start; j > 0; j--) {
          if (this.code[j] === "\n") {
            start_ch = start - j - 1;
            break;
          }
        }

        for (j = end; j < this.code.length; j++) {
          if (this.code[j] === "\n") {
            end_ch = j;
            break;
          }
        }

        console.log(i, start, end, end_ch);

        if (i > start && i < end) {
          var lines = this.code.split("\n").map(x => x.length);
          var s = 0,
            line,
            char;

          for (j = 0; j < lines.length; j++) {
            s += lines[j] + 1;

            console.log(j, s, lines);

            if (s > start) {
              s -= lines[j] + 1;
              line = j;
              console.log(start, s);
              char = start - s;
              break;
            }
          }

          console.log(line, char);

          this.cursor_pos = {
            line: line,
            ch: char
          };
        } else if (i >= start) {
          var lines_current = ch.Change.Current.split("\n");
          var lines_previous = ch.Change.Previous.split("\n");

          var line_difference = lines_current.length - lines_previous.length;

          console.log(lines_current, lines_previous);

          var ch_difference = 0;

          if (i <= end_ch) {
            ch_difference =
              lines_current[lines_current.length - 1].length -
              lines_previous[lines_previous.length - 1].length;

            if (lines_current.length === 1) {
              ch_difference += start_ch;
            }

            if (lines_previous.length === 1) {
              ch_difference -= start_ch;
            }
          }

          this.cursor_pos = {
            line: this.codemirror.doc.getCursor().line + line_difference,
            ch: this.codemirror.doc.getCursor().ch + ch_difference
          };
        }

        this.code = this.code.slice(0, start) + ch.Change.Current + last;
        this.last_code = this.code;
        this.received = true;
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
