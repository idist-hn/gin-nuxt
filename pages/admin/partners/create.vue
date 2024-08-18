<template>
  <a-page-header title="Tạo mới đối tác" @back="() => routerBack()"/>
  <a-form ref="FormData" layout="vertical" :rules="rules" :model="entry">
    <a-row :gutter="16">
      <a-col :md="8">
        <a-card class="mb-10">
          <partner-logo :value="entry"/>
        </a-card>
        <a-card class="mb-10">
          <partner-banner :value="entry"/>
        </a-card>
      </a-col>
      <a-col :md="16">
        <a-card>
          <a-row :gutter="16">
            <a-col :xl="24" :sm="24">
              <partner-name :value="entry"/>
            </a-col>
            <a-col :xl="24" :sm="24">
              <partner-description :value="entry"/>
            </a-col>
            <a-col :xl="24" :sm="24">
              <partner-url :value="entry"/>
            </a-col>
            <a-col :xl="24" :sm="24">
              <partner-active :value="entry"/>
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
                <p>Huỷ tạo mới đơn vị?</p>
                <p>Bạn xác nhận huỷ tạo đơn vị này...</p>
              </template>
              <a-button type="white" size="default" :outlined="true">Huỷ</a-button>
            </a-popconfirm>
          </div>
        </a-card>
      </a-col>
    </a-row>
  </a-form>
</template>

<script>

import PartnerLogo from "~/components/admin/partners/partner-logo.vue";
import PartnerBanner from "~/components/admin/partners/partner-banner.vue";
import PartnerName from "~/components/admin/partners/partner-name.vue";
import PartnerDescription from "~/components/admin/partners/partner-description.vue";
import PartnerUrl from "~/components/admin/partners/partner-url.vue";
import PartnerActive from "~/components/admin/partners/partner-active.vue";
import {usePartnerStore} from "~/stores/partners.js";

export default {
  name: "create",
  components: {PartnerActive, PartnerUrl, PartnerDescription, PartnerName, PartnerBanner, PartnerLogo},
  layout: "admin",
  setup() {
    definePageMeta({
      layout: "admin",
    });
  },
  data: () => ({
    loading: false,
    entry: {
      name: "",
      is_active: true,
      slug: "",
      url: "#",
      logo: "",
      banner: "",
      description: "",
    },
    rules: {
      name: [
        {
          required: true,
          message: "Không được bỏ trống",
          trigger: "blur",
        },
      ],
    },
  }),
  methods: {
    routerBack() {
      this.$router.push({name: "admin-partners"});
    },
    async SubmitForm() {
      this.loading = true;
      let response = await usePartnerStore().CreateEntry(this.entry);
      if (response && response.code === 200) {
        this.$toast(response.message, {autoClose: 2000, type: "success"});
      }
      this.loading = false;
      this.routerBack();
    },
  },
  created() {
  },
  mounted() {
  },
};
</script>

<style scoped></style>
