<template>
  <a-form-item prop="title" name="title" label="Tiêu đề">
    <a-input v-model:value="entry.title" placeholder="Nhập tiêu đề bài viết..." />
  </a-form-item>
</template>

<script>
export default {
  name: "PageTitle",
  props: {
    value: Object,
  },
  data: () => ({
    entry: {
      title: "",
    },
  }),
  methods: {},
  created() {
    this.entry = this.value;
  },
  mounted() {},
};
</script>

<style scoped></style>
