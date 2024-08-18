<template>
  <a-form-item prop='description' name='description' label='Mô tả bài viết'>
    <a-badge class='d-block w-100' :count="entry.description.length">
      <a-textarea :rows='3' v-model:value='entry.description' placeholder='Nhập mô tả bài viết...'/>
    </a-badge>
  </a-form-item>
</template>

<script>
export default {
  name: 'PageDescription',
  props: {
    value: Object,
  },
  data: () => ({
    entry: {
      description: ''
    }
  }),
  methods: {},
  created() {
    this.entry = this.value
  },
  mounted() {
  }
}
</script>

<style scoped>

</style>
