<template>
  <a-form-item prop="category_id" name="category_id" label="Danh mục">
    <a-select mode="default"
              show-search
              label-in-value
              :value="value.category ? { label: value.category.name } : undefined"
              placeholder="Chọn danh mục..."
              style="width: 100%"
              :filter-option="false"
              :not-found-content="loading ? undefined : null"
              @search="fetchData"
              @change="handleChange">
      <template v-slot:notFoundContent>
        <a-spin v-if="loading" slot="notFoundContent" size="small"/>
      </template>
      <a-select-option v-for="e in entries" :key="e.id">
        {{ e.name }}
      </a-select-option>
    </a-select>
  </a-form-item>
</template>

<script>
import debounce from "lodash/debounce";
import {useCategoryStore} from "~/stores/categories";

export default {
  name: "ArticleCategory",
  props: {
    value: Object,
  },
  data() {
    return {
      entries: [],
      loading: true,
      fetchData: debounce(this.getData, 1000),
      pagination: {
        page: 1,
        length: 12,
        search: "",
      },
    };
  },
  methods: {
    async getData(search = "") {
      this.loading = true;
      this.pagination.search = search;
      let response = await useCategoryStore().ListEntries(this.pagination);
      this.entries = response.entries;
      this.loading = false;
    },
    handleChange(entry) {
      let self = this;
      this.entries.map((e) => {
        if (entry.key === e.id) {
          self.value.category = e;
          self.value.category_id = e.id;
        }
      });
    },
  },
  created() {
    this.getData();
  },
};
</script>
