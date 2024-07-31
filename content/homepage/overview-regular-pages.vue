<template>
  <a-card title="Top trang truy cập nhiều nhất (Lượt)" style="width: 100%" class="mb-20">
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
      <a-spin />
    </div>

    <a-row :gutter="24" v-else class="overview-page-container">
      <a-col :md="4">
        <a-statistic title="Thông báo tuyển sinh 2023" :value="112893" @click="routerLink('#abc')" />
      </a-col>
      <a-col :md="4">
        <a-statistic title="Thông báo tuyển sinh 2023" :value="112893" @click="routerLink('#abc')" />
      </a-col>
      <a-col :md="4">
        <a-statistic title="Thông báo tuyển sinh 2023" :value="112893" @click="routerLink('#abc')" />
      </a-col>
      <a-col :md="4">
        <a-statistic title="Thông báo tuyển sinh 2023" :value="112893" @click="routerLink('#abc')" />
      </a-col>
      <a-col :md="4">
        <a-statistic title="Thông báo tuyển sinh 2023" :value="112893" @click="routerLink('#abc')" />
      </a-col>
      <a-col :md="4">
        <a-statistic title="Thông báo tuyển sinh 2023" :value="112893" @click="routerLink('#abc')" />
      </a-col>
    </a-row>
  </a-card>
</template>

<script>
export default {
  name: "overviewRegularPages",
  components: {},
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
  }),
  methods: {
    routerLink(url) {
      this.$router.push(url);
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
      let response = await this.$store.dispatch('overview_dashboards/GetOverviewRegularPages', filter)
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
  created() {},
  mounted() {
    this.getData();
  },
};
</script>
<style lang="scss">
.overview-page-container {
  .ant-statistic {
    display: flex;
    flex-direction: column-reverse;
    background: #fff;
    border: 1px solid #d5d8d7;
    box-shadow: 2px 2px 8px 0px rgba(74, 74, 74, 0.05);
    padding: 5px 12px;
    border-radius: 5px;

    .ant-statistic-content-value {
      color: #c50d0d;
      display: block;
    }
  }
}
</style>
