<template>
  <div class="grade">
    <AuthNavbar />
    <GradeSidebar />

    <ul class="list-group mt-3 float-right w-50 ml-3">
      <li
        class="list-group-item m-1"
        v-for="assignment in getAssignmentsToGrade($route.params.course)"
        :key="assignment.name"
        @click="grade(assignment.name)"
      >
        {{ assignment.name }}
      </li>
    </ul>
  </div>
</template>

<script>
import { mapGetters } from "vuex";

import AuthNavbar from "@/components/AuthNavbar.vue";
import GradeSidebar from "@/components/GradeSidebar.vue";

export default {
  name: "gradeoverview",
  components: {
    AuthNavbar,
    GradeSidebar
  },
  computed: {
    ...mapGetters(["getAssignmentsToGrade"])
  },
  methods: {
    grade(assignment) {
      this.$router.push(
        "/grade/" +
          this.$route.params.course +
          "/" +
          this.$route.params.series +
          "/" +
          this.$route.params.year +
          "/" +
          assignment
      );
    }
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
