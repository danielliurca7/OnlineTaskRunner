<template>
  <div class="dashboard">
    <AuthNavbar />
    <SubjectSidebar />
    <FrontpageTable />
  </div>
</template>

<script>
import AuthNavbar from "@/components/AuthNavbar.vue";
import SubjectSidebar from "@/components/SubjectSidebar.vue";
import FrontpageTable from "@/components/FrontpageTable.vue";

export default {
  name: "dashboard",
  components: {
    AuthNavbar,
    SubjectSidebar,
    FrontpageTable
  }
};
</script>
