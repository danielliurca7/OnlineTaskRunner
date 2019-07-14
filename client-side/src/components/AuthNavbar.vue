<template>
  <b-navbar type="dark" variant="dark">
    <b-navbar-nav>
      <b-nav-item class="ml-2" v-for="tab in tabs" :key="tab.name" :to="{ name: 'subject', params: { subject: tab.name }}">
        {{ tab.name }}
      </b-nav-item>
    </b-navbar-nav>

    <b-navbar-nav class="ml-auto active">
      <b-nav-item to="/dashboard">Dashboard</b-nav-item>
      <b-nav-item to="/profile">{{ username }}</b-nav-item>
    </b-navbar-nav>
  </b-navbar>
</template>

<script>
export default {
  name: "AuthNavbar",
  data() {
    return {
      username: "Gigel",
      tabs: [
        { name: "PC" },
        { name: "APD" },
        { name: "PA" },
        { name: "EGC" }
      ] 
    };
  }
};
</script>
