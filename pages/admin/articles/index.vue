<template>
  <a-page-header title="Bài viết" @back="() => routerBack()">
    <template #extra>
      <nuxt-link :to="{ name: 'admin-articles-create' }">
        <a-button type="success" class="d-flex align-items-center">
          <idist-feather-icons type="plus" classes="mr-5"/>
          Tạo mới
        </a-button>
      </nuxt-link>
    </template>
  </a-page-header>

  <!--@TODO Bổ sung thanh filter-->
  <a-card>
    <a-table
        :loading="loading"
        :columns="columns"
        :data-source="entries"
        :pagination="pagination"
        @change="handleTableChange"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'thumbnail'">
          <a-avatar shape="square" size="large" :src="record.thumbnail"/>
        </template>

        <template v-if="column.key === 'title'">
          <div class="mb-10">{{ record.title }}</div>
          <div>
            <a-tag color="#3b5999">{{ record.category ? record.category.name : "-" }}</a-tag>
          </div>
        </template>
        <template v-if="column.key === 'is_highlight'">
          <a-switch :checked="record.is_highlight" disabled/>
        </template>
        <template v-if="column.key === 'is_hot'">
          <a-switch :checked="record.is_hot" disabled/>
        </template>
        <template v-if="column.key === 'status'">
          <a-tooltip placement="top">
            <template #title>
              <span>{{ record.publisher ? record.publisher.name : ""}} {{ record.published_at ? moment(record.published_at) : moment(record.updated_at) }}</span>
            </template>
            <a-tag :color="getStatusColor(record)">{{ getStatus(record) }}</a-tag>
          </a-tooltip>
        </template>
        <template v-if="column.key === 'count'">
          <a-tag color="blue">{{ record.count_views }} lượt xem</a-tag>
        </template>
        <template v-if="column.key === 'actions'">
          <a-space>
            <nuxt-link :to="{name: 'admin-articles-id-edit', params: { id: record.id } }">
              <a-button type="outline-primary" size="small">
                <idist-feather-icons type="edit"/>
              </a-button>
            </nuxt-link>
            <a-popconfirm
                placement="top"
                title="Xoá bài viết này?"
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
import IdistFeatherIcons from "~/composables/IdistFeatherIcons.vue";
import list from "~/mixins/list"
import {useArticleStore} from '~/stores/articles'

export default {
  name: "index",
  components: {IdistFeatherIcons},
  mixins: [list],
  setup: () => {
    definePageMeta({
      layout: "admin",
    });
  },
  computed: {},
  data: () => ({
    loading: true,
    entries: [],
    columns: [
      {
        dataIndex: "thumbnail",
        key: "thumbnail",
        title: "",
        slots: {title: "thumbnail"},
      },
      {
        dataIndex: "title",
        key: "title",
        title: "Tiêu đề",
        slots: {title: "title"},
      },
      {
        title: "Trạng thái",
        dataIndex: "status",
        key: "status",
        slots: {title: "status"},
      },
      {
        title: "Nổi bật",
        dataIndex: "is_highlight",
        key: "is_highlight",
        slots: {title: "is_highlight"},
      },
      {
        title: "Ưu tiên",
        dataIndex: "is_hot",
        key: "is_hot",
        slots: {title: "is_hot"},
      },
      {
        title: "Bộ đếm",
        key: "count",
        dataIndex: "count",
        slots: {title: "count"},
      },
      {
        title: "",
        key: "actions",
        dataIndex: "actions",
        className: "text-right",
        slots: {title: "actions"},
      },
    ]
  }),
  methods: {
    async getData() {
      this.loading = true;
      let response = await useArticleStore().ListEntries(this.pagination);
      this.entries = response.entries;
      this.pagination = response.pagination;
      this.loading = false;
    },
    getStatus(entry) {
      switch (entry.status) {
        case "draft":
          return "Bài nháp";
        case "published":
          return "Xuất bản";
        case "down":
          return "Đã thu hồi";
        default:
          return "Không xác định";
      }
    },
    getStatusColor(entry) {
      switch (entry.status) {
        case "draft":
          return "orange";
        case "published":
          return "green";
        case "down":
          return "red";
        default:
          return "blue";
      }
    },
    async deleteData(entry) {
      this.loading = true;
      // let response = await this.$store.dispatch("articles/DeleteArticle", entry);
      // if (response && response.code === 200) {
      //   await this.getData();
      // }

      this.$toast.show(response.message, {
        duration: 2000,
        type: response && response.code === 200 ? "success" : "danger",
      });
      this.loading = false;
    },
  },
}
</script>
<style scoped>

</style>