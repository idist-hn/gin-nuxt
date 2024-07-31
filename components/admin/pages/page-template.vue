<template>
  <a-form-item prop="template_id" name="template_id" label="Định dạng trang">
    <a-select mode="default"
              label-in-value
              :value="value.template ? { label: value.template.name } : undefined"
              placeholder="Chọn mẫu định dạng..."
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
import {useTemplateStore} from "~/stores/templates.js";

export default {
  name: "PageTemplate",
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
      let response = await useTemplateStore().ListEntries(this.pagination);
      this.entries = response.entries;
      this.loading = false;
    },
    handleChange(entry) {
      let self = this;
      this.entries.map((e) => {
        if (entry.key === e.id) {
          self.value.template = e;
          self.value.template_id = e.id;
        }
      });
    },
  },
  created() {
    this.getData();
  },
};
</script>
