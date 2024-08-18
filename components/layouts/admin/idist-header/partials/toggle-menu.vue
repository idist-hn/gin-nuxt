<template>
  <a-button @click='toggleCollapsed' :type="theme === 'dark' ? `white` : `dark`">
    <img v-if='collapsed' src='~/assets/images/icons/right.svg' alt='menu' />
    <img v-else src='~/assets/images/icons/left.svg' alt='menu' />
  </a-button>
</template>

<script>
export default {
  name: 'ToggleMenu',
  props: {
    theme: String
  },
  methods: {
    toggleCollapsed() {
      this.$store.dispatch('layout/changeSideNavMode')
    }
  },
  computed: {
    collapsed() {
      // return this.$store.state.layout.collapsedMenu
    }
  },
}
</script>

<style scoped>

</style>
