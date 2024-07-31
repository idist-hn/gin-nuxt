<template>
  <a-form-item prop="name" name="name" label="Tên menu">
    <a-input v-model:value="value.name" placeholder="Nhập tên menu..." />
  </a-form-item>
</template>

<script>
export default {
  name: "MenuName",
  props: {
    value: Object,
  },
};
</script>

<style scoped></style>
