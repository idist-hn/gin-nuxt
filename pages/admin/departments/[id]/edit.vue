<template>
  <div>
    <a-page-header title="Sửa thông tin phòng ban" @back="() => routerBack()"/>
    <a-form v-if="!loading" ref="FormData" layout="vertical" :rules="rules" :model="entry">
      <a-tabs type="card" class="config-tabs-bar">
        <a-tab-pane key="general" tab="Thông tin chung">
          <a-row :gutter="16">
            <a-col :md="12">
              <a-card>
                <a-row :gutter="16">
                  <a-col :xl="24" :sm="24">
                    <department-name :value="entry"/>
                  </a-col>
                  <a-col :xl="24" :sm="24">
                    <department-slug :value="entry"/>
                  </a-col>
                  <a-col :xl="24" :sm="24">
                    <department-address :value="entry"/>
                  </a-col>
                  <a-col :xl="24" :sm="24">
                    <department-phone :value="entry"/>
                  </a-col>
                  <a-col :xl="24" :sm="24">
                    <department-email :value="entry"/>
                  </a-col>
                  <a-col :xl="12" :sm="24">
                    <department-parent :value="entry"/>
                  </a-col>
                  <a-col :xl="12" :sm="24">
                    <department-description :value="entry"/>
                  </a-col>
                  <a-col :xl="12" :sm="24">
                    <department-icon :value="entry"/>
                  </a-col>
                  <a-col :xl="12" :sm="24">
                    <department-active :value="entry"/>
                    <department-sidebar :value="entry"/>
                  </a-col>
                </a-row>
                <div class="mt-10 mb-30">
                  <a-button :loading="loading" :disabled="loading" type="primary" html-type="submit"
                            @click="SubmitForm">
                    Lưu
                  </a-button>
                  <a-popconfirm placement="top" ok-text="Đồng ý" cancel-text="Không" @confirm="routerBack">
                    <template v-slot:title>
                      <p>Huỷ tạo mới phòng ban?</p>
                      <p>Bạn xác nhận huỷ tạo phòng ban này...</p>
                    </template>
                    <a-button type="white" size="default" :outlined="true">Huỷ</a-button>
                  </a-popconfirm>
                </div>
              </a-card>
            </a-col>
            <a-col :md="12">
              <department-pc-style :value="entry"/>
            </a-col>
            <a-col :md="12">
              <department-mobile-style :value="entry"/>
            </a-col>
          </a-row>
        </a-tab-pane>
        <a-tab-pane key="content" tab="Nội dung trên trang">
          <department-extra :value="entry"/>
          <department-content :value="entry"/>
          <div class="mt-10 mb-30">
            <a-button :loading="loading" :disabled="loading" type="primary" html-type="submit" @click="SubmitForm">
              Lưu
            </a-button>
            <a-popconfirm placement="top" ok-text="Đồng ý" cancel-text="Không" @confirm="routerBack">
              <template v-slot:title>
                <p>Huỷ tạo mới phòng ban?</p>
                <p>Bạn xác nhận huỷ tạo phòng ban này...</p>
              </template>
              <a-button type="white" size="default" :outlined="true">Huỷ</a-button>
            </a-popconfirm>
          </div>
        </a-tab-pane>
        <a-tab-pane key="menu" tab="Cấu hình menu">
          <a-row :gutter="16">
            <a-col :md="6">
              <a-collapse accordion class="mb-10">
                <a-collapse-panel key="custom" header="Tuỳ chỉnh">
                  <department-menu-item-custom :value="entry"/>
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
                <menu-tree-items v-if="entry.menu_items.length" :parent_id="-1" :recursive="true" title="title"
                                 :data="entry.menu_items"/>
                <p v-else class="text-center font-italic">Chưa có dữ liệu</p>
              </a-card>
            </a-col>
          </a-row>
        </a-tab-pane>
      </a-tabs>
    </a-form>
    <div v-else class="sd-spin">
      <a-spin/>
    </div>
  </div>
</template>

<script>

import {useDepartmentStore} from "~/stores/departments.js";
import DepartmentName from "~/components/admin/departments/department-name.vue";
import DepartmentSlug from "~/components/admin/departments/department-slug.vue";
import DepartmentParent from "~/components/admin/departments/department-parent.vue";
import DepartmentDescription from "~/components/admin/departments/department-description.vue";
import DepartmentIcon from "~/components/admin/departments/department-icon.vue";
import DepartmentActive from "~/components/admin/departments/department-active.vue";
import DepartmentSidebar from "~/components/admin/departments/department-sidebar.vue";
import DepartmentPcStyle from "~/components/admin/departments/department-pc-style.vue";
import DepartmentMobileStyle from "~/components/admin/departments/department-mobile-style.vue";
import DepartmentExtra from "~/components/admin/departments/department-extra.vue";
import DepartmentContent from "~/components/admin/departments/department-content.vue";
import DepartmentAddress from "~/components/admin/departments/department-address.vue";
import DepartmentPhone from "~/components/admin/departments/department-phone.vue";
import DepartmentEmail from "~/components/admin/departments/department-email.vue";
import MenuName from "~/components/admin/menus/menu-name.vue";
import MenuItemCustom from "~/components/admin/menus/menu-item-custom.vue";
import MenuTreeItems from "~/components/admin/menus/menu-tree-items.vue";
import DepartmentMenuItemCustom from "~/components/admin/departments/department-menu-item-custom.vue";

export default {
  name: "edit",
  components: {
    DepartmentMenuItemCustom,
    MenuTreeItems, MenuItemCustom, MenuName,
    DepartmentEmail,
    DepartmentPhone,
    DepartmentAddress,
    DepartmentContent,
    DepartmentExtra,
    DepartmentMobileStyle,
    DepartmentPcStyle,
    DepartmentSidebar,
    DepartmentActive, DepartmentIcon, DepartmentDescription, DepartmentParent, DepartmentSlug, DepartmentName},
  setup() {
    definePageMeta({
      layout: "admin",
    });
  },
  data: () => ({
    rules: {
      name: [
        {
          required: true,
          message: "Không được bỏ trống",
          trigger: "blur",
        },
      ],
    },
    loading: true,
    entry: {
      name: "",
      description: "",
      is_active: true,
      slug: "",
      icon: "",
      parent: null,
      parent_id: null,
      content: "",
      extras: [],
      menu_items: []
    },
  }),
  methods: {
    routerBack() {
      this.$router.push({name: "admin-departments"});
    },
    async getData() {
      this.loading = true;
      let response = await useDepartmentStore().GetEntry({id: this.$route.params.id});
      if (response.code === 200) {
        this.entry = response.data.entry;
        if (this.entry.menu_items == null) {
          this.entry.menu_items = [];
        }
      }
      this.loading = false;
    },
    async SubmitForm() {
      this.loading = true;
      let response = await useDepartmentStore().UpdateEntry(this.entry);
      if (response && response.code === 200) {
        this.$toast(response.message, {autoClose: 2000, type: "success"});
      }
      this.loading = false;
      this.routerBack();
    },
  },

  mounted() {
    this.getData();
  },
};
</script>

<style scoped></style>
