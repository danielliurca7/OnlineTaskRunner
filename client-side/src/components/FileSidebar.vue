<template>
  <div class="file-sidebar">
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
        { code: "DELETE", label: "Delete" },
        { code: "NEW_FILE", label: "Create File" }
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
