<template>
  <a-form-item prop="name" name="name" label="Tên phòng/ban">
    <a-input v-model:value="value.name" placeholder="Nhập tên phòng ban..." />
  </a-form-item>
</template>
<script>
export default {
  name: "DepartmentName",
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
