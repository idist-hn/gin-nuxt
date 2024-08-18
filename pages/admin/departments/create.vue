<template>
  <a-page-header title="Thêm phòng ban" @back="() => routerBack()"/>
  <a-card>
    <a-form ref="FormData" layout="vertical" :rules="rules" :model="entry">
      <a-row :gutter="16">
        <a-col :xl="12" :sm="24">
          <department-name :value="entry"/>
        </a-col>
        <a-col :xl="12" :sm="24">
          <department-parent :value="entry"/>
        </a-col>
        <a-col :xl="12" :sm="24">
          <department-description :value="entry"/>
        </a-col>
      </a-row>
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
    </a-form>
  </a-card>
</template>

<script>

import DepartmentName from "~/components/admin/departments/department-name.vue";
import DepartmentParent from "~/components/admin/departments/department-parent.vue";
import DepartmentDescription from "~/components/admin/departments/department-description.vue";
import {useDepartmentStore} from "~/stores/departments.js";

export default {
  name: "create",
  components: {DepartmentDescription, DepartmentParent, DepartmentName},
  layout: "admin",
  setup() {
    definePageMeta({
      layout: "admin",
    });
  },
  data: () => ({
    loading: false,
    entry: {
      name: "",
      description: "",
      is_active: true,
      slug: "",
      parent: null,
      parent_id: null,
    },
    rules: {
      name: [
        {
          required: true,
          message: "Không được bỏ trống",
          trigger: "blur",
        },
      ],
    },
  }),
  methods: {
    routerBack() {
      this.$router.push({name: "admin-departments"});
    },
    async SubmitForm() {
      this.loading = true;
      this.loading = true;
      let response = await useDepartmentStore().CreateEntry(this.entry);
      if (response && response.code === 200) {
        this.$toast(response.message, {autoClose: 2000, type: "success"});
        this.routerBack();
      }
      this.loading = false;
    },
  },
  created() {
  },
  mounted() {
  },
};
</script>

<style scoped></style>
