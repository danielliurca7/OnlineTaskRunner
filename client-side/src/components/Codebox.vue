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
import { mapGetters, mapActions } from "vuex";
import axios from "axios";

import { codemirror } from "vue-codemirror";

import "codemirror/lib/codemirror.css";

import "codemirror/mode/clike/clike.js";
import "codemirror/theme/base16-dark.css";

export default {
  name: "Codebox",
  components: {
    codemirror
  },
  props: ["path", "height"],
  watch: {
    path: function(newVal, oldVal) {
      if (oldVal != null) {
        this.changeFile({
          workspace: this.workspace,
          path: oldVal,
          newVal: this.code
        });
      }

      var data = this.getWorkspaceFile(
        this.$route.params.course,
        this.$route.params.series,
        this.$route.params.year,
        this.$route.params.assignmentname,
        this.$route.params.owner,
        newVal
      ).data;

      this.code = data;
      this.last_code = this.code;
    },
    height: function(height) {
      this.codemirror.setSize("100%", height + "px");
    }
  },
  data() {
    return {
      workspace: {},
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
    ...mapGetters([
      "getUserinfo",
      "getUsername",
      "getToken",
      "getWorkspaces",
      "getWorkspaceFile"
    ]),
    codemirror() {
      return this.$refs.myCm.codemirror;
    }
  },
  mounted() {
    this.workspace = {
      owner: this.$route.params.owner,
      course: this.$route.params.course,
      assignmentname: this.$route.params.assignmentname,
      year: parseInt(this.$route.params.year),
      series: this.$route.params.series,
      folder: []
    };

    this.codemirror.setSize("100%", this.height + "px");

    this.$options.sockets.change = change => {
      let ch = change.update;

      if (change.sender != this.getUsername) {
        var sameFile =
          this.workspace.owner === ch.workspace.owner &&
          this.workspace.course === ch.workspace.course &&
          this.workspace.series === ch.workspace.series &&
          this.workspace.year === ch.workspace.year &&
          this.workspace.assignmentname === ch.workspace.assignmentname &&
          JSON.stringify(this.workspace.folder.concat(this.path)) ===
            JSON.stringify(ch.workspace.folder.concat(ch.path));

        var code;

        if (sameFile) {
          code = this.code;
        } else {
          code = this.getWorkspaceFile(
            ch.workspace.course,
            ch.workspace.series,
            ch.workspace.year,
            ch.workspace.assignmentname,
            ch.workspace.owner,
            ch.workspace.folder.concat(ch.path)
          ).data;
        }

        var last = [];
        var start = ch.change.position;
        var end = ch.change.position + ch.change.previous.length;

        if (code.length >= end) {
          last = code.slice(end, code.length);
        }

        if (sameFile) {
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
            if (code[i] === "\n") {
              cursor_pos.line--;
            }
          }

          for (var j = start; j > 0; j--) {
            if (code[j] === "\n") {
              start_ch = start - j - 1;
              break;
            }
          }

          for (j = end; j < code.length; j++) {
            if (code[j] === "\n") {
              end_ch = j;
              break;
            }
          }

          if (i > start && i < end) {
            var lines = code.split("\n").map(x => x.length);
            var s = 0,
              line,
              char;
            for (j = 0; j < lines.length; j++) {
              s += lines[j] + 1;
              if (s > start) {
                s -= lines[j] + 1;
                line = j;
                char = start - s;
                break;
              }
            }
            this.cursor_pos = {
              line: line,
              ch: char
            };
          } else if (i > start) {
            var lines_current = ch.change.current.split("\n");
            var lines_previous = ch.change.previous.split("\n");

            var line_difference = lines_current.length - lines_previous.length;

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
        }

        code = code.slice(0, start) + ch.change.current + last;

        if (sameFile) {
          this.code = code;
          this.last_code = this.code;
          this.received = true;
        } else {
          this.changeFile({
            workspace: ch.workspace,
            path: ch.path,
            newVal: code
          });
        }
      }
    };

    this.$socket.emit("subscribe", {
      token: this.getToken,
      workspace: {
        assignmentname: this.$route.params.assignmentname,
        course: this.$route.params.course,
        folder: [],
        owner: this.$route.params.owner,
        series: this.$route.params.series,
        year: parseInt(this.$route.params.year)
      }
    });
  },
  beforeDestroy() {
    this.$socket.emit("unsubscribe", this.workspace);

    // axios.post("http://localhost:3000/api/stop", this.workspace, {
    //   headers: { Authorization: this.getToken }
    // });

    axios.post("http://localhost:3000/api/commit", this.workspace, {
      headers: { Authorization: this.getToken }
    });

    if (this.path != null) {
      this.changeFile({
        workspace: this.workspace,
        path: this.path,
        newVal: this.code
      });
    }
  },
  methods: {
    ...mapActions(["changeFile"]),
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
        token: this.getToken,
        sender: this.getUsername,
        update: {
          workspace: {
            owner: this.$route.params.owner,
            course: this.$route.params.course,
            assignmentname: this.$route.params.assignmentname,
            year: parseInt(this.$route.params.year),
            series: this.$route.params.series,
            folder: []
          },
          path: this.path,
          change: {
            position: start,
            current: a,
            previous: b
          }
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
