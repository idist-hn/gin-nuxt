<template>
  <a-form-item prop="department_id" name="department_id" label="Phòng ban">
    <a-select mode="default"
              show-search
              label-in-value
              :value="value.department ? { label: value.department.name } : undefined"
              placeholder="Chọn phòng ban..."
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
import {useDepartmentStore} from "~/stores/departments.js";

export default {
  name: "UserDepartment",
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
        pageSize: 1000,
        search: "",
      },
    };
  },
  methods: {
    async getData(search = "") {
      this.loading = true;
      this.pagination.search = search;
      let response = await useDepartmentStore().ListEntries(this.pagination);
      this.entries = response.entries;
      this.loading = false;
    },
    handleChange(entry) {
      let self = this;
      this.entries.map((e) => {
        if (entry.key === e.id) {
          self.value.department = e;
          self.value.department_id = e.id;
        }
      });
    },
  },
  created() {
    this.getData();
  },
};
</script>
