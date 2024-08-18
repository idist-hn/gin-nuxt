<template>
  <a-page-header title="Thêm mới Từ khoá" @back="() => routerBack()"/>
  <a-form ref="FormData" layout="vertical" :rules="rules" :model="entry">
    <a-row :gutter="16">
      <a-col :md="16">
        <a-card>
          <tag-name :value="entry"/>
          <a-space :size="6">
            <a-button :loading="loading" :disabled="loading" type="primary" html-type="submit"
                      @click.prevent="submitForm()">
              Lưu
            </a-button>
            <a-popconfirm placement="top" ok-text="Đồng ý" cancel-text="Không" @confirm="routerBack">
              <template v-slot:title>
                <p>Huỷ tạo mới từ khoá?</p>
                <p>Bạn xác nhận huỷ tạo từ khoá này...</p>
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
import TagName from "~/components/admin/tags/tag-name.vue";
import {useTagStore} from "~/stores/tags.js";

export default {
  name: "create",
  components: {TagName},
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
      hash: ""
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
      this.$router.push({name: "admin-tags"});
    },
    async submitForm() {
      this.loading = true
      let response = await useTagStore().CreateEntry(this.entry);
      if (response && response.code === 200) {
        await this.$router.push({name: "admin-tags"});
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