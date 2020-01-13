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
            :to="{ name: 'subject', params: { name: tab.name } }"
          >
            <span> {{ tab.name }} </span>
            <b-button
              type="button"
              @click.prevent.stop="close(tab.name)"
              class="close ml-2"
              aria-label="Close"
            >
              <span aria-hidden="true">×</span>
            </b-button>
          </b-nav-item>
        </b-navbar-nav>
      </b-collapse>

      <b-navbar-nav class="ml-auto" v-if="getType == '1' || getType == '12'">
        <b-nav-item to="/dashboard">Dashboard</b-nav-item>
      </b-navbar-nav>

      <b-navbar-nav class="ml-auto" v-if="getType == '1' || getType == '12'">
        <b-nav-item to="/subjectoverview">Subjects</b-nav-item>
      </b-navbar-nav>

      <b-navbar-nav class="ml-auto" v-if="getType == '1' || getType == '12'">
        <b-nav-item to="/results">Results</b-nav-item>
      </b-navbar-nav>

      <b-navbar-nav
        class="ml-auto"
        v-if="getType == '2' || getType == '12' || getType == '23'"
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
              :to="{ name: 'subject', params: { name: tab.name } }"
            >
              <span> {{ tab.name }} </span>
              <b-button
                type="button"
                @click.prevent.stop="close(tab.name)"
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
  methods: {
    ...mapActions(["setUsername", "closeTab", "setType"]),
    close: function(tabName) {
      // close the tab with name tabName
      this.closeTab(tabName);

      var tabs = JSON.parse(JSON.stringify(this.getTabs));

      // if there are no more tabs left, go to dashboard
      // else if the current tab has the name tabName
      if (tabs.length === 0) {
        this.$router.push("/dashboard");
      } else if (this.$route.params.name === tabName) {
        this.$router.push("/subject/" + tabs.pop().name);
      }
    }
  },
  computed: {
    ...mapGetters(["getUsername", "getTabs", "getType"])
  },
  mounted() {
    // if the username cookie is active, load it
    if (
      this.$cookie.get("username") !== undefined &&
      this.$cookie.get("username") !== null
    ) {
      this.setUsername(this.$cookie.get("username"));
    }

    // if the type cookie is active, load it
    if (
      this.$cookie.get("type") !== undefined &&
      this.$cookie.get("type") !== null
    ) {
      this.setType(this.$cookie.get("type"));
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
