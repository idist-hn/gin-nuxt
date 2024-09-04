<template>
      <a-row>
        <a-col :md="12" :sm="24">
          <a-card title="Card title">
            <p>Card content</p>
            <p>Card content</p>
            <p>Card content</p>
          </a-card>
        </a-col>
      </a-row>
</template>
<script>

export default {
  name: 'index',
  setup() {
    definePageMeta({
      layout: "common",
    })
  },
  components: {
  },
  mounted() {
  }
}
</script>