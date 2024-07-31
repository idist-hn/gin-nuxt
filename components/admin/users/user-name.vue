<template>
  <a-form-item label="Tên tài khoản" name="name">
    <a-input v-model:value="value.name"/>
  </a-form-item>
</template>

<script>
export default {
  name: "UserName",
  props: {
    value: Object,
  },
};
</script>

<style scoped></style>
