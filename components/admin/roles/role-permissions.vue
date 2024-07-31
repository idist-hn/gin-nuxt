<template>
  <a-row :gutter="12">
    <a-col :md="12" :lg="8" v-for="permission in permissions" :key="permission.id">
      <a-checkbox @click="changeStatus(permission.id)" :checked="checked(permission.id)">{{ permission.name }}</a-checkbox>
    </a-col>
  </a-row>
</template>

<script>
import {usePermissionsStore} from "~/stores/permissions.js";
export default {
  name: "RolePermissions",
  props: {
    value: Object,
  },
  data: () => ({
    loading: false,
    permissions: [],
    pagination: {
      page: 1,
      pageSize: 1000,
    },
  }),
  methods: {
    async getData() {
      this.loading = true;
      let response = await usePermissionsStore().ListEntries(this.pagination);
      this.permissions = response.entries;
      this.loading = false;
    },
    checked(id) {
      return this.value.permission_ids.includes(id);
    },
    changeStatus(id) {
      if (this.value.permission_ids.includes(id)) {
        this.value.permission_ids = this.value.permission_ids.filter((item) => item !== id);
      } else {
        this.value.permission_ids.push(id);
      }
    },
  },
  mounted() {
    this.getData();
  },
};
</script>

<style scoped></style>
