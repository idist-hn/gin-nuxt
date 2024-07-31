<template>
  <a-form-item prop="is_active" name="is_active" label="Hiển thị">
    <a-switch :checked="value.is_active" @change="value.is_active = !value.is_active"/>
  </a-form-item>
</template>

<script>
export default {
  name: "CategoryActive",
  data: () => ({}),
  props: {
    value: Object,
  },
};
</script>

<style scoped></style>
