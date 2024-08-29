<template>
  <div class='support'>
    <a-popover placement='bottomLeft' trigger='click'>
      <template v-slot:content>
        <div class='support-dropdown'>
          <ul>
            <h4>Documentation</h4>
            <li>
              <nuxt-link to='#'>How to customize admin</nuxt-link>
            </li>
            <li>
              <nuxt-link to='#'>How to use</nuxt-link>
            </li>
            <li>
              <nuxt-link to='#'>The relation of vertical spacing</nuxt-link>
            </li>
          </ul>
          <ul>
            <h4>Payments</h4>
            <li>
              <nuxt-link to='#'>How to customize admin</nuxt-link>
            </li>
            <li>
              <nuxt-link to='#'>How to use</nuxt-link>
            </li>
          </ul>
          <ul>
            <h4>Content Planner</h4>
            <li>
              <nuxt-link to='#'>How to customize admin</nuxt-link>
            </li>
            <li>
              <nuxt-link to='#'>How to use</nuxt-link>
            </li>
          </ul>
        </div>
      </template>
      <a to='#' class='head-example'>
        <IdistFeatherIcons type='help-circle' :size='20' />
      </a>
    </a-popover>
  </div>
</template>

<script>
import IdistFeatherIcons from "~/components/commons/IdistFeatherIcons.vue";

export default {
  name: 'header-support',
  components: {IdistFeatherIcons},
  data: () => ({}),
  methods: {},
  created() {
  },
  mounted() {
  }
}
</script>

<style scoped>

</style>
