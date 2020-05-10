<template>
  <div
    class="file-sidebar"
    v-if="this.treeData !== [] && this.treeData.length > 0"
  >
    <b-button-toolbar
      key-nav
      aria-label="Toolbar with button groups"
      class="d-flex justify-content-end"
    >
      <b-button-group class="m-2">
        <b-button @click="build">Build</b-button>
        <b-button @click="stop">Stop</b-button>
        <b-button @click="run">Run</b-button>
      </b-button-group>
    </b-button-toolbar>

    <b-tree-view
      class="float-left"
      :data="treeData"
      :contextMenuItems="menuItems"
      @nodeSelect="select"
      @contextMenuItemSelect="contextMenu"
    >
    </b-tree-view>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
import axios from "axios";

export default {
  name: "FileSidebar",
  computed: {
    ...mapGetters(["getUsername", "getToken"])
  },
  mounted() {
    var workspace = {
      owner: this.$route.params.owner,
      course: this.$route.params.course,
      assignmentname: this.$route.params.assignmentname,
      year: parseInt(this.$route.params.year),
      series: this.$route.params.series,
      folder: []
    };

    axios
      .post("http://localhost:3000/api/files", workspace, {
        headers: { Authorization: this.getToken }
      })
      .then(response => {
        let files = response.data;

        this.treeData = this.constructTreeData(files);
        this.addWorkspace({
          workspace: workspace,
          files: files
        });
      })
      .catch(error => {
        console.log(error);
      });
  },
  data() {
    return {
      menuItems: [
        { code: "NEW_FILE", label: "Create File" },
        { code: "NEW_FOLDER", label: "Create Folder" },
        { code: "DELETE", label: "Delete" }
      ],
      treeData: [],
      pathSelected: ""
    };
  },
  methods: {
    ...mapActions(["addWorkspace"]),
    constructTreeData: function(files) {
      let get_filetree_data = file => {
        let path = file.path;
        let path_len = path.length;

        return {
          data: file.data,
          isDir: file.isdir,
          name: path[path_len - 1],
          path: path,
          children: files
            .filter(
              other =>
                JSON.stringify(other.path.slice(0, path_len)) ===
                  JSON.stringify(path) && other.path.length === path_len + 1
            )
            .map(get_filetree_data)
        };
      };

      return files
        .filter(file => file.path.length === 1)
        .map(get_filetree_data);
    },
    build: function() {
      axios
        .post(
          "http://localhost:3000/api/build",
          {
            owner: "",
            course: this.$route.params.course,
            assignmentname: this.$route.params.assignmentname,
            year: parseInt(this.$route.params.year),
            series: this.$route.params.series,
            folder: ["config"]
          },
          {
            headers: { Authorization: this.getToken }
          }
        )
        .then(response => console.log(response.data));
    },
    run: function() {
      axios
        .post(
          "http://localhost:3000/api/run",
          {
            owner: this.$route.params.owner,
            course: this.$route.params.course,
            assignmentname: this.$route.params.assignmentname,
            year: parseInt(this.$route.params.year),
            series: this.$route.params.series,
            folder: []
          },
          {
            headers: { Authorization: this.getToken }
          }
        )
        .then(response => console.log(response.data));
    },
    stop: function() {
      axios
        .post(
          "http://localhost:3000/api/stop",
          {
            owner: this.$route.params.owner,
            course: this.$route.params.course,
            assignmentname: this.$route.params.assignmentname,
            year: parseInt(this.$route.params.year),
            series: this.$route.params.series,
            folder: []
          },
          {
            headers: { Authorization: this.getToken }
          }
        )
        .then(response => console.log(response.data));
    },
    select: function(object, isSelected) {
      if (object.data.isDir) {
        if (isSelected) {
          object.toggle();
          object.deselect();
        }
      } else {
        if (isSelected) {
          this.$emit("load_path", object.data.path);
        }
      }
    },
    contextMenu: function(itemObject, object) {
      switch (itemObject.code) {
        case "NEW_FILE":
          break;
        case "NEW_FOLDER":
          break;
        case "DELETE":
          axios
            .post("http://localhost:8000/api/delete", {
              owner: this.getUsername,
              subject: "APD",
              assignmentname: "Tema de smecherie",
              year: 2019,
              path: object.data.path
            })
            .then(response => {
              if (response.status === 200) {
                axios
                  .post("http://localhost:8000/api/filetree", {
                    owner: this.getUsername,
                    subject: "APD",
                    assignmentname: "Tema de smecherie",
                    year: 2019
                  })
                  .then(response => {
                    this.treeData = response.data;
                  })
                  .catch(error => {
                    console.log(error);
                  });
              }
            })
            .catch(error => {
              console.log(error);
            });

          break;
      }
    }
  }
};
</script>
