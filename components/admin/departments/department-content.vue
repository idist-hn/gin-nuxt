<template>
  <a-card title="Nội dung">
    <rich-editor v-model:value="value.content" @update="ChangeContent"/>
  </a-card>
</template>

<script>

import RichEditor from "~/components/commons/RichEditor.vue";

export default {
  name: "DepartmentContent",
  components: {RichEditor},
  props: {
    value: Object,
  },
  data: () => ({
    entry: {
      content: ''
    },
  }),
  methods: {
    ChangeContent(content) {
      this.entry.content = content
      this.$emit("update", this.entry)
    }
  },
  created() {
    this.entry = this.value
  }
}
</script>
<style scoped>

</style>
