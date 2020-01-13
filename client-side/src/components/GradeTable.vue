<template>
  <div class="container">
    <table
      class="table table-bordered table-condensed table-striped table-hover"
    >
      <tr>
        <th>Subject</th>
        <th>Assignment</th>
        <th>Student</th>
        <th>Grade</th>
        <th>Grade Time</th>
        <th>Graded by</th>
      </tr>
      <tr
        class="table-row"
        v-for="(assignment, index) in getAssignmentStudents(
          $route.params.subject,
          $route.params.assignment
        )"
        :key="index"
      >
        <td>{{ $route.params.subject }}</td>
        <td>{{ $route.params.assignment }}</td>
        <td>{{ assignment.name }}</td>
        <td>{{ assignment.grade }}</td>
        <td>{{ assignment.gradetime }}</td>
        <td>{{ assignment.graded_by }}</td>
        <td>
          <form class="form-inline">
            <div class="form-group">
              <input
                type="number"
                v-model="grade[index]"
                class="form-control"
                placeholder="Grade"
              />
            </div>
            <b-form @submit.prevent="onSubmit(assignment.name, index)">
              <button type="submit" class="btn btn-primary">Grade</button>
            </b-form>
          </form>
        </td>
      </tr>
    </table>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "GradeTable",
  methods: {
    ...mapActions(["gradeAssignment"]),
    onSubmit: function(student, index) {
      var grade = this.grade[index];
      var gradetime = new Date();
      var graded_by = this.getUsername;

      this.gradeAssignment({
        subject_name: this.$route.params.subject,
        assignment_name: this.$route.params.assignment,
        student_name: student,
        grade: grade,
        gradetime: gradetime,
        graded_by: graded_by
      });

      this.grade = [];
    }
  },
  data() {
    return {
      grade: []
    };
  },
  computed: {
    ...mapGetters(["getUsername", "getAssignmentStudents"])
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
