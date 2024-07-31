<template>
  <div>
    <a-form-item prop="title" name="title" label="Tên hiển thị">
      <a-input v-model:value='entry.title' placeholder='Nhập tên...' />
    </a-form-item>
    <a-form-item prop="url" name="url" label="Đường dẫn">
      <a-input v-model:value='entry.url' placeholder='Nhập link...' />
    </a-form-item>
    <a-form-item prop="description" name="description" label="Mô tả">
      <a-input v-model:value='entry.description' placeholder='Nhập mô tả...' />
    </a-form-item>
    <a-form-item prop="openNewTab" name="openNewTab" label="">
      <a-switch :checked="entry.openNewTab" @change="entry.openNewTab = !entry.openNewTab" />
      Mở trong trang mới
    </a-form-item>
    <a-button v-if="!hideAction" type="primary" size="default" :outlined="true" @click="addEntry">Thêm</a-button>
  </div>
</template>

<script>
import {now} from "lodash";

export default {
  name: "DepartmentMenuItemCustom",
  props: {
    value: Object,
    hideAction: Boolean
  },
  data: () => ({
    entry: {
      id: "",
      title: "",
      description: "",
      url: "",
      openNewTab: true,
      type: "custom",
      children: []
    },
  }),
  methods: {
    addEntry() {
      this.entry.id = now();
      this.value.menu_items.push(this.entry);
      this.entry = {
        id: "",
        title: "",
        url: "",
        type: "custom",
        openNewTab: true,
        children: []
      };
    }
  },
  created() {

  },
  mounted() {
  },
};
</script>

<style scoped></style>
