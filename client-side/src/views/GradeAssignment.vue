<template>
  <div class="grade">
    <AuthNavbar />
    <GradeSidebar />
    <GradeTable />
  </div>
</template>

<script>
import { mapGetters } from "vuex";

import AuthNavbar from "@/components/AuthNavbar.vue";
import GradeSidebar from "@/components/GradeSidebar.vue";
import GradeTable from "@/components/GradeTable.vue";

export default {
  name: "gradeassignment",
  components: {
    AuthNavbar,
    GradeSidebar,
    GradeTable
  },
  computed: {
    ...mapGetters(["getAssignmentsToGrade"])
  },
  methods: {
    grade(subject, assignment) {
      this.$router.push("/grade/" + subject + "/" + assignment);
    }
  }
};
</script>
