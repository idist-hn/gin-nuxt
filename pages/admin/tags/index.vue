<template>
  <a-page-header title="Từ khoá" @back="() => routerBack()">
    <template v-slot:extra>
      <nuxt-link :to="{ name: 'admin-tags-create' }">
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
        <template v-if="column.key === 'hash'">
          <a target="_blank" :href="publicLinkTag(record)">{{ record.hash }}</a>
        </template>
        <template v-if="column.key === 'count'">
          <a-tag color="green">{{ record.count_articles }} Bài viết</a-tag>
        </template>
        <template v-if="column.key === 'actions'">
          <a-space>
            <nuxt-link :to="{name: 'admin-tags-id-edit', params: { id: record.id } }">
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
import {useTagStore} from "~/stores/tags.js";
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
    columns:[
      {
        dataIndex: "name",
        key: "name",
        title: "Tên",
      },
      {
        title: "Hash",
        dataIndex: "hash",
        key: "hash",
      },
      {
        title: "Slug",
        dataIndex: "slug",
        key: "slug",
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
    publicLinkTag(entry) {
      return `${this.$config.public.portalDomain}/tu-khoa/${entry.hash}`;
    },
    async getData() {
      this.loading = true;
      let response = await useTagStore().ListEntries(this.pagination);
      this.entries = response.entries;
      this.pagination = response.pagination;
      this.loading = false;
    },
    async deleteData(entry) {
      this.loading = true;
      let response = await useTagStore().DeleteEntry(entry);
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
