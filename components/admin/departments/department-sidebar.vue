<template>
  <a-form-item prop="show_sidebar" name="show_sidebar" label="">
    <a-switch :checked="value.show_sidebar" @change="onChange" />
    Hiển thị trên menu
  </a-form-item>
</template>

<script>
export default {
  name: "DepartmentSidebar",
  data: () => ({}),
  props: {
    value: Object,
  },
  methods: {
    onChange(checked) {
      this.value.show_sidebar = checked;
    },
  },
  created() {},
  mounted() {},
};
</script>

<style scoped></style>
