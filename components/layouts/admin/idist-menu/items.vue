<template>
  <a-menu v-model:openKeys="openKeys" v-model:selectedKeys="selectedKeys" :mode="mode" :theme="theme">
    <a-menu-item-group key="general">
      <template v-slot:title>
        <p class="sidebar-nav-title">Quản lý</p>
      </template>
      <a-menu-item @click="toggleCollapsed" key="admin-homepage">
        <nuxt-link :to="{ name: 'admin-homepage' }">
          <IdistFeatherIcons type="home"/>
          <span> Trang chủ </span>
        </nuxt-link>
      </a-menu-item>
      <a-menu-item @click="toggleCollapsed" key="admin-articles">
        <nuxt-link :to="{ name: 'admin-articles' }">
          <IdistFeatherIcons type="file"/>
          <span> Tin tức </span>
        </nuxt-link>
      </a-menu-item>
      <a-menu-item @click="toggleCollapsed" key="admin-tags">
        <nuxt-link :to="{ name: 'admin-tags' }">
          <IdistFeatherIcons type="tag"/>
          <span> Từ khoá </span>
        </nuxt-link>
      </a-menu-item>
      <a-menu-item @click="toggleCollapsed" key="admin-categories">
        <nuxt-link :to="{ name: 'admin-categories' }">
          <IdistFeatherIcons type="folder"/>
          <span> Danh mục </span>
        </nuxt-link>
      </a-menu-item>
      <a-menu-item @click="toggleCollapsed" key="admin-menus">
        <nuxt-link :to="{ name: 'admin-menus' }">
          <IdistFeatherIcons type="menu"/>
          <span> Menu </span>
        </nuxt-link>
      </a-menu-item>
      <a-menu-item @click="toggleCollapsed" key="admin-partners">
        <nuxt-link :to="{ name: 'admin-partners' }">
          <IdistFeatherIcons type="pocket"/>
          <span> Đối tác </span>
        </nuxt-link>
      </a-menu-item>
      <a-menu-item @click="toggleCollapsed" key="admin-departments">
        <nuxt-link :to="{ name: 'admin-departments' }">
          <IdistFeatherIcons type="hash"/>
          <span> Phòng ban </span>
        </nuxt-link>
      </a-menu-item>
      <a-menu-item @click="toggleCollapsed" key="admin-pages">
        <nuxt-link :to="{ name: 'admin-pages' }">
          <IdistFeatherIcons type="file-text"/>
          <span> Trang giới thiệu </span>
        </nuxt-link>
      </a-menu-item>
      <a-menu-item @click="toggleCollapsed" key="admin-users">
        <nuxt-link :to="{ name: 'admin-users' }">
          <IdistFeatherIcons type="user"/>
          <span> Người dùng </span>
        </nuxt-link>
      </a-menu-item>
      <a-menu-item @click="toggleCollapsed" key="admin-roles">
        <nuxt-link :to="{ name: 'admin-roles' }">
          <IdistFeatherIcons type="package"/>
          <span> Phân quyền </span>
        </nuxt-link>
      </a-menu-item>
      <a-menu-item @click="toggleCollapsed" key="admin-systems">
        <nuxt-link :to="{ name: 'admin-systems' }">
          <IdistFeatherIcons type="framer"/>
          <span> Hệ thống </span>
        </nuxt-link>
      </a-menu-item>
    </a-menu-item-group>
  </a-menu>
</template>

<script>
import IdistFeatherIcons from "/components/commons/IdistFeatherIcons.vue";
import {mapState} from 'pinia'
import {useLayoutStore} from '~/stores/layout'
export default {
  name: "AsideItems",
  components: {IdistFeatherIcons},
  props: {
    toggleCollapsed: Function,
    events: Object,
    theme: {
      type: String,
      default: () => "light",
    },
  },
  computed: {
    ...mapState(useLayoutStore, ['collapsedMenu']),

  },
  data: () => ({
    mode: "inline",
    selectedKeys: ["admin-homepage"],
    openKeys: ["admin-homepage"],
    preOpenKeys: ["admin-homepage"],
  }),
  watch: {
    openKeys: {
      handler: function (val, oldVal) {
        this.preOpenKeys = oldVal;
      },
      deep: true,
    },
    "$route.matched": {
      handler: function () {
        this.checkRoute();
      },
      deep: true,
    },
  },
  methods: {
    checkRoute() {
      if (this.$route.matched.length) {
        this.selectedKeys = [this.$route.matched[0].name];
        this.openKeys = [this.$route.matched[0].name];
        this.preOpenKeys = [this.$route.matched[0].name];
      }
    },
  },
  mounted() {
    this.checkRoute();
  },
};
</script>
