<template>
    <layout id="admin-layout">
      <auth-header :theme="theme"/>
      <div class="auth-content">
        <slot/>
      </div>
      <auth-footer/>
      <auth-helper/>
    </layout>
</template>

<script>
import {Layout} from "ant-design-vue";

const {Content} = Layout;

import {mapState} from 'pinia'
import {useLayoutStore} from '~/stores/layout'
import AuthHeader from "~/components/layouts/auth/auth-header.vue";
import AuthFooter from "~/components/layouts/auth/auth-footer.vue";
import AuthHelper from "~/components/layouts/auth/auth-helper.vue";

export default {
  name: `DefaultLayout`,
  components: {AuthHelper, AuthFooter, AuthHeader, Layout, Content},
  head: {
    title: "CMS PTIT",
  },
  middleware: ["authenticated"],
  data: () => ({}),
  computed: {
    ...mapState(useLayoutStore, ['theme']),
  },
  methods: {
    async LoadConfig() {
      // get;
    },
  },
  created() {
    if (process.client) {
      if (window.innerWidth < 991) {
        // this.$store.dispatch("layout/changeSideNavMode");
      }
    }
  },
  mounted() {
    console.log("Start admin layout");
    console.log("start get profile");
    // this.$store.dispatch("users/GetProfile");
    console.log("finish get profile");
  },
};
</script>

<style scoped></style>
