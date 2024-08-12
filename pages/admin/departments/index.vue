<template>
  <a-page-header title="Phòng ban" @back="() => routerBack()">
    <template v-slot:extra>
      <nuxt-link :to="{ name: 'admin-departments-create' }">
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
          <span>{{ record.name }}</span>
        </template>
        <template v-if="column.key === 'is_active'">
          <a-switch :checked="record.is_active" disabled/>
        </template>
        <template v-if="column.key === 'show_sidebar'">
          <a-switch :checked="record.show_sidebar" disabled/>
        </template>
        <template v-if="column.key === 'count'">
          <a-tag color="green">{{ record.count_users }} người dùng</a-tag>
        </template>
        <template v-if="column.key === 'actions'">
          <a-space>
            <nuxt-link :to="{name: 'admin-departments-id-edit', params: { id: record.id } }">
              <a-button type="outline-primary" size="small">
                <idist-feather-icons type="edit"/>
              </a-button>
            </nuxt-link>
            <a-popconfirm
                placement="top"
                title="Xoá từ khoá này?"
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
import {useDepartmentStore} from "~/stores/departments.js";
import IdistFeatherIcons from "~/components/commons/IdistFeatherIcons.vue";

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
        title: "Đơn vị quản lý",
        dataIndex: "parent",
        key: "parent",
      },
      {
        title: "Hoạt động",
        dataIndex: "is_active",
        key: "is_active",
      },
      {
        title: "Hiển thị danh mục",
        dataIndex: "show_sidebar",
        key: "show_sidebar",
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
      let response = await useDepartmentStore().ListEntries(this.pagination);
      this.entries = response.entries;
      this.pagination = response.pagination;
      this.loading = false;
    },
    async deleteData(entry) {
      this.loading = true;
      let response = await useDepartmentStore().DeleteEntry(entry);
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
