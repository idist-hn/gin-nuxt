<template>
  <a-row style="z-index: 1">
    <a-col
        :xxl="{ span: 8, offset: 8 }"
        :xl="{ span: 10, offset: 7 }"
        :lg="{ span: 12, offset: 6 }"
        :md="{ span: 20, offset: 2 }"
        :sm="{ span: 20, offset: 2 }"
    >
      <a-card style="margin-top: 10vh">
        <div class="text-center">
          <!--          <img src='~/assets/images/idist-logo.svg' alt='' width='128'>-->
          <img src="@/assets/images/ptit-logo.png" alt="" width="128"/>
          <h2 class="color-dark mt-20 mb-10">Welcome to CMS PTIT</h2>
          <h3 class="color-dark mt-20 mb-30">Login to CMS PTIT</h3>
        </div>
        <a-form ref="FormLogin" layout="vertical" :rules="rules" :model="entry">
          <a-row>
            <a-col :xxl="{ span: 24 }">
              <div>
                <a-alert type="error" v-if="message !== ''" class="mb-20" :message="message" closable/>
              </div>
            </a-col>
            <a-col
                :xxl="{ span: 18, offset: 3 }"
                :xl="{ span: 18, offset: 3 }"
                :lg="{ span: 18, offset: 3 }"
                :md="{ span: 18, offset: 3 }"
                :sm="{ span: 20, offset: 2 }"
            >
              <a-form-item prop="username" name="username" label="Tên đăng nhập">
                <a-input v-model:value="entry.username" placeholder="Tên đăng nhập..."/>
              </a-form-item>

              <a-form-item prop="password" name="password" label="Mật khẩu">
                <a-input-password v-model:value="entry.password" placeholder="Nhập mật khẩu"/>
              </a-form-item>
              <div class="d-flex">
                <a-checkbox class="flex" @change="entry.remember = !entry.remember"> Nhớ đăng nhập</a-checkbox>
              </div>
              <div class="mt-10 mb-30">
                <a-button
                    class="ant-btn-block"
                    size="large"
                    :loading="loading"
                    :disabled="loading"
                    type="primary-6"
                    html-type="submit"
                    @click.prevent="SubmitForm"
                >
                  Đăng nhập
                </a-button>
              </div>
            </a-col>
          </a-row>
        </a-form>
      </a-card>
      <div class="mt-20 text-center">
        <nuxt-link :to="{ name: 'auth-forgot_password' }">Quên mật khẩu?</nuxt-link>
      </div>
    </a-col>
  </a-row>
</template>

<script>

import {mapState} from 'pinia'
import {useProfileStore} from '~/stores/profile'
export default {
  name: "Login",
  setup() {
    definePageMeta({
      layout: "auth",
    })
  },
  data: () => ({
    loading: false,
    rules: {
      username: [
        {
          required: true,
          message: "Không được bỏ trống",
          trigger: "blur",
        },
      ],
      password: [
        {
          required: true,
          message: "Không được bỏ trống",
          trigger: "blur",
        },
      ],
    },
    entry: {
      username: "",
      password: "",
      remember: false,
    },
    message: "",
    message_description: "",
  }),
  methods: {
    async SubmitForm() {
      this.loading = true;
      this.message = "";
      let response = await useProfileStore().PostLogin(this.entry);
      if (response === undefined) {
        await this.$toast("Đăng nhập thất bại", {autoClose: 5000, type: "error"});
        this.loading = false;
      } else if (response && response.code === 200) {
        this.message = "Login success";
        await this.$toast("Đăng nhập thành công", {autoClose: 5000, type: "success"});
        await this.$router.push({name: "admin-homepage"});

      } else {
        this.message = response.message;
        this.message_description = response.error;
        this.loading = false;
      }
      this.loading = false;
    }
  }
};
</script>

<style scoped></style>
