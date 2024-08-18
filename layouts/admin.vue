<template>
  <config-provider :theme="{cssVar: true, hashed: true }">
    <!--    <transition>-->
    <layout id="admin-layout">
      <layouts-admin-idist-header :theme="theme"/>
      <layout class="layout">
        <layouts-admin-idist-menu :theme="theme"/>
        <layout class="atbd-main-layout">
          <content id="main-section">
            <slot/>
          </content>
          <layouts-admin-idist-footer/>
        </layout>
      </layout>
    </layout>
    <!--    </transition>-->
  </config-provider>
</template>

<script>
import {Layout, ConfigProvider} from "ant-design-vue";

const {Content} = Layout;

import {mapState} from 'pinia'
import {useLayoutStore} from '~/stores/layout'
import {useProfileStore} from '~/stores/profile'

export default {
  name: `DefaultLayout`,
  components: {Layout, Content, ConfigProvider},
  head: {
    title: "CMS PTIT",
  },

  setup() {
    definePageMeta({
      middleware: "auth",
      // layout: "admin"
    })
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
    // useProfileStore().GetProfile()
  },
};
</script>

<style scoped></style>
