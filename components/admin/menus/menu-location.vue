<template>
  <a-form-item prop="parent" name="parent" label="Danh mục cha">
    <a-select
        mode='default'
        show-search
        label-in-value
        :value='value.location ? {label:value.location.name }: undefined'
        placeholder='Chọn vị trí...'
        style='width: 100%'
        :filter-option='false'
        :not-found-content='loading ? undefined : null'
        @search='fetchData'
        :allowClear="true"
        @change='handleChange'
    >
      <a-select-option v-for='entry in entries' :key='entry.id'>
        {{ entry.name }}
      </a-select-option>
    </a-select>
  </a-form-item>
</template>

<script>
import debounce from 'lodash/debounce'
import {useMenuStore} from "~/stores/menus.js";

export default {
  name: 'MenuLocation',
  props: {
    value: Object
  },
  data() {
    return {
      entries: [],
      entry: null,
      loading: true,
      fetchData: debounce(this.getData, 1000),
      pagination: {
        page: 1,
        length: 120,
        search: ''
      }
    }
  },
  methods: {
    async getData() {
      this.loading = true
      let response = await useMenuStore().ListMenuLocations(this.pagination)
      this.entries = response.entries
      this.loading = false
    },
    handleChange(entry) {
      if (entry == null) {
        this.value.parent = null
        return
      }
      this.entries.map(e => {
        if (entry.key === e.id) {
          this.value.parent = e
        }
      })
    }
  },
  mounted() {
    this.getData()
  }
}
</script>

<style scoped>

</style>
