<template>
  <div class="subject">
    <AuthNavbar />
    <SubjectSidebar />
    <SubjectTable />
  </div>
</template>

<script>
import AuthNavbar from "@/components/AuthNavbar.vue";
import SubjectSidebar from "@/components/SubjectSidebar.vue";
import SubjectTable from "@/components/SubjectTable.vue";

export default {
  name: "subject",
  components: {
    AuthNavbar,
    SubjectSidebar,
    SubjectTable
  }
};
</script>

<style scoped>
li {
  color: yellowgreen;
  display: block;
  text-decoration: none;
  cursor: pointer;
}

li:hover {
  color: yellowgreen;
  background-color: #ddd;
}
</style>
