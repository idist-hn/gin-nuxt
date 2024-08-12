<template>
  <a-page-header title="Người dùng" @back="() => routerBack()">
    <template v-slot:extra>
      <nuxt-link :to="{ name: 'admin-users-create' }">
        <a-button type="success">
          <idist-feather-icons type="plus"/>
          Tạo mới
        </a-button>
      </nuxt-link>
    </template>
  </a-page-header>
  <a-card>
    <a-table
        :loading="loading"
        :columns="columns"
        :data-source="entries"
        :pagination="pagination"
        @change="handleTableChange"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'name'">
          <div class="mb-10">
            <a-tooltip placement="bottom">
              <template #title>
                <span>{{ record.username }}</span>
              </template>
              <span>{{ record.name }}</span>
            </a-tooltip>
          </div>
        </template>
        <template v-if="column.key === 'department'">
          {{ record.department ? record.department.name : ''}}
        </template>
        <template v-if="column.key === 'roles'">
          <div>
            <a-tag v-for="role in record.roles" :key="role.id" color="gray">{{ role.name }}</a-tag>
          </div>
        </template>
        <template v-if="column.key === 'history'">
          <a-tag color="green">{{ moment(record.last_active_at, "HH:mm DD/MM/YY") }}</a-tag>
        </template>
        <template v-if="column.key === 'count'">
          <a-tag color="blue">{{ record.count_articles ?? 0 }} bài viết</a-tag>
        </template>
        <template v-if="column.key === 'actions'">
          <a-space>
            <nuxt-link :to="{name: 'admin-users-id-edit', params: { id: record.id } }">
              <a-button type="outline-primary" size="small">
                <idist-feather-icons type="edit"/>
              </a-button>
            </nuxt-link>
            <a-popconfirm
                placement="top"
                title="Xoá tài khoản này?"
                ok-text="Đồng ý"
                cancel-text="Không"
                @confirm="deleteData(record)"
            >
              <a-button type="outline-danger" size="small">
                <idist-feather-icons type="trash"/>
              </a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </template>
    </a-table>
  </a-card>
</template>

<script>
import list from "~/mixins/list";
import IdistFeatherIcons from "~/components/commons/IdistFeatherIcons.vue";
import {useUsersStore} from "~/stores/users.js";

export default {
  name: "list",
  components: {IdistFeatherIcons},
  setup() {
    definePageMeta({
      layout: "admin",
    });
  },
  mixins: [list],
  data: () => ({
    loading: true,
    entries: [],
    columns: [
      {
        dataIndex: "name",
        key: "name",
        title: "Tên",
      },
      {
        title: "Bộ phận",
        dataIndex: "department",
        key: "department",
      },
      {
        title: "Quyền/Nhóm quyền",
        dataIndex: "roles",
        key: "roles",
      },
      {
        title: "Hoạt động",
        dataIndex: "history",
        key: "history",
      },
      {
        title: "Thống kê",
        key: "count",
        dataIndex: "count",
      },
      {
        title: "",
        key: "actions",
        dataIndex: "actions",
        align: "right",
      },
    ]
  }),
  methods: {
    routerBack() {
      this.$router.push({name: "admin-homepage"});
    },
    async getData() {
      this.loading = true;
      let response = await useUsersStore().ListEntries(this.pagination);
      this.entries = response.entries;
      this.pagination = response.pagination;
      this.loading = false;
    },
    async deleteData(entry) {
      this.loading = true;
      let response = await useUsersStore().DeleteEntry(entry);
      if (response && response.code === 200) {
        await this.getData();
      }
      this.$toast(response.message, {
        autoClose: 2000,
        type: response && response.code === 200 ? "success" : "danger",
      });
      this.loading = false;
    },
  },
  mounted() {
  },
};
</script>

<style scoped></style>
