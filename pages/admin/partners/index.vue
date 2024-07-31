<template>
  <a-page-header title="Đối tác" @back="() => routerBack()">
    <template v-slot:extra>
      <nuxt-link :to="{ name: 'admin-partners-create' }">
        <a-button type="success">
          <idist-feather-icons type="plus"/>
          Tạo mới
        </a-button>
      </nuxt-link>
    </template>
  </a-page-header>
  <a-row v-if="!loading" :gutter="16">
    <a-col :md="24" v-if="entries.length === 0">
      <a-card>
        <p class="text-center font-italic">Không có dữ liệu</p>
      </a-card>
    </a-col>
    <a-col :md="8" :lg="6" v-for="entry in entries" :key="entry.id">
      <a-card class="text-center partner-item mb-10">
        <div>
          <img :src="entry.logo" alt="" width="120" height="72" style="object-fit: cover; object-position: center"/>
          <div class="partner-status">
            <a-switch :checked="entry.is_active" disabled/>
          </div>
        </div>
        <div class="">
          <h3>{{ entry.name }}</h3>
          <p>{{ entry.description }}</p>
        </div>
        <div class="partner-actions">
          <a-space :size="10">
            <nuxt-link :to="{ name: 'admin-partners-id-edit', params: { id: entry.id } }">
              <a-button type="outline-primary" size="small">
                Chỉnh sửa
              </a-button>
            </nuxt-link>
            <a-popconfirm
                placement="top"
                title="Xoá nội dung này?"
                ok-text="Đồng ý"
                cancel-text="Không"
                @confirm="() => deleteData(entry)"
            >
              <a-button type="outline-danger" size="small">
                Xoá bỏ
              </a-button>
            </a-popconfirm>
          </a-space>
        </div>
      </a-card>
    </a-col>
  </a-row>

  <a-card v-else>
    <div class="text-center" style="padding: 50px 0">
      <a-spin/>
    </div>
  </a-card>
</template>

<script>
import list from "~/mixins/list";
import {usePartnerStore} from "~/stores/partners.js";
import IdistFeatherIcons from "~/composables/IdistFeatherIcons.vue";

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
    async getData() {
      this.loading = true;
      this.pagination.pageSize = 9999;
      let response = await usePartnerStore().ListEntries(this.pagination);
      this.entries = response.entries;
      this.pagination = response.pagination;
      this.loading = false;
    },
    async deleteData(entry) {
      this.loading = true;
      let response = await usePartnerStore().DeleteEntry(entry);
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
