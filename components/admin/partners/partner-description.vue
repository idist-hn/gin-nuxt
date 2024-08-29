<template>
  <a-form-model-item prop="description" name="description" label="Mô tả đơn vị">
    <a-textarea :rows="3" v-model:value="value.description" placeholder="Nhập mô tả..." />
  </a-form-model-item>
</template>
<script>
export default {
  name: "PartnerDescription",
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
