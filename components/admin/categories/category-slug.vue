<template>
  <a-form-item prop='slug' name='slug' label='Đường dẫn'>
    <a-input v-model:value='value.slug' placeholder='Nhập slug...'/>
  </a-form-item>
</template>

<script>
export default {
  name: 'CategorySlug',
  props: {
    value: Object,
    editable: {
      type: Boolean,
      default: () => (true),
    },
  },
  data: () => ({
    is_changed: false,

  }),
  watch: {
    'value.slug': {
      handler: function () {
        if (this.value.slug !== this.slugify(this.value.name)) {
          this.is_changed = true
        }
        if (this.value.slug === '') {
          this.is_changed = false
          this.value.slug = this.slugify(this.value.name);
        } else {
          this.value.slug = this.slugify(this.value.slug)
        }
      },
      deep: true,
    },
    'value.name': {
      handler: function () {
        if (!this.is_changed || this.value.slug === '') {
          this.value.slug = this.slugify(this.value.name);
        }
      },
      deep: true,
    },
  },
  methods: {
    slugify(text) {
      if (text === undefined) return ''
      return text.toString().toLowerCase().replace(/đ/g, 'd').replace(/Đ/g, 'D').normalize('NFD').trim().replace(/\s+/g, '-').replace(/[^\w-]+/g, '').replace(/--+/g, '-');
    },
  },
  mounted() {
    if (this.value.slug === '') {
      this.value.slug = this.slugify(this.value.name);
    }
    if (this.value.slug !== this.slugify(this.value.name) || this.editable === false) {
      this.is_changed = true;
    }
  },
}
</script>

<style scoped>

</style>
