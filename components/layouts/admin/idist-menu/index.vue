<template>
  <a-layout-sider :width='280' :style="{
              margin: '63px 0 0 0',
              padding: '15px 15px 155px 15px',
              // overflowY: 'auto',
              height: '100vh',
              position: 'fixed',
              left: 0,
              zIndex: 998}" :collapsed='collapsedMenu' :theme='theme'>
    <perfect-scrollbar :options='{
                    wheelSpeed: 1,
                    swipeEasing: true,
                    suppressScrollX: true,
                  }'>
      <aside-items :toggleCollapsed='toggleCollapsedMobile' :theme='theme'/>
    </perfect-scrollbar>
    <div @click='toggleCollapsed' class='menu-collapse-button' style=''>
      <IdistFeatherIcons v-if='collapsedMenu' type='chevron-right'/>
      <IdistFeatherIcons v-else type='chevron-left'/>
    </div>
  </a-layout-sider>

</template>

<script>
import {PerfectScrollbar} from 'vue3-perfect-scrollbar'
import AsideItems from './items.vue'
import Layout from "ant-design-vue"
import IdistFeatherIcons from "~/composables/IdistFeatherIcons.vue";


import {mapState} from 'pinia'
import {useLayoutStore} from '~/stores/layout'

export default {
  name: 'index',
  components: {IdistFeatherIcons, PerfectScrollbar, AsideItems},
  props: {
    theme: {
      type: String,
      default: () => ('light')
    }
  },
  computed: {
    ...mapState(useLayoutStore, ['collapsedMenu']),
    // collapsed() {
    //   console.log(layout.state.collapsedMenu)
    //   return layout.state.collapsedMenu
    // }
  },
  methods: {
    toggleCollapsed() {
      useLayoutStore().changeSideNavMode()
    },
    toggleCollapsedMobile() {
      if (process.client && window.innerWidth < 991) {
        useLayoutStore().changeSideNavMode()
      }
    }
  },
}
</script>

<style lang='scss'>
//.ant-layout-sider {
//  .ps {
//    height: calc(100vh - 100px)
//  }
//}
</style>
