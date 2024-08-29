<template>
  <a-card class="department-attributes" title="Định dạng trên PC/Laptop">
    <a-row v-for="(attribute, index) in attributes" :gutter="16" :key="attribute.id" class="mb-10">
      <a-col :md="11">
        <department-style-key :hideTitle="index <= 0" :value="attribute" />
      </a-col>
      <a-col :md="11">
        <department-style-value :hideTitle="index <= 0" :value="attribute" />
      </a-col>
      <a-col :md="2">
        <a-form-item v-if="index <= 0" prop="action" name="action" label=" ">
          <a-button type="danger" @click="deleteAttribute(attribute)">
            <idist-feather-icons type="delete" />
          </a-button>
        </a-form-item>
        <a-button v-else type="danger" @click="deleteAttribute(attribute)">
          <idist-feather-icons type="delete" />
        </a-button>
      </a-col>
    </a-row>
    <a-row class="mt-25">
      <a-button type="success" @click="addAttribute">
        <idist-feather-icons type="plus" />
      </a-button>
    </a-row>
  </a-card>
</template>
<script>

import DepartmentStyleKey from "~/components/admin/departments/department-style-key.vue";
import DepartmentStyleValue from "~/components/admin/departments/department-style-value.vue";
import IdistFeatherIcons from "/components/commons/IdistFeatherIcons.vue";

export default {
  name: "DepartmentPcStyle",
  components: {IdistFeatherIcons, DepartmentStyleValue, DepartmentStyleKey},
  props: {
    value: Object,
  },
  watch: {
    attributes: {
      handler: function () {
        this.value.pc_department_styles = this.attributes;
      },
      deep: true,
    },
  },
  data: () => ({
    attributes: [
      {
        id: 1,
        key: "",
        value: "",
      },
      {
        id: 2,
        key: "",
        value: "",
      },
    ],
  }),
  methods: {
    addAttribute() {
      this.attributes.push({
        id: this.attributes.length + 1,
        key: "",
        value: "",
      });
    },
    deleteAttribute(entry) {
      this.attributes = this.attributes.filter((attribute) => {
        return attribute.id !== entry.id;
      });
    },
  },
  created() {
    this.attributes = this.value.pc_department_styles ?? [];
  },
};
</script>
<style lang="scss">
.department-attributes {
  .ant-form-item {
    margin-bottom: 0;
  }
}
</style>
