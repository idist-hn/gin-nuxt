<template>
  <a-form-item label="Mô tả nhóm quyền" name="description">
    <a-textarea :rows="3" v-model:value="value.description"/>
  </a-form-item>
</template>

<script>
export default {
  name: "RoleDescription",
  props: {
    value: Object,
  },
};
</script>

<style scoped></style>
