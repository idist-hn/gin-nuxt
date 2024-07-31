<template>
  <a-page-header title="Thêm mới trang giới thiệu" @back="() => routerBack()"/>
  <a-form ref="FormData" layout="vertical" :rules="rules" :model="entry">
    <a-row :gutter="16">
      <a-col :xxl="18" :lg="18" :md="18" :sm="24">
        <a-card>
          <page-thumbnail :value="entry"/>
          <page-title :value="entry"/>
          <page-description :value="entry"/>
        </a-card>

        <a-card>
          <page-content :value="entry"/>
        </a-card>
      </a-col>
      <a-col :xxl="6" :lg="6" :md="6" :sm="24">
        <a-card class="mb-20">
          <page-template  :value="entry"/>
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
            <a-popconfirm placement="top" ok-text="Đồng ý" cancel-text="Không" @confirm="routerBack">
              <template v-slot:title>
                <p>Huỷ tạo mới bài viết?</p>
                <p>Bạn xác nhận huỷ tạo bài viết này...</p>
              </template>
              <a-button type="white" size="default" :outlined="true">Huỷ</a-button>
            </a-popconfirm>
          </a-space>
        </a-card>
      </a-col>
    </a-row>
  </a-form>
</template>

<script>
import {usePageStore} from "~/stores/pages.js";
import PageThumbnail from "~/components/admin/pages/page-thumbnail.vue";
import PageTitle from "~/components/admin/pages/page-title.vue";
import PageContent from "~/components/admin/pages/page-content.vue";
import PageDescription from "~/components/admin/pages/page-description.vue";
import PageTemplate from "~/components/admin/pages/page-template.vue";

export default {
  name: "create",
  components: {PageTemplate, PageDescription, PageContent, PageTitle, PageThumbnail},
  setup() {
    definePageMeta({
      layout: "admin",
    });
  },
  data: () => ({
    loading: false,
    entry: {
      title: "",
      description: "",
      slug: "",
      content: "",
      thumbnail: "",
    },
    rules: {
      name: [
        {required: true, message: "Vui lòng nhập tên từ khoá", trigger: "blur"},
        {min: 3, message: "Tên từ khoá phải có ít nhất 3 ký tự", trigger: "blur"},
        {max: 255, message: "Tên từ khoá không được vượt quá 255 ký tự", trigger: "blur"},
      ],
    },
  }),
  methods: {
    routerBack() {
      this.$router.push({name: "admin-pages"});
    },
    async submitForm() {
      this.loading = true
      let response = await usePageStore().CreateEntry(this.entry);
      if (response && response.code === 200) {
        await this.$router.push({name: "admin-pages"});
        await this.$toast("Thêm mới thành công", {autoClose: 5000, type: "success"});
      } else {
        await this.$toast("Thêm mới thất bại", {autoClose: 5000, type: "error"});
      }
      this.loading = false
    },
  },
}
</script>
<style scoped>

</style>