<template>
  <a-form-item prop="is_hot" name="is_hot" label="">
    <a-switch :checked="value.is_hot" @change="onChange" />
    Bài viết tiêu biểu
  </a-form-item>
</template>

<script>
export default {
  name: "ArticleHot",
  data: () => ({}),
  props: {
    value: Object,
  },
  methods: {
    onChange(checked) {
      this.value.is_hot = checked;
    },
  },
  created() {},
  mounted() {},
};
</script>

<style scoped></style>
