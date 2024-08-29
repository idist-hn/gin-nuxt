<template>

  <a-row :gutter="16">
    <a-col :md="6">
      <a-card>
        
      </a-card>
    </a-col>
    <a-col :md="6">
      <a-card>
        1
      </a-card>
    </a-col>
    <a-col :md="6">
      <a-card>
        1
      </a-card>
    </a-col>
    <a-col :md="6">
      <a-card>
        1
      </a-card>
    </a-col>
  </a-row>
</template>

<script>
export default {
  name: "index",
  setup() {
    definePageMeta({
      layout: "admin",
    });
  },
}
</script>

<style scoped lang="scss">

</style>