<template>
  <a-form-item v-if="hideTitle" prop="value" name="value" label="Thuộc tính">
    <a-input v-model:value="value.value" placeholder="Nhập tên" />
  </a-form-item>
  <a-input v-else v-model:value="value.value" placeholder="Nhập tên" />
</template>
<script>
export default {
  name: "DepartmentStyleValue",
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
