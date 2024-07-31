<template>
  <a-card title="Bài viết được truy cập nhiều nhất" style="width: 100%" class="mb-20">
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

    <a-row :gutter="24" v-else class="overview-most-viewed">
      <a-col :sm="24" class="flex most-viewed-item">
        <div class="item-thumbnail">
          <a-avatar shape="square" :size="70" src="@/assets/images/thmbnail-post.png" alt="" />
        </div>
        <div class="item-title">
          <p class="mb-5">
            Hội thảo “Ứng dụng công nghệ trong quản trị tòa soạn số” lần đầu tiên được Học viện và Hội nhà báo Việt Nam
            phối …
          </p>
          <div>
            <strong class="color-danger">
              <idist-feather-icons type="eye" />
              222.5k
            </strong>
            lượt xem
          </div>
        </div>
      </a-col>
      <a-col :sm="24" class="flex most-viewed-item">
        <div class="item-thumbnail">
          <a-avatar shape="square" :size="70" src="@/assets/images/thmbnail-post.png" alt="" />
        </div>
        <div class="item-title">
          <p class="mb-5">
            Hội thảo “Ứng dụng công nghệ trong quản trị tòa soạn số” lần đầu tiên được Học viện và Hội nhà báo Việt Nam
            phối …
          </p>
          <div>
            <strong class="color-danger">
              <idist-feather-icons type="eye" />
              222.5k
            </strong>
            lượt xem
          </div>
        </div>
      </a-col>
      <a-col :sm="24" class="flex most-viewed-item">
        <div class="item-thumbnail">
          <a-avatar shape="square" :size="70" src="@/assets/images/thmbnail-post.png" alt="" />
        </div>
        <div class="item-title">
          <p class="mb-5">
            Hội thảo “Ứng dụng công nghệ trong quản trị tòa soạn số” lần đầu tiên được Học viện và Hội nhà báo Việt Nam
            phối …
          </p>
          <div>
            <strong class="color-danger">
              <idist-feather-icons type="eye" />
              222.5k
            </strong>
            lượt xem
          </div>
        </div>
      </a-col>
      <a-col :sm="24" class="flex most-viewed-item">
        <div class="item-thumbnail">
          <a-avatar shape="square" :size="70" src="@/assets/images/thmbnail-post.png" alt="" />
        </div>
        <div class="item-title">
          <p class="mb-5">
            Hội thảo “Ứng dụng công nghệ trong quản trị tòa soạn số” lần đầu tiên được Học viện và Hội nhà báo Việt Nam
            phối …
          </p>
          <div>
            <strong class="color-danger">
              <idist-feather-icons type="eye" />
              222.5k
            </strong>
            lượt xem
          </div>
        </div>
      </a-col>
      <a-col :sm="24" class="flex most-viewed-item">
        <div class="item-thumbnail">
          <a-avatar shape="square" :size="70" src="@/assets/images/thmbnail-post.png" alt="" />
        </div>
        <div class="item-title">
          <p class="mb-5">
            Hội thảo “Ứng dụng công nghệ trong quản trị tòa soạn số” lần đầu tiên được Học viện và Hội nhà báo Việt Nam
            phối …
          </p>
          <div>
            <strong class="color-danger">
              <idist-feather-icons type="eye" />
              222.5k
            </strong>
            lượt xem
          </div>
        </div>
      </a-col>
      <a-col :sm="24" class="flex most-viewed-item">
        <div class="item-thumbnail">
          <a-avatar shape="square" :size="70" src="@/assets/images/thmbnail-post.png" alt="" />
        </div>
        <div class="item-title">
          <p class="mb-5">
            Hội thảo “Ứng dụng công nghệ trong quản trị tòa soạn số” lần đầu tiên được Học viện và Hội nhà báo Việt Nam
            phối …
          </p>
          <div>
            <strong class="color-danger">
              <idist-feather-icons type="eye" />
              222.5k
            </strong>
            lượt xem
          </div>
        </div>
      </a-col>
      <a-col :sm="24" class="flex most-viewed-item">
        <div class="item-thumbnail">
          <a-avatar shape="square" :size="70" src="@/assets/images/thmbnail-post.png" alt="" />
        </div>
        <div class="item-title">
          <p class="mb-5">
            Hội thảo “Ứng dụng công nghệ trong quản trị tòa soạn số” lần đầu tiên được Học viện và Hội nhà báo Việt Nam
            phối …
          </p>
          <div>
            <strong class="color-danger">
              <idist-feather-icons type="eye" />
              222.5k
            </strong>
            lượt xem
          </div>
        </div>
      </a-col>
      <a-col :sm="24" class="flex most-viewed-item">
        <div class="item-thumbnail">
          <a-avatar shape="square" :size="70" src="@/assets/images/thmbnail-post.png" alt="" />
        </div>
        <div class="item-title">
          <p class="mb-5">
            Hội thảo “Ứng dụng công nghệ trong quản trị tòa soạn số” lần đầu tiên được Học viện và Hội nhà báo Việt Nam
            phối …
          </p>
          <div>
            <strong class="color-danger">
              <idist-feather-icons type="eye" />
              222.5k
            </strong>
            lượt xem
          </div>
        </div>
      </a-col>
    </a-row>
  </a-card>
</template>

<script>
import IdistFeatherIcons from "~/composables/IdistFeatherIcons.vue";

export default {
  name: "overviewMostViewed",
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
      let response = await this.$store.dispatch("overview_dashboards/GetOverviewRegularPosts", filter);
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
.overview-container {
  .ant-statistic {
    display: flex;
    flex-direction: column-reverse;
    background: #ffd7d7;
    padding: 5px 12px;
    border-radius: 5px;
    box-shadow: 2px 2px 8px 0px rgba(74, 74, 74, 0.05);

    .ant-statistic-content-value {
      color: #c50d0d;
      position: relative;
      display: block;

      &:before {
        cursor: pointer;
        position: absolute;
        right: 0;
        top: 2px;
        content: "→";
        color: #c50d0d;
        font-size: 16px;
        font-weight: bold;
      }
    }
  }
}

.overview-most-viewed {
  max-height: 475px;
  overflow-y: auto;
}

.most-viewed-item {
  display: flex;
  margin-bottom: 10px;

  .item-thumbnail {
    margin-right: 10px;
  }

  .item-title {
    p {
      overflow: hidden;
      display: -webkit-box;
      -webkit-line-clamp: 2; /* number of lines to show */
      line-clamp: 2;
      -webkit-box-orient: vertical;
    }
  }
}
</style>
