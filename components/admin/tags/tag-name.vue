<template>
  <a-form-item prop="name" name="name" label="Từ khoá">
    <a-input v-model:value="value.name" placeholder="Nhập từ khoá..." />
  </a-form-item>
</template>

<script>
export default {
  name: "TagName",
  props: {
    value: Object,
  },
};
</script>

<style scoped></style>
