<template>
  <a-page-header title="Cập nhật từ khoá" @back="() => routerBack()"/>
  <a-card class="d-flex " style="flex-direction: column"
          body-style="text-align: center">
    <a-spin tip="Đang phát triển tính năng..."/>
  </a-card>
</template>

<script>
export default {
  name: "edit",
  setup() {
    definePageMeta({
      layout: "admin",
    });
  },
  methods: {
    routerBack() {
      this.$router.push({name: "admin-tags"});
    },
  },
}
</script>

<style scoped lang="scss">

</style>