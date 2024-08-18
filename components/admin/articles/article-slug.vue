<template>
  <a-form-item prop='slug' name='slug' label='Đường dẫn'>
    <a-input v-model:value='entry.slug' placeholder='Nhập slug...' />
  </a-form-item>
</template>

<script>
export default {
  name: 'ArticleSlug',
  props: {
    value: Object,
    editable: {
      type: Boolean,
      default: () => (true),
    },
  },
  data: () => ({
    entry: {
      title: '',
      slug: ''
    },
    is_changed: false,

  }),
  watch: {
    'entry.slug': {
      handler: function() {
        if (this.entry.slug !== this.slugify(this.value.title)) {
          this.is_changed = true
        }
        if (this.entry.slug === '') {
          this.is_changed = false
          this.entry.slug = this.slugify(this.value.title);
        } else {
          this.entry.slug = this.slugify(this.entry.slug)
        }
        this.$emit('update', this.entry);
      },
      deep: true,
    },
    'value.title': {
      handler: function() {
        if (!this.is_changed || this.entry.slug === '') {
          this.entry.slug = this.slugify(this.value.title);
        }
      },
      deep: true,
    },
  },
  methods: {
    slugify(text) {
      return text.toString().
        toLowerCase().
        replace(/đ/g, 'd').
        replace(/Đ/g, 'D').
        normalize('NFD').
        trim().
        replace(/\s+/g, '-').
        replace(/[^\w-]+/g, '').
        replace(/--+/g, '-')
        ;
    },
  },
  created() {
    this.entry = this.value;
    if (this.value.slug === '') {
      this.entry.slug = this.slugify(this.value.title);
    }
    if (this.entry.slug !== this.slugify(this.value.title) || this.editable === false) {
      this.is_changed = true;
    }  },
  mounted() {
  }
}
</script>

<style scoped>

</style>
