<template>
  <a-page-header title="Tạo mới menu" @back="() => routerBack()"/>
  <a-form ref="FormData" layout="vertical" :rules="rules" :model="entry">
    <a-row :gutter="16">
      <a-col :md="6">
        <a-card class="mb-10">
          <menu-name :value="entry"/>
          <menu-location :value="entry"/>
        </a-card>
        <a-collapse accordion class="mb-10">
          <a-collapse-panel key="page" header="Trang bài">
            <p>{{ text }}</p>
          </a-collapse-panel>
          <a-collapse-panel key="post" header="Bài viết">
            <p>{{ text }}</p>
          </a-collapse-panel>
          <a-collapse-panel key="category" header="Danh mục">
            <p>{{ text }}</p>
          </a-collapse-panel>
          <a-collapse-panel key="custom" header="Tuỳ chỉnh">
            <menu-item-custom :value="entry"/>
          </a-collapse-panel>
        </a-collapse>
        <a-card>
          <a-space :size="10">
            <a-button type="success" @click="SubmitForm">
              Lưu
            </a-button>
            <a-button type="danger" @click="routerBack">
              Huỷ bỏ
            </a-button>
          </a-space>
        </a-card>
      </a-col>
      <a-col :md="18">
        <a-card class="mb-10" title="Cấu trúc menu">
          <template v-slot:extra>
            <a-button type="success" @click="SubmitForm" :disabled="loading">
              Cập nhật
            </a-button>
          </template>
          <menu-tree-items v-if="entry.items.length" :parent_id="-1" :recursive="true" title="title"
                           :data="entry.items"/>
          <p v-else class="text-center font-italic">Chưa có dữ liệu</p>
        </a-card>
      </a-col>
    </a-row>
  </a-form>
</template>

<script>
import {useMenuStore} from "~/stores/menus.js";
import MenuName from "~/components/admin/menus/menu-name.vue";
import MenuLocation from "~/components/admin/menus/menu-location.vue";
import MenuItemCustom from "~/components/admin/menus/menu-item-custom.vue";
import MenuItemPage from "~/components/admin/menus/menu-item-page.vue";
import MenuTreeItems from "~/components/admin/menus/menu-tree-items.vue";

export default {
  name: "create",
  components: {MenuTreeItems, MenuItemPage, MenuItemCustom, MenuLocation, MenuName},
  setup() {
    definePageMeta({
      layout: "admin",
    });
  },
  data: () => ({
    loading: true,
    text: "Chưa có dữ liệu",
    entry: {
      name: "",
      location: null,
      items: [],
    },
    rules: {
      name: [
        {required: true, message: "Vui lòng nhập tên menu", trigger: "blur"},
        {min: 3, message: "Tên menu phải có ít nhất 3 ký tự", trigger: "blur"},
        {max: 255, message: "Tên menu không được vượt quá 255 ký tự", trigger: "blur"},
      ],
    },
  }),
  methods: {
    routerBack() {
      this.$router.push({name: "admin-menus"});
    },
    async SubmitForm() {
      this.loading = true
      let response = await useMenuStore().CreateEntry(this.entry);
      if (response && response.code === 200) {
        await this.$router.push({name: "admin-menus"});
        await this.$toast("Thêm mới thành công", {autoClose: 5000, type: "success"});
      } else {
        await this.$toast("Thêm mới thất bại", {autoClose: 5000, type: "error"});
      }
      this.loading = false
    },
  }
}
</script>

<style scoped lang="scss">

</style>