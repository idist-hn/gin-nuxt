<template>
  <a-form-item prop="password" name="password" label="Mật khẩu">
    <a-input-password class="custom-password-field" v-model:value="value.password" placeholder="Nhập mật khẩu"/>
  </a-form-item>
</template>

<script>
export default {
  name: "UserPassword",
  props: {
    value: Object,
  },
};
</script>

<style  lang="scss">
.custom-password-field {
  padding: 0 !important;
}
</style>
