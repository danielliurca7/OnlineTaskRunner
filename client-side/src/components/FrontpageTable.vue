<template>
  <div class="container">
    <table
      class="table table-bordered table-condensed table-striped table-hover"
    >
      <tr>
        <th>Subject</th>
        <th>Assignment</th>
        <th>Deadline</th>
      </tr>
      <tr
        class="table-row"
        v-for="(assignment, index) in getAssignments"
        :key="index"
        @click="openAssignment(assignment)"
      >
        <td>{{ assignment.subjectName }}</td>
        <td>{{ assignment.name }}</td>
        <td>
          {{
            assignment.deadline.getDate() +
              "-" +
              (assignment.deadline.getMonth() + 1) +
              "-" +
              assignment.deadline.getFullYear()
          }}
        </td>
      </tr>
    </table>
  </div>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "FrontpageTable",
  computed: {
    ...mapGetters(["getAssignments"])
  },
  methods: {
    openAssignment(assignment) {
      if (assignment.type === "lab") {
        this.$router.push(
          "/lab/" + assignment.subjectName + ":" + assignment.name
        );
      } else if (assignment.type === "homework") {
        this.$router.push(
          "/homework/" + assignment.subjectName + ":" + assignment.name
        );
      }
    }
  }
};
</script>

<style scoped>
table {
  color: orange;
  background-color: #777;
  margin-top: 100px;
  margin-left: 150px;
}

.table-row {
  cursor: pointer;
}
</style>
