<template>
  <b-navbar type="dark" variant="dark">
    <b-navbar-nav>
      <b-nav-item class="ml-2" v-for="tab in getTabs" :key="tab.name" :to="{ name: 'subject', params: { subject: tab.name }}">
        {{ tab.name }}
      </b-nav-item>
    </b-navbar-nav>

    <b-navbar-nav class="ml-auto active">
      <b-nav-item to="/dashboard">Dashboard</b-nav-item>
      <b-nav-item to="/profile">{{ getUsername }}</b-nav-item>
    </b-navbar-nav>
  </b-navbar>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "AuthNavbar",
  computed: {
    ...mapGetters(["getUsername", "getTabs"])
  }
};
</script>
