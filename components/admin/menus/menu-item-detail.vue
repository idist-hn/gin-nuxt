<template>
  <div>
    <a-form-item prop="title" name="title" label="Tên hiển thị">
      <a-input v-model:value='value.title' placeholder='Nhập tên...' />
    </a-form-item>
    <a-form-item prop="url" name="url" label="Đường dẫn">
      <a-input v-model:value='value.url' placeholder='Nhập link...' />
    </a-form-item>
    <a-form-item prop="description" name="description" label="Mô tả">
      <a-input v-model:value='value.description' placeholder='Nhập mô tả...' />
    </a-form-item>
    <a-form-item prop="openNewTab" name="openNewTab" label="">
      <a-switch :checked="value.openNewTab" @change="onChange"/>
      Mở trong trang mới
    </a-form-item>
  </div>
</template>

<script>

export default {
  name: "MenuItemDetail",
  props: {
    value: Object,
  },
  data: () => ({
    entry: {
      id: "",
      title: "",
      description: "",
      url: "",
      openNewTab: true,
      type: "custom"
    },
  }),
  methods: {
    onChange(checked) {
      this.value.openNewTab = checked;
    },
  }
};
</script>

<style scoped></style>
