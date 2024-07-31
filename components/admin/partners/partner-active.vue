<template>
  <a-form-item prop="is_active" name="is_active" label="">
    <a-switch :checked="value.is_active" @change="onChange" />
    Trạng thái hiển thị
  </a-form-item>
</template>

<script>
export default {
  name: "PartnerActive",
  data: () => ({}),
  props: {
    value: Object,
  },
  methods: {
    onChange(checked) {
      this.value.is_active = checked;
    },
  },
  created() {},
  mounted() {},
};
</script>

<style scoped></style>
