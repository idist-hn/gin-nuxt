<template>
  <a-page-header title="Trang giới thiệu" @back="() => routerBack()">
    <template v-slot:extra>
      <nuxt-link :to="{ name: 'admin-pages-create' }">
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
        <template v-if="column.key === 'count'">
          <a-tag color="green">{{ record.count_articles ?? 0 }} lượt xem</a-tag>
        </template>
        <template v-if="column.key === 'template'">
          <a-tag color="blue">{{ record.template.name ?? "-" }}</a-tag>
        </template>
        <template v-if="column.key === 'actions'">
          <a-space>
            <a :href="getPublicPageLink(record)" target="_blank">
              <a-button type="outline-success" size="small">
                <idist-feather-icons type="eye"/>
              </a-button>
            </a>
            <nuxt-link :to="{name: 'admin-pages-id-edit', params: { id: record.id } }">
              <a-button type="outline-primary" size="small">
                <idist-feather-icons type="edit"/>
              </a-button>
            </nuxt-link>
            <a-popconfirm
                placement="top"
                title="Xoá trang này?"
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
import {usePageStore} from "~/stores/pages.js";

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
        title: "Mẫu định dạng",
        key: "template",
        dataIndex: "template",
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
    getPublicPageLink(record) {
      return `${this.$config.public.portalDomain}/p/${record.slug}`;
    },
    async getData() {
      this.loading = true;
      let response = await usePageStore().ListEntries(this.pagination);
      this.entries = response.entries;
      this.pagination = response.pagination;
      this.loading = false;
    },
    async deleteData(entry) {
      this.loading = true;
      let response = await usePageStore().DeleteEntry(entry);
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
