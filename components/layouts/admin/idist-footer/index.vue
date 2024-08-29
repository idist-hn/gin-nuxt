<template>
  <footer id="idist-footer">
    <a-row class="admin-footer">
      <a-col :md="12" :xs="24">
        <span class="admin-footer__copyright"> Copyright ©2021 Produced by PTIT </span>
      </a-col>
      <a-col :md="12" :xs="24">
        <div class="admin-footer__links">
          <router-link to="#">Điều khoản điều lệ</router-link>
          <router-link to="#">Liên hệ</router-link>
        </div>
      </a-col>
    </a-row>
  </footer>
</template>

<script>
export default {
  name: "index",
  data: () => ({}),
  methods: {},
  created() {},
  mounted() {},
};
</script>

<style scoped></style>
