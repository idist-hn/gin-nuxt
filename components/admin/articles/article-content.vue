<template>
  <a-form-item prop='content' name='content' label=''>
      <rich-editor  v-model:value="value.content" @update="ChangeContent"/>
  </a-form-item>
</template>

<script>
import RichEditor from "~/composables/RichEditor.vue";

export default {
  name: 'ArticleContent',
  components: {RichEditor},
  props: {
    value: Object
  },
  data: () => ({
    entry: {
      content: ''
    },
  }),
  methods: {
    ChangeContent(content){
      this.value.content = content
      this.$emit("update", this.entry)
    }
  },
}
</script>

<style scoped>

</style>
