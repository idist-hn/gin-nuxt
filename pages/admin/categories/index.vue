<template>
  <a-page-header title="Danh mục" @back="() => routerBack()">
    <template v-slot:extra>
      <nuxt-link :to="{ name: 'admin-categories-create' }">
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
        <template v-if="column.key === 'thumbnail'">
          <a-avatar :src="record.thumbnail" shape="square" :size="64"/>
        </template>
        <template v-if="column.key === 'name'">
          <span>{{ record.name }}</span>
        </template>
        <template v-if="column.key === 'description'">
          {{ record.description != "" ? record.description : '-' }}
        </template>
        <template v-if="column.key === 'count'">
          <a-tag color="green">{{ record.count_articles }} Bài viết</a-tag>
        </template>
        <template v-if="column.key === 'actions'">
          <a-space>
            <a target="_blank" :href="publicLinkCategory(record)">
              <a-button type="outline-info" size="small">
                <idist-feather-icons type="eye"/>
              </a-button>
            </a>
            <nuxt-link :to="{name: 'admin-categories-id-edit', params: { id: record.id } }">
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
import IdistFeatherIcons from "/components/commons/IdistFeatherIcons.vue";
import list from '~/mixins/list';
import {useCategoryStore} from "~/stores/categories.js";
import {useTagStore} from "~/stores/tags.js";

export default {
  name: "index",
  components: {IdistFeatherIcons},
  mixins: [list],
  setup() {
    definePageMeta({
      layout: "admin",
    });
  },
  data: () => ({
    loading: true,
    entries: [],
    columns: [
      {
        title: 'Ảnh bìa',
        dataIndex: 'thumbnail',
        key: 'thumbnail',
      },
      {
        title: 'Tên danh mục',
        dataIndex: 'name',
        key: 'name',
      },
      {
        title: 'Đường dẫn',
        dataIndex: 'slug',
        key: 'slug',
      },
      {
        title: 'Mô tả',
        dataIndex: 'description',
        key: 'description',
      },
      {
        title: 'Thống kê',
        dataIndex: 'count',
        key: 'count',
      },
      {
        title: 'Hành động',
        key: 'actions',
        align: 'right',
      },
    ],
  }),
  methods: {
    routerBack() {
      this.$router.push({name: "admin-homepage"});
    },
    publicLinkCategory(entry){
      return `${this.$config.public.portalDomain}/danh-muc/${entry.slug}`;
    },
    async getData() {
      this.loading = true;
      const {entries, pagination} = await useCategoryStore().ListEntries(this.pagination);
      this.entries = entries;

      this.pagination = pagination;
      this.loading = false;
    },
    async deleteData(entry) {
      this.loading = true;
      let response = await useCategoryStore().DeleteEntry(entry);
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
}
</script>

<style scoped lang="scss">

</style>