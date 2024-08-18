<template>
  <div class="settings">
    <a-popover placement="bottomRight" trigger="click">
      <template v-slot:content>
        <div class="setting-dropdown">
          <a-row :gutter="10">
            <a-col :sm="12">
              <nuxt-link to="#">
                <figure class="setting-dropdown__single d-flex">
                  <img src="~/assets/images/feature.png" alt="" />
                  <figcaption>
                    <h4>All Features</h4>
                    <p>Introducing Increment subscriptions</p>
                  </figcaption>
                </figure>
              </nuxt-link>
            </a-col>
            <a-col :sm="12">
              <nuxt-link to="#">
                <figure class="setting-dropdown__single d-flex">
                  <img src="~/assets/images/theme.png" alt="" />
                  <figcaption>
                    <h4>Themes</h4>
                    <p>Third party themes that are compatible</p>
                  </figcaption>
                </figure>
              </nuxt-link>
            </a-col>
            <a-col :sm="12">
              <nuxt-link to="#">
                <figure class="setting-dropdown__single d-flex">
                  <img src="~/assets/images/payment.png" alt="" />
                  <figcaption>
                    <h4>Payments</h4>
                    <p>We handle billions of dollars</p>
                  </figcaption>
                </figure>
              </nuxt-link>
            </a-col>
            <a-col :sm="12">
              <nuxt-link to="#">
                <figure class="setting-dropdown__single d-flex">
                  <img src="~/assets/images/payment.png" alt="" />
                  <figcaption>
                    <h4>Design Mockups</h4>
                    <p>Share planning visuals with clients</p>
                  </figcaption>
                </figure>
              </nuxt-link>
            </a-col>
          </a-row>
        </div>
      </template>
      <a to="#" class="head-example">
        <IdistFeatherIcons type="settings" :size="20" />
      </a>
    </a-popover>
  </div>
</template>

<script>
import IdistFeatherIcons from "~/composables/IdistFeatherIcons.vue";

export default {
  name: "header-settings",
  components: {IdistFeatherIcons},
  data: () => ({}),
  methods: {},
  created() {},
  mounted() {},
};
</script>

<style scoped></style>
