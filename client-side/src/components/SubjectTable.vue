<template>
  <div class="container">
    <b-button-group class="m-4">
      <b-button @click="display = 1" variant="success">
        Labs
      </b-button>
      <b-button @click="display = 2" variant="secondary">
        Homeworks
      </b-button>
      <b-button @click="display = 3" variant="dark">
        Results
      </b-button>
    </b-button-group>

    <h1 class="m-4">
      {{ display === 1 ? "Laboratoare" : display === 2 ? "Teme" : "Rezultate" }}
    </h1>

    <table
      class="table table-bordered table-condensed table-striped table-hover"
      v-if="display === 1"
    >
      <tr>
        <th>Subject</th>
        <th>Assignment</th>
        <th>Total</th>
        <th>Deadline</th>
      </tr>
      <tr
        class="table-row"
        v-for="(lab, index) in getLabs($route.params.name)"
        :key="index"
        @click="openLab($route.params.name + ':' + lab.name)"
      >
        <td>{{ $route.params.name }}</td>
        <td>{{ lab.name }}</td>
        <td>{{ lab.total }}</td>
        <td>
          {{
            lab.deadline.getDate() +
              "-" +
              (lab.deadline.getMonth() + 1) +
              "-" +
              lab.deadline.getFullYear()
          }}
        </td>
      </tr>
    </table>

    <table
      class="table table-bordered table-condensed table-striped table-hover"
      v-if="display === 2"
    >
      <tr>
        <th>Subject</th>
        <th>Assignment</th>
        <th>Total</th>
        <th>Deadline</th>
      </tr>
      <tr
        class="table-row"
        v-for="(homework, index) in getHomeworks($route.params.name)"
        :key="index"
        @click="openHomework($route.params.name + ':' + homework.name)"
      >
        <td>{{ $route.params.name }}</td>
        <td>{{ homework.name }}</td>
        <td>{{ homework.total }}</td>
        <td>
          {{
            homework.deadline.getDate() +
              "-" +
              (homework.deadline.getMonth() + 1) +
              "-" +
              homework.deadline.getFullYear()
          }}
        </td>
      </tr>
    </table>

    <table
      class="table table-bordered table-condensed table-striped table-hover"
      v-if="display === 3"
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
        v-for="(result, index) in getResults($route.params.name)"
        :key="index"
      >
        <td>{{ $route.params.name }}</td>
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
  name: "SubjectTable",
  methods: {
    openLab(name) {
      this.$router.push("/lab/" + name);
    },
    openHomework(name) {
      this.$router.push("/homework/" + name);
    }
  },
  data() {
    return {
      display: 1
    };
  },
  computed: {
    ...mapGetters(["getLabs", "getHomeworks", "getResults"])
  },
  watch: {
    $route() {
      this.display = 1;
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
