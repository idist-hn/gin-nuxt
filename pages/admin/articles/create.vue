<template>
  <a-page-header title="Thêm mới Bài viết" @back="() => routerBack()"/>
  <a-form ref="FormData" layout="vertical" :rules="rules" :model="entry">
    <a-row :gutter="16">
      <a-col :xxl="18" :lg="18" :md="18" :sm="24">
        <a-card class="mb-10">
          <a-row :gutter="10">
            <a-col :md="18">
              <article-title :value="entry"/>
            </a-col>
            <a-col :md="6">
              <article-slug :value="entry"/>
            </a-col>
          </a-row>
          <article-description :value="entry"/>
        </a-card>
        <a-card class="mb-10 p-0" :body-style="{ padding: ' 0 !important' }">
          <article-content :value="entry"/>
        </a-card>
        <a-card class="mb-10">
          <article-relates :value="entry"/>
          <article-tags :value="entry"/>
        </a-card>
      </a-col>
      <a-col :xxl="6" :lg="6" :md="6" :sm="24">
        <a-card class="mb-10">
          <article-thumbnail :value="entry"/>
          <article-published-at :value="entry"/>
          <article-category :value="entry"/>
        </a-card>
        <a-card class="mb-10">
          <a-space :size="6">
            <a-button :loading="loading" :disabled="loading" type="primary" html-type="submit"
                      @click.prevent="submitForm('published')">
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
        <a-card class="mb-20">
          <article-highlight :value="entry"/>
          <article-hot :value="entry"/>
        </a-card>
      </a-col>
    </a-row>
  </a-form>
</template>
<script>

import ArticleTitle from "~/components/admin/articles/article-title.vue";
import ArticleSlug from "~/components/admin/articles/article-slug.vue";
import ArticleDescription from "~/components/admin/articles/article-description.vue";
import ArticleContent from "~/components/admin/articles/article-content.vue";
import ArticleHighlight from "~/components/admin/articles/article-highlight.vue";
import ArticleHot from "~/components/admin/articles/article-hot.vue";
import ArticleRelates from "~/components/admin/articles/article-relates.vue";
import ArticleTags from "~/components/admin/articles/article-tags.vue";
import {useArticleStore} from "~/stores/articles";
import ArticleThumbnail from "~/components/admin/articles/article-thumbnail.vue";
import ArticlePublishedAt from "~/components/admin/articles/article-published-at.vue";
import ArticleCategory from "~/components/admin/articles/article-category.vue";

export default {
  components: {
    ArticleCategory,
    ArticlePublishedAt,
    ArticleThumbnail,
    ArticleTags,
    ArticleRelates,
    ArticleHot, ArticleHighlight, ArticleContent, ArticleDescription, ArticleSlug, ArticleTitle
  },
  setup: () => {
    definePageMeta({
      layout: "admin",
    });
  },
  data: () => ({
    loading: false,
    rules: {
      title: [
        {
          required: true,
          message: "Không được bỏ trống",
          trigger: "blur",
        },
        {
          min: 10,
          message: "Độ dài tối thiểu 10 ký tự",
          trigger: "blur",
        },
        {
          max: 200,
          message: "Độ dài tối đa 200 ký tự",
          trigger: "blur",
        },
      ],
      description: [
        {
          max: 220,
          message: "Độ dài tối đa 220 ký tự",
          trigger: "blur",
        },
      ],
      content: [
        {
          required: true,
          message: "Không được bỏ trống",
          trigger: "blur",
        },
      ],
    },
    entry: {
      title: "",
      slug: "",
      description: "",
      content: "",
      status: "draft",
      thumbnail: "",
      published_at: null,
      category: null,
      related_articles: [],
      related_article_ids: [],
      tag_ids: [],
      tags: []
    },
  }),
  methods: {
    routerBack() {
      this.$router.push({name: "admin-articles"});
    },
    async submitForm() {
      try {
        await this.$refs.FormData.validate();
        let response = await useArticleStore().CreateEntry(this.entry);
        if (response && response.id !== "") {
          this.$router.push({name: "admin-articles"});
        }
        this.$toast.show(response.message, {
          duration: 2000,
          type: response && response.code === 200 ? "success" : "danger",
        });
      } catch (error) {
      }
    },
  },
}
</script>
<style scoped>

</style>