<template>
  <a-form-item prop="tags" name="tags" label="Từ khoá liên quan">
    <a-select
        mode="multiple"
        show-search
        label-in-value
        v-model:value="value.tags"
        placeholder="Chọn từ khoá..."
        style="width: 100%"
        :filter-option="false"
        :not-found-content="loading ? undefined : null"
        @search="fetchUser"
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
import {useTagStore} from '~/stores/tags'

export default {
  name: "ArticleTags",
  props: {
    value: Object,
  },
  data() {
    return {
      entries: [],
      loading: true,
      fetchUser: debounce(this.getData, 1000),
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
      let response = await useTagStore().ListEntries(this.pagination);
      this.entries = response.entries.map((e) => {
        return {
          key: e.id,
          title: e.name,
        };
      });
      this.loading = false;
    },
    handleChange(entries) {
      this.value.tag_ids = [];
      this.value.tags = entries;
      entries.map((e) => {
        this.value.tag_ids.push(e.key);
      });
    },

    handleInitData() {
      if (this.value.tags == null) this.value.tags = []
      this.value.tags = this.value.tags.map(entry => {
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
