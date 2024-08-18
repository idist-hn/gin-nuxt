<template>
  <a-form-item label="Tài khoản" name="username">
    <a-input v-model:value="value.username"/>
  </a-form-item>
</template>

<script>
export default {
  name: "UserUsername",
  props: {
    value: Object,
  },
};
</script>

<style scoped></style>
