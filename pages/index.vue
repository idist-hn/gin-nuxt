<template>
    <a-row :gutter="16">

    </a-row>
</template>
<script>
import IdistFeatherIcons from "/components/commons/IdistFeatherIcons.vue";

export default {
  name: 'index',
  layout: 'default',
  components: {
    IdistFeatherIcons
  },
  mounted() {
    this.$router.push({name:'admin-homepage'})
  }
}
</script>