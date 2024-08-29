<template>
  <a-form-item prop="deparment_parent" name="deparment_parent" label="Phòng ban cấp trên">
    <a-select
        mode="default"
        show-search
        label-in-value
        :value="value.parent ? { label: value.parent.name } : undefined"
        placeholder="Chọn danh mục cha..."
        style="width: 100%"
        :filter-option="false"
        :not-found-content="loading ? undefined : null"
        @search="fetchData"
        @change="handleChange"
    >
      <a-spin v-if="loading" slot="notFoundContent" size="small"/>
      <a-select-option v-for="entry in entries" :key="entry.id">
        {{ entry.name }}
      </a-select-option>
    </a-select>
  </a-form-item>
</template>
<script>
import debounce from "lodash/debounce";
import {useDepartmentStore} from "~/stores/departments.js";

export default {
  name: "DepartmentParent",
  props: {
    value: Object,
  },
  data() {
    return {
      entries: [],
      entry: null,
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
    async getData(search) {
      this.loading = true;
      this.pagination.search = search;
      let response = await useDepartmentStore().ListEntries(this.pagination);
      this.entries = response.entries;
      this.loading = false;
    },
    handleChange(selected) {
      this.entries.map((e) => {
        if (selected.key === e.id) {
          this.value.parent = e;
          this.value.parent_id = e.id;
        }
      });
    },
  },
  mounted() {
    this.getData();
  },
};
</script>
<style scoped></style>
