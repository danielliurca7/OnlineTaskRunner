<template>
  <div class="file-sidebar">
    <b-tree-view
      class="border w-25 mt-3"
      :data="treeData"
      :contextMenuItems="menuItems"
    >
    </b-tree-view>
  </div>
</template>

<script>
export default {
  data() {
    return {
      menuItems: [{code: "DELETE", label: "Delete"}, { code: 'NEW_FILE', label: 'Create File' }],
      treeData: [{"id": 2, "name": "Venus" , "children": [{"id": 3, "name": "Neptune", "children": [{"id": 5, "name": "Mars"}]}, {"id": 4, "name": "Stratus"} ] } ]
    };
  }
};
</script>

<style>
</style>
