<template>
  <div class="helper-modal">
    <a-popover placement="topRight" trigger="hover" overlayClassName="helper-popover">
      <template #content>
        <a class="d-flex align-items-center" style="margin-bottom: 10px" href="mailto:contact@idist.dev">
          <a-space :size="10">
            <IdistFeatherIcons type="mail" :style="{ fontSize: '20px' }" />
            contact@idist.dev
          </a-space>
        </a>
        <a class="d-flex align-items-center" href="tel:0886914796">
          <a-space :size="10">
            <IdistFeatherIcons type="phone" :style="{ fontSize: '20px' }" />
            088.69.14796
          </a-space>
        </a>
      </template>
      <a-button type="primary" class="helper-cta">
        <IdistFeatherIcons type="help-circle" />
      </a-button>
    </a-popover>
  </div>
</template>

<script>
import IdistFeatherIcons from "~/components/commons/IdistFeatherIcons.vue";

export default {
  name: "AuthHelper",
  components: {IdistFeatherIcons},
  data: () => ({}),
  methods: {},
  created() {},
  mounted() {},
};
</script>

<style lang="scss" scoped>
.helper-modal {
  position: absolute;
  bottom: 20px;
  right: 30px;

  .helper-cta {
    width: 40px;
    height: 40px;
    border-radius: 100px;
    display: flex;
    justify-content: center;
    align-items: center;
  }
}
</style>

<style lang="scss">
.helper-popover {
  .ant-popover-arrow {
    border-right-color: #0050b3 !important;
    border-bottom-color: #0050b3 !important;
  }

  .ant-popover-inner {
    background: #0050b3;

    .ant-popover-inner-content a {
      color: white;
    }
  }
}
</style>
