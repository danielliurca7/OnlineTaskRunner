<template>
  <div class="resultsTable">
    <table
      class="table table-bordered table-condensed table-striped table-hover"
    >
      <tr>
        <th>Subject</th>
        <th>Assignment</th>
        <th>Result</th>
        <th>Points</th>
        <th>Total</th>
        <th>Deadline</th>
      </tr>
      <tr
        class="table-row"
        v-for="(result, index) in getAllResults"
        :key="index"
      >
        <td>{{ result.course }}</td>
        <td>{{ result.name }}</td>
        <td>{{ result.result }}</td>
        <td>{{ (result.result * result.total) / 100 }}</td>
        <td>{{ result.total }}</td>
        <td>
          {{
            result.deadline.getDate() +
              "-" +
              (result.deadline.getMonth() + 1) +
              "-" +
              result.deadline.getFullYear()
          }}
        </td>
      </tr>
    </table>
  </div>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "ResultsTable",
  computed: {
    ...mapGetters(["getAllResults"])
  }
};
</script>

<style scoped>
div {
  margin-top: 100px;
  margin-left: 50px;
  margin-right: 50px;
}

table {
  color: orange;
  background-color: #777;
}

.table-row {
  cursor: pointer;
}
</style>
