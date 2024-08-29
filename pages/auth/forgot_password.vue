<template>
  <a-row style='z-index: 1'>
    <a-col :xl='{span:8, offset:8}' :lg='{span:12, offset:6}' :md='{span:20, offset:2}' :sm='{span:20, offset:2}'>
      <a-card style='margin-top: 3rem'>
        <div class='text-center'>
          <img src='~/assets/images/ptit-logo.png' alt='' width='128'>
          <h2 class='color-primary mt-20 mb-30'>
            PTIT's Portal Content Management System
          </h2>
        </div>
        <a-form-model ref='FormLogin' layout='vertical' :rules='rules' :model='entry'>
          <a-row>
            <a-col :xl='{span:12, offset:6}' :lg='{span:12, offset:6}' :md='{span:18, offset:3}'
                   :sm='{span:18, offset:3}'>
              <a-form-model-item prop='account' name='account' label='Tên đăng nhập'>
                <a-input v-model:value='entry.account' placeholder='Tên đăng nhập...'/>
              </a-form-model-item>
              <a-row :gutter='10'>
                <a-col :xs='14'>
                  <a-form-model-item prop='confirm_code' name='confirm_code' label=''>
                    <a-input v-model:value='entry.confirm_code' placeholder='Nhập mã xác thực'/>
                  </a-form-model-item>
                </a-col>
                <a-col :xs='10'>
                  <!--                  <img src='/static/capcha.jpg' alt='' style='width:100%;height: 48px; border-radius: 4px'>-->
                </a-col>
              </a-row>
              <div class='mb-30'>
                <a-button class='ant-btn-block' size='large' :loading='loading' :disabled='loading' type='primary-6'
                          html-type='submit' @click.prevent='SubmitForm'>
                  Hoàn thành
                </a-button>
              </div>
            </a-col>
          </a-row>
        </a-form-model>

      </a-card>
    </a-col>
  </a-row>
</template>

<script>

export default {
  name: 'ForgotPassword',
  layout: 'auth',
  data: () => ({
    loading: false,
    rules: {
      account: [
        {
          required: true,
          message: 'Không được bỏ trống',
          trigger: 'blur'
        }
      ],
      confirm_code: [
        {
          required: true,
          message: 'Không được bỏ trống',
          trigger: 'blur'
        }
      ]
    },
    entry: {
      account: '',
      confirm_code: '',
      remember: false
    }
  }),
  created() {
  },
  mounted() {
  },
  methods: {
    SubmitForm() {
      let self = this
      this.$refs.FormLogin.validate(valid => {
        if (valid) {
          this.loading = true
          setTimeout(function () {
            self.loading = false
          }, 3000)
        } else {
          this.loading = false
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
