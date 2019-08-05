<template>
  <div class="file-sidebar">
    <b-button-toolbar
      key-nav
      aria-label="Toolbar with button groups"
      class="d-flex justify-content-end"
    >
      <b-button-group class="mx-1">
        <b-button>Add</b-button>
        <b-button>Rename</b-button>
        <b-button>Delete</b-button>
      </b-button-group>
    </b-button-toolbar>

    <b-tree-view
      class="float-left"
      :data="treeData"
      :contextMenuItems="menuItems"
    >
    </b-tree-view>
  </div>
</template>

<script>
export default {
  name: "FileSidebar",
  data() {
    return {
      menuItems: [
        { code: "NEW_FILE", label: "Create" },
        { code: "RENAME_FILE", label: "Rename" },
        { code: "DELETE", label: "Delete" }
      ],
      treeData: [
        {
          id: 2,
          name: "Folder 1",
          children: [
            {
              id: 3,
              name: "Folder 2",
              children: [
                {
                  id: 5,
                  name: "File 1"
                }
              ]
            },
            { id: 4, name: "File 2" }
          ]
        }
      ]
    };
  }
};
</script>
