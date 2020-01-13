<template>
  <div>
    <AuthNavbar />
    <SubjectSidebar />
    <h1>Insert subject overview</h1>
  </div>
</template>

<script>
import AuthNavbar from "@/components/AuthNavbar.vue";
import SubjectSidebar from "@/components/SubjectSidebar.vue";

export default {
  name: "SubjectOverview",
  components: {
    AuthNavbar,
    SubjectSidebar
  }
};
</script>
