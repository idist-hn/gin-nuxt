<template>
  <a-form-item prop="address" name="address" label="Địa chỉ">
    <a-input v-model:value="value.address" placeholder="Nhập địa chỉ..." />
  </a-form-item>
</template>
<script>
export default {
  name: "DepartmentAddress",
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
