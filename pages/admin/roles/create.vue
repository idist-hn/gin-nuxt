<template>
  <a-page-header title="Tạo mới nhóm quyền" @back="() => routerBack()"/>
  <a-form ref="FormData" layout="vertical" :rules="rules" :model="entry">
    <a-row :gutter="16">
      <a-col :xxl="18" :lg="18" :md="18" :sm="24">
        <a-card>
          <role-name :value="entry"/>
          <role-description :value="entry"/>
        </a-card>
        <a-card>
          <role-department :value="entry"/>
          <role-permissions :value="entry"/>
        </a-card>
      </a-col>
      <a-col :xxl="6" :lg="6" :md="6" :sm="24">
        <a-card class="mb-20">
          <a-space :size="6">
            <a-button
                :loading="loading"
                :disabled="loading"
                type="primary"
                html-type="submit"
                @click.prevent="submitForm()"
            >
              Lưu
            </a-button>
            <a-popconfirm placement="top" ok-text="Đồng ý" cancel-text="Không" @confirm="routerBack">
              <template v-slot:title>
                <p>Huỷ tạo mới bài viết?</p>
                <p>Bạn xác nhận huỷ tạo bài viết này...</p>
              </template>
              <a-button type="white" size="default" :outlined="true">Huỷ</a-button>
            </a-popconfirm>
          </a-space>
        </a-card>
      </a-col>
    </a-row>
  </a-form>
</template>

<script>
import {useRolesStore} from "~/stores/roles.js";
import RoleName from "~/components/admin/roles/role-name.vue";
import RoleDescription from "~/components/admin/roles/role-description.vue";
import RolePermissions from "~/components/admin/roles/role-permissions.vue";
import RoleDepartment from "~/components/admin/roles/role-department.vue";

export default {
  name: "create",
  components: {RoleDepartment, RolePermissions, RoleDescription, RoleName},
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
      department: null,
      department_id: "",
      permissions: [],
      permission_ids: []
    },
    rules: {
      name: [
        {required: true, message: "Vui lòng nhập tên nhóm quyền", trigger: "blur"},
      ],
    },
  }),
  methods: {
    routerBack() {
      this.$router.push({name: "admin-roles"});
    },
    async submitForm() {
      this.loading = true
      let response = await useRolesStore().CreateEntry(this.entry);
      if (response && response.code === 200) {
        await this.$router.push({name: "admin-roles"});
        await this.$toast("Thêm mới thành công", {autoClose: 5000, type: "success"});
      } else {
        await this.$toast("Thêm mới thất bại", {autoClose: 5000, type: "error"});
      }
      this.loading = false
    },
  },
}
</script>
<style scoped>

</style>