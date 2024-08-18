<template>
  <a-form-item prop="published_at" name="published_at" label="Thời gian xuất bản">
    <a-date-picker
        show-time
        :placeholder="time.format(dateFormat).toString()"
        :format="dateFormat"
        @change="changeTime"
    />
  </a-form-item>
</template>

<script>
import moment from "moment";

export default {
  name: "ArticlePublishedAt",
  props: {
    value: Object,
  },
  data: () => ({
    time: null,
    interval: null,
    dateFormat: "HH:mm DD/MM/YYYY",
    has_change: false,
  }),
  methods: {
    setTimeNow() {
      let self = this;
      clearInterval(this.interval)
      if (this.value.published_at != null) {
        this.time = moment(this.value.published_at);
        this.has_change = true;
      } else {
        self.time = moment();
        this.interval = setInterval(function () {
            self.time = moment();
        }, 1000);
      }
    },
    changeTime(value) {
      this.has_change = true;
      this.time = value;
      if (value == null) {
        this.has_change = false;
        this.value.published_at = null;
        this.setTimeNow();
      } else {
        this.value.published_at = value.format(self.dateFormat).toString();
      }
    },
  },
  created() {
    self.time = moment();
    this.setTimeNow();
  },
};
</script>

<style scoped></style>
