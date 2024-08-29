<template>
  <a-page-header title="Menu" @back="() => routerBack()">
    <template v-slot:extra>
      <nuxt-link :to="{ name: 'admin-menus-create' }">
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
                <span>{{ record.hash }}</span>
              </template>
              <span>{{ record.name }}</span>
            </a-tooltip>
          </div>
        </template>
        <template v-if="column.key === 'is_used'">
          <a-switch :checked="record.is_used" disabled/>
        </template>
        <template v-if="column.key === 'location'">
          {{ record.location ? record.location.name : '-Chưa cấu hình-' }}
        </template>
        <template v-if="column.key === 'actions'">
          <a-space>
            <nuxt-link :to="{name: 'admin-menus-id-edit', params: { id: record.id } }">
              <a-button type="outline-primary" size="small">
                <idist-feather-icons type="edit"/>
              </a-button>
            </nuxt-link>
            <a-popconfirm
                placement="top"
                title="Xoá danh mục menu này?"
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
import IdistFeatherIcons from "~/components/commons/IdistFeatherIcons.vue";
import {useMenuStore} from "~/stores/menus.js";
import list from "~/mixins/list";
import {useTagStore} from "~/stores/tags.js";

export default {
  name: "index",
  mixins: [list],
  components: {IdistFeatherIcons},
  setup() {
    definePageMeta({
      layout: "admin",
    });
  },
  data: () => ({
    loading: false,
    entries: [],
    columns: [
      {
        title: "Tên",
        dataIndex: "name",
        key: "name",
      },
      {
        title: "Sử dụng",
        dataIndex: "is_used",
        key: "is_used",
      },
      {
        title: "Vị trí trên trang",
        dataIndex: "location",
        key: "location",
      },
      {
        title: "Hành động",
        dataIndex: "actions",
        key: "actions",
        align: "right",
      },
    ],
  }),
  methods: {
    routerBack() {
      this.$router.push({name: "admin-homepage"});
    },
    async getData() {
      this.loading = true;
      let response = await useMenuStore().ListEntries(this.pagination)
      this.entries = response.entries;
      this.pagination = response.pagination;
      this.loading = false;
    },

    async deleteData(entry) {
      this.loading = true;
      let response = await useMenuStore().DeleteEntry(entry);
      if (response && response.code === 200) {
        await this.getData();
      }

      this.$toast(response.message, {
        autoClose: 2000,
        type: response && response.code === 200 ? "success" : "danger",
      });
      this.loading = false;
    },
  }
}
</script>

<style scoped lang="scss">

</style>