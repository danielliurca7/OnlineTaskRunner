<template>
  <div class="authenticate-navbar">
    <b-navbar toggleable="sm" type="dark" variant="dark">
      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse v-model="toggled" v-if="!toggled" id="nav-collapse" is-nav>
        <b-navbar-nav>
          <b-nav-item
            class="ml-2"
            v-for="tab in getTabs"
            :key="tab.name"
            @click="open(tab)"
          >
            <span> {{ tab.abbreviation }} {{ tab.name }} </span>
            <b-button
              type="button"
              @click.prevent.stop="close(tab)"
              class="close ml-2"
              aria-label="Close"
            >
              <span aria-hidden="true">×</span>
            </b-button>
          </b-nav-item>
        </b-navbar-nav>
      </b-collapse>

      <b-navbar-nav class="ml-auto" v-if="getUserinfo.student.courses">
        <b-nav-item to="/dashboard">Dashboard</b-nav-item>
      </b-navbar-nav>

      <b-navbar-nav class="ml-auto" v-if="getUserinfo.student.courses">
        <b-nav-item to="/subjectoverview">Subjects</b-nav-item>
      </b-navbar-nav>

      <b-navbar-nav class="ml-auto" v-if="getUserinfo.student.courses">
        <b-nav-item to="/results">Results</b-nav-item>
      </b-navbar-nav>

      <b-navbar-nav
        class="ml-auto"
        v-if="getUserinfo.assistant.courses || getUserinfo.professor.courses"
      >
        <b-nav-item to="/gradeoverview">Grade</b-nav-item>
      </b-navbar-nav>

      <b-navbar-nav class="ml-auto">
        <b-nav-item to="/profile">{{ getUsername }}</b-nav-item>
      </b-navbar-nav>

      <div v-if="toggled">
        <b-collapse v-model="toggled" id="nav-collapse" is-nav>
          <b-navbar-nav>
            <b-nav-item
              class="ml-2"
              v-for="tab in getTabs"
              :key="tab.name"
              @click="open(tab)"
            >
              <span> {{ tab.abbreviation }} {{ tab.name }} </span>
              <b-button
                type="button"
                @click.prevent.stop="close(tab)"
                class="close ml-2"
                aria-label="Close"
              >
                <span aria-hidden="true">×</span>
              </b-button>
            </b-nav-item>
          </b-navbar-nav>
        </b-collapse>
      </div>
    </b-navbar>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "AuthNavbar",
  data() {
    return {
      toggled: false
    };
  },
  computed: {
    ...mapGetters(["getUsername", "getUserinfo", "getTabs"])
  },
  methods: {
    ...mapActions(["closeTab", "setUserinfo", "setUsername", "setToken"]),
    open: function(tab) {
      if (
        this.$route.params.course === tab.course &&
        this.$route.params.series === tab.series &&
        this.$route.params.year === tab.year &&
        this.$route.params.assignmentname === tab.name
      ) {
        this.$router.push(
          [
            "/workspace",
            tab.course,
            tab.series,
            tab.year,
            tab.name,
            this.getUsername
          ].join("/")
        );
      }
    },
    close: function(tab) {
      // close the tab with name tabName
      this.closeTab(tab);

      var tabs = JSON.parse(JSON.stringify(this.getTabs));

      if (tabs.length === 0 && this.$route.fullPath != "/dashboard") {
        this.$router.push("/dashboard");
      } else if (
        this.$route.params.course === tab.course &&
        this.$route.params.series === tab.series &&
        parseInt(this.$route.params.year) === tab.year &&
        this.$route.params.assignmentname === tab.name
      ) {
        var next_tab = tabs.pop();

        this.$router.push(
          [
            "/workspace",
            next_tab.course,
            next_tab.series,
            next_tab.year,
            next_tab.name,
            this.getUsername
          ].join("/")
        );
      }
    }
  },
  created() {
    if (
      localStorage.getItem("username") &&
      localStorage.getItem("userinfo") &&
      localStorage.getItem("token")
    ) {
      this.setUsername(localStorage.getItem("username"));
      this.setUserinfo(JSON.parse(localStorage.getItem("userinfo")));
      this.setToken(localStorage.getItem("token"));
    }
  }
};
</script>

<style scoped>
button {
  color: orangered;
}

button:hover {
  color: red;
}
</style>
