<template>
  <a-card title="Danh mục con" body-style="padding:0 !important">
    <a-list v-if="value.children.length > 0">
      <a-list-item v-for="item in value.children" :key="item.id">
        <a-alert type="info" class="w-100">
          <template #message>
            <a target="_blank" :href="publicLinkCategory(item)">{{ item.name }}</a>
          </template>
        </a-alert>
      </a-list-item>
    </a-list>
    <a-alert v-else type="warning" class="w-100 font-italic">
      <template #message>
        Chưa có danh mục con!
      </template>
    </a-alert>
  </a-card>
</template>

<script>
export default {
  name: "CategoryChildren",
  props: {
    value: Object
  },
  methods: {
    publicLinkCategory(entry) {
      return `${this.$config.public.portalDomain}/danh-muc/${entry.slug}`;
    },
  },
  mounted() {
  }
}
</script>

<style scoped lang="scss">

</style>