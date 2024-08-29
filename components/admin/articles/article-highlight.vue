<template>
  <a-form-item prop="is_highlight" name="is_highlight" label="">
    <a-switch :checked="value.is_highlight" @change="onChange" />
    Bài viết nổi bật
  </a-form-item>
</template>

<script>
export default {
  name: "ArticleHighlight",
  data: () => ({}),
  props: {
    value: Object,
  },
  methods: {
    onChange(checked) {
      this.value.is_highlight = checked;
    },
  },
  created() {},
  mounted() {},
};
</script>

<style scoped></style>
