<template>
  <a-page-header title="Cập nhật người dùng" @back="() => routerBack()"/>
  <a-form v-if="!loading" ref="FormData" layout="vertical" :rules="rules" :model="entry">
    <a-row :gutter="16">
      <a-col :xxl="18" :lg="18" :md="18" :sm="24">
        <a-card title="Thông tin tài khoản">
          <user-name :value="entry"/>
          <user-department :value="entry"/>
        </a-card>
        <a-card title="Thông tin cơ bản">
          <user-username :value="entry"/>
          <user-password :value="entry"/>
          <user-roles :value="entry"/>
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
            <a-popconfirm placem ent="top" ok-text="Đồng ý" cancel-text="Không" @confirm="routerBack">
              <template v-slot:title>
                <p>Huỷ tạo mới người dùng?</p>
                <p>Bạn xác nhận huỷ tạo người dùng này...</p>
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
import {useUsersStore} from "~/stores/users.js";
import UserName from "~/components/admin/users/user-name.vue";
import UserDepartment from "~/components/admin/users/user-department.vue";
import UserPassword from "~/components/admin/users/user-password.vue";
import UserUsername from "~/components/admin/users/user-username.vue";
import UserRoles from "~/components/admin/users/user-roles.vue";

export default {
  name: "create",
  components: {UserRoles, UserUsername, UserPassword, UserDepartment, UserName},
  setup() {
    definePageMeta({
      layout: "admin",
    });
  },
  data: () => ({
    loading: true,
    entry: {
      name: "",
      department: null,
      department_id: null,
      password: ""
    },
    rules: {
      name: [
        {required: true, message: "Vui lòng nhập tên người dùng", trigger: "blur"},
        {min: 3, message: "Tên người dùng phải có ít nhất 3 ký tự", trigger: "blur"},
        {max: 255, message: "Tên người dùng không được vượt quá 255 ký tự", trigger: "blur"},
      ],
    },
  }),
  methods: {
    routerBack() {
      this.$router.push({name: "admin-users"});
    },
    async getData(){
      this.loading = true
      let response = await useUsersStore().GetEntry({id: this.$route.params.id});
      if (response && response.code === 200) {
        this.entry = response.data.entry;
      }
      this.loading = false
    },
    async submitForm() {
      this.loading = true
      let response = await useUsersStore().UpdateEntry(this.entry);
      if (response && response.code === 200) {
        await this.$router.push({name: "admin-users"});
        await this.$toast("Cập nhật thành công", {autoClose: 5000, type: "success"});
      } else {
        await this.$toast("Cập nhật thất bại", {autoClose: 5000, type: "error"});
      }
      this.loading = false
    },
  },
  mounted() {
    this.getData()
  }
}
</script>
<style scoped>

</style>