<template>
  <div class='themes'>
    <a to='#' class='head-example' @click='changeTheme()'>
      <IdistFeatherIcons v-if="theme === 'light'" type='sun' :size='20' />
      <IdistFeatherIcons v-else type='moon' :size='20' />
    </a>
  </div>
</template>

<script>
import IdistFeatherIcons from "~/components/commons/IdistFeatherIcons.vue";

import {mapState} from 'pinia'
import {useLayoutStore} from '~/stores/layout'

export default {
  name: 'header-theme',
  components: {IdistFeatherIcons},
  data: () => ({}),
  methods: {
    changeTheme() {
      useLayoutStore().changeThemeMode()
      // this.$store.dispatch('layout/changeThemeMode')
    }
  },
  computed: {
    ...mapState(useLayoutStore, ['theme']),
  },
  created() {
  },
  mounted() {
  }
}
</script>

<style scoped>

</style>
