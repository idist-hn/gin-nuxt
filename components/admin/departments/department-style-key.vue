<template>
  <a-form-item v-if="hideTitle" prop="key" name="key" label="Thuộc tính">
    <a-input v-model:value="value.key" placeholder="Nhập tên" />
  </a-form-item>
  <a-input v-else v-model:value="value.key" placeholder="Nhập tên" />
</template>
<script>
export default {
  name: "DepartmentStyleKey",
  props: {
    value: {
      type: Object,
      default: () => ({}),
    },
    hideTitle: Boolean,
  },
  methods: {},
};
</script>
<style scoped></style>
