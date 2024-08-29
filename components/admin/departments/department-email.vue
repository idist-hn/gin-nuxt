<template>
  <a-form-item prop="email" name="email" label="Email liên hệ">
    <a-input v-model:value="value.email" placeholder="Nhập email..." />
  </a-form-item>
</template>
<script>
export default {
  name: "DepartmentEmail",
  props: {
    value: {
      type: Object,
      default: () => ({}),
    },
  },
  methods: {},
};
</script>
<style scoped></style>
