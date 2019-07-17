<template>
  <div class="subject">
    <AuthNavbar />

    <b-button-group class="m-3">
      <b-button @click="labs = !labs" variant="success">Labs</b-button>
      <b-button @click="labs = !labs" variant="primary">Homeworks</b-button>
    </b-button-group>

    <h1 class="m-2">{{ labs ? "Laboratoare" : "Teme" }}</h1>

    <ul v-if="labs" class="list-group mt-3 float-left w-50 ml-3">
      <li
        class="list-group-item m-1"
        v-for="lab in getLabs($route.params.subject)"
        :key="lab.name"
      >
        {{ lab.name }}
      </li>
    </ul>
    <ul v-if="!labs" class="list-group mt-3 float-left w-50 ml-3">
      <li
        class="list-group-item m-1"
        v-for="homework in getHomeworks($route.params.subject)"
        :key="homework.name"
      >
        {{ homework.name }}
      </li>
    </ul>
  </div>
</template>

<script>
import AuthNavbar from "@/components/AuthNavbar.vue";
import { mapGetters } from "vuex";

export default {
  name: "subject",
  components: {
    AuthNavbar
  },
  data() {
    return {
      labs: true
    };
  },
  computed: {
    ...mapGetters(["getLabs", "getHomeworks"])
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
