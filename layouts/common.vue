<template>
    <layout id="web-layout">
      <sub-menu></sub-menu>
      <div class="web-content">
      <slot/>
      </div>
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
import SubMenu from "~/components/layouts/common/idist-header/sub-menu.vue";

export default {
  name: `DefaultLayout`,
  components: {SubMenu, AuthHelper, AuthFooter, AuthHeader, Layout, Content},
  head: {
    title: "CIE PTIT",
  },
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
    console.log("Start common layout");
  },
};
</script>

<style scoped></style>
