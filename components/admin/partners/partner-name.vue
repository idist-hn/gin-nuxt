<template>
  <a-form-item prop="name" name="name" label="Tên đơn vị">
    <a-input v-model:value="value.name" placeholder="Nhập tên đơn vị..." />
  </a-form-item>
</template>
<script>
export default {
  name: "PartnerName",
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
