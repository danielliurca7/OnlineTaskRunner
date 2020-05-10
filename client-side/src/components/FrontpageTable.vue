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
        <td>{{ assignment.course }}</td>
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
import { mapGetters, mapActions } from "vuex";

export default {
  name: "FrontpageTable",
  computed: {
    ...mapGetters(["getUsername", "getAssignments", "getTabs"])
  },
  methods: {
    ...mapActions(["openTab"]),
    openAssignment(assignment) {
      var tabs = JSON.parse(JSON.stringify(this.getTabs));

      // if tab is not already opened, open it
      if (tabs.filter(tab => tab.name === assignment.name).length === 0) {
        this.openTab(assignment);
      }

      this.$router.push(
        [
          "/workspace",
          assignment.course,
          assignment.series,
          assignment.year,
          assignment.name,
          this.getUsername
        ].join("/")
      );
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
