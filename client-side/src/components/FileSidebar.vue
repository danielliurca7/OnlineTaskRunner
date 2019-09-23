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
        <b-button @click="clean">Clean</b-button>
        <b-button @click="build">Build</b-button>
        <b-button @click="run">Run</b-button>
        <b-button @click="build_run">Build&Run</b-button>
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
import { mapGetters } from "vuex";
import axios from "axios";

export default {
  name: "FileSidebar",
  mounted() {
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
  },
  computed: {
    ...mapGetters(["getUsername"])
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
    run: function() {
      axios
        .post("http://localhost:8000/api/run", {
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
    build: function() {
      axios
        .post("http://localhost:8000/api/build", {
          owner: this.getUsername,
          subject: "APD",
          assignmentname: "Tema de smecherie",
          year: 2019
        })
        .then(response => {
          console.log(response);
        })
        .catch(error => {
          console.log(error);
        });
    },
    clean: function() {
      axios
        .post("http://localhost:8000/api/clean", {
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
    build_run: function() {
      axios
        .post("http://localhost:8000/api/build", {
          owner: this.getUsername,
          subject: "APD",
          assignmentname: "Tema de smecherie",
          year: 2019
        })
        .then(response => {
          console.log(response.data);

          axios
            .post("http://localhost:8000/api/run", {
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
        })
        .catch(error => {
          console.log(error);
        });
    },
    select: function(object, isSelected) {
      if (object.data.isDir) {
        if (isSelected) {
          object.toggle();
          object.deselect();
        }
      } else {
        if (isSelected) {
          var request = {
            owner: this.getUsername,
            subject: "APD",
            assignmentname: "Tema de smecherie",
            year: 2019,
            path: object.data.path
          };

          axios
            .post("http://localhost:10000/api/get", request)
            .then(response => {
              this.$emit("load_file", {
                info: request,
                data: response.data
              });
            })
            .catch(error => {
              console.log(error);
            });
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
