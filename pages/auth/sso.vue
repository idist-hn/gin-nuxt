<template>
  <a-row style='z-index: 1'>
    <a-col :xxl='{span:8, offset: 8}'
           :xl='{span:10, offset:7}'
           :lg='{span:12, offset:6}'
           :md='{span:20, offset:2}'
           :sm='{span:20, offset:2}'>
      <a-card style='margin-top: 3rem'>
        <div class='text-center'>
          <img src='~/assets/images/byt-logo.svg' alt=''>
          <h2 class='color-primary mt-20 mb-30'>
            HỆ THỐNG QUẢN LÝ QUỐC GIA VỀ ĐĂNG KÝ,
            CẤP PHÉP HÀNH NGHỀ KHÁM BỆNH, CHỮA BỆNH
          </h2>

          <template v-if='$route.query.error'>
            <h3 class='color-danger'>{{ $route.query.error ? 'Truy cập bị từ chối' : '' }}</h3>
            <h4>
              {{ $route.query.error_description }}
            </h4>
            <div>
              <a-button type='danger' class='ant-btn' @click.prevent='CloseWindow'>Đóng</a-button>
            </div>
          </template>

          <div v-else-if='loading' class='spin mt-30' style='position: relative'>
            <a-spin />
          </div>
          <div v-else>
            <template>
              <h2 class='color-success'>Đăng nhập thành công</h2>
              <div>
                <a-button type='success' class='ant-btn' @click.prevent='CloseWindow' :loading='true'>Chuyển hướng sau
                  <span>&nbsp; {{ entry.countdown }} giây</span></a-button>
              </div>
            </template>
          </div>
        </div>

      </a-card>
    </a-col>
  </a-row>
</template>

<script>
export default {
  name: 'sso',
  layout: 'auth',
  data: () => ({
    loading: false,
    entry: {
      countdown: 10
    }
  }),
  methods: {
    CloseWindow() {
      if (process.client) {
        window.close()
      }
    },
    getCountDown() {
      let self = this
      let timer = setInterval(function() {
        if (self.entry.countdown > 0) {
          self.entry.countdown--
        } else {
          self.CloseWindow()
          clearInterval(timer)
        }
      }, 1000)
    },
    async getData() {
      this.loading = true
      let response = await this.$store.dispatch('auth/CodeAuthorization', this.$route.query.code)
      this.entry = response.data
      this.getCountDown()
      this.loading = false
    }
  },
  created() {
    if (this.$route.query.code) {
      this.getData()
    }
  },
  mounted() {

  }
}
</script>

<style scoped>

</style>
