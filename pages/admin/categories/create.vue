<template>
  <a-page-header title="Tạo mới danh mục" @back="() => routerBack()"/>
  <a-row :gutter="16">
    <a-col :xxl="16" :lg="16" :md="16" :sm="24">
      <a-card>
        <a-form v-if="!loading" ref="FormLogin" layout="vertical" :rules="rules" :model="entry">
          <a-row :gutter="16">
            <a-col :xl="12" :sm="24">
              <category-name v-model:value="entry"/>
            </a-col>
            <a-col :xl="12" :sm="24">
              <category-slug v-model:value="entry"/>
            </a-col>
          </a-row>
          <a-row :gutter="16">
            <a-col :xl="12" :sm="24">
              <category-active v-model:value="entry"/>
            </a-col>
            <a-col :xl="12" :sm="24">
                <category-parent v-model:value="entry"/>
            </a-col>
            <a-col :sm="24">
              <category-description v-model:value="entry"/>
            </a-col>
          </a-row>

          <div class="mt-10 mb-30">
            <a-button
                :loading="loading"
                :disabled="loading"
                type="primary"
                html-type="submit"
                @click.prevent="SubmitForm"
            >
              Lưu
            </a-button>
            <a-popconfirm placement="top" ok-text="Đồng ý" cancel-text="Không" @confirm="routerBack">
              <template v-slot:title>
                <p>Huỷ tạo mới danh mục?</p>
                <p>Bạn xác nhận huỷ tạo danh mục này...</p>
              </template>
              <a-button type="white" size="default" :outlined="true">Huỷ</a-button>
            </a-popconfirm>
          </div>
        </a-form>
        <div v-else>
          <a-spin/>
        </div>
      </a-card>
    </a-col>
    <a-col :xxl="8" :lg="8" :md="8" :sm="24">
      <category-children v-model:value="entry"/>
    </a-col>
  </a-row>
</template>

<script>
import CategoryName from "~/components/admin/categories/category-name.vue";
import CategoryParent from "~/components/admin/categories/category-parent.vue";
import CategoryChildren from "~/components/admin/categories/category-children.vue";
import CategorySlug from "~/components/admin/categories/category-slug.vue";
import CategoryActive from "~/components/admin/categories/category_active.vue";
import CategoryDescription from "~/components/admin/categories/category-description.vue";
import {useCategoryStore} from "~/stores/categories.js";

export default {
  name: "create",
  components: {CategoryDescription, CategoryActive, CategorySlug, CategoryChildren, CategoryParent, CategoryName},
  setup() {
    definePageMeta({
      layout: "admin",
    });
  },
  data: () => ({
    loading: false,
    entry: {
      name: "",
      slug: "",
      description: "",
      is_active: true,
      parent: null,
      children: [],
    },
    rules: {
      name: [
        {
          required: true,
          message: "Không được bỏ trống",
          trigger: "blur",
        },
      ],
      slug: [
        {
          required: true,
          message: "Không được bỏ trống",
          trigger: "blur",
        },
      ],
    }
  }),
  methods: {
    routerBack() {
      this.$router.push({name: "admin-categories"});
    },
    async SubmitForm() {
      this.loading = true
      let response = await useCategoryStore().CreateEntry(this.entry);
      if (response && response.code === 200) {
        await this.$router.push({name: "admin-categories"});
        await this.$toast(response.message, {autoClose: 5000, type: "success"});
      } else {
        await this.$toast(response.message, {autoClose: 5000, type: "error"});
      }
      this.loading = false
    },
  },
}
</script>

<style scoped lang="scss">

</style>