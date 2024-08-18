<template>
  <a-form-item prop="name" name="name" label="Danh mục">
    <a-input v-model:value="value.name" placeholder="Nhập tên danh mục..." />
  </a-form-item>
</template>

<script>
export default {
  name: "CategoryName",
  props: {
    value: Object,
  },
};
</script>

<style scoped></style>
