<template>
  <div
    class="file-sidebar"
    v-if="this.treeData !== [] && this.treeData.length > 0"
  >
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
      console.log(object.data);

      switch (itemObject.code) {
        case "NEW_FILE":
          break;
        case "NEW_FOLDER":
          break;
        case "DELETE":
          break;
      }
    }
  }
};
</script>
