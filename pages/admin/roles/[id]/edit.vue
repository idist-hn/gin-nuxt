<template>
  <a-page-header title="Cập nhật trang giới thiệu" @back="() => routerBack()"/>
  <a-form v-if="!loading" ref="FormData" layout="vertical" :rules="rules" :model="entry">
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
  <a-card v-else class="d-flex justify-content-center">
    <a-spin/>
  </a-card>
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
    loading: true,
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

    async getData(){
      this.loading = true
      let response = await useRolesStore().GetEntry({id: this.$route.params.id})
      this.entry = response.data.entry
      if (this.entry.permissions == null) {
        this.entry.permissions = []
        this.entry.permission_ids = []
      }
      this.loading = false
    },
    async submitForm() {
      this.loading = true
      let response = await useRolesStore().UpdateEntry(this.entry);
      if (response && response.code === 200) {
        await this.$router.push({name: "admin-roles"});
        await this.$toast("Thêm mới thành công", {autoClose: 5000, type: "success"});
      } else {
        await this.$toast("Thêm mới thất bại", {autoClose: 5000, type: "error"});
      }
      this.loading = false
    },
  },
  mounted() {
    this.getData()
  }
}
</script>

<style scoped lang="scss">

</style>