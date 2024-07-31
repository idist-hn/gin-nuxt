<template>
  <a-form-item label="Tên nhóm quyền" name="name">
    <a-input v-model:value="value.name"/>
  </a-form-item>
</template>

<script>
export default {
  name: "RoleName",
  props: {
    value: Object,
  },
};
</script>

<style scoped></style>
