<template>
  <a-form-item prop="description" name="description" label="Mô tả phòng/ban">
    <a-textarea :rows="3" v-model:value="value.description" placeholder="Nhập mô tả phòng ban..." />
  </a-form-item>
</template>
<script>
export default {
  name: "DepartmentDescription",
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
