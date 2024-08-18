<template>
  <a-form-item prop="url" name="url" label="Trang giới thiệu">
    <a-input v-model:value="value.url" placeholder="Nhập link giới thiệu..." />
  </a-form-item>
</template>
<script>
export default {
  name: "PartnerUrl",
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
