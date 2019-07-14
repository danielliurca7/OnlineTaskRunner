<template>
  <div class="dashboard">
    <AuthNavbar />
    <ul class="list-group mt-3 float-left w-50 ml-3">
      <li class="list-group-item m-1 btn btn-secondary" v-for="subject in subjects" :key="subject.name">
        <router-link :to="{ name: 'subject', params: { subject: subject.name }}">{{ subject.name }}</router-link>
      </li>
    </ul>
  </div>
</template>

<script>
import AuthNavbar from "@/components/AuthNavbar.vue";

export default {
  name: "dashboard",
  components: {
    AuthNavbar
  },
  data() {
    return {
      subjects: [
        { name: "PC" },
        { name: "APD" },
        { name: "PA" },
        { name: "EGC" }
      ]
    };
  }
};
</script>

<style scoped>
a {
  color: orange;
  display: block;
  text-decoration: none;
  cursor: pointer;
}
</style>
