<template>
  <div class="dashboard">
    <AuthNavbar />
    <ul class="list-group mt-3 float-left w-50 ml-3">
      <li
        class="list-group-item  m-1 btn"
        v-for="subject in getSubjects"
        :key="subject.name"
        @click="setTab(subject.name)"
      >
        {{ subject.name }}
      </li>
    </ul>
  </div>
</template>

<script>
import AuthNavbar from "@/components/AuthNavbar.vue";
import { mapGetters, mapActions } from "vuex";

export default {
  name: "dashboard",
  components: {
    AuthNavbar
  },
  computed: mapGetters(["getSubjects", "getTabs"]),
  methods: {
    ...mapActions(["openTab"]),
    setTab: function(name) {
      var tabs = JSON.parse(JSON.stringify(this.getTabs));

      // if tab is not already opened, open it
      if (tabs.filter(tab => tab.name === name).length === 0) {
        this.openTab(name);
      }

      // go to the clicked tab
      this.$router.push("/subject/" + name);
    }
  }
};
</script>

<style scoped>
li {
  color: orange;
  display: block;
  text-decoration: none;
  cursor: pointer;
}

li:hover {
  color: orange;
  background-color: #ddd;
}
</style>
