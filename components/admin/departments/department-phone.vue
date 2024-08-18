<template>
  <a-form-item prop="phone" name="phone" label="Điện thoại liên hệ">
    <a-input v-model:value="value.phone" placeholder="Nhập số điện thoại..." />
  </a-form-item>
</template>
<script>
export default {
  name: "DepartmentPhone",
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
