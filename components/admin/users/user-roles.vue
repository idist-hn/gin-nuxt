<template>
  <a-form-item prop="roles" name="roles" label="Nhóm quyền">
    <a-select
        mode="multiple"
        show-search
        label-in-value
        v-model:value="value.roles"
        placeholder="Chọn nhóm quyền..."
        style="width: 100%"
        :filter-option="false"
        :not-found-content="loading ? undefined : null"
        @search="fetchData"
        @change="handleChange"
    >
      <template v-slot:notFoundContent>
        <a-spin v-if="loading" slot="notFoundContent" size="small"/>
      </template>
      <a-select-option v-for="e in entries" :key="e.key">
        {{ e.title }}
      </a-select-option>
    </a-select>
  </a-form-item>
</template>

<script>
import debounce from "lodash/debounce";
import {useRolesStore} from '~/stores/roles'

export default {
  name: "UserRoles",
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
    async getData(key = "") {
      this.loading = true;
      this.pagination.search = key;
      let response = await useRolesStore().ListEntries(this.pagination);
      this.entries = response.entries.map((e) => {
        return {
          key: e.id,
          title: e.name,
        };
      });
      this.loading = false;
    },
    handleChange(entries) {
      this.value.role_ids = [];
      this.value.roles = entries;
      entries.map((e) => {
        this.value.role_ids.push(e.key);
      });
    },

    handleInitData() {
      if (this.value.roles == null) this.value.roles = []
      this.value.roles = this.value.roles.map(entry => {
        return {
          key: entry.id ?? entry.key,
          label: entry.title ?? entry.labels
        }
      })
    }
  },
  created() {
    this.entry = this.value;
    this.getData();
  },
  mounted() {
    this.handleInitData()
  },
};
</script>

<style scoped></style>
