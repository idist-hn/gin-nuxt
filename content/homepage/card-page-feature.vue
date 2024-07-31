<template>
  <a-card headless class='p-0 echart-card mb-10' :style='customStyle()'>
    <div v-if='entry.icon' class='echart-card-icon'>
      <img :src='`/` + entry.icon' alt='' />
    </div>
    <div class='card-bar-chart'>
      <h1>{{ entry.view }}</h1>
      <span>Lượt truy cập</span>
    </div>
    <p>
      <span :class="entry.growthType === 'upward' ? 'growth-upward' : 'growth-downward'">
        <idist-feather-icons :type="entry.growthType === 'upward' ? 'arrow-up' : 'arrow-down'" /> {{ entry.growth }}%
      </span>
      <strong>{{ entry.title }}</strong>
    </p>
  </a-card>
</template>
<script>
import IdistFeatherIcons from "~/composables/IdistFeatherIcons.vue";

export default {
  name: "CardPageFeature",
  components: {IdistFeatherIcons},
  props: {
    entry: {
      type: Object,
      default: () => ({}),
    },
  },
  methods: {
    customStyle() {
      return {
        background: 'linear-gradient(135deg, #ffffff, ' + this.entry.color + ', ' + this.entry.color + ')'
      };
    },
  },
};
</script>
