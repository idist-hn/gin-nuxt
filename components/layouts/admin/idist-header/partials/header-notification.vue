<template>
  <div class="notification">
    <a-popover placement="bottomLeft" trigger="click">
      <a-badge dot :offset="[-4, 0]">
        <a to="#" class="head-example">
          <IdistFeatherIcons type="bell" :size="20" />
        </a>
      </a-badge>
      <template v-slot:content>
        <div class="atbd-top-dropdown">
          <h4 class="atbd-top-dropdown__title">
            <span class="title-text">Notifications</span>
            <a-badge class="badge-success" count="3" />
          </h4>
          <perfect-scrollbar
            :options="{
              wheelSpeed: 1,
              swipeEasing: true,
              suppressScrollX: true,
            }"
          >
            <ul class="atbd-top-dropdown__nav notification-list">
              <li>
                <a to="#">
                  <div class="atbd-top-dropdown__content notifications">
                    <div class="notification-icon bg-primary">
                      <IdistFeatherIcons type="hard-drive" />
                    </div>
                    <div class="notification-content d-flex">
                      <div class="notification-text">
                        <h4><span>James</span> sent you a message</h4>
                        <p>5 hours ago</p>
                      </div>
                      <div class="notification-status">
                        <a-badge dot />
                      </div>
                    </div>
                  </div>
                </a>
              </li>
              <li>
                <a to="#">
                  <div class="atbd-top-dropdown__content notifications">
                    <div class="notification-icon bg-secondary">
                      <IdistFeatherIcons type="share-2" />
                    </div>
                    <div class="notification-content d-flex">
                      <div class="notification-text">
                        <h4><span>James</span> sent you a message</h4>
                        <p>5 hours ago</p>
                      </div>

                      <div class="notification-status">
                        <a-badge dot />
                      </div>
                    </div>
                  </div>
                </a>
              </li>
              <li>
                <a to="#">
                  <div class="atbd-top-dropdown__content notifications">
                    <div class="notification-icon bg-secondary">
                      <IdistFeatherIcons type="share-2" />
                    </div>
                    <div class="notification-content d-flex">
                      <div class="notification-text">
                        <h4><span>James</span> sent you a message</h4>
                        <p>5 hours ago</p>
                      </div>

                      <div class="notification-status">
                        <a-badge dot />
                      </div>
                    </div>
                  </div>
                </a>
              </li>
              <li>
                <a to="#">
                  <div class="atbd-top-dropdown__content notifications">
                    <div class="notification-icon bg-secondary">
                      <IdistFeatherIcons type="share-2" />
                    </div>
                    <div class="notification-content d-flex">
                      <div class="notification-text">
                        <h4><span>James</span> sent you a message</h4>
                        <p>5 hours ago</p>
                      </div>

                      <div class="notification-status">
                        <a-badge dot />
                      </div>
                    </div>
                  </div>
                </a>
              </li>
              <li>
                <a to="#">
                  <div class="atbd-top-dropdown__content notifications">
                    <div class="notification-icon bg-secondary">
                      <IdistFeatherIcons type="share-2" />
                    </div>
                    <div class="notification-content d-flex">
                      <div class="notification-text">
                        <h4><span>James</span> sent you a message</h4>
                        <p>5 hours ago</p>
                      </div>

                      <div class="notification-status">
                        <a-badge dot />
                      </div>
                    </div>
                  </div>
                </a>
              </li>
            </ul>
          </perfect-scrollbar>
          <router-link class="btn-seeAll" to="#"> See all incoming activity </router-link>
        </div>
      </template>
    </a-popover>
  </div>
</template>

<script>
import { PerfectScrollbar } from "vue3-perfect-scrollbar";
import IdistFeatherIcons from "~/composables/IdistFeatherIcons.vue";

export default {
  name: "header-notification",
  components: {IdistFeatherIcons, PerfectScrollbar },
  data: () => ({}),
  methods: {},
  created() {},
  mounted() {},
};
</script>

<style scoped></style>
