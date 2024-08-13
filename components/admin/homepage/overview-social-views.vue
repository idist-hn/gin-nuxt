<template>
  <a-card title='Truy cập các trang mạng xã hội (Lượt)' style='width: 100%' class='mb-20'>
    <template #extra>
      <div class='card-nav'>
        <ul>
          <li v-for='tab in tabs' :key='tab.key' :class="activeTable === tab.key ? 'active' : 'deactivate'">
            <a @click='(e) => handleChangeActive(e, tab.key)' to='#' style='font-weight: bold'> {{ tab.name }}</a>
          </li>
        </ul>
      </div>
    </template>

    <div v-if='loading' class='sd-spin'>
      <a-spin />
    </div>

    <div v-else class='overview-traffic-container'>
      <apexchart type='bar' height='350' :options='chartOptions' :series='series' />
    </div>
  </a-card>
</template>
<script>
export default {
  name: 'OverviewSocialViews',
  components: {},
  data: () => ({
    activeTable: 'week',
    loading: true,
    tabs: [
      {
        key: 'week',
        name: 'Tuần'
      },
      {
        key: 'month',
        name: 'Tháng'
      },
      {
        key: 'year',
        name: 'Năm'
      },
    ],
    chartOptions: {
      colors: ['#FF9797', '#7BB4FF'],
      chart: {
        height: 350,
        type: 'bar',
        zoom: {
          enabled: false
        },
      },
      dataLabels: {
        enabled: false
      },
      stroke: {
        width: 2,
        curve: 'straight'
        // dashArray: [0, 8, 5]
      },
      // title: {
      //   text: 'Page Statistics',
      //   align: 'left'
      // },
      legend: {
        tooltipHoverFormatter: function(val, opts) {
          return val + ' - ' + opts.w.globals.series[opts.seriesIndex][opts.dataPointIndex] + ''
        }
      },
      markers: {
        size: 3,
        hover: {
          sizeOffset: 6
        },
      },
      xaxis: {
        categories: [
          '01 Jan',
          '02 Jan',
          '03 Jan',
          '04 Jan',
          '05 Jan',
          '06 Jan',
          '07 Jan',
          '08 Jan',
          '09 Jan',
          '10 Jan',
          '11 Jan',
          '12 Jan'
        ],
      },
      tooltip: {
        y: [
          {
            title: {
              formatter: function(val) {
                return val
              }
            },
          },
          {
            title: {
              formatter: function(val) {
                return val
              }
            },
          },
          {
            title: {
              formatter: function(val) {
                return val
              }
            },
          },
        ],
      },
      grid: {
        borderColor: '#f1f1f1'
      },
    },
    series: [
      {
        name: 'Youtube',
        data: [45, 52, 38, 24, 33, 26, 21, 20, 6, 8, 15, 10]
      },
      {
        name: 'Facebook',
        data: [35, 41, 62, 42, 13, 18, 29, 37, 36, 51, 32, 35]
      },
    ],
  }),
  methods: {
    routerLink(url) {
      this.$router.push(url)
    },
    handleChangeActive(event, activeTable) {
      event.preventDefault()
      this.activeTable = activeTable
      this.getData()
    },
    async getData() {
      this.loading = true
      let filter = {
        type: this.activeTable
      };
      let response = await this.$store.dispatch('overview_dashboards/GetOverviewTraffics', filter)
      this.data = response.data
      this.loading = false
    },
    onTabChange(value, type) {
      if (type === 'key') {
        key.value = value
      } else if (type === 'noTitleKey') {
        noTitleKey.value = value
      }
    },
  },
  created() {
  },
  mounted() {
    this.getData()
  },
};
</script>
<style scoped></style>
