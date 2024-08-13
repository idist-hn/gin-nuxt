<template>
  <a-card headless class="profile-author-box">
    <div class="author-info">
      <figure>
        <img :src="require('~/assets/images/no-image.png')" alt="" width="120" />
        <a-upload :action="getLinkUploadAvatar()" />
      </figure>
      <figcaption>
        <div class="info">
          <h4>{{ user ? user.name : "" }}</h4>
          <p>{{ user && user.role ? user.role.name : "Đang cập nhật" }}</p>
        </div>
      </figcaption>
    </div>
    <nav class="profile-menu-list">
      <ul>
        <li>
          <router-link :to="{ name: 'admin-profile' }">
            <idist-feather-icons type="user" :size="14" />
            Thông tin cá nhân
          </router-link>
        </li>
        <li>
          <router-link :to="{ name: 'admin-profile-settings' }">
            <idist-feather-icons type="settings" :size="14" />
            Cài đặt tài khoản
          </router-link>
        </li>
        <li>
          <router-link :to="{ name: 'admin-profile-change_password' }">
            <idist-feather-icons type="key" :size="14" />
            Đổi mật khẩu
          </router-link>
        </li>
      </ul>
    </nav>
  </a-card>
</template>

<script>
export default {
  name: "OverviewCard",
  data: () => ({}),
  computed: {
    user() {
      return this.$store.state.users.profile ?? null;
    },
  },
  methods: {
    getLinkUploadAvatar() {
      return "";
    },
  },
  created() {},
  mounted() {},
};
</script>

<style scoped></style>
