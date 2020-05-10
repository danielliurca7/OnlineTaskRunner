<template>
  <div class="results">
    <AuthNavbar />
    <ResultsTable />
  </div>
</template>

<script>
import AuthNavbar from "@/components/AuthNavbar.vue";
import ResultsTable from "@/components/ResultsTable.vue";

export default {
  name: "Results",
  components: {
    AuthNavbar,
    ResultsTable
  }
};
</script>
