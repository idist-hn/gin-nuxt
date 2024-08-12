<template>
  <a-card more title="Hiệu suất của trang web" class="mb-0">
    <template #extra>
      <div class="card-nav">
        <ul>
          <li v-for="tab in tabs" :key="tab.key" :class="activeTable === tab.key ? 'active' : 'deactivate'">
            <a @click="(e) => handleChangeActive(e, tab.key)" to="#" style="font-weight: bold"> {{ tab.name }}</a>
          </li>
        </ul>
      </div>
    </template>

    <div v-if='!loading' class='pstates'>
      <div :class="`growth-upward active`" role="button">
        <p>Truy cập trang</p>
        <h1>
          72.6K
          <sub>
            <span> <idist-feather-icons type="arrow-up" size="14" /> 25% </span>
          </sub>
        </h1>
      </div>
      <div :class="`growth-upward`" role="button">
        <p>Truy cập bài viết</p>
        <h1>
          23.5K
          <sub>
            <span> <idist-feather-icons type="arrow-up" size="14" /> 47% </span>
          </sub>
        </h1>
      </div>
      <div :class="`growth-downward`" role="button">
        <p>Click liên kết</p>
        <h1>
          1123
          <sub>
            <span> <idist-feather-icons type="arrow-down" size="14" /> 28% </span>
          </sub>
        </h1>
      </div>
    </div>

    <div v-if="loading" class="sd-spin">
      <a-spin />
    </div>
    <div v-else class="performance-lineChart">
      <apexchart type="area" height="350" :options="chartOptions" :series="series" />
      <ul>
        <li v-for="(item, index) in performanceDatasets" :key="index + 1" class="custom-label">
          <span
            :style="{
              backgroundColor: item.borderColor,
            }"
          />
          {{ item.label }}
        </li>
      </ul>
    </div>
  </a-card>
</template>
<script>
import IdistFeatherIcons from "~/components/commons/IdistFeatherIcons.vue";

export default {
  name: "OverviewPerformance",
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
    chartOptions: {
      colors: ["#C50D0D", "#e5e9f2"],
      maintainAspectRatio: true,
      chart: {
        height: 350,
        type: "area",
        zoom: {
          enabled: false,
        },
      },
      elements: {
        z: 9999,
      },
      legend: {
        display: false,
      },
      hover: {
        mode: "index",
        intersect: false,
      },
      tooltips: {
        mode: "label",
        intersect: false,
        backgroundColor: "#ffffff",
        position: "average",
        enabled: false,
        // custom: customTooltips,
        callbacks: {
          title() {
            return "users";
          },
          label(t, d) {
            const { yLabel, datasetIndex } = t;
            return `<span class='chart-data'>${yLabel}k</span> <span class='data-label'>${d.datasets[datasetIndex].label}</span>`;
          },
        },
      },
      scales: {
        yAxes: [
          {
            gridLines: {
              color: "#e5e9f2",
              borderDash: [3, 3],
              zeroLineColor: "#e5e9f2",
              zeroLineWidth: 1,
              zeroLineBorderDash: [3, 3],
            },
            ticks: {
              beginAtZero: true,
              fontSize: 13,
              fontColor: "#182b49",
              max: 80,
              stepSize: 20,
              callback(label) {
                return `${label}k`;
              },
            },
          },
        ],
        xAxes: [
          {
            gridLines: {
              display: true,
              zeroLineWidth: 2,
              zeroLineColor: "transparent",
              color: "transparent",
              z: 1,
              tickMarkLength: 0,
            },
            ticks: {
              padding: 10,
            },
          },
        ],
      },
    },
    series: [
      {
        name: "Thời điểm hiện tai",
        data: [31, 40, 28, 51, 42, 109, 100],
      },
      {
        name: "Thời điểm trước đó",
        data: [11, 32, 45, 32, 34, 52, 41],
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
      let response = await this.$store.dispatch("overview_dashboards/GetOverviewTraffics", filter);
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
  mounted() {
    this.getData();
  },
};
</script>
<style lang="scss">
.pstates {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  margin: -24px 0 25px;
  @media only screen and (max-width: 767px) {
    margin: -19px 0 25px;
    flex-flow: column;
  }

  > div {
    transition: 0.3s ease;
    padding: 20px;
    @media only screen and (max-width: 1599px) {
      flex: 0 0 50%;
    }

    p {
      font-weight: 500;
    }

    &:hover {
      box-shadow: 0 15px 30px rgba(146, 153, 184, 0.15);

      p {
        color: red;
      }
    }

    &.active {
      background: #f8f9fb;

      &:hover {
        box-shadow: 0 15px 30px #fff;
      }
    }
  }

  .growth-upward,
  .growth-downward {
    cursor: pointer;

    &:focus {
      outline: none;
    }

    h1 {
      font-size: 24px;

      sub {
        span {
          font-weight: 500;
        }
      }
    }
  }
}
</style>
