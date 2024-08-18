<template>
  <a-form-item prop='description' name='description' label='Mô tả'>
    <a-badge class='d-block w-100' :count="value.description.length">
      <a-textarea :rows='3' v-model:value='value.description' placeholder='Nhập mô tả danh mục...'/>
    </a-badge>
  </a-form-item>
</template>

<script>
export default {
  name: 'CategoryDescription',
  props: {
    value: Object,
  },

}
</script>

<style scoped>

</style>
