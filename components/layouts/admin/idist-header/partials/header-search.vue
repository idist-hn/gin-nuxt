<template>
  <div :theme="theme" class="certain-category-search-wrapper" :style="{ width: '100%' }">
    <a-row class="ant-row-middle">
      <a-col :md="2" :xs="1" class="text-right">
        <span class="certain-category-icon">
          <IdistFeatherIcons type="search" />
        </span>
      </a-col>
      <a-col :md="22" :xs="23">
        <a-popover
          placement="bottomLeft"
          v-model="visible"
          title="Kết quả tìm kiếm"
          trigger="click"
          overlayClassName="recommendation-search"
        >
          <template v-slot:content>
            <div v-for="item in filteredData" :key="item.id">
              <nuxt-link to="#">
                {{ item.title }}
                <span class="certain-search-item-count">{{ item.count }} people</span>
              </nuxt-link>
            </div>
            <div v-if="filteredData.length === 0">
              <nuxt-link to="#"> Không tìm thấy kết quả..... </nuxt-link>
            </div>
          </template>
          <a-input placeholder="Tìm kiếm..." @input="(e) => search(e, searchData)" />
        </a-popover>
      </a-col>
    </a-row>
  </div>
</template>

<script>

import {mapState} from 'pinia'
import {useLayoutStore} from '~/stores/layout'
import IdistFeatherIcons from "~/composables/IdistFeatherIcons.vue";
export default {
  name: "header-search",
  components: {IdistFeatherIcons},
  data: () => ({
    innerWidth: 1000,
    activeSearch: false,
    visible: false,
    searchData: [
      {
        id: 1,
        title: "AntDesign",
        count: 10000,
      },
      {
        id: 2,
        title: "Design UI",
        count: 10600,
      },

      {
        id: 3,
        title: "Bootstrap Design UI",
        count: 60100,
      },
      {
        id: 4,
        title: "Meterial design",
        count: 30010,
      },

      {
        id: 5,
        title: "AntDesign design language",
        count: 100000,
      },
    ],
    filteredData: [
      {
        id: 1,
        title: "AntDesign",
        count: 10000,
      },
      {
        id: 2,
        title: "Design UI",
        count: 10600,
      },

      {
        id: 3,
        title: "Bootstrap Design UI",
        count: 60100,
      },
      {
        id: 4,
        title: "Meterial design",
        count: 30010,
      },

      {
        id: 5,
        title: "AntDesign design language",
        count: 100000,
      },
    ],
  }),
  computed: {
    ...mapState(useLayoutStore, ['theme']),
  },
  methods: {
    toggleSearch() {
      this.activeSearch = !this.activeSearch;
    },
    search(e, searchData) {
      const data = searchData.filter((item) => {
        return item.title.startsWith(e.target.value);
      });
      this.filteredData = data;
    },
  },
  created() {
    if (process.client) {
      this.innerWidth = window.innerWidth;
    }
  },
  mounted() {},
};
</script>

<style scoped></style>
