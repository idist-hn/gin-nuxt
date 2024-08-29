<template>
  <config-provider :theme="{cssVar: true, hashed: false }">
    <a-layout id="auth-layout">
      <auth-header/>
      <a-layout-content class="auth-content">
        <slot/>
      </a-layout-content>
      <auth-footer/>
      <auth-helper/>
      <!--    <img-->
      <!--        src="~/assets/images/layout-element-1.png"-->
      <!--        :style="setElementPosition()"-->
      <!--        v-anime="{ rotate: 360, duration: 3000 }"-->
      <!--        alt=""-->
      <!--    />-->
      <!--    <img-->
      <!--        src="~/assets/images/layout-element-2.png"-->
      <!--        :style="setElementPosition()"-->
      <!--        v-anime="{ translateX: 180, duration: 3000 }"-->
      <!--        alt=""-->
      <!--    />-->
      <!--    <img-->
      <!--        src="~/assets/images/layout-element-3.png"-->
      <!--        :style="setElementPosition()"-->
      <!--        v-anime="{ rotate: 360, duration: 3000 }"-->
      <!--        alt=""-->
      <!--    />-->
      <!--    <img-->
      <!--        src="~/assets/images/layout-element-4.png"-->
      <!--        :style="setElementPosition()"-->
      <!--        v-anime="{ rotate: 360, duration: 3000 }"-->
      <!--        alt=""-->
      <!--    />-->
    </a-layout>
  </config-provider>
</template>

<script>

import AuthHelper from "~/components/layouts/auth/auth-helper.vue";
import AuthHeader from "~/components/layouts/auth/auth-header.vue";
import AuthFooter from "~/components/layouts/auth/auth-footer.vue";
import {ConfigProvider} from "ant-design-vue";

export default {
  name: "LayoutAuth",
  components: {ConfigProvider, AuthFooter, AuthHeader, AuthHelper},
  head: {
    title: "CMS PTIT",
  },
  setup() {
    // definePageMeta({
    //   head: {
    //     title: "CMS PTIT",
    //   }
    // });
  },
  methods: {
    setElementPosition() {
      return {
        top: Math.floor(Math.random() * 100) + "vh",
        right: Math.floor(Math.random() * 100) + "vw",
        position: "absolute",
      };
    },
  },
};
</script>

<style scoped></style>
