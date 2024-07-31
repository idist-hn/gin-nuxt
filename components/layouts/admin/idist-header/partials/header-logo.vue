<template>
  <nuxt-link class="top-menu" :to="{ name: 'admin-homepage' }" style="margin-left: 1rem">
    <img src="@/assets/images/ptit-logo.png" alt="logo" height="40" />
    <span class="color-primary" style="font-weight: 600"> PTIT CMS </span>
  </nuxt-link>
</template>

<script>
export default {
  name: "header-logo",
  data: () => ({}),
  methods: {},
  created() {},
  mounted() {},
};
</script>

<style scoped></style>
