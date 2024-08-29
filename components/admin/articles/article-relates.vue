<template>
  <a-form-item prop='related_articles' name='related_articles' label='Bài viết liên quan'>
    <a-select
        mode='multiple'
        show-search
        label-in-value
        v-model:value='value.related_articles'
        placeholder='Chọn bài viết...'
        style='width: 100%'
        :filter-option='true'
        :not-found-content='loading ? undefined : null'
        @search='fetchData'
        @change="handleChange"
    >
      <template v-slot:notFoundContent>
        <a-spin size='small'/>
      </template>
      <!--      <template v-slot:tagRender #tagRender="{ value: val, label, closable, onClose, option }">-->
      <!--        <a-tag :closable="closable" style="margin-right: 3px" @close="onClose">-->
      <!--          demo - {{ val.title }}&nbsp;-->
      <!--          <span role="img" :aria-label="val">{{ option.icon }}</span>-->
      <!--        </a-tag>-->
      <!--      </template>-->
      <!--      <template v-slot:option #option="{ value: val, label, icon }">-->
      <!--        <span role="img" :aria-label="val"> option - {{ icon }}</span>-->
      <!--        &nbsp;&nbsp;{{ label }}-->
      <!--      </template>-->
      <a-select-option v-for='entry in entries' :key='entry.id' :id="entry.id">
        {{ entry.title }}
      </a-select-option>
    </a-select>
  </a-form-item>
</template>

<script>
import debounce from 'lodash/debounce'
import {useArticleStore} from '~/stores/articles'

export default {
  name: 'ArticleRelates',
  props: {
    value: Object
  },
  data() {
    return {
      entries: [],
      loading: true,
      fetchData: debounce(this.getData, 1000),
      pagination: {
        page: 1,
        length: 12,
        search: ''
      }
    }
  },
  methods: {
    async getData(key = '') {
      this.loading = true
      this.pagination.search = key
      let response = await useArticleStore().ListRelateArticles(this.pagination)
      this.entries = response.entries
      this.loading = false
    },
    handleChange(entries) {
      this.value.related_article_ids = []
      entries.map(e => {
        this.value.related_article_ids.push(e.key)
      })
    },
    handleInitData() {
      if (this.value.related_articles == null) this.value.related_articles = []
      this.value.related_articles = this.value.related_articles.map(entry => {
        return {
          key: entry.id ?? entry.key,
          label: entry.title ?? entry.labels
        }
      })
    }
  },
  created() {
    this.handleInitData()
    this.getData()

  },
  mounted() {
  }
}
</script>

<style scoped>

</style>
