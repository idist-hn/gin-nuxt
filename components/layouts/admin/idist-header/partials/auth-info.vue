<template>
  <div class="info-wrapper">
    <header-message />
    <header-notification />
    <header-settings />
    <header-support />
    <header-theme />
    <header-user />
  </div>
</template>

<script>
import HeaderMessage from "./header-message.vue";
import HeaderNotification from "./header-notification.vue";
import HeaderSettings from "~/components/layouts/admin/idist-header/partials/header-settings.vue";
import HeaderSupport from "~/components/layouts/admin/idist-header/partials/header-support.vue";
import HeaderUser from "~/components/layouts/admin/idist-header/partials/header-user.vue";
import HeaderTheme from "~/components/layouts/admin/idist-header/partials/header-theme.vue";

export default {
  name: "auth-info",
  components: { HeaderTheme, HeaderUser, HeaderSupport, HeaderSettings, HeaderNotification, HeaderMessage },
  data: () => ({}),
  methods: {},
  created() {},
  mounted() {},
};
</script>

<style scoped></style>
