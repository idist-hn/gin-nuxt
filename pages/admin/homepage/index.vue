<template>
  <a-page-header title="Tổng quan"/>
  <a-row :gutter="16" justify="center">
    <a-col :md="12">
      <homepage-categories/>
      <!--      <overview-most-viewed/>-->
      <!--      <overview-social-views/>-->
    </a-col>

    <a-col :md="12" class="mb-20">
      <!--      <overview-performance/>-->
      <!--      <overview-post-performance/>-->
      <!--      <overview-prefered-performance/>-->
    </a-col>
  </a-row>
</template>
<script>
import HomepageCategories from "~/content/homepage/homepage-categories.vue";

export default {
  name: "Homepage",
  components: {HomepageCategories},
  setup() {
    definePageMeta({
      layout: "admin",
      // middleware: "auth",
    });
    // nextTick(() => {
    //
    // });
    // useNuxtApp().$toast.info('toastify success');
  },
};
</script>