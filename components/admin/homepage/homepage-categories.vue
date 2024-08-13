<template>
  <a-card title="Trang được truy cập nhiều nhất" style="width: 100%" class="mb-20">
    <template #extra>
      <div class="card-nav">
        <ul>
          <li v-for="tab in tabs" :key="tab.key" :class="activeTable === tab.key ? 'active' : 'deactivate'">
            <a @click="(e) => handleChangeActive(e, tab.key)" to="#" style="font-weight: bold"> {{ tab.name }}</a>
          </li>
        </ul>
      </div>
    </template>

    <div v-if="loading" class="sd-spin">
      <a-spin/>
    </div>
    <a-row v-else :gutter="20">
      <a-col v-for='entry in entries' :key='entry.id' :md='8'>
        <a-card headless class='p-0 echart-card mb-10' :style='customStyle()'>
          <div v-if='entry.icon' class='echart-card-icon'>
            <img :src='`/` + entry.icon' alt=''/>
          </div>
          <div class='card-bar-chart'>
            <h1>{{ entry.view }}</h1>
            <span>Lượt truy cập</span>
          </div>
          <p>
            <span :class="entry.growthType === 'upward' ? 'growth-upward' : 'growth-downward'">
              <idist-feather-icons :type="entry.growthType === 'upward' ? 'arrow-up' : 'arrow-down'"/> {{ entry.growth }}%
            </span>
            <strong>{{ entry.title }}</strong>
          </p>
        </a-card>
      </a-col>
    </a-row>
  </a-card>
</template>

<script>
import {useHomepageStore} from "~/stores/homepage.js";
import IdistFeatherIcons from "~/composables/IdistFeatherIcons.vue";

export default {
  name: "HomepageCategories",
  components: {IdistFeatherIcons},
  data: () => ({
    activeTable: "week",
    loading: true,
    tabs: [
      {
        key: "week",
        name: "Tuần",
      },
      {
        key: "month",
        name: "Tháng",
      },
      {
        key: "year",
        name: "Năm",
      },
    ],
    entries: [
      {
        id: 1,
        view: 12333212,
        title: "Khoa học công nghệ",
        color: '#FFEBEB',
        growth: 25,
        growthType: "upward",
        icon: 'magic-star.svg'
      },
      {
        id: 2,
        view: 17835423,
        title: 'Khoa học công nghệ',
        color: '#FFF9E3',
        growth: 5,
        growthType: 'downward',
        icon: 'heart.svg'
      },
      {
        id: 3,
        view: 45262345,
        title: 'Khoa học công nghệ',
        color: '#EEFDED',
        growth: 5,
        growthType: 'upward',
        icon: 'like.svg'
      },
      {
        id: 4,
        view: 45262345,
        title: 'Khoa học công nghệ',
        color: '#edfdfb',
        growth: 5,
        growthType: 'upward'
      },
      {
        id: 5,
        view: 45262345,
        title: 'Khoa học công nghệ',
        color: '#edfdfb',
        growth: 5,
        growthType: 'upward'
      },
      {
        id: 6,
        view: 45262345,
        title: 'Khoa học công nghệ',
        color: '#edfdfb',
        growth: 5,
        growthType: 'upward'
      },
    ],
  }),
  methods: {
    routerLink(url) {
      this.$router.push(url);
    },
    customStyle() {
      return {
        background: 'linear-gradient(135deg, #ffffff, ' + this.entry.color + ', ' + this.entry.color + ')'
      };
    },
    handleChangeActive(event, activeTable) {
      event.preventDefault();
      this.activeTable = activeTable;
      this.getData();
    },
    async getData() {
      this.loading = true;
      let filter = {
        type: this.activeTable,
      };
      // let response = await this.$store.dispatch("overview_dashboards/GetOverviewRegularPosts", filter);
      let response = await useHomepageStore().GetOverviewRegularPosts(filter);
      this.data = response.data;
      this.loading = false;
    },
    onTabChange(value, type) {
      if (type === "key") {
        key.value = value;
      } else if (type === "noTitleKey") {
        noTitleKey.value = value;
      }
    },
  },
  created() {
  },
  mounted() {
    this.getData();
  },
};
</script>

<style lang="scss"></style>
